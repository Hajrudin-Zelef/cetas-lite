---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-171
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [24795, 24921]
sha256: 40c3a7951f28d5aae39e88d9d3854e60f9e7e6a3620f916e42675894ca63710a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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

                    The VPN instance IPv6 address family takes effect only after the RD is configured.
                    The RDs of different VPN instance IPv6 address families on a PE must be different.

                          NOTE

                         If you perform this step in the view of a newly created VPN instance, the IPv6 address
                         family is automatically enabled and the IPv6 address family view is automatically displayed.

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
         Step 7 (Optional) Configure the maximum number of route prefixes allowed for the VPN
                instance IPv6 address family.
                    prefix limit number { alert-percent [ route-unchanged ] | simply-alert }

                    This configuration prevents a VPN instance IPv6 address family on a PE from
                    receiving too many route prefixes.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         391
VPN Configuration
VPN Configuration                                                                 4 IPv6 L3VPN Configuration


         Step 8 (Optional) Configure the device to assign a unique label to each VPNv6 route sent
                from the current VPN instance IPv6 address family to its peer PE and forward the
                data packets received from its peer PE through outbound interfaces found in the
                local ILM.
                    apply-label per-route pop-go

                    The apply-label per-route pop-go command allows the local device to record in
                    its ILM table the mapping between the label assigned to a VPNv6 route and the
                    outbound interface for the route before the device sends the route to a BGP-
                    VPNv6 peer. After the local device receives a labeled data packet from its BGP-
                    VPNv6 peer, the local device directly searches the ILM table for an outbound
                    interface based on the label carried in the packet, instead of searching the IP
                    forwarding table based on the longest-match rule. The device then removes the
                    label and forwards the packet through the found outbound interface. This
                    implementation significantly accelerates packet forwarding.

         Step 9 (Optional) Configure MPLS label distribution based on the VPN instance IPv6
                address family (known as one-label-per-instance), so that only one label is
                assigned to all the routes of the VPN instance IPv6 address family.
                    apply-label per-instance

                    ----End

4.14.3 Configuring Route-related Attributes in the VPN
Instance

Context
                    In Hub-Spoke networking, you can configure VPN targets on the Hub-PE and
                    Spoke-PEs to control the advertisement of VPN routes. The import VPN target list
                    configured on the Hub-PE must contain the export VPN targets configured on all
                    the Spoke-PEs. The export VPN target list configured on the Hub-PE must contain
                    the import VPN targets configured on all the Spoke-PEs.

                    Controlling route import by the VPN instance IPv6 address family by configuring
                    VPN targets is also a key part of the Hub-Spoke solution.

Procedure
                    ●    Configure the Spoke-PE.
                         a.   Enter the system view.
                              system-view

                         b.   Enter the VPN instance view of VPN-in.
                              ip vpn-instance vpn-instance-name

                         c.   Enter the VPN instance IPv6 address family view.
                              ipv6-family

                         d.   Configure an RD for the VPN instance IPv6 address family.
                              route-distinguisher route-distinguisher

                         e.   Configure import VPN targets for the VPN instance, so that it can receive
                              VPNv6 routes advertised by the Hub-PE.
                              vpn-target vpn-target2 &<1-8> import-extcommunity


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                            392
VPN Configuration
VPN Configuration                                                                4 IPv6 L3VPN Configuration


                             vpn-target2 must be in the export VPN target list configured on the Hub-
                             PE.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        f.   Configure export VPN targets for the VPN instance, so that it adds these
                             VPN targets to routes imported from sites accessed by the Spoke-PE
                             when advertising these routes.
                             vpn-target vpn-target1 &<1-8> export-extcommunity

                             vpn-target1 must be in the import VPN target list configured on the Hub-
                             PE.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        g.   (Optional) Configure an import route-policy for the VPN instance IPv6
                             address family.
                             import route-policy policy-name

