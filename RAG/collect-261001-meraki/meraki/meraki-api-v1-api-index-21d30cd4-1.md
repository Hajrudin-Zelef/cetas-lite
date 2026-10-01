---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-1
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [1, 25]
sha256: b5c7b15ae3db4e5dcadc3b45dc34c61605b9d6e8a305bbac0af164cbeecd8b9b
---

# meraki-api-v1-api-index-21d30cd4

|  | Operation | Path Parameters | Request Parameters | Response Parameters | OAuth Scopes | 
|---|---|---|---|---|---|
|  | POST /administered/assistant/chat/completions Create a synchronous AI assistant chat completion for user-wide threads (BETA)  >  createAdministeredAssistantChatCompletion | `` | content, country, data, language, mediaType, platform, query, source, text, threadId, type | content, followUps, messageId, sources, threadId | dashboard:general:telemetry:write | 
|  | GET /administered/assistant/chat/threads List all active user-wide conversation threads for the authenticated user. (BETA)  >  getAdministeredAssistantChatThreads | perPage, sort, sortOrder, from, to | `` | count, dateModified, items, meta, remaining, tag, threadId, threadName, total | dashboard:general:telemetry:read | 
|  | POST /administered/assistant/chat/threads Create a user-wide conversation thread for multi-turn AI assistant interactions. (BETA)  >  createAdministeredAssistantChatThread | `` | threadName | threadId | dashboard:general:telemetry:write | 
|  | GET /administered/assistant/chat/threads/{threadId} Return a single user-wide conversation thread. (BETA)  >  getAdministeredAssistantChatThread | threadId | `` | dateModified, tag, threadId, threadName | dashboard:general:telemetry:read | 
|  | PUT /administered/assistant/chat/threads/{threadId} Update the name of a user-wide conversation thread. (BETA)  >  updateAdministeredAssistantChatThread | threadId | threadName | `` | dashboard:general:telemetry:write | 
|  | DELETE /administered/assistant/chat/threads/{threadId} Delete a user-wide conversation thread and all its messages. (BETA)  >  deleteAdministeredAssistantChatThread | threadId | `` | `` | dashboard:general:telemetry:write | 
|  | GET /administered/assistant/chat/threads/{threadId}/messages List messages in a user-wide conversation thread. (BETA)  >  getAdministeredAssistantChatThreadMessages | threadId, perPage, sortOrder | `` | comment, content, count, feedback, followUps, format, id, items, meta, reason, remaining, sender, sources, tag, timestamp, total, vote | dashboard:general:telemetry:read | 
|  | POST /administered/assistant/chat/threads/{threadId}/messages Create a new chat message in an existing user-wide thread. (BETA)  >  createAdministeredAssistantChatThreadMessage | threadId | content, data, mediaType, platform, source, text, type | messageId, runId, threadId | dashboard:general:telemetry:write | 
|  | GET /administered/assistant/chat/threads/{threadId}/messages/{messageId} Return a single message in a user-wide conversation thread. (BETA)  >  getAdministeredAssistantChatThreadMessage | threadId, messageId | `` | comment, content, feedback, followUps, format, id, reason, sender, sources, tag, timestamp, vote | dashboard:general:telemetry:read | 
|  | GET /administered/identities/me Returns the identity of the current user.  >  getAdministeredIdentitiesMe | `` | `` | api, authentication, created, email, enabled, key, lastUsedDashboardAt, mode, name, saml, twoFactor | `` | 
|  | GET /administered/identities/me/api/keys List the non-sensitive metadata associated with the API keys that belong to the user  >  getAdministeredIdentitiesMeApiKeys | `` | `` | createdAt, suffix | `` | 
|  | POST /administered/identities/me/api/keys/generate Generates an API key for an identity  >  generateAdministeredIdentitiesMeApiKeys | `` | `` | key | `` | 
|  | POST /administered/identities/me/api/keys/{suffix}/revoke Revokes an identity's API key, using the last four characters of the key  >  revokeAdministeredIdentitiesMeApiKeys | suffix | `` | `` | `` | 
|  | GET /administered/licensing/subscription/entitlements Retrieve the list of purchasable entitlements  >  getAdministeredLicensingSubscriptionEntitlements | skus, subscriptionType | `` | featureTier, isAddOn, isFree, name, productClass, productType, sku | dashboard:licensing:config:read | 
|  | POST /administered/licensing/subscription/networks/featureTiers/batchUpdate Batch change networks to their desired feature tier for specified product types  >  batchAdministeredLicensingSubscriptionNetworksFeatureTiersUpdate | `` | featureTier, id, isAtomic, items, network, productType, productTypes | error, errors, featureTier, id, items, network, productType, productTypes | dashboard:licensing:config:write | 
|  | GET /administered/licensing/subscription/subscriptions List available subscriptions  >  getAdministeredLicensingSubscriptionSubscriptions | perPage, startingAfter, endingBefore, subscriptionIds, organizationIds, statuses, productTypes, skus, name, startDate, endDate | `` | account, assigned, available, counts, description, domain, endDate, enterpriseAgreement, entitlements, id, lastUpdatedAt, limit, name, networks, organizations, productTypes, renewalRequested, seats, sku, smartAccount, startDate, status, subscriptionId, suites, type, webOrderId | dashboard:licensing:config:read | 
|  | POST /administered/licensing/subscription/subscriptions/claim Claim a subscription into an organization.  >  claimAdministeredLicensingSubscriptionSubscriptions | validate | claimKey, description, name, organizationId | account, assigned, available, counts, description, domain, endDate, enterpriseAgreement, entitlements, id, lastUpdatedAt, limit, name, networks, organizations, productTypes, renewalRequested, seats, sku, smartAccount, startDate, status, subscriptionId, suites, type, webOrderId | dashboard:licensing:config:write | 
|  | POST /administered/licensing/subscription/subscriptions/claimKey/validate Find a subscription by claim key  >  validateAdministeredLicensingSubscriptionSubscriptionsClaimKey | `` | claimKey | account, assigned, available, counts, description, domain, endDate, enterpriseAgreement, entitlements, id, lastUpdatedAt, limit, name, networks, organizations, productTypes, renewalRequested, seats, sku, smartAccount, startDate, status, subscriptionId, suites, type, webOrderId | dashboard:licensing:config:write | 
|  | GET /administered/licensing/subscription/subscriptions/compliance/statuses Get compliance status for requested subscriptions  >  getAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses | organizationIds, subscriptionIds | `` | byProductClass, entitlements, gracePeriodEndsAt, id, missing, name, productClass, quantity, sku, status, subscription, violations | dashboard:licensing:config:read | 
|  | POST /administered/licensing/subscription/subscriptions/{subscriptionId}/bind Bind networks to a subscription  >  bindAdministeredLicensingSubscriptionSubscription | subscriptionId, validate | networkIds | errors, id, insufficientEntitlements, name, networks, quantity, sku, subscriptionId | dashboard:licensing:config:write | 
|  | POST /administered/organizations/permissions/resolve Resolve the authenticated caller admin's permissions across multiple organizations (BETA)  >  resolveAdministeredOrganizationsPermissions | `` | endingBefore, organizationIds, perPage, startingAfter | action, controller, counts, id, items, meta, name, organizationId, remaining, resource, total, type | dashboard:iam:config:read | 
|  | GET /administered/search/live List the appropriate results for a given global search utilizing live_search_react (BETA)  >  getAdministeredSearchLive | query, organizationId, networkId | `` | deviceList, deviceType, externalLink, itemId, keywords, label, list | `` | 
|  | GET /devices/{serial} Return a single device  >  getDevice | serial | `` | address, beaconIdParams, details, firmware, floorPlanId, lanIp, lat, lng, mac, major, minor, model, name, networkId, notes, serial, tags, url, uuid, value | dashboard:general:config:read | 
