package main

import (
	"context"
	"flag"

	populatorMachinery "github.com/kubernetes-csi/lib-volume-populator/v3/populator-machinery"
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

	// Example steps:
	// 1. Retrieve data source details from the defined CRD.
	// 2. Initiate a data transfer job to params.PvcPrime using params.KubeClient.
	// 3. Report the volume population status to the original PVC's through params.Recorder.
	// 4. You should check if the transfer job already exists before creating it, otherwise the transfer job might
	// get created multiple times if everytime you use a unique name.

	klog.Infof("Run populateFn")
	return nil
}

func populateCompleteFn(ctx context.Context, params populatorMachinery.PopulatorParams) (bool, error) {
	// Implement the provider-specific logic to determine the status of volume population.
	// This may involve calling cloud-native APIs or checking the completion status of
	// temporary Kubernetes resources like Pods or Jobs.

	// Example steps:
	// 1. Fetch the transfer job using params.KubeClient.
	// 2. Verify if the job has finished successfully (returns true) or is still running (returns false).
	// 3. If the transfer job encountered an error, evaluate the need for cleanup.
	// 4. Report the volume population status to the original PVC through params.Recorder.

	klog.Infof("Run populateCompleteFn")
	return true, nil
}

func populateCleanupFn(ctx context.Context, params populatorMachinery.PopulatorParams) error {
	// Implement the provider-specific logic to clean up any temporary resources
	// that were created during the volume population process. This step happens after PV rebind to the original PVC
	// and before the PVC' gets deleted.

	// Example steps:
	// 1. Fetch the transfer job using params.KubeClient.
	// 2. If the transfer job still exists delete the job.
	// 3. Report the volume population status to the original PVC through params.Recorder.

	klog.Infof("Run populateCleanupFn")
	return nil
}
