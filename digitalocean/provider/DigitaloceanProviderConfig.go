package provider


type DigitaloceanProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs#alias DigitaloceanProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The URL to use for the DigitalOcean API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs#api_endpoint DigitaloceanProvider#api_endpoint}
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// The access key ID for Spaces API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs#spaces_access_id DigitaloceanProvider#spaces_access_id}
	SpacesAccessId *string `field:"optional" json:"spacesAccessId" yaml:"spacesAccessId"`
	// The URL to use for the DigitalOcean Spaces API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs#spaces_endpoint DigitaloceanProvider#spaces_endpoint}
	SpacesEndpoint *string `field:"optional" json:"spacesEndpoint" yaml:"spacesEndpoint"`
	// The secret access key for Spaces API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs#spaces_secret_key DigitaloceanProvider#spaces_secret_key}
	SpacesSecretKey *string `field:"optional" json:"spacesSecretKey" yaml:"spacesSecretKey"`
	// The token key for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs#token DigitaloceanProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

