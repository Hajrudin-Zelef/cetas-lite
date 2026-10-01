---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-54
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [1057, 1082]
sha256: b602c441e0a88ceb4392819bc3ea06d4779f9637570939b3deb68d124a8aa814
---

# meraki-api-v1-api-index-21d30cd4

|  | GET /organizations/{organizationId}/iam/users/idps/productIntegrations List all available IdP Product Integration urls for the organization (BETA)  >  getOrganizationIamUsersIdpsProductIntegrations | organizationId | `` | description, name, productIntegrationId, url | `` | 
|  | POST /organizations/{organizationId}/iam/users/idps/search Search all IdPs for an organization (BETA)  >  createOrganizationIamUsersIdpsSearch | organizationId | authZone, endingBefore, id, idpIds, perPage, startingAfter, type | counts, createdAt, description, idpConfig, idpId, items, lastUpdatedAt, meta, name, remaining, syncType, syncable, total, type | `` | 
|  | GET /organizations/{organizationId}/iam/users/idps/sync/history Get the IdP sync status records for an organization (BETA)  >  getOrganizationIamUsersIdpsSyncHistory | organizationId, perPage, startingAfter, endingBefore, idpId | `` | counts, createdAt, idpId, idpSyncId, items, lastUpdatedAt, message, meta, remaining, status, total | `` | 
|  | GET /organizations/{organizationId}/iam/users/idps/sync/latest Get the latest IdP sync status records for all IdPs in an organization (BETA)  >  getOrganizationIamUsersIdpsSyncLatest | organizationId, perPage, startingAfter, endingBefore, idpIds, authZoneId, authZoneType | `` | createdAt, idpId, idpSyncId, items, lastUpdatedAt, message, status, syncedBy | `` | 
|  | POST /organizations/{organizationId}/iam/users/idps/testConnectivity Test connectivity to an Entra ID identity provider. (BETA)  >  createOrganizationIamUsersIdpsTestConnectivity | organizationId | client, id, idpConfig, idpId, secret, tenant | code, errors, message, result | `` | 
|  | POST /organizations/{organizationId}/iam/users/idps/users Create a Meraki user (BETA)  >  createOrganizationIamUsersIdpsUser | organizationId | displayName, email, password, sendPassword | accessTypes, createdAt, displayName, externalId, groups, id, idp, idpUserId, lastUpdatedAt, name, type, upn | `` | 
|  | PUT /organizations/{organizationId}/iam/users/idps/users/{id} Update a Meraki user (BETA)  >  updateOrganizationIamUsersIdpsUser | organizationId, id | displayName, email, password, sendPassword | accessTypes, createdAt, displayName, externalId, groups, id, idp, idpUserId, lastUpdatedAt, name, type, upn | `` | 
|  | DELETE /organizations/{organizationId}/iam/users/idps/users/{id} Delete a Meraki end user (BETA)  >  deleteOrganizationIamUsersIdpsUser | organizationId, id | `` | `` | `` | 
|  | POST /organizations/{organizationId}/iam/users/idps/{idpId}/sync Trigger an IdP sync for an identity provider (BETA)  >  createOrganizationIamUsersIdpSync | organizationId, idpId | emails, force | createdAt, idpId, idpSyncId, lastUpdatedAt, message, status, syncedBy | `` | 
|  | GET /organizations/{organizationId}/iam/users/idps/{idpId}/sync/latest Get the latest IdP sync status for an identity provider (BETA)  >  getOrganizationIamUsersIdpSyncLatest | organizationId, idpId | `` | createdAt, idpSyncId, lastUpdatedAt, message, status | `` | 
|  | PUT /organizations/{organizationId}/iam/users/idps/{id} Update an identity provider (BETA)  >  updateOrganizationIamUsersIdp | organizationId, id | clientId, clientSecret, description, idpConfig, name, syncType, tenantId | createdAt, description, idpConfig, idpId, lastUpdatedAt, name, syncType, syncable, type | `` | 
|  | DELETE /organizations/{organizationId}/iam/users/idps/{id} Delete a identity provider from an organization (BETA)  >  deleteOrganizationIamUsersIdp | organizationId, id | `` | `` | `` | 
|  | GET /organizations/{organizationId}/iam/users/idps/{id}/authZones List all auth zones for an identity provider (BETA)  >  getOrganizationIamUsersIdpAuthZones | organizationId, id | `` | id, items, name, network, type | `` | 
|  | POST /organizations/{organizationId}/iam/users/search List the end users and their associated identity providers for an organization. (BETA)  >  searchOrganizationUsers | organizationId | accessTypes, endingBefore, groupIds, idpIds, perPage, searchQuery, sortKey, sortOrder, startingAfter, statuses, userIds | accessTypes, counts, createdAt, displayName, email, externalId, groups, id, idp, idpUsers, items, lastUpdatedAt, meta, name, remaining, total, type, upn, userId, username | `` | 
|  | GET /organizations/{organizationId}/iam/users/summaryPanel Get the count of users and user groups for an organization. (BETA)  >  getOrganizationIamUsersSummaryPanel | organizationId | `` | userCount, userGroupCount | `` | 
|  | GET /organizations/{organizationId}/insight/applications List all Insight tracked applications  >  getOrganizationInsightApplications | organizationId | `` | applicationId, byNetwork, goodput, name, networkId, responseDuration, thresholds, type | sdwan:telemetry:read | 
|  | POST /organizations/{organizationId}/insight/applications Add an Insight tracked application (BETA)  >  createOrganizationInsightApplication | organizationId | counterSetRuleId, enableSmartThresholds, goodput, responseTime, thresholds | applicationId, byNetwork, goodput, name, networkId, responseDuration, thresholds, type | `` | 
|  | PUT /organizations/{organizationId}/insight/applications/{applicationId} Update an Insight tracked application (BETA)  >  updateOrganizationInsightApplication | organizationId, applicationId | enableSmartThresholds, goodput, responseTime, thresholds | applicationId, byNetwork, goodput, name, networkId, responseDuration, thresholds, type | `` | 
|  | DELETE /organizations/{organizationId}/insight/applications/{applicationId} Delete an Insight tracked application (BETA)  >  deleteOrganizationInsightApplication | organizationId, applicationId | `` | `` | `` | 
|  | GET /organizations/{organizationId}/insight/monitoredMediaServers List the monitored media servers for this organization  >  getOrganizationInsightMonitoredMediaServers | organizationId | `` | address, bestEffortMonitoringEnabled, id, name | `` | 
|  | POST /organizations/{organizationId}/insight/monitoredMediaServers Add a media server to be monitored for this organization  >  createOrganizationInsightMonitoredMediaServer | organizationId | address, bestEffortMonitoringEnabled, name | address, bestEffortMonitoringEnabled, id, name | `` | 
|  | GET /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} Return a monitored media server for this organization  >  getOrganizationInsightMonitoredMediaServer | organizationId, monitoredMediaServerId | `` | address, bestEffortMonitoringEnabled, id, name | `` | 
|  | PUT /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} Update a monitored media server for this organization  >  updateOrganizationInsightMonitoredMediaServer | organizationId, monitoredMediaServerId | address, bestEffortMonitoringEnabled, name | address, bestEffortMonitoringEnabled, id, name | `` | 
|  | DELETE /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} Delete a monitored media server from this organization  >  deleteOrganizationInsightMonitoredMediaServer | organizationId, monitoredMediaServerId | `` | `` | `` | 
|  | GET /organizations/{organizationId}/insight/speedTestResults List the speed tests for the given devices under this organization (BETA)  >  getOrganizationInsightSpeedTestResults | organizationId, serials, timespan, t0, t1 | `` | average, interface, networkId, request, results, serial, speedTestId, speeds, startedAt | `` | 
|  | GET /organizations/{organizationId}/insight/webApps Lists all default web applications rules with counter set rule ids (BETA)  >  getOrganizationInsightWebApps | organizationId | `` | category, counterSetRuleId, expression, goodput, host, name, net, port, responseDelay, signature, signatureType, thresholds | `` | 
