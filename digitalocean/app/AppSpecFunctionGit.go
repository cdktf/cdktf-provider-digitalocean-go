// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecFunctionGit struct {
	// The name of the branch to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/app#branch App#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The clone URL of the repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.37.1/docs/resources/app#repo_clone_url App#repo_clone_url}
	RepoCloneUrl *string `field:"optional" json:"repoCloneUrl" yaml:"repoCloneUrl"`
}

