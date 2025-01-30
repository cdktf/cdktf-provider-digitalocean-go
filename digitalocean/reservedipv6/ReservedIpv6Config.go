// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reservedipv6

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReservedIpv6Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/reserved_ipv6#region_slug ReservedIpv6#region_slug}.
	RegionSlug *string `field:"required" json:"regionSlug" yaml:"regionSlug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/reserved_ipv6#droplet_id ReservedIpv6#droplet_id}.
	DropletId *float64 `field:"optional" json:"dropletId" yaml:"dropletId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/reserved_ipv6#id ReservedIpv6#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/reserved_ipv6#ip ReservedIpv6#ip}.
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
}

