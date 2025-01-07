// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerGlbSettings struct {
	// target port rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/loadbalancer#target_port Loadbalancer#target_port}
	TargetPort *float64 `field:"required" json:"targetPort" yaml:"targetPort"`
	// target protocol rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/loadbalancer#target_protocol Loadbalancer#target_protocol}
	TargetProtocol *string `field:"required" json:"targetProtocol" yaml:"targetProtocol"`
	// cdn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/loadbalancer#cdn Loadbalancer#cdn}
	Cdn *LoadbalancerGlbSettingsCdn `field:"optional" json:"cdn" yaml:"cdn"`
	// fail-over threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/loadbalancer#failover_threshold Loadbalancer#failover_threshold}
	FailoverThreshold *float64 `field:"optional" json:"failoverThreshold" yaml:"failoverThreshold"`
	// region priority map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.47.0/docs/resources/loadbalancer#region_priorities Loadbalancer#region_priorities}
	RegionPriorities *map[string]*float64 `field:"optional" json:"regionPriorities" yaml:"regionPriorities"`
}

