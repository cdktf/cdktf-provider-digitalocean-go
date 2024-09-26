// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterNodePool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#size KubernetesCluster#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#auto_scale KubernetesCluster#auto_scale}.
	AutoScale interface{} `field:"optional" json:"autoScale" yaml:"autoScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#labels KubernetesCluster#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#max_nodes KubernetesCluster#max_nodes}.
	MaxNodes *float64 `field:"optional" json:"maxNodes" yaml:"maxNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#min_nodes KubernetesCluster#min_nodes}.
	MinNodes *float64 `field:"optional" json:"minNodes" yaml:"minNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#node_count KubernetesCluster#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.42.0/docs/resources/kubernetes_cluster#taint KubernetesCluster#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
}

