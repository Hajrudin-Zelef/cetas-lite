---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-30
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [616, 638]
sha256: f0b123bfca790057313049f7c546deeab097e16e0e332223dbff04b6819c24c2
---

# meraki-api-v1-api-index-21d30cd4

|  | PUT /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies Update the device type group policies for the SSID  >  updateNetworkWirelessSsidDeviceTypeGroupPolicies | networkId, number | devicePolicy, deviceType, deviceTypePolicies, enabled, groupPolicyId | devicePolicy, deviceType, deviceTypePolicies, enabled, groupPolicyId | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/eapOverride Return the EAP overridden parameters for an SSID  >  getNetworkWirelessSsidEapOverride | networkId, number | `` | eapolKey, identity, maxRetries, retries, timeout, timeoutInMs | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/eapOverride Update the EAP overridden parameters for an SSID.  >  updateNetworkWirelessSsidEapOverride | networkId, number | eapolKey, identity, maxRetries, retries, timeout, timeoutInMs | eapolKey, identity, maxRetries, retries, timeout, timeoutInMs | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules Return the L3 firewall rules for an SSID on an MR network  >  getNetworkWirelessSsidFirewallL3FirewallRules | networkId, number | `` | allowLanAccess, comment, destCidr, destPort, ipVer, policy, protocol, rules | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules Update the L3 firewall rules of an SSID on an MR network  >  updateNetworkWirelessSsidFirewallL3FirewallRules | networkId, number | allowLanAccess, comment, destCidr, destPort, ipVer, policy, protocol, rules | allowLanAccess, comment, destCidr, destPort, ipVer, policy, protocol, rules | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules Return the L7 firewall rules for an SSID on an MR network  >  getNetworkWirelessSsidFirewallL7FirewallRules | networkId, number | `` | policy, rules, type, value | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules Update the L7 firewall rules of an SSID on an MR network  >  updateNetworkWirelessSsidFirewallL7FirewallRules | networkId, number | policy, rules, type, value | policy, rules, type, value | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/hotspot20 Return the Hotspot 2.0 settings for an SSID  >  getNetworkWirelessSsidHotspot20 | networkId, number | `` | authenticationTypes, credentials, domains, eapInnerAuthentication, enabled, format, id, mcc, mccMncs, methods, mnc, naiRealms, name, networkAccessType, nonEapInnerAuthentication, operator, roamConsortOis, tunneledEapMethodCredentials, type, venue | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/hotspot20 Update the Hotspot 2.0 settings of an SSID  >  updateNetworkWirelessSsidHotspot20 | networkId, number | authenticationTypes, credentials, domains, eapInnerAuthentication, enabled, format, id, mcc, mccMncs, methods, mnc, naiRealms, name, networkAccessType, nonEapInnerAuthentication, operator, realm, roamConsortOis, tunneledEapMethodCredentials, type, venue | authenticationTypes, credentials, domains, eapInnerAuthentication, enabled, format, id, mcc, mccMncs, methods, mnc, naiRealms, name, networkAccessType, nonEapInnerAuthentication, operator, roamConsortOis, tunneledEapMethodCredentials, type, venue | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/identityPsks List all Identity PSKs in a wireless network  >  getNetworkWirelessSsidIdentityPsks | networkId, number | `` | email, expiresAt, groupPolicyId, id, name, passphrase, wifiPersonalNetworkId | wireless:config:read | 
|  | POST /networks/{networkId}/wireless/ssids/{number}/identityPsks Create an Identity PSK  >  createNetworkWirelessSsidIdentityPsk | networkId, number | expiresAt, groupPolicyId, name, passphrase | email, expiresAt, groupPolicyId, id, name, passphrase, wifiPersonalNetworkId | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} Return an Identity PSK  >  getNetworkWirelessSsidIdentityPsk | networkId, number, identityPskId | `` | email, expiresAt, groupPolicyId, id, name, passphrase, wifiPersonalNetworkId | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} Update an Identity PSK  >  updateNetworkWirelessSsidIdentityPsk | networkId, number, identityPskId | expiresAt, groupPolicyId, name, passphrase | email, expiresAt, groupPolicyId, id, name, passphrase, wifiPersonalNetworkId | wireless:config:write | 
|  | DELETE /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} Delete an Identity PSK  >  deleteNetworkWirelessSsidIdentityPsk | networkId, number, identityPskId | `` | `` | wireless:config:write | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/openRoaming Update the OpenRoaming setting for the SSID  >  updateNetworkWirelessSsidOpenRoaming | networkId, number | enabled, tenantId | enabled, tenantId | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/overrides Display the overrides for this SSID  >  getNetworkWirelessSsidOverrides | networkId, number | `` | ccxNameIeEnabled | `` | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/overrides Update the overrides for this SSID  >  updateNetworkWirelessSsidOverrides | networkId, number | ccxNameIeEnabled | ccxNameIeEnabled | `` | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion Update the client exclusion status configuration for a given SSID (BETA)  >  updateNetworkWirelessSsidPoliciesClientExclusion | networkId, number | dynamic, enabled, reasons, static, timeout | dynamic, enabled, id, name, network, number, reasons, ssid, static, timeout | `` | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion/static/exclusions Replace the static client exclusion list for the given SSID (use PUT /exclusions) (BETA)  >  updateNetworkWirelessSsidPoliciesClientExclusionStaticExclusions | networkId, number | macs | id, macs, name, network, number, ssid | `` | 
|  | POST /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion/static/exclusions/bulkAdd Add MAC addresses to the existing static client exclusion list for the given SSID (use POST /bulkAdd) (BETA)  >  createNetworkWirelessSsidPoliciesClientExclusionStaticExclusionsBulkAdd | networkId, number | macs | id, macs, name, network, number, ssid | `` | 
|  | POST /networks/{networkId}/wireless/ssids/{number}/policies/clientExclusion/static/exclusions/bulkRemove Remove MAC addresses from the existing static client exclusion list for the given SSID (use POST /bulkRemove) (BETA)  >  createNetworkWirelessSsidPoliciesClientExclusionStaticExclusionsBulkRemove | networkId, number | macs | `` | `` | 
|  | GET /networks/{networkId}/wireless/ssids/{number}/schedules List the outage schedule for the SSID  >  getNetworkWirelessSsidSchedules | networkId, number | `` | enabled, end, endDay, endTime, ranges, rangesInSeconds, start, startDay, startTime | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/ssids/{number}/schedules Update the outage schedule for the SSID  >  updateNetworkWirelessSsidSchedules | networkId, number | enabled, end, endDay, endTime, ranges, rangesInSeconds, start, startDay, startTime | enabled, end, endDay, endTime, ranges, rangesInSeconds, start, startDay, startTime | wireless:config:write | 
