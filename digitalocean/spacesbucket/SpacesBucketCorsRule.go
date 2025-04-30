// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spacesbucket


type SpacesBucketCorsRule struct {
	// A list of HTTP methods (e.g. GET) which are allowed from the specified origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.52.0/docs/resources/spaces_bucket#allowed_methods SpacesBucket#allowed_methods}
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// A list of hosts from which requests using the specified methods are allowed.
	//
	// A host may contain one wildcard (e.g. http://*.example.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.52.0/docs/resources/spaces_bucket#allowed_origins SpacesBucket#allowed_origins}
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// A list of headers that will be included in the CORS preflight request's Access-Control-Request-Headers.
	//
	// A header may contain one wildcard (e.g. x-amz-*).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.52.0/docs/resources/spaces_bucket#allowed_headers SpacesBucket#allowed_headers}
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.52.0/docs/resources/spaces_bucket#max_age_seconds SpacesBucket#max_age_seconds}.
	MaxAgeSeconds *float64 `field:"optional" json:"maxAgeSeconds" yaml:"maxAgeSeconds"`
}

