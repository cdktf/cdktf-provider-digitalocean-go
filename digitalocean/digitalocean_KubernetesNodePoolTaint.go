// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type KubernetesNodePoolTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_node_pool#effect KubernetesNodePool#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_node_pool#key KubernetesNodePool#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_node_pool#value KubernetesNodePool#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

