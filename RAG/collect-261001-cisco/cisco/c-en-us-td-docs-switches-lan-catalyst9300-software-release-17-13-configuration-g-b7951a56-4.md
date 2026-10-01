---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56.md
source_anchor: ""
source_lines: [184, 209]
sha256: 3cc902fb4b4a5babb4d457e9826a61653af3d71a72ad45d9f12a39114c0f2e9a
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56

This example shows how to enable sticky port security on a port, to manually configure MAC addresses for data VLAN and voice VLAN, and to set the total maximum number of secure addresses to 20 (10 for data VLAN and 10 for voice VLAN).
Device> enable
Device# configure terminal
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
Device(config-if)# end
This table provides release and related information for features explained in this module.
These features are available on all releases subsequent to the one they were introduced in, unless noted otherwise.
| Release | Feature | Feature Information | 
|---|---|---|
| Cisco IOS XE Everest 16.5.1a | Port Security | The Port Security feature restricts input to an interface by limiting and identifying MAC addresses of the stations allowed to access the port. | 
| Cisco IOS XE Everest 16.5.1a | Port Security MAC Aging | When devices are added or removed from a network, the device updates the address table, adding new dynamic addresses and aging out those that are not in use. | 
| Cisco IOS XE 17.13.1 | Management Traffic Protocol | Management traffic control feature allows traffic to enter through a user-defined physical interface and restricts traffic to any other interfaces that is not defined by the user. When the network protocol and the IP address is set according to the user’s preference, traffic flow is allowed only through the defined interface. | 
Use Cisco Feature Navigator to find information about platform and software image support. To access Cisco Feature Navigator, go to https://cfnng.cisco.com.
