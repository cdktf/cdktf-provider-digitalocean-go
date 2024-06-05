// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJobImageDeployOnPush struct {
	// Whether to automatically deploy images pushed to DOCR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.39.2/docs/resources/app#enabled App#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

