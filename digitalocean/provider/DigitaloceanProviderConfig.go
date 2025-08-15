// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type DigitaloceanProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#alias DigitaloceanProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The URL to use for the DigitalOcean API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#api_endpoint DigitaloceanProvider#api_endpoint}
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// The maximum number of retries on a failed API request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#http_retry_max DigitaloceanProvider#http_retry_max}
	HttpRetryMax *float64 `field:"optional" json:"httpRetryMax" yaml:"httpRetryMax"`
	// The maximum wait time (in seconds) between failed API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#http_retry_wait_max DigitaloceanProvider#http_retry_wait_max}
	HttpRetryWaitMax *float64 `field:"optional" json:"httpRetryWaitMax" yaml:"httpRetryWaitMax"`
	// The minimum wait time (in seconds) between failed API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#http_retry_wait_min DigitaloceanProvider#http_retry_wait_min}
	HttpRetryWaitMin *float64 `field:"optional" json:"httpRetryWaitMin" yaml:"httpRetryWaitMin"`
	// The rate of requests per second to limit the HTTP client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#requests_per_second DigitaloceanProvider#requests_per_second}
	RequestsPerSecond *float64 `field:"optional" json:"requestsPerSecond" yaml:"requestsPerSecond"`
	// The access key ID for Spaces API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#spaces_access_id DigitaloceanProvider#spaces_access_id}
	SpacesAccessId *string `field:"optional" json:"spacesAccessId" yaml:"spacesAccessId"`
	// The URL to use for the DigitalOcean Spaces API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#spaces_endpoint DigitaloceanProvider#spaces_endpoint}
	SpacesEndpoint *string `field:"optional" json:"spacesEndpoint" yaml:"spacesEndpoint"`
	// The secret access key for Spaces API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#spaces_secret_key DigitaloceanProvider#spaces_secret_key}
	SpacesSecretKey *string `field:"optional" json:"spacesSecretKey" yaml:"spacesSecretKey"`
	// The token key for API operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.65.0/docs#token DigitaloceanProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

