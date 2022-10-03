package datadigitaloceanrecords


type DataDigitaloceanRecordsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/records#key DataDigitaloceanRecords#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/records#values DataDigitaloceanRecords#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/records#all DataDigitaloceanRecords#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/records#match_by DataDigitaloceanRecords#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

