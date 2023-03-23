package loadbalancer


type LoadbalancerFirewall struct {
	// the rules for ALLOWING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer#allow Loadbalancer#allow}
	Allow *[]*string `field:"optional" json:"allow" yaml:"allow"`
	// the rules for DENYING traffic to the LB (strings in the form: 'ip:1.2.3.4' or 'cidr:1.2.0.0/16').
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/loadbalancer#deny Loadbalancer#deny}
	Deny *[]*string `field:"optional" json:"deny" yaml:"deny"`
}

