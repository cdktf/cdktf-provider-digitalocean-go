// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecAlert struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#rule App#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#disabled App#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

