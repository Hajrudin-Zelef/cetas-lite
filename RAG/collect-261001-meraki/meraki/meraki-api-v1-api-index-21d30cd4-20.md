---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-20
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [416, 443]
sha256: 881cb3880d6b31faa1bffd77efa1d1a7709464a77ee266a48c6c7f5880062844
---

# meraki-api-v1-api-index-21d30cd4

|  | GET /networks/{networkId}/sm/devices/{deviceId}/desktopLogs Return historical records of various Systems Manager network connection details for desktop devices.  >  getNetworkSmDeviceDesktopLogs | networkId, deviceId, perPage, startingAfter, endingBefore | `` | dhcpServer, dnsServer, gateway, ip, measuredAt, networkDevice, networkDriver, networkMTU, publicIP, subnet, ts, user, wifiAuth, wifiBssid, wifiChannel, wifiNoise, wifiRssi, wifiSsid | sm:telemetry:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/deviceCommandLogs Return historical records of commands sent to Systems Manager devices  >  getNetworkSmDeviceDeviceCommandLogs | networkId, deviceId, perPage, startingAfter, endingBefore | `` | action, dashboardUser, details, name, ts | sm:telemetry:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/deviceProfiles Get the installed profiles associated with a device  >  getNetworkSmDeviceDeviceProfiles | networkId, deviceId | `` | deviceId, id, isEncrypted, isManaged, name, profileData, profileIdentifier, version | sm:config:read | 
|  | POST /networks/{networkId}/sm/devices/{deviceId}/installApps Install applications on a device  >  installNetworkSmDeviceApps | networkId, deviceId | appIds, force | `` | `` | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/networkAdapters List the network adapters of a device  >  getNetworkSmDeviceNetworkAdapters | networkId, deviceId | `` | dhcpServer, dnsServer, gateway, id, ip, mac, name, subnet | sm:config:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/performanceHistory Return historical records of various Systems Manager client metrics for desktop devices.  >  getNetworkSmDevicePerformanceHistory | networkId, deviceId, perPage, startingAfter, endingBefore | `` | c, cpuPercentUsed, diskUsage, memActive, memFree, memInactive, memWired, networkReceived, networkSent, space, swapUsed, ts, used | sm:telemetry:read | 
|  | POST /networks/{networkId}/sm/devices/{deviceId}/refreshDetails Refresh the details of a device  >  refreshNetworkSmDeviceDetails | networkId, deviceId | `` | `` | sm:config:write | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/restrictions List the restrictions on a device  >  getNetworkSmDeviceRestrictions | networkId, deviceId | `` | profile, restrictions | sm:config:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/securityCenters List the security centers on a device  >  getNetworkSmDeviceSecurityCenters | networkId, deviceId | `` | antiVirusName, fireWallName, hasAntiVirus, hasFireWallInstalled, id, isAutoLoginDisabled, isDiskEncrypted, isFireWallEnabled, isRooted, runningProcs | sm:config:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/softwares Get a list of softwares associated with a device  >  getNetworkSmDeviceSoftwares | networkId, deviceId | `` | appId, bundleSize, createdAt, deviceId, dynamicSize, id, identifier, installedAt, iosRedemptionCode, isManaged, itunesId, licenseKey, name, path, redemptionCode, shortVersion, status, toInstall, toUninstall, uninstalledAt, updatedAt, vendor, version | sm:config:read | 
|  | POST /networks/{networkId}/sm/devices/{deviceId}/unenroll Unenroll a device  >  unenrollNetworkSmDevice | networkId, deviceId | `` | success | sm:config:write | 
|  | POST /networks/{networkId}/sm/devices/{deviceId}/uninstallApps Uninstall applications on a device  >  uninstallNetworkSmDeviceApps | networkId, deviceId | appIds | `` | `` | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/wlanLists List the saved SSID names on a device  >  getNetworkSmDeviceWlanLists | networkId, deviceId | `` | createdAt, id, xml | sm:config:read | 
|  | GET /networks/{networkId}/sm/profiles List profiles in a network  >  getNetworkSmProfiles | networkId, perPage, startingAfter, endingBefore, payloadTypes | `` | description, id, name, payloadTypes, scope, tags | sm:config:read | 
|  | GET /networks/{networkId}/sm/targetGroups List the target groups in this network  >  getNetworkSmTargetGroups | networkId, withDetails | `` | id, name, scope, tags | sm:config:read | 
|  | POST /networks/{networkId}/sm/targetGroups Add a target group  >  createNetworkSmTargetGroup | networkId | name, scope | id, name, scope, tags | sm:config:write | 
|  | GET /networks/{networkId}/sm/targetGroups/{targetGroupId} Return a target group  >  getNetworkSmTargetGroup | networkId, targetGroupId, withDetails | `` | id, name, scope, tags | sm:config:read | 
|  | PUT /networks/{networkId}/sm/targetGroups/{targetGroupId} Update a target group  >  updateNetworkSmTargetGroup | networkId, targetGroupId | name, scope | id, name, scope, tags | sm:config:write | 
|  | DELETE /networks/{networkId}/sm/targetGroups/{targetGroupId} Delete a target group from a network  >  deleteNetworkSmTargetGroup | networkId, targetGroupId | `` | `` | sm:config:write | 
|  | GET /networks/{networkId}/sm/trustedAccessConfigs List Trusted Access Configs  >  getNetworkSmTrustedAccessConfigs | networkId, perPage, startingAfter, endingBefore | `` | accessEndAt, accessStartAt, additionalEmailText, id, name, notifyTimeBeforeAccessEnds, scope, sendExpirationEmails, ssidName, tags, timeboundType | sm:config:read | 
|  | GET /networks/{networkId}/sm/userAccessDevices List User Access Devices and its Trusted Access Connections  >  getNetworkSmUserAccessDevices | networkId, perPage, startingAfter, endingBefore | `` | downloadedAt, email, id, lastConnectedAt, mac, name, scepCompletedAt, systemType, tags, trustedAccessConfigId, trustedAccessConnections, username | sm:config:read | 
|  | DELETE /networks/{networkId}/sm/userAccessDevices/{userAccessDeviceId} Delete a User Access Device  >  deleteNetworkSmUserAccessDevice | networkId, userAccessDeviceId | `` | `` | sm:config:write | 
|  | GET /networks/{networkId}/sm/users List the owners in an SM network with various specified fields and filters  >  getNetworkSmUsers | networkId, ids, usernames, emails, scope | `` | adGroups, asmGroups, azureAdGroups, displayName, email, fullName, hasIdentityCertificate, hasPassword, id, isExternal, samlGroups, tags, userThumbnail, username | sm:config:read | 
|  | GET /networks/{networkId}/sm/users/{userId}/deviceProfiles Get the profiles associated with a user  >  getNetworkSmUserDeviceProfiles | networkId, userId | `` | deviceId, id, isEncrypted, isManaged, name, profileData, profileIdentifier, version | sm:config:read | 
|  | GET /networks/{networkId}/sm/users/{userId}/softwares Get a list of softwares associated with a user  >  getNetworkSmUserSoftwares | networkId, userId | `` | appId, bundleSize, createdAt, deviceId, dynamicSize, id, identifier, installedAt, iosRedemptionCode, isManaged, itunesId, licenseKey, name, path, redemptionCode, shortVersion, status, toInstall, toUninstall, uninstalledAt, updatedAt, vendor, version | sm:config:read | 
|  | GET /networks/{networkId}/snmp Return the SNMP settings for a network  >  getNetworkSnmp | networkId | `` | access, authentication, communityString, passphrase, privacy, protocol, username, users | dashboard:general:telemetry:read | 
|  | PUT /networks/{networkId}/snmp Update the SNMP settings for a network  >  updateNetworkSnmp | networkId | access, authentication, communityString, passphrase, privacy, protocol, username, users | access, authentication, communityString, passphrase, privacy, protocol, username, users | dashboard:general:telemetry:write | 
|  | PUT /networks/{networkId}/snmp/traps Update the SNMP trap configuration for the specified network (BETA)  >  updateNetworkSnmpTraps | networkId | address, community, mode, name, passphrase, port, receiver, users, v2, v3 | address, community, id, mode, name, network, port, receiver, users, v2, v3 | `` | 
