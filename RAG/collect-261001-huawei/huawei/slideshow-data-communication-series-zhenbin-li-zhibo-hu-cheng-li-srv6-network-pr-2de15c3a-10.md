---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-10
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: ["China", "Huawei", "Intel"]
dates: []
keywords: ["distribution", "ethernet", "intel"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1327, 1462]
sha256: 45709b1bac85a4afcca683f15b3af207091bc9fcfdbcd558469ad3acfb73021f
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

Huawei proposed POF to provide programmable forwarding logic for
switches, which is missing from OpenFlow.
Similar to OpenFlow, the architecture of POF includes two parts:
the control plane (POF controller) and the data plane (POF Forwarding
Element (FE)).
64.
SRv6 Background ◾17
POF makes the forwarding plane totally protocol-oblivious. Te POF
FE has no need to understand the packet format. All an FE needs to do
is, under the instruction of its controller, extract and assemble the search
keys (located by one or more<ofset, length>tuples) from the packet
header, conduct table lookups, and then execute the associated instruc-
tions. As a result, the FE can easily support new protocols and forwarding
requirements in the future.[21]
In summary, POF is an SDN technology that comprehensively abstracts
network processing into a more interoperable, protocol-oblivious process,
while providing the capabilities of programming forwarding rules and
forwarding logic. Figure 1.8 shows the architecture of the POF hardware
and sofware switches.
BecausePOFsupportsprotocol-obliviousforwarding,itcanbedeployed
on any network, including non-Ethernet networks such as Named Data
Network (NDN) and Content-Centric Network (CCN). In addition, the
POF FE supports adding metadata to packets so that stateful processing
can be supported.
(a) (b)
FIGURE 1.8 POF hardware switch (a) and POF sofware switch (b).
65.
18 ◾ SRv6Network Programming
However, POF is more complex than OpenFlow. Furthermore, a Flow
Instruction Set[21] needs to be defned to implement complex instruction-
based scheduling, which afects forwarding performance to some extent.
Terefore, POF has not made much progress in commercial use.
1.4.3 P4
In order to solve the insufcient programmability of OpenFlow, Professors
Nick McKeown (Stanford) and Jennifer Rexford (Princeton), among oth-
ers, proposed P4.[21] P4 is a high-level language for programming protocol-
independent packet processors. P4 can be used to confgure switches to
express how to parse, process, and forward packets.
P4 supports programmable packet processing on a device by defning a
P4 program containing the following components: headers, parsers, tables,
actions, and control programs.[22] A P4 program can be compiled and run
on the common abstract forwarding model-based devices. Figure 1.9
shows the common abstract forwarding model.
With P4 you can confgure and program how a switch processes
packets even afer the switch is deployed, thereby eliminating the need to
purchase new devices to support new features. Tis innovation solves the
insufcient programmability of OpenFlow. In addition, P4 programma-
bility enables switches to support protocol-independent packet forward-
ing, which means that diferent protocols can be supported by defning
associated P4 programs.
FIGURE 1.9 Common abstract forwarding model.
66.
SRv6 Background ◾19
Although P4 has certain technical advantages, it has not made much
progress in commercial deployment. One reason is that networks are
evolving nowhere fast. Te standardization of new network features is not
solely up to one device vendor. It is usually discussed for several years
among carriers, device vendors, and other relevant parties in the industry.
With such a long process, the beneft of supporting new features in a short
time provided by P4 is no longer signifcant. Another reason is that fully
centralized SDN has insufcient reliability and response speed but high
requirements on the controller. However, carriers do not need to com-
pletely reconstruct their networks by overturning the distributed routing
protocol architecture. Instead, the better solution is to provide the global
trafc optimization based on existing distributed routing protocols. So
far, P4 has not yet been put into large-scale commercial use. On June 11,
2019, Barefoot, a P4 startup, was ultimately acquired by Intel.
1.4.4 SR
OpenFlow, POF, and P4 were originally designed to provide programma-
bility for networks. But is revolutionary innovation necessary for network
programmability? Te answer is no.
In 2013, SR was proposed, which is a transitional extension based
on the existing network and provides network programmability. SR is
a source routing paradigm which can program the forwarding path of
packets by allowing the ingress of the path to insert forwarding instruc-
tions into packets. Te core idea of SR is to combine diferent segments
into a path and insert segment information into packets at the ingress of
the path to guide packet forwarding. A segment is an instruction, which
is executed on a node for packet forwarding or processing. It can refer to
a specifc interface or the shortest path to a node through which a packet
is forwarded. Such a segment is identifed by a Segment Identifer (SID).
Currently, there are two data planes for SR: MPLS and IPv6. When SR
is applied to the MPLS data plane, it is called SR-MPLS, and the SID is
encoded as an MPLS label. When SR is applied to the IPv6 data plane, it is
called SRv6, and the SID is encoded as an IPv6 address.
Te design of SR is easily comparable to many real-life examples, for
instance, traveling by train or plane. Te following uses travel as an exam-
ple to further explain SR.
If fying from Haikou to London requires a stop in Guangzhou and in
Beijing, we need to buy three tickets: Haikou to Guangzhou, Guangzhou
to Beijing, and Beijing to London.
67.
20 ◾ SRv6Network Programming
With three tickets in hand, we know we will fy from Haikou to
Guangzhou on fight HU7009 according to the frst ticket; when arriving
in Guangzhou, we take fight HU7808 to Beijing, according to the second
ticket; when arriving in Beijing, we fy to London by fight CA937, accord-
ing to the last ticket. Ultimately, we fy to London segment by segment,
using the three tickets.
Te packet forwarding on an SR network is similar. As shown in
Figure 1.10, a packet enters the SR network from node A. From the des-
tination address, node A knows that the packet needs to pass through
nodes B and C before reaching node D. Terefore, node A inserts the SIDs
of nodes B, C, and D into the packet header to indicate how to steer the
packet. Te packet will be sent to node B, node C, and fnally node D
according to the SID information in the packet header.
Compared with RSVP-TE MPLS, SR-MPLS has the following
advantages:
• Simplifes the control plane. In SR-MPLS, label distribution proto-
cols such as RSVP-TE are not needed any more. Instead, only IGP
and Border Gateway Protocol (BGP) extensions are required in the
SR-MPLS control plane, thereby reducing the number of control
plane protocols.
• Simplifes network states. When RSVP-TE is used on an MPLS net-
work, nodes need to maintain per-fow states. In contrast, on an
SR-MPLS network, only the ingress node needs to maintain per-fow
states, while the transit and egress nodes do not.
FIGURE 1.10 Packet forwarding process with SR.
68.
SRv6 Background ◾21
SR-MPLS provides network programming capabilities by reusing the
existing MPLS forwarding mechanism. Terefore, it can support smooth
upgrades from existing MPLS networks to SR-MPLS networks. In this
way, SR-MPLS, as the transitional innovation, is much easier to be adopted
by the industry. In addition, SR retains the distributed intelligence of the
network while introducing the global trafc optimization of the SDN con-
troller. Tis makes implementation more practical, and it will go further.
So far, SR has already become the de facto SDN standard.
Although SR-MPLS based on MPLS data plane can provide good pro-
grammability, it cannot satisfy services that need to carry metadata, such
as SFC and IOAM, as MPLS encapsulation has relatively poor extensibil-
ity. Compared with SR-MPLS, SRv6, which is based on IPv6 data plane,
not only inherits all the advantages of SR-MPLS, but also provides better
extensibility.
1.5 KEY TO ALL IP 2.0: SRv6
As mentioned before, the key to speeding up IPv6 deployment lies in fnd-
