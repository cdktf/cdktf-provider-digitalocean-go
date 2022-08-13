// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanRegionsSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/regions#key DataDigitaloceanRegions#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/regions#direction DataDigitaloceanRegions#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

