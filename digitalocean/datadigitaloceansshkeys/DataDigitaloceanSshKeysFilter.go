// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceansshkeys


type DataDigitaloceanSshKeysFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/ssh_keys#key DataDigitaloceanSshKeys#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/ssh_keys#values DataDigitaloceanSshKeys#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/ssh_keys#all DataDigitaloceanSshKeys#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/ssh_keys#match_by DataDigitaloceanSshKeys#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

