// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerHealthcheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#port Loadbalancer#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#protocol Loadbalancer#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#check_interval_seconds Loadbalancer#check_interval_seconds}.
	CheckIntervalSeconds *float64 `field:"optional" json:"checkIntervalSeconds" yaml:"checkIntervalSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#healthy_threshold Loadbalancer#healthy_threshold}.
	HealthyThreshold *float64 `field:"optional" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#path Loadbalancer#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#response_timeout_seconds Loadbalancer#response_timeout_seconds}.
	ResponseTimeoutSeconds *float64 `field:"optional" json:"responseTimeoutSeconds" yaml:"responseTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/loadbalancer#unhealthy_threshold Loadbalancer#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"optional" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

