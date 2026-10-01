---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-26
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [540, 566]
sha256: e62b88c08ea410cdd5096a952d86a7745db9a55cda397dfd521185191ede7b17
---

# meraki-api-v1-api-index-21d30cd4

|  | GET /networks/{networkId}/vlanProfiles/{iname} Get an existing VLAN profile of a network  >  getNetworkVlanProfile | networkId, iname | `` | adaptivePolicyGroup, id, iname, isDefault, name, vlanGroups, vlanId, vlanIds, vlanNames | dashboard:general:config:read | 
|  | PUT /networks/{networkId}/vlanProfiles/{iname} Update an existing VLAN profile of a network  >  updateNetworkVlanProfile | networkId, iname | adaptivePolicyGroup, id, name, vlanGroups, vlanId, vlanIds, vlanNames | adaptivePolicyGroup, id, iname, isDefault, name, vlanGroups, vlanId, vlanIds, vlanNames | dashboard:general:config:write | 
|  | DELETE /networks/{networkId}/vlanProfiles/{iname} Delete a VLAN profile of a network  >  deleteNetworkVlanProfile | networkId, iname | `` | `` | dashboard:general:config:write | 
|  | GET /networks/{networkId}/webhooks/httpServers List the HTTP servers for a network  >  getNetworkWebhooksHttpServers | networkId | `` | enabled, id, name, networkId, payloadTemplate, payloadTemplateId, url | dashboard:general:telemetry:read | 
|  | POST /networks/{networkId}/webhooks/httpServers Add an HTTP server to a network  >  createNetworkWebhooksHttpServer | networkId | name, payloadTemplate, payloadTemplateId, sharedSecret, url | enabled, id, name, networkId, payloadTemplate, payloadTemplateId, url | dashboard:general:telemetry:write | 
|  | GET /networks/{networkId}/webhooks/httpServers/{httpServerId} Return an HTTP server for a network  >  getNetworkWebhooksHttpServer | networkId, httpServerId | `` | enabled, id, name, networkId, payloadTemplate, payloadTemplateId, url | dashboard:general:telemetry:read | 
|  | PUT /networks/{networkId}/webhooks/httpServers/{httpServerId} Update an HTTP server  >  updateNetworkWebhooksHttpServer | networkId, httpServerId | name, payloadTemplate, payloadTemplateId, sharedSecret | enabled, id, name, networkId, payloadTemplate, payloadTemplateId, url | dashboard:general:telemetry:write | 
|  | DELETE /networks/{networkId}/webhooks/httpServers/{httpServerId} Delete an HTTP server from a network  >  deleteNetworkWebhooksHttpServer | networkId, httpServerId | `` | `` | dashboard:general:telemetry:write | 
|  | GET /networks/{networkId}/webhooks/payloadTemplates List the webhook payload templates for a network  >  getNetworkWebhooksPayloadTemplates | networkId | `` | adminsCanModify, body, byNetwork, headers, name, payloadTemplateId, sharing, template, type | dashboard:general:telemetry:read | 
|  | POST /networks/{networkId}/webhooks/payloadTemplates Create a webhook payload template for a network  >  createNetworkWebhooksPayloadTemplate | networkId | body, bodyFile, headers, headersFile, name, template | adminsCanModify, body, byNetwork, headers, name, payloadTemplateId, sharing, template, type | dashboard:general:telemetry:write | 
|  | GET /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} Get the webhook payload template for a network  >  getNetworkWebhooksPayloadTemplate | networkId, payloadTemplateId | `` | adminsCanModify, body, byNetwork, headers, name, payloadTemplateId, sharing, template, type | dashboard:general:telemetry:read | 
|  | DELETE /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} Destroy a webhook payload template for a network  >  deleteNetworkWebhooksPayloadTemplate | networkId, payloadTemplateId | `` | `` | dashboard:general:telemetry:write | 
|  | PUT /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} Update a webhook payload template for a network  >  updateNetworkWebhooksPayloadTemplate | networkId, payloadTemplateId | body, bodyFile, headers, headersFile, name, template | adminsCanModify, body, byNetwork, headers, name, payloadTemplateId, sharing, template, type | dashboard:general:telemetry:write | 
|  | POST /networks/{networkId}/webhooks/webhookTests Send a test webhook for a network  >  createNetworkWebhooksWebhookTest | networkId | alertTypeId, payloadTemplateId, payloadTemplateName, sharedSecret, url | id, status, url | dashboard:general:telemetry:write | 
|  | GET /networks/{networkId}/webhooks/webhookTests/{webhookTestId} Return the status of a webhook test for a network  >  getNetworkWebhooksWebhookTest | networkId, webhookTestId | `` | id, status, url | dashboard:general:telemetry:read | 
|  | GET /networks/{networkId}/wireless/airMarshal List Air Marshal scan results from a network  >  getNetworkWirelessAirMarshal | networkId, t0, timespan | `` | bssid, bssids, channels, contained, detectedBy, device, encryption, firstSeen, lastSeen, manufacturers, rssi, ssid, types, wiredLastSeen, wiredMacs, wiredVlans | wireless:config:read | 
|  | POST /networks/{networkId}/wireless/airMarshal/rules Creates a new rule  >  createNetworkWirelessAirMarshalRule | networkId | match, string, type | createdAt, id, match, name, network, ruleId, string, type, updatedAt | `` | 
|  | PUT /networks/{networkId}/wireless/airMarshal/rules/{ruleId} Update a rule  >  updateNetworkWirelessAirMarshalRule | networkId, ruleId | match, string, type | createdAt, id, match, name, network, ruleId, string, type, updatedAt | `` | 
|  | DELETE /networks/{networkId}/wireless/airMarshal/rules/{ruleId} Delete an Air Marshal rule.  >  deleteNetworkWirelessAirMarshalRule | networkId, ruleId | `` | `` | `` | 
|  | PUT /networks/{networkId}/wireless/airMarshal/settings Updates Air Marshal settings.  >  updateNetworkWirelessAirMarshalSettings | networkId | defaultPolicy | defaultPolicy, networkId | `` | 
|  | GET /networks/{networkId}/wireless/alternateManagementInterface Return alternate management interface and devices with IP assigned  >  getNetworkWirelessAlternateManagementInterface | networkId | `` | `` | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/alternateManagementInterface Update alternate management interface and device static IP  >  updateNetworkWirelessAlternateManagementInterface | networkId | accessPoints, alternateManagementIp, dns1, dns2, enabled, gateway, protocols, serial, subnetMask, vlanId | `` | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/billing Return the billing settings of this network  >  getNetworkWirelessBilling | networkId | `` | bandwidthLimits, currency, id, limitDown, limitUp, plans, price, timeLimit | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/billing Update the billing settings  >  updateNetworkWirelessBilling | networkId | bandwidthLimits, currency, id, limitDown, limitUp, plans, price, timeLimit | bandwidthLimits, currency, id, limitDown, limitUp, plans, price, timeLimit | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/bluetooth/settings Return the Bluetooth settings for a network. Bluetooth settings must be enabled on the network.  >  getNetworkWirelessBluetoothSettings | networkId | `` | advertised, advertisingEnabled, eslEnabled, major, majorMinorAssignmentMode, minor, power, scanningEnabled, transmit, uuid | wireless:config:read | 
|  | PUT /networks/{networkId}/wireless/bluetooth/settings Update the Bluetooth settings for a network  >  updateNetworkWirelessBluetoothSettings | networkId | advertised, advertisingEnabled, major, majorMinorAssignmentMode, minor, power, scanningEnabled, transmit, uuid | advertised, advertisingEnabled, eslEnabled, major, majorMinorAssignmentMode, minor, power, scanningEnabled, transmit, uuid | wireless:config:write | 
|  | GET /networks/{networkId}/wireless/channelUtilizationHistory Return AP channel utilization over time for a device or network client  >  getNetworkWirelessChannelUtilizationHistory | networkId, t0, t1, timespan, resolution, autoResolution, clientId, deviceSerial, apTag, band | `` | endTs, startTs, utilization80211, utilizationNon80211, utilizationTotal | wireless:telemetry:read | 
