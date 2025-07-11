// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanprojects


type DataDigitaloceanProjectsSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/data-sources/projects#key DataDigitaloceanProjects#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.59.0/docs/data-sources/projects#direction DataDigitaloceanProjects#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

