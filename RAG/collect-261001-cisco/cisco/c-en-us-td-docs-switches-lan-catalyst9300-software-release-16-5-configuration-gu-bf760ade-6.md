---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade-6
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade.md
source_anchor: ""
source_lines: [294, 328]
sha256: 3b54c4e0db7537c334d8a7e1295389bfeedb02e1c105a13f930658baac626034
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade

This example shows how to enable port security on a port and to set the maximum number of secure addresses to 50. The violation mode is the default, no static secure MAC addresses are configured, and sticky learning is enabled.
Device(config)# interface gigabitethernet1/0/1
Device(config-if)# switchport mode access
Device(config-if)# switchport port-security
Device(config-if)# switchport port-security maximum 50
Device(config-if)# switchport port-security mac-address sticky
This example shows how to configure a static secure MAC address on VLAN 3 on a port:
Device(config)# interface gigabitethernet1/0/2
Device(config-if)# switchport mode trunk
Device(config-if)# switchport port-security
Device(config-if)# switchport port-security mac-address 0000.0200.0004 vlan 3
This example shows how to enable sticky port security on a port, to manually configure MAC addresses for data VLAN and voice VLAN, and to set the total maximum number of secure addresses to 20 (10 for data VLAN and 10 for voice VLAN).
Device(config)# interface tengigabitethernet1/0/1
Device(config-if)# switchport access vlan 21
Device(config-if)# switchport mode access
Device(config-if)# switchport voice vlan 22
Device(config-if)# switchport port-security
Device(config-if)# switchport port-security maximum 20
Device(config-if)# switchport port-security violation restrict
Device(config-if)# switchport port-security mac-address sticky
Device(config-if)# switchport port-security mac-address sticky 0000.0000.0002
Device(config-if)# switchport port-security mac-address 0000.0000.0003
Device(config-if)# switchport port-security mac-address sticky 0000.0000.0001 vlan voice
Device(config-if)# switchport port-security mac-address 0000.0000.0004 vlan voice
Device(config-if)# switchport port-security maximum 10 vlan access
Device(config-if)# switchport port-security maximum 10 vlan voice
Additional References for Port-Based Traffic Control
MIBs
| MIB | MIBs Link | 
|---|---|
| All the supported MIBs for this release. | To locate and download MIBs for selected platforms, Cisco IOS releases, and feature sets, use Cisco MIB Locator found at the following URL: http://www.cisco.com/go/mibs | 
Technical Assistance
| Description | Link | 
|---|---|
| The Cisco Support website provides extensive online resources, including documentation and tools for troubleshooting and resolving technical issues with Cisco products and technologies. To receive security and technical information about your products, you can subscribe to various services, such as the Product Alert Tool (accessed from Field Notices), the Cisco Technical Services Newsletter, and Really Simple Syndication (RSS) Feeds. Access to most tools on the Cisco Support website requires a Cisco.com user ID and password. | http://www.cisco.com/support |
