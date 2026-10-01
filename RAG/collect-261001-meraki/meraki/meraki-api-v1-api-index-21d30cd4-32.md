---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-32
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [658, 680]
sha256: 3bdaa0c2d77f7ae4dc5acc6e4ff540064a7f4e7d76bde0881ce91d01c0ed746f
---

# meraki-api-v1-api-index-21d30cd4

|  | POST /organizations/{organizationId}/adaptivePolicy/acls Creates new adaptive policy ACL  >  createOrganizationAdaptivePolicyAcl | organizationId | description, dstPort, ipVersion, log, name, policy, protocol, rules, srcPort, tcpEstablished | aclId, createdAt, description, dstPort, ipVersion, log, name, policy, protocol, rules, srcPort, tcpEstablished, updatedAt | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/adaptivePolicy/acls/{aclId} Returns the adaptive policy ACL information  >  getOrganizationAdaptivePolicyAcl | organizationId, aclId | `` | aclId, createdAt, description, dstPort, ipVersion, log, name, policy, protocol, rules, srcPort, tcpEstablished, updatedAt | dashboard:general:config:read | 
|  | PUT /organizations/{organizationId}/adaptivePolicy/acls/{aclId} Updates an adaptive policy ACL  >  updateOrganizationAdaptivePolicyAcl | organizationId, aclId | description, dstPort, ipVersion, log, name, policy, protocol, rules, srcPort, tcpEstablished | aclId, createdAt, description, dstPort, ipVersion, log, name, policy, protocol, rules, srcPort, tcpEstablished, updatedAt | dashboard:general:config:write | 
|  | DELETE /organizations/{organizationId}/adaptivePolicy/acls/{aclId} Deletes the specified adaptive policy ACL  >  deleteOrganizationAdaptivePolicyAcl | organizationId, aclId | `` | `` | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/adaptivePolicy/groups List adaptive policy groups in a organization  >  getOrganizationAdaptivePolicyGroups | organizationId | `` | createdAt, description, groupId, id, isDefaultGroup, name, policyObjects, requiredIpMappings, sgt, updatedAt | dashboard:general:config:read | 
|  | POST /organizations/{organizationId}/adaptivePolicy/groups Creates a new adaptive policy group  >  createOrganizationAdaptivePolicyGroup | organizationId | description, id, name, policyObjects, sgt | createdAt, description, groupId, id, isDefaultGroup, name, policyObjects, requiredIpMappings, sgt, updatedAt | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/adaptivePolicy/groups/{id} Returns an adaptive policy group  >  getOrganizationAdaptivePolicyGroup | organizationId, id | `` | createdAt, description, groupId, id, isDefaultGroup, name, policyObjects, requiredIpMappings, sgt, updatedAt | dashboard:general:config:read | 
|  | PUT /organizations/{organizationId}/adaptivePolicy/groups/{id} Updates an adaptive policy group  >  updateOrganizationAdaptivePolicyGroup | organizationId, id | description, id, name, policyObjects, sgt | createdAt, description, groupId, id, isDefaultGroup, name, policyObjects, requiredIpMappings, sgt, updatedAt | dashboard:general:config:write | 
|  | DELETE /organizations/{organizationId}/adaptivePolicy/groups/{id} Deletes the specified adaptive policy group and any associated policies and references  >  deleteOrganizationAdaptivePolicyGroup | organizationId, id | `` | `` | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/adaptivePolicy/overview Returns adaptive policy aggregate statistics for an organization  >  getOrganizationAdaptivePolicyOverview | organizationId | `` | aclsInAPolicy, allowPolicies, counts, customAcls, customGroups, denyPolicies, groups, limits, policies, policyObjects, rulesInAnAcl | dashboard:general:config:read | 
|  | GET /organizations/{organizationId}/adaptivePolicy/policies List adaptive policies in an organization  >  getOrganizationAdaptivePolicyPolicies | organizationId | `` | acls, adaptivePolicyId, createdAt, destinationGroup, id, lastEntryRule, name, sgt, sourceGroup, updatedAt | dashboard:general:config:read | 
|  | POST /organizations/{organizationId}/adaptivePolicy/policies Add an Adaptive Policy  >  createOrganizationAdaptivePolicyPolicy | organizationId | acls, destinationGroup, id, lastEntryRule, name, sgt, sourceGroup | acls, adaptivePolicyId, createdAt, destinationGroup, id, lastEntryRule, name, sgt, sourceGroup, updatedAt | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/adaptivePolicy/policies/{id} Return an adaptive policy  >  getOrganizationAdaptivePolicyPolicy | organizationId, id | `` | acls, adaptivePolicyId, createdAt, destinationGroup, id, lastEntryRule, name, sgt, sourceGroup, updatedAt | dashboard:general:config:read | 
|  | PUT /organizations/{organizationId}/adaptivePolicy/policies/{id} Update an Adaptive Policy  >  updateOrganizationAdaptivePolicyPolicy | organizationId, id | acls, destinationGroup, id, lastEntryRule, name, sgt, sourceGroup | acls, adaptivePolicyId, createdAt, destinationGroup, id, lastEntryRule, name, sgt, sourceGroup, updatedAt | dashboard:general:config:write | 
|  | DELETE /organizations/{organizationId}/adaptivePolicy/policies/{id} Delete an Adaptive Policy  >  deleteOrganizationAdaptivePolicyPolicy | organizationId, id | `` | `` | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/adaptivePolicy/settings Returns global adaptive policy settings in an organization  >  getOrganizationAdaptivePolicySettings | organizationId | `` | enabledNetworks | dashboard:general:config:read | 
|  | PUT /organizations/{organizationId}/adaptivePolicy/settings Update global adaptive policy settings  >  updateOrganizationAdaptivePolicySettings | organizationId | enabledNetworks | enabledNetworks | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/admins List the dashboard administrators in this organization  >  getOrganizationAdmins | organizationId, networkIds | `` | access, accountStatus, authenticationMethod, email, hasApiKey, id, lastActive, name, networks, orgAccess, tag, tags, twoFactorAuthEnabled | dashboard:iam:config:read | 
|  | POST /organizations/{organizationId}/admins Create a new dashboard administrator  >  createOrganizationAdmin | organizationId | access, authenticationMethod, email, id, name, networks, orgAccess, tag, tags | access, accountStatus, authenticationMethod, email, hasApiKey, id, lastActive, name, networks, orgAccess, tag, tags, twoFactorAuthEnabled | dashboard:iam:config:write | 
|  | PUT /organizations/{organizationId}/admins/{adminId} Update an administrator  >  updateOrganizationAdmin | organizationId, adminId | access, id, name, networks, orgAccess, tag, tags | access, accountStatus, authenticationMethod, email, hasApiKey, id, lastActive, name, networks, orgAccess, tag, tags, twoFactorAuthEnabled | dashboard:iam:config:write | 
|  | DELETE /organizations/{organizationId}/admins/{adminId} Revoke all access for a dashboard administrator within this organization  >  deleteOrganizationAdmin | organizationId, adminId | `` | `` | dashboard:iam:config:write | 
|  | GET /organizations/{organizationId}/alerts/profiles List all organization-wide alert configurations  >  getOrganizationAlertsProfiles | organizationId | `` | alertCondition, bit_rate_bps, description, duration, emails, enabled, httpServerIds, id, interface, networkTags, recipients, type, window | sdwan:telemetry:read | 
|  | POST /organizations/{organizationId}/alerts/profiles Create an organization-wide alert configuration  >  createOrganizationAlertsProfile | organizationId | alertCondition, bit_rate_bps, description, duration, emails, httpServerIds, interface, jitter_ms, latency_ms, loss_ratio, mos, networkTags, recipients, type, window | alertCondition, bit_rate_bps, description, duration, emails, enabled, httpServerIds, id, interface, networkTags, recipients, type, window | sdwan:telemetry:write | 
