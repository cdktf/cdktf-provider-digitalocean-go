// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package partnerattachment


type PartnerAttachmentBgp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/partner_attachment#auth_key PartnerAttachment#auth_key}.
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/partner_attachment#local_router_ip PartnerAttachment#local_router_ip}.
	LocalRouterIp *string `field:"optional" json:"localRouterIp" yaml:"localRouterIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/partner_attachment#peer_router_asn PartnerAttachment#peer_router_asn}.
	PeerRouterAsn *float64 `field:"optional" json:"peerRouterAsn" yaml:"peerRouterAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/partner_attachment#peer_router_ip PartnerAttachment#peer_router_ip}.
	PeerRouterIp *string `field:"optional" json:"peerRouterIp" yaml:"peerRouterIp"`
}

