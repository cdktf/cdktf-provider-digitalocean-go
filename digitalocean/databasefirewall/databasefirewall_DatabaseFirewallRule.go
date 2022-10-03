package databasefirewall


type DatabaseFirewallRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/database_firewall#type DatabaseFirewall#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/database_firewall#value DatabaseFirewall#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

