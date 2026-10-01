---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke-e0593c3a-1
title: "c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2014-03-28"]
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a.md
source_anchor: ""
source_lines: [1, 46]
sha256: 5b518e90711fa345e69aeb757de040951ed0570adca0aeb94ab35f1bbcafd060
---

# c-en-us-td-docs-routers-ios-config-17-x-sec-vpn-b-security-vpn-m-sec-flex-spoke--e0593c3a

Prerequisites for FlexVPN Spoke to Spoke
IKEv2, the FlexVPN server, and the FlexVPN spoke must be configured.
Last Published Date: March 28, 2014
The FlexVPN Spoke to Spoke feature enables a FlexVPN client to establish a direct crypto tunnel with another FlexVPN client leveraging virtual tunnel interfaces (VTI), Internet Key Exchange Version 2 (IKEv2), and Next Hop Resolution Protocol (NHRP) to build spoke-to-spoke connections.
FlexVPN is Cisco’s implementation of the IKEv2 standard featuring a unified paradigm and CLI that combines site to site, remote access, hub and spoke topologies and partial meshes (spoke to spoke direct). FlexVPN offers a simple but modular framework that extensively uses the tunnel interface paradigm while remaining compatible with legacy VPN implementations using the crypto maps.
The FlexVPN server provides the server side functionality of FlexVPN. The FlexVPN client establishes a secure IPsec VPN tunnel between a FlexVPN client and another FlexVPN server.
NHRP is an Address Resolution Protocol (ARP)-like protocol that alleviates nonbroadcast multiaccess (NBMA) network problems. With NHRP, NHRP entities attached to an NBMA network dynamically learn the NBMA address of the other entities that are part of that network, allowing these entities to directly communicate without requiring traffic to use an intermediate hop.
The FlexVPN Spoke to Spoke feature integrates NHRP and FlexVPN client (spoke) to establish a direct crypto channel with another client in an existing FlexVPN network. The connections are built using virtual tunnel interfaces (VTI), IKEv2 and NHRP, where NHRP is used for resolving the FlexVPN clients in the network.
The following is recommended in FlexVPN:
Routing entries are not exchanged between spokes.
Different profiles are used for the spokes and the config-exchange command is not configured for the spokes.
The FlexVPN IPv6 Direct Spoke to Spoke feature supports the use of IPv6 addresses for FlexVPN spokes. The support for IPv6 addresses provides support for IPv6 over IPv4, IPv4 over IPv6, and IPv6 over IPv6 transports.
| Note | Spoke to Spoke FlexVPN does not support dynamic AAA authorization. | 
The following diagram illustrates the NHRP resolution request and reply in FlexVPN.
Due to bidirectional traffic, similar events occur in both directions at Spoke1, Spoke2, and hub. For clarity, events from Host1 to Host2 are discussed. Assume that there is a network N1 (192.168.1.0/24) behind Spoke1 and another network N2 (192.168.2.0/24) behind Spoke2. The network between the two spokes is matched through an access control list (ACL). This is because ACLs are applied on the IKEv2 policies on both spokes.
The network along with its prefix information from both the spokes is conveyed to the hub via IKEv2 information payload exchanges. This causes a route addition in the routing table by IKEv2 at the hub as follows:
192.168.1.0/24—Connected via virtual access interface1
192.168.2.0/24—Connected via virtual access interface2
The hub will push a summarized route via IKEv2 to both spokes, and the spokes will install the route in their routing table as follows:
192.168.0.0/16—next hop <tunnel address of the hub> - interface Tunnel 1
| Note | The routing protocol can also add the route to the routing table. | 
Assuming that traffic moves from N1 to N2, the traffic flow is as follows:
Host1 sends traffic destined to Host2. The traffic reaches the LAN interface of spoke1, looks up the route, hits the summarized route, and routes the packet to interface tunnel 1.
When the traffic reaches the hub’s virtual access interface1, the traffic looks up the route table for a route entry for N2, either directly connected over virtual access interface 2 or via a point-to-point tunnel interface.
The traffic from Host1 to Host2 traverses the hub through virtual access interface1 and virtual access interface2. The hub determines that ingress and the egress interfaces (virtual access interface1 and virtual access interface2) belong to same NHRP network (network D configured on both the interfaces). The hub sends out an NHRP redirect message to spoke1 on virtual access interface1.
On receiving the redirect, Spoke1 initiates a resolution request for Host2 over the point-to-point tunnel interface (the same interface over which it received the redirect). The resolution request traverses the routed path (Spoke1-hub-spoke2). On receiving the resolution request, Spoke2 determines that it is the exit point and needs to respond to the resolution request.
Spoke2 receives the resolution request on the tunnel interface and retrieves the virtual template number from the tunnel interface. The virtual template number is used to create the virtual access interface to start a crypto channel and establishes IKEv2 and IPsec security associations (SAs). Once the crypto SAs between the two spokes are up, Spoke2 installs the necessary NHRP cache entries for Spoke1 and its network under the newly created virtual access interface and sends out the resolution reply over the virtual access interface.
After receiving the resolution request over the virtual access interface, Spoke1 installs the necessary cache entries for Spoke2 and its network. Spoke1 also deletes the temporary cache entry pointing to the hub to resolve the network under tunnel interface1.
NHRP adds shortcut routes as next-hop override (NHO) or H route. For more information on shortcut switching, refer to Shortcut Switching Enhancements for NHRP in DMVPN Networks.
The FlexVPN server and client must be configured.
| Note | If you are using EIGRP as the routing protocol it is important to configure route filtering on the Hub to ensure that the Spoke tunnel addresses are not advertised via EIGRP. This is important to ensure that there is symmetric routing between the hub and spoke interfaces even after spoke- spoke tunnel is established. | 
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface virtual-template number type tunnel Example: Device(config)# interface virtual-template 1 type tunnel | Creates a virtual template interface that can be configured and applied dynamically to create virtual access interfaces. | 
| Step 4 | ip unnumbered loopback number Example: Device(config-if)# ip unnumbered loopback 0 | Assigns the IP address of an existing interface (usually a loopback interface) to the virtual tunnel interface. | 
| Step 5 | Do one of the following:  Example: Device(config-if)# ip nhrp network-id 1Example: Device(config-if)# ipv6 nhrp network-id 1 | Enables NHRP on the interface. | 
| Step 6 | ip nhrp redirect [timeout seconds] Example: Device(config-if)# ip nhrp redirect | Enables redirect traffic indication if traffic is forwarded with the NHRP network. To avoid sending duplicate redirects, use the timeout keyword and the seconds argument to indicate when to expire a redirect entry created. | 
| Step 7 | exit Example: Device(config-if)# exit | Exits interface configuration mode and returns to global configuration mode. | 
Perform this task to configure NHRP shortcuts on the tunnel interface on the FlexVPN spoke.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | interface tunnel number Example: Device(config)# interface tunnel 1 | Configures the FlexVPN client interface and enters interface configuration mode. | 
