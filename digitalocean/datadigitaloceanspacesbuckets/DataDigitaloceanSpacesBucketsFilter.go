// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanspacesbuckets


type DataDigitaloceanSpacesBucketsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.0/docs/data-sources/spaces_buckets#key DataDigitaloceanSpacesBuckets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.0/docs/data-sources/spaces_buckets#values DataDigitaloceanSpacesBuckets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.0/docs/data-sources/spaces_buckets#all DataDigitaloceanSpacesBuckets#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.0/docs/data-sources/spaces_buckets#match_by DataDigitaloceanSpacesBuckets#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

