// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type MonitorAlertAlertsSlack struct {
	// The Slack channel to send alerts to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/monitor_alert#channel MonitorAlert#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The webhook URL for Slack.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/monitor_alert#url MonitorAlert#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

