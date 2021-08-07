module github.com/kubernetes-csi/lib-volume-populator/example/hello-populator

go 1.16

require (
	github.com/kubernetes-csi/lib-volume-populator v0.0.0
	k8s.io/apimachinery v0.22.0
	k8s.io/klog/v2 v2.9.0
)

replace github.com/kubernetes-csi/lib-volume-populator => ../..
