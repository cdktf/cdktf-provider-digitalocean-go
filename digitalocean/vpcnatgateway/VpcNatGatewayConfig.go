// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcnatgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcNatGatewayConfig struct {
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
	// Name of the VPC NAT Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#name VpcNatGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Region of the VPC NAT Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#region VpcNatGateway#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Size of the VPC NAT Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#size VpcNatGateway#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Type of the VPC NAT Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#type VpcNatGateway#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// vpcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#vpcs VpcNatGateway#vpcs}
	Vpcs interface{} `field:"required" json:"vpcs" yaml:"vpcs"`
	// ICMP connection timeout (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#icmp_timeout_seconds VpcNatGateway#icmp_timeout_seconds}
	IcmpTimeoutSeconds *float64 `field:"optional" json:"icmpTimeoutSeconds" yaml:"icmpTimeoutSeconds"`
	// TCP connection timeout (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#tcp_timeout_seconds VpcNatGateway#tcp_timeout_seconds}
	TcpTimeoutSeconds *float64 `field:"optional" json:"tcpTimeoutSeconds" yaml:"tcpTimeoutSeconds"`
	// UDP connection timeout (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/vpc_nat_gateway#udp_timeout_seconds VpcNatGateway#udp_timeout_seconds}
	UdpTimeoutSeconds *float64 `field:"optional" json:"udpTimeoutSeconds" yaml:"udpTimeoutSeconds"`
}

