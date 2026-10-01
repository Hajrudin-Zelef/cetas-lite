---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-121
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [17106, 17235]
sha256: 258df71a7f993de9dcdfe4a0fa26de8816bd8930ed41388c25dd86e395e9fc6f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.

         Step 7 Set the maximum number of route prefixes allowed for the VPN instance IPv6
                address family.
                    prefix limit number { alert-percent [ route-unchanged ] | simply-alert }

                    This configuration prevents a VPN instance IPv6 address family on a PE from
                    receiving too many route prefixes.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         271
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


                          NOTE

                         After the prefix limit command is run to increase the maximum number of route prefixes
                         allowed for a VPN instance IPv6 address family or the undo prefix limit command is run to
                         cancel the limit, the device adds newly received route prefixes of various protocols to the
                         VRF table for the VPN instance IPv6 address family.
                         When the number of route prefixes exceeds the limit, direct routes and static routes can still
                         be added to the VRF table for the VPN instance IPv6 address family.

         Step 8 (Optional) Configure an import route-policy for the VPN instance IPv6 address
                family.
                    import route-policy policy-name

                    In addition to using VPN targets to control VPN route import and export, an
                    import route-policy can be configured to better control VPN route import. An
                    import route-policy can be used to filter routes to be imported into the VPN
                    instance IPv6 address family or modify route attributes.
                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
         Step 9 (Optional) Configure an export route-policy for the VPN instance IPv6 address
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
                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
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

4.5.2 Binding an Interface to the IPv6 VPN Instance

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      272
VPN Configuration
VPN Configuration                                                                    4 IPv6 L3VPN Configuration


Prerequisites
                    A VPN instance has been created, and the IPv6 address family has been enabled
                    for the VPN instance.

Context
                    If the interface belonging to a VPN instance is not bound to the VPN instance, the
                    interface functions as a public network interface and cannot forward VPN data.
                    An interface becomes a VPN interface after being bound to a VPN instance. You
                    must configure an IP address for the interface, so that the PE can exchange
                    routing information with its connected CE through this interface.
                    After an interface is bound to a VPN instance, Layer 3 configurations, such as IP
                    address and routing protocol configurations, are deleted from the interface.
                    After you disable an address family in the VPN instance, the corresponding
                    address configuration is deleted from the interface. If no address family
                    configuration exists in the VPN instance, the interface is unbound from the VPN
                    instance.

                          NOTE

                         Loopback interfaces can be bound to VPN instances for connectivity testing between VPNs,
                         under the prerequisite that VLANIF or physical interfaces have been bound to the VPN
                         instances prior to loopback interfaces. In actual service scenarios, a specific VLANIF,
                         physical, or tunnel interface must be bound to a VPN instance.


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

                          NOTE

                         Running the ip binding vpn-instance command on an interface deletes Layer 3 (including
                         IPv4 and IPv6) configurations, such as IP address and routing protocol configurations, on
                         the interface. If needed, reconfigure them after running the command.

         Step 5 Enable IPv6 on the interface.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   273

