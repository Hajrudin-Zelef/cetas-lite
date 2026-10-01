---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-12
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [663, 694]
sha256: 24d7416d66d9c7698882e62ddc03db63384f6568415c52bf8e1817e66654157a
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

| show ip ospf neighbors [ neighbor-id ] [ detail ] [ interface - type number ] [ vrf { vrf-name \| all \| default \| management }] [ summary ] | Displays the list of OSPFv2 neighbors. | 
| show ip ospf request-list neighbor-id interface - type number [ vrf { vrf-name \| all \| default \| management }] | Displays the list of OSPFv2 link-state requests. | 
| show ip ospf retransmission-list neighbor-id interface - type number [ vrf { vrf-name \| all \| default \| management }] | Displays the list of OSPFv2 link-state retransmissions. | 
| show ip ospf route [ ospf-route ] [ summary ] [ vrf { vrf-name \| all \| default \| management }] | Displays the internal OSPFv2 routes. | 
| show ip ospf summary-address [ vrf { vrf-name \| all \| default \| management }] | Displays information about the OSPFv2 summary addresses. | 
| show ip ospf virtual-links [ brief ] [ vrf { vrf-name \| all \| default \| management }] | Displays information about OSPFv2 virtual links. | 
| show ip ospf vrf { vrf-name \| all \| default \| management } | Displays information about the VRF-based OSPFv2 configuration. | 
| show running-configuration ospf | Displays the current running OSPFv2 configuration. | 
Monitoring OSPFv2
To display OSPFv2 statistics, use the following commands:
|  |  | 
|---|---|
| show ip ospf policy statistics area area-id filter-list { in \| out } [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 route policy statistics for an area. | 
| show ip ospf policy statistics redistribute { bgp id \| direct \| eigrp id \| isis id \| ospf id \| rip id \| static } [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 route policy statistics. | 
| show ip ospf statistics [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 event counters. | 
| show ip ospf traffic [ interface - type number ] [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 packet counters. | 
Configuration Examples for OSPFv2
The following example shows how to configure OSPFv2:
OSPF RFC Compatibility Mode Example
The following example shows how to configure OSPF to be compatible with routers that comply with RFC 1583:
Note You must configure RFC 1583 compatibility on any VRF that connects to routers running only RFC1583 compatible OSPF.
Additional References
For additional information related to implementing OSPF, see the following sections:
Related Documents
|  |  | 
|---|---|
| OSPFv3 for IPv6 networks | Chapter 6, “Configuring OSPFv3” | 
| Route maps | Chapter 15, “Configuring Route Policy Manager” | 
MIBs
|  |  | 
|---|---|
| MIBs related to OSPFv2 | To locate and download supported MIBs, go to the following URL: ftp://ftp.cisco.com/pub/mibs/supportlists/nexus9000/Nexus9000MIBSupportList.html |
