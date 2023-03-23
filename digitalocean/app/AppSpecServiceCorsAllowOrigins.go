package app


type AppSpecServiceCorsAllowOrigins struct {
	// Exact string match.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#exact App#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Prefix-based match.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#prefix App#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// RE2 style regex-based match.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#regex App#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

