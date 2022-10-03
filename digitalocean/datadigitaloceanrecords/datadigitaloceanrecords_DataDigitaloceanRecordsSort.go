package datadigitaloceanrecords


type DataDigitaloceanRecordsSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/records#key DataDigitaloceanRecords#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/records#direction DataDigitaloceanRecords#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

