// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecServiceAutoscalingMetrics struct {
	// cpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/resources/app#cpu App#cpu}
	Cpu *AppSpecServiceAutoscalingMetricsCpu `field:"optional" json:"cpu" yaml:"cpu"`
}

