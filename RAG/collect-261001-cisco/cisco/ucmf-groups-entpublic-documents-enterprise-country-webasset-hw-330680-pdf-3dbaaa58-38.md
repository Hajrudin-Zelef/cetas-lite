---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-38
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [5463, 5653]
sha256: 6cae5849b743edc81a782ca3822d0680ba39fdfe747d883d4b8ce856bd40a821
---

# A line starting with the # sign is comments.

Figure 3-39 MAM
……
BC0 for CT0
BC1 for CT1
BC7 for 
CT7
Maximum 
Reservable 
Bandwidth
In the MAM, the total bandwidth of CTi along an LSP cannot be larger than that of BCi (0
≤ i ≤ 7). The total bandwidth of BCs cannot be larger than the maximum reservable
bandwidth.
For example, the bandwidth of a link is 100 Mbit/s, MAM is used, and three CTs are
supported, that is, CT0, CT1, and CT2. BC0 is 20 Mbit/s and transmits CT0 traffic (for
example, BE traffic); BC1 is 50 Mbit/s and transmits CT1 traffic (for example, AF traffic);
BC2 is 30 Mbit/s and transmits CT2 traffic (for example, EF traffic). The total bandwidth
of all LSPs transmitting BE traffic cannot be larger than 20 Mbit/s; the total bandwidth of
all LSPs transmitting AF traffic cannot be larger than 50 Mbit/s; the total bandwidth of all
LSPs transmitting EF traffic cannot be larger than 30 Mbit/s.
l Extended-MAM: Similar to the MAM, extended-MAM maps one BC to one CT and CTS
cannot share the bandwidth. The BC mode ID of the extended-MAM is 254.
The extended-MAM supports eight more implicit CTs (the combination of CT0 and eight
priorities). This is different from the MAM.
Extended-MAM redefines Unreserved Bandwidth Sub-TLV and Bandwidth Constraint
Sub-TLV advertised by an IGP. Unreserved Bandwidth Sub-TLV carries unreserved
bandwidth for eight TE-classes. Bandwidth Constraint Sub-TLV carries information about
the BC model and unreserved bandwidth for eight TE-classes so that the device supports
a maximum of 16 TE-classes.
When device A that has Extended-MAM configured functions as the transit or egress node,
device B configured with DS-TE in non-IETF mode functions as the ingress node and
creates a dynamic CR-LSP, and the TE-class (<CT0, priority>, 0 ≤ priority ≤ 7) of the
CR-LSP is not defined in the TE-class mapping table specified on device A, the CR-LSP
creation request is valid.
Table 3-22 lists the comparisons between the three bandwidth constraints models.
Table 3-22 Comparisons between the three bandwidth constraints models
Item RDM MAM/Extended-MAM
BC-CT mapping Maps one BC to one or more
CTs.
Maps one BC to one CT,
which is easy for bandwidth
management.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
112

Item RDM MAM/Extended-MAM
Bandwidth preemption Is unable to divide CT
bandwidth and requires
preemption to provide
sufficient bandwidth for CTs.
Divides CT bandwidth and
provides sufficient
bandwidth for CTs.
Bandwidth use efficiency Efficiently uses bandwidth. Wastes bandwidth.
 
3.2.10.3 Implementation
Basic Implementation
Edge nodes in the Diff-Serv model divide the traffic into several classes, and add class
information into the DSCP field in packets. The internal node selects a proper PHB for a packet
according to the DSCP field.
The EXP field in the MPLS packet header contains information relavant to the Diff-Serv model.
The key to implement DS-TE is how to map the DSCP field (with a maximum of 64 values) to
the EXP field (with a maximum of eight values). RFC 3270 defines the following solutions:
l Label-Only-Inferred-PSC LSP (L-LSP): The drop priority is specified in the EXP field and
the PHB is determined by the label value. During packet forwarding, the label determines
the packet forwarding path and allocates a PHB for the path.
l EXP-Inferred-PSC LSP (E-LSP): The PHB and the drop priority are specified in the EXP
field of the MPLS label. During packet forwarding, the label value determines the packet
forwarding path and the EXP value determines a PHB. The E-LSP is applicable to networks
that supports a maximum of eight PHBs.
The device implements the E-LSP. The device maps DSCP or EXP priorities to local priorities.
Table 3-23 lists the default mapping. DS-TE map local priorities to CTs and separately allocates
resources to each CT. Therefore, DS-TE LSPs are set up according to the CT. That is, the DS-
TE calculates the path and reserves resources based on the CT and its bandwidth.
Table 3-23 Default mapping between DSCP priorities, local priorities, and EXP priorities
DSCP Priority Local Priority EXP Priority
0-7 0 0
8-15 1 1
16-23 2 2
24-31 3 3
32-39 4 4
40-47 5 5
48-55 6 6
56-63 7 7
 
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
113

RSVP Extensions
IETF extends RSVP to implement DS-TE in IETF mode. RFC 4124 defines a CLASSTYPE
object for Path messages. The IETF draft (draft-minei-diffserv-te-multi-class-02) defines the
extended-classtype object that carries CT information about the E-LSP. For details, see RFC
4124 and draft-minei-diffserv-te-multi-class-02.
DS-TE LSP Setup
The DS-TE LSP setup process is similar to MPLS-TE LSP Setup. The difference is as follows:
l The RSVP Path message contains CT information.
l When receiving the RSVP Path message carrying CT information, the LSR check whether
<CT, priority> exists in the local TE-class mapping table and whether CT bandwidth is
sufficient. If <CT, priority> exists in the local TE-class mapping table and CT bandwidth
is sufficient, a new LSP can be set up.
l After the LSP is successfully set up, the LSR recalculates the reservable bandwidth for each
CT. Information about the reservable bandwidth is then sent to an IGP, and the IGP
advertises the information to other nodes over the network.
Figure 3-40 DS-TE LSP setup process
Path1 Path2 Path3
Resv5Resv6Ingress Transit Egress Resv4
Transit
Check whether <CT, priority> exists in the local TE-
Class table and whether CT bandwidth is sufficient
 
DS-TE Modes
The device provides the IETF mode and non-IETF mode:
l IETF mode: indicates the mode defined by the IETF. Eight CTs are combined with eight
priorities and the combinations specify 64 TE-classes. A maximum of eight TE-classes can
be configured on the device.
l Non-IETF mode: indicates the mode not defined by the IETF. Each of the two CTs is
combined with each of eight priorities, so 16 TE-classes are available.
Table 3-24 describes their differences.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
114

Table 3-24 Differences between the IETF mode and non-IETF mode
Item Non-IETF Mode IETF Mode
Bandwidth
constraints
model
Supports the MAM and
RDM.
Supports the RDM, MAM, and extended-
MAM.
CT Supports CT0 and CT1. Supports CT0 to CT7.
BC type Supports BC0 and BC1. Supports BC0 to BC7.
TE-class
mapping table
A TE-class mapping table can
be configured but cannot take
effect.
Supports the configuration and application
of the TE-class mapping table.
IGP message l The Unreserved
bandwidth Sub-TLV
carries the unreserved
bandwidth for eight TE-
classes corresponding to
CT0, in byte/s.
l The sub-TLV
(Unreserved Bandwidth
for Class-Type 1, type
0x8001) carries the
unreserved bandwidth for
eight TE-classes
corresponding to CT1, in
byte/s.
The CT information is carried in the sub-
TLVs.
The sub-TLVs are as follows:
l Unreserved Bandwidth Sub-TLV
– For RDM and MAM, it carries the
unreserved bandwidth for eight TE-
classes, in byte/s.
– For extended-MAM, it carries the
unreserved bandwidth for eight TE-
classes corresponding to CT0, in
byte/s.
l Bandwidth Constraints Sub-TLV
– For RDM and MAM, it carries
information about the BC model and
the BC bandwidth, in byte/s.
– For extended-MAM, it carries
information about the BC model and
unreserved bandwidth for eight TE-
classes, in byte/s.
RSVP
messages
The ADSPEC object carries
CT information.
Different objects carry CT information as
follows:
l Single CT: The CLASSTYPE object
carries CT information.
l Multi-CT: The EXTENDED_CLASS-
TYPE object carries CT information.
 
