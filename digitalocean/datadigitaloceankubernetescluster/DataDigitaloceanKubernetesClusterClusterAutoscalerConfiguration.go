// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceankubernetescluster


type DataDigitaloceanKubernetesClusterClusterAutoscalerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/kubernetes_cluster#expanders DataDigitaloceanKubernetesCluster#expanders}.
	Expanders *[]*string `field:"optional" json:"expanders" yaml:"expanders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/kubernetes_cluster#scale_down_unneeded_time DataDigitaloceanKubernetesCluster#scale_down_unneeded_time}.
	ScaleDownUnneededTime *string `field:"optional" json:"scaleDownUnneededTime" yaml:"scaleDownUnneededTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.69.0/docs/data-sources/kubernetes_cluster#scale_down_utilization_threshold DataDigitaloceanKubernetesCluster#scale_down_utilization_threshold}.
	ScaleDownUtilizationThreshold *float64 `field:"optional" json:"scaleDownUtilizationThreshold" yaml:"scaleDownUtilizationThreshold"`
}

