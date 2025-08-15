// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package partnerattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PartnerAttachmentConfig struct {
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
	// The connection bandwidth in Mbps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#connection_bandwidth_in_mbps PartnerAttachment#connection_bandwidth_in_mbps}
	ConnectionBandwidthInMbps *float64 `field:"required" json:"connectionBandwidthInMbps" yaml:"connectionBandwidthInMbps"`
	// The NaaS provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#naas_provider PartnerAttachment#naas_provider}
	NaasProvider *string `field:"required" json:"naasProvider" yaml:"naasProvider"`
	// The name of the Partner Attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#name PartnerAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The region where the Partner Attachment will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#region PartnerAttachment#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The list of VPC IDs to attach the Partner Attachment to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#vpc_ids PartnerAttachment#vpc_ids}
	VpcIds *[]*string `field:"required" json:"vpcIds" yaml:"vpcIds"`
	// bgp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#bgp PartnerAttachment#bgp}
	Bgp *PartnerAttachmentBgp `field:"optional" json:"bgp" yaml:"bgp"`
	// The UUID of the Parent Partner Attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#parent_uuid PartnerAttachment#parent_uuid}
	ParentUuid *string `field:"optional" json:"parentUuid" yaml:"parentUuid"`
	// The redundancy zone for the NaaS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#redundancy_zone PartnerAttachment#redundancy_zone}
	RedundancyZone *string `field:"optional" json:"redundancyZone" yaml:"redundancyZone"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/partner_attachment#timeouts PartnerAttachment#timeouts}
	Timeouts *PartnerAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

