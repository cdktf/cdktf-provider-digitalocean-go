package monitoralert


type MonitorAlertAlerts struct {
	// List of email addresses to sent notifications to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/monitor_alert#email MonitorAlert#email}
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/monitor_alert#slack MonitorAlert#slack}
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
}

