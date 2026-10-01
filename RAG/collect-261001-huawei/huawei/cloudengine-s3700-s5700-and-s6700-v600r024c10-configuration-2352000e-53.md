---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-53
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [6713, 6886]
sha256: 14c372f9957d298fa4fd8e76c508c0c8f576da1f5748e3539b0d3d17d640a6c7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                              Delayed route selection ensures that after the primary path recovers, the
                              device on the primary path performs route selection only after the
                              corresponding forwarding entries on the device are stable. This prevents
                              traffic loss during traffic switchback.
                    ●   Configure VPN BGP auto FRR.
                               NOTE

                              Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S,
                              S6750-S, S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and
                              S5735-S-V2 series support BGP.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        d.    Enable VPN BGP auto FRR.
                              auto-frr

                        e.    Configure delayed route selection.
                              route-select delay delay-value

                    ----End

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                             106
VPN Configuration
VPN Configuration                                                                  3 IPv4 L3VPN Configuration


Verifying the Configuration
                    Run the display ip routing-table vpn-instance vpn-instance-name [ ipv4-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IP route in the routing table.

3.10.3 Configuring VPN IP FRR

Context
                    IP FRR applies to networks where different protocols run between PEs and CEs. In
                    a VPN IP FRR scenario, if the next hop of the primary route along a link between
                    the PE and CEs becomes unreachable, traffic is immediately switched to another
                    link between the PE and CEs, minimizing IP service interruption.

                    VPN IP FRR applies to delay- and packet loss-sensitive IP services on a VPN. As
                    shown in Figure 3-16, the PE forwards data to the site vpn1 through Link_A, and
                    Link_B serves as a backup link. When the PE detects that the route to CE1 is
                    unreachable, it immediately switches traffic to Link_B and then performs
                    operations such as VPN route convergence to minimize the impact of the link
                    failure on VPN services.

                    Figure 3-16 VPN IP FRR networking




Procedure
                    ●   Configure IP FRR.
                              NOTE

                             This function is not supported by the S3710-H, S5735R-L-V2, S5735E-L-V2, S5735-L-V2,
                             and S5735I-L-V2.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VPN instance view.
                             ip vpn-instance vpn-instance-name

                        c.   Enter the VPN instance IPv4 address family view.
                             ipv4-family

                        d.   Enable VPN IP FRR.
                             ip frr


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  107
VPN Configuration
VPN Configuration                                                                 3 IPv4 L3VPN Configuration


                    ●   Configure VPN BGP auto FRR.
                               NOTE

                              Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S,
                              S6750-S, S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and
                              S5735-S-V2 series support BGP.
                        a.    Enter the system view.
                              system-view
                        b.    Enter the BGP view.
                              bgp as-number
                        c.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name
                        d.    Enable VPN BGP auto FRR.
                              auto-frr
                        e.    Configure delayed route selection.
                              route-select delay delay-value

                    ----End

Verifying the Configuration
                    Run the display ip routing-table vpn-instance vpn-instance-name [ ipv4-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IP route in the routing table.

3.10.4 Configuring VPNv4 FRR
Context
                    In an H-VPN scenario in which VPNv4 FRR is configured on an SPE, the SPE
                    immediately switches VPN services to the backup link if the primary link fails.

                    Figure 3-17 VPNv4 FRR networking




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                108
VPN Configuration
VPN Configuration                                                                  3 IPv4 L3VPN Configuration


Prerequisites
                    Before configuring VPNv4 FRR, you have completed the following tasks:
                    ●    Configure a routing protocol on the devices for them to communicate.
                    ●    Ensure that two unequal-cost routes destined for the PE are available on the
                         SPE.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4

         Step 4 Enable VPNv4 FRR.
                    auto-frr

         Step 5 Configure the device to select a VPNv4 route only when the next hop recurses to a
                tunnel. This configuration prevents packet loss during a revertive switchover.
                    bestroute nexthop-resolved tunnel

         Step 6 (Optional) Configure delayed route selection.
                    route-select delay delay-value

                    Delayed route selection ensures that after the primary path recovers, the device on
                    the primary path performs route selection only after the corresponding forwarding
                    entries on the device are stable. This prevents traffic loss during traffic switchback.

                    ----End

Verifying the Configuration
                    Run the display ip routing-table vpn-instance vpn-instance-name [ ipv4-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IP route in the routing table.

3.10.5 Configuring IP+VPNv4 Hybrid FRR
Context
                          NOTE

                        This function is not supported by the S3710-H, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and
                        S5735I-L-V2.

                    After IP+VPNv4 hybrid FRR is configured, if the next hop of the route from a PE to
                    a CE is unreachable, the PE can switch traffic to the link between its standby PE
                    and the CE for transmission.
                    If a local PE has learned a VPN route from a CE and learned a VPNv4 route with
                    the same prefix as the learned VPN route from a peer PE, you can configure IP
                    +VPNv4 hybrid FRR on the local PE. After IP+VPNv4 hybrid FRR is configured, the

