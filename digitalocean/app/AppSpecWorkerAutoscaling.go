// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerAutoscaling struct {
	// The maximum amount of instances for this component. Must be more than min_instance_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/app#max_instance_count App#max_instance_count}
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/app#metrics App#metrics}
	Metrics *AppSpecWorkerAutoscalingMetrics `field:"required" json:"metrics" yaml:"metrics"`
	// The minimum amount of instances for this component. Must be less than max_instance_count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.67.0/docs/resources/app#min_instance_count App#min_instance_count}
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

