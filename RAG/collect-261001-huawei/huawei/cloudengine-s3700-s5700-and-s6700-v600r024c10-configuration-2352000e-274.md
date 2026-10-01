---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-274
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["compute", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40285, 40407]
sha256: 2d4a2bd32fc7829652a7659e1fc69dfd49b4811cd65dc1857de0dcdf9d1ccd21
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     BGP VPLS                   ● BGP VPLS-enabled          ● Labels are wasted.
                                                  devices exchange BGP      ● Site ID management
                                                  Update messages that        is needed.
                                                  carry label block
                                                  information and           ● There is no standard
                                                  compute and                 mechanism of
                                                  compare received            clearing MAC address
                                                  label block                 entries.
                                                  information with the
                                                  local label block
                                                  information before
                                                  establishing a PW
                                                  with each other.
                                                ● BGP VPLS
                                                  implementation is
                                                  simpler than BGP AD
                                                  VPLS implementation.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          646
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                     VPLS Mode                  Advantage                   Disadvantage

                     BGP AD VPLS                ● Label resources are       ● To establish a PW
                                                  saved.                      with each other, BGP
                                                ● BGP AD VPLS-enabled         AD VPLS-enabled
                                                  devices can                 devices need to use
                                                  communicate with            LDP signaling
                                                  PWE3-enabled                messages to exchange
                                                  devices by using LDP        label information
                                                  FEC 128.                    after VPLS member
                                                                              discovery.
                                                                            ● BGP AD VPLS
                                                                              depends on BGP
                                                                              capabilities and its
                                                                              implementation is
                                                                              complex.



                    Table 6-7 Comparison between BGP AD VPLS configuration and LDP VPLS
                    configuration when existing BGP sessions are used after new nodes are added
                     VPLS Mode                  VSI Configurations on       Additional
                                                New Nodes                   Configurations on
                                                                            Existing Nodes

                     LDP VPLS                   vsi-id company1             vsi-id company1
                                                 pwsignal ldp                pwsignal ldp
                                                   vsi-id 2                    vsi-id 2
                                                   peer x.x.x.x                peer y.y.y.y
                                                ...                         ...

                     BGP AD VPLS                vsi-id company1             No additional
                                                 bgp-ad
                                                   vpls-id 10
                                                                            configurations need to
                                                ...                         be performed.
                                                 l2vpn-ad-family
                                                   peer x.x.x.x enable      BGP AD VPLS-enabled
                                                                            devices use extended
                                                                            BGP Update messages
                                                                            carrying VPLS member
                                                                            information to
                                                                            automatically discover
                                                                            VPLS members and use
                                                                            LDP FEC 129 to
                                                                            negotiate and establish
                                                                            VPLS PWs. In this way,
                                                                            VPLS members are
                                                                            automatically discovered
                                                                            and VPLS PWs are
                                                                            automatically
                                                                            established. Therefore,
                                                                            existing VSIs do not need
                                                                            additional
                                                                            configurations.



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            647
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    On the network shown in Figure 6-18, BGP sessions are established between PE1,
                    PE2, and PE3, and BGP AD VPLS is configured on PE1 and PE2, which reside in the
                    same VPLS domain. PE3 needs to be added to the VPLS domain for network
                    expansion. To achieve this goal, it is necessary to configure the ID of the VPLS
                    domain for a VSI on PE3. VPLS configurations do not need to be modified on BGP
                    AD VPLS-enabled PE1 and PE2. After PE3 joins the VPLS domain, PWs can be
                    automatically established using BGP AD between PE1 and PE3 and between PE2
                    and PE3. This simplifies VPLS configuration.

                    Figure 6-18 Full-mesh BGP AD VPLS networking




                    BGP AD also supports HVPLS. The PWs of BGP AD VSIs can be set to spoke PWs so
                    that remote peers of a PE are used as user-side devices on an HVPLS network.

6.8.2 Enabling BGP Peers to Exchange VPLS Information
Context
                    BGP AD VPLS shares a TCP connection with BGP. Most BGP AD VPLS
                    configurations are the same as BGP configurations. A major difference between
                    BGP and BGP AD VPLS is that the latter requires PEs to function as BGP peers to
                    exchange VPLS member information in the L2VPN AD address family view.
                    Perform the following steps on the PEs at both ends of a PW.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

