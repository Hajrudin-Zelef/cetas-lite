---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-140
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [19856, 20010]
sha256: 7d6a2ddf94e144e3be5ee937602a1d0f23b8e2e81391429c6026fefe5315cee6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                             A NET specifies the current IS-IS area address and the system ID of the
                             device.
                        d.   (Optional) Set an IS-IS level.
                             is-level { level-1 | level-1-2 | level-2 }

                             The default level is Level-1-2.
                        e.   Enable IPv6 for the IS-IS process.
                             ipv6 enable [ topology { compatible [ enable-mt-spf ] | ipv6 | standard } ]

                             Before enabling IPv6 for the IS-IS process, enable IPv6 in the system view
                             first.
                        f.   Import BGP routes.
                             ipv6 import-route bgp inherit-cost [ { level-1 | level-2 | level-1-2 } | tag tag | route-policy
                             route-policy-name ] *

                             If the IS-IS level is not specified in the command, BGP routes will be
                             imported into the Level-2 IS-IS routing table.
                        g.   Return to the system view.
                             quit

                        h.   Enter the view of the interface bound to the VPN instance.
                             interface interface-type interface-number

                        i.   Switch the interface working mode from Layer 2 to Layer 3.
                             undo portswitch

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command. Determine whether to perform this step based on
                             the current interface working mode.
                        j.   Enable IPv6 on the interface.
                             ipv6 enable


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              315
VPN Configuration
VPN Configuration                                                                        4 IPv6 L3VPN Configuration


                        k.   Enable IS-ISv6 on the interface.
                             isis ipv6 enable [ process-id ]

                        l.   Return to the system view.
                             quit

                        m. Enter the BGP view.
                             bgp as-number

                        n.   Enter the BGP VPN instance IPv6 address family view.
                             ipv6-family vpn-instance vpn-instance-name

                        o.   Import IS-IS routes into the routing table for the BGP VPN instance IPv6
                             address family.
                             import-route isis process-id [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.   Enter the system view.
                             system-view

                        b.   Create an IS-IS process, bind it to a VPN instance, and enter the IS-IS
                             view.
                             isis process-id vpn-instance vpn-instance-name

                             An IS-IS process can be bound to only one VPN instance. If an IS-IS
                             process is not bound to any VPN instance before it is started, this process
                             becomes a public network process and cannot be bound to a VPN
                             instance later.
                        c.   Set a NET.
                             network-entity net-addr

                             A NET specifies the current IS-IS area address and the system ID of the
                             device.
                        d.   (Optional) Set an IS-IS level.
                             is-level { level-1 | level-1-2 | level-2 }

                             The default level is Level-1-2.
                        e.   Enable IPv6 for the IS-IS process.
                             ipv6 enable [ topology { compatible [ enable-mt-spf ] | ipv6 | standard } ]

                             Before enabling IPv6 for the IS-IS process, enable IPv6 in the system view
                             first.
                        f.   Return to the system view.
                             quit

                        g.   Enter the view of the interface bound to the VPN instance.
                             interface interface-type interface-number

                        h.   Switch the interface working mode from Layer 2 to Layer 3.
                             undo portswitch

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command. Determine whether to perform this step based on
                             the current interface working mode.
                        i.   Enable IPv6 on the interface.
                             ipv6 enable


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    316
VPN Configuration
VPN Configuration                                                                           4 IPv6 L3VPN Configuration


                        j.    Enable IS-ISv6 on the interface.
                              isis ipv6 enable [ process-id ]

                    ----End

Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ verbose ]
                    command on the MCE to check the IPv6 routing table of the VPN instance.

4.9.5 Configuring RIPng Between an MCE and a PE

Prerequisites
                    Before configuring MCE, complete the following tasks:

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 4.5.1 Configuring an IPv6 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind the MCE's interfaces and the PE's interface connected to the MCE to VPN
                        instances and configure IP addresses for these interfaces. For details, see 4.5.2
                        Binding an Interface to the IPv6 VPN Instance.

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support RIPng.

                    Deleting a VPN instance or disabling the IPv6 address family of a VPN instance
                    will also delete the related RIPng processes.

Procedure
                    ●   Configure the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Create a RIPng process, bind it to a VPN instance, and enter the RIPng
                              view.
                              ripng process-id vpn-instance vpn-instance-name

                              A RIPng process can be bound to only one VPN instance.
                        c.    Import BGP routes.
                              import-route bgp [ permit-ibgp ] [ cost cost | inherit-cost | route-policy route-policy-name ] *

                        d.    Return to the system view.
                              quit

