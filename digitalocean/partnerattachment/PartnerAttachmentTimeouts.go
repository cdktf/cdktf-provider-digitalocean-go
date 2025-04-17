// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package partnerattachment


type PartnerAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/partner_attachment#create PartnerAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/partner_attachment#delete PartnerAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

