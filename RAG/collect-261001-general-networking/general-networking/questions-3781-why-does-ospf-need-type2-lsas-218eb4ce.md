---
id: collect-261001-general-networking/general-networking/questions-3781-why-does-ospf-need-type2-lsas-218eb4ce
title: "questions-3781-why-does-ospf-need-type2-lsas-218eb4ce"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "latency", "research"]
source: docs/RAG/collect-261001-general-networking/questions-3781-why-does-ospf-need-type2-lsas-218eb4ce.md
source_anchor: ""
source_lines: [1, 71]
sha256: 62f0f2839930fdcc3fa704ff7d56f86c76dd7f138c25ba8a1e3de009a5181c67
---

# questions-3781-why-does-ospf-need-type2-lsas-218eb4ce

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
13
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Learning more about OSPF for CCNP studies. I'm looking at how OSPF builds its links, and have just covered Type1 LSAs. Looking at Type1 LSAs, I'm wondering why these are even necessary?
The book I'm reading implies that Type2 LSAs are used to help the router build the 'puzzle' of the topology, as if just using Type1 LSAs it couldn't figure out all of the links in the topology. It appears as if the Type1 LSA gives enough information for the router to be able to derive how two or even more routers are linked. Maybe the book I'm reading has poor examples, but I can't see what OSPF gains from the Type2 LSAs and it's hard to understand how they work.
It's important to note that type 2 LSA's are only generated on segments where a DR/BDR has been elected - this includes BMA (Broadcast Multi-Access) and NBMA (Non-Broadcast Multi-Access) networks. The DR is what generates the type 2 LSA. This behavior can be bypassed by configuring your Ethernet interfaces that you're choosing to run OSPF on as point-to-point (this will also prevent the DR election process).
Type 2 LSA's are beneficial when running OSPF over a Broadcast (Ethernet) or Non-Broadcast Multi-Access (Frame Relay) medium. Put simply, yes, the routers could use type 1 LSA's and detail every router's links to all other routers, but this is inefficient and will introduce unnecessary bloat into the OSPF LSDB. To mitigate this, the Type 2 (network) LSA is used to represent the broadcast subnet. Each router LSA then has a link to the broadcast subnet’s network LSA, and the network LSA has links to each of the router LSAs. It's a math problem - with every router using type 1 LSA's, you have n * (n - 1) links in the link state database. With type 2 LSA's, this number is reduced to n * 2.
I highly recommend reading John Moy's book on OSPF. He also wrote the initial RFC's for the protocol.
In addition: Type-2 LSA use only as "virtual instance" of a router in MA segment, this pseudonode has an adjacency to all attached routers (include DR/BDR) on the network and lists all attached routers (RID) to that segment. For transfer LSA they(DR/BDR) use Type-1 LSA as well.
R1# sh ip ospf database
OSPF Router with ID (1.1.1.1) (Process ID 1)
Router Link States (Area 0)
Link ID ADV Router Age Seq# Checksum Link count
1.1.1.1 1.1.1.1 708 0x80000003 0x008686 2
2.2.2.2 2.2.2.2 709 0x80000003 0x00CB0C 2
Net Link States (Area 0)
Link ID ADV Router Age Seq# Checksum
192.168.0.2 2.2.2.2 709 0x80000001 0x0014A6
R1# sh ip ospf database network
OSPF Router with ID (1.1.1.1) (Process ID 1)
Net Link States (Area 0)
Routing Bit Set on this LSA in topology Base with MTID 0
LS age: 780
Options: (No TOS-capability, DC)
LS Type: Network Links
Link State ID: 1.1.1.1 (address of Designated Router)
Advertising Router: 1.1.1.1
LS Seq Number: 80000001
Checksum: 0x14A6
Length: 32
Network Mask: /24
Attached Router: 2.2.2.2
Attached Router: 1.1.1.1
R1#sh ip ospf database router self-originate
OSPF Router with ID (1.1.1.1) (Process ID 1)
Router Link States (Area 0)
LS age: 400
Options: (No TOS-capability, DC)
LS Type: Router Links
Link State ID: 1.1.1.1
Advertising Router: 1.1.1.1
LS Seq Number: 80000002
Checksum: 0x729C
Length: 48
Number of Links: 2
Link connected to: a Stub Network
(Link ID) Network/subnet number: 11.11.11.11
(Link Data) Network Mask: 255.255.255.255
Number of MTID metrics: 0
TOS 0 Metrics: 1
Link connected to: a Transit Network
(Link ID) Designated Router address: 192.168.0.1
(Link Data) Router Interface address: 192.168.0.1
Number of MTID metrics: 0
TOS 0 Metrics: 10
Here's an example of whe LSA 2 could be useful (not found in the original answer):
R1----|----R2----|----R3 - all connected on broadcast medium.
Let's say R3 link goes down:
R1----|----R2----|
R2 will detect R3 going down when the dead timer expires. But how does R1 find out about R3 going down, because R2 will not change it's type 1 LSA (R2's link towards R3 is still up). The answer is that R2 will flood a type 2 LSA in which it says R3 is not part of the pseudonode anymore. Upon receiving this update, R1 will delete the routes that used R3 as transit. Interestingly, R1 still has R3 type 1 LSA. It just sees that the graph is interrupted (from the type 2 lsa sent by R2).
I think a reason is that in a Router-LSA the network is only represented as the IP address (no netmask) of the DR of that network while both the IP and the netmask are included in the Network-LSA.
Conceptually it is the DR that identifies the network, not an average router linked to the network.
Another reason is that such a Network-LSA will be sent to others and time out as a single unit. For example, a retiring DR can flush its old Network-LSA so that that network will be deleted from other routers' link state DB.
Link state advertisements form the basis of this type of protocol. without them and their hello and dead timers there would be no way to assure the topology and links were still active.
Link state protocols depend on these, whereas EIGRP and other distance vector protocols depend more upon the path of data and path cost determined by bandwidth availability, latency etc. They also do not have regular "updates" updates are sent when necessary such as when a link is found to be inactive.
With OSPF and LSAs whole updates of the topology table are sent regularly, they depend on similar items, such as distance and bandwidth but they are calculated differently due to the algorithm used in OSPF.
I prefer EIGRP but that is not an option in non Cisco land, its just a more efficient and simpler protocol to configure IMO.
I live in an all Juniper world, so eIGRP is a thing of the past, OSPF and the different type of LSA advertisements are a necessity to know.
