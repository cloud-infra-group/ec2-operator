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

package controller

import (
	"context"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ec2operatorv1alpha1 "github.com/cloud-infra-group/ec2-operator/api/v1alpha1"
	"github.com/cloud-infra-group/ec2-operator/internal/ec2client"
)

const (
	awsVpcEndpointServiceFinalizer = "finalizer.cloud-infra-group.io/awsvpcendpointservice"
)

// AWSVPCEndpointServiceReconciler reconciles a AWSVPCEndpointService object
type AWSVPCEndpointServiceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	EC2API ec2client.EC2API
}

// +kubebuilder:rbac:groups=ec2operator.cloud-infra-group.io.cloud-infra-group.io,resources=awsvpcendpointservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2operator.cloud-infra-group.io.cloud-infra-group.io,resources=awsvpcendpointservices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2operator.cloud-infra-group.io.cloud-infra-group.io,resources=awsvpcendpointservices/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *AWSVPCEndpointServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling AWSVPCEndpointService", "name", req.NamespacedName)

	if r.Client == nil {
		return ctrl.Result{}, fmt.Errorf("client is nil")
	}

	// Fetch the AWSVPCEndpointService resource
	instance := &ec2operatorv1alpha1.AWSVPCEndpointService{}
	if err := r.Get(ctx, req.NamespacedName, instance); err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Failed to fetch AWSVPCEndpointService")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// If the resource is being deleted, delete the VPCEndpointService
	if !instance.ObjectMeta.DeletionTimestamp.IsZero() {
		logger.Info("AWSVPCEndpointService is marked for deletion")
		if err := r.delete(ctx, instance); err != nil {
			logger.Error(err, "Failed to delete AWSVPCEndpointService from EC2")
			return ctrl.Result{}, err
		}
		if controllerutil.ContainsFinalizer(instance, awsVpcEndpointServiceFinalizer) {
			controllerutil.RemoveFinalizer(instance, awsVpcEndpointServiceFinalizer)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, fmt.Errorf("failed to remove finalizer from AWSVPCEndpointService: %w", err)
			}
		}
		return ctrl.Result{}, nil

	}
	if !controllerutil.ContainsFinalizer(instance, awsVpcEndpointServiceFinalizer) {
		// Set the finalizer
		controllerutil.AddFinalizer(instance, awsVpcEndpointServiceFinalizer)
		logger.Info("Updating AWSVPCEndpointService to set the finalizer", "metadata", instance.ObjectMeta)
		if err := r.Update(ctx, instance); err != nil {
			logger.Error(err, "Failed to update AWSVPCEndpointService")
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// Ensure the AWS VPC Endpoint Service exists in AWS
	serviceID, err := r.ensureVPCEndpointService(ctx, instance)
	if err != nil {
		logger.Error(err, "Failed to ensure VPC Endpoint Service")
		return ctrl.Result{}, err
	}
	logger.Info("Ensured VPC Endpoint Service", "serviceID", serviceID)

	// Update the status
	instance.Status.ServiceID = serviceID
	instance.Status.State = "Available"
	logger.Info("Updating AWSVPCEndpointService.status", "metadata", instance.ObjectMeta)
	if err := r.Status().Update(ctx, instance); err != nil {
		logger.Error(err, "Failed to update AWSVPCEndpointService status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *AWSVPCEndpointServiceReconciler) delete(ctx context.Context, instance *ec2operatorv1alpha1.AWSVPCEndpointService) error {
	logger := log.FromContext(ctx).WithName("delete")

	deleteInput := &ec2.DeleteVpcEndpointServiceConfigurationsInput{ServiceIds: []string{instance.Status.ServiceID}}
	deleteOutput, err := r.EC2API.DeleteVpcEndpointServiceConfigurations(ctx, deleteInput)
	if err != nil {
		logger.Error(err, "Failed to delete AWSVPCEndpointService")
		return err
	}
	if deleteOutput == nil {
		err := fmt.Errorf("DeleteVpcEndpointServiceConfigurations returned nil")
		logger.Error(err, "Failed to delete AWSVPCEndpointService")
		return err
	}
	for _, item := range deleteOutput.Unsuccessful {
		if item.Error != nil {
			err := fmt.Errorf("delete AWSVPCEndpointService %q unsuccessful: %s", *item.ResourceId, *item.Error.Message)
			logger.Error(err, "Failed to delete AWSVPCEndpointService")
			return err
		}
	}
	return nil
}

// ensureVPCEndpointService ensures the VPC endpoint service exists in AWS
func (r *AWSVPCEndpointServiceReconciler) ensureVPCEndpointService(ctx context.Context, instance *ec2operatorv1alpha1.AWSVPCEndpointService) (string, error) {
	logger := log.FromContext(ctx).WithName("ensureVPCEndpointService")

	// Describe existing VPC endpoint services to check if it already exists
	describeInput := &ec2.DescribeVpcEndpointServiceConfigurationsInput{}
	describeOutput, err := r.EC2API.DescribeVpcEndpointServiceConfigurations(ctx, describeInput)
	if err != nil {
		logger.Error(err, "Failed to describe existing VPC endpoint services")
		return "", err
	}

	// Check if a service with the specified NLB ARN already exists
	var existingServiceID string
	for _, service := range describeOutput.ServiceConfigurations {
		for _, arn := range service.NetworkLoadBalancerArns {
			if arn == instance.Spec.NetworkLoadBalancerARN {
				existingServiceID = *service.ServiceId
				break
			}
		}
		if existingServiceID != "" {
			break
		}
	}

	if existingServiceID != "" {
		logger.Info("VPC Endpoint Service already exists", "serviceID", existingServiceID)
		return existingServiceID, nil
	}

	// Create a new VPC endpoint service if one does not exist
	createInput := &ec2.CreateVpcEndpointServiceConfigurationInput{
		AcceptanceRequired: aws.Bool(instance.Spec.AcceptanceRequired),
		NetworkLoadBalancerArns: []string{
			instance.Spec.NetworkLoadBalancerARN,
		},
	}

	if len(instance.Spec.Tags) > 0 {
		var tags []ec2types.Tag
		for key, value := range instance.Spec.Tags {
			tags = append(tags, ec2types.Tag{
				Key:   aws.String(key),
				Value: aws.String(value),
			})
		}
		createInput.TagSpecifications = []ec2types.TagSpecification{
			{
				ResourceType: "vpc-endpoint-service",
				Tags:         tags,
			},
		}
	}

	createOutput, err := r.EC2API.CreateVpcEndpointServiceConfiguration(ctx, createInput)
	if err != nil {
		logger.Error(err, "Failed to create VPC Endpoint Service")
		return "", err
	}

	if createOutput == nil || createOutput.ServiceConfiguration == nil || createOutput.ServiceConfiguration.ServiceId == nil {
		return "", fmt.Errorf("CreateVpcEndpointServiceConfiguration returned nil output")
	}
	serviceID := *createOutput.ServiceConfiguration.ServiceId
	logger.Info("Created new VPC Endpoint Service", "serviceID", serviceID)

	return serviceID, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AWSVPCEndpointServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ec2operatorv1alpha1.AWSVPCEndpointService{}).
		Named("awsvpcendpointservice").
		Complete(r)
}
