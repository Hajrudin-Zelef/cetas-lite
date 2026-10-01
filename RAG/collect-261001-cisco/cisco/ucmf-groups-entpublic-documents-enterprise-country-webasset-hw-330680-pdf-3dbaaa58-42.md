---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-42
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [6253, 6401]
sha256: b30752c62bd954718f8e2b2c5dcf3e374e71f5e4407f12c1886a3527f681bad2
---

# A line starting with the # sign is comments.

– dPeerME: indicates a peer maintenance entity defect. A dPeerME defect is the
server-layer defect that occurs on a peer maintenance entity outside the MPLS
subnet. The defects of this type are reported by other network layers connected to
the MPLS subnet to MPLS OAM for handling.
l MPLS layer defects
– dLOCV: indicates the defect of connectivity verification loss.
A dLOCV defect occurs when no CV or FFD packet is received within three
consecutive intervals for sending CV or FFD packets.
– dTTSI_Mismatch: indicates the defect of TTSI mismatching.
A dTTSI_Mismatch defect occurs when no CV or FFD packet with a correct TTSI
is received within three consecutive intervals for sending CV or FFD packets.
– dTTSI_Mismerge: indicates the defect of TTSI mismerging.
A dTTSI_Mismerge defect occurs when CV or FFD packets with both correct and
incorrect TTSIs are received within three consecutive intervals for sending CV or
FFD packets.
– dExcess: indicates the defect of the excessive rate of receiving connectivity detection
packets.
A dExcess defect occurs when five or more correct CV or FFD packets are received
within three consecutive intervals for sending CV or FFD packets.
l Other defects
dUnknown: indicates an unknown defect in an MPLS network.
A defect can be defined as dUnknown. For example, if the egress detects a defect that
both CV packets and FFD packets are sent along the same LSP, this special defect that
is not defined in the protocol can be identified as a dUnknown defect.
4.2.2 Reverse Tunnel
When the basic OAM function is configured, the LSP to be detected needs to be bound to a
reverse tunnel for transmitting BDI packets.
BDI packets are transmitted through the reverse tunnel. A reverse tunnel can be an LSP with the
ingress and egress being opposite to those on the detected LSP. The reverse tunnel can also be
a non-MPLS path that connects the ingress and the egress of the detected LSP.
The reverse tunnel transmitting BDI packets can be one of the following types:
l Private reverse LSP
l Shared reverse LSP
l Non-MPLS reverse path
NOTE
Currently, Huawei only supports a TE tunnel as the reverse tunnel.
4.2.3 MPLS OAM Auto Protocol
MPLS OAM defined in ITU-T Recommendation Y.1710 and Y.1711 has the following
drawbacks:
l On an LSP, if the ingress is enabled with the OAM function later than the egress, or OAM
is enabled on the egress and disabled on the ingress, a dLOCV defect occurs.
Enterprise Data Communication Products
Feature Description - MPLS 4 MPLS OAM
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
129

l The dLOCV defect also occurs when OAM is disabled. You must disable OAM on the
ingress and egress before changing the type or updating the interval for sending detection
packets.
l The OAM parameters need to be set on the ingress and egress respectively. This, however,
may cause the detection packet type and the interval for sending detection packets on the
ingress to be different from those on the egress.
On the Huawei devices, the OAM auto protocol can address the preceding problems.
If the OAM auto protocol is enabled on the egress, the functions of first packet triggering OAM
and the dynamic enabling or disabling OAM are provided.
The MPLS OAM auto protocol is the patent of Huawei.
4.2.4 Protection Switching
Protection switching refers to that a protection tunnel (namely, the bypass tunnel) is pre-set for
the primary tunnel and assigned with bandwidth. The primary tunnel and the bypass tunnel form
a protection group. When the primary tunnel is faulty, data traffic can be quickly switched to
the bypass tunnel. This decreases the packet loss ratio or shortens the delay due to the LSP
failure, and enhances reliability of networks. Protection switching refers to the end-to-end
protection.
With MPLS OAM for fast fault detection, the protection switching can be performed in
milliseconds.
On the Huawei devices, protection switching can be performed in 1:1 mode or N:1 mode.
l In 1:1 mode, a primary tunnel and a bypass tunnel are set up between the ingress and egress.
– Normally, data is transmitted through the primary tunnel.
– When the ingress detects a fault on the primary tunnel through MPLS OAM, protection
switching is performed and the ingress switches data to the bypass tunnel for
transmission.
l In N:1 mode, a tunnel functions as the bypass tunnel for multiple primary tunnels. When
any primary tunnel fails, data is switched to the shared bypass tunnel. The N:1 mode is
used to save bandwidth in a network with the mesh topology.
Figure 4-2 Schematic diagram of the MPLS OAM tunnel protection
SwitchA SwitchB
Primary tunnel
Primary tunnel
Protection tunnel
Backward tunnel
Data flow when primary
tunnel is normal
Data flow when primary
tunnel is falled
 
Enterprise Data Communication Products
Feature Description - MPLS 4 MPLS OAM
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
130

4.3 References
The following table lists the references of this document.
Huawei implements MPLS OAM based on ITU-T recommendations; however, the Request for
Comments (RFCs) are only for reference.
Document Description Remarks
ITU-T
Recommendation
Y.1710
Requirements for Operation &
Maintenance functionality for
MPLS networks
Huawei implement MPLS OAM in
compliance with this
recommendation.
ITU-T
Recommendation
Y.1711
Operation & Maintenance
mechanism for MPLS networks
Huawei implement MPLS OAM in
compliance with this
recommendation.
ITU-T
Recommendation
Y.1720
Protection switching for MPLS
networks
Huawei implement MPLS OAM in
compliance with this
recommendation.
RFC 3429 Assignment of the 'OAM Alert
Label' for Multiprotocol Label
Switching Architecture (MPLS)
Operation and Maintenance
(OAM) Functions
This RFC is only reference for
Huawei to implement MPLS
OAM.
RFC 4377 Operations and Management
(OAM) Requirements for Multi-
Protocol Label Switched (MPLS)
Networks
This RFC is only reference for
Huawei to implement MPLS
OAM.
RFC 4378 A Framework for Multi-Protocol
Label Switching (MPLS)
Operations and Management
(OAM)
This RFC is only reference for
Huawei to implement MPLS
OAM.
Enterprise Data Communication Products
Feature Description - MPLS 4 MPLS OAM
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
131
