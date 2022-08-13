// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type MonitorAlertAlerts struct {
	// List of email addresses to sent notifications to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/monitor_alert#email MonitorAlert#email}
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// slack block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/monitor_alert#slack MonitorAlert#slack}
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
}

