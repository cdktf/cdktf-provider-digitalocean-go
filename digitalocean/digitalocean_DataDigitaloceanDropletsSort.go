// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanDropletsSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/droplets#key DataDigitaloceanDroplets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/droplets#direction DataDigitaloceanDroplets#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

