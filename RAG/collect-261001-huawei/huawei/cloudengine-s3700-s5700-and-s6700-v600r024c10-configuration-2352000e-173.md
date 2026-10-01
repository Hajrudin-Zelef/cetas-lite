---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-173
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [25034, 25179]
sha256: ff7aafa9ffa8f78208ac973b542b925d3d6631a7976dc930a30e3998b14b5767
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                             The VPN instance IPv6 address family takes effect only after the RD is
                             configured. Before configuring an RD, you can configure only the
                             description about the VPN instance. No other parameters can be
                             configured.
                        m. Configure export VPN targets for the export of routes from all Hub and
                           Spoke sites.
                             vpn-target vpn-target2 &<1-8> export-extcommunity

                             The vpn-target2 list here must contain the import VPN targets configured
                             on all Spoke-PEs.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        n.   (Optional) Configure an import route-policy for the VPN instance IPv6
                             address family.
                             import route-policy policy-name

                             In addition to using VPN targets to control VPN route import and export,
                             an import route-policy can be configured to better control VPN route
                             import. An import route-policy can be used to filter routes to be imported
                             into the VPN instance IPv6 address family or modify route attributes.
                             This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                             S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                             S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                             S5755-H.
                        o.   (Optional) Configure an export route-policy for the VPN instance IPv6
                             address family.
                             export route-policy policy-name [ add-ert-first ]

                             In addition to using VPN targets to control VPN route import and export,
                             an export route-policy can be configured to better control VPN route
                             export. An export route-policy can be used to filter routes to be
                             advertised to other PEs or modify route attributes.
                             By default, export VPN targets are added to VPN routes after these routes
                             are matched against an export route-policy. If the export route-policy
                             contains VPN target-related filtering rules, it cannot apply to VPN routes.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          395
VPN Configuration
VPN Configuration                                                                  4 IPv6 L3VPN Configuration


                               To address this issue, configure the add-ert-first parameter. This instructs
                               the device to add export VPN targets to VPN routes before matching
                               these routes against the export route-policy.
                               This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                               S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2,
                               S5735I-S-V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and
                               S5755-H.
                    ----End

4.14.4 Binding an Interface to a VPN Instance

Context
                    After an interface is bound to a VPN instance, the interface becomes a part of the
                    VPN. Packets entering the interface will be forwarded based on the VRF table of
                    the VPN.
                    In a Hub-CE dual-link access scenario, two interfaces or sub-interfaces are
                    required on the Hub-PE:
                    ●    One is bound to VPN-in for importing routes advertised by Spoke-PEs.
                    ●    The other is bound to VPN-out for exporting the routes of all Hub and Spoke
                         sites.
                    If a Hub-PE and a Hub-CE are connected through a single link, only one interface
                    or sub-interface on the Hub-PE needs to be bound to the VPN instance vpnhub.
                    Perform the following steps on the Hub-PE and all Spoke-PEs.

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

                          NOTE

                        After the ip binding vpn-instance command is run on the interface, Layer 3 features such
                        as the IP address and routing protocol configured on the interface are deleted.

         Step 5 Enable IPv6 on the interface.
                    ipv6 enable


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 396
VPN Configuration
VPN Configuration                                                                              4 IPv6 L3VPN Configuration


         Step 6 Configure an IPv6 address for the interface.
                    ipv6 address { ipv6-address prefix-length | ipv6-address/prefix-length }

                    ----End

4.14.5 Configuring Route Exchange Between the Hub-PE and
Spoke-PE

Context
                    MP-IBGP, which introduces extended community attributes into BGP, can advertise
                    VPNv4 routes between PEs.

                    MP-IBGP peer relationships need be established between the Hub-PE and each
                    Spoke-PE. Spoke-PEs do not need to exchange routes directly and therefore do not
                    need to establish MP-IBGP peer relationships.

                    Perform the following steps on the Hub-PE and all Spoke-PEs.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Configure the remote PE as a peer.
                    peer peer-address as-number as-number

         Step 4 Specify an interface for setting up a TCP connection with the BGP peer.
                    peer peer-address connect-interface loopback interface-number

                          NOTE

                         A PE must use a loopback interface address with a 32-bit mask to set up an MP-IBGP peer
                         relationship with the peer PE so that VPN routes can recurse to tunnels. The route to the
                         local loopback interface is advertised to the peer PE using IGP on the MPLS backbone
                         network.

         Step 5 Enter the BGP-VPNv6 address family view.
                    ipv6-family vpnv6 [ unicast ]

         Step 6 Enable the function to exchange VPNv6 routing information with the peer.
                    peer peer-address enable

                    ----End

