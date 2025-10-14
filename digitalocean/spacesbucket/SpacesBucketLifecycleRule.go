// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spacesbucket


type SpacesBucketLifecycleRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/spaces_bucket#enabled SpacesBucket#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/spaces_bucket#abort_incomplete_multipart_upload_days SpacesBucket#abort_incomplete_multipart_upload_days}.
	AbortIncompleteMultipartUploadDays *float64 `field:"optional" json:"abortIncompleteMultipartUploadDays" yaml:"abortIncompleteMultipartUploadDays"`
	// expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/spaces_bucket#expiration SpacesBucket#expiration}
	Expiration *SpacesBucketLifecycleRuleExpiration `field:"optional" json:"expiration" yaml:"expiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/spaces_bucket#id SpacesBucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// noncurrent_version_expiration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/spaces_bucket#noncurrent_version_expiration SpacesBucket#noncurrent_version_expiration}
	NoncurrentVersionExpiration *SpacesBucketLifecycleRuleNoncurrentVersionExpiration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.68.0/docs/resources/spaces_bucket#prefix SpacesBucket#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

