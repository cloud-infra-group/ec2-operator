/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSVPCEndpointServicePrincipalSpec defines the desired state of AWSVPCEndpointServicePrincipal.
type AWSVPCEndpointServicePrincipalSpec struct {
	AWSVPCEndpointServiceRef AWSVPCEndpointServiceRef `json:"awsVpcEndpointServiceRef"`
	PrincipalARN             string                   `json:"principalARN"`
}

// AWSVPCEndpointServicePrincipalStatus defines the observed state of AWSVPCEndpointServicePrincipal.
type AWSVPCEndpointServicePrincipalStatus struct {
	Ready bool `json:"ready"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AWSVPCEndpointServicePrincipal is the Schema for the awsvpcendpointserviceprincipals API.
type AWSVPCEndpointServicePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSVPCEndpointServicePrincipalSpec   `json:"spec,omitempty"`
	Status AWSVPCEndpointServicePrincipalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSVPCEndpointServicePrincipalList contains a list of AWSVPCEndpointServicePrincipal.
type AWSVPCEndpointServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSVPCEndpointServicePrincipal `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSVPCEndpointServicePrincipal{}, &AWSVPCEndpointServicePrincipalList{})
}
