---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1000178168-67f5f3d1-example-for-configuring-vlan-id-based-227b0469
title: "Configure VLAN mapping on GE0/0/1."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1000178168-67f5f3d1-example-for-configuring-vlan-id-based--227b0469.md
source_anchor: ""
source_lines: [1, 29]
sha256: 97437ed7f6d5b631a7297a126c0d635432b9225c4bb3f328ae0d8abe34d5dddb
---

# Configure VLAN mapping on GE0/0/1.

This document describes the configuration of Ethernet services, including configuring link aggregation, VLANs, Voice VLAN, VLAN mapping, QinQ, GVRP, MAC table, STP/RSTP/MSTP, SEP, and so on.
Example for Configuring VLAN ID-based N:1 VLAN Mapping
Example for Configuring VLAN ID-based N:1 VLAN Mapping
Networking Requirements
In Figure 11-6, a large number of switches need to be deployed at the corridor so that the same service
used by different users can be sent on different VLANs. To save VLAN resources, configure the VLAN aggregation function (N:1) on the switches so that same services are sent on the same VLAN.
Figure 11-6 Networking diagram for configuring N:1 VLAN mapping
Configuration Roadmap
The configuration roadmap is as follows:
Create the original VLAN and the translated VLAN on the Switch and add GE0/0/1 to the VLANs in tagged mode.
[Switch] interface gigabitethernet0/0/1[Switch-GigabitEthernet0/0/1] port link-type hybrid[Switch-GigabitEthernet0/0/1] port hybrid tagged vlan 10 100 to 109
# Configure VLAN mapping on GE0/0/1.
[Switch-GigabitEthernet0/0/1] qinq vlan-translation enable[Switch-GigabitEthernet0/0/1] port vlan-mapping vlan 100 to 109 map-vlan 10
Verify the configuration.
Verify that users in VLAN 100 to VLAN 109 can connect to the Internet through the Switch.
Configuration Files
Switch configuration file
#
sysname Switch
#
vlan batch 10 100 to 109
#
interface gigabitethernet0/0/1 port link-type hybrid
qinq vlan-translation enable
port hybrid tagged vlan 10 100 to 109
port vlan-mapping vlan 100 to 109 map-vlan 10
#
return
We use cookies on this site, in order for the site to work properly and to analyse traffic, offer enhanced functionality and personalise content.Learn more
