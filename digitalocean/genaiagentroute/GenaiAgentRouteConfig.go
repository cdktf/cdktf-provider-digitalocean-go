// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiagentroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GenaiAgentRouteConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The UUID of the child agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/genai_agent_route#child_agent_uuid GenaiAgentRoute#child_agent_uuid}
	ChildAgentUuid *string `field:"required" json:"childAgentUuid" yaml:"childAgentUuid"`
	// The UUID of the parent agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/genai_agent_route#parent_agent_uuid GenaiAgentRoute#parent_agent_uuid}
	ParentAgentUuid *string `field:"required" json:"parentAgentUuid" yaml:"parentAgentUuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/genai_agent_route#id GenaiAgentRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// if-case condition for the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/genai_agent_route#if_case GenaiAgentRoute#if_case}
	IfCase *string `field:"optional" json:"ifCase" yaml:"ifCase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/genai_agent_route#rollback GenaiAgentRoute#rollback}.
	Rollback interface{} `field:"optional" json:"rollback" yaml:"rollback"`
	// A name for the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/genai_agent_route#route_name GenaiAgentRoute#route_name}
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
}

