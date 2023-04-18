package app


type AppSpecServiceLogDestinationPapertrail struct {
	// Papertrail syslog endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#endpoint App#endpoint}
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
}

