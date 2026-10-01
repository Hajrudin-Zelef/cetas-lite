---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-3
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [97, 172]
sha256: 64115860a692921926ab69faac551bbdea4567f09f964a3ca5edf28517386faf
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

By default, both UDP and ND forwarding are enabled if a helper address has been defined for an interface. The description for the ip forward-protocol interface configuration command in the Cisco IOS IP Command Reference, Volume 1 of 3: Addressing and Services lists the ports that are forwarded by default if you do not specify any UDP ports.
Broadcast Packet Handling
After configuring an IP interface address, you can enable routing and configure one or more routing protocols, or you can configure the way the Device responds to network broadcasts. A broadcast is a data packet destined for all hosts on a physical network. The Device supports two kinds of broadcasting:
- 
		  A directed broadcast packet is sent to a specific network or series of networks. A directed broadcast address includes the network or subnet fields.
- 
		  A flooded broadcast packet is sent to every network. 
 Note
 You can also limit broadcast, unicast, and multicast traffic on Layer 2 interfaces by using the storm-control interface configuration command to set traffic suppression levels.
Routers provide some protection from broadcast storms by limiting their extent to the local cable. Bridges (including intelligent bridges), because they are Layer 2 devices, forward broadcasts to all network segments, thus propagating broadcast storms. The best solution to the broadcast storm problem is to use a single broadcast address scheme on a network. In most modern IP implementations, you can set the address to be used as the broadcast address. Many implementations, including the one in the Device, support several addressing schemes for forwarding broadcast messages.
IP Broadcast Flooding
You can allow IP broadcasts to be flooded throughout your internetwork in a controlled fashion by using the database created by the bridging STP. Using this feature also prevents loops. To support this capability, bridging must be configured on each interface that is to participate in the flooding. If bridging is not configured on an interface, it still can receive broadcasts. However, the interface never forwards broadcasts it receives, and the router never uses that interface to send broadcasts received on a different interface.
Packets that are forwarded to a single network address using the IP helper-address mechanism can be flooded. Only one copy of the packet is sent on each network segment.
To be considered for flooding, packets must meet these criteria. (Note that these are the same conditions used to consider packet forwarding using IP helper addresses.)
-  
		  The packet must be a MAC-level broadcast.
-  
		  The packet must be an IP-level broadcast.
-  
		  The packet must be a TFTP, DNS, Time, NetBIOS, ND, or BOOTP packet, or a UDP specified by the ip forward-protocol udp global configuration command.
-  
		  The time-to-live (TTL) value of the packet must be at least two.
A flooded UDP datagram is given the destination address specified with the ip broadcast-address interface configuration command on the output interface. The destination address can be set to any address. Thus, the destination address might change as the datagram propagates through the network. The source address is never changed. The TTL value is decremented.
When a flooded UDP datagram is sent out an interface (and the destination address possibly changed), the datagram is handed to the normal IP output routines and is, therefore, subject to access lists, if they are present on the output interface.
In the Device, the majority of packets are forwarded in hardware; most packets do not go through the Device CPU. For those packets that do go to the CPU, you can speed up spanning tree-based UDP flooding by a factor of about four to five times by using turbo-flooding. This feature is supported over Ethernet interfaces configured for ARP encapsulation.
How to Configure IP Routing
By default, IP routing is disabled on the Device, and you must enable it before routing can take place. For detailed IP routing configuration information, see the Cisco IOS IP Configuration Guide.
In the following procedures, the specified interface must be one of these Layer 3 interfaces:
-  
		  A routed port: a physical port configured as a Layer 3 port by using the no switchport interface configuration command.
-  
		  A switch virtual interface (SVI): a VLAN interface created by using the interface vlan vlan_id global configuration command and by default a Layer 3 interface. 
 Note
 On enabling ip routing, the VLAN configured as SVI will also learn broadcast ARP requests which are not self destined.
-  
		  An EtherChannel port channel in Layer 3 mode: a port-channel logical interface created by using the interface port-channel port-channel-number global configuration command and binding the Ethernet interface into the channel group. For more information, see the “Configuring Layer 3 EtherChannels” chapter in the Layer 2 Configuration Guide. 
 Note
 The switch does not support tunnel interfaces for unicast routed traffic.
All Layer 3 interfaces on which routing will occur must have IP addresses assigned to them.
| Note | A Layer 3 switch can have an IP address assigned to each routed port and SVI. The number of routed ports and SVIs that you can configure is limited to 128, exceeding the recommended number and volume of features being implemented might impact CPU utilization because of hardware limitations. | 
Configuring routing consists of several main procedures:
-  
		  To support VLAN interfaces, create and configure VLANs on the Device or switch stack, and assign VLAN membership to Layer 2 interfaces. For more information, see the "Configuring VLANs” chapter in the VLAN Configuration Guide.
-  
		  Configure Layer 3 interfaces.
-  
		  Enable IP routing on the switch.
-  
		  Assign IP addresses to the Layer 3 interfaces.
-  
		  Enable selected routing protocols on the switch.
-  
		  Configure routing protocol parameters (optional).
How to Configure IP Addressing
A required task for configuring IP routing is to assign IP addresses to Layer 3 network interfaces to enable the interfaces and allow communication with the hosts on those interfaces that use IP. The following sections describe how to configure various IP addressing features. Assigning IP addresses to the interface is required; the other procedures are optional.
Default IP Addressing Configuration
| Table 1 Default Addressing Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| IP address | None defined. | 
| ARP | No permanent entries in the Address Resolution Protocol (ARP) cache. Encapsulation: Standard Ethernet-style ARP. Timeout: 14400 seconds (4 hours). | 
| IP broadcast address | 255.255.255.255 (all ones). | 
| IP classless routing | Enabled. | 
| IP default gateway | Disabled. | 
| IP directed broadcast | Disabled (all IP directed broadcasts are dropped). | 
| IP domain | Domain list: No domain names defined. Domain lookup: Enabled. Domain name: Enabled. | 
| IP forward-protocol | If a helper address is defined or User Datagram Protocol (UDP) flooding is configured, UDP forwarding is enabled on default ports. Any-local-broadcast: Disabled. Spanning Tree Protocol (STP): Disabled. Turbo-flood: Disabled. | 
| IP helper address | Disabled. | 
| IP host | Disabled. | 
| IRDP | Disabled. Defaults when enabled: | 
| IP proxy ARP | Enabled. | 
| IP routing | Disabled. | 
| IP subnet-zero | Disabled. | 
Assigning IP Addresses to Network Interfaces
An IP address identifies a location to which IP packets can be sent. Some IP addresses are reserved for special uses and cannot be used for host, subnet, or network addresses. RFC 1166, “Internet Numbers,” contains the official description of IP addresses.
