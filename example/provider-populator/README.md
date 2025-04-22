
# Provider Populator Example

This example demonstrates an extremely simple provider specific populator implementation.

## How To

Install kubernetes 1.33 or later

Install the CRD:

`kubectl apply -f crd.yaml`

Install the controller:

`kubectl apply -f deploy.yaml`

Create a CR:

```
kubectl create -f - << EOF
apiVersion: provider.example.com/v1alpha1
kind: Provider
metadata:
  name: provider1
spec:
  dataSourceName: example-source
EOF
```

Create a PVC:

```
kubectl create -f - << EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc1
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 10Mi
  dataSourceRef:
    apiGroup: provider.example.com
    kind: Provider
    name: provider1
EOF
```

Get the logs from the pod to verify that controller starts:

`kubectl get pod -n provider`

```
NAME                                  READY   STATUS    RESTARTS   AGE
provider-populator-7f6c786fb5-q2vsq   1/1     Running   0          11s
```

`kubectl logs provider-populator-7f6c786fb5-q2vsq -n provider`

```
I0331 08:06:03.602113       1 controller.go:236] Starting populator controller for Provider.provider.example.com
W0331 08:06:03.604038       1 client_config.go:667] Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.
I0331 08:06:03.606450       1 metrics.go:109] Metrics path successfully registered at /metrics
I0331 08:06:03.606920       1 metrics.go:121] Metrics http server successfully started on :8080, /metrics
```

### To build the image from code:

`make all CMDS=provider-populator`

Make sure you have a repo you can push to, and set the variable

`YOUR_REPO=...`

Push the image to your repo:

```
docker tag provider-populator:latest ${YOUR_REPO}/provider-populator:latest
docker push ${YOUR_REPO}/provider-populator:latest
```

To use the image, update deploy.yaml before installing the controller.
