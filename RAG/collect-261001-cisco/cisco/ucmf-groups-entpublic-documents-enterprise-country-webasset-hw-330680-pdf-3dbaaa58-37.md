---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-37
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [5337, 5462]
sha256: 976f5483bc696bbd5fafcf4086796b5fa367f4ce7dcddbb9199f6c120e46862b
---

# A line starting with the # sign is comments.

IGP Extension
To support DS-TE. RFC 4124 uses a Bandwidth Constraints Sub-TLV into the Interior Gateway
Protocol (IGP) and redefines the Unreserved Bandwidth Sub-TLV. These Sub-TLVs are used
to collect and advertise information about the reservable bandwidth for each CT along a link.
For details, see RFC 4124.
Single-CT LSP and Multi-CT LSP
A single-CT LSP transmits traffic of only one CT.
A multi-CT LSP transmits traffic of multiple CTs.
For the multi-CT, the resource reservation, LSP establishment, or bandwidth preemption can be
successfully performed only when the bandwidth of all the CTs is sufficient.
LSP Preemption and TE-Class Mapping
If no path meets the bandwidth requirement of a desired CR-LSP, a device can tear down an
established CR-LSP and use the bandwidth assigned to that CR-LSP to establish a desired CR-
LSP. This process is called preemption. DS-TE uses setup and holding priorities to determine
whether to preempt resources.
DS-TE specifies a preemption priority for each CT. A TE-class is a combination of a CT and a
priority, which is described as follows:
TE-Class[n] = <CTi, priority>
i and the priority value range from 0 to 7, and n identifies a TE class.
The priority indicates the priority of CR-LSP preemption, and is not the value of the EXP field
in the MPLS packet header. The value of preemption priority ranges from 0 to 7. A smaller value
indicates a higher priority. A CR-LSP can be set up only when both the combination of its CT
and setup priority (<CT, setup-priority>) and the combination of its CT and holding priority
(<CT, hold-priority>) exist in the TE-class mapping table. For example, the TE-class mapping
table of a certain node contains only TE-class[0] = <CT0, 6> and TE-class[1] = <CT0, 7>.
Only the following types of CR-LSPs can be set up successfully:
l Class-Type = CT0, setup-priority = 6, hold-priority = 6
l Class-Type = CT0, setup-priority = 7, hold-priority = 6
l Class-Type = CT0, setup-priority = 7, hold-priority = 7
NOTE
The CR-LSPs of "Class-Type = CT0, setup-priority = 6, hold-priority = 7" cannot be configured. This is
because the setup priority of the CR-LSP cannot be higher than its holding priority.
Each of eight CTs can be combined with any of eight priorities, so there are 64 TE-classes. On
the device, eight TE-classes can be configured manually.
A TE-class mapping table consists of a set of TE-classes. You are advised to configure all the
LSRs with the same TE-class mapping table over an MPLS network. The device has the default
TE-class mapping table.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
109

Table 3-21 Default TE-class mapping table
TE-Class CT Priority
TE-Class[0] 0 0
TE-Class[1] 1 0
TE-Class[2] 2 0
TE-Class[3] 3 0
TE-Class[4] 0 7
TE-Class[5] 1 7
TE-Class[6] 2 7
TE-Class[7] 3 7
 
Bandwidth
MPLS DS-TE involves the following types of bandwidth:
l Total link bandwidth: is physical link bandwidth.
l Maximum reservable bandwidth: is the maximum bandwidth that a link can reserve for an
MPLS TE tunnel to be established. The maximum reservable bandwidth must be lower
than or equal to the total link bandwidth.
l CT bandwidth: is the bandwidth of service traffic of each type on each DS-TE tunnel.
l BC bandwidth: is the bandwidth reserved for all CTs along a link.
Figure 3-37 Relationship between bandwidths
CT0
CT0
CT1
CT1
Bandwidth of BC0
Bandwidth of BC1
Total link bandwidth
Maximum Reservable Bandwidth
Link through 
which a tunnel 
passes Tunnel 
B
Tunnel 
A
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
110

Bandwidth Constraints Model
Bandwidth constraint model defines the maximum number of bandwidth constraints and which
CTs each bandwidth constraint applies to and how to use BC bandwidth.
The IETF defines the following bandwidth constraints models:
l Russian Dolls Model (RDM): CTs can share bandwidth. The BC model ID of the RDM is
0.
The bandwidth of BC0 is less than or equal to the maximum reservable bandwidth of a link.
In Figure 3-38:
– Total bandwidth of all LSPs from CT0, CT1, ... CT7 ≤ Bandwidth of BC0 ≤ Maximum
reservable bandwidth
– Total bandwidth of all LSPs from CT1, CT2, and CT7 ≤ Bandwidth of BC1
– ...
– Total bandwidth of all LSPs from CT7 ≤ Bandwidth of BC7
Figure 3-38 RDM
……BC0 for 
CT0+CT1+
…+CT7
BC1 for 
CT1+CT2+
…+CT7 BC7 for
CT7
Maximum 
Reservable 
Bandwidth
For example, the bandwidth of a link is 100 Mbit/s, RDM is used, and three CTs are
supported, that is, CT0, CT1, and CT2. CT0, CT1, and CT2 transmit BE, AF, and EF traffic
respectively. The bandwidths of BC0, BC1, and BC2 are 100 Mbit/s, 50 Mbit/s, and 20
Mbit/s respectively. The total bandwidth of all LSPs transmitting EF traffic cannot be larger
than 20 Mbit/s; the total bandwidth of all LSPs transmitting AF and EF traffic cannot be
larger than 50 Mbit/s; the total bandwidth of all LSPs cannot be larger than 100 Mbit/s.
The RDM allows bandwidth preemption between CTs. If 0 ≤ m < n ≤ 7 and 0 ≤ i < j
≤ 7, the CTi of priority m can preempt the bandwidth of CTi of priority n and the bandwidth
of CTj of priority n. For example, CT0 with priority 3 can preempt the bandwidth of CT0
with priority 5 and the bandwidth of CT1 with priority 0. The total bandwidth of CTi of all
LSPs cannot exceed the bandwidth of BCi.
l Maximum Allocation Model (MAM): One BC is mapped to one CT, and CTs cannot share
bandwidth. The BC mode ID of the MAM is 1.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
111

