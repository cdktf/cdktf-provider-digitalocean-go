// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package project

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectConfig struct {
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
	// the human-readable name for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#name Project#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// the description of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#description Project#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// the environment of the project's resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#environment Project#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#id Project#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// determine if the project is the default or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#is_default Project#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// the purpose of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#purpose Project#purpose}
	Purpose *string `field:"optional" json:"purpose" yaml:"purpose"`
	// the resources associated with the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#resources Project#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/project#timeouts Project#timeouts}
	Timeouts *ProjectTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

