---
id: collect-261001-meraki/meraki/meraki-api-v1-api-index-21d30cd4-80
title: "meraki-api-v1-api-index-21d30cd4"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-api-v1-api-index-21d30cd4.md
source_anchor: ""
source_lines: [1542, 1546]
sha256: 8c7f4dbbaad995d5454c944d5d1b84b401b22329b8c306c47a7f8be0098bef77
---

# meraki-api-v1-api-index-21d30cd4

|  | GET /organizations/{organizationId}/wirelessController/devices/redundancy/failover/history List the failover events of wireless LAN controllers in an organization  >  getOrganizationWirelessControllerDevicesRedundancyFailoverHistory | organizationId, serials, t0, t1, timespan, perPage, startingAfter, endingBefore | `` | active, chassis, counts, failed, items, meta, name, reason, remaining, serial, total, ts | `` | 
|  | GET /organizations/{organizationId}/wirelessController/devices/redundancy/statuses List redundancy details of wireless LAN controllers in an organization  >  getOrganizationWirelessControllerDevicesRedundancyStatuses | organizationId, serials, perPage, startingAfter, endingBefore | `` | counts, enabled, failover, items, last, meta, mobilityMac, mode, reason, remaining, serial, total, ts | `` | 
|  | GET /organizations/{organizationId}/wirelessController/devices/system/utilization/history/byInterval List cpu utilization data of wireless LAN controllers in an organization  >  getOrganizationWirelessControllerDevicesSystemUtilizationHistoryByInterval | organizationId, serials, t0, t1, timespan, perPage, startingAfter, endingBefore | `` | average, byCore, counts, endTs, intervals, items, meta, name, overall, percentage, remaining, serial, startTs, total, usage | `` | 
|  | GET /organizations/{organizationId}/wirelessController/overview/byDevice List the overview information of wireless LAN controllers in an organization and it is updated every minute.  >  getOrganizationWirelessControllerOverviewByDevice | organizationId, networkIds, serials, perPage, startingAfter, endingBefore | `` | address, addresses, byStatus, chassisName, clients, connections, counts, firmware, id, items, management, meta, network, offline, online, redundancy, redundantSerial, remaining, role, serial, shortName, total, version | `` | 
|  | POST /organizations/{organizationId}/wirelessController/regulatoryDomain/package/generate Generate the regulatory domain package (BETA)  >  generateOrganizationWirelessControllerRegulatoryDomainPackage | organizationId | networkIds | certificates, content, country, counts, createdAt, createdBy, details, devices, email, id, mac, method, organizationId, purpose, regulatoryDomain, schemaVersion, serial, signature | `` |
