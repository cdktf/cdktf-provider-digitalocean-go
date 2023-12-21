// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterMaintenancePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/kubernetes_cluster#day KubernetesCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.34.1/docs/resources/kubernetes_cluster#start_time KubernetesCluster#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

