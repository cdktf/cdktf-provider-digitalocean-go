// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecJob struct {
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// alert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#alert App#alert}
	Alert interface{} `field:"optional" json:"alert" yaml:"alert"`
	// bitbucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#bitbucket App#bitbucket}
	Bitbucket *AppSpecJobBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// An optional build command to run while building this component from source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#build_command App#build_command}
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#dockerfile_path App#dockerfile_path}
	DockerfilePath *string `field:"optional" json:"dockerfilePath" yaml:"dockerfilePath"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#env App#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// An environment slug describing the type of this app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#environment_slug App#environment_slug}
	EnvironmentSlug *string `field:"optional" json:"environmentSlug" yaml:"environmentSlug"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#git App#git}
	Git *AppSpecJobGit `field:"optional" json:"git" yaml:"git"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#github App#github}
	Github *AppSpecJobGithub `field:"optional" json:"github" yaml:"github"`
	// gitlab block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#gitlab App#gitlab}
	Gitlab *AppSpecJobGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#image App#image}
	Image *AppSpecJobImage `field:"optional" json:"image" yaml:"image"`
	// The amount of instances that this component should be scaled to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#instance_count App#instance_count}
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The instance size to use for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#instance_size_slug App#instance_size_slug}
	InstanceSizeSlug *string `field:"optional" json:"instanceSizeSlug" yaml:"instanceSizeSlug"`
	// The type of job and when it will be run during the deployment process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#kind App#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#log_destination App#log_destination}
	LogDestination interface{} `field:"optional" json:"logDestination" yaml:"logDestination"`
	// An optional run command to override the component's default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#run_command App#run_command}
	RunCommand *string `field:"optional" json:"runCommand" yaml:"runCommand"`
	// An optional path to the working directory to use for the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#source_dir App#source_dir}
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
	// termination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/app#termination App#termination}
	Termination *AppSpecJobTermination `field:"optional" json:"termination" yaml:"termination"`
}

