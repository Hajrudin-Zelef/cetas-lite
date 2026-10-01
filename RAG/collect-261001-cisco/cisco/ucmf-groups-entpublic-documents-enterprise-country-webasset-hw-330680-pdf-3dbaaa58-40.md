---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-40
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [5900, 6074]
sha256: f79e3be3dc4367a56c9c5774f1cd8ea0b8ba2ea03a0570ddcae43fc3f0841305
---

# A line starting with the # sign is comments.

Service QoS
Requireme
nts
Reliability Requirements Security Requirements
Busines
s VPN
l Bandwid
th needs
to be
reserved.
l Medium
QoS
requirem
ents
 
Networking Description
The IP MAN consists of backbone and access subnetworks. The IP MAN sends services to users.
Figure 3-42 shows E2E residential service networking. Figure 3-43 shows E2E enterprise
service networking.
Figure 3-42 Residential service networking
HSI
VOD
VOIP
IP/MPLS 
BackBoneMAN
SR
BRAS
UPE
SoftX
HSI
VOD/VOIP
MPLS TE+VLL/VPLS
PE-AGG
PE-AGG
MPLS TE+VLL/VPLS
DSLAM
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
120

Figure 3-43 Enterprise service networking
IP/MPLS 
BackBone
Enterprise 
service
MAN
SR
BRAS
UPE
L3VPN or L2VPN
MPLS TE
MPLS TE Hotstanby 
BFD for CR-LSP
Feature Deployment
Services shown in Figure 3-42 and Figure 3-43 are core services of carriers. These services
have specific bandwidth, QoS, and reliability requirements. To meet these requirements, MPLS
TE tunnels can be used over virtual private networks (VPNs). Table 3-28 lists deployment
solutions.
Table 3-28 MPLS TE over an IP MAN
Item L3VPN L2VPN
Service Business VPN l HSI
l VoD
l VoIP
Public
network
tunnels for
VPN services
MPLS TE tunnels MPLS TE tunnels
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
121

Item L3VPN L2VPN
Reliability l Network reliability:
– Link protection: provided
using TE Hot-standby and
bidirectional forwarding
detection (BFD) for
constraints-routed label
switched path (CR-LSP).
– Node protection: provided
using VPN fast reroute (FRR)
and BFD for TE tunnel.
l Device reliability: provided using
RSVP graceful restart (GR) or
non-stop routing (NSR).
l Network reliability:
– Link protection: provided
using TE Hot-standby and
BFD for CR-LSP.
– Node protection: provided
using virtual leased line
(VLL) FRR and BFD for TE
tunnel.
l Device reliability: provided using
RSVP GR or NSR.
QoS E2E QoS functions need to be deployed on a link between a user-end provider
edge device (UPE) and a broadband remote access server (BRAS) or a link
between a UPE and a service router (SR).
Security Message digest 5 (MD5) or keychain is used to authenticate RSVP messages.
 
Key deployment points are as follows:
l Explicit paths are configured to separately establish primary and backup CR-LSPs. The
two paths do not overlap in important areas.
3.3.2 DS-TE Applications
Application Scenario: Access of Different Services to One VPN
On VPNs with MPLS-TE tunnels, one VPN may transmit EF, AF, and BE services
simultaneously. One MPLS TE tunnel may transmit different types of services with different
QoS requirements.
To prevent services on one MPLS TE tunnel from affecting each other, set up specific VPNs
and TE tunnels to transmit specific services. This is because resources may be wasted when
multiple VPNs and tunnels are set up for transmitting difference types of services over a network
simultaneously.
Alternatively, you can deploy DS-TE and use a multi-CT LSP to transmit services over one
VPN. A multi-CT LSP can reserve up to eight CTs. Each CT can transmit one type of services
of one VPN. Services among different CTs does not affect each other.
As shown in Figure 3-44, VPN1 transmits EF, AF, and BE services. One DS-TE tunnel needs
to be set up and configured with CT0 (100 Mbit/s), CT2 (50 Mbit/s), and CT5 (10 Mbit/s). The
tunnel is bound to VPN1 on the ingress. After traffic of VPN1 is classified, the traffic enters
corresponding CT queues.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
122

Figure 3-44 One MPLS TE tunnel transmitting different services on one VPN
PE PE
CE CEVPN1
Site 1
VPN1
Site 2
CT0 for BE: 100Mbit/s 
CT2 for AF: 50Mbit/s
CT5 for EF: 10Mbit/s 
MPLS TE tunnel
 
Application Scenario: Access of Different Services to Different VPNs
On a VPN with MPLS-TE tunnels, multiple VPNs may share one TE tunnel. These VPNs require
specific QoS. They compete with each other for resources, and QoS requirements of services
over VPNs cannot be met.
The solutions to the preceding scenario are as follows:
l Multiple VPNs transmit different types of services.
One TE tunnel can be used to transmit a maximum of eight types of services.
For example, VPN1 and VPN2 can access the TE tunnel simultaneously. VPN1 transmits
EF and BE services and VPN2 transmits AF services. One TE tunnel needs to be set up
and each type of services on each VPN is configured with a specific CT. The number of
CTs is equal to the sum of service types on VPN1 and the number of service types on VPN2.
Three CTs are supported.
l Multiple VPNs transmit the same type of services.
The number of TE tunnels to be set up is equal to the number of VPNs. The number of CTs
on each tunnel is equal to the number of corresponding service types on VPNs.
For example, VPN1 and VPN2 can access the TE tunnel simultaneousl. VPN1 bears EF
and BE services and VPN2 also transmits EF and BE services. Two TE tunnels can be set
up for VPN1 and VPN2. Each type of services on each tunnel is configured with a specific
CT.
l Multiple VPNs transmit services (some services are the same).
Each VPN needs a tunnel. The number of CTs on each tunnel is equal to the number of
corresponding service types on VPNs.
Application Scenario: Access of Traffic to VPNs and Non-VPNs
QoS requirements vary with VPN traffic and non-VPN traffic. If one TE tunnel transmits all the
traffic, the VPN traffic and non-VPN traffic may compete with each other for resources, and
QoS requirements of services cannot be met.
The solutions to the preceding scenario are as follows:
l The VPN and non-VPN transmit different types of services.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
123

