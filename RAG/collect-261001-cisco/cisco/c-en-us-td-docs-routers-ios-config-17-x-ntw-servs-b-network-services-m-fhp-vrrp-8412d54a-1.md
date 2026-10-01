---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp-8412d54a-1
title: "c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory", "preemption"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a.md
source_anchor: ""
source_lines: [1, 98]
sha256: 987e8c92e7fc6dc58cf27ebf7d1c0cd37c17b3b75f756903ed6a20a8ed099d8e
---

# c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
The Virtual Router Redundancy Protocol (VRRP) is an election protocol that dynamically assigns responsibility for one or
more virtual routers to the VRRP routers on a LAN, allowing several routers on a multiaccess link to utilize the same virtual
IP address. A VRRP router is configured to run the VRRP protocol in conjunction with one or more other routers attached to
a LAN. In a VRRP configuration, one router is elected as the virtual primary router, with the other routers acting as backups
in case the virtual primary router fails.
This module explains the concepts related to VRRP and describes how to configure VRRP in a network.
Restrictions for VRRP
VRRP is designed for use over multiaccess, multicast, or broadcast capable Ethernet LANs. VRRP is not intended as a replacement
for existing dynamic protocols.
VRRP is supported on Ethernet, Fast Ethernet, Bridge Group Virtual Interface (BVI), and Gigabit Ethernet interfaces, and
on Multiprotocol Label Switching (MPLS) Virtual Private Networks (VPNs), VRF-aware MPLS VPNs, and VLANs.
Because of the forwarding delay that is associated with the initialization of a BVI interface, you must configure the VRRP
advertise timer to a value equal to or greater than the forwarding delay on the BVI interface. This setting prevents a VRRP
router on a recently initialized BVI interface from unconditionally taking over the primary role. Use the bridgeforward-time command to set the forwarding delay on the BVI interface. Use the vrrptimersadvertise command to set the VRRP advertisement timer.
Information About VRRP
VRRP Operation
There are several ways a LAN client can determine which router should be the first hop to a particular remote destination.
The client can use a dynamic process or static configuration. Examples of dynamic router discovery are as follows:
Proxy ARP—The client uses Address Resolution Protocol (ARP) to get the destination it wants to reach, and a router will respond
to the ARP request with its own MAC address.
Routing protocol—The client listens to dynamic routing protocol updates (for example, from Routing Information Protocol [RIP])
and forms its own routing table.
ICMP Router Discovery Protocol (IRDP) client—The client runs an Internet Control Message Protocol (ICMP) router discovery
client.
The drawback to dynamic discovery protocols is that they incur some configuration and processing overhead on the LAN client.
Also, in the event of a router failure, the process of switching to another router can be slow.
An alternative to dynamic discovery protocols is to statically configure a default router on the client. This approach simplifies
client configuration and processing, but creates a single point of failure. If the default gateway fails, the LAN client is
limited to communicating only on the local IP network segment and is cut off from the rest of the network.
VRRP can solve the static configuration problem. VRRP enables a group of routers to form a single
virtualrouter. The LAN clients can then be configured with the virtual router as their default gateway. The virtual router, representing
a group of routers, is also known as a VRRP group.
VRRP is supported on Ethernet, Fast Ethernet, BVI, and Gigabit Ethernet interfaces, and on MPLS VPNs, VRF-aware MPLS VPNs,
and VLANs.
The figure below shows a LAN topology in which VRRP is configured. In this example, Routers A, B, and C are VRRP routers
(routers running VRRP) that comprise a virtual router. The IP address of the virtual router is the same as that configured
for the Ethernet interface of Router A (10.0.0.1).
Because the virtual router uses the IP address of the physical Ethernet interface of Router A, Router A assumes the role
of the virtual primary router and is also known as the IP address owner. As the virtual primary router, Router A controls
the IP address of the virtual router and is responsible for forwarding packets sent to this IP address. Clients 1 through
3 are configured with the default gateway IP address of 10.0.0.1.
Routers B and C function as virtual router backups. If the virtual primary router fails, the router configured with the higher
priority will become the virtual primary router and provide uninterrupted service for the LAN hosts. When Router A recovers,
it becomes the virtual primary router again. For more detail on the roles that VRRP routers play and what happens if the virtual
primary router fails, see the VRRP Router Priority and Preemption section.
The figure below shows a LAN topology in which VRRP is configured so that Routers A and B share the traffic to and from clients
1 through 4 and that Routers A and B act as virtual router backups to each other if either router fails.
In this topology, two virtual routers are configured. (For more information, see the Multiple Virtual Router Support section.) For virtual router 1, Router A is the owner of IP address 10.0.0.1 and virtual primary router, and Router B is
the virtual router backup to Router A. Clients 1 and 2 are configured with the default gateway IP address of 10.0.0.1.
For virtual router 2, Router B is the owner of IP address 10.0.0.2 and virtual primary router, and Router A is the virtual
router backup to Router B. Clients 3 and 4 are configured with the default gateway IP address of 10.0.0.2.
VRRP Benefits
Redundancy
VRRP enables you to
configure multiple routers as the default gateway router, which reduces the
possibility of a single point of failure in a network.
Load Sharing
You can configure
VRRP in such a way that traffic to and from LAN clients can be shared by
multiple routers, thereby sharing the traffic load more equitably among
available routers.
Multiple Virtual
Routers
Multiple IP Addresses
The virtual router
can manage multiple IP addresses, including secondary IP addresses. Therefore,
if you have multiple subnets configured on an Ethernet interface, you can
configure VRRP on each subnet.
Preemption
The redundancy scheme of VRRP enables you to preempt a virtual router backup that has taken over for a failing virtual primary
router with a higher priority virtual router backup that has become available.
Authentication
VRRP message digest
5 (MD5) algorithm authentication protects against VRRP-spoofing software and
uses the industry-standard MD5 algorithm for improved reliability and security.
Advertisement
Protocol
VRRP uses a
dedicated Internet Assigned Numbers Authority (IANA) standard multicast address
(224.0.0.18) for VRRP advertisements. This addressing scheme minimizes the
number of routers that must service the multicasts and allows test equipment to
accurately identify VRRP packets on a segment. The IANA assigned VRRP the IP
protocol number 112.
VRRP Object Tracking
VRRP object tracking provides a way to ensure the best VRRP router is the virtual primary router for the group by altering
VRRP priorities to the status of tracked objects such as the interface or IP route states.
Multiple Virtual Router
Support
Router processing
capability
Router memory
capability
Router interface
support of multiple MAC addresses
In a topology where multiple virtual routers are configured on a router interface, the interface can act as primary for one
virtual router and as a backup for one or more virtual routers.
VRRP Router Priority and Preemption
