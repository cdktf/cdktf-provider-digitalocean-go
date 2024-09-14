// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecFunctionRoutes struct {
	// Path specifies an route by HTTP path prefix.
	//
	// Paths must start with / and must be unique within the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/app#path App#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// An optional flag to preserve the path that is forwarded to the backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.41.0/docs/resources/app#preserve_path_prefix App#preserve_path_prefix}
	PreservePathPrefix interface{} `field:"optional" json:"preservePathPrefix" yaml:"preservePathPrefix"`
}

