// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package droplet


type DropletBackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/droplet#hour Droplet#hour}.
	Hour *float64 `field:"optional" json:"hour" yaml:"hour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/droplet#plan Droplet#plan}.
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/droplet#weekday Droplet#weekday}.
	Weekday *string `field:"optional" json:"weekday" yaml:"weekday"`
}

