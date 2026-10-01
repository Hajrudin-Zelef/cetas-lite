---
id: collect-261001-huawei/huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a-14
title: "slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["alignment"]
source: docs/RAG/collect-261001-huawei/slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a.md
source_anchor: ""
source_lines: [1878, 2072]
sha256: 3f2d79f0e9a31efd6325536c705bd58e1e1a8550371889276baebf704cd44b07
---

# slideshow-data-communication-series-zhenbin-li-zhibo-hu-cheng-li-srv6-network-pr-2de15c3a

2.4 NETWORK PROGRAM: SRv6 EXTENSION HEADER
2.4.1 SRv6 Extension Header Design
To implement SRv6, the original IPv6 header is extended to defne a new
type of routing header called SRH,[3] which carries segment lists and other
related information to explicitly specify an SRv6 path.
Figure 2.4 shows the format of the SRH.
Table 2.1 describes the felds in an SRH.
Te SRH stores a SID list used for implementing network services,
which is similar to an ordered list of computer instructions. Te segments
represented by Segment List [0] through Segment List [n] are similar to
FIGURE 2.4 IPv6 SRH format.
84.
SRv6 Fundamentals ◾37
TABLE 2.1 Fields in an SRH
Field Length
Next Header 8 bits
Hdr Ext Len 8 bits
Routing 8 bits
Type
Segments 8 bits
Lef (SL)
Last Entry 8 bits
Flags 8 bits
Tag 16 bit
Segment 128 bits × number
List [n] of segments
Optional Variable
TLV
Description
Type of the header following the SRH. Common
header types are as follows:
• 4: IPv4 encapsulation
• 41: IPv6 encapsulation
• 43: IPv6 routing header (IPv6-Route)
• 58: Internet Control Message Protocol version 6
(ICMPv6)
• 59: no Next Header for IPv6
Length of an SRH, excluding the frst 8 bytes.
Type of the extension routing header. SRHs have a
routing type value of 4.
Number of remaining segments.
Index of the last element in a segment list.
Reserved for special processing, such as Operations,
Administration, and Maintenance (OAM).
Whether a packet is part of a group of packets, such
as packets sharing the same set of properties.
Te nth segment in a segment list, expressed using a
128-bit IPv6 address.
Optional TLVs, such as Padding TLVs and
Hash-based Message Authentication Code
(HMAC) TLVs.
the instructions of a computer program, and Segment List [n] indicates
the frst instruction that needs to be executed. Te SL feld, which is simi-
lar to the Program Counter (PC) of a computer program, points to the
instruction that is being executed and can be initially set to n (the index of
the last segment in a segment list). Each time an instruction is executed,
the SL value is decremented by 1 to point to the next instruction to be
executed. From this, we can see that the SRv6 forwarding process can be
easily simulated using a computer program.
To make it easier to explain the SRv6 data forwarding process in this
book, a simplifed version of the SRH shown in Figure 2.5 (a) is used and
can be further simplifed as shown in Figure 2.5 (b).
2.4.2 SRH TLVs
As shown in Figure 2.4, SRHs use optional TLVs to ofer a higher level
of network programmability.[3] Te TLVs provide better extensibility and
85.
38 ◾ SRv6Network Programming
FIGURE 2.5 Abstract SRH.
can carry variable-length data, including encryption, authentication,
and performance monitoring information. Additionally, as we described
in previous sections, the SIDs in a segment list defne instructions to be
executed by nodes. Tis means that each SID can be executed only by the
node advertising it, and other nodes only perform basic table lookup and
forwarding operations. In contrast, the information carried in TLVs can
be processed by any node advertising a SID in the corresponding segment
list. To improve forwarding efciency, a local TLV processing policy that
determines, for example, whether to ignore all TLVs or process only certain
types of TLVs according to interface confgurations, needs to be defned.
SRHs currently support two types of TLVs: Padding TLVs and HMAC
TLVs.[3,4]
2.4.2.1 Padding TLV
Two types of Padding TLVs (Pad1 TLV and PadN TLV) are defned for
SRHs. If the length of subsequent TLVs is not a multiple of 8 bytes, Padding
TLVs can be used for alignment, thereby ensuring that the overall SRH
length is a multiple of 8 bytes. Because Padding TLVs are not instructive,
they are ignored by nodes during SRH processing.
Te Pad1 TLV defnes only the Type feld, which is used for padding a
single byte. If two or more bytes are required, this TLV cannot be used.
Figure 2.6 shows the format of the Pad1 TLV.
Te PadN TLV is used for padding multiple bytes. Figure 2.7 shows the
format of the PadN TLV.
Table 2.2 describes the felds in the PadN TLV.
86.
SRv6 Fundamentals ◾39
FIGURE 2.6 Pad1 TLV.
FIGURE 2.7 PadN TLV.
TABLE 2.2 Fields in the PadN TLV
Field Length Description
Type 8 bits TLV type.
Length 8 bits Length of the Padding feld.
Padding Variable Padding bits, which are used when the packet length is not a
multiple of 8 bytes. As the padding bits are not instructive, they
must be set to 0s before transmission and ignored upon receipt.
2.4.2.2 HMAC TLV
Te HMAC TLV is used to prevent key information in an SRH from being
tampered with. It is optional and requires 8-byte alignment. Figure 2.8
shows the format of the HMAC TLV.
Table 2.3 describes the felds in the HMAC TLV.
FIGURE 2.8 HMAC TLV.
87.
40 ◾ SRv6Network Programming
TABLE 2.3 Fields in the HMAC TLV
Field Length
Type 8 bits
Length 8 bits
D 1 bit
Reserved 15 bits
HMAC 32 bits
Key ID
HMAC Variable (up
to 256 bits)
Description
TLV type.
Length of the HMAC feld.
When this bit is set to 1, DA verifcation is disabled.
Tis feld is mainly used in reduced SRH mode.
Reserved. Tis feld is flled with 0s before transmission
and ignored upon receipt.
Identifes the pre-shared key and algorithm used for
HMAC computation. If this feld is set to 0, the TLV
does not contain the HMAC feld.
HMAC computation result.
2.4.3 SRv6 Instruction Set: Endpoint Node Behaviors
Afer describing SRv6 instructions and SRHs in previous sections, let’s
move onto the details of these instructions. Te IETF draf SRv6 Network
Programming[1] defnes multiple behaviors or instructions. Each SID is
bound to an instruction to specify the action to be taken during SID pro-
cessing. An SRH can encapsulate an ordered list of SIDs, providing packet-
related services including forwarding, encapsulation, and decapsulation.
SRv6 instructions defne the atomic functions of networks. Before
introducing SRv6 instructions, let’s learn the naming rules of these
instructions.
• End: the most basic instruction executed by an endpoint node,
directing the node to terminate the current instruction and start the
next instruction. Te corresponding forwarding behavior is to dec-
rement the SL feld by 1 and copy the SID pointed by the SL feld to
the DA feld in the IPv6 header.
• X: forwards packets through one or a group of Layer 3 outbound
interfaces.
• T: searches a specifed routing table and forwards packets.
• D: decapsulates packets by removing the IPv6 header and related
extension headers.
• V: searches a specifed table for packet forwarding based on Virtual
Local Area Network (VLAN) information.
88.
SRv6 Fundamentals ◾41
• U: searches a specifed table for packet forwarding based on unicast
MAC address information.
• M: searches a Layer 2 forwarding table for multicast forwarding.
• B6: binding to an SRv6 Policy.
• BM: binding to an SR-MPLS Policy.
Table 2.4 describes the functions of common SRv6 instructions, each of
which integrates one or more atomic functions previously listed.
TABLE 2.4 Functions of Common SRv6 Instructions
Instruction Function Description
End Terminates the current
instruction and executes the
next one.
End.X Forwards a packet through a
specifed outbound interface.
End.T Searches a specifed IPv6 routing
table for packet forwarding.
End.DX6 Decapsulates a packet and
forwards it over a specifed IPv6
Layer 3 adjacency.
End.DX4 Decapsulates a packet and
forwards it over a specifed IPv4
Layer 3 adjacency.
End.DT6 Decapsulates a packet and
searches a specifed IPv6 routing
table for packet forwarding.
End.DT4 Decapsulates a packet and
searches a specifed IPv4 routing
table for packet forwarding.
End.DT46 Decapsulates a packet and
searches a specifed IPv4 or IPv6
routing table for packet
forwarding.
Application Scenario
Used for packet forwarding
through a specifed node. An
End SID is similar to an
SR-MPLS node SID.
Used for packet forwarding
through a specifed outbound
