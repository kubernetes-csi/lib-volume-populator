/*
Copyright 2024 The Kubernetes Authors.

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
	"fmt"
	"reflect"
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/dynamic/dynamiclister"
	dynamicfake "k8s.io/client-go/dynamic/fake"
	kubeinformers "k8s.io/client-go/informers"
	informercorev1 "k8s.io/client-go/informers/core/v1"
	informerstoragev1 "k8s.io/client-go/informers/storage/v1"
	kubefake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"
	gatewayInformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
)

type testCase struct {
	// Name of the test
	name string
	// Key added to the notifyMap
	key string
	// PVC to be processed
	pvcName      string
	pvcNamespace string

	// Object to insert into fake kubeclient/dynClient/gatewayClient before the test starts
	initialObjects []runtime.Object
	// Args for the populator pod
	populatorArgs func(b bool, u *unstructured.Unstructured) ([]string, error)
	// Provider specific data population function
	populateFn func(context.Context, PopulatorParams) error
	// Provider specific data population completeness check function, return true when data transfer gets completed.
	populateCompleteFn func(context.Context, PopulatorParams) (bool, error)
	// The original PVC gets deleted or not
	pvcDeleted bool
	// PvcPrimeMutator is the mutator function for pvcPrime
	pvcPrimeMutator func(PvcPrimeMutatorParams) (*v1.PersistentVolumeClaim, error)
	// Expected errors
	expectedResult error
	// Expected keys in the notifyMap
	expectedKeys []string
	// Expected objects after the test runs
	expectedObjects *vpObjects
}

// vpObjects includes the objects we want to compare after tests run
type vpObjects struct {
	pvc *corev1.PersistentVolumeClaim
	// When set to true, only compare PVC Finalizers
	onlyComparePVCFinalizers bool
	pvcPrime                 *corev1.PersistentVolumeClaim
	// When set to true, only compare mutate feilds
	onlyComparePVCPrimeMutateFields bool
	pod                             *v1.Pod
	pv                              *corev1.PersistentVolume
}

const (
	testPrefix                         = "volume.populator.test"
	testMutatorSuffix                  = "mutate"
	testVpWorkingNamespace             = "test"
	testPvcNamespace                   = "default"
	testPvcName                        = "test-pvc"
	testPvcUid                         = "test-uid"
	testApiGroup                       = "test.api.group"
	testApiVersion                     = testApiGroup + "/v1alpha1"
	testDatasourceKind                 = "TestKind"
	testDataSourceName                 = "test-data-source-name"
	testStorageClassName               = "test-sc"
	testPvcPrimeName                   = populatorPvcPrefix + "-" + testPvcUid
	testPvName                         = "test-pv"
	testNodeName                       = "test-node-name"
	testPodName                        = populatorPodPrefix + "-" + testPvcUid
	testProvisioner                    = "test.provisioner"
	testPopulationOperationStartFailed = "Test populate operation start failed"
	testPopulateCompleteFailed         = "Test populate operation complete failed"
	testMutatePVCPrimeFailed           = "Test mutate pvcPrime failed"
	dataSourceKey                      = "unstructured/" + testPvcNamespace + "/" + testDataSourceName
	storageClassKey                    = "sc/" + testStorageClassName
	podKey                             = "pod/" + testVpWorkingNamespace + "/" + testPodName
	pvcPrimeKey                        = "pvc/" + testVpWorkingNamespace + "/" + testPvcPrimeName
	pvKey                              = "pv/" + testPvName
	pvFinalizer                        = "kubernetes.io/pvc-protection"
	vpFinalizer                        = testPrefix + "/" + "populate-target-protection"
	testImage                          = ""
	testMountPath                      = ""
)

func pvc(name, namespace, nodeName, scName, volumeName string, uid types.UID, finalizers []string, datasourceRef *v1.TypedObjectReference, phase v1.PersistentVolumeClaimPhase, accessMode v1.PersistentVolumeAccessMode) *v1.PersistentVolumeClaim {
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			UID:         uid,
			Annotations: map[string]string{},
			Finalizers:  finalizers,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				accessMode,
			},
			StorageClassName: &scName,
			VolumeName:       volumeName,
			DataSourceRef:    datasourceRef,
		},
		Status: v1.PersistentVolumeClaimStatus{
			Phase: phase,
		},
	}

	if nodeName != "" {
		pvc.Annotations[annSelectedNode] = nodeName
	}

	return pvc
}

func dsf(apiGp, kind, name, namespace string) *v1.TypedObjectReference {
	return &v1.TypedObjectReference{
		APIGroup:  &apiGp,
		Kind:      kind,
		Name:      name,
		Namespace: &namespace,
	}
}

func ust() *unstructured.Unstructured {
	return &unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": testApiVersion,
			"kind":       testDatasourceKind,
			"metadata": map[string]any{
				"name":      testDataSourceName,
				"namespace": testPvcNamespace,
			},
		},
	}
}

func sc(scName string, volumeBindMode storagev1.VolumeBindingMode) *storagev1.StorageClass {
	p := testProvisioner
	r := corev1.PersistentVolumeReclaimDelete
	return &storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: scName,
		},
		Provisioner:       p,
		VolumeBindingMode: &volumeBindMode,
		ReclaimPolicy:     &r,
	}
}

func pod(phase corev1.PodPhase, nodeName string) *v1.Pod {
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      testPodName,
			Namespace: testVpWorkingNamespace,
		},
		Spec: MakePopulatePodSpec(testPvcPrimeName),
	}
	pod.Spec.Volumes[0].VolumeSource.PersistentVolumeClaim.ClaimName = testPvcPrimeName
	con := &pod.Spec.Containers[0]
	con.Image = testImage
	con.Args, _ = populatorArgs(false, nil)
	con.VolumeMounts = []corev1.VolumeMount{
		{
			Name:      populatorPodVolumeName,
			MountPath: testMountPath,
		},
	}
	if phase != "" {
		pod.Status.Phase = phase
	}
	if nodeName != "" {
		pod.Spec.NodeName = nodeName
	}
	return pod
}

func pv(pvcName, pvcNamespace, pvcUid string) *v1.PersistentVolume {
	return &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: testPvName,
			Annotations: map[string]string{
				fmt.Sprintf("%s/populated-from", testPrefix): fmt.Sprintf("%s/%s", testPvcNamespace, testDataSourceName),
			},
		},
		Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeReclaimPolicy: corev1.PersistentVolumeReclaimDelete,
			ClaimRef: &corev1.ObjectReference{
				Name:      pvcName,
				Namespace: pvcNamespace,
				UID:       types.UID(pvcUid),
			},
		},
	}
}

func populatorArgs(b bool, u *unstructured.Unstructured) ([]string, error) {
	var args []string
	return args, nil
}

func populateOperationStartError(ctx context.Context, p PopulatorParams) error {
	return fmt.Errorf(testPopulationOperationStartFailed)
}

func PopulateOperationStartSuccess(ctx context.Context, p PopulatorParams) error {
	return nil
}

func populateCompleteError(ctx context.Context, p PopulatorParams) (bool, error) {
	return false, fmt.Errorf(testPopulateCompleteFailed)
}

func populateNotComplete(ctx context.Context, p PopulatorParams) (bool, error) {
	return false, nil
}

func populateCompleteSuccess(ctx context.Context, p PopulatorParams) (bool, error) {
	return true, nil
}

func pvcPrimeMutateAccessModeRWX(mp PvcPrimeMutatorParams) (*v1.PersistentVolumeClaim, error) {
	accessMode := v1.ReadWriteMany
	mp.PvcPrime.Spec.AccessModes[0] = accessMode
	return mp.PvcPrime, nil
}

func pvcPrimeMutateError(mp PvcPrimeMutatorParams) (*v1.PersistentVolumeClaim, error) {
	return mp.PvcPrime, fmt.Errorf(testMutatePVCPrimeFailed)
}

func pvcPrimeMutatePVCPrimeNil(mp PvcPrimeMutatorParams) (*v1.PersistentVolumeClaim, error) {
	return nil, nil
}

func initTest(test testCase) (
	*controller,
	informercorev1.PersistentVolumeClaimInformer,
	cache.SharedIndexInformer,
	informerstoragev1.StorageClassInformer,
	informercorev1.PodInformer,
	informercorev1.PersistentVolumeInformer,
) {
	gvr := schema.GroupVersionResource{
		Group:    testApiGroup,
		Version:  "v1alpha1",
		Resource: "testdatasources",
	}
	gk := schema.GroupKind{
		Group: testApiGroup,
		Kind:  testDatasourceKind,
	}

	kubeClient := kubefake.NewSimpleClientset()
	dynClient := dynamicfake.NewSimpleDynamicClient(runtime.NewScheme())
	gatewayClient := gatewayfake.NewSimpleClientset()

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	dynInformerFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynClient, time.Second*30)

	pvcInformer := kubeInformerFactory.Core().V1().PersistentVolumeClaims()
	pvInformer := kubeInformerFactory.Core().V1().PersistentVolumes()
	podInformer := kubeInformerFactory.Core().V1().Pods()
	scInformer := kubeInformerFactory.Storage().V1().StorageClasses()
	unstInformer := dynInformerFactory.ForResource(gvr).Informer()

	gatewayInformerFactory := gatewayInformers.NewSharedInformerFactory(gatewayClient, time.Second*30)
	referenceGrants := gatewayInformerFactory.Gateway().V1beta1().ReferenceGrants()

	var podConfig *PodConfig
	if test.populatorArgs != nil {
		podConfig = &PodConfig{
			ImageName:     testImage,
			DevicePath:    "",
			MountPath:     testMountPath,
			PopulatorArgs: test.populatorArgs,
		}
	}

	var providerFunctionConfig *ProviderFunctionConfig
	if test.populateFn != nil || test.populateCompleteFn != nil {
		providerFunctionConfig = &ProviderFunctionConfig{
			PopulateFn:         test.populateFn,
			PopulateCompleteFn: test.populateCompleteFn,
		}
	}

	var mutatorConfig *MutatorConfig
	if test.pvcPrimeMutator != nil {
		mutatorConfig = &MutatorConfig{
			PvcPrimeMutator: test.pvcPrimeMutator,
		}
	}

	c := &controller{
		kubeClient:             kubeClient,
		populatorNamespace:     testVpWorkingNamespace,
		populatedFromAnno:      testPrefix + "/" + populatedFromAnnoSuffix,
		pvcFinalizer:           testPrefix + "/" + pvcFinalizerSuffix,
		pvcLister:              pvcInformer.Lister(),
		pvcSynced:              pvcInformer.Informer().HasSynced,
		pvLister:               pvInformer.Lister(),
		pvSynced:               pvInformer.Informer().HasSynced,
		podLister:              podInformer.Lister(),
		podSynced:              podInformer.Informer().HasSynced,
		scLister:               scInformer.Lister(),
		scSynced:               scInformer.Informer().HasSynced,
		unstLister:             dynamiclister.New(unstInformer.GetIndexer(), gvr),
		unstSynced:             unstInformer.HasSynced,
		notifyMap:              make(map[string]*stringSet),
		cleanupMap:             make(map[string]*stringSet),
		workqueue:              workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		gk:                     gk,
		metrics:                initMetrics(),
		recorder:               getRecorder(kubeClient, testPrefix+"-"+controllerNameSuffix),
		referenceGrantLister:   referenceGrants.Lister(),
		referenceGrantSynced:   referenceGrants.Informer().HasSynced,
		podConfig:              podConfig,
		providerFunctionConfig: providerFunctionConfig,
		mutatorConfig:          mutatorConfig,
	}
	return c, pvcInformer, unstInformer, scInformer, podInformer, pvInformer
}

func compareResult(want error, got error) bool {
	if want == nil {
		return got == nil
	}
	if got == nil {
		return want == nil
	}
	return want.Error() == got.Error()
}

func compareNotifyMap(want []string, got map[string]*stringSet) error {
	if len(want) != len(got) {
		return fmt.Errorf("The number of keys expected is different from actual. Expect %v, got %v", len(want), len(got))
	}
	for _, key := range want {
		if got[key] == nil {
			return fmt.Errorf("Expected key %s not found in the notifyMap", key)
		}
	}
	return nil
}

func compareObjects(want *vpObjects, got *vpObjects) error {
	if want != nil && got == nil {
		return fmt.Errorf("Expected vpObjects is different from actual vpObjects. Expected %+v\n, actual %+v\n", want, got)
	}

	if want.onlyComparePVCFinalizers {
		if !reflect.DeepEqual(want.pvc.Finalizers, got.pvc.Finalizers) {
			return fmt.Errorf("Expected pvc Finalizers is different from actual pvc Finalizers. Expected %+v\n, actual %+v\n",
				want.pvc.Finalizers, got.pvc.Finalizers)
		}
	} else {
		if !reflect.DeepEqual(want.pvc, got.pvc) {
			return fmt.Errorf("Expected pvc is different from actual pvc. Expected %+v\n, actual %+v\n",
				want.pvc, got.pvc)
		}
	}

	if want.onlyComparePVCPrimeMutateFields {
		if !reflect.DeepEqual(want.pvcPrime.Spec.AccessModes, got.pvcPrime.Spec.AccessModes) {
			return fmt.Errorf("Expected pvcPrime AccessModes is different from actual pvcPrime AccessModes. Expected %+v\n, actual %+v\n",
				want.pvcPrime.Spec.AccessModes, got.pvcPrime.Spec.AccessModes)
		}
	} else {
		if !reflect.DeepEqual(want.pvcPrime, got.pvcPrime) {
			return fmt.Errorf("Expected pvcPrime is different from actual pvcPrime. Expected %+v\n, actual %+v\n",
				want.pvcPrime, got.pvcPrime)
		}
	}

	if !reflect.DeepEqual(want.pod, got.pod) {
		return fmt.Errorf("Expected pod is different from actual pod. Expected %+v\n, actual %+v\n",
			want.pod, got.pod)
	}

	if !reflect.DeepEqual(want.pv, got.pv) {
		return fmt.Errorf("Expected pv is different from actual pv. Expected %+v\n, actual %+v\n",
			want.pv, got.pv)
	}
	return nil
}

func runSyncPvcTests(tests []testCase, t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, pvcInformer, unstInformer, scInformer, podInformer, pvInformer := initTest(test)
			for _, obj := range test.initialObjects {
				switch obj.(type) {
				case *v1.PersistentVolumeClaim:
					pvc := obj.(*v1.PersistentVolumeClaim)
					_, err := c.kubeClient.CoreV1().PersistentVolumeClaims(pvc.ObjectMeta.Namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
					if err != nil {
						t.Fatalf("Create pvc failed: %s", err.Error())
					}
					pvcInformer.Informer().GetStore().Add(obj)
				case *unstructured.Unstructured:
					unstInformer.GetStore().Add(obj)
				case *storagev1.StorageClass:
					scInformer.Informer().GetStore().Add(obj)
				case *v1.Pod:
					pod := obj.(*v1.Pod)
					_, err := c.kubeClient.CoreV1().Pods(pod.ObjectMeta.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
					if err != nil {
						t.Fatalf("Create pod failed: %s", err.Error())
					}
					podInformer.Informer().GetStore().Add(obj)
				case *v1.PersistentVolume:
					pv := obj.(*v1.PersistentVolume)
					_, err := c.kubeClient.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
					if err != nil {
						t.Fatalf("Create pv failed: %s", err.Error())
					}
					pvInformer.Informer().GetStore().Add(obj)
				default:
					t.Fatalf("Unknown initalObject type: %+v", obj)
				}
			}

			result := c.syncPvc(context.TODO(), test.key, test.pvcNamespace, test.pvcName)
			if !compareResult(test.expectedResult, result) {
				t.Errorf("Error: expected result %t, got %t", test.expectedResult, result)
			}
			err := compareNotifyMap(test.expectedKeys, c.notifyMap)
			if err != nil {
				t.Errorf("Compare notifyMap failed, error: %+v", err.Error())
			}
			actualObjects := &vpObjects{}
			actualPvc, err := c.kubeClient.CoreV1().PersistentVolumeClaims(test.pvcNamespace).Get(context.TODO(), test.pvcName, metav1.GetOptions{})
			if err != nil {
				if !errors.IsNotFound(err) {
					t.Errorf("Get pvc failed, error: %+v", err.Error())
				}
			}
			if actualPvc != nil && !errors.IsNotFound(err) {
				actualObjects.pvc = actualPvc
			}
			actualPvcPrime, err := c.kubeClient.CoreV1().PersistentVolumeClaims(testVpWorkingNamespace).Get(context.TODO(), testPvcPrimeName, metav1.GetOptions{})
			if err != nil {
				if !errors.IsNotFound(err) {
					t.Errorf("Get pvcPrime failed, error: %+v", err.Error())
				}
			}
			if actualPvcPrime != nil && !errors.IsNotFound(err) {
				actualObjects.pvcPrime = actualPvcPrime
			}
			actualPod, err := c.kubeClient.CoreV1().Pods(testVpWorkingNamespace).Get(context.TODO(), testPodName, metav1.GetOptions{})
			if err != nil {
				if !errors.IsNotFound(err) {
					t.Errorf("Get pod failed, error: %+v", err.Error())
				}
			}
			if actualPod != nil && !errors.IsNotFound(err) {
				actualObjects.pod = actualPod
			}
			actualPV, err := c.kubeClient.CoreV1().PersistentVolumes().Get(context.TODO(), testPvName, metav1.GetOptions{})
			if err != nil {
				if !errors.IsNotFound(err) {
					t.Errorf("Get pv failed, error: %+v", err.Error())
				}
			}
			if actualPV != nil && !errors.IsNotFound(err) {
				actualObjects.pv = actualPV
			}
			if test.expectedObjects != nil {
				err = compareObjects(test.expectedObjects, actualObjects)
				if err != nil {
					t.Errorf("Compare vpObjects failed, error: %+v", err.Error())
				}
			}
		})
	}
}

func TestSyncPvcWithPopulatorPod(t *testing.T) {
	tests := []testCase{
		{
			name:         "Ignore PVCs in controller's working namespace",
			key:          "pvc/" + testVpWorkingNamespace + "/" + testPvcName,
			pvcNamespace: testVpWorkingNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testVpWorkingNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testVpWorkingNamespace), "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testVpWorkingNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testVpWorkingNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:           "Orginal PVC not found",
			key:            "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace:   testPvcNamespace,
			pvcName:        testPvcName,
			initialObjects: []runtime.Object{},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
		},
		{
			name:           "Ignore PVCs without a data source",
			key:            "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace:   testPvcNamespace,
			pvcName:        testPvcName,
			initialObjects: []runtime.Object{pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce)},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce)},
		},
		{
			name:         "Ignore PVCs that aren't for this populator to handle, apiGroup not match",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf("test.api.group1", testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf("test.api.group1", testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Ignore PVCs that aren't for this populator to handle, kind not match",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, "TestKind1", testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, "TestKind1", testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Ignore PVCs that aren't for this populator to handle, data source name not exist",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, "", testPvcNamespace), "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, "", testPvcNamespace), "", v1.ReadWriteOnce)},
		},
		{
			name:         "Original PVC and data source in different namespace without grant",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, "default1"), "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: fmt.Errorf("accessing default1/test-data-source-name of TestKind dataSource from default/test-pvc isn't allowed"),
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, "default1"), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Data source not exists",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{dataSourceKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "StorageClass not exists",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{storageClassKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "PVC not bound to a node",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Create pvcPrime mutate succeeded",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
			},
			populatorArgs:   populatorArgs,
			pvcPrimeMutator: pvcPrimeMutateAccessModeRWX,
			expectedResult:  nil,
			expectedKeys:    []string{podKey, pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				// Only compare PVC finlizeres here because reflect.DeepEqual returns false after PVC gets pathched even
				// if all the fields look exactly the same. The only field we want to verify here is the finalizers gets updated as expected.
				onlyComparePVCFinalizers: true,
				pvcPrime:                 pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, "", "", []string{}, nil, "", v1.ReadWriteMany),
				// Only compare pvcPrime mutate fields here because reflect.DeepEqual returns false after PVC gets mutate even
				// if all the fields look exactly the same.
				onlyComparePVCPrimeMutateFields: true,
				pod:                             pod("", ""),
			},
		},
		{
			name:         "Create pvcPrime mutate error",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
			},
			populatorArgs:   populatorArgs,
			pvcPrimeMutator: pvcPrimeMutateError,
			expectedResult:  fmt.Errorf(testMutatePVCPrimeFailed),
			expectedKeys:    []string{podKey, pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Create pvcPrime mutate, return pvcPrimeMutator pvcPrime nil",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
			},
			populatorArgs:   populatorArgs,
			pvcPrimeMutator: pvcPrimeMutatePVCPrimeNil,
			expectedResult:  fmt.Errorf("pvcPrime must not be nil"),
			expectedKeys:    []string{podKey, pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Create populator pod",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{podKey, pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv:       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
				pod:      pod("", testNodeName),
			},
		},
		{
			name:         "Wait populator pod succeed",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
				pod(corev1.PodRunning, testNodeName),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{podKey, pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv:       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
				pod:      pod(corev1.PodRunning, testNodeName),
			},
		},
		{
			name:         "Populator pod failed",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
				pod(corev1.PodFailed, testNodeName),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{podKey, pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv:       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
		},
		{
			name:         "PV not exists",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pod(corev1.PodSucceeded, testNodeName),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{podKey, pvcPrimeKey, pvKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pod:      pod(corev1.PodSucceeded, testNodeName),
			},
		},
		{
			name:         "Wait for the bind controller to rebind the PV",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
				pod(corev1.PodSucceeded, testNodeName),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{podKey, pvcPrimeKey, pvKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv:       pv(testPvcName, testPvcNamespace, testPvcUid),
				pod:      pod(corev1.PodSucceeded, testNodeName),
			},
		},
		{
			name:         "Clean up populator pod, pvcPrime and other temporary resources",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimLost, v1.ReadWriteOnce),
				pv(testPvcName, testPvcNamespace, testPvcUid),
				pod(corev1.PodSucceeded, testNodeName),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pv: pv(testPvcName, testPvcNamespace, testPvcUid),
			},
		},
		{
			name:         "Delete original PVC",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "Terminating", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
				pod(corev1.PodRunning, testNodeName),
			},
			populatorArgs:  populatorArgs,
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "Terminating", v1.ReadWriteOnce),
				pv: pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
		},
	}

	runSyncPvcTests(tests, t)
}

func TestSyncPvcWithProviderImplementation(t *testing.T) {
	tests := []testCase{
		{
			name:         "Ignore PVCs in controller's working namespace",
			key:          "pvc/" + testVpWorkingNamespace + "/" + testPvcName,
			pvcNamespace: testVpWorkingNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testVpWorkingNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testVpWorkingNamespace), "", v1.ReadWriteOnce),
			},
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testVpWorkingNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testVpWorkingNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:           "Orginal PVC not found",
			key:            "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace:   testPvcNamespace,
			pvcName:        testPvcName,
			initialObjects: []runtime.Object{},
			expectedResult: nil,
			expectedKeys:   []string{},
		},
		{
			name:           "Ignore PVCs without a data source",
			key:            "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace:   testPvcNamespace,
			pvcName:        testPvcName,
			initialObjects: []runtime.Object{pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce)},
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Ignore PVCs that aren't for this populator to handle, apiGroup not match",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf("test.api.group1", testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf("test.api.group1", testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Ignore PVCs that aren't for this populator to handle, kind not match",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, "TestKind1", testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, "TestKind1", testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Ignore PVCs that aren't for this populator to handle, data source name not exist",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, "", testPvcNamespace), "", v1.ReadWriteOnce),
			},
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, "", testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Original PVC and data source in different namespace without grant",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, "default1"), "", v1.ReadWriteOnce),
			},
			expectedResult: fmt.Errorf("accessing default1/test-data-source-name of TestKind dataSource from default/test-pvc isn't allowed"),
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, "default1"), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Data source not exists",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
			expectedResult: nil,
			expectedKeys:   []string{dataSourceKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "StorageClass not exists",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
			},
			expectedResult: nil,
			expectedKeys:   []string{storageClassKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "PVC not bound to a node",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
			},
			expectedResult: nil,
			expectedKeys:   []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Create pvcPrime mutate succeeded",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
			},
			pvcPrimeMutator: pvcPrimeMutateAccessModeRWX,
			populateFn:      populateOperationStartError,
			expectedResult:  nil,
			expectedKeys:    []string{pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				// Only compare PVC finlizeres here because reflect.DeepEqual returns false after PVC gets pathched even
				// if all the fields look exactly the same. The only field we want to verify here is the finalizers gets updated as expected.
				onlyComparePVCFinalizers: true,
				pvcPrime:                 pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, "", []string{}, nil, "", v1.ReadWriteMany),
				// Only compare pvcPrime mutate fields here because reflect.DeepEqual returns false after PVC gets mutate even
				// if all the fields look exactly the same.
				onlyComparePVCPrimeMutateFields: true,
			},
		},
		{
			name:         "Create pvcPrime mutate error",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
			},
			pvcPrimeMutator: pvcPrimeMutateError,
			populateFn:      populateOperationStartError,
			expectedResult:  fmt.Errorf(testMutatePVCPrimeFailed),
			expectedKeys:    []string{pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Create pvcPrime mutate, return pvcPrime nil",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
			},
			pvcPrimeMutator: pvcPrimeMutatePVCPrimeNil,
			populateFn:      populateOperationStartError,
			expectedResult:  fmt.Errorf("pvcPrime must not be nil"),
			expectedKeys:    []string{pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
			},
		},
		{
			name:         "Populate operation start return an error",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
				pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimBound, v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
			populateFn:     populateOperationStartError,
			expectedResult: fmt.Errorf(testPopulationOperationStartFailed),
			expectedKeys:   []string{pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimBound, v1.ReadWriteOnce),
				pv:       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
		},
		{
			name:         "Populate completeness check return an error",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
				pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimBound, v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
			populateFn:         PopulateOperationStartSuccess,
			populateCompleteFn: populateCompleteError,
			expectedResult:     fmt.Errorf(testPopulateCompleteFailed),
			expectedKeys:       []string{pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimBound, v1.ReadWriteOnce),
				pv:       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
		},
		{
			name:         "Populate not complete",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
				pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimBound, v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
			populateFn:         PopulateOperationStartSuccess,
			populateCompleteFn: populateNotComplete,
			expectedResult:     fmt.Errorf(reasonWaitForDataPopulationFinished),
			expectedKeys:       []string{pvcPrimeKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimBound, v1.ReadWriteOnce),
				pv:       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
		},
		{
			name:         "Wait for the bind controller to rebind the PV",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
				pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
			populateFn:         PopulateOperationStartSuccess,
			populateCompleteFn: populateCompleteSuccess,
			expectedResult:     nil,
			expectedKeys:       []string{pvcPrimeKey, pvKey},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				pvcPrime: pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv:       pv(testPvcName, testPvcNamespace, testPvcUid),
			},
		},
		{
			name:         "Clean up pvcPrime",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingImmediate),
				pvc(testPvcPrimeName, testVpWorkingNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, corev1.ClaimLost, v1.ReadWriteOnce),
				pv(testPvcName, testPvcNamespace, testPvcUid),
			},
			populateFn:         PopulateOperationStartSuccess,
			populateCompleteFn: populateCompleteSuccess,
			expectedResult:     nil,
			expectedKeys:       []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, "", testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "", v1.ReadWriteOnce),
				// Only compare PVC finlizeres here because reflect.DeepEqual returns false after PVC gets pathched even
				// if all the fields look exactly the same. The only field we want to verify here is the finalizers gets updated as expected.
				onlyComparePVCFinalizers: true,
				pv:                       pv(testPvcName, testPvcNamespace, testPvcUid),
			},
		},
		{
			name:         "Delete original PVC",
			key:          "pvc/" + testPvcNamespace + "/" + testPvcName,
			pvcNamespace: testPvcNamespace,
			pvcName:      testPvcName,
			initialObjects: []runtime.Object{
				pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer, vpFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "Terminating", v1.ReadWriteOnce),
				ust(),
				sc(testStorageClassName, storagev1.VolumeBindingWaitForFirstConsumer),
				pvc(testPvcPrimeName, testVpWorkingNamespace, testNodeName, testStorageClassName, testPvName, testPvcUid, []string{pvFinalizer}, nil, "", v1.ReadWriteOnce),
				pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
			populateFn:         PopulateOperationStartSuccess,
			populateCompleteFn: populateNotComplete,
			expectedKeys:       []string{},
			expectedObjects: &vpObjects{
				pvc: pvc(testPvcName, testPvcNamespace, testNodeName, testStorageClassName, "", testPvcUid, []string{pvFinalizer},
					dsf(testApiGroup, testDatasourceKind, testDataSourceName, testPvcNamespace), "Terminating", v1.ReadWriteOnce),
				// Only compare PVC finlizeres here  because reflect.DeepEqual returns false after PVC gets pathched even
				// if all the fields look exactly the same. The only field we want to verify here is the finalizers gets updated as expected.
				onlyComparePVCFinalizers: true,
				pv:                       pv(testPvcPrimeName, testVpWorkingNamespace, testPvcUid),
			},
		},
	}

	runSyncPvcTests(tests, t)
}
