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

	"k8s.io/apimachinery/pkg/types"
	k8smetrics "k8s.io/component-base/metrics"
	"k8s.io/klog/v2"
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

// ProviderMetric is an empty interface, serving as a placeholder or marker interface.
// It can be used to group types that represent metrics, even if they don't share any common methods.
type ProviderMetric interface{}

// VolumePopulationMetric represents a metric for tracking volume population operations.
type VolumePopulationMetric struct {
	// Method is the name of the method that was called and tracked by this metric.
	// It indicates which specific operation was performed.
	// Known methods:
	// - controller.syncPvc: This covers the core business logic such as:
	//   - The creation of PVCs, data source custom resources, and populator pods.
	//   - Volume population routines and provider specific populator functions.
	//   - Binding the PV on completion.
	//   - The deletion of temporary PVCs and populator pods.
	Method string
	// Error is the error that occurred during the volume population operation.
	// A nil value in this field indicates a successful operation.
	Error error
}

// ProviderMetricManager holds the configuration for handling provider specific-metric handling.
//
// The `ProviderMetric` parameter is an interface, allowing the provider to use any data type as a metric.
// It's essential for the provider to implement the necessary type handling within the `HandleMetricFn`
// to process each metric type appropriately and decide what actions to take with the data.
//
// In the provider's implementation, the `HandleMetricFn` function would likely send the metric data
// to an external system for collection and analysis. This could involve:
// - Serializing the metric data into a suitable format (e.g., JSON, Protobuf).
// - Sending the data over a network to a metrics collection service or database.
// - Potentially handling errors or retries in case of network issues or service unavailability.
type ProviderMetricManager struct {
	HandleMetricFn func(ProviderMetric) error
}

// handleVolumePopulationMetric handlles a metric indicating that a volume population operation was performed.
func (pm *ProviderMetricManager) handleVolumePopulationMetric(method string, err error) {
	if err := pm.HandleMetricFn(&VolumePopulationMetric{
		Method: method,
		Error:  err,
	}); err != nil {
		klog.Errorf("Failed to handle volume population metric: %+v", err)
	}
}
