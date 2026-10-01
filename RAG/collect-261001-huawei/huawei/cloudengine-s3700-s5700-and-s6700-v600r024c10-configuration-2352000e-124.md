---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-124
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [17499, 17674]
sha256: 4f5c096794ce743046476065b591046bf0db6620fd5d02b2aed8a3a7b28a3140
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 100 200
                          #
                          ip vpn-instance vpna
                           ipv6-family
                            route-distinguisher 100:1
                            vpn-target 111:1 export-extcommunity
                            vpn-target 111:1 import-extcommunity
                            vpn-target 222:2 import-extcommunity
                          #
                          ip vpn-instance vpnb
                           ipv6-family
                            route-distinguisher 100:2
                            vpn-target 222:2 export-extcommunity
                            vpn-target 222:2 import-extcommunity
                            vpn-target 111:1 import-extcommunity
                          #
                          interface Vlanif100
                           ip binding vpn-instance vpna
                           ipv6 enable
                           ipv6 address 2001:DB8::11:2/112
                          #
                          interface Vlanif200
                           ip binding vpn-instance vpnb
                           ipv6 enable
                           ipv6 address 2001:DB8::12:2/112
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 100
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           port trunk allow-pass vlan 100 200
                          #
                          bgp 100
                           #
                           ipv6-family unicast
                           #
                           ipv6-family vpn-instance vpna
                            import-route direct
                           #
                           ipv6-family vpn-instance vpnb
                            import-route direct
                          #
                          return
                    ●     CE1
                          #
                          sysname CE1
                          #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         278
VPN Configuration
VPN Configuration                                                              4 IPv6 L3VPN Configuration

                         vlan batch 100
                         #
                         interface Vlanif100
                          ipv6 enable
                          ipv6 address 2001:DB8::11:1/112
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100
                         #
                         ipv6 route-static 2001:DB8::12:1 112 2001:DB8::11:2
                         #
                         return

                    ●    CE2
                         #
                         sysname CE2
                         #
                         vlan batch 100
                         #
                         interface Vlanif100
                          ipv6 enable
                          ipv6 address 2001:DB8::12:1/112
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100
                         #
                         ipv6 route-static 2001:DB8::11:1 112 2001:DB8::12:2
                         #
                         return



4.6 Configuring Basic IPv6 L3VPN over MPLS Functions

4.6.1 Configuring an IPv6 VPN Instance on a PE

Prerequisites
                    Before configuring an IPv6 VPN instance, you have completed the following task:

                    ●    Configure link layer protocol parameters for interfaces to ensure that these
                         interfaces work properly.


Context
                    A VPN instance is also called a VRF or a per-site forwarding table.

                    VPN instances are used to isolate VPN routes from public network routes. Routes
                    of different VPN instances are isolated from one another.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         279
VPN Configuration
VPN Configuration                                                                              4 IPv6 L3VPN Configuration


                          NOTE

                         The name of a VPN instance is case-sensitive. For example, vpn1 and VPN1 are considered
                         different VPN instances.

         Step 3 (Optional) Configure a description for the VPN instance.
                    description description-information

         Step 4 Enable the VPN instance IPv6 address family, and enter the VPN instance IPv6
                address family view.
                    ipv6-family

                    Configurations in a VPN instance can be performed only after an address family is
                    enabled for the VPN instance based on the advertised route and type of forwarded
                    data.
         Step 5 Configure an RD for the VPN instance IPv6 address family.
                    route-distinguisher route-distinguisher

                    A VPN instance IPv6 address family takes effect only after having an RD
                    configured. The RDs of different VPN instance IPv6 address families on a PE must
                    be different.

                          NOTE

                         If you perform this step in the view of a newly created VPN instance, the VPN instance IPv6
                         address family is automatically enabled and the VPN instance IPv6 address family view is
                         automatically displayed.
                         The S3710-H series products do not support dynamic routing protocols. Therefore, this
                         command does not take effect after being delivered.

         Step 6 Configure VPN targets for the VPN instance IPv6 address family.
                    vpn-target vpn-target &<1-8> [ both | export-extcommunity | import-extcommunity ]

                    VPN targets are a type of BGP extended community attribute used to control the
                    import and export of VPN routes. You can configure a maximum of eight import
                    VPN targets and eight export VPN targets each time the vpn-target command is
                    run. Run this command multiple times if you want to configure more VPN targets
                    in the VPN instance IPv6 address family.
                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
         Step 7 Set the maximum number of route prefixes allowed for the VPN instance IPv6
                address family.
                    prefix limit number { alert-percent [ route-unchanged ] | simply-alert }

                    This configuration prevents a VPN instance IPv6 address family on a PE from
                    receiving too many route prefixes.
                          NOTE

