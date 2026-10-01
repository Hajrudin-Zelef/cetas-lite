---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-17
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [2435, 2597]
sha256: 3ce3d42781e5b695541aaa31870c1baf749a3ffecf8580a9964394d3d6f5a749
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

ing. Terefore, each End.DT2U SID needs to be bound to a Layer 2 table.
Te End.DT2U SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Learns the source MAC address of the inner Ethernet frame.
3. Saves the source MAC address to the Layer 2 table bound to the End.
DT2U SID.
4. Searches the Layer 2 table bound to the End.DT2U SID based on the
destination MAC address of the inner Ethernet frame.
5. Forwards the corresponding packet accordingly.
Te pseudocode of an End.DT2U SID is as follows:
S01. If (Upper-Layer Header type != 143) {
S02. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the
offset of the upper-layer header, interrupt packet processing
and discard the packet
S03. }
S04. Remove the IPv6 header and all its extension headers
S05. Learn the exposed inner MAC Source Address in the specific
L2 table (T)
S06. Lookup the exposed inner MAC Destination Address in table T
S07. If (matched entry in T) {
S08. Forward via the matched table T entry
S09. } Else {
S10. Forward via all outgoing L2 interfaces entries in table T
S11. }
2.4.3.12 End.DT2M SID
End.DT2M (short for decapsulation and L2 table fooding) supports
packet decapsulation, inner source MAC address learning and saving
to a specifed Layer 2 table, and packet forwarding to Layer 2 outbound
interfaces excluding the specifed one. End.DT2M SIDs are mainly used
in EVPN bridging Broadcast & Unknown-unicast & Multicast (BUM)[5]
and EVPN E-Tree[6] scenarios.
End.DT2M can be disassembled into End+D+T2M, where D indi-
cates decapsulation and T2M indicates Layer 2 table lookup for multicast
98.
SRv6 Fundamentals ◾51
forwarding. Terefore, each End.DT2M SID needs to be bound to a Layer
2 table and may carry EVPN ESI fltering or EVPN E-Tree parameters in
order to exclude specifed outbound interfaces during fooding.
Te End.DT2M SID instruction includes the following operations:
1. Removes the outer IPv6 header and SRH.
2. Learns the source MAC address of the inner Ethernet frame.
3. Saves the source MAC address to the Layer 2 table bound to the End.
DT2M SID.
4. Forwards the inner Ethernet frame through Layer 2 outbound inter-
faces excluding the interface specifed by the parameters carried in
the End.DT2M SID.
Te pseudocode of an End.DT2M SID is as follows:
S01. If (Upper-Layer Header type != 143) {
S02. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the offset
of the upper-layer header, interrupt packet processing and
discard the packet
S03. }
S04. Remove the IPv6 header and all its extension headers
S05. Learn the exposed inner MAC Source Address in the specific L2 table
S06. Forward via all outgoing L2 interfaces excluding the one specified
in Arg.FE2
2.4.3.13 End.B6.Insert SID
End.B6.Insert (short for endpoint bound to an SRv6 Policy with Insert)
supports the application of specifed SRv6 Policies[7] to packets. End.
B6.Insert SIDs instantiate binding SIDs in SRv6 and are used in scenarios
where TE can be fexibly implemented across SRv6 domains.
End.B6.Insert can be disassembled into End+B6+Insert, where B6
indicates the application of an SRv6 Policy and Insert indicates the inser-
tion of an SRH afer the IPv6 header. Terefore, each End.B6.Insert SID
needs to be bound to an SRv6 Policy.
Te End.B6.Insert SID instruction includes the following operations:
1. Inserts an SRH (including segment lists) afer the IPv6 header.
2. Sets the DA to the frst SID of a specifed SRv6 Policy.
3. Searches the IPv6 routing table.
4. Forwards the new IPv6 packet accordingly.
99.
52 ◾ SRv6Network Programming
Te pseudocode of an End.B6.Insert SID is as follows:
S01. When an SRH is processed {
S02. If (Segments Left == 0) {
S03. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the
offset of the upper-layer header, interrupt packet processing
and discard the packet
S04. }
S04. If (IPv6 Hop Limit <= 1) {
S05. Send an ICMP Time Exceeded message to the Source Address,
code 0 (Hop limit exceeded in transit), interrupt packet
processing and discard the packet
S06. }
S07. max_LE = (Hdr Ext Len / 2) - 1
S08. If ((Last Entry > max_LE) or (Segments Left > (Last Entry+1)){
S09. Send an ICMP Parameter Problem to the Source Address, code 0
(Erroneous header field encountered), pointer set to the
Segments Left field, interrupt packet processing and discard
the packet
S11. }
S12. Decrement Hop Limit by 1
S13. Insert a new SRH after the IPv6 Header and the SRH
contains the list of segments of the specific SRv6 Policy
S14. Set the IPv6 DA to the first segment of the SRv6 Policy
S15. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new IPv6 destination
S16. }
Inaddition,theEnd.B6.Insert.RedinstructionisanoptimizationoftheEnd.
B6.Insert instruction, except that the former requires a reduced SRH (i.e.,
the SRH excluding the frst SID of the involved SRv6 Policy) to be added.
2.4.3.14 End.B6.Encaps SID
End.B6.Encaps (short for endpoint bound to an SRv6 Policy with Encaps)
supports the application of specifed SRv6 Policies[7] to packets. End.
B6.Encaps SIDs also instantiate binding SIDs in SRv6 and are used in sce-
narios where TE can be fexibly implemented across SRv6 domains.
End.B6.Encaps can be disassembled into End+B6+Encaps, where
B6 indicates the application of an SRv6 Policy and Encaps indicates the
encapsulation of an outer IPv6 header and SRH. Terefore, each End.
B6.Encaps SID needs to be bound to an SRv6 Policy.
Te End.B6.Encaps SID instruction includes the following operations:
1. Decrements the SL value of the inner SRH by 1.
2. Updates the DA feld of the inner IPv6 header with the SID to which
the SL feld is pointing.
100.
SRv6 Fundamentals ◾53
3. Encapsulates an IPv6 header and SRH (including segment lists).
4. Sets the source address to the address of the current node and the
DA to the frst SID of the involved SRv6 Policy.
5. Sets other felds in the outer IPv6 header.
6. Searches the IPv6 routing table.
7. Forwards the new IPv6 packet accordingly.
Te pseudocode of an End.B6.Encaps SID is as follows:
S01. When an SRH is processed {
S02. If (Segments Left == 0) {
S03. Send an ICMP Parameter Problem message to the Source Address,
code 4 (SR Upper-layer Header Error), pointer set to the offset
of the upper-layer header, interrupt packet processing and
discard the packet
S04. }
S05. If (IPv6 Hop Limit <= 1) {
S06. Send an ICMP Time Exceeded message to the Source Address, code
0 (Hop limit exceeded in transit), interrupt packet processing
and discard the packet
S07. }
S08. max_LE = (Hdr Ext Len / 2) - 1
S09. If ((Last Entry > max_LE) or (Segments Left > (Last Entry+1)) {
S10. Send an ICMP Parameter Problem to the Source Address, code 0
(Erroneous header field encountered), pointer set to the
Segments Left field, interrupt packet processing and discard
the packet
S11. }
S12. Decrement Hop Limit by 1
S13. Decrement Segments Left by 1
S14. Update the inner IPv6 DA with inner Segment List[Segments Left]
S15. Push a new IPv6 header with its own SRH containing the list of
segments of the SRv6 Policy
S16. Set the outer IPv6 SA to itself
S17. Set the outer IPv6 DA to the first SID of the SRv6 Policy
S18. Set the outer Payload Length, Traffic Class, Flow Label and
Next Header fields
S19. Resubmit the packet to the egress IPv6 FIB lookup and transmission
to the new IPv6 destination
S20. }
In addition, the End.B6.Encaps.Red instruction is an optimization of the
End.B6.Encaps instruction, with the diference lying in the fact that the
former requires a reduced SRH to be added.
2.4.3.15 End.BM SID
End.BM (short for endpoint bound to an SR-MPLS Policy) supports the
application of a specifed SR-MPLS Policy[7] to packets. End.BM SIDs
