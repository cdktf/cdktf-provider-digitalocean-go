package app


type AppSpecWorkerLogDestinationPapertrail struct {
	// Papertrail syslog endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#endpoint App#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}
