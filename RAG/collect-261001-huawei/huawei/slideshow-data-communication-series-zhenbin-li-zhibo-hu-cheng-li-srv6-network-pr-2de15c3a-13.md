---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-13
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1737, 1877]
sha256: ef6c4e4d77e7cba02bb9f32dd2e69e6d5df3c4e98d0ac28e0a5c531df1fff76c
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

In Chapter 1, we introduced the development of network programming
and the advantages of SRv6 in this regard. Network programming stems
from computer programming, through which we can convert our intent
into a series of instructions that computers can understand and execute to
meet our requirements. Similarly, if service intent can be translated into a
series of device-executable forwarding instructions on a network, network
programming can be achieved, ultimately meeting service customization
requirements. Figure 2.1 illustrates the implementation of network pro-
gramming as compared to computer programming.
SRv6 was introduced to translate network functions into instructions
and encapsulate the instructions into 128-bit IPv6 addresses. As such,
service requirements can be translated into an ordered list of instruc-
tions, which can then be executed by network devices along the service
31
79.
32 ◾ SRv6Network Programming
FIGURE 2.1 Computer programming and network programming.
forwarding path, thereby achieving fexible orchestration and on-demand
customization of services on an SRv6 network.
2.2 NETWORK INSTRUCTIONS: SRv6 SEGMENTS
A computer instruction typically consists of an opcode and an operand. Te
former determines the operation to be performed; the latter determines the
data, memory address, or both to be used in the computation. Similarly, net-
work instructions, which are called SRv6 segments, need to be defned for
SRv6 network programming, and they are identifed using 128-bit SIDs.[2]
Each SRv6 SID usually consists of three felds, as shown in Figure 2.2.
Te three felds are described as follows:
1. Te Locator feld identifes the location of a network node and is
used for other nodes to route and forward packets to this identi-
fed node. A locator has two important characteristics: routable and
aggregatable. “Routable” means that the corresponding locator route
can be advertised by a node to other nodes on the network through
FIGURE 2.2 SRv6 SID.
80.
SRv6 Fundamentals ◾33
an IGP so that those nodes can forward packets to the node through
the locator route. “Aggregatable” means that the locator route can
be aggregated. Te two characteristics help tackle problems such as
high network complexity and large network scale.
Te locator length is variable, so SRv6 SIDs can be used on small
to large networks.
Te following is an example of locator confguration.
<HUAWEI> system-view
[~HUAWEI] segment-routing ipv6
[~HUAWEI-segment-routing-ipv6] locator test1 ipv6-prefix 2001:db8:100:: 64
Afer the preceding confgurations are complete, a locator with the
64-bit prefx 2001:db8:100:: is generated to guide packet forwarding
to the node generating the corresponding instruction, thereby mak-
ing the instruction addressable.
2. Te Function feld specifes the forwarding behavior to be performed
and is similar to the opcode in a computer instruction. In SRv6
network programming, forwarding behaviors are expressed using
diferent functions. You can think of SIDs as similar to computer
instructions, in that each type of SID identifes corresponding for-
warding behaviors, such as forwarding packets to a specifed link or
searching a specifed table for packet forwarding.
Te following is an example of End.X function confguration.
[~HUAWEI-segment-routing-ipv6-locator] opcode ::1 end-x interface
GigabitEthernet3/0/0 next-hop 2001:db8:200::1
Te preceding example defnes an End.X function with an opcode
of ::1. If the Arguments feld is not specifed, the locator pre-
fx 2001:db8:100:: and function opcode ::1 form the SRv6 SID
2001:db8:100::1, which is a manually confgured End.X SID. An
End.X SID identifes one or a group of Layer 3 network adjacencies
connected to a network node. Te forwarding behavior bound to
the SID is to update the next SID into the IPv6 Destination Address
(DA) feld and forward packets to the corresponding neighbor-
ing node through the interface specifed by the End.X SID. In the
preceding example, the End.X function instructs the local node
to forward packets from GigabitEthernet 3/0/0 to the next hop at
2001:db8:200::1.
81.
34 ◾ SRv6Network Programming
3. Te Arguments (Args) feld is optional. It is used to defne param-
eters for instruction execution and can contain fow, service, and any
other related information. For example, the Arg.FE2 parameter can
be specifed for an End.DT2M SID to exclude a specifc or a group
of outbound interfaces when the corresponding Layer 2 forward-
ing table is searched for multicast replication in Ethernet Segment
Identifer (ESI) fltering and Ethernet Virtual Private Network
(EVPN) Ethernet Tree (E-Tree) scenarios.
Much like computers that use limited instruction sets to implement
various computing functions, SRv6 also defnes instruction sets for for-
warding purposes.
Instructions are essential to network programming. Any E2E ser-
vice connection requirement can be expressed using an ordered set of
instructions. In current SR implementation, the source node encapsulates
ordered segment lists into packets to instruct specifed nodes to execute
corresponding instructions, thereby achieving network programmability.
With the increase of SRv6 application scenarios, SRv6 instruction sets will
evolve continuously.
2.3 NETWORK NODES: SRv6 NODES
An SRv6 network may contain nodes of the following roles[3]:
• SRv6 source node: a source node that encapsulates packets with
SRv6 headers
• Transit node: an IPv6 node that forwards SRv6 packets but does not
perform SRv6 processing
• SRv6 segment endpoint node (endpoint node for short): a node that
receives and processes SRv6 packets whose IPv6 DA is a local SID of
the node
A node plays a role based on the task it takes in SRv6 packet forwarding,
and it may play two or more roles. For example, it can be the source node
on one SRv6 path and a transit or endpoint node on another SRv6 path.
82.
SRv6 Fundamentals ◾35
2.3.1 SRv6 Source Node
An SRv6 source node steers a packet using an SRv6 segment list. If the
SRv6 segment list contains only one SID, and no Type Length Value (TLV)
or other information needs to be added to the packet, the DA feld of the
packet is set to the SID, without requiring SRH encapsulation.
An SRv6 source node can be either an SRv6-capable host where IPv6
packets originate or an edge device in an SRv6 domain.
2.3.2 Transit Node
A transit node is an IPv6 node that does not participate in SRv6 process-
ing on the SRv6 packet forwarding path; that is, the transit node just for-
wards IPv6 packets. Afer receiving an SRv6 packet, a transit node parses
the IPv6 DA feld in the packet. If the value of this feld is neither a locally
confgured SRv6 SID nor a local interface address, the transit node consid-
ers the SRv6 packet as an ordinary IPv6 packet. As such, it searches the
corresponding IPv6 routing table according to the longest match rule for
packet processing and forwarding. In this process, processing the DA as
an SRv6 SID or processing SRHs is not required.
A transit node can be either an ordinary IPv6 node or an SRv6-capable
node.
2.3.3 Endpoint Node
An endpoint node is a node that receives a packet destined for itself
(a packet of which the IPv6 DA is a local SID). Unlike transit nodes,
endpoint nodes have to process both SRv6 SIDs and SRHs.
To sum up, an SRv6 source node encapsulates packets with SRv6 head-
ers, a transit node processes and forwards the packets as common IPv6
packets, and an endpoint node processes both SRv6 SIDs and SRHs in the
packets, as shown in Figure 2.3.
Note: As recommended by RFC 3849, to avoid address conficts with live-
network addresses, IPv6 addresses with the prefx 2001:DB8::/32 are usu-
ally used as examples in books. However, 2001:DB8::/32 is too long to ft
neatly into fgures. Terefore, short addresses such as A1::1 and 1::1 are
ofen used in this book. Unless otherwise specifed, all IPv6 addresses in
this book are examples only.
83.
36 ◾ SRv6Network Programming
FIGURE 2.3 SRv6 nodes.
