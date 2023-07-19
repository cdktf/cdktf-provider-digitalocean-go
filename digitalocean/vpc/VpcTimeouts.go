package vpc


type VpcTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/vpc#delete Vpc#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

