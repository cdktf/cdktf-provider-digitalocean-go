// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppConfig struct {
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
	// dedicated_ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#dedicated_ips App#dedicated_ips}
	DedicatedIps interface{} `field:"optional" json:"dedicatedIps" yaml:"dedicatedIps"`
	// Number of deployments to request per page when polling for deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#deployment_per_page App#deployment_per_page}
	DeploymentPerPage *float64 `field:"optional" json:"deploymentPerPage" yaml:"deploymentPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#id App#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#project_id App#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#spec App#spec}
	Spec *AppSpec `field:"optional" json:"spec" yaml:"spec"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#timeouts App#timeouts}
	Timeouts *AppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

