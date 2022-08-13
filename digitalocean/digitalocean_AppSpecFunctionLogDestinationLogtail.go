// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecFunctionLogDestinationLogtail struct {
	// Logtail token.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#token App#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

