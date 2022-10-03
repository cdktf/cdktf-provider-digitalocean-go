package databasecluster


type DatabaseClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/digitalocean/r/database_cluster#create DatabaseCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

