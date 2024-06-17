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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha3

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	apisv1alpha3 "sigs.k8s.io/gateway-api/apis/applyconfiguration/apis/v1alpha3"
	v1alpha3 "sigs.k8s.io/gateway-api/apis/v1alpha3"
	scheme "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/scheme"
)

// BackendTLSPoliciesGetter has a method to return a BackendTLSPolicyInterface.
// A group's client should implement this interface.
type BackendTLSPoliciesGetter interface {
	BackendTLSPolicies(namespace string) BackendTLSPolicyInterface
}

// BackendTLSPolicyInterface has methods to work with BackendTLSPolicy resources.
type BackendTLSPolicyInterface interface {
	Create(ctx context.Context, backendTLSPolicy *v1alpha3.BackendTLSPolicy, opts v1.CreateOptions) (*v1alpha3.BackendTLSPolicy, error)
	Update(ctx context.Context, backendTLSPolicy *v1alpha3.BackendTLSPolicy, opts v1.UpdateOptions) (*v1alpha3.BackendTLSPolicy, error)
	UpdateStatus(ctx context.Context, backendTLSPolicy *v1alpha3.BackendTLSPolicy, opts v1.UpdateOptions) (*v1alpha3.BackendTLSPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.BackendTLSPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.BackendTLSPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.BackendTLSPolicy, err error)
	Apply(ctx context.Context, backendTLSPolicy *apisv1alpha3.BackendTLSPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.BackendTLSPolicy, err error)
	ApplyStatus(ctx context.Context, backendTLSPolicy *apisv1alpha3.BackendTLSPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.BackendTLSPolicy, err error)
	BackendTLSPolicyExpansion
}

// backendTLSPolicies implements BackendTLSPolicyInterface
type backendTLSPolicies struct {
	client rest.Interface
	ns     string
}

// newBackendTLSPolicies returns a BackendTLSPolicies
func newBackendTLSPolicies(c *GatewayV1alpha3Client, namespace string) *backendTLSPolicies {
	return &backendTLSPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backendTLSPolicy, and returns the corresponding backendTLSPolicy object, and an error if there is any.
func (c *backendTLSPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.BackendTLSPolicy, err error) {
	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackendTLSPolicies that match those selectors.
func (c *backendTLSPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.BackendTLSPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha3.BackendTLSPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backendTLSPolicies.
func (c *backendTLSPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backendTLSPolicy and creates it.  Returns the server's representation of the backendTLSPolicy, and an error, if there is any.
func (c *backendTLSPolicies) Create(ctx context.Context, backendTLSPolicy *v1alpha3.BackendTLSPolicy, opts v1.CreateOptions) (result *v1alpha3.BackendTLSPolicy, err error) {
	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendTLSPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backendTLSPolicy and updates it. Returns the server's representation of the backendTLSPolicy, and an error, if there is any.
func (c *backendTLSPolicies) Update(ctx context.Context, backendTLSPolicy *v1alpha3.BackendTLSPolicy, opts v1.UpdateOptions) (result *v1alpha3.BackendTLSPolicy, err error) {
	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(backendTLSPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendTLSPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backendTLSPolicies) UpdateStatus(ctx context.Context, backendTLSPolicy *v1alpha3.BackendTLSPolicy, opts v1.UpdateOptions) (result *v1alpha3.BackendTLSPolicy, err error) {
	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(backendTLSPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backendTLSPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backendTLSPolicy and deletes it. Returns an error if one occurs.
func (c *backendTLSPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backendTLSPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendtlspolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backendTLSPolicy.
func (c *backendTLSPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.BackendTLSPolicy, err error) {
	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied backendTLSPolicy.
func (c *backendTLSPolicies) Apply(ctx context.Context, backendTLSPolicy *apisv1alpha3.BackendTLSPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.BackendTLSPolicy, err error) {
	if backendTLSPolicy == nil {
		return nil, fmt.Errorf("backendTLSPolicy provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(backendTLSPolicy)
	if err != nil {
		return nil, err
	}
	name := backendTLSPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("backendTLSPolicy.Name must be provided to Apply")
	}
	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *backendTLSPolicies) ApplyStatus(ctx context.Context, backendTLSPolicy *apisv1alpha3.BackendTLSPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.BackendTLSPolicy, err error) {
	if backendTLSPolicy == nil {
		return nil, fmt.Errorf("backendTLSPolicy provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(backendTLSPolicy)
	if err != nil {
		return nil, err
	}

	name := backendTLSPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("backendTLSPolicy.Name must be provided to Apply")
	}

	result = &v1alpha3.BackendTLSPolicy{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("backendtlspolicies").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
