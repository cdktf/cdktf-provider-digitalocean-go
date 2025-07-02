// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package droplet


type DropletTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/droplet#create Droplet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/droplet#delete Droplet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/droplet#update Droplet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

