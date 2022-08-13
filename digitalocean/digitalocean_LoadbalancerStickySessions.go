// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type LoadbalancerStickySessions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer#cookie_name Loadbalancer#cookie_name}.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer#cookie_ttl_seconds Loadbalancer#cookie_ttl_seconds}.
	CookieTtlSeconds *float64 `field:"optional" json:"cookieTtlSeconds" yaml:"cookieTtlSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer#type Loadbalancer#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

