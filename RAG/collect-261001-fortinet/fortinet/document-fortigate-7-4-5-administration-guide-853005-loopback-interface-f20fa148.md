---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-5-administration-guide-853005-loopback-interface-f20fa148
title: "document-fortigate-7-4-5-administration-guide-853005-loopback-interface-f20fa148"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-5-administration-guide-853005-loopback-interface-f20fa148.md
source_anchor: ""
source_lines: [1, 29]
sha256: 44ea77c6e511f28583c6cd8f73c72ac71af657101fd8e0e7574830aca879ac80
---

# document-fortigate-7-4-5-administration-guide-853005-loopback-interface-f20fa148

Loopback interface
Loopback interface
A loopback interface is a logical interface that is always up. Its IP address does not depend on one specific physical port, and the attached subnet is always present in the routing table. Therefore, it can be accessed through several physical or VLAN interfaces.
Typically, a loopback interface can be used with management access, BGP peering, PIM rendezvous points, and SD-WAN.
A loopback interface requires appropriate firewall policies to allow traffic to the interface. For example, see IPsec tunnel terminated on a loopback interface.
Multiple loopback interfaces can be configured in either non-VDOM mode or in each VDOM.
Dynamic routing protocols can be enabled on loopback interfaces. For example, loopback interfaces are a good practice for OSPF. To make it easier to troubleshoot OSPF, set the OSPF router ID to the same value as the loopback IP address to access a specific FortiGate using that IP address and SSH.
A loopback interface is configured using similar steps as a physical interface (see Configuring an interface).
IPsec tunnel terminated on a loopback interface
As mentioned above, a loopback interface requires appropriate firewall policies to allow traffic to the interface. In other words, traffic ingressing on an interface that is destined for the IP address associated with a loopback interface requires an appropriate firewall policy from that interface to the loopback interface otherwise the traffic will be dropped.
For example, consider the following topology where an IPsec tunnel is terminated on a loopback interface, VPN_LO, on the FortiGate FGT-1 and on a WAN interface on the FortiGate FGT-2.
We will focus on the configuration required for FortiGate FGT-1.
IPsec tunnel terminates on a loopback interface, VPN_LO, which has an associated IP address that the remote peer will use as its IPsec remote gateway address.
The IPsec tunnel uses wan1 as its underlay interface.
In this scenario, the administrator of the FortiGate FGT-1 device must configure a firewall policy from the wan1 interface to the VPN_LO interface that allows incoming traffic from the remote peer to reach the VPN_LO interface for proper IPsec tunnel connectivity.
For example:
config firewall policy
    edit 4
        set name "Loopback-In"
        set srcintf "wan1"
        set dstintf "VPN_LO"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set logtraffic disable
    next
end
