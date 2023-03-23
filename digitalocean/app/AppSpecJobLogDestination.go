package app


type AppSpecJobLogDestination struct {
	// Name of the log destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#datadog App#datadog}
	Datadog *AppSpecJobLogDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// logtail block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#logtail App#logtail}
	Logtail *AppSpecJobLogDestinationLogtail `field:"optional" json:"logtail" yaml:"logtail"`
	// papertrail block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#papertrail App#papertrail}
	Papertrail *AppSpecJobLogDestinationPapertrail `field:"optional" json:"papertrail" yaml:"papertrail"`
}

