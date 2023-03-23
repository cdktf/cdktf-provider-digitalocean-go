package datadigitaloceandomains


type DataDigitaloceanDomainsFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/domains#key DataDigitaloceanDomains#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/domains#values DataDigitaloceanDomains#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/domains#all DataDigitaloceanDomains#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/domains#match_by DataDigitaloceanDomains#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

