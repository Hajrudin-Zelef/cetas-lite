---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-48
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5896, 6041]
sha256: dd3c5c0a3a22e8d75d45a1a713eb71216bcba6539282c836fbfdeed4a2311177
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 3.5.1 Configuring an IPv4 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind each MCE interface and the PE interface connecting to the MCE to the
                        VPN instance, and configure IP addresses for the interfaces. For detailed
                        configurations, see 3.5.2 Binding an Interface to an IPv4 VPN Instance.

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support IS-IS.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP.

                    Deleting a VPN instance or disabling a VPN instance IPv4 address family will
                    delete all the IS-IS processes bound to the VPN instance and the VPN instance
                    IPv4 address family.

Procedure
                    ●   Configure the PE.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  93
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration


                        a.   Enter the system view.
                             system-view

                        b.   Create an IS-IS process, bind it to a VPN instance, and enter the IS-IS
                             view.
                             isis process-id vpn-instance vpn-instance-name

                             An IS-IS process can be bound to only one VPN instance.
                        c.   Set a network entity title (NET).
                             network-entity net-addr

                             A NET specifies the current IS-IS area address and the system ID.
                        d.   (Optional) Set an IS-IS level.
                             is-level { level-1 | level-1-2 | level-2 }

                             The default level is Level-1-2.
                        e.   Import BGP routes.
                             import-route bgp [ cost-type { external | internal } | cost cost | tag tag | route-policy route-
                             policy-name | [ level-1 | level-2 | level-1-2 ] ] *

                             If the IS-IS level is not specified in the command, BGP routes will be
                             imported into the Level-2 IS-IS routing table.
                        f.   Return to the system view.
                             quit

                        g.   Enter the view of the interface to be bound to the VPN instance.
                             interface interface-type interface-number

                        h.   Switch the interface working mode from Layer 2 to Layer 3.
                             undo portswitch

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command. Determine whether to perform this step based on
                             the current interface working mode.
                        i.   Enable IS-IS on the interface.
                             isis enable [ process-id ]

                        j.   Return to the system view.
                             quit

                        k.   Enter the BGP view.
                             bgp as-number

                        l.   Enter the BGP VPN instance IPv4 address family view.
                             ipv4-family vpn-instance vpn-instance-name

                        m. Import IS-IS routes into the routing table of the BGP VPN instance IPv4
                           address family.
                             import-route isis process-id [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.   Enter the system view.
                             system-view

                        b.   Create an IS-IS process, bind it to a VPN instance, and enter the IS-IS
                             view.
                             isis process-id vpn-instance vpn-instance-name


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                               94
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                              An IS-IS process can be bound to only one VPN instance. If an IS-IS
                              process is not bound to any VPN instance before it is started, this process
                              becomes a public network process and cannot be bound to a VPN
                              instance later.
                        c.    Set a NET.
                              network-entity net-addr

                              A NET specifies the current IS-IS area address and the system ID.
                        d.    (Optional) Set an IS-IS level.
                              is-level { level-1 | level-1-2 | level-2 }

                              The default level is Level-1-2.
                        e.    Return to the system view.
                              quit

                        f.    Enter the view of the interface to be bound to the VPN instance.
                              interface interface-type interface-number

                        g.    Switch the interface working mode from Layer 2 to Layer 3.
                              undo portswitch

                              Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                              S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                              be switched from Layer 2 mode to Layer 3 mode using the undo
                              portswitch command. Determine whether to perform this step based on
                              the current interface working mode.
                        h.    Enable IS-IS on the interface.
                              isis enable [ process-id ]

                    ----End

3.9.6 Configuring RIP Between an MCE and a PE

Prerequisites
                    Before configuring MCE, you have completed the following tasks:

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 3.5.1 Configuring an IPv4 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind each MCE interface and the PE interface connecting to the MCE to the
                        VPN instance, and configure IP addresses for the interfaces. For detailed
                        configurations, see 3.5.2 Binding an Interface to an IPv4 VPN Instance.


Context
                         NOTE

