---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-43
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5206, 5344]
sha256: 2b1ddaa6f6f862e69caa61a0f2e49039759924c8a2407ddeef0eac529b219150
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        81
VPN Configuration
VPN Configuration                                                                         3 IPv4 L3VPN Configuration


                    ●   Import direct routes from the VPN instance into the public network instance's
                        routing table for direct routes.
                        ip import-rib vpn-instance vpn-instance-name protocol direct [ route-policy route-policy-name ]

                    ●   Import static routes from the VPN instance into the public network instance's
                        routing table for static routes.
                        ip import-rib vpn-instance vpn-instance-name protocol static [ valid-route ] [ route-policy route-
                        policy-name ]

                    ●   Import OSPF routes from the VPN instance into the public network instance's
                        routing table for OSPF routes.
                        ip import-rib vpn-instance vpn-instance-name protocol ospf process-id [ valid-route ] [ route-
                        policy route-policy-name ]

                    ●   Import IS-IS routes from the VPN instance into the public network instance's
                        routing table for IS-IS routes.
                        ip import-rib vpn-instance vpn-instance-name protocol isis process-id [ valid-route ] [ route-policy
                        route-policy-name ]

                    ●   Import BGP routes from the VPN instance into the public network instance's
                        routing table for BGP routes.
                        bgp as-number
                        ipv4-family unicast
                        import-rib vpn-instance vpn-instance-name [ valid-route ] [ route-policy route-policy-name ]

                    ----End


Verifying the Configuration
                    ●   Run the display ip routing-table vpn-instance vpn-instance-name command
                        to check IPv4 routes imported into a specified VPN instance.
                    ●   Run the display ip routing-table command to check IPv4 routes on the
                        public network.

3.7.3 Importing the Public Network Instance's Routes to the
Routing Tables of an IPv4 VPN Instance

Prerequisites
                    Before configuring a device to import routes from a public network instance into
                    the routing tables of an IPv4 VPN instance, you have completed the following
                    tasks:

                    ●   3.5.1 Configuring an IPv4 VPN Instance on a PE
                    ●   3.5.2 Binding an Interface to an IPv4 VPN Instance


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Import different types of routes from the public network instance into the VPN
                instance's routing tables.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               82
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration


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
                    ●   Import direct routes from the public network instance into the VPN instance's
                        routing table for direct routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib public protocol direct [ route-policy route-policy-name ]

                    ●   Import static routes from the public network instance into the VPN instance's
                        routing table for static routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib public protocol static [ valid-route ] [ route-policy route-policy-name ]

                    ●   Import OSPF routes from the public network instance into the VPN instance's
                        routing table for OSPF routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib public protocol ospf process-id [ valid-route ] [ route-policy route-policy-name ]

                    ●   Import IS-IS routes from the public network instance into the VPN instance's
                        routing table for IS-IS routes.
                        ip vpn-instance vpn-instance-name
                        ipv4-family
                        import-rib public protocol isis process-id [ valid-route ] [ route-policy route-policy-name ]

                    ●   Import BGP routes from the public network instance into the VPN instance's
                        routing table for BGP routes.
                        bgp as-number
                        ipv4-family vpn-instance vpn-instance-name
                        import-rib public [ valid-route ] [ route-policy route-policy-name ]

                    ----End

Verifying the Configuration
                    ●   Run the display ip routing-table vpn-instance vpn-instance-name command
                        to check IPv4 routes imported into a specified VPN instance.
                    ●   Run the display ip routing-table command to check IPv4 routes on the
                        public network.

3.7.4 Configuring an IPv4 VPN Instance to Import Routes from
Other IPv4 VPN Instances

