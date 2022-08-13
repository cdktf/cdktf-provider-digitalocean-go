// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type DigitaloceanProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean#alias DigitaloceanProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The URL to use for the DigitalOcean API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean#api_endpoint DigitaloceanProvider#api_endpoint}
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// The access key ID for Spaces API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean#spaces_access_id DigitaloceanProvider#spaces_access_id}
	SpacesAccessId *string `field:"optional" json:"spacesAccessId" yaml:"spacesAccessId"`
	// The URL to use for the DigitalOcean Spaces API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean#spaces_endpoint DigitaloceanProvider#spaces_endpoint}
	SpacesEndpoint *string `field:"optional" json:"spacesEndpoint" yaml:"spacesEndpoint"`
	// The secret access key for Spaces API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean#spaces_secret_key DigitaloceanProvider#spaces_secret_key}
	SpacesSecretKey *string `field:"optional" json:"spacesSecretKey" yaml:"spacesSecretKey"`
	// The token key for API operations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean#token DigitaloceanProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

