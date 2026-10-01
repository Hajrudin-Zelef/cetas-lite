---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-104
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [14561, 14672]
sha256: 1315323535930a2acc9bff5a7abd0ba2fa86e7ff4f66cef9ef522b0dd4b5f142
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        d.   Enter the VPN instance IPv4 address family view.
                             ipv4-family
                        e.   Configure an RD for the VPN instance IPv4 address family.
                             route-distinguisher route-distinguisher

                             A VPN instance IPv4 address family takes effect only after being
                             configured with an RD. Before configuring an RD, you can configure only
                             the description about the VPN instance. No other parameters can be
                             configured.
                        f.   Configure import VPN targets for the VPN instance, so that it can receive
                             VPNv4 routes advertised by all Spoke-PEs.
                             vpn-target vpn-target1 &<1-8> import-extcommunity

                             The vpn-target1 list here must contain the export VPN targets configured
                             on all Spoke-PEs.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        g.   (Optional) Configure an import route-policy for the VPN instance IPv4
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
                        h.   (Optional) Configure an export route-policy for the VPN instance IPv4
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
                        i.   Return to the system view.
                             quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          231
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                        j.   Enter the VPN instance view of vpn_out.
                             ip vpn-instance vpn-instance-name2
                        k.   (Optional) Configure a description for the VPN instance.
                             description description-information

                             The description is used to record the purpose of creating the VPN
                             instance and the CEs associated with the VPN instance.
                        l.   Enter the VPN instance IPv4 address family view.
                             ipv4-family
                        m. Configure an RD for the VPN instance IPv4 address family.
                             route-distinguisher route-distinguisher

                             A VPN instance IPv4 address family takes effect only after being
                             configured with an RD. Before configuring an RD, you can configure only
                             the description about the VPN instance. No other parameters can be
                             configured.
                        n.   Configure export VPN targets for the export of routes from all hub and
                             spoke sites.
                             vpn-target vpn-target2 &<1-8> export-extcommunity

                             The vpn-target2 list here must contain the import VPN targets configured
                             on all Spoke-PEs.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        o.   (Optional) Configure an import route-policy for the VPN instance IPv4
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
                        p.   (Optional) Configure an export route-policy for the VPN instance IPv4
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

