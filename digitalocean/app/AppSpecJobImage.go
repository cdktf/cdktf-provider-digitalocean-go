// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJobImage struct {
	// The registry type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#registry_type App#registry_type}
	RegistryType *string `field:"required" json:"registryType" yaml:"registryType"`
	// The repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#repository App#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// deploy_on_push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#deploy_on_push App#deploy_on_push}
	DeployOnPush interface{} `field:"optional" json:"deployOnPush" yaml:"deployOnPush"`
	// The image digest. Cannot be specified if tag is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#digest App#digest}
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// The registry name. Must be left empty for the DOCR registry type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#registry App#registry}
	Registry *string `field:"optional" json:"registry" yaml:"registry"`
	// Access credentials for third-party registries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#registry_credentials App#registry_credentials}
	RegistryCredentials *string `field:"optional" json:"registryCredentials" yaml:"registryCredentials"`
	// The repository tag. Defaults to latest if not provided. Cannot be specified if digest is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/app#tag App#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

