package app


type AppTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/app#create App#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

