package app


type AppSpecDomain struct {
	// The hostname for the domain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the domain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#type App#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Indicates whether the domain includes all sub-domains, in addition to the given domain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#wildcard App#wildcard}
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
	// If the domain uses DigitalOcean DNS and you would like App Platform to automatically manage it for you, set this to the name of the domain on your account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#zone App#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

