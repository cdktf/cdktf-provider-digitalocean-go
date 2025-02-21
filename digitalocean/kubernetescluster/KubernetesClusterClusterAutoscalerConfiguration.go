// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterClusterAutoscalerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/kubernetes_cluster#scale_down_unneeded_time KubernetesCluster#scale_down_unneeded_time}.
	ScaleDownUnneededTime *string `field:"optional" json:"scaleDownUnneededTime" yaml:"scaleDownUnneededTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.49.0/docs/resources/kubernetes_cluster#scale_down_utilization_threshold KubernetesCluster#scale_down_utilization_threshold}.
	ScaleDownUtilizationThreshold *float64 `field:"optional" json:"scaleDownUtilizationThreshold" yaml:"scaleDownUtilizationThreshold"`
}

