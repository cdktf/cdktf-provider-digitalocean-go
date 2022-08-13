// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanTagsSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/tags#key DataDigitaloceanTags#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/tags#direction DataDigitaloceanTags#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

