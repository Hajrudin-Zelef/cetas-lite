---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-41
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [6075, 6252]
sha256: 737d462be8b3f877ac46229822de4a9150f9c8af74a7ade0f181dead823f18ed
---

# A line starting with the # sign is comments.

One TE tunnel needs to be set up. Services on the VPN and non-VPN are configured with
different CTs. The number of CTs is equal to the sum of the number of service types on
the VPN and the number of service types on the non-VPN.
l The VPN and non-VPN transmit the same type of services.
Two TE tunnels need to be set up for the VPN and non-VPN services. Specific types of
services on each tunnel are configured with specific CTs.
l The VPN and non-VPN transmit services (some services are the same).
Two TE tunnels need to be set up for the VPN and non-VPN services. Specific types of
services on each tunnel are configured with specific CTs.
3.4 References
The following table lists the references of this document.
Document Description Rem
arks
RFC 2205 Resource Reservation Protocol -
RFC 2209 Resource Reservation Protocol (RSVP) - Version 1 Message
Processing Rules
-
RFC 2370 The OSPF Opaque LSA Option -
RFC 2547 BGP/MPLS VPNs -
RFC 2702 Requirements for Traffic Engineering Over MPLS -
RFC 2747 RSVP Cryptographic Authentication -
RFC 2961 RSVP Refresh Overhead Reduction Extensions -
RFC 3031 Multiprotocol Label Switching Architecture -
RFC 3032 MPLS Label Stack Encoding -
RFC 3034 Use of Label Switching on Frame Relay Networks Specification -
RFC 3209 RSVP-TE: Extensions to RSVP for LSP Tunnels -
RFC 3210 Applicability Statement for Extensions to RSVP for LSP-Tunnels -
RFC 3473 Generalized Multi-Protocol Label Switching (GMPLS) Signaling
Resource Reservation Protocol-Traffic Engineering (RSVP-TE)
Extensions
-
RFC 3630 Traffic Engineering (TE) Extensions to OSPF Version 2 -
RFC 3784 Intermediate System to Intermediate System (IS-IS) Extensions for
Traffic Engineering (TE)
-
RFC 4124 Protocol Extensions for Support of Diffserv-aware MPLS Traffic
Engineering
-
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
124

Document Description Rem
arks
RFC 4127 Russian Dolls Bandwidth Constraints Model for Diffserv-aware
MPLS Traffic Engineering
-
RFC 4128 Bandwidth Constraints Models for Differentiated Services
(Diffserv)-aware MPLS Traffic Engineering: Performance
Evaluation
-
RFC 4139 Requirements for Generalized MPLS (GMPLS) Signaling Usage
and Extensions for Automatically Switched Optical Network
(ASON)
-
RFC 4090 Fast Reroute Extensions to RSVP-TE for LSP Tunnels -
draft-ietf-
mpls-nodeid-
subobject-01
Definition of an RRO node-id subobject -
draft-ietf-
tewg-diff-te-
proto-02
Protocol extensions for support of Diff-Serv-aware MPLS Traffic
Engineering
-
draft-ietf-
mpls-diff-te-
reqts-00
Requirements for support of Diff-Serv-aware MPLS Traffic
Engineering
-
draft-ietf-
mpls-diff-
ext-07
MPLS Support of Differentiated Services -
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
125

4 MPLS OAM
About This Chapter
4.1 Introduction to MPLS OAM
4.2 Principles
4.3 References
Enterprise Data Communication Products
Feature Description - MPLS 4 MPLS OAM
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
126

4.1 Introduction to MPLS OAM
Definition
Operation, Administration and Maintenance (OAM) is an important means to cut costs in
network maintenance. The MPLS OAM mechanism manages operation and maintenance of
Multiprotocol Label Switching (MPLS) networks.
MPLS supports different Layer 2 and Layer 3 protocols such as IP, Frame Relay (FR), and
Asynchronous Transfer Mode (ATM). In an MPLS network, the OAM mechanism is provided
totally independent of any upper or lower layer, which implements the following features on the
MPLS user plane:
l Detects connectivity of label switched paths (LSPs).
l Assesses utilization and performance of an MPLS network.
l Performs protection switching when a defect or fault occurs on a link to provide services
in compliance with the signed service level agreements (SLAs) signed.
Purpose
As an extensible key technology of the next generation network, MPLS provides multiple
services guaranteed by quality of service (QoS). MPLS introduces a unique network layer that
may cause faults. Therefore, MPLS networks need to support OAM.
The protocols (such as Synchronous Optical Network (SONET)/Synchronous Digital Hierarchy
(SDH)) at the server layer below the MPLS network layer and the protocols (such as IP, FR, and
ATM) at the client layer above the MPLS network layer have their respective OAM mechanisms.
Failures of the MPLS network cannot be rectified thoroughly through the OAM mechanism of
other layers. In addition, the network technology hierarchy also requires MPLS to have its
independent OAM mechanism to decrease dependency between layers on each other.
The MPLS OAM mechanism can detect, identify, and locate a defect at the MPLS layer
effectively. Then, the MPLS OAM mechanism the reports and handles the defect. In addition,
when a failure occurs, the MPLS OAM mechanism can trigger protection switching.
4.2 Principles
4.2.1 MPLS OAM Detection
MPLS OAM packets can be classified into the following types:
l Connectivity detection packets
– Fast Failure Detection (FFD) packets
– Connectivity Verification (CV) packets
l Forward Defect Indication (FDI) packets
l Backward Defect Indication (BDI) packets
MPLS OAM monitors:
Enterprise Data Communication Products
Feature Description - MPLS 4 MPLS OAM
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
127

l TE LSP Connectivity
MPLS OAM periodically sends CV or FFD packets along a TE LSP.
MPLS OAM for a TE LSP
Figure 4-1 MPLS OAM for a TE LSP
CV/FFD
CV/FFD
BDIBDI
Ingress
lSR
Engress
lSR
 
As shown in Figure 4-1, MPLS OAM works as follows:
1. The ingress sends a CV packet or an FFD packet along an LSP to be detected. The packet
passes through the LSP and arrives at the egress.
2. The egress compares the packet type, interval, and Trail Termination Source Identifier
(TTSI) in the received packet with the local values to check the correctness of the packet.
In addition, the egress counts the number of received correct and incorrect packets within
a detection cycle. In this manner, MPLS OAM detects connectivity of the LSP.
3. The detection interval of CV packet is a fixed value, and the detection cycle of FFD packet
is three times the detection interval.
4. When the egress detects an LSP defect, the egress analyzes the defect type and sends a BDI
packet carrying the defect information to the ingress through a reverse tunnel. In this
manner, the ingress is notified of the defect of a specific type in time. If the protection group
is configured correctly, protection switching is triggered.
The detected defect is of one of the following types:
l Non-MPLS layer defects
– dServer: indicates a server-layer defect. A dServer defect is the server-layer defect
that occurs below an MPLS network. The defects of this type are reported by the
server layer to MPLS OAM for handling.
The lower layer network that bears MPLS services may have its own protection and
defect detection mechanism. When a lower-layer defect occurs on an LSP, a
downstream label switch router (LSR) that is closest to the defect can notify the
egress of the defect. The lower-layer defect should not trigger the switchover but be
only notified to the network management device. In addition, the lower-layer defect
can be notified to the ingress through a proper method (of sending BDI packets).
Enterprise Data Communication Products
Feature Description - MPLS 4 MPLS OAM
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
128

