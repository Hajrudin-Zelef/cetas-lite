---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-261
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [38282, 38404]
sha256: 86c95b1281c260fb5e1197ca2e7a12e94b1576d91e426cfa79f3dfaf13a37866
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

6.6.4 (Optional) Configuring VPLS Flow Label-based Load
Balancing
Context
                    Packets of multiple data flows on the same PW carry the same VC labels which
                    are encapsulated on a PE. When these packets arrive at a P, the P forwards them
                    along only one path, regardless of whether multiple load-balancing paths exist.
                    To load-balance different data flows, configure flow label-based load balancing on
                    the PE. After the configurations are complete, the PE adds flow labels following
                    private network labels when encapsulating data packets, and the P load-balances
                    these data flows based on the flow labels.
                    On the network shown in Figure 6-14, two different data flows need to be
                    transmitted to the peer end through an L2VPN.
                    1.   PE1 calculates flow labels based on the source and destination IP addresses of
                         the two data flows. The following assumes that the flow labels are 1025 and
                         1026.
                    2.   PE1 adds the flow labels following private network labels in the packets of
                         the two data flows.
                    3.   When the two data flows reach P1, P1 performs a hash calculation based on
                         the flow labels and maps the two data flows onto different paths. The
                         following assumes that the next hop of the data flow with flow label 1025 is
                         P2, and the next hop of the data flow with flow label 1026 is P3.
                    4.   When the two data flows reach PE2, the private network labels and flow
                         labels are sequentially removed. PE2 then forwards the two data flows to
                         their destination CEs based on their private network labels.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     612
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration


                    Figure 6-9 Flow label-based load balancing




                    Flow label-based load balancing applies to an L2VPN on which multiple links exist
                    between Ps. Flow label-based load balancing allows data flows on the same VPN
                    to be load-balanced among different paths based on flow labels for higher
                    network resource utilization.

Procedure
                    ●   Enable flow label-based load balancing for all PWs in a VSI.

                        Perform the following steps on the PEs at both ends of a PW.

                        a.   Enter the system view.
                             system-view

                        b.   Enter the VSI view.
                             vsi vsi-name

                        c.   Enter the VSI-LDP view.
                             pwsignal ldp

                        d.   Enable flow label-based load balancing for the VSI.
                             flow-label { both | send | receive } [ static ]

                             By default, flow label-based load balancing is disabled for a VSI.

                                   NOTE

                                  ● Flow label-based load balancing can be enabled only when any of the
                                    following conditions is true:
                                      ●     The receive parameter is configured on the local end, and the send
                                            parameter is configured on the remote end.
                                      ●     The send parameter is configured on the local end, and the receive
                                            parameter is configured on the remote end.
                                      ●     Both the send and receive parameters are configured on the local and
                                            remote ends.
                                  ● If static is configured, flow label-based load balancing is statically configured.
                                      If the static flow label-based load balancing configuration does not match on
                                      both ends, the device discards packets with flow labels, causing packet loss.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        613
VPN Configuration
VPN Configuration                                                                              6 VPLS Configuration


                    ●   Enable flow label-based load balancing for a specified PW in a VSI.

                        Perform the following steps on the PEs at both ends of a PW.

                        a.    Enter the system view.
                              system-view

                        b.    Enter the VSI view.
                              vsi vsi-name

                        c.    Enter the VSI-LDP view.
                              pwsignal ldp

                        d.    Configure a VSI peer and enter the VSI-LDP-PW view.
                              peer peer-address [ negotiation-vc-id vc-id ] pw pw-name

                        e.    (Optional) Disable flow label-based load balancing for the PW.
                              flow-label disable

                              By default, a PW has the same flow label-based load balancing capability
                              as the VSI to which it belongs.

                                    NOTE

                                   ● If flow label-based load balancing is configured for both a VSI and a PW in
                                     this VSI, the configuration of the PW takes precedence.
                                   ● If the flow-label command has been run for a VSI but not for a PW, you can
                                     run the flow-label disable command in the VSI-LDP-PW view to disable flow
                                     label-based load balancing of the PW. If flow label-based load balancing for
                                     the PW is required again, you can run the undo flow-label disable command
                                     to enable flow label-based load balancing for the PW.
                                   ● If the undo flow-label command is run in the VSI-LDP-PW view, only flow
                                     label-based load balancing that has been configured using the flow-label
                                     command in the VSI-LDP-PW view is disabled.
                        f.    Enable flow label-based load balancing for the PW.
                              flow-label { both | send | receive } [ static ]

                              By default, a PW has the same flow label-based load balancing capability
                              as the VSI to which it belongs.

                                    NOTE

