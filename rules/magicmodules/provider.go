// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Rules is a list of rules generated from Magic Modules
var Rules = []tflint.Rule{
	NewGoogleAccessContextManagerServicePerimeterInvalidPerimeterTypeRule(),
	NewGoogleActiveDirectoryDomainTrustInvalidTrustDirectionRule(),
	NewGoogleActiveDirectoryDomainTrustInvalidTrustTypeRule(),
	NewGoogleAppEngineDomainMappingInvalidOverrideStrategyRule(),
	NewGoogleAppEngineFirewallRuleInvalidActionRule(),
	NewGoogleAppEngineFlexibleAppVersionInvalidServingStatusRule(),
	NewGoogleBigqueryRoutineInvalidDeterminismLevelRule(),
	NewGoogleBigqueryRoutineInvalidLanguageRule(),
	NewGoogleBigqueryRoutineInvalidRoutineTypeRule(),
	NewGoogleBinaryAuthorizationPolicyInvalidGlobalPolicyEvaluationModeRule(),
	NewGoogleCloudAssetFolderFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetOrganizationFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetProjectFeedInvalidContentTypeRule(),
	NewGoogleCloudiotDeviceInvalidLogLevelRule(),
	NewGoogleCloudiotRegistryInvalidLogLevelRule(),
	NewGoogleComputeAddressInvalidAddressTypeRule(),
	NewGoogleComputeAddressInvalidNameRule(),
	NewGoogleComputeAddressInvalidNetworkTierRule(),
	NewGoogleComputeAddressInvalidPurposeRule(),
	NewGoogleComputeBackendBucketInvalidNameRule(),
	NewGoogleComputeBackendBucketSignedUrlKeyInvalidNameRule(),
	NewGoogleComputeBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeBackendServiceInvalidProtocolRule(),
	NewGoogleComputeBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeBackendServiceSignedUrlKeyInvalidNameRule(),
	NewGoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule(),
	NewGoogleComputeFirewallInvalidDirectionRule(),
	NewGoogleComputeForwardingRuleInvalidIpProtocolRule(),
	NewGoogleComputeForwardingRuleInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeForwardingRuleInvalidNetworkTierRule(),
	NewGoogleComputeGlobalAddressInvalidAddressTypeRule(),
	NewGoogleComputeGlobalAddressInvalidIpVersionRule(),
	NewGoogleComputeGlobalAddressInvalidPurposeRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidIpProtocolRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidIpVersionRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeGlobalNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeInterconnectAttachmentInvalidBandwidthRule(),
	NewGoogleComputeInterconnectAttachmentInvalidNameRule(),
	NewGoogleComputeInterconnectAttachmentInvalidTypeRule(),
	NewGoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeRegionBackendServiceInvalidProtocolRule(),
	NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeRegionNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeRouteInvalidNameRule(),
	NewGoogleComputeRouterNatInvalidNatIpAllocateOptionRule(),
	NewGoogleComputeRouterNatInvalidSourceSubnetworkIpRangesToNatRule(),
	NewGoogleComputeRouterPeerInvalidAdvertiseModeRule(),
	NewGoogleComputeSslPolicyInvalidMinTlsVersionRule(),
	NewGoogleComputeSslPolicyInvalidProfileRule(),
	NewGoogleComputeTargetHttpsProxyInvalidQuicOverrideRule(),
	NewGoogleComputeTargetInstanceInvalidNatPolicyRule(),
	NewGoogleComputeTargetSslProxyInvalidProxyHeaderRule(),
	NewGoogleComputeTargetTcpProxyInvalidProxyHeaderRule(),
	NewGoogleDataCatalogEntryGroupInvalidEntryGroupIdRule(),
	NewGoogleDataCatalogEntryInvalidTypeRule(),
	NewGoogleDataCatalogEntryInvalidUserSpecifiedSystemRule(),
	NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule(),
	NewGoogleDataCatalogTagTemplateInvalidTagTemplateIdRule(),
	NewGoogleDataLossPreventionJobTriggerInvalidStatusRule(),
	NewGoogleDatastoreIndexInvalidAncestorRule(),
	NewGoogleDeploymentManagerDeploymentInvalidCreatePolicyRule(),
	NewGoogleDeploymentManagerDeploymentInvalidDeletePolicyRule(),
	NewGoogleDialogflowAgentInvalidApiVersionRule(),
	NewGoogleDialogflowAgentInvalidMatchModeRule(),
	NewGoogleDialogflowAgentInvalidTierRule(),
	NewGoogleDialogflowEntityTypeInvalidKindRule(),
	NewGoogleDialogflowIntentInvalidWebhookStateRule(),
	NewGoogleDnsManagedZoneInvalidVisibilityRule(),
	NewGoogleFilestoreInstanceInvalidTierRule(),
	NewGoogleFirestoreIndexInvalidQueryScopeRule(),
	NewGoogleHealthcareFhirStoreInvalidVersionRule(),
	NewGoogleKmsCryptoKeyInvalidPurposeRule(),
	NewGoogleKmsKeyRingImportJobInvalidImportMethodRule(),
	NewGoogleKmsKeyRingImportJobInvalidProtectionLevelRule(),
	NewGoogleMonitoringAlertPolicyInvalidCombinerRule(),
	NewGoogleMonitoringCustomServiceInvalidServiceIdRule(),
	NewGoogleMonitoringMetricDescriptorInvalidLaunchStageRule(),
	NewGoogleMonitoringMetricDescriptorInvalidMetricKindRule(),
	NewGoogleMonitoringMetricDescriptorInvalidValueTypeRule(),
	NewGoogleMonitoringSloInvalidCalendarPeriodRule(),
	NewGoogleMonitoringSloInvalidSloIdRule(),
	NewGoogleOsConfigPatchDeploymentInvalidPatchDeploymentIdRule(),
	NewGoogleRedisInstanceInvalidConnectModeRule(),
	NewGoogleRedisInstanceInvalidNameRule(),
	NewGoogleRedisInstanceInvalidTierRule(),
	NewGoogleSccSourceInvalidDisplayNameRule(),
	NewGoogleSpannerDatabaseInvalidNameRule(),
	NewGoogleSpannerInstanceInvalidNameRule(),
	NewGoogleSqlSourceRepresentationInstanceInvalidDatabaseVersionRule(),
	NewGoogleStorageBucketAccessControlInvalidRoleRule(),
	NewGoogleStorageDefaultObjectAccessControlInvalidRoleRule(),
	NewGoogleStorageHmacKeyInvalidStateRule(),
	NewGoogleStorageObjectAccessControlInvalidRoleRule(),
}
