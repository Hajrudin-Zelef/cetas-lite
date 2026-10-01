---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-63
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [1230, 1256]
sha256: 944fe92ac5b27b1be2629058e0711a3d8d4a53cfabb6cca37354cb7d010255fa
---

# meraki-api-v1-api-index-21d30cd4

|  | DELETE /organizations/{organizationId}/policies/global/group/policies/{policyId} Delete an Organization-Wide Policy  >  deleteOrganizationPoliciesGlobalGroupPolicy | organizationId, policyId | `` | `` | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/policyObjects Lists Policy Objects belonging to the organization.  >  getOrganizationPolicyObjects | organizationId, perPage, startingAfter, endingBefore | `` | category, cidr, createdAt, groupIds, id, name, networkIds, type, updatedAt | dashboard:general:config:read | 
|  | POST /organizations/{organizationId}/policyObjects Creates a new Policy Object  >  createOrganizationPolicyObject | organizationId | category, cidr, fqdn, groupIds, ip, mask, name, type | category, cidr, createdAt, groupIds, id, name, networkIds, type, updatedAt | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/policyObjects/groups Lists Policy Object Groups belonging to the organization.  >  getOrganizationPolicyObjectsGroups | organizationId, perPage, startingAfter, endingBefore | `` | category, createdAt, id, name, networkIds, objectIds, updatedAt | dashboard:general:config:read | 
|  | POST /organizations/{organizationId}/policyObjects/groups Creates a new Policy Object Group.  >  createOrganizationPolicyObjectsGroup | organizationId | category, name, objectIds | category, createdAt, id, name, networkIds, objectIds, updatedAt | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} Shows details of a Policy Object Group.  >  getOrganizationPolicyObjectsGroup | organizationId, policyObjectGroupId | `` | category, createdAt, id, name, networkIds, objectIds, updatedAt | dashboard:general:config:read | 
|  | PUT /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} Updates a Policy Object Group.  >  updateOrganizationPolicyObjectsGroup | organizationId, policyObjectGroupId | name, objectIds | category, createdAt, id, name, networkIds, objectIds, updatedAt | dashboard:general:config:write | 
|  | DELETE /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} Deletes a Policy Object Group.  >  deleteOrganizationPolicyObjectsGroup | organizationId, policyObjectGroupId | `` | `` | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/policyObjects/{policyObjectId} Shows details of a Policy Object.  >  getOrganizationPolicyObject | organizationId, policyObjectId | `` | category, cidr, createdAt, groupIds, id, name, networkIds, type, updatedAt | dashboard:general:config:read | 
|  | PUT /organizations/{organizationId}/policyObjects/{policyObjectId} Updates a Policy Object  >  updateOrganizationPolicyObject | organizationId, policyObjectId | cidr, fqdn, groupIds, ip, mask, name | category, cidr, createdAt, groupIds, id, name, networkIds, type, updatedAt | dashboard:general:config:write | 
|  | DELETE /organizations/{organizationId}/policyObjects/{policyObjectId} Deletes a Policy Object.  >  deleteOrganizationPolicyObject | organizationId, policyObjectId | `` | `` | dashboard:general:config:write | 
|  | GET /organizations/{organizationId}/routing/vrfs List existing organization-wide VRFs (Virtual Routing and Forwarding). (BETA)  >  getOrganizationRoutingVrfs | organizationId, vrfIds | `` | appliance, autoRd, counts, description, enabled, index, items, meta, name, routeDistinguisher, routeTarget, switchFabricId, total, vpn, vrfId | switch:config:read | 
|  | POST /organizations/{organizationId}/routing/vrfs Add an organization-wide VRF (Virtual Routing and Forwarding) (BETA)  >  createOrganizationRoutingVrf | organizationId | appliance, description, enabled, index, name, routeDistinguisher, routeTarget | appliance, autoRd, description, enabled, index, name, routeDistinguisher, routeTarget, switchFabricId, vpn, vrfId | switch:config:write | 
|  | GET /organizations/{organizationId}/routing/vrfs/overview/byVrf List existing organization-wide VRFs (Virtual Routing and Forwarding) overviews. (BETA)  >  getOrganizationRoutingVrfsOverviewByVrf | organizationId, vrfIds | `` | appliance, counts, description, enabled, index, items, meta, name, networks, routeDistinguisher, routeTarget, total, vrfId | switch:config:read | 
|  | PUT /organizations/{organizationId}/routing/vrfs/{vrfId} Update an organization-wide VRF (Virtual Routing and Forwarding) (BETA)  >  updateOrganizationRoutingVrf | organizationId, vrfId | appliance, description, enabled, index, name, routeDistinguisher, routeTarget | appliance, autoRd, description, enabled, index, name, routeDistinguisher, routeTarget, switchFabricId, vpn, vrfId | switch:config:write | 
|  | DELETE /organizations/{organizationId}/routing/vrfs/{vrfId} Delete a VRF (Virtual Routing and Forwarding) from a organization (BETA)  >  deleteOrganizationRoutingVrf | organizationId, vrfId | `` | `` | switch:config:write | 
|  | GET /organizations/{organizationId}/saml Returns the SAML SSO enabled settings for an organization.  >  getOrganizationSaml | organizationId | `` | enabled, idpId, spInitiated, subdomain | dashboard:iam:config:read | 
|  | PUT /organizations/{organizationId}/saml Updates the SAML SSO enabled settings for an organization.  >  updateOrganizationSaml | organizationId | enabled, idpId, spInitiated, subdomain | enabled, idpId, spInitiated, subdomain | dashboard:iam:config:write | 
|  | GET /organizations/{organizationId}/saml/idps List the SAML IdPs in your organization.  >  getOrganizationSamlIdps | organizationId | `` | consumerUrl, idpId, sloLogoutUrl, ssoLoginUrl, visionConsumerUrl, x509certSha1Fingerprint | dashboard:iam:config:read | 
|  | POST /organizations/{organizationId}/saml/idps Create a SAML IdP for your organization.  >  createOrganizationSamlIdp | organizationId | sloLogoutUrl, ssoLoginUrl, x509certSha1Fingerprint | consumerUrl, idpId, sloLogoutUrl, ssoLoginUrl, visionConsumerUrl, x509certSha1Fingerprint | dashboard:iam:config:write | 
|  | PUT /organizations/{organizationId}/saml/idps/{idpId} Update a SAML IdP in your organization  >  updateOrganizationSamlIdp | organizationId, idpId | sloLogoutUrl, ssoLoginUrl, x509certSha1Fingerprint | consumerUrl, idpId, sloLogoutUrl, ssoLoginUrl, visionConsumerUrl, x509certSha1Fingerprint | dashboard:iam:config:write | 
|  | GET /organizations/{organizationId}/saml/idps/{idpId} Get a SAML IdP from your organization.  >  getOrganizationSamlIdp | organizationId, idpId | `` | consumerUrl, idpId, sloLogoutUrl, ssoLoginUrl, visionConsumerUrl, x509certSha1Fingerprint | dashboard:iam:config:read | 
|  | DELETE /organizations/{organizationId}/saml/idps/{idpId} Remove a SAML IdP in your organization.  >  deleteOrganizationSamlIdp | organizationId, idpId | `` | `` | dashboard:iam:config:write | 
|  | GET /organizations/{organizationId}/samlRoles List the SAML roles for this organization  >  getOrganizationSamlRoles | organizationId | `` | access, camera, id, networks, orgAccess, orgWide, role, tag, tags | dashboard:iam:config:read | 
|  | POST /organizations/{organizationId}/samlRoles Create a SAML role  >  createOrganizationSamlRole | organizationId | access, id, networks, orgAccess, role, tag, tags | access, camera, id, networks, orgAccess, orgWide, role, tag, tags | dashboard:iam:config:write | 
|  | GET /organizations/{organizationId}/samlRoles/{samlRoleId} Return a SAML role  >  getOrganizationSamlRole | organizationId, samlRoleId | `` | access, camera, id, networks, orgAccess, orgWide, role, tag, tags | dashboard:iam:config:read | 
|  | PUT /organizations/{organizationId}/samlRoles/{samlRoleId} Update a SAML role  >  updateOrganizationSamlRole | organizationId, samlRoleId | access, id, networks, orgAccess, role, tag, tags | access, camera, id, networks, orgAccess, orgWide, role, tag, tags | dashboard:iam:config:write | 
