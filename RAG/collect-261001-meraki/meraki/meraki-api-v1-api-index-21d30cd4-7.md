---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-7
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [138, 163]
sha256: ab7adfeaff70dc1dda2cbd49bc078a5adb7028a21b99cf09863dc70797846a0f
---

# meraki-api-v1-api-index-21d30cd4

|  | POST /devices/{serial}/switch/routing/staticRoutes Create a layer 3 static route for a switch  >  createDeviceSwitchRoutingStaticRoute | serial | advertiseViaOspfEnabled, leakRouteToDefaultVrf, name, nextHopIp, preferOverOspfRoutesEnabled, subnet, vrf | advertiseViaOspfEnabled, leakRouteToDefaultVrf, managementNextHop, name, nextHopIp, preferOverOspfRoutesEnabled, staticRouteId, subnet, vrf | switch:config:write | 
|  | GET /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} Return a layer 3 static route for a switch  >  getDeviceSwitchRoutingStaticRoute | serial, staticRouteId | `` | advertiseViaOspfEnabled, leakRouteToDefaultVrf, managementNextHop, name, nextHopIp, preferOverOspfRoutesEnabled, staticRouteId, subnet, vrf | switch:config:read | 
|  | PUT /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} Update a layer 3 static route for a switch  >  updateDeviceSwitchRoutingStaticRoute | serial, staticRouteId | advertiseViaOspfEnabled, leakRouteToDefaultVrf, managementNextHop, name, nextHopIp, preferOverOspfRoutesEnabled, subnet, vrf | advertiseViaOspfEnabled, leakRouteToDefaultVrf, managementNextHop, name, nextHopIp, preferOverOspfRoutesEnabled, staticRouteId, subnet, vrf | switch:config:write | 
|  | DELETE /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} Delete a layer 3 static route for a switch  >  deleteDeviceSwitchRoutingStaticRoute | serial, staticRouteId | `` | `` | switch:config:write | 
|  | GET /devices/{serial}/switch/warmSpare Return warm spare configuration for a switch  >  getDeviceSwitchWarmSpare | serial | `` | enabled, primarySerial, spareSerial | switch:config:read | 
|  | PUT /devices/{serial}/switch/warmSpare Update warm spare configuration for a switch  >  updateDeviceSwitchWarmSpare | serial | enabled, spareSerial | enabled, primarySerial, spareSerial | switch:config:write | 
|  | PUT /devices/{serial}/wireless/alternateManagementInterface/ipv6 Update alternate management interface IPv6 address  >  updateDeviceWirelessAlternateManagementInterfaceIpv6 | serial | address, addresses, assignmentMode, gateway, nameservers, prefix, protocol | address, addresses, assignmentMode, gateway, nameservers, prefix, protocol | wireless:config:write | 
|  | GET /devices/{serial}/wireless/bluetooth/settings Return the bluetooth settings for a wireless device  >  getDeviceWirelessBluetoothSettings | serial | `` | advertised, interval, major, minor, power, transmit, uuid | wireless:config:read | 
|  | PUT /devices/{serial}/wireless/bluetooth/settings Update the bluetooth settings for a wireless device  >  updateDeviceWirelessBluetoothSettings | serial | advertised, interval, major, minor, power, transmit, uuid | advertised, interval, major, minor, power, transmit, uuid | wireless:config:write | 
|  | GET /devices/{serial}/wireless/connectionStats Aggregated connectivity info for a given AP on this network  >  getDeviceWirelessConnectionStats | serial, t0, t1, timespan, band, ssid, apTag | `` | assoc, auth, connectionStats, dhcp, dns, serial, success | wireless:telemetry:read | 
|  | GET /devices/{serial}/wireless/electronicShelfLabel Return the ESL settings of a device  >  getDeviceWirelessElectronicShelfLabel | serial | `` | apEslId, channel, enabled, hostname, networkId, provider, serial | `` | 
|  | PUT /devices/{serial}/wireless/electronicShelfLabel Update the ESL settings of a device  >  updateDeviceWirelessElectronicShelfLabel | serial | channel, enabled | apEslId, channel, enabled, hostname, networkId, provider, serial | `` | 
|  | GET /devices/{serial}/wireless/healthScores Fetch the health scores for a given AP on this network (BETA)  >  getDeviceWirelessHealthScores | serial | `` | device, latest, onboarding, performance, serial | `` | 
|  | GET /devices/{serial}/wireless/latencyStats Aggregated latency info for a given AP on this network  >  getDeviceWirelessLatencyStats | serial, t0, t1, timespan, band, ssid, apTag, vlan, fields | `` | avg, backgroundTraffic, bestEffortTraffic, latencyStats, rawDistribution, serial, videoTraffic, voiceTraffic | wireless:telemetry:read | 
|  | GET /devices/{serial}/wireless/radio/afc/position Return the position for a wireless device (BETA)  >  getDeviceWirelessRadioAfcPosition | serial | `` | aboveFloor, aboveGround, antenna, cableLength, gps, height, id, name, network, serial, uncertainty, value | `` | 
|  | PUT /devices/{serial}/wireless/radio/afc/position Update the position attributes for this device (BETA)  >  updateDeviceWirelessRadioAfcPosition | serial | aboveFloor, aboveGround, antenna, cableLength, gps, height, uncertainty, value | aboveFloor, aboveGround, antenna, cableLength, gps, height, id, name, network, serial, uncertainty, value | `` | 
|  | GET /devices/{serial}/wireless/radio/afc/powerLimits Return the AFC power limits for a wireless device (BETA)  >  getDeviceWirelessRadioAfcPowerLimits | serial | `` | byChannel, channel, channelWidth, expiresAt, id, lastSuccessAt, lastUpdatedAt, lat, limit, lng, location, name, network, serial, status, type, uncertainty | `` | 
|  | GET /devices/{serial}/wireless/radio/overrides Return the radio overrides of a device  >  getDeviceWirelessRadioOverrides | serial | `` | band, channel, channelWidth, enabled, id, index, network, radios, rfProfile, serial, targetPower | `` | 
|  | PUT /devices/{serial}/wireless/radio/overrides Update 2.4 GHz, 5 GHz, and 6 GHz radio settings (channel, channel width, power, and enable/disable) that override RF profiles  >  updateDeviceWirelessRadioOverrides | serial | channel, channelWidth, enabled, id, index, radios, rfProfile, targetPower | band, channel, channelWidth, enabled, id, index, network, radios, rfProfile, serial, targetPower | `` | 
|  | GET /devices/{serial}/wireless/radio/settings Return the manually configured radio settings overrides of a device, which take precedence over RF profiles.  >  getDeviceWirelessRadioSettings | serial | `` | channel, channelWidth, fiveGhzSettings, rfProfileId, serial, targetPower, twoFourGhzSettings | wireless:config:read | 
|  | PUT /devices/{serial}/wireless/radio/settings Update 2.4 GHz and 5 GHz radio settings (channel, channel width, power) that override RF profiles  >  updateDeviceWirelessRadioSettings | serial | channel, channelWidth, fiveGhzSettings, rfProfileId, targetPower, twoFourGhzSettings | channel, channelWidth, fiveGhzSettings, rfProfileId, serial, targetPower, twoFourGhzSettings | wireless:config:write | 
|  | GET /devices/{serial}/wireless/radio/status Show the status of this device's radios (BETA)  >  getDeviceWirelessRadioStatus | serial | `` | band, channel, dfs, mode, power, radarDetected, status, transmit, value, width | `` | 
|  | GET /devices/{serial}/wireless/status Return the SSID statuses of an access point  >  getDeviceWirelessStatus | serial | `` | band, basicServiceSets, broadcasting, bssid, channel, channelWidth, enabled, power, ssidName, ssidNumber, visible | wireless:telemetry:read | 
|  | POST /devices/{serial}/wireless/zigbee/enrollments Enqueue a job to start enrolling door locks on zigbee configured wireless devices  >  createDeviceWirelessZigbeeEnrollment | serial | `` | enrollmentId, request, serial, status, url | `` | 
|  | GET /devices/{serial}/wireless/zigbee/enrollments/{enrollmentId} Return an enrollment  >  getDeviceWirelessZigbeeEnrollment | serial, enrollmentId | `` | doorLockId, doorLocks, enrolledAt, enrollmentId, enrollmentStartedAt, eui64, gateway, id, lastSeenAt, lqi, name, network, request, rssi, serial, shortId, status, url | `` | 
|  | GET /networks/{networkId} Return a network  >  getNetwork | networkId | `` | details, enrollmentString, id, isBoundToConfigTemplate, name, notes, organizationId, productType, productTypes, tags, timeZone, url, value | dashboard:general:config:read | 
