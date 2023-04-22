package databasecluster


type DatabaseClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.28.0/docs/resources/database_cluster#create DatabaseCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

