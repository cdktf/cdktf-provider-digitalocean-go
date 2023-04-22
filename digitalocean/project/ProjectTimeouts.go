package project


type ProjectTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/resources/project#delete Project#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

