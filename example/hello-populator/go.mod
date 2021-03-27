module github.com/kubernetes-csi/lib-volume-populator/example/hello-populator

go 1.15

require (
	github.com/kubernetes-csi/lib-volume-populator v0.0.0
	k8s.io/apimachinery v0.19.9
	k8s.io/klog/v2 v2.8.0
)

replace github.com/kubernetes-csi/lib-volume-populator => ../..
