/*
Copyright 2021 The Kubernetes Authors.

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
	"flag"
	"fmt"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"

	populator_machinery "github.com/kubernetes-csi/lib-volume-populator/populator-machinery"
)

const (
	prefix     = "hello.example.com"
	mountPath  = "/mnt"
	devicePath = "/dev/block"
)

var version = "unknown"

func main() {
	var (
		mode         string
		fileName     string
		fileContents string
		httpEndpoint string
		metricsPath  string
		masterURL    string
		kubeconfig   string
		imageName    string
		showVersion  bool
		namespace    string
	)
	klog.InitFlags(nil)
	// Main arg
	flag.StringVar(&mode, "mode", "", "Mode to run in (controller, populate)")
	// Populate args
	flag.StringVar(&fileName, "file-name", "", "File name to populate")
	flag.StringVar(&fileContents, "file-contents", "", "Contents to populate file with")
	// Controller args
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&imageName, "image-name", "", "Image to use for populating")
	// Metrics args
	flag.StringVar(&httpEndpoint, "http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string, which means the server is disabled.")
	flag.StringVar(&metricsPath, "metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")
	// Other args
	flag.BoolVar(&showVersion, "version", false, "display the version string")
	flag.StringVar(&namespace, "namespace", "hello", "Namespace to deploy controller")
	flag.Parse()

	if showVersion {
		fmt.Println(os.Args[0], version)
		os.Exit(0)
	}

	switch mode {
	case "controller":
		const (
			groupName  = "hello.example.com"
			apiVersion = "v1alpha1"
			kind       = "Hello"
			resource   = "hellos"
		)
		var (
			gk  = schema.GroupKind{Group: groupName, Kind: kind}
			gvr = schema.GroupVersionResource{Group: groupName, Version: apiVersion, Resource: resource}
		)
		populator_machinery.RunController(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath,
			namespace, prefix, gk, gvr, mountPath, devicePath, getPopulatorPodArgs)
	case "populate":
		populate(fileName, fileContents)
	default:
		klog.Fatalf("Invalid mode: %s", mode)
	}
}

func populate(fileName, fileContents string) {
	if "" == fileName || "" == fileContents {
		klog.Fatalf("Missing required arg")
	}
	f, err := os.Create(fileName)
	if nil != err {
		klog.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	if !strings.HasSuffix(fileContents, "\n") {
		fileContents += "\n"
	}

	_, err = f.WriteString(fileContents)
	if nil != err {
		klog.Fatalf("Failed to write to file: %v", err)
	}
}

type Hello struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HelloSpec `json:"spec"`
}

type HelloSpec struct {
	FileName     string `json:"fileName"`
	FileContents string `json:"fileContents"`
}

func getPopulatorPodArgs(rawBlock bool, u *unstructured.Unstructured) ([]string, error) {
	var hello Hello
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &hello)
	if nil != err {
		return nil, err
	}
	args := []string{"--mode=populate"}
	if rawBlock {
		args = append(args, "--file-name="+devicePath)
	} else {
		args = append(args, "--file-name="+mountPath+"/"+hello.Spec.FileName)
	}
	args = append(args, "--file-contents="+hello.Spec.FileContents)
	return args, nil
}
