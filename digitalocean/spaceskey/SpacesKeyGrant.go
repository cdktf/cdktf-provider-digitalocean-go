// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spaceskey


type SpacesKeyGrant struct {
	// The name of the bucket to grant the key access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/spaces_key#bucket SpacesKey#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The permission to grant the key. Valid values are `read`, `readwrite`, or `fullaccess`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/spaces_key#permission SpacesKey#permission}
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

