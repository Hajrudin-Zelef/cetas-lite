---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-32
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3577, 3714]
sha256: 30644c164deaf266a5539310a583ae978f2e23ce685553a1aa1e601c07896e1f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

                          NOTE

                         The name of a VPN instance is case sensitive. For example, vpn1 and VPN1 are considered
                         different VPN instances.

         Step 3 (Optional) Configure a description for the VPN instance.
                    description description-information

         Step 4 Enable the IPv4 address family for the VPN instance, and enter the VPN instance
                IPv4 address family view.
                    ipv4-family

                    Configurations in a VPN instance can be performed only after an address family is
                    enabled for the VPN instance based on the advertised route and forwarded data
                    type.
         Step 5 Configure an RD for the VPN instance IPv4 address family.
                    route-distinguisher route-distinguisher

                    A VPN instance IPv4 address family takes effect only after being configured with
                    an RD. The RDs of different VPN instance IPv4 address families on a PE must be
                    different.

                          NOTE

                         If you configure an RD for the VPN instance IPv4 address family in the created VPN
                         instance view, the VPN instance IPv4 address family is automatically enabled and the VPN
                         instance IPv4 address family view is automatically displayed.
                         The S3710-H series does not support dynamic routing protocols. Therefore, this command
                         does not take effect after being delivered.

         Step 6 Configure VPN targets for the VPN instance IPv4 address family.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    56
VPN Configuration
VPN Configuration                                                                              3 IPv4 L3VPN Configuration

                    vpn-target vpn-target &<1-8> [ both | export-extcommunity | import-extcommunity ]

                    VPN targets are a type of BGP extended community attribute and used to control
                    the import and export of VPN routes. You can configure a maximum of eight
                    import VPN targets and eight export VPN targets each time the vpn-target
                    command is run. Run this command multiple times if you want to configure more
                    VPN targets.

                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S6750E-
                    S, S6750-S, S5755-S, S5755E-H, and S5755-H.

         Step 7 (Optional) Configure the maximum number of route prefixes allowed by the VPN
                instance IPv4 address family.
                    prefix limit number { alert-percent [ route-unchanged ] | simply-alert }

                    This configuration prevents a VPN instance IPv4 address family on a PE from
                    receiving too many route prefixes.
                          NOTE

                         After the prefix limit command is run to increase the allowed maximum number of route
                         prefixes in the VPN instance IPv4 address family or the undo prefix limit command is run
                         to cancel the limit, the system adds newly received route prefixes of various protocols to
                         the VPN IP routing table.
                         After the number of route prefixes exceeds the maximum limit, direct and static routes can
                         still be added to the routing table of the VPN instance IPv4 address family.

         Step 8 (Optional) Configure an import route-policy for the VPN instance IPv4 address
                family.
                    import route-policy policy-name

                    In addition to using VPN targets to control VPN route import and export, an
                    import route-policy can be configured to better control VPN route import. An
                    import route-policy can be used to filter routes to be imported into the VPN
                    instance IPv4 address family or modify route attributes.

                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.

         Step 9 (Optional) Configure an export route-policy for the VPN instance IPv4 address
                family.
                    export route-policy policy-name [ add-ert-first ]

                    In addition to using VPN targets to control VPN route import and export, an
                    export route-policy can be configured to better control VPN route export. An
                    export route-policy can be used to filter routes to be advertised to other PEs or
                    modify route attributes.

                    By default, export VPN targets are added to VPN routes after these routes are
                    matched against an export route-policy. If the export route-policy contains VPN
                    target-related filtering rules, it cannot apply to VPN routes. To address this issue,
                    configure the add-ert-first parameter. This instructs the device to add export VPN
                    targets to VPN routes before matching these routes against the export route-
                    policy.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          57
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S6750E-
                    S, S6750-S, S5755-S, S5755E-H, and S5755-H.

        Step 10 (Optional) Disable local route leaking in L3VPN scenarios.
                    local-cross unicast disable

                    By default, local route leaking is enabled in L3VPN scenarios. If you do not want
                    the routes of a VPN instance to be leaked to other VPN instances, run the local-
                    cross unicast disable command in the view of that VPN instance to disable local
                    route leaking.

                    This command is supported only by the S6780-H, S6730E-H-V2, S6730-H-V2,
                    S5735E-S-V2, S5735-S-V2, S5735R-S-V2, S5735I-S-V2, S5755E-H, S5755-H, S5732-
                    H-V2, S5735I-H-V2, S6750E-S, S6750-S, S6750-H, and S5755-S.

                    ----End

3.6.3 Binding an Interface to an IPv4 VPN Instance

Prerequisites
                    A VPN instance has been created and the IPv4 address family has been enabled
                    for the VPN instance.

Context
                    If the interface belonging to the VPN is not bound to the VPN instance, the
                    interface functions as a public network interface and cannot forward VPN data.

                    An interface bound to a VPN instance becomes a VPN interface on a PE and must
                    be assigned an IP address so that the PE can exchange routes with its connected
                    CE.

                    After an interface is bound to a VPN instance, Layer 3 features, such as the IP
                    address and routing protocol configured on the interface, are deleted.

