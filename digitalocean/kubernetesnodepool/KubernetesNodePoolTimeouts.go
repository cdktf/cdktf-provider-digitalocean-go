package kubernetesnodepool


type KubernetesNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_node_pool#create KubernetesNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_node_pool#delete KubernetesNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

