// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecEnv struct {
	// The name of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/app#key App#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The visibility scope of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/app#scope App#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The type of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/app#type App#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value of the environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/app#value App#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

