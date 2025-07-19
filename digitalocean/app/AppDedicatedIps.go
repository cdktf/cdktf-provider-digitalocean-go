// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppDedicatedIps struct {
	// The ID of the dedicated egress IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/app#id App#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IP address of the dedicated egress IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/app#ip App#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// The status of the dedicated egress IP: 'UNKNOWN', 'ASSIGNING', 'ASSIGNED', or 'REMOVED'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/resources/app#status App#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

