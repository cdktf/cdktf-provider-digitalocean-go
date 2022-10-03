package datadigitaloceanspacesbuckets


type DataDigitaloceanSpacesBucketsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/spaces_buckets#key DataDigitaloceanSpacesBuckets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/spaces_buckets#values DataDigitaloceanSpacesBuckets#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/spaces_buckets#all DataDigitaloceanSpacesBuckets#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/spaces_buckets#match_by DataDigitaloceanSpacesBuckets#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

