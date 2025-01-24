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

// GRPCRouteMatchApplyConfiguration represents a declarative configuration of the GRPCRouteMatch type for use
// with apply.
type GRPCRouteMatchApplyConfiguration struct {
	Method  *GRPCMethodMatchApplyConfiguration  `json:"method,omitempty"`
	Headers []GRPCHeaderMatchApplyConfiguration `json:"headers,omitempty"`
}

// GRPCRouteMatchApplyConfiguration constructs a declarative configuration of the GRPCRouteMatch type for use with
// apply.
func GRPCRouteMatch() *GRPCRouteMatchApplyConfiguration {
	return &GRPCRouteMatchApplyConfiguration{}
}

// WithMethod sets the Method field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Method field is set to the value of the last call.
func (b *GRPCRouteMatchApplyConfiguration) WithMethod(value *GRPCMethodMatchApplyConfiguration) *GRPCRouteMatchApplyConfiguration {
	b.Method = value
	return b
}

// WithHeaders adds the given value to the Headers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Headers field.
func (b *GRPCRouteMatchApplyConfiguration) WithHeaders(values ...*GRPCHeaderMatchApplyConfiguration) *GRPCRouteMatchApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithHeaders")
		}
		b.Headers = append(b.Headers, *values[i])
	}
	return b
}