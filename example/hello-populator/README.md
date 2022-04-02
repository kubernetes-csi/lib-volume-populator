
# Hello World Populator Example

This example demonstrates an extremely simple populator implementation.

## How To

Install kubernetes 1.17 or later, and enable the AnyVolumeDataSource
feature gate.

Install the CRD:

`kubectl apply -f crd.yaml`

Install the controller:

`kubectl apply -f deploy.yaml`

Create a CR:

```
kubectl create -f - << EOF
apiVersion: hello.example.com/v1alpha1
kind: Hello
metadata:
  name: hello1
spec:
  fileName: test.txt
  fileContents: Hello, world!
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
    apiGroup: hello.example.com
    kind: Hello
    name: hello1
EOF
```

Create a job to print the contents of the pre-populated file inside the PVC:

```
kubectl create -f - << EOF
apiVersion: batch/v1
kind: Job
metadata:
  name: job1
spec:
  template:
    spec:
      containers:
        - name: test
          image: busybox:1.31.1
          command:
            - cat
            - /mnt/test.txt
          volumeMounts:
          - mountPath: /mnt
            name: vol
      restartPolicy: Never
      volumes:
        - name: vol
          persistentVolumeClaim:
            claimName: pvc1
EOF
```

Wait for the job to complete:

`kubectl wait --for=condition=Complete --timeout=120s job/job1`

Get the logs from the job to verify that it worked:

`kubectl logs job/job1`

### To build the image from code:

`make all`

Make sure you have a repo you can push to, and set the variable

`YOUR_REPO=...`

Push the image to your repo:

```
docker tag hello-populator:latest ${YOUR_REPO}/hello-populator:latest
docker push ${YOUR_REPO}/hello-populator:latest
```

To use the image, update deploy.yaml before installing the controller.
