package spacesbucket


type SpacesBucketLifecycleRuleExpiration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/spaces_bucket#date SpacesBucket#date}.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/spaces_bucket#days SpacesBucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/spaces_bucket#expired_object_delete_marker SpacesBucket#expired_object_delete_marker}.
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
}

