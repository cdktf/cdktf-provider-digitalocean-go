package app


type AppSpecFunctionAlert struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#operator App#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#rule App#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#value App#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#window App#window}.
	Window *string `field:"required" json:"window" yaml:"window"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/app#disabled App#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

