// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.64.0/docs/resources/kubernetes_cluster#create KubernetesCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

