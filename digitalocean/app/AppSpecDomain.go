// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecDomain struct {
	// The hostname for the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.57.0/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.57.0/docs/resources/app#type App#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Indicates whether the domain includes all sub-domains, in addition to the given domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.57.0/docs/resources/app#wildcard App#wildcard}
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
	// If the domain uses DigitalOcean DNS and you would like App Platform to automatically manage it for you, set this to the name of the domain on your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.57.0/docs/resources/app#zone App#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

