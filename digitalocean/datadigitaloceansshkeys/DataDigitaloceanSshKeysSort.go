package datadigitaloceansshkeys


type DataDigitaloceanSshKeysSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/ssh_keys#key DataDigitaloceanSshKeys#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/ssh_keys#direction DataDigitaloceanSshKeys#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

