---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-158
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22665, 22824]
sha256: ee0b76dd7413db8962ee17c14d28af047035f56baaab93fce6520b22aceb22d5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

4.11.1 Configuring a VPN Instance
Context
                    An instance is created to comprise the VPN forwarding information for each VPN
                    in an IPv6 L3VPN. This instance is called a VPN instance or VPN routing and
                    forwarding (VRF) table. In related standards (BGP/MPLS IP VPNs), a VPN instance
                    is also called a per-site forwarding table. VPN instances are required in all IPv6
                    L3VPN networking solutions.

                    In an IPv6 VPN provider edge (6VPE) scenario, both IPv4 and IPv6 address families
                    must be enabled in a VPN instance. The IPv4 routing table of the VPN instance
                    manages routes used to establish VPN LDP LSPs, whereas the IPv6 routing table of
                    the VPN instance manages received VPN IPv6 routes.

Procedure
         Step 1 Enter the system view.
                    system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         357
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

                            NOTE

                           The name of a VPN instance is case-sensitive. For example, vpn1 and VPN1 are considered
                           different VPN instances.

         Step 3 (Optional) Configure a description for the VPN instance.
                    description description-information

         Step 4 Enable the VPN instance IPv4 address family, and enter the VPN instance IPv4
                address family view.
                    ipv4-family

                    Configurations in a VPN instance can be performed only after an address family is
                    enabled for the VPN instance based on the advertised route and type of forwarded
                    data.

         Step 5 Configure an RD for the VPN instance IPv4 address family.
                    route-distinguisher route-distinguisher

                    A VPN instance IPv4 address family takes effect only after having an RD
                    configured. The RDs of different VPN instance IPv4 address families on a PE must
                    be different.

                            NOTE

                           If you perform this step in the view of a newly created VPN instance, the VPN instance IPv4
                           address family is automatically enabled and the VPN instance IPv4 address family view is
                           automatically displayed.

         Step 6 Configure VPN targets for the VPN instance IPv4 address family.
                    vpn-target vpn-target &<1-8> [ both | export-extcommunity | import-extcommunity ]

                    VPN targets are a type of BGP extended community attribute used to control the
                    import and export of VPN routes. You can configure a maximum of eight import
                    VPN targets and eight export VPN targets each time the vpn-target command is
                    run. Run this command multiple times if you want to configure more VPN targets
                    in the VPN instance IPv4 address family.

                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.

         Step 7 Exit the VPN instance IPv4 address family view.
                    quit

         Step 8 Enable the VPN instance IPv6 address family, and enter the VPN instance IPv6
                address family view.
                    ipv6-family

                    Configurations in a VPN instance can be performed only after an address family is
                    enabled for the VPN instance based on the advertised route and type of forwarded
                    data.

         Step 9 Configure an RD for the VPN instance IPv6 address family.
                    route-distinguisher route-distinguisher


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                     358
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


                    A VPN instance IPv6 address family takes effect only after having an RD
                    configured. The RDs of different VPN instance IPv6 address families on a PE must
                    be different.

                          NOTE

                         If you perform this step in the view of a newly created VPN instance, the VPN instance IPv6
                         address family is automatically enabled and the VPN instance IPv6 address family view is
                         automatically displayed.

        Step 10 Configure VPN targets for the VPN instance IPv6 address family.
                    vpn-target vpn-target &<1-8> [ both | export-extcommunity | import-extcommunity ]

                    VPN targets are a type of BGP extended community attribute used to control the
                    import and export of VPN routes. You can configure a maximum of eight import
                    VPN targets and eight export VPN targets each time the vpn-target command is
                    run. Run this command multiple times if you want to configure more VPN targets
                    in the VPN instance IPv6 address family.

                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.

                    ----End

4.11.2 Binding an Interface to a VPN Instance

Context
                    If the interface belonging to a VPN instance is not bound to the VPN instance, the
                    interface functions as a public network interface and cannot forward VPN data.

                    Perform the following steps on each PE that connects to a CE.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface to be bound to a VPN instance.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.

         Step 4 Bind the interface to a VPN instance.
                    ip binding vpn-instance vpn-instance-name

                    By default, an interface functions as a public network interface that is not bound
                    to any VPN instance.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    359
VPN Configuration
VPN Configuration                                                                     4 IPv6 L3VPN Configuration


                          NOTE

                         Running the ip binding vpn-instance command will delete Layer 3 features such as the IP
                         address and routing protocol configured on the interface. To use these features, you need to
                         reconfigure them.

         Step 5 Configure an IP address for the interface.
                    ip address ip-address { mask | mask-length }

