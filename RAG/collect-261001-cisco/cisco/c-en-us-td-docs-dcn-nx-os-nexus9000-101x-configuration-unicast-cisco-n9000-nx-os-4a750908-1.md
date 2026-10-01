---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908-1
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908.md
source_anchor: ""
source_lines: [1, 57]
sha256: d770ad0714e637aa4905f794105c424250312246113c455b2bb0bb7f10b291e9
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908

About Basic BGP
Cisco NX-OS supports BGP version 4, which includes multiprotocol extensions that allow BGP to carry routing information for IP multicast routes and multiple Layer 3 protocol address families. BGP uses TCP as a reliable transport protocol to create TCP sessions with other BGP-enabled devices.
BGP uses a path-vector routing algorithm to exchange routing information between BGP-enabled networking devices or BGP speakers. Based on this information, each BGP speaker determines a path to reach a particular destination while detecting and avoiding paths with routing loops. The routing information includes the actual route prefix for a destination, the path of autonomous systems to the destination, and other path attributes.
BGP selects a single path, by default, as the best path to a destination host or network. Each path carries well-known mandatory, well-known discretionary, and optional transitive attributes that are used in BGP best-path analysis. You can influence BGP path selection by altering some of these attributes by configuring BGP policies. See the Route policies and resetting BGP sessions section for more information.
BGP also supports load balancing or equal-cost multipath (ECMP). See the Load Sharing and Multipath section for more information.
BGP autonomous systems
An autonomous system (AS) is a network controlled by a single administration entity. An autonomous system forms a routing domain with one or more interior gateway protocols (IGPs) and a consistent set of routing policies.
BGP supports 16-bit and 32-bit autonomous system numbers. For more information, see the Autonomous Systems section.
Separate BGP autonomous systems dynamically exchange routing information through external BGP (eBGP) peering sessions. BGP speakers within the same autonomous system can exchange routing information through internal BGP (iBGP) peering sessions.
4-Byte AS number support
BGP supports 2-byte autonomous system (AS) numbers in plain-text notation or as.dot notation and 4-byte AS numbers in plain-text notation.
When BGP is configured with a 4-byte AS number, the route-target auto VXLAN command cannot be used because the AS number along with the VNI (which is already a 3-byte value) is used to generate the route target. For more information, see the Cisco Nexus 9000 Series NX-OS VXLAN Configuration Guide.
Administrative distance
An administrative distance is a rating of the trustworthiness of a routing information source. By default, BGP uses the administrative distances shown in the table.
| Table 1. BGP Default                                     				Administrative Distances |  |  | 
|---|---|---|
| Distance | Default Value | Function | 
|---|---|---|
| External | 20 | Applied to routes learned from eBGP. | 
| Internal | 200 | Applied to routes learned from iBGP. | 
| Local | 220 | Applied to routes originated by the router. | 
| Note | The administrative distance does not influence the BGP path selection algorithm, but it does influence whether BGP-learned routes are installed in the IP routing table. | 
For more information, see the Administrative Distance section.
BGP peers
A BGP peer is a BGP speaker that has an active TCP connection to another BGP speaker.
A BGP speaker does not discover another BGP speaker automatically. You must configure the relationships between BGP speakers.
BGP sessions
BGP sessions are TCP connections established between Border Gateway Protocol (BGP) peers that enable the exchange of routing information. BGP uses TCP port 179 to create these sessions, which serve as the communication channel for exchanging routing updates and maintaining network topology awareness.
After this initial exchange, the BGP peers send only incremental updates when a topology change occurs in the network or when a routing policy change occurs. In the periods of inactivity between these updates, peers exchange special messages called keepalives. The hold time is the maximum time limit that can elapse between receiving consecutive BGP update or keepalive messages.
Cisco NX-OS supports these peer configuration options:
- 
                                    				
                                    Individual IPv4 or IPv6 address: BGP establishes a session with the BGP speaker that matches the remote address and AS number.
- 
                                    				
                                    IPv4 or IPv6 prefix peers for a single AS number: BGP establishes sessions with BGP speakers that match the prefix and the AS number.
- 
                                    				
                                    Dynamic AS number prefix peers: BGP establishes sessions with BGP speakers that match the prefix and an AS number from a list of configured AS numbers.
Dynamic AS numbers for prefix peers and interface peers
Dynamic AS numbers are a feature in Cisco NX-OS BGP configurations that allow you to specify a range or list of Autonomous System (AS) numbers for establishing BGP sessions with prefix peers or interface peers. This means BGP will accept sessions only from peers whose AS numbers match those in the configured list or range.
Cisco NX-OS accepts a range or list of AS numbers to establish BGP sessions. For example, if you configure BGP to use IPv4 prefix 192.0.2.0/8 and AS numbers 33, 66, and 99, BGP establishes a session with 192.0.2.1 with AS number 66 but rejects a session from 192.0.2.2 with AS number 50.
Beginning with Cisco NX-OS Release 9.3(6), support for dynamic AS numbers is extended to interface peers in addition to prefix peers. See Configure BGP Interface Peering via IPv6 Link-Local for IPv4 and IPv6 Address Families.
Cisco NX-OS does not associate prefix peers with dynamic AS numbers as either interior BGP (iBGP) or external BGP (eBGP) sessions until after the session is established. See Configuring Advanced BGP for more information on iBGP and eBGP.
| Note | The dynamic AS number prefix peer configuration overrides the individual AS number configuration that is inherited from a BGP template. For more information, see Configuring Advanced BGP. | 
BGP router identifier
To establish BGP sessions between peers, BGP must have a router ID, which is sent to BGP peers in the OPEN message when a BGP session is established. The BGP router ID is a 32-bit value that is often represented by an IPv4 address. You can configure the router ID. By default, Cisco NX-OS sets the router ID to the IPv4 address of a loopback interface on the router. If no loopback interface is configured on the router, the software chooses the highest IPv4 address configured to a physical interface on the router to represent the BGP router ID. The BGP router ID must be unique to the BGP peers in a network.
If BGP does not have a router ID, it cannot establish any peering sessions with BGP peers.
Each routing process has an associated router ID. You can configure the router ID to any interface in the system. If you do not configure the router ID, Cisco NX-OS selects the router ID based on the following criteria:
- 
                                 				
                                 Cisco NX-OS prefers loopback0 over any other interface. If loopback0 does not exist, then Cisco NX-OS prefers the first loopback interface over any other interface type.
- 
                                 				
                                 If you have not configured a loopback interface, Cisco NX-OS uses the first interface in the configuration file as the router ID. If you configure any loopback interface after Cisco NX-OS selects the router ID, the loopback interface becomes the router ID. If the loopback interface is not loopback0 and you configure loopback0 with an IP address, the router ID changes to the IP address of loopback0.
- 
                                 				
