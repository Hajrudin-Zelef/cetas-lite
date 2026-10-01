---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-2
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [26, 50]
sha256: 02131f341909dab8d361b3f46631330f5fcabe2fdc3aaf49259530687e43130b
---

# meraki-api-v1-api-index-21d30cd4

|  | PUT /devices/{serial} Update the attributes of a device  >  updateDevice | serial | address, floorPlanId, lat, lng, moveMapMarker, name, notes, switchProfileId, tags | address, beaconIdParams, details, firmware, floorPlanId, lanIp, lat, lng, mac, major, minor, model, name, networkId, notes, serial, tags, url, uuid, value | dashboard:general:config:write | 
|  | GET /devices/{serial}/appliance/dhcp/subnets Return the DHCP subnet information for an appliance  >  getDeviceApplianceDhcpSubnets | serial | `` | freeCount, subnet, usedCount, vlanId | sdwan:telemetry:read | 
|  | POST /devices/{serial}/appliance/interfaces/ports/update Update configurations for an appliance's specified port  >  createDeviceApplianceInterfacesPortsUpdate | serial | access, allowedVlans, downlink, duplex, enabled, id, interface, layer, mode, nativeVlan, number, personality, policy, profile, sgt, slot, speed, subslot, trunk, type, uplink, vlan | access, allowedVlans, downlink, enabled, id, interface, isFlexible, layer, mode, name, nativeVlan, number, personality, policy, primary, profile, sgt, slot, subslot, trunk, type, uplink, vlan | `` | 
|  | PUT /devices/{serial}/appliance/interfaces/ports/{number} Update configurations for an appliance's specified port  >  updateDeviceApplianceInterfacesPort | serial, number | access, allowedVlans, downlink, duplex, enabled, id, layer, mode, nativeVlan, personality, policy, profile, sgt, speed, trunk, type, uplink, vlan | access, allowedVlans, downlink, enabled, id, interface, isFlexible, layer, mode, name, nativeVlan, number, personality, policy, primary, profile, sgt, slot, subslot, trunk, type, uplink, vlan | `` | 
|  | GET /devices/{serial}/appliance/performance Return the performance score for a single Secure Appliance or Secure Router  >  getDeviceAppliancePerformance | serial, t0, t1, timespan | `` | perfScore | sdwan:telemetry:read | 
|  | GET /devices/{serial}/appliance/prefixes/delegated Return current delegated IPv6 prefixes on an appliance.  >  getDeviceAppliancePrefixesDelegated | serial | `` | assigned, available, counts, description, expiresAt, interface, isPreferred, method, origin, prefix, staticDelegatedPrefixId | sdwan:telemetry:read | 
|  | GET /devices/{serial}/appliance/prefixes/delegated/vlanAssignments Return prefixes assigned to all IPv6 enabled VLANs on an appliance.  >  getDeviceAppliancePrefixesDelegatedVlanAssignments | serial | `` | address, id, interface, ipv6, linkLocal, name, origin, prefix, solicitedNodeMulticast, status, vlan | sdwan:telemetry:read | 
|  | GET /devices/{serial}/appliance/radio/settings Return the radio settings of an appliance  >  getDeviceApplianceRadioSettings | serial | `` | channel, channelWidth, fiveGhzSettings, rfProfileId, serial, targetPower, twoFourGhzSettings | sdwan:config:read | 
|  | PUT /devices/{serial}/appliance/radio/settings Update the radio settings of an appliance  >  updateDeviceApplianceRadioSettings | serial | channel, channelWidth, fiveGhzSettings, rfProfileId, targetPower, twoFourGhzSettings | channel, channelWidth, fiveGhzSettings, rfProfileId, serial, targetPower, twoFourGhzSettings | sdwan:config:write | 
|  | GET /devices/{serial}/appliance/uplinks/settings Return the uplink settings for a secure router or security appliance  >  getDeviceApplianceUplinksSettings | serial | `` | address, addresses, assignmentMode, authentication, borderRelay, enabled, gateway, interfaceId, interfaces, ipv4, ipv6, isp, local, mode, name, nameservers, peerSgtCapable, pppoe, sgt, svis, transition, username, vlanId, vlanTagging, wan1, wan2 | sdwan:config:read | 
|  | PUT /devices/{serial}/appliance/uplinks/settings Update the uplink settings for a secure router or security appliance  >  updateDeviceApplianceUplinksSettings | serial | address, addresses, assignmentMode, authentication, borderRelay, enabled, gateway, interfaceId, interfaces, ipv4, ipv6, isp, local, mode, name, nameservers, password, peerSgtCapable, pppoe, sgt, svis, transition, username, vlanId, vlanTagging, wan1, wan2 | address, addresses, assignmentMode, authentication, borderRelay, enabled, gateway, interfaceId, interfaces, ipv4, ipv6, isp, local, mode, name, nameservers, peerSgtCapable, pppoe, sgt, svis, transition, username, vlanId, vlanTagging, wan1, wan2 | sdwan:config:write | 
|  | POST /devices/{serial}/appliance/vmx/authenticationToken Generate a new vMX authentication token  >  createDeviceApplianceVmxAuthenticationToken | serial | `` | expiresAt, token | sdwan:config:write | 
|  | POST /devices/{serial}/blinkLeds Blink the LEDs on a device (DEPRECATED)  >  blinkDeviceLeds | serial | duration, duty, period | duration, duty, period | dashboard:general:config:write | 
|  | GET /devices/{serial}/camera/analytics/live Returns live state from camera analytics zones (DEPRECATED)  >  getDeviceCameraAnalyticsLive | serial | `` | person, ts, zoneId, zones | `` | 
|  | GET /devices/{serial}/camera/analytics/overview Returns an overview of aggregate analytics data for a timespan (DEPRECATED)  >  getDeviceCameraAnalyticsOverview | serial, t0, t1, timespan, objectType | `` | averageCount, endTs, entrances, startTs, zoneId | `` | 
|  | GET /devices/{serial}/camera/analytics/recent Returns most recent record for analytics zones (DEPRECATED)  >  getDeviceCameraAnalyticsRecent | serial, objectType | `` | averageCount, endTs, entrances, startTs, zoneId | `` | 
|  | GET /devices/{serial}/camera/analytics/zones Returns all configured analytic zones for this camera (DEPRECATED)  >  getDeviceCameraAnalyticsZones | serial | `` | id, label, regionOfInterest, type, x0, x1, y0, y1 | camera:config:read | 
|  | GET /devices/{serial}/camera/analytics/zones/{zoneId}/history Return historical records for analytic zones (DEPRECATED)  >  getDeviceCameraAnalyticsZoneHistory | serial, zoneId, t0, t1, timespan, resolution, objectType | `` | averageCount, endTs, entrances, startTs | `` | 
|  | GET /devices/{serial}/camera/clip Generate a video clip of up to 5 minutes long.  >  clipDeviceCamera | serial, startTimestamp, endTimestamp, imagerId | `` | expiry, url | camera:telemetry:write | 
|  | GET /devices/{serial}/camera/customAnalytics Return custom analytics settings for a camera  >  getDeviceCameraCustomAnalytics | serial | `` | artifactId, enabled, name, parameters, value | camera:config:read | 
|  | PUT /devices/{serial}/camera/customAnalytics Update custom analytics settings for a camera  >  updateDeviceCameraCustomAnalytics | serial | artifactId, enabled, name, parameters, value | artifactId, enabled, name, parameters, value | camera:config:write | 
|  | POST /devices/{serial}/camera/generateSnapshot Generate a snapshot of what the camera sees at the specified time and return a link to that image.  >  generateDeviceCameraSnapshot | serial | fullframe, timestamp | expiry, url | camera:telemetry:write | 
|  | GET /devices/{serial}/camera/qualityAndRetention Returns quality and retention settings for the given camera  >  getDeviceCameraQualityAndRetention | serial | `` | audioRecordingEnabled, motionBasedRetentionEnabled, motionDetectorVersion, profileId, quality, resolution, restrictedBandwidthModeEnabled | camera:config:read | 
|  | PUT /devices/{serial}/camera/qualityAndRetention Update quality and retention settings for the given camera  >  updateDeviceCameraQualityAndRetention | serial | audioRecordingEnabled, motionBasedRetentionEnabled, motionDetectorVersion, profileId, quality, resolution, restrictedBandwidthModeEnabled | audioRecordingEnabled, motionBasedRetentionEnabled, motionDetectorVersion, profileId, quality, resolution, restrictedBandwidthModeEnabled | camera:config:write | 
|  | GET /devices/{serial}/camera/sense Returns sense settings for a given camera  >  getDeviceCameraSense | serial | `` | audioDetection, detectionModelId, enabled, mqttBrokerId, mqttTopics, senseEnabled | camera:config:read | 
