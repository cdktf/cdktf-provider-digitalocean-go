// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceandroplets


type DataDigitaloceanDropletsSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/data-sources/droplets#key DataDigitaloceanDroplets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/data-sources/droplets#direction DataDigitaloceanDroplets#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

