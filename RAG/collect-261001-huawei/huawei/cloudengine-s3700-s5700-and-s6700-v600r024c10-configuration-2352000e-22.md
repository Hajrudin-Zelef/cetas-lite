---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-22
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [2255, 2420]
sha256: f9db19a25792a6bce4b8fc180522538edecf015db6ba7e92b80320073b1c64a9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of a tunnel interface.
                    interface tunnel interface-number

         Step 3 Enable the GRE traffic statistics collection function.
                    statistics enable

                    By default, the traffic statistics collection function is disabled on a tunnel interface.
         Step 4 Check traffic statistics on GRE tunnel interfaces.
                    display interface tunnel

                    ----End

2.5.3 Resetting Keepalive Message Statistics on Tunnel
Interfaces
Context
                    You can reset the traffic statistics of a tunnel interface by resetting statistics about
                    the keepalive messages and keepalive response messages sent and received by the
                    tunnel interface.

Procedure
                    ●    Reset the statistics of a specified tunnel interface in the user view.
                         reset interface counters tunnel [ interface-number ]

                    ●    Reset keepalive message statistics on a specified tunnel interface.
                         system-view
                         interface tunnel interface-number
                         reset keepalive packets count

                    ----End




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             31
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration




                                    3          IPv4 L3VPN Configuration


                    3.1 Overview of IPv4 L3VPN
                    3.2 Understanding IPv4 L3VPN
                    3.3 Configuration Precautions for IPv4 L3VPN
                    3.4 Default Settings for IPv4 L3VPN
                    3.5 Configuring Mutual Access Between Local IPv4 L3VPNs
                    3.6 Configuring Basic IPv4 L3VPN over MPLS
                    3.7 Importing Routes Between Instances
                    3.8 Setting a Router ID for a BGP VPN Instance IPv4 Address Family
                    3.9 Configuring the MCE Function
                    3.10 Configuring IPv4 L3VPN FRR
                    3.11 Configuring IPv4 L3VPN over MPLS Inter-AS Option A
                    3.12 Configuring IPv4 L3VPN over MPLS Inter-AS Option B (Basic Networking)
                    3.13 Configuring IPv4 L3VPN over MPLS Inter-AS Option B (ASBRs Also
                    Functioning as PEs)
                    3.14 Configuring IPv4 L3VPN over MPLS HVPN
                    3.15 Configuring IPv4 L3VPN over MPLS Hub-Spoke
                    3.16 Maintaining IPv4 L3VPN


3.1 Overview of IPv4 L3VPN
                         NOTE

                        Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755E-H,
                        S5755-H, and S5732-H-V2 series can use MPLS to forward VPN packets on backbone
                        networks of SPs.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 32
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration


Definition
                    Virtual private network (VPN) technology is used to establish private networks
                    over a public network. The reason why a VPN is called a virtual network is that a
                    connection between any two nodes within an entire VPN network does not require
                    an E2E physical link, which is required by a traditional private network. A VPN
                    network is built over a network platform provided by a public network service
                    provider.

Purpose
                    The expansion of the network scale boosts the internal interconnection
                    requirements of enterprises. Initially, Internet service providers (ISPs) provide
                    services for enterprises through leased lines. This mode has many shortcomings,
                    such as lengthy construction, high costs, management difficulties, and poor
                    security.
                    VPN is developed to solve the preceding problems. VPN has the following
                    characteristics:
                    ●   Private: It provides the same private network experience as a traditional
                        private network. Resources between the VPN and the underlying transport
                        network are independent, that is, VPN resources cannot be consumed by
                        users who do not belong to the VPN. In addition, the VPN security can be
                        fully guaranteed to protect internal information of the VPN from external
                        interference.
                    ●   Virtual: Users in a VPN communicate with each other over a public network,
                        which can also be used by non-VPN users. A VPN is only a logical private
                        network. A public network that carries a VPN is called a VPN backbone
                        network.

Benefits
                    VPN brings the following benefits:
                    ●   Reduces enterprise interconnection costs and improves enterprise
                        interconnection flexibility.
                    ●   Supports overlapping address spaces, which solves the problem of IP address
                        shortage and facilitates network planning.
                    ●   Isolates data forwarding between VPN instances, improving security.


3.2 Understanding IPv4 L3VPN

3.2.1 Basic Concepts of IPv4 L3VPN
Definition
                    As shown in Figure 3-1, a basic IPv4 L3VPN has the following characteristics:
                    ●   Transmits routes through BGP extensions.
                    ●   Encapsulates and transmits VPN packets over recursed public network
                        tunnels.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               33
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                    ●   A device cannot function as two or three of the P, PE, and CE roles at the
                        same time.

                    Figure 3-1 Basic IPv4 L3VPN networking




Related Concepts
                    ●   Site
                        Site is often mentioned when VPN is discussed. The meaning of site can be
                        understood from the following aspects:
                        A site means a group of IP systems with IP connectivity. Such IP connectivity
                        does not need to be achieved by the SP network.
                        As shown in the left part of Figure 3-2, Company X has its HQ network at
                        Site A in City A and its branch network at Site B in City B. IP devices within
                        each site can communicate without using the SP network.

                        Figure 3-2 Sites




                        Sites are divided based on topological connections between devices, rather
                        than the geographical locations, even though devices at a site are

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                34
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


