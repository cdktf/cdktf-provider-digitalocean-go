// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecJobEnv struct {
	// The name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#key App#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The visibility scope of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#scope App#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The type of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#type App#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#value App#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

