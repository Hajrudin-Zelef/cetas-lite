---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-139
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [19723, 19855]
sha256: 9b3143aa7f745ad8b9298e6d97e760bf4f78cfe0b6cbb3815a557744265e9eb6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                             Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                             S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                             be switched from Layer 2 mode to Layer 3 mode using the undo
                             portswitch command. Determine whether to perform this step based on
                             the current interface working mode.
                        i.   Enable IPv6 on the interface.
                             ipv6 enable
                        j.   Enable OSPFv3 on the interface.
                             ospfv3 process-id area area-id [ instance instance-id ]
                        k.   Return to the system view.
                             quit
                        l.   Enter the BGP view.
                             bgp as-number
                        m. Enter the BGP VPN instance IPv6 address family view.
                             ipv6-family vpn-instance vpn-instance-name
                        n.   Import OSPFv3 routes into the routing table for the BGP VPN instance
                             IPv6 address family.
                             import-route ospfv3 process-id [ med med-value | route-policy route-policy-name ] *
                    ●   Configure the MCE.
                        a.   Enter the system view.
                             system-view
                        b.   Create an OSPFv3 process, bind it to a VPN instance, and enter the
                             OSPFv3 view.
                             ospfv3 [ process-id ] vpn-instance vpnname

                             An OSPFv3 process can be bound to only one VPN instance.
                             A router ID needs to be specified when an OSPF process is started after it
                             is bound to a VPN instance. The router ID must be different from the
                             public network router ID configured in the system view. If the router ID is
                             not specified, OSPFv3 selects the IP address of one of the interfaces
                             bound to the VPN instance as the router ID based on a certain rule.
                        c.   Configure a router ID.
                             router-id router-id

                             A router ID uniquely identifies an OSPFv3 process in an AS. If no router ID
                             is configured, the OSPFv3 process cannot run.
                        d.   (Optional) Configure a domain ID for the OSPFv3 process.
                             domain-id domain-idvalue [ secondary ]

                             The domain ID can be either an integer or a dotted decimal number.
                             Generally, the routes that are imported from a PE are advertised as
                             External-LSAs. The routes that belong to different nodes of the same
                             OSPFv3 domain are advertised as Type 3 LSAs (intra-domain routes). This
                             requires that different nodes in the same OSPFv3 domain have the same
                             domain ID.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      313
VPN Configuration
VPN Configuration                                                                       4 IPv6 L3VPN Configuration


                        e.    Disable routing loop detection.
                              vpn-instance-capability simple

                              If OSPFv3 multi-instance is deployed on the MCE, the MCE receives LSAs
                              with the Down (DN) bit set to 1 from the PE. Because routing loop
                              detection has been enabled for VPN instances on the MCE, the MCE
                              cannot use these LSAs to calculate routes. In this case, run the vpn-
                              instance-capability simple command to disable OSPFv3 routing loop
                              detection. The MCE can then use LSAs to calculate OSPFv3 routes
                              without checking the DN bits and route tags carried in these LSAs.
                        f.    Enter the view of the interface bound to a VPN instance.
                              interface interface-type interface-number

                        g.    Switch the interface working mode from Layer 2 to Layer 3.
                              undo portswitch

                              Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                              S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                              be switched from Layer 2 mode to Layer 3 mode using the undo
                              portswitch command. Determine whether to perform this step based on
                              the current interface working mode.
                        h.    Enable IPv6 on the interface.
                              ipv6 enable

                        i.    Enable OSPFv3 on the interface.
                              ospfv3 process-id area area-id [ instance instance-id ]

                    ----End


Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ verbose ]
                    command on the MCE to check the IPv6 routing table of the VPN instance.

4.9.4 Configuring IS-ISv6 Between an MCE and a PE

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
                    Deleting a VPN instance or disabling the IPv6 address family of a VPN instance
                    will also delete the related IS-ISv6 processes.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   314
VPN Configuration
VPN Configuration                                                                           4 IPv6 L3VPN Configuration


                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support IS-ISv6.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP4+.


Procedure
                    ●   Configure the PE.
                        a.   Enter the system view.
                             system-view

                        b.   Create an IS-IS process, bind it to a VPN instance, and enter the IS-IS
                             view.
                             isis process-id vpn-instance vpn-instance-name

                             An IS-IS process can be bound to only one VPN instance.
                        c.   Set a network entity title (NET).
                             network-entity net-addr

