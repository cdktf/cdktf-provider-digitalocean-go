// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceansizes


type DataDigitaloceanSizesSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.0/docs/data-sources/sizes#key DataDigitaloceanSizes#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.0/docs/data-sources/sizes#direction DataDigitaloceanSizes#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

