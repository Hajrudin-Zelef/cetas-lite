---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-138
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [19601, 19722]
sha256: 7f81de266d9e51276903fe0dae6d7619e79b88322cac985bb8b3435e24d6f0a8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind the MCE's interfaces and the PE's interface connected to the MCE to VPN
                        instances and configure IP addresses for these interfaces. For details, see 4.5.2
                        Binding an Interface to the IPv6 VPN Instance.

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP4+.


Procedure
                    ●   Configure the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Configure a static route for the IPv6 address family of a specified VPN
                              instance.
                              ipv6 route-static vpn-instance vpn-source-name dest-ipv6-address prefix-length interface-type
                              interface-number [ nexthop-ipv6-address ] [ preference preference | tag tag ] *
                        c.    Enter the BGP view.
                              bgp as-number

                        d.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        e.    Import the configured static route into the routing table for the BGP VPN
                              instance IPv6 address family.
                              import-route static [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.    Enter the system view.
                              system-view

                        b.    Configure a static route for the IPv6 address family of a specified VPN
                              instance.
                              ipv6 route-static vpn-instance vpn-source-name destination-ipv6-address prefix-length
                              interface-type interface-number [ nexthop-ipv6-address ] [ preference preference | tag tag ] *

                    ----End

Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ verbose ]
                    command on the MCE to check the IPv6 routing table of the VPN instance.

4.9.3 Configuring OSPFv3 Between an MCE and a PE

Prerequisites
                    Before configuring MCE, complete the following tasks:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               311
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration


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
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support OSPFv3.

                    Deleting a VPN instance or disabling the IPv6 address family of a VPN instance
                    will also delete the related OSPFv3 processes.

Procedure
                    ●   Configure the PE.
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
                             OSPFv3 domain are advertised as Type-3 LSAs (intra-domain routes). This
                             requires that different nodes in the same OSPFv3 domain have the same
                             domain ID.
                        e.   Import BGP routes.
                             import-route bgp [ cost cost | tag tag | type type | route-policy route-policy-name ]*
                        f.   Return to the system view.
                             quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         312
VPN Configuration
VPN Configuration                                                                       4 IPv6 L3VPN Configuration


                        g.   Enter the view of the interface bound to the VPN instance.
                             interface interface-type interface-number
                        h.   Switch the interface working mode from Layer 2 to Layer 3.
                             undo portswitch

