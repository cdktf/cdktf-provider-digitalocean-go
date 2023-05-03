package datadigitaloceantags


type DataDigitaloceanTagsSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/data-sources/tags#key DataDigitaloceanTags#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/data-sources/tags#direction DataDigitaloceanTags#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

