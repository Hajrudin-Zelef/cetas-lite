---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-105
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [14673, 14791]
sha256: 32ae996030c89d0183c717301ee259627703632c03748b436a3476c18c5c6631
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             232
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        q.   Configure MPLS label distribution based on the VPN instance IPv4
                             address family (known as one-label-per-instance), so that only one label
                             is assigned to all the routes of the VPN instance IPv4 address family.
                             apply-label per-instance

                    ●   Configure the Hub-PE (which connects to a Hub-CE over a single link).
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VPN instance view of vpnhub.
                             ip vpn-instance vpn-instance-name1

                        c.   (Optional) Configure a description for the VPN instance.
                             description description-information

                             The description is used to record the purpose of creating the VPN
                             instance and the CEs associated with the VPN instance.
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
                        g.   Configure export VPN targets for the VPN instance, so that it adds these
                             VPN targets to hub site routes when advertising these routes to all spoke
                             sites.
                             vpn-target vpn-target2 &<1-8> export-extcommunity

                             The vpn-target2 list here must contain the import VPN targets configured
                             on all Spoke-PEs.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        h.   (Optional) Configure an import route-policy for the VPN instance IPv4
                             address family.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             233
VPN Configuration
VPN Configuration                                                                 3 IPv4 L3VPN Configuration

                              import route-policy policy-name

                              In addition to using VPN targets to control VPN route import and export,
                              an import route-policy can be configured to better control VPN route
                              import. An import route-policy can be used to filter routes to be imported
                              into the VPN instance IPv4 address family or modify route attributes.
                              This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                              S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                              S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                              S5755-H.
                        i.    (Optional) Configure an export route-policy for the VPN instance IPv4
                              address family.
                              export route-policy policy-name [ add-ert-first ]

                              Before performing this step, you must create a route-policy that filters
                              default routes and then run the export route-policy command to enable
                              the Hub-PE to advertise only default routes to Spoke-PEs.
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
                        j.    Configure the device to assign a unique label to each route sent from the
                              current VPN instance IPv4 address family to its peer PE and forward the
                              data packets received from its peer PE through outbound interfaces found
                              in the local ILM.
                              apply-label per-route pop-go

                              The apply-label per-route pop-go command allows the local device to
                              record in its ILM table the mapping between the label assigned to a BGP-
                              VPNv4 route and the outbound interface for the route before the Hub-PE
                              sends the route to other Spoke-PEs. After the Hub-PE receives a labeled
                              data packet from a Spoke-PE, the local device directly searches the ILM
                              table for an outbound interface based on the label carried in the packet,
                              instead of searching the IP forwarding table based on the longest-match
                              rule. The device then removes the label and forwards the packet through
                              the found outbound interface. This implementation prevents dat packets
                              from being directly forwarded to other Spoke-PEs without passing
                              through the Hub-CE.
                    ----End

3.15.3 Binding an Interface to a VPN Instance




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             234

