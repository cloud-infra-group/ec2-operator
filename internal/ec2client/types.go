package ec2client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . EC2API
type EC2API interface {
	DescribeVpcEndpointServiceConfigurations(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
	CreateVpcEndpointServiceConfiguration(ctx context.Context, input *ec2.CreateVpcEndpointServiceConfigurationInput, optFns ...func(*ec2.Options)) (*ec2.CreateVpcEndpointServiceConfigurationOutput, error)
	ModifyVpcEndpointServicePermissions(ctx context.Context, input *ec2.ModifyVpcEndpointServicePermissionsInput, optFns ...func(*ec2.Options)) (*ec2.ModifyVpcEndpointServicePermissionsOutput, error)
	DescribeVpcEndpointServicePermissions(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error)
	DeleteVpcEndpointServiceConfigurations(ctx context.Context, input *ec2.DeleteVpcEndpointServiceConfigurationsInput, optFns ...func(*ec2.Options)) (*ec2.DeleteVpcEndpointServiceConfigurationsOutput, error)
}
