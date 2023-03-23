package uptimealert


type UptimeAlertNotificationsSlack struct {
	// The Slack channel to send alerts to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/uptime_alert#channel UptimeAlert#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The webhook URL for Slack.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/uptime_alert#url UptimeAlert#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

