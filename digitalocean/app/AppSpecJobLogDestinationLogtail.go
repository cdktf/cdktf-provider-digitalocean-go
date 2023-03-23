package app


type AppSpecJobLogDestinationLogtail struct {
	// Logtail token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#token App#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

