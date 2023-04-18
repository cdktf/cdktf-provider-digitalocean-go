package app


type AppSpec struct {
	// The name of the app. Must be unique across all apps in the same account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#name App#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// alert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#alert App#alert}
	Alert interface{} `field:"optional" json:"alert" yaml:"alert"`
	// database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#database App#database}
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#domain App#domain}
	Domain interface{} `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#domains App#domains}.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#env App#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#function App#function}
	Function interface{} `field:"optional" json:"function" yaml:"function"`
	// job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#job App#job}
	Job interface{} `field:"optional" json:"job" yaml:"job"`
	// The slug for the DigitalOcean data center region hosting the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#region App#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#service App#service}
	Service interface{} `field:"optional" json:"service" yaml:"service"`
	// static_site block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#static_site App#static_site}
	StaticSite interface{} `field:"optional" json:"staticSite" yaml:"staticSite"`
	// worker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.27.1/docs/resources/app#worker App#worker}
	Worker interface{} `field:"optional" json:"worker" yaml:"worker"`
}

