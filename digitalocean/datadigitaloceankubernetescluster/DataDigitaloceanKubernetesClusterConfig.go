// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadigitaloceankubernetescluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDigitaloceanKubernetesClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#name DataDigitaloceanKubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// amd_gpu_device_metrics_exporter_plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#amd_gpu_device_metrics_exporter_plugin DataDigitaloceanKubernetesCluster#amd_gpu_device_metrics_exporter_plugin}
	AmdGpuDeviceMetricsExporterPlugin *DataDigitaloceanKubernetesClusterAmdGpuDeviceMetricsExporterPlugin `field:"optional" json:"amdGpuDeviceMetricsExporterPlugin" yaml:"amdGpuDeviceMetricsExporterPlugin"`
	// amd_gpu_device_plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#amd_gpu_device_plugin DataDigitaloceanKubernetesCluster#amd_gpu_device_plugin}
	AmdGpuDevicePlugin *DataDigitaloceanKubernetesClusterAmdGpuDevicePlugin `field:"optional" json:"amdGpuDevicePlugin" yaml:"amdGpuDevicePlugin"`
	// cluster_autoscaler_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#cluster_autoscaler_configuration DataDigitaloceanKubernetesCluster#cluster_autoscaler_configuration}
	ClusterAutoscalerConfiguration interface{} `field:"optional" json:"clusterAutoscalerConfiguration" yaml:"clusterAutoscalerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#id DataDigitaloceanKubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#kubeconfig_expire_seconds DataDigitaloceanKubernetesCluster#kubeconfig_expire_seconds}.
	KubeconfigExpireSeconds *float64 `field:"optional" json:"kubeconfigExpireSeconds" yaml:"kubeconfigExpireSeconds"`
	// routing_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#routing_agent DataDigitaloceanKubernetesCluster#routing_agent}
	RoutingAgent *DataDigitaloceanKubernetesClusterRoutingAgent `field:"optional" json:"routingAgent" yaml:"routingAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.66.0/docs/data-sources/kubernetes_cluster#tags DataDigitaloceanKubernetesCluster#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

