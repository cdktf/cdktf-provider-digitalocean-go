// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJobAlertDestinationsSlackWebhooks struct {
	// The Slack channel to send notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/app#channel App#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The Slack webhook URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/app#url App#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

