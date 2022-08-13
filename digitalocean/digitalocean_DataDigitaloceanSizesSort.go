// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanSizesSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/sizes#key DataDigitaloceanSizes#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/sizes#direction DataDigitaloceanSizes#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

