// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecServiceAlertDestinations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/app#emails App#emails}.
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
	// slack_webhooks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.62.0/docs/resources/app#slack_webhooks App#slack_webhooks}
	SlackWebhooks interface{} `field:"optional" json:"slackWebhooks" yaml:"slackWebhooks"`
}

