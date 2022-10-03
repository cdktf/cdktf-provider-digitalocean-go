package vpc


type VpcTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/vpc#delete Vpc#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}
