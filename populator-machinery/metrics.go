/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package populator_machinery

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"

	k8smetrics "k8s.io/component-base/metrics"
)

const (
	subSystem   = "volume_populator"
	labelResult = "result"
)

type metricsManager struct {
	mu               sync.Mutex
	srv              *http.Server
	cache            map[types.UID]time.Time
	registry         k8smetrics.KubeRegistry
	opLatencyMetrics *k8smetrics.HistogramVec
	opInFlight       *k8smetrics.Gauge
}

var metricBuckets = []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10, 15, 30, 60, 120, 300, 600}
var inFlightCheckInterval = 30 * time.Second

func initMetrics() *metricsManager {

	m := new(metricsManager)
	m.cache = make(map[types.UID]time.Time)
	m.registry = k8smetrics.NewKubeRegistry()

	m.opLatencyMetrics = k8smetrics.NewHistogramVec(
		&k8smetrics.HistogramOpts{
			Subsystem: subSystem,
			Name:      "operation_seconds",
			Help:      "Time taken by each populator operation",
			Buckets:   metricBuckets,
		},
		[]string{labelResult},
	)
	m.opInFlight = k8smetrics.NewGauge(
		&k8smetrics.GaugeOpts{
			Subsystem: subSystem,
			Name:      "operations_in_flight",
			Help:      "Total number of operations in flight",
		},
	)

	k8smetrics.RegisterProcessStartTime(m.registry.Register)
	m.registry.MustRegister(m.opLatencyMetrics)
	m.registry.MustRegister(m.opInFlight)

	go m.scheduleOpsInFlightMetric()

	return m
}

func (m *metricsManager) scheduleOpsInFlightMetric() {
	for range time.Tick(inFlightCheckInterval) {
		func() {
			m.mu.Lock()
			defer m.mu.Unlock()
			m.opInFlight.Set(float64(len(m.cache)))
		}()
	}
}

type promklog struct{}

func (pl promklog) Println(v ...interface{}) {
	klog.Error(v...)
}

func (m *metricsManager) startListener(httpEndpoint, metricsPath string) {
	if "" == httpEndpoint || "" == metricsPath {
		return
	}

	mux := http.NewServeMux()
	mux.Handle(metricsPath, k8smetrics.HandlerFor(
		m.registry,
		k8smetrics.HandlerOpts{
			ErrorLog:      promklog{},
			ErrorHandling: k8smetrics.ContinueOnError,
		}))

	klog.Infof("Metrics path successfully registered at %s", metricsPath)

	l, err := net.Listen("tcp", httpEndpoint)
	if err != nil {
		klog.Fatalf("failed to listen on address[%s], error[%v]", httpEndpoint, err)
	}
	m.srv = &http.Server{Addr: l.Addr().String(), Handler: mux}
	go func() {
		if err := m.srv.Serve(l); err != http.ErrServerClosed {
			klog.Fatalf("failed to start endpoint at:%s/%s, error: %v", httpEndpoint, metricsPath, err)
		}
	}()
	klog.Infof("Metrics http server successfully started on %s, %s", httpEndpoint, metricsPath)
}

func (m *metricsManager) stopListener() {
	if m.srv == nil {
		return
	}

	err := m.srv.Shutdown(context.Background())
	if err != nil {
		klog.Errorf("Failed to shutdown metrics server: %s", err.Error())
	}

	klog.Infof("Metrics server successfully shutdown")
}

// operationStart starts a new operation
func (m *metricsManager) operationStart(pvcUID types.UID) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.cache[pvcUID]; !exists {
		m.cache[pvcUID] = time.Now()
	}
	m.opInFlight.Set(float64(len(m.cache)))
}

// dropOperation drops an operation
func (m *metricsManager) dropOperation(pvcUID types.UID) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.cache, pvcUID)
	m.opInFlight.Set(float64(len(m.cache)))
}

// recordMetrics emits operation metrics
func (m *metricsManager) recordMetrics(pvcUID types.UID, result string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	startTime, exists := m.cache[pvcUID]
	if !exists {
		// the operation has not been cached, return directly
		return
	}

	operationDuration := time.Since(startTime).Seconds()
	m.opLatencyMetrics.WithLabelValues(result).Observe(operationDuration)

	delete(m.cache, pvcUID)
	m.opInFlight.Set(float64(len(m.cache)))
}

// prometheusMetricsManager is designed to manage Prometheus metrics related to population operations.
type prometheusMetricsManager struct {
	// opErrorCount is a Prometheus counter to track the number of errors encountered during operations
	opErrorCount *prometheus.CounterVec
	// opRequestCount is a Prometheus counter to track the total number of operation requests received.
	opRequestCount *prometheus.CounterVec
	// opSuccessCount is a Prometheus counter to track the total number of successful operations.
	opSuccessCount *prometheus.CounterVec
}

// initPrometheusMetrics initializes a `prometheusMetricsManager` struct for tracking
// metrics related to "populator operations" using Prometheus.
func initPrometheusMetrics() *prometheusMetricsManager {
	pm := new(prometheusMetricsManager)

	pm.opErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "volume_population_error_count",
			Help: "Number of failed volume populator operations",
		},
		[]string{"method", "pvc_uid", "error_code"},
	)

	pm.opRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "volume_population_count",
			Help: "Number of total volume populator operations requests",
		},
		[]string{"pvc_uid"},
	)

	pm.opSuccessCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "volume_population_success_count",
			Help: "Number of successful volume populator operations",
		},
		[]string{"pvc_uid"},
	)

	prometheus.MustRegister(pm.opErrorCount)
	prometheus.MustRegister(pm.opRequestCount)
	prometheus.MustRegister(pm.opSuccessCount)

	return pm
}

// recordError is responsible for recording errors that occur during operations and updating the corresponding Prometheus metric.
func (pm *prometheusMetricsManager) recordError(method string, pvcUID types.UID, err error) {
	internalErr, _ := status.FromError(err)
	code := internalErr.Code().String()
	pm.opErrorCount.WithLabelValues(method, string(pvcUID), code).Inc()
}

// recordRequest is responsible for recording requests that occur during operations and updating the corresponding Prometheus metric.
func (pm *prometheusMetricsManager) recordRequest(pvcUID types.UID) {
	pm.opRequestCount.WithLabelValues(string(pvcUID)).Inc()
}

// recordSuccess is responsible for recording successes that occur during operations and updating the corresponding Prometheus metric.
func (pm *prometheusMetricsManager) recordSuccess(pvcUID types.UID) {
	pm.opSuccessCount.WithLabelValues(string(pvcUID)).Inc()
}

// startListener initiates an HTTP server to expose Prometheus metrics.
func (pm *prometheusMetricsManager) startListener(httpEndpoint, metricsPath string) {
	if "" == httpEndpoint || "" == metricsPath {
		return
	}

	http.Handle(metricsPath, promhttp.Handler())
	go http.ListenAndServe(httpEndpoint, nil)

	klog.Infof("Prometheus metrics http server successfully started on %s, %s", httpEndpoint, metricsPath)
}
