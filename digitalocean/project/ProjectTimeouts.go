package project


type ProjectTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/project#delete Project#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

