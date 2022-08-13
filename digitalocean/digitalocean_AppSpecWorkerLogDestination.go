// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecWorkerLogDestination struct {
	// Name of the log destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#datadog App#datadog}
	Datadog *AppSpecWorkerLogDestinationDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// logtail block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#logtail App#logtail}
	Logtail *AppSpecWorkerLogDestinationLogtail `field:"optional" json:"logtail" yaml:"logtail"`
	// papertrail block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#papertrail App#papertrail}
	Papertrail *AppSpecWorkerLogDestinationPapertrail `field:"optional" json:"papertrail" yaml:"papertrail"`
}

