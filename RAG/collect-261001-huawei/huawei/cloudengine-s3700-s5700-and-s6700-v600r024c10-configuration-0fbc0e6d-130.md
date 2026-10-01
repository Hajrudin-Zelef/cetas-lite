---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-130
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [18435, 18620]
sha256: b225f478a9c0eabfb7925f5d00adde830f12be3966fca1eb4dc285f0576253d7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Application Scenarios
                 When RSVP-TE is used to establish dynamic MPLS TE tunnels for carrying services,
                 you can deploy RSVP-TE GR helpers to improve neighbor reliability.

Benefits
                 This feature ensures uninterrupted data transmission when an active/standby
                 switchover occurs in the control plane of a failed node, offering device-level
                 reliability for MPLS TE nodes.

4.11.2 Configuring an RSVP-TE GR Helper
Prerequisites
                 Before configuring an RSVP-TE GR helper, you have completed the following tasks:
                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Configure RSVP-TE Hello extension on the GR node and GR helper.

Context
                 A device can only function as a GR helper to help a neighbor complete an RSVP-TE
                 GR. Therefore, this configuration needs to be performed only when you deploy
                 RSVP-TE GR on a neighbor that supports this feature. If all the neighbors of the

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          309
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


                 local device run the same version as the local device, you do not need to perform
                 this configuration.

                 Perform the following configuration on the GR helper.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Enable RSVP-TE.
                 mpls rsvp-te

         Step 4 Enable the device to support RSVP-TE GR.
                 mpls rsvp-te hello support-peer-gr

                 ----End

4.11.3 (Optional) Establishing a Hello Session Between RSVP-
TE GR Nodes

Context
                 If TE FRR is deployed, a Hello session needs to be established between the PLR
                 and MP. This configuration is required when the neighboring node does not
                 support the function of automatically establishing a Hello session.

                 Perform the following configuration on the PLR and MP of the bypass tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Establish a Hello session with an RSVP-TE neighboring node.
                 mpls rsvp-te hello nodeid-session ip-address

                 The ip-address parameter indicates the LSR ID of the RSVP-TE neighboring node.

                 ----End

4.11.4 Verifying the Configuration

Procedure
                 ●      Run the display mpls rsvp-te graceful-restart command to check the local
                        RSVP-TE GR status.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         310
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 ●      Run the display mpls rsvp-te graceful-restart peer [ { interface interface-
                        type interface-number | node-id } [ ip-address ] ] command to check the
                        RSVP-TE GR status of the neighbor.
                 ----End


4.12 Configuring Dynamic BFD for RSVP

4.12.1 Understanding BFD for RSVP
Definition
                 Bidirectional forwarding detection (BFD) is an end-to-end fast fault detection
                 mechanism. It can be used in MPLS TE to rapidly detect faults in links that a
                 tunnel traverses, thereby triggering path switchovers and improving networkwide
                 reliability. BFD for MPLS TE can be categorized into BFD for RSVP, BFD for TE
                 tunnel, and BFD for CR-LSP based on detection objects. This section specifically
                 covers BFD for RSVP.
                 BFD for RSVP uses BFD to monitor RSVP neighbor relationships, implementing
                 millisecond-level fault detection. It can also work with RSVP to rapidly detect RSVP
                 neighbor faults.

Context
                 When a Layer 2 device is deployed on a link between two RSVP nodes, an RSVP
                 node can only use the Hello mechanism to detect a link fault. For example, on the
                 network shown in Figure 4-21, a Layer 2 device (Switch A) exists between P1 and
                 P2. If the link between Switch A and P2 fails, P1 cannot quickly detect the fault
                 through the link layer because Switch A isolates the fault. Instead, P1 can detect
                 the fault only through the RSVP Hello mechanism within seconds. This causes a
                 large amount of data to be lost. In this case, you can configure BFD for RSVP for
                 faster fault detection, thereby triggering TE FRR switching and improving network
                 reliability.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          311
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


                 Figure 4-21 Network diagram of BFD for RSVP




Implementation
                 Unlike BFD for CR-LSP and BFD for TE tunnel that support multi-hop BFD sessions,
                 BFD for RSVP only establishes single-hop BFD sessions between RSVP neighbors to
                 monitor the IP layer.

                 BFD for RSVP, BFD for OSPF, BFD for IS-IS, and BFD for BGP can share a BFD
                 session. When protocol-specific BFD parameters are set for a BFD session shared
                 by RSVP and other protocols, the smallest values take effect. The parameters
                 include the minimum intervals at which BFD packets are sent, minimum intervals
                 at which BFD packets are received, and local detection multipliers.

Application Scenarios
                 BFD for RSVP applies to a TE FRR scenario where a Layer 2 device exists between
                 a PLR and an RSVP neighbor on the primary CR-LSP.

Benefits
                 This feature improves the reliability of an MPLS TE network with Layer 2 devices.

4.12.2 Enabling BFD Globally

Prerequisites
                 Before configuring dynamic BFD for RSVP, you have completed the following task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 To configure dynamic BFD for RSVP, you need to enable BFD globally on both
                 RSVP neighbors.

Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                         312
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Perform the following configuration on the two RSVP neighbors between which a
                 Layer 2 device exists.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable BFD globally.
                 bfd

                 ----End

4.12.3 Enabling BFD for RSVP
Context
                 You can enable BFD for RSVP in either of the following ways:
                 ●      Enabling BFD for RSVP globally: This mode is recommended if BFD for RSVP
                        needs to be enabled for most RSVP interfaces on the local node.
                 ●      Enabling BFD for RSVP on specified RSVP interfaces: This mode is
                        recommended if BFD for RSVP needs to be enabled for only a small number
                        of RSVP interfaces on the local node.
                 Select one way as required. Perform the following configuration on the two RSVP
                 neighbors between which a Layer 2 device exists.

