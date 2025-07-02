// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJobGithub struct {
	// The name of the branch to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/app#branch App#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Whether to automatically deploy new commits made to the repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/app#deploy_on_push App#deploy_on_push}
	DeployOnPush interface{} `field:"optional" json:"deployOnPush" yaml:"deployOnPush"`
	// The name of the repo in the format `owner/repo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.58.0/docs/resources/app#repo App#repo}
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

