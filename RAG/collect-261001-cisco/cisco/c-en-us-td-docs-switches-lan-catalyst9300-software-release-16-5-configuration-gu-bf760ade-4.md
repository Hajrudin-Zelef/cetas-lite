---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade.md
source_anchor: ""
source_lines: [199, 240]
sha256: 6ef45d3728c126b37ff325fab00395059406f4975c445fb5a9b4c649dcbba428
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-bf760ade

                                       Voice VLAN is only supported on access ports and not on trunk ports, even though the configuration is allowed.
- 
                                       					
                                       When you enable port security on an interface that is also configured with a voice VLAN, set the maximum allowed secure addresses on the port to two. When the port is connected to a Cisco IP phone, the IP phone requires one MAC address. The Cisco IP phone address is learned on the voice VLAN, but is not learned on the access VLAN. If you connect a single PC to the Cisco IP phone, no additional MAC addresses are required. If you connect more than one PC to the Cisco IP phone, you must configure enough secure addresses to allow one for each PC and one for the phone.
- 
                                       					
                                       When a trunk port configured with port security and assigned to an access VLAN for data traffic and to a voice VLAN for voice traffic, entering the switchport voice and switchport priority extend interface configuration commands has no effect. When a connected device uses the same MAC address to request an IP address for the access VLAN and then an IP address for the voice VLAN, only the access VLAN is assigned an IP address.
- 
                                       					
                                       When you enter a maximum secure address value for an interface, and the new value is greater than the previous value, the new value overwrites the previously configured value. If the new value is less than the previous value and the number of configured secure addresses on the interface exceeds the new value, the command is rejected.
- 
                                       					
                                       The switch does not support port security aging of sticky secure MAC addresses.
This table summarizes port security compatibility with other port-based features.
| Table 5. Port Security Compatibility                                           		  with Other Switch Features |  | 
|---|---|
| Type of Port or Feature on Port | Compatible with Port Security | 
|---|---|
| DTP 4 port 5 | No | 
| Trunk port | Yes | 
| Dynamic-access port 6 | No | 
| Routed port | No | 
| SPAN source port | Yes | 
| SPAN destination port | No | 
| EtherChannel | No | 
| Tunneling port | Yes | 
| Protected port | Yes | 
| IEEE 802.1x port | Yes | 
| Voice VLAN port 7 | Yes | 
| IP source guard | Yes | 
| Dynamic Address Resolution Protocol (ARP) inspection | Yes | 
| Flex Links | Yes | 
Overview of Port-Based Traffic Control
Port-based traffic control is a set of Layer 2 features on the Cisco Catalyst switches used to filter or block packets at the port level in response to specific traffic conditions. The following port-based traffic control features are supported:
- 
                                    				
                                    Storm Control
- 
                                    		  
                                    Protected Ports
- 
                                    		  
