package app


type AppSpecWorkerLogDestination struct {
	// Name of the log destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#datadog App#datadog}
	Datadog *AppSpecWorkerLogDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// logtail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#logtail App#logtail}
	Logtail *AppSpecWorkerLogDestinationLogtail `field:"optional" json:"logtail" yaml:"logtail"`
	// papertrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#papertrail App#papertrail}
	Papertrail *AppSpecWorkerLogDestinationPapertrail `field:"optional" json:"papertrail" yaml:"papertrail"`
}

