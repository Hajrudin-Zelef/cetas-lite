---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98.md
source_anchor: ""
source_lines: [128, 151]
sha256: 9872156516ecbd088405bfa484da56da0fd9bc2d4b1e44d7eb898fff002c2a55
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-12-configuration-g-b39e1a98

                                 Static address—A manually entered unicast address that does not age and that is not lost when the device resets.
The address table lists the destination MAC address, the associated VLAN ID, and port number associated with the address and the type (static or dynamic).
| Note | For complete syntax and usage information for the commands used in this section, see the command reference for this release. | 
MAC Address Table Creation
With multiple MAC addresses supported on all ports, you can connect any port on the device to other network devices. The device provides dynamic addressing by learning the source address of packets it receives on each port and adding the address and its associated port number to the address table. As devices are added or removed from the network, the device updates the address table, adding new dynamic addresses and aging out those that are not in use.
The aging interval is globally configured. However, the device maintains an address table for each VLAN, and STP can accelerate the aging interval on a per-VLAN basis.
The device sends packets between any combination of ports, based on the destination address of the received packet. Using the MAC address table, the device forwards the packet only to the port associated with the destination address. If the destination address is on the port that sent the packet, the packet is filtered and not forwarded. The device always uses the store-and-forward method: complete packets are stored and checked for errors before transmission.
MAC Addresses and VLANs
All addresses are associated with a VLAN. An address can exist in more than one VLAN and have different destinations in each. Unicast addresses, for example, could be forwarded to port 1 in VLAN 1 and ports 9, 10, and 1 in VLAN 5.
Each VLAN maintains its own logical address table. A known address in one VLAN is unknown in another until it is learned or statically associated with a port in the other VLAN.
Default MAC Address Table Settings
The following table shows the default settings for the MAC address table.
| Table 2.  Default Settings for the MAC Address |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Aging time | 300 seconds | 
| Dynamic addresses | Automatically learned | 
| Static addresses | None configured | 
ARP Table Management
To communicate with a device (over Ethernet, for example), the software first must learn the 48-bit MAC address or the local data link address of that device. The process of learning the local data link address from an IP address is called address resolution .
The Address Resolution Protocol (ARP) associates a host IP address with the corresponding media or MAC addresses and the VLAN ID. Using an IP address, ARP finds the associated MAC address. When a MAC address is found, the IP-MAC address association is stored in an ARP cache for rapid retrieval. Then the IP datagram is encapsulated in a link-layer frame and sent over the network. Encapsulation of IP datagrams and ARP requests and replies on IEEE 802 networks other than Ethernet is specified by the Subnetwork Access Protocol (SNAP). By default, standard Ethernet-style ARP encapsulation (represented by the arpa keyword) is enabled on the IP interface.
ARP entries added manually to the table do not age and must be manually removed.
For CLI procedures, see the Cisco IOS Release 12.4 documentation on Cisco.com.
