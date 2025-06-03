// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package uptimecheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UptimeCheckConfig struct {
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
	// A human-friendly display name for the check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.55.0/docs/resources/uptime_check#name UptimeCheck#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The endpoint to perform healthchecks on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.55.0/docs/resources/uptime_check#target UptimeCheck#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// A boolean value indicating whether the check is enabled/disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.55.0/docs/resources/uptime_check#enabled UptimeCheck#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An array containing the selected regions to perform healthchecks from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.55.0/docs/resources/uptime_check#regions UptimeCheck#regions}
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// The type of health check to perform. Enum: 'ping' 'http' 'https'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.55.0/docs/resources/uptime_check#type UptimeCheck#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

