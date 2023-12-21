// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanimages


type DataDigitaloceanImagesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/data-sources/images#key DataDigitaloceanImages#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/data-sources/images#values DataDigitaloceanImages#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/data-sources/images#all DataDigitaloceanImages#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/data-sources/images#match_by DataDigitaloceanImages#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

