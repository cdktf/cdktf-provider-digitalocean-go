package app


type AppSpecJobImage struct {
	// The registry type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#registry_type App#registry_type}
	RegistryType *string `field:"required" json:"registryType" yaml:"registryType"`
	// The repository name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#repository App#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The registry name. Must be left empty for the DOCR registry type.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#registry App#registry}
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// The repository tag. Defaults to latest if not provided.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#tag App#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

