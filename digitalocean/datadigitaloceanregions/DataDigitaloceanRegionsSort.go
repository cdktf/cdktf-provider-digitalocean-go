// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanregions


type DataDigitaloceanRegionsSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/regions#key DataDigitaloceanRegions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.60.0/docs/data-sources/regions#direction DataDigitaloceanRegions#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

