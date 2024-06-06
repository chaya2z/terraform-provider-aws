// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package opsworks

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	endpoints_sdkv1 "github.com/aws/aws-sdk-go/aws/endpoints"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	opsworks_sdkv1 "github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceApplication,
			TypeName: "aws_opsworks_application",
		},
		{
			Factory:  ResourceCustomLayer,
			TypeName: "aws_opsworks_custom_layer",
			Name:     "Custom Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceECSClusterLayer,
			TypeName: "aws_opsworks_ecs_cluster_layer",
			Name:     "ECS Cluster Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceGangliaLayer,
			TypeName: "aws_opsworks_ganglia_layer",
			Name:     "Ganglia Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceHAProxyLayer,
			TypeName: "aws_opsworks_haproxy_layer",
			Name:     "HAProxy Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_opsworks_instance",
		},
		{
			Factory:  ResourceJavaAppLayer,
			TypeName: "aws_opsworks_java_app_layer",
			Name:     "Java App Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMemcachedLayer,
			TypeName: "aws_opsworks_memcached_layer",
			Name:     "Memcached Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceMySQLLayer,
			TypeName: "aws_opsworks_mysql_layer",
			Name:     "MySQL Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceNodejsAppLayer,
			TypeName: "aws_opsworks_nodejs_app_layer",
			Name:     "NodeJS App Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourcePermission,
			TypeName: "aws_opsworks_permission",
		},
		{
			Factory:  ResourcePHPAppLayer,
			TypeName: "aws_opsworks_php_app_layer",
			Name:     "PHP App Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceRailsAppLayer,
			TypeName: "aws_opsworks_rails_app_layer",
			Name:     "Rails App Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceRDSDBInstance,
			TypeName: "aws_opsworks_rds_db_instance",
		},
		{
			Factory:  ResourceStack,
			TypeName: "aws_opsworks_stack",
			Name:     "Stack",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceStaticWebLayer,
			TypeName: "aws_opsworks_static_web_layer",
			Name:     "Static Web Layer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceUserProfile,
			TypeName: "aws_opsworks_user_profile",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.OpsWorks
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*opsworks_sdkv1.OpsWorks, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	cfg := aws_sdkv1.Config{}

	if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
		tflog.Debug(ctx, "setting endpoint", map[string]any{
			"tf_aws.endpoint": endpoint,
		})
		cfg.Endpoint = aws_sdkv1.String(endpoint)

		if sess.Config.UseFIPSEndpoint == endpoints_sdkv1.FIPSEndpointStateEnabled {
			tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
			cfg.UseFIPSEndpoint = endpoints_sdkv1.FIPSEndpointStateDisabled
		}
	}

	return opsworks_sdkv1.New(sess.Copy(&cfg)), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
