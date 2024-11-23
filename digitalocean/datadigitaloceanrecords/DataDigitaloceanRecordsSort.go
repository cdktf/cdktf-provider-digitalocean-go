// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanrecords


type DataDigitaloceanRecordsSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.44.1/docs/data-sources/records#key DataDigitaloceanRecords#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.44.1/docs/data-sources/records#direction DataDigitaloceanRecords#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

