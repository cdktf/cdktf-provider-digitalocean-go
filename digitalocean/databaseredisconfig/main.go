// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databaseredisconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfig",
		reflect.TypeOf((*DatabaseRedisConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aclChannelsDefault", GoGetter: "AclChannelsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "aclChannelsDefaultInput", GoGetter: "AclChannelsDefaultInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdInput", GoGetter: "ClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ioThreads", GoGetter: "IoThreads"},
			_jsii_.MemberProperty{JsiiProperty: "ioThreadsInput", GoGetter: "IoThreadsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lfuDecayTime", GoGetter: "LfuDecayTime"},
			_jsii_.MemberProperty{JsiiProperty: "lfuDecayTimeInput", GoGetter: "LfuDecayTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lfuLogFactor", GoGetter: "LfuLogFactor"},
			_jsii_.MemberProperty{JsiiProperty: "lfuLogFactorInput", GoGetter: "LfuLogFactorInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxmemoryPolicy", GoGetter: "MaxmemoryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "maxmemoryPolicyInput", GoGetter: "MaxmemoryPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifyKeyspaceEvents", GoGetter: "NotifyKeyspaceEvents"},
			_jsii_.MemberProperty{JsiiProperty: "notifyKeyspaceEventsInput", GoGetter: "NotifyKeyspaceEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfDatabases", GoGetter: "NumberOfDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfDatabasesInput", GoGetter: "NumberOfDatabasesInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "persistence", GoGetter: "Persistence"},
			_jsii_.MemberProperty{JsiiProperty: "persistenceInput", GoGetter: "PersistenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pubsubClientOutputBufferLimit", GoGetter: "PubsubClientOutputBufferLimit"},
			_jsii_.MemberProperty{JsiiProperty: "pubsubClientOutputBufferLimitInput", GoGetter: "PubsubClientOutputBufferLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAclChannelsDefault", GoMethod: "ResetAclChannelsDefault"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIoThreads", GoMethod: "ResetIoThreads"},
			_jsii_.MemberMethod{JsiiMethod: "resetLfuDecayTime", GoMethod: "ResetLfuDecayTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetLfuLogFactor", GoMethod: "ResetLfuLogFactor"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxmemoryPolicy", GoMethod: "ResetMaxmemoryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifyKeyspaceEvents", GoMethod: "ResetNotifyKeyspaceEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumberOfDatabases", GoMethod: "ResetNumberOfDatabases"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPersistence", GoMethod: "ResetPersistence"},
			_jsii_.MemberMethod{JsiiMethod: "resetPubsubClientOutputBufferLimit", GoMethod: "ResetPubsubClientOutputBufferLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsl", GoMethod: "ResetSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "ssl", GoGetter: "Ssl"},
			_jsii_.MemberProperty{JsiiProperty: "sslInput", GoGetter: "SslInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DatabaseRedisConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-digitalocean.databaseRedisConfig.DatabaseRedisConfigConfig",
		reflect.TypeOf((*DatabaseRedisConfigConfig)(nil)).Elem(),
	)
}
