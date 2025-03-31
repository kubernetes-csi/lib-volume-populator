package main

import (
	"context"
	"flag"

	populatorMachinery "github.com/kubernetes-csi/lib-volume-populator/populator-machinery"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/klog/v2"
)

const (
	groupName  = "provider.example.com"
	kind       = "Provider"
	apiVersion = "v1alpha1"
	resource   = "providers"
	prefix     = "provider.example.com"
)

var (
	namespace      = flag.String("namespace", "provider", "Namespace to deploy controller")
	mode           = flag.String("mode", "", "Mode to run the application in, supported values is controller")
	kubeconfigPath = flag.String("kubeconfig-path", "", "kubeconfig path")
	httpEndpoint   = flag.String("http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string.")
	metricsPath    = flag.String("metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")

	groupKind = schema.GroupKind{
		Group: groupName,
		Kind:  kind,
	}
	versionResource = schema.GroupVersionResource{
		Group:    groupKind.Group,
		Version:  apiVersion,
		Resource: resource,
	}
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	switch *mode {
	case "controller":
		klog.Infof("Run provider-populator controller")
		pfcfg := &populatorMachinery.ProviderFunctionConfig{
			PopulateFn:         populateFn,
			PopulateCompleteFn: populateCompleteFn,
			PopulateCleanupFn:  populateCleanupFn,
		}

		vpcfg := &populatorMachinery.VolumePopulatorConfig{
			MasterURL:              "",
			Kubeconfig:             *kubeconfigPath,
			HttpEndpoint:           *httpEndpoint,
			MetricsPath:            *metricsPath,
			Namespace:              *namespace,
			Prefix:                 prefix,
			Gk:                     groupKind,
			Gvr:                    versionResource,
			ProviderFunctionConfig: pfcfg,
			CrossNamespace:         false,
		}

		populatorMachinery.RunControllerWithConfig(*vpcfg)
	default:
		klog.Infof("Mode %s is not supported", *mode)
	}
}

func populateFn(ctx context.Context, params populatorMachinery.PopulatorParams) error {
	// Implement the provider-specific logic to initiate volume population.
	// This may involve calling cloud-native APIs or creating temporary Kubernetes resources
	// such as Pods or Jobs for data transfer.
	klog.Infof("Run populateFn")
	return nil
}

func populateCompleteFn(ctx context.Context, params populatorMachinery.PopulatorParams) (bool, error) {
	// Implement the provider-specific logic to determine the status of volume population.
	// This may involve calling cloud-native APIs or checking the completion status of
	// temporary Kubernetes resources like Pods or Jobs.
	klog.Infof("Run populateCompleteFn")
	return true, nil
}

func populateCleanupFn(ctx context.Context, params populatorMachinery.PopulatorParams) error {
	// Implement the provider-specific logic to clean up any temporary resources
	// that were created during the volume population process.
	klog.Infof("Run populateCleanupFn")
	return nil
}
