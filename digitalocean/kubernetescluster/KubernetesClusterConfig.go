// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#name KubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#node_pool KubernetesCluster#node_pool}
	NodePool *KubernetesClusterNodePool `field:"required" json:"nodePool" yaml:"nodePool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#region KubernetesCluster#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#version KubernetesCluster#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#auto_upgrade KubernetesCluster#auto_upgrade}.
	AutoUpgrade interface{} `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#cluster_subnet KubernetesCluster#cluster_subnet}.
	ClusterSubnet *string `field:"optional" json:"clusterSubnet" yaml:"clusterSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#destroy_all_associated_resources KubernetesCluster#destroy_all_associated_resources}.
	DestroyAllAssociatedResources interface{} `field:"optional" json:"destroyAllAssociatedResources" yaml:"destroyAllAssociatedResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#ha KubernetesCluster#ha}.
	Ha interface{} `field:"optional" json:"ha" yaml:"ha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#id KubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#maintenance_policy KubernetesCluster#maintenance_policy}
	MaintenancePolicy *KubernetesClusterMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#registry_integration KubernetesCluster#registry_integration}.
	RegistryIntegration interface{} `field:"optional" json:"registryIntegration" yaml:"registryIntegration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#service_subnet KubernetesCluster#service_subnet}.
	ServiceSubnet *string `field:"optional" json:"serviceSubnet" yaml:"serviceSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#surge_upgrade KubernetesCluster#surge_upgrade}.
	SurgeUpgrade interface{} `field:"optional" json:"surgeUpgrade" yaml:"surgeUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#tags KubernetesCluster#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#timeouts KubernetesCluster#timeouts}
	Timeouts *KubernetesClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.46.0/docs/resources/kubernetes_cluster#vpc_uuid KubernetesCluster#vpc_uuid}.
	VpcUuid *string `field:"optional" json:"vpcUuid" yaml:"vpcUuid"`
}

