// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// DefaultNetworkDefinitionApplyConfiguration represents an declarative configuration of the DefaultNetworkDefinition type for use
// with apply.
type DefaultNetworkDefinitionApplyConfiguration struct {
	Type                *v1.NetworkType                        `json:"type,omitempty"`
	OpenShiftSDNConfig  *OpenShiftSDNConfigApplyConfiguration  `json:"openshiftSDNConfig,omitempty"`
	OVNKubernetesConfig *OVNKubernetesConfigApplyConfiguration `json:"ovnKubernetesConfig,omitempty"`
	KuryrConfig         *KuryrConfigApplyConfiguration         `json:"kuryrConfig,omitempty"`
}

// DefaultNetworkDefinitionApplyConfiguration constructs an declarative configuration of the DefaultNetworkDefinition type for use with
// apply.
func DefaultNetworkDefinition() *DefaultNetworkDefinitionApplyConfiguration {
	return &DefaultNetworkDefinitionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DefaultNetworkDefinitionApplyConfiguration) WithType(value v1.NetworkType) *DefaultNetworkDefinitionApplyConfiguration {
	b.Type = &value
	return b
}

// WithOpenShiftSDNConfig sets the OpenShiftSDNConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpenShiftSDNConfig field is set to the value of the last call.
func (b *DefaultNetworkDefinitionApplyConfiguration) WithOpenShiftSDNConfig(value *OpenShiftSDNConfigApplyConfiguration) *DefaultNetworkDefinitionApplyConfiguration {
	b.OpenShiftSDNConfig = value
	return b
}

// WithOVNKubernetesConfig sets the OVNKubernetesConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OVNKubernetesConfig field is set to the value of the last call.
func (b *DefaultNetworkDefinitionApplyConfiguration) WithOVNKubernetesConfig(value *OVNKubernetesConfigApplyConfiguration) *DefaultNetworkDefinitionApplyConfiguration {
	b.OVNKubernetesConfig = value
	return b
}

// WithKuryrConfig sets the KuryrConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KuryrConfig field is set to the value of the last call.
func (b *DefaultNetworkDefinitionApplyConfiguration) WithKuryrConfig(value *KuryrConfigApplyConfiguration) *DefaultNetworkDefinitionApplyConfiguration {
	b.KuryrConfig = value
	return b
}
