// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customimage


type CustomImageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs/resources/custom_image#create CustomImage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

