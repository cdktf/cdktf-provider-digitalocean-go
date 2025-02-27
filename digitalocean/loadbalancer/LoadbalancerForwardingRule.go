// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerForwardingRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#entry_port Loadbalancer#entry_port}.
	EntryPort *float64 `field:"required" json:"entryPort" yaml:"entryPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#entry_protocol Loadbalancer#entry_protocol}.
	EntryProtocol *string `field:"required" json:"entryProtocol" yaml:"entryProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#target_port Loadbalancer#target_port}.
	TargetPort *float64 `field:"required" json:"targetPort" yaml:"targetPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#target_protocol Loadbalancer#target_protocol}.
	TargetProtocol *string `field:"required" json:"targetProtocol" yaml:"targetProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#certificate_id Loadbalancer#certificate_id}.
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#certificate_name Loadbalancer#certificate_name}.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/loadbalancer#tls_passthrough Loadbalancer#tls_passthrough}.
	TlsPassthrough interface{} `field:"optional" json:"tlsPassthrough" yaml:"tlsPassthrough"`
}

