---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-21
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2754, 2892]
sha256: 5fdaa923cbaafc1e90a66751bbeb9f524de6012d561471a9d4be189b572647ad
---

# A line starting with the # sign is comments.

NOTE
l A static CR-LSP is manually established, and there is no need to use the information advertisement or
the path calculation.
l A dynamic CR-LSP is dynamically established by signaling. Therefore, all the preceding functions are
used to establish a dynamic CR-LSP.
When deploying MPLS TE on a network, you need to configure link attributes and tunnel
attributes to automatically establish an MPLS TE tunnel. After a tunnel is established, traffic
has to be imported to it and forwarded through it.
3.2.3 Information Advertisement
MPLS TE uses routing protocols to advertise resource allocation information about each node
on a network. Each node on an MPLS TE network especially the ingress node of a tunnel
determines the nodes through which a tunnel passes based on advertised information.
Contents to Be Advertised
The network resource information to be advertised includes the following items:
l Link status information: interface IP addresses, link types, and link metric values, which
are collected by an Interior Gateway Protocol (IGP).
l Bandwidth information, such as maximum link bandwidth and maximum reservable
bandwidth.
l TE metric: TE link metric, which is the same as the IGP metric by default.
l Link administrative group: link color.
l Affinity Attributes: color of the link required by TE.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
56

l SRLG: shared risk link group. It is a constraint for path calculation of a backup path. SRLG
helps prevent backup and primary paths from overlapping over links with the same risk
level.
Advertisement Methods
MPLS TE advertises information using extended link-state-based routing protocols, including
OSPF TE and IS-IS TE. Open Shortest Path First (OSPF) TE and Intermediate System to
Intermediate System (IS-IS) TE automatically collect TE information and flood it to MPLS TE
nodes.
OSPF TE
Open Shortest Path First (OSPF) is a link-state-based routing protocol and features strong
scalability. OSPF defines label state advertisements (LSAs) of Type 1 to Type 5, and Type 7 to
carry the routing information of the intra-area, inter-area, and autonomous system external (AS-
external) for route calculation. These LSAs have fixed formats and cannot meet MPLS TE
requirements. Therefore, Opaque LSA and TE LSA are introduced.
l Opaque LSA
The Opaque LSA consists of three types of LSAs, that is, Type 9, Type 10, and Type 11.
Type 9 Opaque LSA can be flooded only on an interface; Type 10 Opaque LSA can be
flooded only within an area; Type 11 Opaque LSA that is similar to Type 5 LSA, can be
flooded within the entire AS outside the stub area and the not-so-stubby area (NSSA).
The Opaque LSA has the similar header format as that of other types of LSAs. The
difference is that the 4-byte Link State ID field is divided into Opaque Type and Opaque
ID, as shown in Figure 3-9.
Figure 3-9 Format of the Opaque LSA
0 23 15 31
LS age Options LS type=9, 10, 11
Advertising Router
LS sequence number
LS checksum Length
Opaque Type Opaque ID
Opaque Information
7
The first byte is Opaque Type that is used to differentiate between application types of this
LSA; the other three bytes are Opaque ID that is used to differentiate LSAs of the same
application type. The Opaque LSA of the same type may have 255 types of applications,
and each application can have 16777216 LSAs in a flooding scope.
For example, the LSA applied in the OSPF graceful restart (GR) is Type 9 LSA with the
application type as 3; the LSA applied in TE extensions is Type 10 LSA with the application
type as 1.
The LSA carries information in the Opaque Information field. The information format is
defined by the application with different requirements. Usually, the extensible type-length-
value (TLV) format is used.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
57

Figure 3-10 TLV format
0 15 31
Type
Value
Length
– Type: indicates the message type.
– Length: defines the length of the Value field, in bytes.
– Value: indicates the message carried in TLV. It can be a TLV format. The nested TLV
is called sub-TLV.
l TE LSA Extensions
The LSA applied in TE extensions is called TE LSA. The TE LSA is Type 10 LSA, with
the application type as 1. Therefore, the TE LSA has the Link State ID in the format of
1.x.x.x; however, the flooding scope is restricted to an area.
Figure 3-11 shows the typical format of the TE LSA.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
58

Figure 3-11 TE LSA format
0 23 15 31
LS age
Opaque ID
Options LS type=10
Advertising Router
LS sequence number
LS checksum length=132
Link ID
External Route Tag
Local IP Address
Remote IP Address
TE Metric
Maximum Bandwidth
Maximum Reservable Bandwidth
TLV Type=1 TLV length=4
Router Address
TLV Type=2 TLV length=100
Sub-TLV Type=1 Sub-TLV length=1
Sub-TLV length=4NSub-TLV Type=3
Sub-TLV Type=4 Sub-TLV length=4N
Sub-TLV Type=5 Sub-TLV length=4
Sub-TLV Type=6 Sub-TLV length=4
Sub-TLV Type=7 Sub-TLV length=4
Opq Type=1
Sub-TLV Type=2 Sub-TLV length=4
PaddingLink Type=1
Sub-TLV Type=8 Sub-TLV length=32
Unreserved Bandwidth-Priority 0
Unreserved Bandwidth-Priority 1
...
Unreserved Bandwidth-Priority 7
Sub-TLV Type=9 Sub-TLV length=4
Administrative  Group  
The TE LSA uses the TLV format to carry the needed information. At present, two types
of TLVs are defined as follows:
– TLV Type 1
Router address TLV: uniquely identifies an MPLS node. In CSPF, this is known as the
router ID.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
59

