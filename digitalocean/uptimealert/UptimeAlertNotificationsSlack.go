// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package uptimealert


type UptimeAlertNotificationsSlack struct {
	// The Slack channel to send alerts to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.31.0/docs/resources/uptime_alert#channel UptimeAlert#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The webhook URL for Slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.31.0/docs/resources/uptime_alert#url UptimeAlert#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

