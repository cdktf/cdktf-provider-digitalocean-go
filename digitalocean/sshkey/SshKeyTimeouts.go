// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sshkey


type SshKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/resources/ssh_key#create SshKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

