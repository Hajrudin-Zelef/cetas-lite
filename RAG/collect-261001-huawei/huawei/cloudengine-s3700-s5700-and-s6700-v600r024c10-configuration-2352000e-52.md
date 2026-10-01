---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-52
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [6569, 6712]
sha256: bc9288008e2793b8acae6a3f707999cd09b98cf397e45d644de2e6ec9cc0db08
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

3.10 Configuring IPv4 L3VPN FRR




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         103
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


3.10.1 Understanding IPv4 L3VPN FRR
Context
                    IPv4 L3VPN FRR applies to services that are very sensitive to packet loss and delay
                    on a VPN. It provides four FRR functions: VPN FRR, VPN IP FRR, VPNv4 FRR, and IP
                    +VPNv4 hybrid FRR.

Comparison Between Different Types of FRR

                    Table 3-3 Comparison between different types of FRR
                     FRR Type          Application         Protected Object          Description
                                       Scenario

                     VPN FRR           Applies to          Protects traffic          It is implemented
                                       scenarios in        transmitted over links    using BGP.
                                       which a PE is       between a PE and its      Remotely leaked
                                       connected to        connected PEs.            routes back up
                                       multiple PEs.                                 each other.

                     VPN IP FRR        Applies to          Protects traffic          It is implemented
                                       scenarios in        transmitted over links    using RM. Routes
                                       which multiple      between a PE and its      of different
                                       CEs are             connected CEs.            protocols can
                                       connected to                                  back up each
                                       the same PE.                                  other.

                     VPNv4 FRR         Applies to inter-   Protects the following    It is implemented
                                       AS VPN Option       objects in inter-AS VPN   using BGP. VPNv4
                                       B and H-VPN         Option B and H-VPN        routes back up
                                       scenarios.          scenarios:                each other.
                                                           ● In an inter-AS VPN
                                                             Option B scenario,
                                                             VPNv4 FRR protects
                                                             traffic transmitted
                                                             over links between
                                                             PEs and ASBRs.
                                                           ● In an H-VPN
                                                             scenario, VPNv4 FRR
                                                             protects traffic
                                                             transmitted over
                                                             links between SPEs
                                                             and PEs.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          104
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                     FRR Type           Application       Protected Object          Description
                                        Scenario

                     IP+VPNv4           Applies to        Protects traffic          ● IP FRR: applies
                     hybrid FRR         scenarios in      transmitted over links      to networks
                                        which a CE is     between a CE and its        where a non-
                                        connected to      connected PEs.              BGP protocol
                                        two PEs.                                      runs between
                                                                                      PEs and CEs.
                                                                                    ● Auto FRR:
                                                                                      applies to
                                                                                      networks
                                                                                      where BGP
                                                                                      runs between
                                                                                      PEs and CEs.


3.10.2 Configuring VPN FRR

Prerequisites
                    Before configuring VPN FRR, you have completed the following tasks:
                    ●   Configure a routing protocol on the devices for them to communicate.
                    ●   Ensure that two unequal-cost routes destined for the CE are available on the
                        PE.
                    ●   Create a VPN.

Context
                    If a CE is dual-homed to two PEs, you can configure VPN FRR to ensure that if the
                    primary link between PEs fails, VPN services are switched to the backup link.
                    VPN FRR applies to services that are very sensitive to packet loss and delay on a
                    VPN. On the network shown in 3.10.2 Configuring VPN FRR, CE1 is dual-homed
                    to PE2 and PE3, and VPN FRR is configured on PE1. When the link between PE1
                    and PE2 fails, VPN traffic needs to be fast switched to the link between PE1 and
                    PE3.
                    Two methods for configuring VPN FRR are available:
                    ●   Configure VPN FRR.
                    ●   Configure VPN BGP auto FRR.

                    Figure 3-15 VPN FRR networking




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         105
VPN Configuration
VPN Configuration                                                                 3 IPv4 L3VPN Configuration


Procedure
                    ●   Configuring VPN FRR
                               NOTE

                              Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S,
                              S6750-S, S5755-S, S5755E-H, and S5755-H series support this function.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the VPN instance view.
                              ip vpn-instance vpn-instance-name

                        c.    Enter the VPN instance IPv4 address family view.
                              ipv4-family

                        d.    Enable VPN FRR.
                              vpn frr

                        e.    Return to the VPN instance view.
                              quit

                        f.    Return to the system view.
                              quit

                        g.    (Optional) Enter the BGP view.
                              bgp as-number

                        h.    (Optional) Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        i.    (Optional) Configure delayed route selection.
                              route-select delay delay-value

