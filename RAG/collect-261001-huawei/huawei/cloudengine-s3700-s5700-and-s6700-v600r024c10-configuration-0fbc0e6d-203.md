---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-203
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [29443, 29585]
sha256: afb9c9329bf98c60993df9e347721f997f74c8b413d70218938187a1a66c7867
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

4.28 Configuring Static BFD for TE Tunnel
4.28.1 Understanding Static BFD for TE Tunnel
Definition
                 BFD is a bidirectional detection mechanism mainly used to detect communication
                 faults between forwarding engines. It can detect the connectivity of a data
                 protocol run on the same path between two systems. The path can be either a
                 physical or logical link, for example, a TE tunnel. BFD for TE tunnel uses BFD to
                 detect the entire TE tunnel and triggers VPN FRR to rapidly switch traffic if the
                 primary path fails, minimizing adverse impacts on services.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      487
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


Context
                 On an MPLS VPN network utilizing MPLS TE for encapsulating and transmitting
                 VPN data packets across the public network, service restoration following a
                 primary MPLS TE tunnel failure depends solely on E2E route convergence and LSP
                 convergence. The duration of service convergence is significantly influenced by the
                 volume of routes within the MPLS VPN and the number of hops across the
                 transport network. Consequently, an increase in the number of VPN routes leads
                 to extended convergence time, resulting in service interruptions. To address this
                 issue, VPN FRR was introduced. This approach involves pre-configuring primary
                 and backup forwarding entries on the remote PE, which point to the primary and
                 backup PEs respectively. Additionally, BFD for TE tunnel is utilized to quickly
                 identify faults in the primary TE tunnel and switch traffic to the backup TE tunnel,
                 effectively mitigating the problem of prolonged E2E service convergence time
                 resulting from tunnel faults.


Fundamentals
                 As shown in Figure 4-52, in normal cases, CE1 accesses CE2 through the primary
                 tunnel. However, if PE2 fails, CE1 accesses CE2 through the backup tunnel.

                 If BFD for TE tunnel is not configured and PE2 encounters a fault, PE1 will detect
                 the fault, preferentially select the route advertised by PE3, and re-deliver
                 forwarding entries. The time it takes for the network to converge closely aligns
                 with the IGP convergence time. Before PE1 re-delivers the forwarding entry for the
                 route advertised by PE3, CE1 will experience a period of inability to access CE2.
                 This occurs because PE2, which is the endpoint of the outer TE tunnel indicated by
                 the forwarding entry, has failed. Consequently, E2E services are interrupted.

                 After BFD for TE tunnel is enabled, if PE2 experiences a failure, PE1 leverages BFD
                 for rapid detection of the primary TE tunnel's unavailability between PE1 and PE2.
                 Subsequently, PE1 updates the primary TE tunnel status table to reflect its
                 unavailability and communicates this change to the forwarding engine. This action
                 triggers VPN FRR, facilitating rapid traffic switching to the backup tunnel. The
                 fault convergence time is reduced to milliseconds.


                 Figure 4-52 VPN FRR networking




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                            488
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


Application Scenarios
                 On a VPN FRR network, an MPLS TE tunnel is established between PEs, and the
                 BFD for TE tunnel mechanism is used to detect faults in the TE tunnel. If BFD
                 detects a fault in the tunnel, a VPN FRR switchover is performed in milliseconds.


Differences Between BFD for CR-LSP and BFD for TE Tunnel
                 ●      Fault notification objects: BFD for CR-LSP notifies a TE tunnel of a fault and
                        triggers service traffic switching to the backup CR-LSP in the same TE tunnel.
                        However, BFD for TE tunnel notifies applications of faults. If both the primary
                        and backup CR-LSPs fail, BFD for TE tunnel can rapidly detect the tunnel fault
                        and notify applications, such as VPN, of the fault, triggering service traffic
                        switching to the backup TE tunnel interface.
                 ●      Supported sessions: BFD for CR-LSP supports both dynamic and static
                        sessions, but BFD for TE tunnel supports only static sessions.
                 ●      Application scenarios: BFD for CR-LSP applies to hot-standby CR-LSP
                        scenarios, but BFD for TE tunnel applies to VPN FRR scenarios.

4.28.2 Enabling BFD Globally

Prerequisites
                 Before configuring static BFD for TE tunnel, complete the following task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 Bidirectional forwarding detection (BFD) is an end-to-end fast fault detection
                 mechanism. It can be used in MPLS TE to rapidly detect faults in links that a
                 tunnel traverses and triggers path switchover, improving service reliability
                 networkwide. BFD for MPLS TE can be categorized into BFD for RSVP, BFD for TE
                 tunnel, and BFD for CR-LSP based on detection objects.

                 BFD for TE tunnel uses BFD to detect the entire TE tunnel and triggers VPN FRR to
                 rapidly switch traffic if the primary path fails, minimizing adverse impact on
                 services.

                 The difference between BFD for TE tunnel and BFD for CR-LSP lies in the objects
                 to which BFD reports faults. In BFD for TE tunnel, BFD notifies applications such as
                 VPN of faults and triggers traffic switchover between different TE tunnel
                 interfaces. However, in BFD for CR-LSP, BFD notifies TE tunnels of faults and
                 triggers traffic switchover between different CR-LSPs in the same TE tunnel.

                 On a VPN FRR network, a TE tunnel is established between PEs, and the BFD
                 mechanism is used to detect faults in the tunnel. If BFD detects a fault on the TE
                 tunnel, VPN FRR switchover is performed in milliseconds.

                 Before configuring static BFD for TE tunnel, enable BFD globally on the ingress
                 and egress of a tunnel. Perform the following configuration on the ingress and
                 egress of an MPLS TE tunnel.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           489
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable BFD globally.
                 bfd

                 ----End

4.28.3 Setting BFD Parameters on an Ingress

Context
                 BFD parameters that can be configured on a tunnel ingress include the local
                 discriminator, remote discriminator, local interval for sending BFD packets, local
                 interval for receiving BFD packets, and local BFD detection multiplier. These
                 parameters affect the establishment of a BFD session.

                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

