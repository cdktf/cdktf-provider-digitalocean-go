// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterControlPlaneFirewall struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/kubernetes_cluster#allowed_addresses KubernetesCluster#allowed_addresses}.
	AllowedAddresses *[]*string `field:"required" json:"allowedAddresses" yaml:"allowedAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.48.1/docs/resources/kubernetes_cluster#enabled KubernetesCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

