// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerImageDeployOnPush struct {
	// Whether to automatically deploy images pushed to DOCR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/app#enabled App#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

