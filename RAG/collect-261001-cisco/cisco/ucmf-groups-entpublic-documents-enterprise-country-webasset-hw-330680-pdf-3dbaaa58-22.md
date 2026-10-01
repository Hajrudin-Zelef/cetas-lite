---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-22
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [2893, 3061]
sha256: 98ff1960e05a596bfb1822b20f8fe7962d6ecaafc892939dd19bfcd6fa97365f
---

# A line starting with the # sign is comments.

– TLV Type 2
Link TLV: carries the attributes of a link enabled with MPLS TE. Table 3-2 shows the
sub-TLVs that can be carried in the Link TLV.
Table 3-2 Sub-TLVs that can be carried in the Link TLV
Sub-TLV Description
Type1: Link Type (the length of
the Value field is 1 byte)
Indicates the link type.
l 1: indicates point-to-point links.
l 2: indicates multi-access links.
There is padding of three bytes after the value field
of the Type1 sub-TLV.
Type2: Link ID (the length of the
Value field is 4 bytes)
Indicates the link ID. It is in the format of an IP
address.
l For a point-to-point link, this field indicates the
OSPF router ID of the neighbor.
l For a multi-access link, this field indicates the
interface IP address of the designated router
(DR).
Type3: Local IP Address (the
length of the Value field is 4N
bytes)
Indicates the IP address of the local interface. It can
be IP addresses of several local interfaces. Each IP
address occupies 4 bytes.
Type4: Remote IP Address (the
length of the Value field is 4N
bytes)
Indicates the IP address of the remote interface. It
can be IP addresses of several remote interfaces.
Each IP address occupies 4 bytes.
l For a point-to-point link, this field is set as the
remote IP address.
l For a multi-access link, this field can be set as
0.0.0.0 or skipped.
Type5: Traffic Engineering
Metric (the length of the Value
field is 4 bytes)
Indicates the TE metric configured on a TE link.
The data format is ULONG.
Type6: Maximum Bandwidth (the
length of the Value field is 4 bytes)
Indicates the maximum bandwidth of a link. The
data format is 4 bytes in floating point.
Type7: Maximum Reservable
Bandwidth (the length of the
Value field is 4 bytes)
Indicates the maximum reservable bandwidth of a
link. The data format is 4 bytes in floating point.
Type8: Unreserved Bandwidth
(the length of the Value field is 32
bytes)
Indicates reservable bandwidth of eight priorities
of a link. Each priority is in the format of 4 bytes in
floating point.
Type9: Administrative Group (the
length of the Value field is 4 bytes)
Indicates the administrative group attribute.
 
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
60

If a link is identified as an MPLS TE link, OSPF runs on the link, and OSPF neighbors are
established; then the TE extensions to OSPF generates a corresponding TE LSA and
advertises it to the area according to this TE link. If other devices in the area also support
TE extensions, a network topology consisting of TE links is generated among these devices.
Each device that advertises the TE LSA must have a unique Router Address.
The Opaque LSA of Type 10 is advertised within the OSPF area. Therefore, CSPF
calculation is area-based. The inter-area LSPs need to be calculated in segments.
IS-IS TE
IS-IS is a routing protocol based on link status. It can be extended to advertise TE information.
In extended IS-IS, two TLVs are supported:
l Type 135: Wide Metric
IS-IS has two metrics:
– Narrow metric: 6 bits
– Wide metric: 32 bits. This TLV is not used in route calculation. Instead, it is for TE
information transmission only.
Narrow metric provides only 64 metric values. So, it cannot meet the requirements of large-
scale traffic engineering. Wide metric is introduced to transfer TE information.
During the transition from the narrow metric to the wide metric, IS-IS TE must support the
following metrics:
– Compatible: capable of receiving and sending the packets with the metric types being
narrow or wide
– Wide Compatible: capable of receiving the packets with the metric types being narrow
or wide and sending only the packets with the metric types being wide
l Type 22: IS reachability TLV
The IS reachability TLV is TLV type 22. Figure 3-12 shows the format of the IS
reachability TLV.
Figure 3-12 Format of the IS reachability TLV
15 0 23 31
System ID and pseudonode number ( 7 octets ) Link metric
 ( 3 octets )
Link metric ( continued ) sub-TLV length ( 1 octets )
sub-TLVs ( 0~244 octets)
The IS reachability TLV consisting of:
– System ID and pseudo node ID
– Default link metric
– Length of sub-TLVs
– Sub-TLVs with changeable length
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
61

Table 3-3 describes sub-TLVs of different types.
Table 3-3 Sub-TLVs in IS-IS TE
Sub-TLV Description
Type3: Administrative group (4
bytes)
Indicates the administrative group attribute. It has
four bytes in length. Each set bit corresponds to one
administrative group assigned to the interface.
Type6: IPv4 interface address (4N
bytes)
Indicates the IP address of the local interface. It can
be IP addresses of several local interfaces. Each IP
address occupies four bytes.
Type8: IPv4 neighbor address (4N
bytes)
Indicates the IP address of the remote interface. It
can be IP addresses of several remote interfaces.
Each IP address occupies four bytes.
l For a point-to-point link, this field is set as the
remote IP address.
l For a multi-access link, this field is set as
0.0.0.0.
Type9: Maximum link bandwidth
(4 bytes)
Indicates the maximum bandwidth of a link.
Type10: Reservable link
bandwidth (4 bytes)
Indicates the maximum reservable bandwidth of a
link.
Type11: Unreserved bandwidth
(32 bytes)
Indicates reservable bandwidth of eight priorities of
a link.
Type18: TE Default metric (3
bytes)
Indicates the TE metric configured on a TE link.
 
When to Advertise Information
OSPF TE or IS-IS TE floods link information so that each node can save area-wide link
information in a traffic engineering database (TEDB). Information flooding is triggered by the
establishment of an MPLS TE tunnel, or one of the following conditions:
l A specific IGP TE flooding interval elapses.
l A link is activated or deactivated.
l A CR-LSP in an MPLS TE tunnel fails to be established because of insufficient bandwidth.
l Link attributes, such as the administrative group attribute or affinity attribute change.
l The link bandwidth changes.
When the available bandwidth of an MPLS interface changes, the system automatically
updates information in the TEDB and floods it. When a lot of tunnels are to be established
on a node, the node reserves bandwidth and frequently updates information in the TEDB
and floods it. For example, the bandwidth of a link is 100 Mbit/s. If 100 TE tunnels, each
with bandwidth of 1 Mbit/s, are established, the system floods link information 100 times.
To help suppress the frequency at which TEDB information is updated and flooded, the
flooding is triggered based on either of the following conditions:
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
62

