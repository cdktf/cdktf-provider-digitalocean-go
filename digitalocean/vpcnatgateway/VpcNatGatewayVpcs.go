// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcnatgateway


type VpcNatGatewayVpcs struct {
	// ID of the ingress VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/vpc_nat_gateway#vpc_uuid VpcNatGateway#vpc_uuid}
	VpcUuid *string `field:"required" json:"vpcUuid" yaml:"vpcUuid"`
	// Indicates if this is the default VPC NAT Gateway in the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/vpc_nat_gateway#default_gateway VpcNatGateway#default_gateway}
	DefaultGateway interface{} `field:"optional" json:"defaultGateway" yaml:"defaultGateway"`
}

