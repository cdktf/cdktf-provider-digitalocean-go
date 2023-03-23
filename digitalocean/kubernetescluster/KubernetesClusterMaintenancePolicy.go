package kubernetescluster


type KubernetesClusterMaintenancePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_cluster#day KubernetesCluster#day}.
	Day *string `field:"optional" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/kubernetes_cluster#start_time KubernetesCluster#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

