// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpc


type VpcTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.45.0/docs/resources/vpc#delete Vpc#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

