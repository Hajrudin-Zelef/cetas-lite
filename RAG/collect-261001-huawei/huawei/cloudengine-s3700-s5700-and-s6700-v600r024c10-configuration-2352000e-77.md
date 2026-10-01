---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-77
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10501, 10673]
sha256: 60de6fe4dba9e4184412745c80100b2ff1de17457a85c1f48e071c59691f7de0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.
         Step 4 Configure an IP address for the interface.
                    ip address ip-address { mask | mask-length }

         Step 5 Enable MPLS.
                    mpls

         Step 6 Return to the system view.
                    quit

         Step 7 Enter the BGP view.
                    bgp as-number

         Step 8 Establish an IBGP peer relationship between the PE and ASBR in the same AS.
                    peer peer-address as-number as-number

         Step 9 Configure a loopback interface as the outbound interface of the BGP session.
                    peer peer-address connect-interface loopback interface-number

        Step 10 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]

        Step 11 Enable the function to exchange VPNv4 routes between the PE and ASBR in the
                same AS.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             166
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration

                    peer peer-address enable

                    ----End

3.13.4 Configuring ASBRs Not to Filter VPNv4 Routes Based on
VPN Targets

Context
                    In inter-AS VPN Option B, ASBRs do not have VPN instances. If you want ASBRs to
                    keep received VPNv4 routes, configure ASBRs not to filter VPNv4 routes based on
                    VPN targets.

                    For configuration details, see Configuring ASBRs Not to Filter VPNv4 Routes
                    Based on VPN Targets.

3.13.5 (Optional) Using a Route-Policy to Control VPNv4
Route Import and Export on ASBRs

Context
                    ASBRs can use a route-policy to filter undesired VPNv4 routes based on VPN
                    targets or RDs.

                    For configuration details, see (Optional) Using a Route-Policy to Control VPNv4
                    Route Import and Export on ASBRs.

3.13.6 (Optional) Configuring One-Label-per-Next-Hop Label
Distribution on ASBRs

Context
                    In an inter-AS VPN Option B scenario, after one-label-per-next-hop label
                    distribution is enabled on an ASBR, the ASBR assigns only one label to VPNv4
                    routes that share the same next hop and same outgoing label. Compared with on-
                    label-per-route label distribution, one-label-per-next-hop label distribution
                    significantly saves label resources on the ASBR.

                    Perform the following steps on each ASBR.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         167
VPN Configuration
VPN Configuration                                                                              3 IPv4 L3VPN Configuration


         Step 4 Enable one-label-per-next-hop label distribution for VPNv4 routes on the ASBR.
                    apply-label per-nexthop




                        NOTICE

                    After one-label-per-next-hop label distribution is enabled or disabled on an ASBR,
                    the labels assigned by the ASBR to routes change. As a result, temporary packet
                    loss may occur.

                    ----End

3.13.7 Configuring a VPN Instance on ASBRs
Context
                    If an ASBR also functions as a PE, you need to configure a VPN instance enabled
                    with the IPv4 address family on the ASBR to manage VPN routes.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

         Step 3 Enable the IPv4 address family for the VPN instance, and enter the VPN instance
                IPv4 address family view.
                    ipv4-family

         Step 4 Configure an RD for the VPN instance IPv4 address family.
                    route-distinguisher route-distinguisher

         Step 5 Configure VPN targets for the VPN instance IPv4 address family.
                    vpn-target vpn-target &<1-8> [ both | export-extcommunity | import-extcommunity ]

                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
         Step 6 (Optional) Configure the maximum number of route prefixes allowed by the VPN
                instance IPv4 address family.
                    prefix limit number { alert-percent [ route-unchanged ] | simply-alert }

                    This configuration prevents a VPN instance IPv4 address family on a PE from
                    receiving too many route prefixes.
         Step 7 (Optional) Configure an import route-policy for the VPN instance IPv4 address
                family.
                    import route-policy policy-name

                    In addition to using VPN targets to control VPN route import and export, an
                    import route-policy can be configured to better control VPN route import. An
                    import route-policy can be used to filter routes to be imported into the VPN
                    instance IPv4 address family or modify route attributes.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         168
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration


                    This command is supported only by the S6780-H, S6750-H, S6730E-H-V2, S6730-
                    H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-V2, S5735I-
                    H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.

         Step 8 (Optional) Configure an export route-policy for the VPN instance IPv4 address
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

                    ----End

