---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-205
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [29726, 29844]
sha256: c0bd0ea7468b9165751747be29dae246f60e76e929cb1cff2e78abe332721004
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 When the BFD session state changes, the BFD module notifies the application
                 protocol of the change. Fast switching between the primary and backup CR-LSPs
                 can be then performed. If an LSP or a TE tunnel is used as a reverse path to notify
                 the ingress of a fault, the process-pst command can be run to allow the reverse
                 path to switch traffic if the BFD session goes down. If an IP link is used as a
                 reverse path, this command can be executed only when the IP link has only one
                 hop. This command does not apply to a multi-hop IP link used as a reverse path.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              492
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


         Step 6 (Optional) Set local BFD parameters as required. For details, see Table 4-29.
                 You can perform one or more steps in Table 4-29 as required.

                 Table 4-29 Configuring local BFD parameters

                  Operation                     Command                     Purpose

                  Set the minimum               min-tx-interval interval    The default interval is
                  interval at which the                                     1000 milliseconds.
                  local device sends BFD
                  packets.

                  Set the minimum               min-rx-interval interval    The default interval is
                  interval at which the                                     1000 milliseconds.
                  local device receives BFD
                  packets

                  Set the local BFD             detect-multiplier           The default local BFD
                  detection multiplier.         multiplier                  detection multiplier is 3.




                 Actual local interval for sending BFD packets = max {Locally configured interval
                 for sending BFD packets, Remotely configured interval for receiving BFD packets}
                 Actual local interval for receiving BFD packets = max {Remotely configured
                 interval for sending BFD packets, Locally configured interval for receiving BFD
                 packets}
                 Actual local detection interval = Actual local interval for receiving BFD packets x
                 Remotely configured BFD detection multiplier
                 For example:
                 ●      On the local device, the interval for sending BFD packets is set to 200 ms, the
                        interval for receiving BFD packets is set to 300 ms, and the detection
                        multiplier is set to 4.
                 ●      On the remote device, the interval for sending BFD packets is set to 100 ms,
                        the interval for receiving BFD packets is set to 600 ms, and the detection
                        multiplier is set to 5.
                 Then:
                 ●      On the local device, the actual interval for sending BFD packets is 600 ms, as
                        calculated using the formula max {200 ms, 600 ms}; the actual interval for
                        receiving BFD packets is 300 ms, as calculated using the formula max {100
                        ms, 300 ms}; the detection multiplier is 1500 ms, as calculated using the
                        formula 300 ms x 5.
                 ●      On the remote device, the actual interval for sending BFD packets is 300 ms,
                        as calculated using the formula max {100 ms, 300 ms}; the actual interval for
                        receiving BFD packets is 600 ms, as calculated using the formula max {200
                        ms, 600 ms}; the detection multiplier is 2400 ms, as calculated using the
                        formula 600 ms x 4.

                 ----End

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             493
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


4.28.5 Verifying the Configuration
Procedure
                 ●      Run the display bfd session mpls-te interface tunnel tunnel-num
                        [ verbose ] command to check BFD session information on a tunnel ingress.
                 ●      Run the following commands to check BFD session information on a tunnel
                        egress:
                        –   Run the display bfd session all [ for-ip | for-lsp | for-te ] [ verbose ]
                            command to check the configurations of all BFD sessions.
                        –   Run the display bfd session static [ for-ip | for-lsp | for-te ] [ verbose ]
                            command to check the configurations of static BFD sessions.
                        –   Run the display bfd session peer-ip peer-ip [ vpn-instance vpn-name ]
                            [ verbose ] command to check the configuration of a BFD session of
                            which the reverse path is an IP link.
                        –   Run the display bfd session ldp-lsp peer-ip peer-ip [ nexthop nexthop-
                            ip [ interface interface-type interface-number ] ] [ verbose ] command
                            to check information about BFD sessions with reverse LDP LSPs.
                        – Run the display bfd session mpls-te interface tunnel-name te-lsp
                            [ verbose ] command to check the configuration of a BFD session of
                            which the reverse path is a CR-LSP.
                        – Run the display bfd session mpls-te interface tunnel-name [ verbose ]
                            command to check the configuration of a BFD session of which the
                            reverse path is a TE tunnel.
                 ●      Run the following commands to check BFD statistics:
                        – Run the display bfd statistics session all [ for-ip | for-lsp | for-te ]
                            command to check statistics about all BFD sessions.
                        – Run the display bfd statistics session static [ for-ip | for-lsp | for-te ]
                            command to check statistics about static BFD sessions.
                        – Run the display bfd statistics session peer-ip peer-ip [ vpn-instance
                            vpn-instance-name ] command to check statistics about BFD sessions
                            with IP links.
                        – Run the display bfd statistics session ldp-lsp peer-ip peer-ip [ nexthop
                            nexthop-ip [ interface interface-type interface-number ] ] command to
                            check statistics about BFD sessions with LDP LSPs.
                        – Run the display bfd statistics session mpls-te interface tunnel
                            interface-number te-lsp command to check statistics about BFD sessions
                            with CR-LSPs.
                        – Run the display bfd statistics session mpls-te interface tunnel
                            interface-number command to check statistics about BFD sessions with
                            TE tunnels.
                 ----End

4.28.6 Example for Configuring Static BFD for TE Tunnel
Networking Requirements
                 On the MPLS network shown in Figure 4-53, a Layer 2 device (a switch) is
                 deployed between PE1 and PE2. Configure VPN FRR on PE1 and MPLS TE tunnels.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                494

