// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dropletautoscale

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DropletAutoscaleConfig struct {
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
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/droplet_autoscale#config DropletAutoscale#config}
	Config *DropletAutoscaleConfigA `field:"required" json:"config" yaml:"config"`
	// droplet_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/droplet_autoscale#droplet_template DropletAutoscale#droplet_template}
	DropletTemplate *DropletAutoscaleDropletTemplate `field:"required" json:"dropletTemplate" yaml:"dropletTemplate"`
	// Name of the Droplet autoscale pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/droplet_autoscale#name DropletAutoscale#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

