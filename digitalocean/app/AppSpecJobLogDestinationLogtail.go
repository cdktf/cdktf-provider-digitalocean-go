package app


type AppSpecJobLogDestinationLogtail struct {
	// Logtail token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#token App#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

