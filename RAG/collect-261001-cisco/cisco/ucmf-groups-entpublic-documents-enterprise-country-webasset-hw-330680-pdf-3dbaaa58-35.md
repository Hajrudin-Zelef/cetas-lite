---
id: collect-261001-cisco/cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58-35
title: "A line starting with the # sign is comments."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-04-15"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-cisco/ucmf-groups-entpublic-documents-enterprise-country-webasset-hw-330680-pdf-3dbaaa58.md
source_anchor: ""
source_lines: [5014, 5171]
sha256: a61f579dd30ce767aa7d861bb8c7bc156cf86a43d7d96baa99cc80adf60c490e
---

# A line starting with the # sign is comments.

BFD for RSVP can share BFD sessions with BFD for Open Shortest Path First (OSPF), BFD
for Intermediate System to Intermediate System (IS-IS), or BFD for Border Gateway Protocol
(BGP). The local node selects the smallest values of parameters between the two ends of the
shared BFD session as local BFD parameters. The parameters include the interval at which BFD
packets are sent, interval at which BFD packets are received, and local detection multiplier.
BFD for CR-LSP
BFD monitors CR-LSPs. After BFD detects a fault in a CR-LSP, the BFD module immediately
instructs the forwarding plane to trigger a rapid traffic switchover. BFD for CR-LSP is used
together with a hot-standby CR-LSP or a tunnel protection group.
A BFD session is bound to a CR-LSP. This means that a BFD session is set up between the
ingress and egress. A BFD packet is sent by the ingress to the egress along a CR-LSP. After the
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
101

egress receives the packet, the egress responds to the BFD packet. The ingress can rapidly detect
the status of links through which the CR-LSP passes based on whether a reply packet is received.
If a link fault is detected, the BFD module notifies the forwarding plane of the fault. The
forwarding plane searches for a backup CR-LSP and switches traffic to the backup CR-LSP. In
addition, the forwarding plane reports the fault to the control plane. If dynamic BFD for CR-
LSP is used, the control plane proactively creates a BFD session to monitor the backup CR-LSP.
A static BFD session can also be used to monitor the backup CR-LSP.
Figure 3-30 Traffic forwarding of a BFD session before and after a traffic switchover
Primary CR-LSP
Backup CR-LSP
BFD session
LSRA LSRB LSRC
LSRD
LSRA LSRB LSRC
LSRD
Link fault
Before a link fault occurs
After a link fault occurs
 
BFD for TE Tunnel
BFD for TE tunnel uses BFD sessions to detect the entire TE tunnel to trigger traffic switchover
of the applications such as VPN FRR.
BFD for TE tunnel and BFD for CR-LSP send fault information to different objects. BFD for
TE tunnel notifies applications (VPN for example) of faults and triggers a traffic switchover
between different TE tunnel interfaces. BFD for CR-LSPs notifies a TE tunnel of faults and
triggers a traffic switchover between different CR-LSPs in the same TE tunnel.
Differences
Table 3-20 lists differences between BFD for RSVP, BFD for CR-LSP, and BFD for TE tunnel.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
102

Table 3-20 Differences between BFD for CR-LSP, BFD for RSVP, and BFD for TE tunnel
Detection
Technique
Detection
Object
Node Usage
Scenario
BFD Session
Support
BFD for RSVP RSVP neighbor
relationships
Two ends of an
RSVP session
Can be used
with TE FRR
Dynamic
BFD for CR-
LSP
CR-LSPs Ingress and
egress
Can be used
with a hot-
backup CR-LSP
l Dynamic
l Static
BFD for TE
tunnel
MPLS TE
tunnels
Ingress and
egress
Can be used
with VPN FRR
Static
 
3.2.9.9 RSVP GR
RSVP graceful restart (GR) ensures uninterrupted transmission on the forwarding plane while
an active main board (AMB)/standby main board (SMB) switchover is performed on the control
plane.
Background
GR applies to provider edge (PE) routers on the provider network shown in Figure 3-31. User
nodes access the provider network through only a single PE. MPLS TE tunnels are established
between PEs on the network to implement TE or transmit VPN traffic. If a PE fails or a
maintenance measure (such as a software upgrade) is taken, an AMB/SMB switchover is
performed on the PE. To prevent traffic loss during a traffic switchover, RSVP GR can be
implemented to ensure the uninterrupted transmission of critical services.
Figure 3-31 RSVP GR application scenario
PE1 PE2
PE3 PE4
CE1 CE2
CE3 CE4
VPNA
VPNB
VPNA
VPNB
 
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
103

Related Concepts
RSVP GR is a rapid status restoration mechanism for RSVP-TE that is implemented based on
non-stop forwarding (NSF).
In the RSVP GR process, two roles are defined based on their functions. GR restarter performs
a graceful restart; while GR helper assists the GR restarter in implementing a graceful restart.
RSVP GR supports the following messages:
l Hello message carrying a GR extension: This message is used to detect the GR status of a
neighboring node.
l GR Path message: This message is sent by an upstream node and carries the contents of
the latest refreshed Path message.
l Recovery Path message: This message is sent by a downstream node and carries the contents
of the last Path message that is received by the downstream node.
Implementation
RSVP GR uses the Hello extension to monitor the GR status of neighboring nodes.
The principles of RSVP GR are as follows:
When the GR restarter performs GR on the network shown in Figure 3-32, it stops sending Hello
messages to its neighboring nodes. If a GR-enabled neighboring node (GR-Helper) fails to
receive consecutive three Hello messages, the neighbor node considers that the GR restarter is
performing GR and retains all forwarding information. In addition, an interface board on the
neighboring node continues transmitting services and waits for the GR restarter to restore the
GR status.
After the GR restarter is restarted and receives a Hello message from its neighboring node, it
replies with a Hello message. The processing modes of Hello messages are different on an
upstream node and a downstream node.
l If an upstream GR helper receives a Hello message, it sends a GR Path message to the GR
restarter.
l If a downstream GR helper receives a Hello message, it sends a Recovery Path message to
the GR restarter.
Figure 3-32 Schematic diagram of RSVP GR
GR-Helper GR-Restarter
Upstream Downstream
Hello Hello
GR Path Recovery 
Path GR-Helper
 
When the restarter receives GR Path and Recovery Path messages, the path state block (PSB)
on the GR restarter restores CR-LSP status information on the local control plane is restored.
If a downstream GR helper cannot send a Recovery Path message, the local PSB can restore
CR-LSP status information using only a GR Path message.
Enterprise Data Communication Products
Feature Description - MPLS 3 MPLS TE
Issue 04 (2013-04-15) Huawei Proprietary and Confidential
Copyright © Huawei Technologies Co., Ltd.
104

