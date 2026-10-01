---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-19
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [389, 415]
sha256: a4440593e40166afeb5c60911d407ae54e03635730e30110a1e344e40511dcd6
---

# meraki-api-v1-api-index-21d30cd4

|  | GET /networks/{networkId}/sensor/alerts/profiles/{id} Show details of a sensor alert profile for a network.  >  getNetworkSensorAlertsProfile | networkId, id | `` | ambient, apparentPower, celsius, co2, concentration, conditions, current, direction, door, draw, duration, emails, fahrenheit, frequency, httpServerIds, humidity, id, includeSensorUrl, indoorAirQuality, level, message, metric, name, noise, open, outageDetected, percentage, pm25, powerFactor, present, profileId, quality, realPower, recipients, relativePercentage, schedule, score, serials, smsNumbers, temperature, threshold, tvoc, upstreamPower, voltage, water | sensor:telemetry:read | 
|  | PUT /networks/{networkId}/sensor/alerts/profiles/{id} Updates a sensor alert profile for a network.  >  updateNetworkSensorAlertsProfile | networkId, id | ambient, apparentPower, celsius, co2, concentration, conditions, current, direction, door, draw, duration, emails, fahrenheit, frequency, httpServerIds, humidity, id, includeSensorUrl, indoorAirQuality, level, message, metric, name, noise, open, outageDetected, percentage, pm25, powerFactor, present, quality, realPower, recipients, relativePercentage, schedule, score, serials, smsNumbers, temperature, threshold, tvoc, upstreamPower, voltage, water | ambient, apparentPower, celsius, co2, concentration, conditions, current, direction, door, draw, duration, emails, fahrenheit, frequency, httpServerIds, humidity, id, includeSensorUrl, indoorAirQuality, level, message, metric, name, noise, open, outageDetected, percentage, pm25, powerFactor, present, profileId, quality, realPower, recipients, relativePercentage, schedule, score, serials, smsNumbers, temperature, threshold, tvoc, upstreamPower, voltage, water | sensor:telemetry:write | 
|  | DELETE /networks/{networkId}/sensor/alerts/profiles/{id} Deletes a sensor alert profile from a network.  >  deleteNetworkSensorAlertsProfile | networkId, id | `` | `` | sensor:telemetry:write | 
|  | GET /networks/{networkId}/sensor/mqttBrokers List the sensor settings of all MQTT brokers for this network  >  getNetworkSensorMqttBrokers | networkId | `` | enabled, mqttBrokerId | sensor:telemetry:read | 
|  | GET /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} Return the sensor settings of an MQTT broker  >  getNetworkSensorMqttBroker | networkId, mqttBrokerId | `` | enabled, mqttBrokerId | sensor:telemetry:read | 
|  | PUT /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} Update the sensor settings of an MQTT broker  >  updateNetworkSensorMqttBroker | networkId, mqttBrokerId | enabled | enabled, mqttBrokerId | sensor:telemetry:write | 
|  | GET /networks/{networkId}/sensor/relationships List the sensor roles for devices in a given network  >  getNetworkSensorRelationships | networkId | `` | device, livestream, name, productType, relatedDevices, relationships, serial | sensor:config:read | 
|  | GET /networks/{networkId}/sensor/schedules Returns a list of all sensor schedules. (BETA)  >  getNetworkSensorSchedules | networkId | `` | id, name | `` | 
|  | GET /networks/{networkId}/settings Return the settings for a network  >  getNetworkSettings | networkId | `` | authentication, enabled, fips, localStatusPage, localStatusPageEnabled, namedVlans, remoteStatusPageEnabled, securePort, username | dashboard:general:config:read | 
|  | PUT /networks/{networkId}/settings Update the settings for a network  >  updateNetworkSettings | networkId | authentication, enabled, fips, localStatusPage, localStatusPageEnabled, namedVlans, password, remoteStatusPageEnabled, securePort, username | authentication, enabled, fips, localStatusPage, localStatusPageEnabled, namedVlans, remoteStatusPageEnabled, securePort, username | dashboard:general:config:write | 
|  | POST /networks/{networkId}/sites/buildings Create a new building (BETA)  >  createNetworkSitesBuilding | networkId | floors, id, name, number, plan | buildingId, counts, floors, id, items, meta, name, network, number, plan, remaining, total | dashboard:general:config:write | 
|  | DELETE /networks/{networkId}/sites/buildings/{buildingId} Delete a building (BETA)  >  deleteNetworkSitesBuilding | networkId, buildingId | `` | `` | dashboard:general:config:write | 
|  | PUT /networks/{networkId}/sites/buildings/{buildingId} Update a building (BETA)  >  updateNetworkSitesBuilding | networkId, buildingId | floors, id, name, number, plan | buildingId, counts, floors, id, items, meta, name, network, number, plan, remaining, total | dashboard:general:config:write | 
|  | POST /networks/{networkId}/sm/bypassActivationLockAttempts Bypass activation lock attempt  >  createNetworkSmBypassActivationLockAttempt | networkId | ids | `` | `` | 
|  | GET /networks/{networkId}/sm/bypassActivationLockAttempts/{attemptId} Bypass activation lock attempt status  >  getNetworkSmBypassActivationLockAttempt | networkId, attemptId | `` | `` | `` | 
|  | GET /networks/{networkId}/sm/devices List the devices enrolled in an SM network with various specified fields and filters  >  getNetworkSmDevices | networkId, fields, wifiMacs, serials, ids, uuids, systemTypes, scope, perPage, startingAfter, endingBefore | `` | id, ip, name, notes, osName, serial, serialNumber, ssid, systemModel, tags, uuid, wifiMac | sm:config:read | 
|  | POST /networks/{networkId}/sm/devices/checkin Force check-in a set of devices  >  checkinNetworkSmDevices | networkId | ids, scope, serials, wifiMacs | ids | sm:config:write | 
|  | PUT /networks/{networkId}/sm/devices/fields Modify the fields of a device  >  updateNetworkSmDevicesFields | networkId | deviceFields, id, name, notes, serial, wifiMac | id, name, notes, serial, wifiMac | sm:config:write | 
|  | POST /networks/{networkId}/sm/devices/lock Lock a set of devices  >  lockNetworkSmDevices | networkId | ids, pin, scope, serials, wifiMacs | ids | sm:config:write | 
|  | POST /networks/{networkId}/sm/devices/modifyTags Add, delete, or update the tags of a set of devices  >  modifyNetworkSmDevicesTags | networkId | ids, scope, serials, tags, updateAction, wifiMacs | id, serial, tags, wifiMac | sm:config:write | 
|  | POST /networks/{networkId}/sm/devices/move Move a set of devices to a new network  >  moveNetworkSmDevices | networkId | ids, newNetwork, scope, serials, wifiMacs | ids, newNetwork | sm:config:write | 
|  | POST /networks/{networkId}/sm/devices/reboot Reboot a set of endpoints  >  rebootNetworkSmDevices | networkId | ids, kextPaths, notifyUser, rebuildKernelCache, requestRequiresNetworkTether, scope, serials, wifiMacs | ids | `` | 
|  | POST /networks/{networkId}/sm/devices/shutdown Shutdown a set of endpoints  >  shutdownNetworkSmDevices | networkId | ids, scope, serials, wifiMacs | ids | `` | 
|  | POST /networks/{networkId}/sm/devices/wipe Wipe a device  >  wipeNetworkSmDevices | networkId | id, pin, serial, wifiMac | id | sm:config:write | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory Return the client's daily cellular data usage history  >  getNetworkSmDeviceCellularUsageHistory | networkId, deviceId | `` | received, sent, ts | sm:telemetry:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/certs List the certs on a device  >  getNetworkSmDeviceCerts | networkId, deviceId | `` | certPem, deviceId, id, issuer, name, notValidAfter, notValidBefore, subject | sm:config:read | 
|  | GET /networks/{networkId}/sm/devices/{deviceId}/connectivity Returns historical connectivity data (whether a device is regularly checking in to Dashboard).  >  getNetworkSmDeviceConnectivity | networkId, deviceId, perPage, startingAfter, endingBefore | `` | firstSeenAt, lastSeenAt | sm:telemetry:read | 
