// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecMaintenance struct {
	// Indicates whether the app should be archived. Setting this to true implies that enabled is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#archive App#archive}
	Archive interface{} `field:"optional" json:"archive" yaml:"archive"`
	// Indicates whether maintenance mode should be enabled for the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#enabled App#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A custom offline page to display when maintenance mode is enabled or the app is archived.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.70.0/docs/resources/app#offline_page_url App#offline_page_url}
	OfflinePageUrl *string `field:"optional" json:"offlinePageUrl" yaml:"offlinePageUrl"`
}

