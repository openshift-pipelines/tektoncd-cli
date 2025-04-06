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
)

// ObjectMetaApplyConfiguration represents a declarative configuration of the ObjectMeta type for use
// with apply.
type ObjectMetaApplyConfiguration struct {
	Name                       *string                            `json:"name,omitempty"`
	GenerateName               *string                            `json:"generateName,omitempty"`
	Namespace                  *string                            `json:"namespace,omitempty"`
	UID                        *types.UID                         `json:"uid,omitempty"`
	ResourceVersion            *string                            `json:"resourceVersion,omitempty"`
	Generation                 *int64                             `json:"generation,omitempty"`
	CreationTimestamp          *metav1.Time                       `json:"creationTimestamp,omitempty"`
	DeletionTimestamp          *metav1.Time                       `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64                             `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string                  `json:"labels,omitempty"`
	Annotations                map[string]string                  `json:"annotations,omitempty"`
	OwnerReferences            []OwnerReferenceApplyConfiguration `json:"ownerReferences,omitempty"`
	Finalizers                 []string                           `json:"finalizers,omitempty"`
}

// ObjectMetaApplyConfiguration constructs a declarative configuration of the ObjectMeta type for use with
// apply.
func ObjectMeta() *ObjectMetaApplyConfiguration {
	return &ObjectMetaApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithName(value string) *ObjectMetaApplyConfiguration {
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithGenerateName(value string) *ObjectMetaApplyConfiguration {
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithNamespace(value string) *ObjectMetaApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithUID(value types.UID) *ObjectMetaApplyConfiguration {
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithResourceVersion(value string) *ObjectMetaApplyConfiguration {
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithGeneration(value int64) *ObjectMetaApplyConfiguration {
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ObjectMetaApplyConfiguration {
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ObjectMetaApplyConfiguration {
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ObjectMetaApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ObjectMetaApplyConfiguration {
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ObjectMetaApplyConfiguration) WithLabels(entries map[string]string) *ObjectMetaApplyConfiguration {
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
func (b *ObjectMetaApplyConfiguration) WithAnnotations(entries map[string]string) *ObjectMetaApplyConfiguration {
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
func (b *ObjectMetaApplyConfiguration) WithOwnerReferences(values ...*OwnerReferenceApplyConfiguration) *ObjectMetaApplyConfiguration {
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
func (b *ObjectMetaApplyConfiguration) WithFinalizers(values ...string) *ObjectMetaApplyConfiguration {
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ObjectMetaApplyConfiguration) GetName() *string {
	return b.Name
}
