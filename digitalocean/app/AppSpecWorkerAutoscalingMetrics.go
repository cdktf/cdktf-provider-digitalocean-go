// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecWorkerAutoscalingMetrics struct {
	// cpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/app#cpu App#cpu}
	Cpu *AppSpecWorkerAutoscalingMetricsCpu `field:"optional" json:"cpu" yaml:"cpu"`
}

