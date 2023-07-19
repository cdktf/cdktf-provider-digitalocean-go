package kubernetesnodepool


type KubernetesNodePoolTaint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/kubernetes_node_pool#effect KubernetesNodePool#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/kubernetes_node_pool#key KubernetesNodePool#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.29.0/docs/resources/kubernetes_node_pool#value KubernetesNodePool#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

