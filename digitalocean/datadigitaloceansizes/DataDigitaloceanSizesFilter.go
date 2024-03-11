// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceansizes


type DataDigitaloceanSizesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.35.0/docs/data-sources/sizes#key DataDigitaloceanSizes#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.35.0/docs/data-sources/sizes#values DataDigitaloceanSizes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.35.0/docs/data-sources/sizes#all DataDigitaloceanSizes#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.35.0/docs/data-sources/sizes#match_by DataDigitaloceanSizes#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

