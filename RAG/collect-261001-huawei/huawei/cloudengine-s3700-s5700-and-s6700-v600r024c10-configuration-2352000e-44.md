---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-44
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5345, 5475]
sha256: e90fda8b39617217e44d287ed3995f0bdc42ce26f79a2d171efaf78d7322558a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Prerequisites
                    Before configuring an IPv4 VPN instance to import routes from other IPv4 VPN
                    instances, you have completed the following tasks:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                             83
VPN Configuration
VPN Configuration                                                                         3 IPv4 L3VPN Configuration


                    ●   3.5.1 Configuring an IPv4 VPN Instance on a PE
                    ●   3.5.2 Binding an Interface to an IPv4 VPN Instance

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Import different types of routes from another VPN instance into the VPN
                instance's routing tables.
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support the route-
                        policy route-policy-name parameter.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support OSPF.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support IS-IS and BGP.

                         NOTE

                        If you do not want a VPN instance to change the next hops of routes imported from the
                        public network instance or other VPN instances when advertising these routes to its IBGP
                        peers, run the import-rib route next-hop-invariable command for the VPN instance.
                    ●   Import direct routes from another VPN instance into the VPN instance's
                        routing table for direct routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib vpn-instance vpn-instance-name protocol direct [ route-policy route-policy-name ]
                    ●   Import static routes from another VPN instance into the VPN instance's
                        routing table for static routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib vpn-instance vpn-instance-name protocol static [ valid-route ] [ route-policy route-
                        policy-name ]
                    ●   Import OSPF routes from another VPN instance into the VPN instance's
                        routing table for OSPF routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib vpn-instance vpn-instance-name protocol ospf process-id [ valid-route ] [ route-policy
                        route-policy-name ]
                    ●   Import IS-IS routes from another VPN instance into the VPN instance's routing
                        table for IS-IS routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib vpn-instance vpn-instance-name protocol isis process-id [ valid-route ] [ route-policy
                        route-policy-name ]
                    ●   Import BGP routes from another VPN instance into the BGP routing table in
                        the BGP VPN instance IPv4 address family.
                        bgp as-number
                        ipv4-family vpn-instance vpn-instance-name
                        import-rib vpn-instance vpn-instance-name [ valid-route ] [ route-policy route-policy-name ]

                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                 84
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration


Verifying the Configuration
                    ●   Run the display ip routing-table vpn-instance vpn-instance-name command
                        to check IPv4 routes imported into a specified VPN instance.
                    ●   Run the display ip routing-table command to check IPv4 routes on the
                        public network.


3.8 Setting a Router ID for a BGP VPN Instance IPv4
Address Family

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP.

                    Because no router ID is configured for a BGP VPN instance IPv4 address family by
                    default, the BGP router ID is used. As a result, different BGP VPN instance IPv4
                    address families on the same device have the same router ID. However, some
                    cases require different router IDs. For example, they are required if BGP peer
                    relationships need to be established between different BGP VPN instance IPv4
                    address families on the same PE.
                    Two methods are available to configure a router ID for a BGP VPN instance IPv4
                    address family. Choose either of the following methods as required:
                    ●   Set a router ID for all BGP VPN instance IPv4 address families.
                    ●   Set a router ID for a specified BGP VPN instance IPv4 address family.
                    The router ID configured in the BGP VPN instance IPv4 address family view takes
                    precedence over that configured in the BGP view.


                        NOTICE

                    If a BGP session has been established in a BGP VPN instance IPv4 address family,
                    changing or deleting the configured router ID resets the BGP session.


Procedure
                    ●   Set a router ID for all BGP VPN instance IPv4 address families.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the BGP view.
                             bgp as-number
                        c.   Set a router ID for all BGP VPN instance IPv4 address families.
                             router-id vpn-instance auto-select

                             The router-id vpn-instance auto-select command takes precedence over
                             the router-id ipv4-address command in the BGP view.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   85
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


                                    NOTE

