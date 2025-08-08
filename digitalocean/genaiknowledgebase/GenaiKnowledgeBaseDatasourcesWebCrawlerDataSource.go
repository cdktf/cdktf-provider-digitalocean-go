// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package genaiknowledgebase


type GenaiKnowledgeBaseDatasourcesWebCrawlerDataSource struct {
	// The base URL to crawl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#base_url GenaiKnowledgeBase#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Options for specifying how URLs found on pages should be handled.
	//
	// - UNKNOWN: Default unknown value
	// - SCOPED: Only include the base URL.
	// - PATH: Crawl the base URL and linked pages within the URL path.
	// - DOMAIN: Crawl the base URL and linked pages within the same domain.
	// - SUBDOMAINS: Crawl the base URL and linked pages for any subdomain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#crawling_option GenaiKnowledgeBase#crawling_option}
	CrawlingOption *string `field:"optional" json:"crawlingOption" yaml:"crawlingOption"`
	// Whether to embed media content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/digitalocean/digitalocean/2.63.0/docs/resources/genai_knowledge_base#embed_media GenaiKnowledgeBase#embed_media}
	EmbedMedia interface{} `field:"optional" json:"embedMedia" yaml:"embedMedia"`
}

