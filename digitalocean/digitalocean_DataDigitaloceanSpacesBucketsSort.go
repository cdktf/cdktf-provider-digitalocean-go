// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DataDigitaloceanSpacesBucketsSort struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/spaces_buckets#key DataDigitaloceanSpacesBuckets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/d/spaces_buckets#direction DataDigitaloceanSpacesBuckets#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

