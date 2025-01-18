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
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	ec2operatorcloudinfragroupiov1alpha1 "github.com/cloud-infra-group/ec2-operator/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var awsvpcendpointservicelog = logf.Log.WithName("awsvpcendpointservice-resource")

// SetupAWSVPCEndpointServiceWebhookWithManager registers the webhook for AWSVPCEndpointService in the manager.
func SetupAWSVPCEndpointServiceWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&ec2operatorcloudinfragroupiov1alpha1.AWSVPCEndpointService{}).
		WithValidator(&AWSVPCEndpointServiceCustomValidator{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-ec2operator-cloud-infra-group-io-cloud-infra-group-io-v1alpha1-awsvpcendpointservice,mutating=false,failurePolicy=fail,sideEffects=None,groups=ec2operator.cloud-infra-group.io.cloud-infra-group.io,resources=awsvpcendpointservices,verbs=create;update,versions=v1alpha1,name=vawsvpcendpointservice-v1alpha1.kb.io,admissionReviewVersions=v1

// AWSVPCEndpointServiceCustomValidator struct is responsible for validating the AWSVPCEndpointService resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type AWSVPCEndpointServiceCustomValidator struct {
	//TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &AWSVPCEndpointServiceCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type AWSVPCEndpointService.
func (v *AWSVPCEndpointServiceCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	awsvpcendpointservice, ok := obj.(*ec2operatorcloudinfragroupiov1alpha1.AWSVPCEndpointService)
	if !ok {
		return nil, fmt.Errorf("expected a AWSVPCEndpointService object but got %T", obj)
	}
	awsvpcendpointservicelog.Info("Validation for AWSVPCEndpointService upon creation", "name", awsvpcendpointservice.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type AWSVPCEndpointService.
func (v *AWSVPCEndpointServiceCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	awsvpcendpointservice, ok := newObj.(*ec2operatorcloudinfragroupiov1alpha1.AWSVPCEndpointService)
	if !ok {
		return nil, fmt.Errorf("expected a AWSVPCEndpointService object for the newObj but got %T", newObj)
	}
	awsvpcendpointservicelog.Info("Validation for AWSVPCEndpointService upon update", "name", awsvpcendpointservice.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type AWSVPCEndpointService.
func (v *AWSVPCEndpointServiceCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	awsvpcendpointservice, ok := obj.(*ec2operatorcloudinfragroupiov1alpha1.AWSVPCEndpointService)
	if !ok {
		return nil, fmt.Errorf("expected a AWSVPCEndpointService object but got %T", obj)
	}
	awsvpcendpointservicelog.Info("Validation for AWSVPCEndpointService upon deletion", "name", awsvpcendpointservice.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
