package app


type AppSpecJobImageDeployOnPush struct {
	// Whether to automatically deploy images pushed to DOCR.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#enabled App#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

