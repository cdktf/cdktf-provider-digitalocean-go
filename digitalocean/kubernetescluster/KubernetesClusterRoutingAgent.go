// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterRoutingAgent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.52.0/docs/resources/kubernetes_cluster#enabled KubernetesCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

