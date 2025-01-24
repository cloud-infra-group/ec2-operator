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

package controller_test

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest/komega"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	ec2operatorv1alpha1 "github.com/cloud-infra-group/ec2-operator/api/v1alpha1"
	"github.com/cloud-infra-group/ec2-operator/internal/controller"
	mock "github.com/cloud-infra-group/ec2-operator/internal/ec2client/ec2clientfakes"
)

var _ = Describe("AWSVPCEndpointServicePrincipal Controller", func() {
	var (
		ctx                  context.Context
		mockEC2Client        *mock.FakeEC2API
		controllerReconciler *controller.AWSVPCEndpointServicePrincipalReconciler
		principal            = "arn:aws:iam::987654321012:root"

		waitShort       = 10 * time.Second
		pollingInterval = 500 * time.Millisecond
	)
	komega.SetClient(k8sClient)
	Context("When reconciling a resource", func() {
		const (
			resourceName = "test-principal-resource"
		)

		ctx = context.Background()
		mockEC2Client = &mock.FakeEC2API{}

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		awsvpcendpointserviceprincipal := &ec2operatorv1alpha1.AWSVPCEndpointServicePrincipal{}

		BeforeEach(func() {

			By("creating the custom resource for the Kind AWSVPCEndpointServicePrincipal")
			err := k8sClient.Get(ctx, typeNamespacedName, awsvpcendpointserviceprincipal)
			if err != nil && errors.IsNotFound(err) {
				resource := &ec2operatorv1alpha1.AWSVPCEndpointServicePrincipal{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: ec2operatorv1alpha1.AWSVPCEndpointServicePrincipalSpec{
						AWSVPCEndpointServiceRef: ec2operatorv1alpha1.AWSVPCEndpointServiceRef{Name: resourceName},
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			resource := &ec2operatorv1alpha1.AWSVPCEndpointServicePrincipal{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance AWSVPCEndpointServicePrincipal")
			Expect(client.IgnoreNotFound(err)).To(Succeed())
		})
		It("should successfully reconcile the resource and add the Service Principal to the VPC endpoint service", func() {
			mockEC2Client.DescribeVpcEndpointServicePermissionsReturns(
				&ec2.DescribeVpcEndpointServicePermissionsOutput{
					AllowedPrincipals: []ec2types.AllowedPrincipal{}},
				nil)
			mockEC2Client.ModifyVpcEndpointServicePermissionsReturns(
				&ec2.ModifyVpcEndpointServicePermissionsOutput{
					AddedPrincipals: []ec2types.AddedPrincipal{
						{Principal: &principal},
					},
				}, nil)
			By("Reconciling the created resource")
			controllerReconciler = &controller.AWSVPCEndpointServicePrincipalReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(mockEC2Client.DescribeVpcEndpointServicePermissionsCallCount()).To(Equal(1))
			Expect(mockEC2Client.ModifyVpcEndpointServicePermissionsCallCount()).To(Equal(1))

			updatedResource := &ec2operatorv1alpha1.AWSVPCEndpointServicePrincipal{
				ObjectMeta: metav1.ObjectMeta{
					Name:      typeNamespacedName.Name,
					Namespace: typeNamespacedName.Namespace,
				},
			}
			By(fmt.Sprintf("Waiting for AWSVPCEndpointServicePrincipal resource %s status to be 'Ready'", typeNamespacedName.Name))
			Eventually(komega.Object(updatedResource), waitShort, pollingInterval).Should(HaveField("Status.Ready", true))
		})
		It("should successfully reconcile the resource and do nothing if the principal already exists on the VPC endpoint service", func() {})
		It("should successfully reconcile the resource and delete the Service Principal from the VPC endpoint service", func() {})
		It("should successfully reconcile the resource and do nothing if the Service Principal doesn't exist on the VPC endpoint service", func() {})
		It("should fail to reconcile the resource and add the Service Principal to a non-existent VPC endpoint service", func() {
			mockEC2Client.DescribeVpcEndpointServicePermissionsReturns(
				&ec2.DescribeVpcEndpointServicePermissionsOutput{
					AllowedPrincipals: []ec2types.AllowedPrincipal{}},
				nil)
			mockEC2Client.ModifyVpcEndpointServicePermissionsReturns(
				&ec2.ModifyVpcEndpointServicePermissionsOutput{
					AddedPrincipals: []ec2types.AddedPrincipal{
						{Principal: &principal},
					},
				}, nil)
			mockEC2Client.DescribeVpcEndpointServiceConfigurationsReturns(&ec2.DescribeVpcEndpointServiceConfigurationsOutput{ServiceConfigurations: []ec2types.ServiceConfiguration{}}, nil)

			By("Reconciling the created resource")
			controllerReconciler = &controller.AWSVPCEndpointServicePrincipalReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("not found"))
			Expect(mockEC2Client.DescribeVpcEndpointServiceConfigurationsCallCount()).To(Equal(0))
			Expect(mockEC2Client.DescribeVpcEndpointServicePermissionsCallCount()).To(Equal(0))
			Expect(mockEC2Client.ModifyVpcEndpointServicePermissionsCallCount()).To(Equal(0))

			updatedResource := &ec2operatorv1alpha1.AWSVPCEndpointServicePrincipal{
				ObjectMeta: metav1.ObjectMeta{
					Name:      typeNamespacedName.Name,
					Namespace: typeNamespacedName.Namespace,
				},
			}
			Expect(komega.Object(updatedResource)).Should(HaveField("Status.Ready", false))
		})
	})
})
