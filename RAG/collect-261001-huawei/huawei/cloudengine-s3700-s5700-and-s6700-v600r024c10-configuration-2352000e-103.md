---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-103
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [14429, 14560]
sha256: e68c9d1b4e18e6f33d92f17db0749fce70454e9d6efaa639e99ae93adec60b71
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        Figure 3-47 EBGP running between the Hub-CE and Hub-PE, and IGP running
                        between Spoke-PEs and Spoke-CEs




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          228
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


                        The networking topology is similar to that shown in Figure 3-45. The AS_Path
                        attribute of a Spoke-CE route forwarded by the Hub-CE to the Hub-PE
                        contains the AS number of the Hub-PE. The Hub-PE must therefore be
                        configured to allow local AS number repetition in the AS_Path attribute of a
                        route.

3.15.2 Configuring a VPN Instance

Context
                    In hub-spoke networking, the PE connected to a central site (hub site) is called a
                    Hub-PE and the PE connected to a non-central site (spoke site) is called a Spoke-
                    PE. Spoke-PEs and Hub-PEs must have VPN instances configured. If the Hub-CE
                    and Hub-PE are connected over dual links, the Hub-PE must have two VPN
                    instances configured, for example, vpn_in and vpn_out. If the Hub-PE and Hub-CE
                    are connected over a single link, the Hub-PE needs only one VPN instance, for
                    example, vpnhub.

                         NOTE

                        Steps 1 to 8 are used to configure a VPN instance. Configurations of different VPN instances
                        are similar. If different VPN instances are configured on the same device, these VPN
                        instances must have different names, RDs, and descriptions.


Procedure
                    ●   Configure the Spoke-PE.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VPN instance view of vpn_in.
                             ip vpn-instance vpn-instance-name

                        c.   (Optional) Configure a description for the VPN instance.
                             description description-information

                        d.   Enter the VPN instance IPv4 address family view.
                             ipv4-family

                        e.   Configure an RD for the VPN instance IPv4 address family.
                             route-distinguisher route-distinguisher

                        f.   Configure import VPN targets for the VPN instance, so that it can receive
                             VPNv4 routes advertised by the Hub-PE.
                             vpn-target vpn-target2 &<1-8> import-extcommunity

                             vpn-target2 must be in the export VPN target list configured on the Hub-
                             PE.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        g.   Configure export VPN targets for the VPN instance, so that it adds these
                             VPN targets to routes imported from sites accessed by the Spoke-PE
                             when advertising these routes.
                             vpn-target vpn-target1 &<1-8> export-extcommunity


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   229
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                             vpn-target1 must be in the import VPN target list configured on the Hub-
                             PE.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        h.   (Optional) Configure an import route-policy for the VPN instance IPv4
                             address family.
                             import route-policy policy-name

                             In addition to using VPN targets to control VPN route import and export,
                             an import route-policy can be configured to better control VPN route
                             import. An import route-policy can be used to filter routes to be imported
                             into the VPN instance IPv4 address family or modify route attributes.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        i.   (Optional) Configure an export route-policy for the VPN instance IPv4
                             address family.
                             export route-policy policy-name [ add-ert-first ]

                             In addition to using VPN targets to control VPN route import and export,
                             an export route-policy can be configured to better control VPN route
                             export. An export route-policy can be used to filter routes to be
                             advertised to other PEs or modify route attributes.
                             By default, export VPN targets are added to VPN routes after these routes
                             are matched against an export route-policy. If the export route-policy
                             contains VPN target-related filtering rules, it cannot apply to VPN routes.
                             To address this issue, configure the add-ert-first parameter. This instructs
                             the device to add export VPN targets to VPN routes before matching
                             these routes against the export route-policy.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        j.   (Optional) Configure MPLS label distribution based on the VPN instance
                             IPv4 address family (known as one-label-per-instance), so that only one
                             label is assigned to all the routes of the VPN instance IPv4 address family.
                             apply-label per-instance
                    ●   Configure the Hub-PE (which connects to a Hub-CE over dual links).
                        a.   Enter the system view.
                             system-view
                        b.   Enter the VPN instance view of vpn_in.
                             ip vpn-instance vpn-instance-name1
                        c.   (Optional) Configure a description for the VPN instance.
                             description description-information

                             The description is used to record the purpose of creating the VPN
                             instance and the CEs associated with the VPN instance.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             230
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


