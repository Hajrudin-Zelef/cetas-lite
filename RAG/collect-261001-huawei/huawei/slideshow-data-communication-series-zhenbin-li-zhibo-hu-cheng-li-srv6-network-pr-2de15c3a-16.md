---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-16
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [2277, 2434]
sha256: 4f07f0e68d0c0d165d44e172231320043474fe328cbe77c3e1f1e861436a0099
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

Te End.DX4 SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Forwards the inner IPv4 packet to the IPv4 Layer 3 adjacency bound
to the End.DX4 SID.
Te pseudocode of an End.DX4 SID is as follows:
S01. If (Upper-Layer Header type != 4) {
S02. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the
offset of the upper-layer header, interrupt packet processing and
discard the packet
S03. }
S04. Remove the outer IPv6 Header with all its extension headers
S05. Forward the exposed IPv4 packet to a member of specific
L3 adjacencies
2.4.3.6 End.DT6 SID
End.DT6 (short for decapsulation and specifc IPv6 table lookup) sup-
ports packet decapsulation and lookup in a specifed IPv6 routing table for
packet forwarding. End.DT6 SIDs are mainly used in L3VPNv6 scenarios
as per-VPN SIDs.
End.DT6 can be disassembled into End+D+T6, where D indi-
cates decapsulation and T6 indicates IPv6 table lookup for forwarding.
Terefore, each End.DT6 SID needs to be bound to an IPv6 routing table,
which can be either the IPv6 routing table of a VPN instance or a public
IPv6 routing table.
Te End.DT6 SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Searches the IPv6 routing table bound to the End.DT6 SID to for-
ward the inner IPv6 packet.
Te pseudocode of an End.DT6 SID is as follows:
S01. If (Upper-Layer Header type != 41) {
S02. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the
offset of the upper-layer header, interrupt packet processing
94.
SRv6 Fundamentals ◾47
and discard the packet
S03. }
S04. Remove the outer IPv6 Header with all its extension headers
S05. Set the packet's associated FIB table to the specific IPv6 FIB
S06. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new IPv6 destination
2.4.3.7 End.DT4 SID
End.DT4 (short for decapsulation and specifc IPv4 table lookup) supports
packet decapsulation and lookup in a specifed IPv4 routing table for
packet forwarding. End.DT4 SIDs are mainly used in L3VPNv4 scenarios
as per-VPN SIDs.
End.DT4 can be disassembled into End+D+T4, where D indi-
cates decapsulation and T4 indicates IPv4 table lookup for forwarding.
Terefore, each End.DT4 SID needs to be bound to an IPv4 routing table,
which can be either the IPv4 routing table of a VPN instance or a public
IPv4 routing table.
Te End.DT4 SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Searches the IPv4 routing table bound to the End.DT4 SID to
forward the inner IPv4 packet.
Te pseudocode of an End.DT4 SID is as follows:
S01. If (Upper-layer Header type == 4) {
S02. Remove the outer IPv6 Header with all its extension headers
S03. Set the packet's associated FIB table to the specific IPv4 FIB
S04. Resubmit the packet to the egress IPv4 FIB lookup and transmission
to the new IPv4 destination
S05. } Else if (Upper-layer Header type == 41) {
S06. Remove the outer IPv6 Header with all its extension headers
S07. Set the packet's associated FIB table to the specific IPv6 FIB
S08. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new IPv6 destination
S09. } Else {
S10. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the offset
of the upper-layer header, interrupt packet processing and discard
the packet
S11. }
2.4.3.8 End.DT46 SID
End.DT46 (short for decapsulation and specifc IP table lookup) supports
packet decapsulation and lookup in a specifed IPv4 or IPv6 routing table
95.
48 ◾ SRv6Network Programming
for packet forwarding. End.DT46 SIDs are mainly used in L3VPN sce-
narios as per-VPN SIDs.
End.DT46 can be disassembled into End+D+T46, where D indicates
decapsulation and T46 indicates IPv4 or IPv6 table lookup for forward-
ing. Terefore, each End.DT46 SID needs to be bound to an IPv4 or IPv6
routing table, which can be either the IPv4 or IPv6 routing table of a VPN
instance or a public IPv4 or IPv6 routing table.
Te End.DT46 SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. According to the Layer 3 protocol type of the inner IP packet,
searches the IPv4 or IPv6 routing table bound to the End.DT46 SID
to forward the inner IP packet.
Te pseudocode of an End.DT46 SID is as follows:
S01. If (Upper-layer Header type == 4) {
S02. Remove the outer IPv6 Header with all its extension headers
S03. Set the packet's associated FIB table to the specific IPv4 FIB
S04. Resubmit the packet to the egress IPv4 FIB lookup and transmission
to the new IPv4 destination
S05. } Else if (Upper-layer Header type == 41) {
S06. Remove the outer IPv6 Header with all its extension headers
S07. Set the packet's associated FIB table to the specific IPv6 FIB
S08. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new IPv6 destination
S09. } Else {
S10. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the offset
of the upper-layer header, interrupt packet processing and discard
the packet
S11. }
2.4.3.9 End.DX2 SID
End.DX2 (short for decapsulation and L2 cross-connect) supports packet
decapsulation and forwarding through a specifed Layer 2 outbound inter-
face. End.DX2 SIDs are mainly used in Layer 2 Virtual Private Network
(L2VPN) and EVPN Virtual Private Wire Service (VPWS)[5] scenarios.
End.DX2 can be disassembled into End+D+X2, where D indicates
decapsulation and X2 indicates cross-connect (meaning that a packet
needs to be directly forwarded through a specifed Layer 2 interface).
Terefore, each End.DX2 SID needs to be bound to a Layer 2 outbound
interface.
96.
SRv6 Fundamentals ◾49
Te End.DX2 SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Forwards the inner Ethernet frame to the Layer 2 outbound interface
bound to the End.DX2 SID.
Te pseudocode of an End.DX2 SID is as follows:
S01. If (Upper-Layer Header type != 143) {
S02. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the offset
of the upper-layer header, interrupt packet processing and
discard the packet
S03. }
S04. Remove the outer IPv6 Header with all its extension headers and
forward the Ethernet frame to the specific outgoing L2 interface
2.4.3.10 End.DX2V SID
End.DX2V (short for decapsulation and VLAN L2 table lookup) supports
packet decapsulation and lookup in a specifed Layer 2 table based on the
inner VLAN information of the packet to be forwarded. End.DX2V SIDs
are mainly used in EVPN fexible cross-connect scenarios.[5]
End.DX2V can be disassembled into End+D+X2V, where D indicates
decapsulation and X2V indicates VLAN Layer 2 table lookup. Terefore,
each End.DX2V SID needs to be bound to a Layer 2 table.
Te End.DX2V SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Searches the Layer 2 table bound to the End.DX2V SID based on the
VLAN information in the inner Ethernet frame.
Te pseudocode of an End.DX2V SID is based on that of an End.DX2 SID,
but with line S04 being modifed as follows:
S04. Remove the outer IPv6 Header with all its extension headers, lookup
the exposed inner VLANs in the specific L2 table, and forward via
the matched table entry.
2.4.3.11 End.DT2U SID
End.DT2U (short for decapsulation and unicast MAC L2 table lookup)
supports packet decapsulation, inner source MAC address learning and
97.
50 ◾ SRv6Network Programming
saving to a specifed Layer 2 table, and lookup in the table for packet for-
warding based on the inner destination MAC address. End.DT2U SIDs
are mainly used in EVPN bridging unicast scenarios.[5]
End.DT2U can be disassembled into End+D+T2U, where D indicates
decapsulation and T2U indicates Layer 2 table lookup for unicast forward-
