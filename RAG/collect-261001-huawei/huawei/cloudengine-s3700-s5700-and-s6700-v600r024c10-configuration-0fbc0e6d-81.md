---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-81
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11415, 11549]
sha256: f62bb373766186a09f9a9c896c0514b051c6b8c041db2fe698a0f1d353fa10f5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      Run the display this command in the MPLS-LDP view and check whether
                        information similar to the following is displayed:
                        propagate mapping for ip-prefix abc (The value varies according to the actual situation.)
                        If so, check whether some LSPs are not included in the IP-prefix-based policy
                        abc.
                 ●      Run the display ip ip-prefix and check whether information similar to the
                        following is displayed:
                            index: 10           permit 1.1.1.1/32
                            index: 20           permit 2.2.2.2/32 (The value varies according to the actual situation.)
                        If so, LSPs can be established only for the routes 1.1.1.1/32 and 2.2.2.2/32.
                 ●      If the preceding policies are configured, add the routing information
                        corresponding to the LSP to the policies.
                 ●      If the preceding policies are not configured, go to Step 5.
         Step 5 Collect the following information and contact technical support personnel:
                 ●      Results of the preceding steps
                 ●      Configuration file, logs, and alarms of the device

                 ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               193
MPLS Configuration
MPLS Configuration                                                         4 MPLS TE Configuration




                                       4         MPLS TE Configuration


                 4.1 Overview of MPLS TE
                 4.2 Understanding MPLS TE
                 4.3 Configuration Precautions for MPLS TE
                 4.4 Default Settings for MPLS TE
                 4.5 Configuring Static MPLS TE Tunnels
                 4.6 Configuring Static Bidirectional Associated LSPs
                 4.7 Configuring Dynamic MPLS TE Tunnels
                 4.8 Steering Traffic to an MPLS TE Tunnel
                 4.9 Adjusting RSVP-TE Signaling Parameters
                 4.10 Configuring RSVP-TE Authentication
                 4.11 Configuring an RSVP-TE GR Helper
                 4.12 Configuring Dynamic BFD for RSVP
                 4.13 Configuring the MPLS TE Bandwidth Flooding Threshold
                 4.14 Configuring a Metric for MPLS TE Tunnel Path Selection
                 4.15 Configuring MPLS TE Tunnel Priorities
                 4.16 Configuring the Link Administrative Group and Affinity Attributes of MPLS TE
                 4.17 Configuring MPLS TE Explicit Paths
                 4.18 Configuring a Hop Limit for an MPLS TE Tunnel
                 4.19 Configuring an MPLS TE SRLG
                 4.20 Configuring MPLS TE Tunnel Re-optimization
                 4.21 Configuring Delayed Switching and Deletion in MPLS TE
                 4.22 Configuring Synchronization Between CR-LSP Establishment and Overload
                 Status
                 4.23 Configuring CR-LSP Backup

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                         194
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 4.24 Configuring Static BFD for CR-LSP
                 4.25 Configuring Dynamic BFD for CR-LSP
                 4.26 Configuring Manual MPLS TE FRR
                 4.27 Configuring MPLS TE Auto FRR
                 4.28 Configuring Static BFD for TE Tunnel
                 4.29 Configuring Synchronization Between TE Tunnel Status and BFD Session
                 Status
                 4.30 Configuring Synchronization Between TE FRR and CR-LSP Backup
                 4.31 Verifying TE Tunnels Using Ping and Tracert
                 4.32 Maintaining MPLS TE
                 4.33 Troubleshooting MPLS TE


4.1 Overview of MPLS TE
Definition
                 Multiprotocol Label Switching (MPLS) traffic engineering (MPLS TE) manages and
                 controls network traffic based on the MPLS technology. MPLS TE establishes
                 constraint-based routed label switched path (CR-LSPs) to plan paths for network
                 traffic and optimize network resource usage. MPLS TE provides various reliability
                 technologies to ensure end-to-end bandwidth and quality of service (QoS)
                 guarantee.

Purpose
                 On a traditional IP network, a node typically selects the shortest path as the
                 optimal route, without considering other factors such as bandwidth availability.
                 Such a routing mechanism can lead to congestion on the shortest path while
                 leaving resources underutilized on other available paths. Traffic engineering can
                 address congestion caused by uneven resource allocation by directing some traffic
                 onto less utilized or idle links, thereby optimizing the overall resource utilization.
                 Before the advent of MPLS TE, the following TE mechanisms were utilized:
                 ●      IP TE: adjusts path metrics to control traffic transmission paths. This
                        mechanism helps avoid congestion on certain links, though it might shift
                        congestion to other links. Adjusting a metric on a complex network is
                        challenging because altering one link can impact multiple routes.
                 ●      Asynchronous Transfer Mode (ATM) TE: Traditional Interior Gateway Protocols
                        (IGPs) choose routes solely based on connectivity, lacking the capability to
                        distribute traffic according to the bandwidth and specific attributes of links.
                        The IP over ATM overlay model can address this limitation. This overlay model
                        establishes virtual links for steering certain traffic, ensuring more effective
                        traffic distribution and better QoS guarantees. However, ATM TE incurs
                        significant additional costs and complicates network expansion efforts.
                 A scalable and simple solution is necessary for implementing TE on a large-scale
                 backbone network. As an overlay model, MPLS can establish a virtual topology

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            195
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                 over a physical topology, directing traffic through this virtual framework.
                 Therefore, MPLS TE stands as an ideal solution.


Benefits
                 As a TE solution, MPLS TE offers the following advantages:
                 ●      Maximizes network resource utilization, offering bandwidth and QoS
                        guarantees without requiring hardware upgrades. This approach significantly
                        reduces enterprise costs.
                 ●      Offers ease of deployment and maintenance on live networks, leveraging
                        existing MPLS technologies.
                 ●      Delivers a range of reliability features to achieve carrier- and device-class
                        reliability.
                 ●      Supports L2VPN and L3VPN services, ensuring high security and reliable QoS
                        for VPN services.


4.2 Understanding MPLS TE

4.2.1 Basic Concepts of MPLS TE

