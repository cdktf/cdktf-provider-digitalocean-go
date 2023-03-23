package app


type AppSpecServiceGit struct {
	// The name of the branch to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#branch App#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The clone URL of the repo.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#repo_clone_url App#repo_clone_url}
	RepoCloneUrl *string `field:"optional" json:"repoCloneUrl" yaml:"repoCloneUrl"`
}

