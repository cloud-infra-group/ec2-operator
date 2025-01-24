/*
Copyright 2025.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSVPCEndpointServiceSpec defines the desired state of AWSVPCEndpointService.
type AWSVPCEndpointServiceSpec struct {
	// NetworkLoadBalancerARN specifies the ARN of the NLB used for the endpoint service
	// This field is immutable after the resource is created.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf", message="field is immutable"
	NetworkLoadBalancerARN string `json:"networkLoadBalancerARN"`

	// AcceptanceRequired indicates whether endpoint connections require acceptance
	AcceptanceRequired bool `json:"acceptanceRequired"`

	// Tags are optional metadata tags to apply to the VPC endpoint service.
	// Keys must match the regex pattern `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`, have a minimum length of 1,
	// and a maximum length of 128. Values must match the same regex pattern and have a maximum length of 256.
	//
	// +kubebuilder:validation:MaxProperties=50
	Tags map[string]string `json:"tags,omitempty"`
}

// AWSVPCEndpointServiceStatus defines the observed state of AWSVPCEndpointService.
type AWSVPCEndpointServiceStatus struct {
	// ServiceID is the unique identifier of the VPC Endpoint Service in AWS
	ServiceID string `json:"serviceID,omitempty"`

	// State represents the current state of the VPC Endpoint Service (e.g., Available, Pending, Failed)
	State string `json:"state,omitempty"`

	// AcceptedPrincipals is a list of AWS principal ARNs that currently have permissions on the service
	AcceptedPrincipals []string `json:"acceptedPrincipals,omitempty"`

	// Conditions represent the latest available observations of the resource's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastSyncTime is the last time the controller successfully reconciled this resource
	LastSyncTime metav1.Time `json:"lastSyncTime,omitempty"`

	// ErrorMessage provides details if there was an error during reconciliation
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AWSVPCEndpointService is the Schema for the awsvpcendpointservices API.
type AWSVPCEndpointService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSVPCEndpointServiceSpec   `json:"spec,omitempty"`
	Status AWSVPCEndpointServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AWSVPCEndpointServiceList contains a list of AWSVPCEndpointService.
type AWSVPCEndpointServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSVPCEndpointService `json:"items"`
}

type AWSVPCEndpointServiceRef corev1.ObjectReference

func init() {
	SchemeBuilder.Register(&AWSVPCEndpointService{}, &AWSVPCEndpointServiceList{})
}
