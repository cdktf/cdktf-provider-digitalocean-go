package app


type AppTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#create App#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

