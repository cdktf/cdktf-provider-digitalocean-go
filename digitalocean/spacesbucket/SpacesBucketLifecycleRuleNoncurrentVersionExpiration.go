// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spacesbucket


type SpacesBucketLifecycleRuleNoncurrentVersionExpiration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/spaces_bucket#days SpacesBucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

