// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecStaticSiteCors struct {
	// Whether browsers should expose the response to the client-side JavaScript code when the request’s credentials mode is `include`.
	//
	// This configures the Access-Control-Allow-Credentials header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/app#allow_credentials App#allow_credentials}
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// The set of allowed HTTP request headers. This configures the Access-Control-Allow-Headers header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/app#allow_headers App#allow_headers}
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// The set of allowed HTTP methods. This configures the Access-Control-Allow-Methods header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/app#allow_methods App#allow_methods}
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// allow_origins block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/app#allow_origins App#allow_origins}
	AllowOrigins *AppSpecStaticSiteCorsAllowOrigins `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// The set of HTTP response headers that browsers are allowed to access. This configures the Access-Control-Expose-Headers header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/app#expose_headers App#expose_headers}
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// An optional duration specifying how long browsers can cache the results of a preflight request.
	//
	// This configures the Access-Control-Max-Age header. Example: `5h30m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.51.0/docs/resources/app#max_age App#max_age}
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

