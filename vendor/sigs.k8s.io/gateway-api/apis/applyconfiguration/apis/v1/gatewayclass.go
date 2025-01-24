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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	internal "sigs.k8s.io/gateway-api/apis/applyconfiguration/internal"
	gatewayapiapisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// GatewayClassApplyConfiguration represents a declarative configuration of the GatewayClass type for use
// with apply.
type GatewayClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *GatewayClassSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *GatewayClassStatusApplyConfiguration `json:"status,omitempty"`
}

// GatewayClass constructs a declarative configuration of the GatewayClass type for use with
// apply.
func GatewayClass(name string) *GatewayClassApplyConfiguration {
	b := &GatewayClassApplyConfiguration{}
	b.WithName(name)
	b.WithKind("GatewayClass")
	b.WithAPIVersion("gateway.networking.k8s.io/v1")
	return b
}

// ExtractGatewayClass extracts the applied configuration owned by fieldManager from
// gatewayClass. If no managedFields are found in gatewayClass for fieldManager, a
// GatewayClassApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// gatewayClass must be a unmodified GatewayClass API object that was retrieved from the Kubernetes API.
// ExtractGatewayClass provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractGatewayClass(gatewayClass *gatewayapiapisv1.GatewayClass, fieldManager string) (*GatewayClassApplyConfiguration, error) {
	return extractGatewayClass(gatewayClass, fieldManager, "")
}

// ExtractGatewayClassStatus is the same as ExtractGatewayClass except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractGatewayClassStatus(gatewayClass *gatewayapiapisv1.GatewayClass, fieldManager string) (*GatewayClassApplyConfiguration, error) {
	return extractGatewayClass(gatewayClass, fieldManager, "status")
}

func extractGatewayClass(gatewayClass *gatewayapiapisv1.GatewayClass, fieldManager string, subresource string) (*GatewayClassApplyConfiguration, error) {
	b := &GatewayClassApplyConfiguration{}
	err := managedfields.ExtractInto(gatewayClass, internal.Parser().Type("io.k8s.sigs.gateway-api.apis.v1.GatewayClass"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(gatewayClass.Name)

	b.WithKind("GatewayClass")
	b.WithAPIVersion("gateway.networking.k8s.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithKind(value string) *GatewayClassApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithAPIVersion(value string) *GatewayClassApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithName(value string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithGenerateName(value string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithNamespace(value string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithUID(value types.UID) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithResourceVersion(value string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithGeneration(value int64) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithCreationTimestamp(value metav1.Time) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *GatewayClassApplyConfiguration) WithLabels(entries map[string]string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *GatewayClassApplyConfiguration) WithAnnotations(entries map[string]string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *GatewayClassApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *GatewayClassApplyConfiguration) WithFinalizers(values ...string) *GatewayClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *GatewayClassApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithSpec(value *GatewayClassSpecApplyConfiguration) *GatewayClassApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *GatewayClassApplyConfiguration) WithStatus(value *GatewayClassStatusApplyConfiguration) *GatewayClassApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *GatewayClassApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}