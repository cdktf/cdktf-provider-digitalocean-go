// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceandroplets


type DataDigitaloceanDropletsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/data-sources/droplets#key DataDigitaloceanDroplets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/data-sources/droplets#values DataDigitaloceanDroplets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/data-sources/droplets#all DataDigitaloceanDroplets#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/data-sources/droplets#match_by DataDigitaloceanDroplets#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

