
# Cross-namespace Snapshot Populator for csi-hostpath Driver

This example demonstrates a populator implementation for populating volume from cross-namespace snapshot.
This populator only works for csi-hostpath driver.

## How To

1. Deploy k8s with AnyVolumeDataSource feature gate enabled (It is enabled by default in k8s 1.24 or later)

ex) In k/k directory:

- k8s 1.23 or prior:

```bash
FEATURE_GATES=AnyVolumeDataSource=true hack/local-up-cluster.sh
```

- k8s 1.24 or later:

```bash
hack/local-up-cluster.sh
```

2. Deploy volumeSnapshot controller and csi-hostpath driver

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/client/config/crd/snapshot.storage.k8s.io_volumesnapshotclasses.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/client/config/crd/snapshot.storage.k8s.io_volumesnapshotcontents.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/client/config/crd/snapshot.storage.k8s.io_volumesnapshots.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/master/deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml

cd /tmp
git clone --depth=1 https://github.com/kubernetes-csi/csi-driver-host-path.git
cd csi-driver-host-path
./deploy/kubernetes-latest/deploy.sh
```

3. Prepare for using this feature

- Deploy VolumePopulator  CRD and volume-data-source-validator

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/volume-data-source-validator/master/client/config/crd/populator.storage.k8s.io_volumepopulators.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/volume-data-source-validator/master/deploy/kubernetes/rbac-data-source-validator.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/volume-data-source-validator/master/deploy/kubernetes/setup-data-source-validator.yaml
```

- Deploy ReferencePolicy CRD and VolumeSnapshotLink CRD

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/master/config/crd/experimental/gateway.networking.k8s.io_referencepolicies.yaml
curl -L https://raw.githubusercontent.com/mkimuram/lib-volume-populator/csi-hostpath-transfer/example/csi-hostpath-transfer/crd.yaml | sed 's|controller-gen.kubebuilder.io/version: v0.7.0|api-approved.kubernetes.io: https://github.com/kubernetes/enhancements/pull/3295|' | kubectl apply -f -
```

- Build and deploiy csi-hostpath-transfer

```bash
cd /tmp
git clone --depth=1 -b csi-hostpath-transfer https://github.com/mkimuram/lib-volume-populator.git
cd lib-volume-populator
make container-csi-hostpath-transfer
kubectl apply -f example/csi-hostpath-transfer/deploy.yaml
```

Depending on the environment, the container image above may need to be imported.
```bash
ctr images list
ctr image export /tmp/csi-hostpath-transfer.tar docker.io/library/csi-hostpath-transfer:latest
ctr -n=k8s.io images import /tmp/csi-hostpath-transfer.tar
```

4. Create StorageClass, PVC, and VolumeSnapshot by the examples in the csi-hostpath repo

Note that touching /var/lib/csi-hostpath-data/${volumeHandle}/dummydata for the original PV is done, before taking snapshot. So, this file should be included in the volume restored from the snapshot.

```bash
cd /tmp/csi-driver-host-path
kubectl apply -f examples/csi-storageclass.yaml
kubectl apply -f examples/csi-pvc.yaml

kubectl get pvc,pv
NAME                            STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS      AGE
persistentvolumeclaim/csi-pvc   Bound    pvc-43b57144-0666-41f0-9554-cff9a42daeb5   1Gi        RWO            csi-hostpath-sc   5s

NAME                                                        CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM             STORAGECLASS      REASON   AGE
persistentvolume/pvc-43b57144-0666-41f0-9554-cff9a42daeb5   1Gi        RWO            Delete           Bound    default/csi-pvc   csi-hostpath-sc            4s


touch /var/lib/csi-hostpath-data/$(kubectl get pv $(kubectl get pvc csi-pvc -o jsonpath='{.spec.volumeName}') -o jsonpath='{.spec.csi.volumeHandle}')/dummydata


cat examples/csi-snapshot-v1beta1.yaml | sed 's|v1beta1|v1|' | kubectl apply -f -

kubectl get volumesnapshot,volumesnapshotcontent
NAME                                                       READYTOUSE   SOURCEPVC   SOURCESNAPSHOTCONTENT   RESTORESIZE   SNAPSHOTCLASS            SNAPSHOTCONTENT                                    CREATIONTIME   AGE
volumesnapshot.snapshot.storage.k8s.io/new-snapshot-demo   true         csi-pvc                             1Gi           csi-hostpath-snapclass   snapcontent-20514895-2c63-4f1c-8f43-990208ae8f87   5s             5s

NAME                                                                                             READYTOUSE   RESTORESIZE   DELETIONPOLICY   DRIVER                VOLUMESNAPSHOTCLASS      VOLUMESNAPSHOT      VOLUMESNAPSHOTNAMESPACE   AGE
volumesnapshotcontent.snapshot.storage.k8s.io/snapcontent-20514895-2c63-4f1c-8f43-990208ae8f87   true         1073741824    Delete           hostpath.csi.k8s.io   csi-hostpath-snapclass   new-snapshot-demo   default                   5s
```

5. Test creating PVC in ns1 namespace from VolumeSnapshot in default namespace 

- Create ReferencePolicy, VolumeSnapshotLink, and PersistentVolumeClaim

```bash
kubectl create ns ns1

cat << EOF | kubectl apply -f -
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: ReferencePolicy
metadata:
  name: bar
  namespace: default
spec:
  from:
  - group: snapshot.storage.k8s.io
    kind: VolumeSnapshotLink
    namespace: ns1
  to:
  - group: snapshot.storage.k8s.io
    kind: VolumeSnapshot
    name: new-snapshot-demo
EOF

cat << EOF | kubectl apply -f -
apiVersion: snapshot.storage.k8s.io/v1alpha1
kind: VolumeSnapshotLink
metadata:
  name: foolink
  namespace: ns1
spec:
  source:
    name: new-snapshot-demo
    namespace: default
EOF

cat << EOF | kubectl apply -f -
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: foo-pvc
  namespace: ns1
spec:
  storageClassName: csi-hostpath-sc
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  dataSourceRef:
    apiGroup: snapshot.storage.k8s.io
    kind: VolumeSnapshotLink
    name: foolink
  volumeMode: Filesystem
EOF
```

- Confirm that the PVC and PV are created from the volumeSnapshot

```bash
kubectl get pvc,pv -n ns1 
NAME                            STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS      AGE
persistentvolumeclaim/foo-pvc   Bound    pvc-bbd237a9-4d92-4b10-b597-7154909e88a4   1Gi        RWO            csi-hostpath-sc   2m33s

NAME                                                        CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM             STORAGECLASS      REASON   AGE
persistentvolume/pvc-43b57144-0666-41f0-9554-cff9a42daeb5   1Gi        RWO            Delete           Bound    default/csi-pvc   csi-hostpath-sc            3m51s
persistentvolume/pvc-bbd237a9-4d92-4b10-b597-7154909e88a4   1Gi        RWO            Delete           Bound    ns1/foo-pvc       csi-hostpath-sc            26s
```

- Confirm that that the PV contains the dummydata

```bash
ls -l /var/lib/csi-hostpath-data/$(kubectl get pv $(kubectl get pvc foo-pvc -n ns1 -o jsonpath='{.spec.volumeName}') -o jsonpath='{.spec.csi.volumeHandle}')
-rw-r--r--. 1 root root 0 Nov  2 21:25 dummydata
```
