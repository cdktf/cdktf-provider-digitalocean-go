// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.loadbalancer.Loadbalancer",
		reflect.TypeOf((*Loadbalancer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "algorithm", GoGetter: "Algorithm"},
			_jsii_.MemberProperty{JsiiProperty: "algorithmInput", GoGetter: "AlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableLetsEncryptDnsRecords", GoGetter: "DisableLetsEncryptDnsRecords"},
			_jsii_.MemberProperty{JsiiProperty: "disableLetsEncryptDnsRecordsInput", GoGetter: "DisableLetsEncryptDnsRecordsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dropletIds", GoGetter: "DropletIds"},
			_jsii_.MemberProperty{JsiiProperty: "dropletIdsInput", GoGetter: "DropletIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dropletTag", GoGetter: "DropletTag"},
			_jsii_.MemberProperty{JsiiProperty: "dropletTagInput", GoGetter: "DropletTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableBackendKeepalive", GoGetter: "EnableBackendKeepalive"},
			_jsii_.MemberProperty{JsiiProperty: "enableBackendKeepaliveInput", GoGetter: "EnableBackendKeepaliveInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableProxyProtocol", GoGetter: "EnableProxyProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "enableProxyProtocolInput", GoGetter: "EnableProxyProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "firewall", GoGetter: "Firewall"},
			_jsii_.MemberProperty{JsiiProperty: "firewallInput", GoGetter: "FirewallInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "forwardingRule", GoGetter: "ForwardingRule"},
			_jsii_.MemberProperty{JsiiProperty: "forwardingRuleInput", GoGetter: "ForwardingRuleInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthcheck", GoGetter: "Healthcheck"},
			_jsii_.MemberProperty{JsiiProperty: "healthcheckInput", GoGetter: "HealthcheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpIdleTimeoutSeconds", GoGetter: "HttpIdleTimeoutSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "httpIdleTimeoutSecondsInput", GoGetter: "HttpIdleTimeoutSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putFirewall", GoMethod: "PutFirewall"},
			_jsii_.MemberMethod{JsiiMethod: "putForwardingRule", GoMethod: "PutForwardingRule"},
			_jsii_.MemberMethod{JsiiMethod: "putHealthcheck", GoMethod: "PutHealthcheck"},
			_jsii_.MemberMethod{JsiiMethod: "putStickySessions", GoMethod: "PutStickySessions"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redirectHttpToHttps", GoGetter: "RedirectHttpToHttps"},
			_jsii_.MemberProperty{JsiiProperty: "redirectHttpToHttpsInput", GoGetter: "RedirectHttpToHttpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlgorithm", GoMethod: "ResetAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableLetsEncryptDnsRecords", GoMethod: "ResetDisableLetsEncryptDnsRecords"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropletIds", GoMethod: "ResetDropletIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropletTag", GoMethod: "ResetDropletTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableBackendKeepalive", GoMethod: "ResetEnableBackendKeepalive"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableProxyProtocol", GoMethod: "ResetEnableProxyProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirewall", GoMethod: "ResetFirewall"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthcheck", GoMethod: "ResetHealthcheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpIdleTimeoutSeconds", GoMethod: "ResetHttpIdleTimeoutSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectHttpToHttps", GoMethod: "ResetRedirectHttpToHttps"},
			_jsii_.MemberMethod{JsiiMethod: "resetSize", GoMethod: "ResetSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizeUnit", GoMethod: "ResetSizeUnit"},
			_jsii_.MemberMethod{JsiiMethod: "resetStickySessions", GoMethod: "ResetStickySessions"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcUuid", GoMethod: "ResetVpcUuid"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInput", GoGetter: "SizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sizeUnit", GoGetter: "SizeUnit"},
			_jsii_.MemberProperty{JsiiProperty: "sizeUnitInput", GoGetter: "SizeUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "stickySessions", GoGetter: "StickySessions"},
			_jsii_.MemberProperty{JsiiProperty: "stickySessionsInput", GoGetter: "StickySessionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "urn", GoGetter: "Urn"},
			_jsii_.MemberProperty{JsiiProperty: "vpcUuid", GoGetter: "VpcUuid"},
			_jsii_.MemberProperty{JsiiProperty: "vpcUuidInput", GoGetter: "VpcUuidInput"},
		},
		func() interface{} {
			j := jsiiProxy_Loadbalancer{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerConfig",
		reflect.TypeOf((*LoadbalancerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerFirewall",
		reflect.TypeOf((*LoadbalancerFirewall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerFirewallOutputReference",
		reflect.TypeOf((*LoadbalancerFirewallOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allow", GoGetter: "Allow"},
			_jsii_.MemberProperty{JsiiProperty: "allowInput", GoGetter: "AllowInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deny", GoGetter: "Deny"},
			_jsii_.MemberProperty{JsiiProperty: "denyInput", GoGetter: "DenyInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllow", GoMethod: "ResetAllow"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeny", GoMethod: "ResetDeny"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerFirewallOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerForwardingRule",
		reflect.TypeOf((*LoadbalancerForwardingRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerForwardingRuleList",
		reflect.TypeOf((*LoadbalancerForwardingRuleList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerForwardingRuleList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerForwardingRuleOutputReference",
		reflect.TypeOf((*LoadbalancerForwardingRuleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateId", GoGetter: "CertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "certificateIdInput", GoGetter: "CertificateIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateName", GoGetter: "CertificateName"},
			_jsii_.MemberProperty{JsiiProperty: "certificateNameInput", GoGetter: "CertificateNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "entryPort", GoGetter: "EntryPort"},
			_jsii_.MemberProperty{JsiiProperty: "entryPortInput", GoGetter: "EntryPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "entryProtocol", GoGetter: "EntryProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "entryProtocolInput", GoGetter: "EntryProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateId", GoMethod: "ResetCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateName", GoMethod: "ResetCertificateName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsPassthrough", GoMethod: "ResetTlsPassthrough"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetPort", GoGetter: "TargetPort"},
			_jsii_.MemberProperty{JsiiProperty: "targetPortInput", GoGetter: "TargetPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetProtocol", GoGetter: "TargetProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "targetProtocolInput", GoGetter: "TargetProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tlsPassthrough", GoGetter: "TlsPassthrough"},
			_jsii_.MemberProperty{JsiiProperty: "tlsPassthroughInput", GoGetter: "TlsPassthroughInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerForwardingRuleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerHealthcheck",
		reflect.TypeOf((*LoadbalancerHealthcheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerHealthcheckOutputReference",
		reflect.TypeOf((*LoadbalancerHealthcheckOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "checkIntervalSeconds", GoGetter: "CheckIntervalSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "checkIntervalSecondsInput", GoGetter: "CheckIntervalSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThreshold", GoGetter: "HealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThresholdInput", GoGetter: "HealthyThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckIntervalSeconds", GoMethod: "ResetCheckIntervalSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthyThreshold", GoMethod: "ResetHealthyThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseTimeoutSeconds", GoMethod: "ResetResponseTimeoutSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnhealthyThreshold", GoMethod: "ResetUnhealthyThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "responseTimeoutSeconds", GoGetter: "ResponseTimeoutSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "responseTimeoutSecondsInput", GoGetter: "ResponseTimeoutSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThreshold", GoGetter: "UnhealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThresholdInput", GoGetter: "UnhealthyThresholdInput"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerHealthcheckOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerStickySessions",
		reflect.TypeOf((*LoadbalancerStickySessions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.loadbalancer.LoadbalancerStickySessionsOutputReference",
		reflect.TypeOf((*LoadbalancerStickySessionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cookieName", GoGetter: "CookieName"},
			_jsii_.MemberProperty{JsiiProperty: "cookieNameInput", GoGetter: "CookieNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "cookieTtlSeconds", GoGetter: "CookieTtlSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "cookieTtlSecondsInput", GoGetter: "CookieTtlSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookieName", GoMethod: "ResetCookieName"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookieTtlSeconds", GoMethod: "ResetCookieTtlSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerStickySessionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
