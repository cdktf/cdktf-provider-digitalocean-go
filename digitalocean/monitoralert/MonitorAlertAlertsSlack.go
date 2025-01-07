// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoralert


type MonitorAlertAlertsSlack struct {
	// The Slack channel to send alerts to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/monitor_alert#channel MonitorAlert#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The webhook URL for Slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/monitor_alert#url MonitorAlert#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

