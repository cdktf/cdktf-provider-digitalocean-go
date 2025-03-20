// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerDomains struct {
	// domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.2/docs/resources/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// name of certificate required for TLS handshaking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.2/docs/resources/loadbalancer#certificate_name Loadbalancer#certificate_name}
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// flag indicating if domain is managed by DigitalOcean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.2/docs/resources/loadbalancer#is_managed Loadbalancer#is_managed}
	IsManaged interface{} `field:"optional" json:"isManaged" yaml:"isManaged"`
}

