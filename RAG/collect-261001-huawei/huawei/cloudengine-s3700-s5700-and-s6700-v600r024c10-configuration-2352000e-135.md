---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-135
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [19178, 19313]
sha256: a535ecd32b702f52b7ee743e32538ff457c9f3cf434f8ebbbd9f71f5cdc1c2fc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support the route-
                        policy route-policy-name parameter.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support OSPFv3.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support IS-ISv6 and BGP4+.

                         NOTE

                        If you do not want a VPN instance to change the next hops of routes imported from the
                        public network instance or other VPN instances when advertising these routes to its IBGP
                        peers, run the import-rib route next-hop-invariable command for the VPN instance.
                    ●   Import direct routes from the VPN instance into the public network instance's
                        routing table for direct routes.
                        ipv6 import-rib vpn-instance vpn-instance-name protocol direct [ route-policy route-policy-name ]

                    ●   Import static routes from the VPN instance into the public network instance's
                        routing table for static routes.
                        ipv6 import-rib vpn-instance vpn-instance-name protocol static [ valid-route ] [ route-policy
                        route-policy-name ]
                    ●   Import OSPFv3 routes from the VPN instance into the public network
                        instance's routing table for OSPFv3 routes.
                        ipv6 import-rib vpn-instance vpn-instance-name protocol ospfv3 process-id [ valid-route ] [ route-
                        policy route-policy-name ]

                    ●   Import IS-ISv6 routes from the VPN instance into the public network
                        instance's routing table for IS-ISv6 routes.
                        ipv6 import-rib vpn-instance vpn-instance-name protocol isis process-id [ valid-route ] [ route-
                        policy route-policy-name ]

                    ●   Import BGP routes from the VPN instance into the public network instance's
                        routing table for BGP routes.
                        bgp as-number
                        ipv6-family unicast
                        import-rib vpn-instance vpn-instance-name [ valid-route ] [ route-policy route-policy-name ]

                    ----End

Verifying the Configuration
                    ●   Run the display ipv6 routing-table vpn-instance vpn-instance-name
                        command to check information about IPv6 routes imported into a specified
                        VPN instance.
                    ●   Run the display ipv6 routing-table command to check public network IPv6
                        routes.

4.7.3 Configuring Route Import from the Public Network
Instance to an IPv6 VPN Instance's Routing Tables




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                304
VPN Configuration
VPN Configuration                                                                           4 IPv6 L3VPN Configuration


Prerequisites
                    Before configuring route import from the public network instance to an IPv6 VPN
                    instance's routing tables, you have completed the following task:

                    ●   4.5.1 Configuring an IPv6 VPN Instance on a PE
                    ●   4.5.2 Binding an Interface to the IPv6 VPN Instance


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Import different types of routes from the public network instance into the VPN
                instance's routing tables.
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support the route-
                        policy route-policy-name parameter.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support OSPFv3.
                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support IS-ISv6 and BGP4+.
                    ●   Import direct routes from the public network instance into the VPN instance's
                        routing table for direct routes.
                        ip vpn-instance vpn-instance-name
                        ipv6-family
                        import-rib public protocol direct [ route-policy route-policy-name ]

                    ●   Import static routes from the public network instance into the VPN instance's
                        routing table for static routes.
                        ip vpn-instance vpn-instance-name
                        ipv6-family
                        import-rib public protocol static [ valid-route ] [ route-policy route-policy-name ]

                    ●   Import OSPFv3 routes from the public network instance into the VPN
                        instance's routing table for OSPFv3 routes.
                        ip vpn-instance vpn-instance-name
                        ipv6-family
                        import-rib public protocol ospfv3 process-id [ valid-route ] [ route-policy route-policy-name ]

                    ●   Import IS-ISv6 routes from the public network instance into the VPN
                        instance's routing table for IS-ISv6 routes.
                        ip vpn-instance vpn-instance-name
                        ipv6-family
                        import-rib public protocol isis process-id [ valid-route ] [ route-policy route-policy-name ]

                    ●   Import BGP routes from the public network instance into the VPN instance's
                        routing table for BGP routes.
                        bgp as-number
                        ipv6-family vpn-instance vpn-instance-name
                        import-rib public [ valid-route ] [ route-policy route-policy-name ]

                    ----End


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               305
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


Verifying the Configuration
                    ●   Run the display ipv6 routing-table vpn-instance vpn-instance-name
                        command to check information about IPv6 routes imported into a specified
                        VPN instance.
                    ●   Run the display ipv6 routing-table command to check public network IPv6
                        routes.

4.7.4 Configuring Route Import from One IPv6 VPN Instance
into Another IPv6 VPN Instance

