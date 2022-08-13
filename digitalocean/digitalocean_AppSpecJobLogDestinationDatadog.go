// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecJobLogDestinationDatadog struct {
	// Datadog API key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#api_key App#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Datadog HTTP log intake endpoint.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#endpoint App#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

