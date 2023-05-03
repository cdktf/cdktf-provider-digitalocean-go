package app


type AppSpecWorkerLogDestinationLogtail struct {
	// Logtail token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/resources/app#token App#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

