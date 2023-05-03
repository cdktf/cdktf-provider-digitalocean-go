package app


type AppSpecServiceLogDestination struct {
	// Name of the log destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/resources/app#datadog App#datadog}
	Datadog *AppSpecServiceLogDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// logtail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/resources/app#logtail App#logtail}
	Logtail *AppSpecServiceLogDestinationLogtail `field:"optional" json:"logtail" yaml:"logtail"`
	// papertrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/resources/app#papertrail App#papertrail}
	Papertrail *AppSpecServiceLogDestinationPapertrail `field:"optional" json:"papertrail" yaml:"papertrail"`
}

