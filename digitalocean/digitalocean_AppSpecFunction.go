// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type AppSpecFunction struct {
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// alert block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#alert App#alert}
	Alert interface{} `field:"optional" json:"alert" yaml:"alert"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#cors App#cors}
	Cors *AppSpecFunctionCors `field:"optional" json:"cors" yaml:"cors"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#env App#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// git block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#git App#git}
	Git *AppSpecFunctionGit `field:"optional" json:"git" yaml:"git"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#github App#github}
	Github *AppSpecFunctionGithub `field:"optional" json:"github" yaml:"github"`
	// gitlab block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#gitlab App#gitlab}
	Gitlab *AppSpecFunctionGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	// log_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#log_destination App#log_destination}
	LogDestination interface{} `field:"optional" json:"logDestination" yaml:"logDestination"`
	// routes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#routes App#routes}
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
	// An optional path to the working directory to use for the build.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#source_dir App#source_dir}
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
}

