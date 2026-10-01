---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-330
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["compute", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [49305, 49494]
sha256: 587e73704106eef107c353bb9f5c761e11c51d0e0f47cda0bb32f4b6b8fabbb7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   CE2
                        #
                        sysname CE1
                        #
                        vlan batch 10 100
                        #
                        stp region-configuration
                         instance 1 vlan 10 100
                        #
                        erps ring 1
                         control-vlan 100
                         protected-instance 1
                         version v2
                         sub-ring
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         undo port trunk allow-pass vlan 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   791
VPN Configuration
VPN Configuration                                                                                      6 VPLS Configuration

                           port trunk allow-pass vlan 10 100
                           stp disable
                           erps ring 1
                          #
                          interface 10GE1/0/2
                           port link-type trunk
                           undo port trunk allow-pass vlan 1
                           port trunk allow-pass vlan 10 100
                           stp disable
                           erps ring 1 rpl owner
                          #
                          return



6.19 Maintaining VPLS

6.19.1 Collecting the Traffic Statistics of a VPLS PW
Context
                    To monitor the network operating status and locate faults more easily on a VPLS
                    network, configure the traffic statistics collection function.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the VSI view.
                    vsi vsi-name [ static | auto ]

         Step 3 Enable traffic statistics collection for VPLS PWs. Run the following commands
                based on the VPLS type:
                    ●     LDP VPLS
                          a.    Configure LDP as the PW signaling protocol and enter the VSI-LDP view.
                                pwsignal ldp

                          b.    Enable traffic statistics collection for LDP VPLS PWs.

                                ▪     Enable traffic statistics collection for PWs globally.
                                      traffic-statistics enable

                                ▪     Enable traffic statistics collection for a specified LDP VPLS PW.
                                      traffic-statistics peer peer-address [ negotiation-vc-id vc-id ] enable

                    ●     BGP VPLS
                          a.    Configure BGP as the PW signaling protocol and enter the VSI BGP view.
                                pwsignal bgp

                          b.    Enable traffic statistics collection for BGP VPLS PWs.
                                traffic-statistics peer peer-address remote-site site-id enable

                    ●     BGP AD VPLS
                          a.    Enter the VSI-BGPAD view.
                                bgp-ad

                          b.    Enable traffic statistics collection for BGP AD VPLS PWs.

                                ▪     Enable traffic statistics collection for PWs globally.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          792
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                                   traffic-statistics enable

                              ▪    Enable traffic statistics collection for a specified BGP AD VPLS PW.
                                   traffic-statistics peer peer-address enable

                    ----End

6.19.2 Checking the Traffic Statistics of a VPLS PW

Context
                    After configuring traffic statistics collection for a VPLS PW, you can check the
                    collected traffic statistics of the VPLS PW.

                         NOTE

                        If a PW goes down within 5 minutes, traffic statistics collected before the PW goes down
                        cannot be used to compute the traffic rate in the 5-minute period.


Procedure
                    ●   Run the display traffic-statistics vsi vsi-name [ peer peer-address
                        [ negotiation-vc-id vc-id | ldp129 | remote-site remote-site-id ]] command
                        in any view to check traffic statistics of a specified VSI or peer.

                    ----End

6.19.3 Clearing the Traffic Statistics of a VPLS PW

Context
                    To clear traffic statistics of a VPLS PW, run the commands listed in the following
                    table in the user view.



                        NOTICE

                    Traffic statistics cannot be restored after you clear them. Exercise caution when
                    clearing traffic statistics.



                    Table 6-13 Clearing the traffic statistics of a VPLS PW

                     Operation                     Command

                     Clear the traffic             reset traffic-statistics vsi all
                     statistics of all VSIs.

                     Clear the traffic             reset traffic-statistics vsi name vsi-name [ peer peer-
                     statistics of a               address [ negotiation-vc-id vc-id | ldp129 | remote-site
                     specified VSI.                remote-site-id ] ]



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    793
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


6.19.4 Monitoring the VPLS Operating Status
Context
                    In routine maintenance, you can run the commands listed in the following table in
                    any view to check the VPLS operating status.

                    Table 6-14 Monitoring the VPLS operating status
                     Operation                Command

                     Check VSI                display vsi [ name vsi-name ] [ verbose ]
                     information.

                     Check remote VSI         display vsi remote ldp [ [ router-id ip-address ] [ pw-
                     information.             id pw-id ] | [ verbose ] | [ unmatch ] ]

                     Check VPLS               display vpls connection [ ldp | bgp | vsi vsi-name ]
                     connection               [ down | up ] [ verbose ]
                     information.

                     Check information        display vsi pw out-interface [ vsi vsi-name ]
                     about the outbound
                     interface of a VSI PW.

                     Check information        display l2vpn vsi-list tunnel-policy policy-name
                     about the tunnel
                     policy used by a VSI.

                     Check the forwarding     display vpls forwarding-info [ vsi vsi-name [ peer
                     information of a VSI.    peer-address [ negotiation-vc-id vc-id | remote-site
                                              site-id ] ] | state { up | down } ] [ verbose ]
                     Check information        display vsi services { vsi-name | all | interface
                     about the AC             interface-type interface-number | vlan vlan-id |
                     interface associated     interface interface-name }
                     with VSIs.

                     Check the binding        display admin-vsi binding [ admin-vsi vsi-name ]
                     between the
                     management VSI and
                     service VSI.




