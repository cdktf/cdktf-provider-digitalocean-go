// Prebuilt digitalocean Provider for Terraform CDK (cdktf)
package digitalocean


type KubernetesClusterNodePoolTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_cluster#effect KubernetesCluster#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_cluster#key KubernetesCluster#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_cluster#value KubernetesCluster#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

