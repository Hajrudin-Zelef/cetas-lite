---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-39
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [5654, 5899]
sha256: eb4064fdf64e774b1c0376e87564cab7f34f8498532acc2887ce4e69fe281eac
---

# A line starting with the # sign is comments.

DS-TE Mode Switching
On the device, the non-IETF mode and the IETF mode can be switched to each other. DS-TE
mode switching is described in Table 3-25.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
115

Table 3-25 DS-TE mode switching
Item Non-IETF Mode → IETF Mode IETF Mode → Non-IETF Mode
Change in
the
Bandwidt
h
Constraint
s model
The bandwidth model is
unchanged.
The bandwidth models are changed as
follows:
l The extended-MAN is changed to
MAM.
l The RDM is unchanged.
l The MAM is unchanged.
Change in
the
bandwidth
The bandwidth values of BC0 and
BC1 are unchanged.
Other BC values are reset to zero except
values of BC0 and BC1.
Change in
the TE-
Class
mapping
table
If the TE-class mapping table is
configured, it is applied.
Otherwise, the default one is
applied.
For information about the default
TE-class mapping table, see Table
3-21.
The TE-class mapping table is not applied.
l If a TE-class mapping table is
configured, it is not deleted.
l If no TE-class mapping table is
configured, the default one is deleted.
LSP
deletion
LSPs whose <CT, set-priority> or
<CT, hold-priority> is not in the
TE-class mapping table are deleted
on the ingress node and the transit
node.
The following LSPs are deleted on the
ingress node and the transit node:
l Multi-CT LSPs
l LSPs of single CT from CT2 to CT7
 
DS-TE Scheduling
An ingress node marks local priorities of packets on the inbound interface based on priority
mapping or complex traffic classification. Each local priority is mapped to a CT. Packets arrive
at the outbound interface with local priorities. On the outbound interface, HQoS is used to
allocate bandwidth for DS-TE traffic, as shown in Figure 3-41.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
116

Figure 3-41 HQoS scheduling
CT0
CT1
CT2
CT0
CT1
CT2
  BC0-CTs
  BC1-CTs
  BC2-CTs
  BC7-CTs
LSP1
LSP2
Queue0
Queue1
Queue2
……
Queue7
Reserved 
bandwidth for TE
   Other traffic
Physical 
interface
  TE-Class CT0 0
  TE-Class CT1 0
  TE-Class CT2 0
DSCP0~7-LP0
EXP0-LP0
DSCP8~15-LP1
EXP1-LP1
DSCP16~23-LP2
EXP2-LP2
Traffic mapping
Bandwidth Constraint
l Reserved bandwidth for TE tunnels
Certain bandwidth is reserved for all TE tunnels on a physical interface by a separate level
of queuing. This prevents other types of traffic from occupying the reserved bandwidth.
l Bandwidth guarantee for CTs on TE tunnels
A TE tunnel supports multiple CTs to transmit services of different types. End-to-end
bandwidth is reserved for each CT when a CR-LSP is set up. The device allocate bandwidth
to each CT from the bandwidth reserved for TE tunnels when they set up a CR-LSP. The
allocated bandwidth conforms to the bandwidth constraints in the RDM or MAM model.
The guaranteed bandwidth for a CT is specified by the committed information rate (CIR).
l Traffic scheduling between CTs (total traffic of the same CT on different LSPs)
Each CT maps a service type. High-priority services (such as voice services) can be
scheduled using priority queuing (PQ). Services requiring bandwidth guarantee (such as
protocol and data services) can be scheduled using weighted fair queuing (WFQ). CTs are
mapped to local priorities in a one-to-one mode. You can associate CTs with queue profiles
and configure different scheduling modes in the queue profiles to provide differentiated
services for CTs.
The AR provides 32 queue profiles globally configure scheduling modes for CTs. The following
table lists the default mappings between CTs and fair queues (FQs), as shown in Table 3-26.
Table 3-26 CT scheduling
Local Priority CT FQ Configurable
Scheduling Mode
7 (CS7) CT7 7 WFQ
6 (CS6) CT6 6 WFQ
5 (EF) CT5 5 WFQ
4 (AF4) CT4 4 WFQ
3 (AF3) CT3 3 WFQ
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
117

Local Priority CT FQ Configurable
Scheduling Mode
2 (AF2) CT2 2 WFQ
1 (AF1) CT1 1 WFQ
0 (BE) CT0 0 WFQ
 
DS-TE Reliability
DS-TE provides the following reliability methods:
l TE FRR
DS-TE is applied as follows:
– When bandwidth protection is required, the CTs and bandwidth are configured manually
on the bypass CR-LSP in manually-configured FRR. The QoS is guaranteed. The
protection modes are the 1:1 and N:1. In the automatic FRR, the bypass CR-LSP inherits
the CTs and bandwidth of the primary CR-LSP. The QoS is guaranteed. The protection
mode is 1:1 only.
– When bandwidth protection is not required, both manually-configured and automatic
FRR support 1:1 protection and N:1 protection, irrespective of the CTs and their
bandwidth on the bypass tunnel.
l CR-LSP backup
The bypass CR-LSP inherits the CTs and their bandwidth from the primary CR-LSP. The
best-effort path cannot guarantee QoS and it does not inherit the CTs and bandwidth from
the primary CR-LSP.
Interworking Between Devices
During network deployment or device version upgrade, non-DS-TE devices may work with DS-
TE devices, or devices in non-IETF mode work with devices in IETF mode.
The device supports the interworking between the following devices:
l Interworking between DS-TE devices and non-DS-TE devices
– Supports the establishment of non-DS-TE tunnels from non-DS-TE devices to DS-TE
devices.
– Supports the establishment of non-DS-TE tunnels from DS-TE devices to non-DS-TE
devices.
l Interworking between non-Huawei DS-TE devices that do not support the CLASSTYPE
object
The device can parse the following Path messages with CT information sent by non-Huawei
devices:
– L-LSP CT information that is carried by the EXTENDED_CLASSTYPE object
– CTO information that is carried by the EXTENDED_CLASSTYPE object
3.3 Applications
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
118

3.3.1 MPLS TE Applications on an IP MAN
Service Overview
As technology advances, service bearing networks of carriers are becoming integrated and IP/
MPLS technology is of great importance after network integration. Voice, video, leased line,
and data services can be uniformly transmitted over the integrated IP/MPLS network. Services
transmitted over MANs are classified into the following two types based on user types:
l Residential services: include high speed Internet (HSI), video on demand (VoD), and voice
over IP (VoIP) services.
l Enterprise services: include VPN services for the headquarters and branches of large
enterprises. L3VPN services include Business VPNs and L2VPN services include data,
real-time video, and real-time voice services.
Table 3-27 lists services, their quality of service (QoS), reliability, and security requirements.
Table 3-27 IP MAN services
Service QoS
Requireme
nts
Reliability Requirements Security Requirements
HSI l Bandwid
th does
not need
to be
reserved.
l Low QoS
requirem
ents
l A standby link is established
for end-to-end (E2E)
services. If an active link
fails, the services can switch
to the standby link.
l Voice services must be
transmitted in real time and
be rapidly switched to a
standby link if an active link
fails.
l Different types of services
are separately transmitted.
l The IP bearer network
infrastructure can defend
against attacks and viruses
to ensure proper network
operation.
VoD l Bandwid
th needs
to be
reserved.
l Medium
QoS
requirem
ents
VoIP l Bandwid
th needs
to be
reserved.
l High
QoS
requirem
ents
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
119

