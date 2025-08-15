// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanfirewall

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanFirewallConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/data-sources/firewall#firewall_id DataDigitaloceanFirewall#firewall_id}.
	FirewallId *string `field:"required" json:"firewallId" yaml:"firewallId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/data-sources/firewall#droplet_ids DataDigitaloceanFirewall#droplet_ids}.
	DropletIds *[]*float64 `field:"optional" json:"dropletIds" yaml:"dropletIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/data-sources/firewall#id DataDigitaloceanFirewall#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inbound_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/data-sources/firewall#inbound_rule DataDigitaloceanFirewall#inbound_rule}
	InboundRule interface{} `field:"optional" json:"inboundRule" yaml:"inboundRule"`
	// outbound_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/data-sources/firewall#outbound_rule DataDigitaloceanFirewall#outbound_rule}
	OutboundRule interface{} `field:"optional" json:"outboundRule" yaml:"outboundRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/data-sources/firewall#tags DataDigitaloceanFirewall#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

