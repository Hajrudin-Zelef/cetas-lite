---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-15
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [2073, 2276]
sha256: e5bcce64254e22621295dda12c28179087fbf888608707aa3e2f22260d8eca52
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

interface. An End.X SID is
similar to an SR-MPLS
adjacency SID.
Used in scenarios where
multiple routing tables exist.
Used in L3VPNv6 scenarios
where packets are forwarded to
a Customer Edge (CE) through
a specifed IPv6 adjacency.
Used in L3VPNv4 scenarios
where packets are forwarded to
a CE through a specifed IPv4
adjacency.
Used in L3VPNv6 scenarios.
Used in L3VPNv4 scenarios.
Used in L3VPNv4/L3VPNv6
scenarios.
(Continued)
89.
42 ◾ SRv6Network Programming
TABLE 2.4 (Continued) Functions of Common SRv6 Instructions
Instruction Function Description
End.DX2 Decapsulates a packet and
forwards it through a specifed
Layer 2 outbound interface.
End.DX2V Decapsulates a packet and
searches a specifed Layer 2 table
for packet forwarding based on
inner VLAN information.
End.DT2U Decapsulates a packet, learns the
inner source MAC address in a
specifed Layer 2 table, and
searches the table for packet
forwarding based on the inner
destination MAC address.
End.DT2M Decapsulates a packet, learns the
inner source MAC address in a
specifed Layer 2 table, and
forwards the packet to Layer 2
outbound interfaces excluding
the specifed one.
End.B6.Insert Inserts an SRH and applies a
specifed SRv6 Policy.
End.B6.Insert. Inserts a reduced SRH and
Red applies a specifed SRv6 Policy.
End.B6.Encaps Encapsulates an outer IPv6
header and SRH, and applies a
specifed SRv6 Policy.
End.B6.Encaps. Encapsulates an outer IPv6
Red header and reduced SRH, and
applies a specifed SRv6 Policy.
End.BM Inserts an MPLS label stack and
applies a specifed SR-MPLS
Policy.
Application Scenario
Used in EVPN VPWS scenarios.
Used in EVPN Virtual Private
LAN Service (VPLS) scenarios.
Used in EVPN VPLS unicast
scenarios.
Used in EVPN VPLS multicast
scenarios.
Used in scenarios such as trafc
steering into an SRv6 Policy in
Insert mode, tunnel stitching,
and Sofware Defned Wide
Area Network (SD-WAN)
route selection.
Used in scenarios such as trafc
steering into an SRv6 Policy in
Insert or Reduced mode,
tunnel stitching, and SD-WAN
route selection.
Used in scenarios such as trafc
steering into an SRv6 Policy
in Encaps mode, tunnel
stitching, and SD-WAN route
selection.
Used in scenarios such as trafc
steering into an SRv6 Policy in
Encaps or Reduced mode,
tunnel stitching, and SD-WAN
route selection.
Used in SRv6 and SR-MPLS
interworking scenarios where
trafc needs to be steered into
an SR-MPLS Policy.
90.
SRv6 Fundamentals ◾43
2.4.3.1 End SID
End is the most basic SRv6 instruction. A SID bound to the End instruc-
tion is called an End SID, which identifes a node. An End SID instructs
a network node to forward a packet to the node that advertises the SID.
Afer receiving the packet, the node performs the operations defned by
the SID for packet processing.
Te End SID instruction includes the following operations:
1. Decrements the SL value by 1.
2. Obtains the next SID from the SRH based on the SL value.
3. Updates the DA feld in the IPv6 header with the next SID.
4. Searches the routing table for packet forwarding.
Other parameters, such as Hop Limit, are processed according to the
normal forwarding process.
Te pseudocode of an End SID is as follows:
S01. When an SRH is processed {
S02. If (Segments Left == 0) {
S03. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the
offset of the upper-layer header, interrupt packet processing
and discard the packet
S04. }
S05. If (IPv6 Hop Limit <= 1) {
S06. Send an ICMP Time Exceeded message to the Source Address,
code 0 (Hop limit exceeded in transit), interrupt packet
processing and discard the packet
S07. }
S08. max_LE = (Hdr Ext Len / 2) - 1
S09. If ((Last Entry > max_LE) or (Segments Left > Last Entry+1)) {
S10. Send an ICMP Parameter Problem to the Source Address, code
0 (Erroneous header field encountered), pointer set to the
Segments Left field, interrupt packet processing and discard
the packet
S11. }
S12. Decrement Hop Limit by 1
S13. Decrement Segments Left by 1
S14. Update IPv6 DA with Segment List[Segments Left]
S15. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new destination
S16. }
2.4.3.2 End.X SID
End.X (short for Layer-3 cross-connect) supports packet forwarding
to a Layer 3 adjacency over a specifed link. End.X SIDs can be used in
91.
44 ◾ SRv6Network Programming
scenarios including Topology Independent Loop Free Alternate (TI-LFA)
and strict explicit path-based TE.
In essence, End.X SIDs are developed based on End SIDs. End.X can
be disassembled into End+X, where X indicates cross-connect (meaning
that a packet needs to be directly forwarded to a specifed Layer 3 adja-
cency). Terefore, each End.X SID needs to be bound to one or a group of
Layer 3 adjacencies.
Te End.X SID instruction includes the following operations:
1. Decrements the SL value by 1.
2. Obtains the next SID from the SRH based on the SL value.
3. Updates the DA feld in the IPv6 header with the next SID.
4. Forwards the IPv6 packet to the Layer 3 adjacency bound to the
End.X SID.
Te pseudocode of an End.X SID is based on that of an End SID, but with
line S15 being modifed as follows:
S15. Resubmit the packet to the IPv6 module for transmission to the new
IPv6 destination via a member of specific L3 adjacencies
2.4.3.3 End.T SID
End.T (short for specifc IPv6 table lookup) supports packet forwarding
by searching the specifed IPv6 routing table. End.T SIDs can be used in
common IPv6 routing and VPN scenarios.
End.T SIDs are also developed based on End SIDs. End.T can be
disassembled into End+T, where T indicates table lookup for packet
forwarding. Terefore, each End.T SID needs to be bound to an IPv6
routing table.
Te End.T SID instruction includes the following operations:
1. Decrements the SL value by 1.
2. Obtains the next SID from the SRH based on the SL value.
3. Updates the DA feld in the IPv6 header with the next SID.
4. Searches the specifed routing table for IPv6 packet forwarding.
92.
SRv6 Fundamentals ◾45
Te pseudocode of an End.T SID is based on that of an End SID, but with
line S15 being modifed as follows:
S15.1. Set the packet's associated FIB table to the specific IPv6 FIB
S15.2. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new IPv6 destination
2.4.3.4 End.DX6 SID
End.DX6(shortfordecapsulationandIPv6cross-connect)supportspacket
decapsulation and forwarding to a specifed IPv6 Layer 3 adjacency. End.
DX6 SIDs are mainly used in L3VPNv6 scenarios as per-CE VPN SIDs.
End.DX6 can be disassembled into End+D+X6, where D indicates
decapsulation, and X6 indicates IPv6 cross-connect (meaning that a
packet needs to be directly forwarded to a specifed IPv6 Layer 3 adja-
cency). Terefore, each End.DX6 SID needs to be bound to one or a group
of IPv6 Layer 3 adjacencies.
Te End.DX6 SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Forwards the inner IPv6 packet to the IPv6 Layer 3 adjacency bound
to the End.DX6 SID.
Te pseudocode of an End.DX6 SID is as follows:
S01. If (Upper-Layer Header type != 41) {
S02. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the offset
of the upper-layer header, interrupt packet processing and discard
the packet
S03. }
S04. Remove the outer IPv6 Header with all its extension headers
S05. Forward the exposed IPv6 packet to a member of
specific L3 adjacencies
2.4.3.5 End.DX4 SID
End.DX4 (short for decapsulation and IPv4 cross-connect) supports packet
decapsulation and forwarding to a specifed IPv4 Layer 3 adjacency. End.
DX4 SIDs are mainly used in L3VPNv4 scenarios as per-CE VPN SIDs.
End.DX4 can be disassembled into End+D+X4, where D indicates
decapsulation, and X4 indicates IPv4 cross-connect (meaning that a packet
needs to be directly forwarded to a specifed IPv4 Layer 3 adjacency).
93.
46 ◾ SRv6Network Programming
Terefore, each End.DX4 SID needs to be bound to one or a group of IPv4
Layer 3 adjacencies.
