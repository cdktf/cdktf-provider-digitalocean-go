package datadigitaloceansshkeys


type DataDigitaloceanSshKeysSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/data-sources/ssh_keys#key DataDigitaloceanSshKeys#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/data-sources/ssh_keys#direction DataDigitaloceanSshKeys#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

