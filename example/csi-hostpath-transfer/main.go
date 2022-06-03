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

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	snapclientset "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gatewayclientset "sigs.k8s.io/gateway-api/pkg/client/clientset/gateway/versioned"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	utilexec "k8s.io/utils/exec"

	populator_machinery "github.com/kubernetes-csi/lib-volume-populator/populator-machinery"
)

const (
	prefix           = "snapshot.storage.k8s.io"
	mountPath        = "/mnt"
	devicePath       = "/dev/block"
	stateDir         = "/csi-data-dir"
	snapshotExt      = ".snap"
	driverName       = "hostpath.csi.k8s.io"
	snapshotKind     = "VolumeSnapshot"
	snapshotAPIGroup = "snapshot.storage.k8s.io"
	snapshotLinkKind = "VolumeSnapshotLink"
)

var version = "unknown"

func main() {
	var (
		mode               string
		volumeMode         string
		sourceNamespace    string
		destNamespace      string
		sourceSnapshotName string
		httpEndpoint       string
		metricsPath        string
		masterURL          string
		kubeconfig         string
		imageName          string
		showVersion        bool
		namespace          string
	)
	// Main arg
	flag.StringVar(&mode, "mode", "", "Mode to run in (controller, populate)")
	// Populate args
	flag.StringVar(&volumeMode, "volume-mode", "", "volume mode to populate")
	flag.StringVar(&sourceNamespace, "source-namespace", "", "source namespace")
	flag.StringVar(&destNamespace, "dest-namespace", "", "dest namespace")
	flag.StringVar(&sourceSnapshotName, "source-snapshot-name", "", "source snapshot name")
	// Controller args
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&imageName, "image-name", "", "Image to use for populating")
	// Metrics args
	flag.StringVar(&httpEndpoint, "http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string, which means the server is disabled.")
	flag.StringVar(&metricsPath, "metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")
	// Other args
	flag.BoolVar(&showVersion, "version", false, "display the version string")
	flag.StringVar(&namespace, "namespace", "csi-hostpath-transfer", "Namespace to deploy controller")
	flag.Parse()

	if showVersion {
		fmt.Println(os.Args[0], version)
		os.Exit(0)
	}

	switch mode {
	case "controller":
		const (
			groupName  = "snapshot.storage.k8s.io"
			apiVersion = "v1alpha1"
			kind       = "VolumeSnapshotLink"
			resource   = "volumesnapshotlinks"
		)
		var (
			gk  = schema.GroupKind{Group: groupName, Kind: kind}
			gvr = schema.GroupVersionResource{Group: groupName, Version: apiVersion, Resource: resource}
		)
		populator_machinery.RunController(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath,
			namespace, prefix, gk, gvr, mountPath, devicePath, getPopulatorPodArgs, patchPodFunc)
	case "populate":
		cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
		if err != nil {
			klog.Fatalf("Failed to create config: %v", err)
		}

		// gatewayclientset.NewForConfig creates a new Clientset for GatewayV1alpha2Client
		gcl, err := gatewayclientset.NewForConfig(cfg)
		if err != nil {
			klog.Fatalf("Failed to create gateway client: %v", err)
		}

		// Check if populating the volume is allowed.
		if !isPopulatingAllowed(gcl, sourceNamespace, destNamespace, sourceSnapshotName) {
			klog.Fatalf("Populating volume from snapshot %s/%s is not allowed from namespace %s", sourceNamespace, sourceSnapshotName, destNamespace)
		}

		// snapclientset.NewForConfig creates a new Clientset for  VolumesnapshotV1Client
		snapClient, err := snapclientset.NewForConfig(cfg)
		if err != nil {
			klog.Fatalf("Failed to create snapshot client: %v", err)
		}

		// Get sourcePath and destPath from the source.
		sourcePath, err := getSourcePath(snapClient, sourceNamespace, sourceSnapshotName)
		if err != nil {
			klog.Fatalf("Failed to get source path from snapshot %s/%s: %v", sourceNamespace, sourceSnapshotName, err)
		}

		var destPath string
		switch volumeMode {
		case "file":
			destPath = mountPath
		case "block":
			destPath = devicePath
		default:
			klog.Fatalf("unknown mode: %v", volumeMode)
		}

		// Populate the volume.
		populate(volumeMode, sourcePath, destPath)
	default:
		klog.Fatalf("Invalid mode: %s", mode)
	}
}

func populate(volumeMode, sourcePath, destPath string) {
	if "" == sourcePath || "" == destPath {
		klog.Fatalf("Missing required arg")
	}

	// Below logics are almost the same as loadFromSnapshot.
	// https://github.com/kubernetes-csi/csi-driver-host-path/blob/master/pkg/hostpath/hostpath.go#L285
	var cmd []string
	switch volumeMode {
	case "file":
		cmd = []string{"tar", "zxvf", sourcePath, "-C", destPath}
	case "block":
		cmd = []string{"dd", "if=" + sourcePath, "of=" + destPath}
	default:
		klog.Fatalf("unknown mode: %v", volumeMode)
	}

	executor := utilexec.New()
	klog.Infof("Command Start: %v", cmd)
	out, err := executor.Command(cmd[0], cmd[1:]...).CombinedOutput()
	klog.Infof("Command Finish: %v", string(out))
	if err != nil {
		klog.Fatalf("failed to populate data (%v): %w: %s", volumeMode, err, out)
	}
}

type VolumeSnapshotLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec VolumeSnapshotLinkSpec `json:"spec"`
}

type VolumeSnapshotLinkSpec struct {
	Source VolumeSnapshotLinkSource `json:"source"`
}

type VolumeSnapshotLinkSource struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func getPopulatorPodArgs(rawBlock bool, u *unstructured.Unstructured) ([]string, error) {
	var vsl VolumeSnapshotLink
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &vsl)
	if nil != err {
		return nil, err
	}

	// Create args.
	args := []string{"--mode=populate"}
	if rawBlock {
		args = append(args, "--volume-mode=block")
	} else {
		args = append(args, "--volume-mode=file")
	}
	args = append(args, "--source-namespace="+vsl.Spec.Source.Namespace)
	args = append(args, "--dest-namespace="+vsl.Namespace)
	args = append(args, "--source-snapshot-name="+vsl.Spec.Source.Name)

	return args, nil
}

func isPopulatingAllowed(gcl gatewayclientset.Interface, sourceNamespace, destNamespace, sourceSnapshotName string) bool {
	referencePolicies, err := gcl.GatewayV1alpha2().ReferencePolicies(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		// TODO: Handle error properly.
		klog.Fatalf("Failed to get referencePolicies: %v", err)
		return false
	}

	var allowed bool
	// Check that accessing to snapshotLinkObj is allowed.
	for _, policy := range referencePolicies.Items {
		var validFrom bool
		for _, from := range policy.Spec.From {
			if from.Group == snapshotAPIGroup && from.Kind == snapshotLinkKind &&
				string(from.Namespace) == destNamespace {
				validFrom = true
				break
			}
		}
		// Skip unrelated policy by checking From field
		if !validFrom {
			continue
		}

		for _, to := range policy.Spec.To {
			if to.Group != snapshotAPIGroup || to.Kind != snapshotKind {
				continue
			}
			if to.Name == nil || string(*to.Name) == "" || string(*to.Name) == sourceSnapshotName {
				allowed = true
				break
			}
		}

		if allowed {
			break
		}
	}

	return allowed
}

func getSourcePath(scl snapclientset.Interface, sourceNamespace, sourceSnapshotName string) (string, error) {
	// Below logics are partially copied and modified from getSnapshotSource.
	// https://github.com/kubernetes-csi/external-provisioner/blob/master/pkg/controller/controller.go#L1067
	snapshotObj, err := scl.SnapshotV1().VolumeSnapshots(sourceNamespace).Get(context.TODO(), sourceSnapshotName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("error getting snapshot %s/%s from api server: %v", sourceNamespace, sourceSnapshotName, err)
	}

	if snapshotObj.ObjectMeta.DeletionTimestamp != nil {
		return "", fmt.Errorf("snapshot %s/%s is currently being deleted", sourceNamespace, sourceSnapshotName)
	}

	if snapshotObj.Status == nil || snapshotObj.Status.BoundVolumeSnapshotContentName == nil {
		return "", fmt.Errorf("snapshot %s/%s not bound", sourceNamespace, sourceSnapshotName)
	}

	snapContentObj, err := scl.SnapshotV1().VolumeSnapshotContents().Get(context.TODO(), *snapshotObj.Status.BoundVolumeSnapshotContentName, metav1.GetOptions{})
	if err != nil {
		klog.Warningf("error getting snapshotcontent %s for snapshot %s/%s from api server: %s", *snapshotObj.Status.BoundVolumeSnapshotContentName, snapshotObj.Namespace, snapshotObj.Name, err)
		return "", fmt.Errorf("snapshot %s/%s not bound", sourceNamespace, sourceSnapshotName)
	}

	if snapContentObj.Spec.Driver != driverName {
		return "", fmt.Errorf("snapshot %s/%s is for %s, not for %s", sourceNamespace, sourceSnapshotName, snapContentObj.Spec.Driver, driverName)
	}

	if snapshotObj.Status.ReadyToUse == nil || *snapshotObj.Status.ReadyToUse == false {
		return "", fmt.Errorf("snapshot %s/%s is not Ready", sourceNamespace, sourceSnapshotName)
	}

	if snapContentObj.Status == nil || snapContentObj.Status.SnapshotHandle == nil {
		return "", fmt.Errorf("snapshot %s/%s is not available", sourceNamespace, sourceSnapshotName)

	}

	snapshotID := *snapContentObj.Status.SnapshotHandle

	// Below logics are partially copied and modified from getSnapshotPath.
	// https://github.com/kubernetes-csi/csi-driver-host-path/blob/master/pkg/hostpath/hostpath.go#L137-L140
	path := filepath.Join(stateDir, fmt.Sprintf("%s%s", snapshotID, snapshotExt))

	return path, nil
}

func patchPodFunc(pod *corev1.Pod, pvc *corev1.PersistentVolumeClaim, dataSource *unstructured.Unstructured) (*corev1.Pod, error) {
	pod.Spec.Volumes = append(pod.Spec.Volumes, corev1.Volume{
		Name: "csi-data-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/csi-hostpath-data/",
			},
		},
	})
	con := &pod.Spec.Containers[0]
	con.VolumeMounts = append(con.VolumeMounts, corev1.VolumeMount{
		Name:      "csi-data-dir",
		MountPath: "/csi-data-dir",
	})
	pod.Spec.ServiceAccountName = "csi-hostpath-transfer-account"

	return pod, nil
}
