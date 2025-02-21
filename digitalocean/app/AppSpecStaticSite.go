// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecStaticSite struct {
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// bitbucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#bitbucket App#bitbucket}
	Bitbucket *AppSpecStaticSiteBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// An optional build command to run while building this component from source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#build_command App#build_command}
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The name of the document to use as the fallback for any requests to documents that are not found when serving this static site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#catchall_document App#catchall_document}
	CatchallDocument *string `field:"optional" json:"catchallDocument" yaml:"catchallDocument"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#cors App#cors}
	Cors *AppSpecStaticSiteCors `field:"optional" json:"cors" yaml:"cors"`
	// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#dockerfile_path App#dockerfile_path}
	DockerfilePath *string `field:"optional" json:"dockerfilePath" yaml:"dockerfilePath"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#env App#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// An environment slug describing the type of this app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#environment_slug App#environment_slug}
	EnvironmentSlug *string `field:"optional" json:"environmentSlug" yaml:"environmentSlug"`
	// The name of the error document to use when serving this static site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#error_document App#error_document}
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#git App#git}
	Git *AppSpecStaticSiteGit `field:"optional" json:"git" yaml:"git"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#github App#github}
	Github *AppSpecStaticSiteGithub `field:"optional" json:"github" yaml:"github"`
	// gitlab block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#gitlab App#gitlab}
	Gitlab *AppSpecStaticSiteGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// The name of the index document to use when serving this static site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#index_document App#index_document}
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// An optional path to where the built assets will be located, relative to the build context.
	//
	// If not set, App Platform will automatically scan for these directory names: `_static`, `dist`, `public`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#output_dir App#output_dir}
	OutputDir *string `field:"optional" json:"outputDir" yaml:"outputDir"`
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#routes App#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
	// An optional path to the working directory to use for the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/app#source_dir App#source_dir}
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
}

