// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanpartnerattachment


type DataDigitaloceanPartnerAttachmentBgp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/data-sources/partner_attachment#local_router_ip DataDigitaloceanPartnerAttachment#local_router_ip}.
	LocalRouterIp *string `field:"optional" json:"localRouterIp" yaml:"localRouterIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/data-sources/partner_attachment#peer_router_asn DataDigitaloceanPartnerAttachment#peer_router_asn}.
	PeerRouterAsn *float64 `field:"optional" json:"peerRouterAsn" yaml:"peerRouterAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.54.0/docs/data-sources/partner_attachment#peer_router_ip DataDigitaloceanPartnerAttachment#peer_router_ip}.
	PeerRouterIp *string `field:"optional" json:"peerRouterIp" yaml:"peerRouterIp"`
}

