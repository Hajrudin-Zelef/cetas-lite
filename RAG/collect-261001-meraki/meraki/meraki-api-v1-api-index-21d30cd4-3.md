---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-3
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [51, 77]
sha256: 729dfcbfe6ecc8cc806495810ca3871b19d8f4ef606a0afb541eedee57cba486
---

# meraki-api-v1-api-index-21d30cd4

|  | PUT /devices/{serial}/camera/sense Update sense settings for the given camera  >  updateDeviceCameraSense | serial | audioDetection, detectionModelId, enabled, mqttBrokerId, senseEnabled | audioDetection, detectionModelId, enabled, mqttBrokerId, mqttTopics, senseEnabled | camera:config:write | 
|  | GET /devices/{serial}/camera/sense/objectDetectionModels Returns the MV Sense object detection model list for the given camera  >  getDeviceCameraSenseObjectDetectionModels | serial | `` | description, id | camera:config:read | 
|  | GET /devices/{serial}/camera/video/settings Returns video settings for the given camera  >  getDeviceCameraVideoSettings | serial | `` | externalRtspEnabled, rtspUrl | camera:config:read | 
|  | PUT /devices/{serial}/camera/video/settings Update video settings for the given camera  >  updateDeviceCameraVideoSettings | serial | externalRtspEnabled | externalRtspEnabled, rtspUrl | camera:config:write | 
|  | GET /devices/{serial}/camera/videoLink Returns video link to the specified camera  >  getDeviceCameraVideoLink | serial, timestamp | `` | url, visionUrl | camera:config:read | 
|  | GET /devices/{serial}/camera/wirelessProfiles Returns wireless profile assigned to the given camera  >  getDeviceCameraWirelessProfiles | serial | `` | backup, ids, primary, secondary | camera:config:read | 
|  | PUT /devices/{serial}/camera/wirelessProfiles Assign wireless profiles to the given camera  >  updateDeviceCameraWirelessProfiles | serial | backup, ids, primary, secondary | backup, ids, primary, secondary | camera:config:write | 
|  | PUT /devices/{serial}/cellular/geolocations Update the enablement of the geolocation feature for a device  >  updateDeviceCellularGeolocations | serial | enabled | enabled | `` | 
|  | GET /devices/{serial}/cellular/sims Return the SIM and APN configurations for a cellular device.  >  getDeviceCellularSims | serial | `` | allowedIpTypes, apns, authentication, enabled, iccid, imsi, ip, isPrimary, msisdn, name, password, simFailover, simOrdering, sims, slot, status, timeout, translation, type, username | sdwan:config:read | 
|  | PUT /devices/{serial}/cellular/sims Updates the SIM and APN configurations for a cellular device.  >  updateDeviceCellularSims | serial | allowedIpTypes, apns, authentication, enabled, ip, isPrimary, name, password, simFailover, simOrder, simOrdering, sims, slot, timeout, translation, type, username | allowedIpTypes, apns, authentication, enabled, iccid, imsi, ip, isPrimary, msisdn, name, password, simFailover, simOrdering, sims, slot, status, timeout, translation, type, username | sdwan:config:write | 
|  | POST /devices/{serial}/cellular/uplinks/bands/masks/update Update the cellular band masks for a device  >  createDeviceCellularUplinksBandsMasksUpdate | serial | masked, slot, type | bySignalType, bySlot, enabled, masked, slot, supported, type | `` | 
|  | GET /devices/{serial}/cellularGateway/lan Show the LAN Settings of a MG  >  getDeviceCellularGatewayLan | serial | `` | comment, deviceLanIp, deviceName, deviceSubnet, end, fixedIpAssignments, ip, mac, name, reservedIpRanges, start | sdwan:config:read | 
|  | PUT /devices/{serial}/cellularGateway/lan Update the LAN Settings for a single MG.  >  updateDeviceCellularGatewayLan | serial | comment, end, fixedIpAssignments, ip, mac, name, reservedIpRanges, start | comment, deviceLanIp, deviceName, deviceSubnet, end, fixedIpAssignments, ip, mac, name, reservedIpRanges, start | sdwan:config:write | 
|  | GET /devices/{serial}/cellularGateway/portForwardingRules Returns the port forwarding rules for a single MG.  >  getDeviceCellularGatewayPortForwardingRules | serial | `` | access, allowedIps, lanIp, localPort, name, protocol, publicPort, rules | sdwan:config:read | 
|  | PUT /devices/{serial}/cellularGateway/portForwardingRules Updates the port forwarding rules for a single MG.  >  updateDeviceCellularGatewayPortForwardingRules | serial | access, allowedIps, lanIp, localPort, name, protocol, publicPort, rules | access, allowedIps, lanIp, localPort, name, protocol, publicPort, rules | sdwan:config:write | 
|  | POST /devices/{serial}/certificates/{certificateSerial}/revoke Revoke a device certificate (BETA)  >  revokeDeviceCertificate | serial, certificateSerial | reason | certificate, feature, id, message, result, serial, type | dashboard:iam:config:write | 
|  | GET /devices/{serial}/clients List the clients of a device, up to a maximum of a month ago  >  getDeviceClients | serial, t0, timespan | `` | adaptivePolicyGroup, description, dhcpHostname, id, ip, mac, mdnsName, namedVlan, recv, sent, switchport, usage, user, vlan | dashboard:general:telemetry:read | 
|  | POST /devices/{serial}/liveTools/aclHitCount Enqueue a job to perform an ACL hit count for the device (BETA)  >  createDeviceLiveToolsAclHitCount | serial | callback, httpServer, id, payloadTemplate, sharedSecret, url | aclHitCountId, callback, id, request, serial, status, url | dashboard:general:telemetry:write | 
|  | GET /devices/{serial}/liveTools/aclHitCount/{id} Return an ACL hit count live tool job. (BETA)  >  getDeviceLiveToolsAclHitCount | serial, id | `` | aclHitCountId, acls, address, counts, destination, error, ipProtocol, ipVersion, ipv4, ipv6, number, objectGroup, operator, policy, port, ports, request, serial, source, status, total, type, url | dashboard:general:telemetry:read | 
|  | POST /devices/{serial}/liveTools/arpTable Enqueue a job to perform a ARP table request for the device  >  createDeviceLiveToolsArpTable | serial | address, callback, httpServer, id, ip, payloadTemplate, sharedSecret, url | address, arpTableId, callback, id, ip, request, serial, status, url | dashboard:general:telemetry:write | 
|  | GET /devices/{serial}/liveTools/arpTable/{arpTableId} Return an ARP table live tool job.  >  getDeviceLiveToolsArpTable | serial, arpTableId | `` | address, arpTableId, entries, error, interface, ip, lastUpdatedAt, mac, request, serial, status, url, vlanId, vrfName | dashboard:general:telemetry:read | 
|  | POST /devices/{serial}/liveTools/cableTest Enqueue a job to perform a cable test for the device on the specified ports  >  createDeviceLiveToolsCableTest | serial | callback, httpServer, id, payloadTemplate, ports, sharedSecret, url | cableTestId, callback, id, ports, request, serial, status, url | dashboard:general:config:write | 
|  | GET /devices/{serial}/liveTools/cableTest/{id} Return a cable test live tool job.  >  getDeviceLiveToolsCableTest | serial, id | `` | cableTestId, error, index, lengthMeters, pairs, port, ports, request, results, serial, speedMbps, status, url | dashboard:general:config:read | 
|  | GET /devices/{serial}/liveTools/clients/disconnect/{id} Return a client disconnect job. (BETA)  >  getDeviceLiveToolsClientsDisconnect | serial, id | `` | error, id, mac, request, results, serial, status, success, url | wireless:config:read | 
|  | POST /devices/{serial}/liveTools/dhcpLeases Enqueue a job to perform a DHCP leases request for the device (BETA)  >  createDeviceLiveToolsDhcpLease | serial | callback, httpServer, id, payloadTemplate, sharedSecret, url | callback, dhcpLeasesId, id, request, serial, status, url | dashboard:general:telemetry:write | 
|  | GET /devices/{serial}/liveTools/dhcpLeases/{dhcpLeasesId} Return a DHCP leases live tool job. (BETA)  >  getDeviceLiveToolsDhcpLease | serial, dhcpLeasesId | `` | dhcpLeases, dhcpLeasesId, error, expiresAt, ip, mac, request, serial, status, url | dashboard:general:telemetry:read | 
|  | POST /devices/{serial}/liveTools/leds/blink Enqueue a job to blink LEDs on a device  >  createDeviceLiveToolsLedsBlink | serial | callback, duration, httpServer, id, payloadTemplate, sharedSecret, url | callback, duration, error, id, ledsBlinkId, request, serial, status, url | dashboard:general:config:write | 
