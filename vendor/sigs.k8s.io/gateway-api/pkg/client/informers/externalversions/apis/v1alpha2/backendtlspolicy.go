/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apisv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	versioned "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	internalinterfaces "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "sigs.k8s.io/gateway-api/pkg/client/listers/apis/v1alpha2"
)

// BackendTLSPolicyInformer provides access to a shared informer and lister for
// BackendTLSPolicies.
type BackendTLSPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.BackendTLSPolicyLister
}

type backendTLSPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBackendTLSPolicyInformer constructs a new informer for BackendTLSPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackendTLSPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackendTLSPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBackendTLSPolicyInformer constructs a new informer for BackendTLSPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackendTLSPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1alpha2().BackendTLSPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1alpha2().BackendTLSPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&apisv1alpha2.BackendTLSPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *backendTLSPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackendTLSPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backendTLSPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1alpha2.BackendTLSPolicy{}, f.defaultInformer)
}

func (f *backendTLSPolicyInformer) Lister() v1alpha2.BackendTLSPolicyLister {
	return v1alpha2.NewBackendTLSPolicyLister(f.Informer().GetIndexer())
}
