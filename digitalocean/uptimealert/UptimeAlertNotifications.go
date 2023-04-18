package uptimealert


type UptimeAlertNotifications struct {
	// List of email addresses to sent notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/uptime_alert#email UptimeAlert#email}
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/uptime_alert#slack UptimeAlert#slack}
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
}

