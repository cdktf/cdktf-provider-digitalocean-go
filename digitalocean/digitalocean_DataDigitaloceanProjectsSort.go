// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanProjectsSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/projects#key DataDigitaloceanProjects#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/projects#direction DataDigitaloceanProjects#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

