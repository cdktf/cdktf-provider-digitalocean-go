// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanImagesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/images#key DataDigitaloceanImages#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/images#values DataDigitaloceanImages#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/images#all DataDigitaloceanImages#all}.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/images#match_by DataDigitaloceanImages#match_by}.
	MatchBy *string `field:"optional" json:"matchBy" yaml:"matchBy"`
}

