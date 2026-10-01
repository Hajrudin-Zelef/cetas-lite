---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2-2
title: "Run the show running-config command to check the interface configuration."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2.md
source_anchor: ""
source_lines: [48, 99]
sha256: fbf804dcc7a562b9fa217e49070fae772faf81d1ea219c17a05fed4a515fb9bb
---

# Run the show running-config command to check the interface configuration.

VTP supports three working modes: VTP server, VTP client, and VTP transparent. For details, see Table 1-3.
Table 1-3 Working mode of VTP
| Working Mode | Description | 
| Server | A VTP server maintains all VLAN lists in the local VTP domain. It can create, delete, and modify VLANs, send advertisement packets, and synchronize VLAN information to other switches in the local VTP domain. VLAN information is saved in the nonvolatitle RAM (NVRAM). By default, a Cisco switch is used as the VTP server. | 
| Client | A VTP client learns VTP information from a VTP server. It cannot create, delete, or modify VLANs, but can forward advertisement packets. VLAN information is not saved in NVRAM. | 
| Transparent | A VTP transparent switch is an independent switch that does not participate in VTP implementation or learn VLAN information from the VTP server. It only maintains local VLAN information. The VTP transparent switch can create, delete, and modify only the local VLAN information. When VTP version 1 is used, VTP transparent switches can only forward VTP packets of other switches in the same VTP domain. When VTP version 2 is used, VTP transparent switches can forward VTP packets of switches in a different VTP domain. | 
Advertisement Packets
Switches use VTP advertisement packets to transmit VLAN information. Table 1-4 describes three types of VTP advertisement packets.
Table 1-4 Format of advertisement packets
| Format | Applicable Scenario | 
| Summary Advertisement | l By default, a VTP server sends a Summary Advertisement packet every 300s to inform adjacent switches of the current VTP domain name and the configuration revision number. When a switch receives a summary advertisement packet, the following situations occur: 1. The switch compares the VTP domain name with its own VTP domain name. If the names are different, the switch ignores the packet. 2. If the names are the same, the switch compares the configuration revision number with its configuration revision number. 3. If its configuration revision number is higher than or equal to the configuration revision number in the Summary Advertisement packet, the switch ignores the packet. If its configuration revision number is lower than the configuration revision number in the Summary Advertisement packet, the switch sends an Advertisement Request packet. l When a switch receives an Advertisement Request packet, it sends a Summary Advertisement packet, and then sends one or several Subset Advertisement packets. | 
| Subset Advertisement | When you add, delete, or change a VLAN on a switch, the VTP server where the changes are made increments the configuration revision number and sends a Summary Advertisement packet. Then the VTP server sends one or more Subset Advertisement packets. A subset advertisement contains a list of VLAN information. If there are several VLANs, the VTP server needs to send more than one Subset Advertisement packet to advertise all the VLANs. | 
| Advertisement Request | A switch needs an Advertisement Request packet in the following situations: l The switch restarts. l The VTP domain name has been changed. l The switch has received a Summary Advertisement packet with a higher configuration revision number than its own. | 
VTP advertisement packets have the following characteristics:
l VTP advertisement packets are transmitted in multicast mode through trunk interfaces in VLAN 1.
l VTP advertisement packets are sent to the destination MAC address 01-00-0C-CC-CC-CC.
l VTP advertisement packets are sent in either Inter-Switch Link (ISL) or IEEE 802.1Q (dot1q) frames.
The differences between Cisco VTP and Huawei VCMP are as follows.
l Multiple servers can exist in a Cisco VTP domain, and any switch can function as a VTP server. VTP servers synchronize information to each other.
l Only one switch in a VCMP domain functions as the VCMP server to control all VLAN configurations in the domain.
Table 1-5 Differences in command formats
| Function | Command on Huawei S Series Switches | Command on Cisco Switches | Description | 
| Configure the device role or mode. | vcmp role { client \| server \| silent \| transparent } | vtp mode { client \| off \| server \| transparent } | A switch used as a VCMP silent in a Huawei VCMP domain is similar to the switch in off mode in a Cisco VTP domain, and directly discards received protocol packets. | 
| Configure the domain name. | vcmp domain domain-name | vtp domain domain-name vtp domain domain-name | - | 
| Configure the domain ID. | vcmp device-id device-id | Not supported | Cisco VTP does not support the configuration. | 
| Configure an authentication password for the domain. | vcmp authentication sha2-256 password password | vtp password password | - | 
| Configuring the protocol version number. | Not supported | vtp version number | Huawei VCMP does not support the configuration. | 
| Check the protocol status. | display vcmp status | show vtp status | - | 
VTP and VCMP are proprietary protocols, and cannot interwork. Huawei S series switches and Cisco switches can be used on the entire network. Configurations can be performed on the switch that is directly connected to Huawei and Cisco switches to implement interworking between Huawei and Cisco switches. The following describes three types of hybrid networking models.
l 1.5.1 Hybrid Networking 1: C-H Model
In the C-H model, a Cisco switch directly connects to a Huawei S series switch that has no downstream Cisco switch connected.
l 1.5.2 Hybrid Networking 2: C-H-C Model
In the C-H-C model, a Cisco switch directly connects to a Huawei S series switch that has a downstream Cisco switch connected.
l 1.5.3 Hybrid Networking 3: C-H-H-C Model
In the C-H-H-C model, a Cisco switch directly connects to a Huawei S series switch, and another edge switch of the VCMP network connects to a Cisco switch.
Huawei S series switches can replace switches in a Cisco VTP domain.
l Replacing the transparent switch
In Figure 1-3, a Huawei S series switch replaces the VTP transparent switch on a Cisco network. After the replacement, you only need to create a VLAN manually on the Huawei S series switch and add interfaces to the VLAN. For details, see Huawei S series switch configuration in 1.5.1 Hybrid Networking 1: C-H Model.
Figure 1-3 Networking for replacing the transparent switch
l Replacing the client
In Figure 1-4, a Huawei S series switch replaces the VTP client on a Cisco network. After the replacement, you need to configure the Huawei S series switch to transparently transmit VTP packets. For details, see Huawei S series switch configuration in 1.5.2 Hybrid Networking 2: C-H-C Model.
Figure 1-4 Networking for replacing the VTP client
l Replacing the server
In Figure 1-5, no VTP server exists in the VTP domain after a Huawei S series switch replace it. If the Cisco network runs VTP version 1 or 2, any Cisco switch can function as the VTP server. If the Cisco network runs VTP version 3 alone or with VTP version1, you need to find a switch running VTP version 3 and run the vtp primary vlan command to specify the switch as the VTP server to manage the VTP domain.
The Huawei S series switch only needs to transparently transmit VTP packets. For details, see Huawei S series switch configuration in 1.5.2 Hybrid Networking 2: C-H-C Model.
Figure 1-5 Networking for replacing the VTP server
Overview
When VTP is enabled on the Cisco switch to synchronize VLAN information, the Huawei S series switch cannot process VTP packets. Therefore, a VLAN needs to be configured manually on the Huawei S series switch.
Configuration Notes
l This example applies to Huawei S series switches of all versions.
l If switchport dynamic auto or switchport dynamic desirable is configured on the Cisco switch interface before the Cisco switch interface is directly connects to the Huawei S series switch, change it to switchport mode trunk to prevent DTP negotiation failure.
Networking Requirements
