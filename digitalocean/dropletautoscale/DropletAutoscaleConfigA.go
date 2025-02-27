// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dropletautoscale


type DropletAutoscaleConfigA struct {
	// Cooldown duration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/droplet_autoscale#cooldown_minutes DropletAutoscale#cooldown_minutes}
	CooldownMinutes *float64 `field:"optional" json:"cooldownMinutes" yaml:"cooldownMinutes"`
	// Max number of members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/droplet_autoscale#max_instances DropletAutoscale#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Min number of members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/droplet_autoscale#min_instances DropletAutoscale#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// CPU target threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/droplet_autoscale#target_cpu_utilization DropletAutoscale#target_cpu_utilization}
	TargetCpuUtilization *float64 `field:"optional" json:"targetCpuUtilization" yaml:"targetCpuUtilization"`
	// Memory target threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/droplet_autoscale#target_memory_utilization DropletAutoscale#target_memory_utilization}
	TargetMemoryUtilization *float64 `field:"optional" json:"targetMemoryUtilization" yaml:"targetMemoryUtilization"`
	// Target number of members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.1/docs/resources/droplet_autoscale#target_number_instances DropletAutoscale#target_number_instances}
	TargetNumberInstances *float64 `field:"optional" json:"targetNumberInstances" yaml:"targetNumberInstances"`
}

