// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceanspacesbucketobjects

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanSpacesBucketObjectsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#bucket DataDigitaloceanSpacesBucketObjects#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#region DataDigitaloceanSpacesBucketObjects#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#delimiter DataDigitaloceanSpacesBucketObjects#delimiter}.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#encoding_type DataDigitaloceanSpacesBucketObjects#encoding_type}.
	EncodingType *string `field:"optional" json:"encodingType" yaml:"encodingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#id DataDigitaloceanSpacesBucketObjects#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#max_keys DataDigitaloceanSpacesBucketObjects#max_keys}.
	MaxKeys *float64 `field:"optional" json:"maxKeys" yaml:"maxKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/spaces_bucket_objects#prefix DataDigitaloceanSpacesBucketObjects#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

