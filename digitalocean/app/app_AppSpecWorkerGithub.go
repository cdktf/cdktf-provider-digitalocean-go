package app


type AppSpecWorkerGithub struct {
	// The name of the branch to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#branch App#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Whether to automatically deploy new commits made to the repo.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#deploy_on_push App#deploy_on_push}
	DeployOnPush interface{} `field:"optional" json:"deployOnPush" yaml:"deployOnPush"`
	// The name of the repo in the format `owner/repo`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#repo App#repo}
	Repo *string `field:"optional" json:"repo" yaml:"repo"`
}

