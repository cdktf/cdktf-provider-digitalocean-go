package datadigitaloceanrecords


type DataDigitaloceanRecordsSort struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/data-sources/records#key DataDigitaloceanRecords#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.1/docs/data-sources/records#direction DataDigitaloceanRecords#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

