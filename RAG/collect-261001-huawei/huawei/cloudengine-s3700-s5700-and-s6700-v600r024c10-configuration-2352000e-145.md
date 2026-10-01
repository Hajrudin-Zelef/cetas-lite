---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-145
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20606, 20782]
sha256: a1c5f233b561d916a62831215a81a00c9ceb6fe4967691f088fbaafd4d7d2c5a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Configure a routing protocol on the devices for them to communicate.
                    ●   Ensure that two unequal-cost routes destined for the CE are available on the
                        PE.
                    ●   Create a VPN.


Procedure
         Step 1 Enter the system view.
                    system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         326
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration


         Step 2 Enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

         Step 3 Enter the VPN instance IPv6 address family view.
                    ipv6-family

         Step 4 Enable IPv6 VPN FRR.
                    vpn frr

         Step 5 Return to the VPN instance view.
                    quit

         Step 6 Return to the system view.
                    quit

         Step 7 (Optional) Enter the BGP view.
                    bgp as-number

         Step 8 (Optional) Enter the BGP VPN instance IPv6 address family view.
                    ipv6-family vpn-instance vpn-instance-name

         Step 9 (Optional) Configure delayed route selection.
                    route-select delay delay-value

                    Delayed route selection ensures that after the primary path recovers, the device on
                    the primary path performs route selection only after the corresponding forwarding
                    entries on the device are stable. This prevents traffic loss during traffic switchback.

                    ----End

Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ ipv6-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IPv6 route in the routing table.

4.10.2 Configuring VPN IPv6 FRR

Context
                    VPN IPv6 FRR applies to delay- and packet loss-sensitive IP services on a VPN.
                    After VPN IPv6 FRR is configured, if traffic from a PE to a CE fails to be forwarded,
                    the traffic is quickly switched to a link connected to another CE. This ensures non-
                    stop forwarding of IP services.

                    The device currently supports two VPN IPv6 FRR modes:

                    ●      IPv6 FRR: applies to networks where different protocols run between PEs and
                           CEs.
                    ●      VPN BGP auto FRR: applies to networks where BGP runs between PEs and
                           CEs.

Prerequisites
                    Before configuring VPN IPv6 FRR, you have completed the following tasks:

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          327
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


                    ●   Configure IPv6 L3VPN.
                    ●   Configure a PE to learn IPv6 VPN routes with the same prefix from different
                        CEs connected to it.


Procedure
                    ●   Configure IPv6 FRR.
                               NOTE

                              This function is not supported by the S3710-H, S5735R-L-V2, S5735E-L-V2, S5735-L-V2,
                              and S5735I-L-V2.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the VPN instance view.
                              ip vpn-instance vpn-instance-name

                        c.    Enter the VPN instance IPv6 address family view.
                              ipv6-family

                        d.    Enable VPN IPv6 FRR.
                              ipv6 frr

                    ●   Configure VPN BGP auto FRR.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        d.    Enable VPN BGP auto FRR.
                              auto-frr

                        e.    Configure a delay for route selection.
                              route-select delay delay-value

                    ----End


Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ ipv6-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IPv6 route in the routing table.

4.10.3 Configuring VPNv6 FRR

Context
                    VPNv6 FRR applies to scenarios where IPv6 services are sensitive to packet loss
                    and delay and the requirements for VPNs transmitting IPv6 services are high. If a
                    fault occurs on the backbone network after VPNv6 FRR is enabled, IPv6 VPN
                    services can be quickly switched to another link, ensuring service continuity.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   328
VPN Configuration
VPN Configuration                                                                  4 IPv6 L3VPN Configuration


Prerequisites
                    Before configuring VPNv6 FRR, you have completed the following tasks:
                    ●    Configure IPv6 L3VPN.
                    ●    Ensure that the PE can receive IPv6 VPN routes with the same prefix from
                         different VPNv6 peers.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Enter the BGP VPNv6 address family view.
                    ipv6-family vpnv6

         Step 4 Enable VPNv6 FRR.
                    auto-frr

         Step 5 (Optional) Configure delayed route selection.
                    route-select delay delay-value

                    Delayed route selection ensures that after the primary path recovers, the device on
                    the primary path performs route selection only after the corresponding forwarding
                    entries on the device are stable. This prevents traffic loss during traffic switchback.

                    ----End

Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ ipv6-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IPv6 route in the routing table.

4.10.4 Configuring IPv6+VPNv6 Hybrid FRR
Context
                          NOTE

                        This function is not supported by the S3710-H, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and
                        S5735I-L-V2.

