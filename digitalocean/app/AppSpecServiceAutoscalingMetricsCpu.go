// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package app


type AppSpecServiceAutoscalingMetricsCpu struct {
	// The average target CPU utilization for the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.43.0/docs/resources/app#percent App#percent}
	Percent *float64 `field:"required" json:"percent" yaml:"percent"`
}

