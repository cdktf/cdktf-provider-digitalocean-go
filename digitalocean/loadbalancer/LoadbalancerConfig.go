// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// forwarding_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#forwarding_rule Loadbalancer#forwarding_rule}
	ForwardingRule interface{} `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#name Loadbalancer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#region Loadbalancer#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#algorithm Loadbalancer#algorithm}.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#disable_lets_encrypt_dns_records Loadbalancer#disable_lets_encrypt_dns_records}.
	DisableLetsEncryptDnsRecords interface{} `field:"optional" json:"disableLetsEncryptDnsRecords" yaml:"disableLetsEncryptDnsRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#droplet_ids Loadbalancer#droplet_ids}.
	DropletIds *[]*float64 `field:"optional" json:"dropletIds" yaml:"dropletIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#droplet_tag Loadbalancer#droplet_tag}.
	DropletTag *string `field:"optional" json:"dropletTag" yaml:"dropletTag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#enable_backend_keepalive Loadbalancer#enable_backend_keepalive}.
	EnableBackendKeepalive interface{} `field:"optional" json:"enableBackendKeepalive" yaml:"enableBackendKeepalive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#enable_proxy_protocol Loadbalancer#enable_proxy_protocol}.
	EnableProxyProtocol interface{} `field:"optional" json:"enableProxyProtocol" yaml:"enableProxyProtocol"`
	// firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#firewall Loadbalancer#firewall}
	Firewall *LoadbalancerFirewall `field:"optional" json:"firewall" yaml:"firewall"`
	// healthcheck block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#healthcheck Loadbalancer#healthcheck}
	Healthcheck *LoadbalancerHealthcheck `field:"optional" json:"healthcheck" yaml:"healthcheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#http_idle_timeout_seconds Loadbalancer#http_idle_timeout_seconds}.
	HttpIdleTimeoutSeconds *float64 `field:"optional" json:"httpIdleTimeoutSeconds" yaml:"httpIdleTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#id Loadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#project_id Loadbalancer#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#redirect_http_to_https Loadbalancer#redirect_http_to_https}.
	RedirectHttpToHttps interface{} `field:"optional" json:"redirectHttpToHttps" yaml:"redirectHttpToHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#size Loadbalancer#size}.
	Size *string `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#size_unit Loadbalancer#size_unit}.
	SizeUnit *float64 `field:"optional" json:"sizeUnit" yaml:"sizeUnit"`
	// sticky_sessions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#sticky_sessions Loadbalancer#sticky_sessions}
	StickySessions *LoadbalancerStickySessions `field:"optional" json:"stickySessions" yaml:"stickySessions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/loadbalancer#vpc_uuid Loadbalancer#vpc_uuid}.
	VpcUuid *string `field:"optional" json:"vpcUuid" yaml:"vpcUuid"`
}

