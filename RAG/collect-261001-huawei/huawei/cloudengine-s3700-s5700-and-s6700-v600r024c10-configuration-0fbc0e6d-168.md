---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-168
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24354, 24474]
sha256: 2c01fd38284006556c857ce843c1d1c3d4cdc30ad303fdc1631b141f8f7df18f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 can be then performed. If an LSP or a TE tunnel is used as a reverse path to notify
                 the ingress of a fault, the process-pst command can be run to allow the reverse
                 path to switch traffic if the BFD session goes down. If an IP link is used as a
                 reverse path, this command can be executed only when the IP link has only one
                 hop. This command does not apply to a multi-hop IP link used as a reverse path.
         Step 6 (Optional) Set local BFD parameters as required. For details, see Table 4-21.
                 You can perform one or more steps in Table 4-21 as required.

                 Table 4-21 Configuring local BFD parameters
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

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             405
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                        receiving BFD packets is 600 ms, as calculated using the formula max {200
                        ms, 600 ms}; the detection multiplier is 2400 ms, as calculated using the
                        formula 600 ms x 4.

                 ----End

4.24.5 Verifying the Configuration
Procedure
                 ●      Run the display bfd session mpls-te interface tunnel-name te-lsp
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
                        –   Run the display bfd session ldp-lsp peer-ip ip-address [ nexthop
                            nexthop-ip [ interface interface-type interface-number ] ] [ verbose ]
                            command to check the configuration of a BFD session of which the
                            reverse path is an LDP LSP.
                        –   Run the display bfd session mpls-te interface tunnel-name te-lsp
                            [ verbose ] command to check the configuration of a BFD session of
                            which the reverse path is a CR-LSP.
                        –   Run the display bfd session mpls-te interface tunnel-name [ verbose ]
                            command to check the configuration of a BFD session of which the
                            reverse path is a TE tunnel.
                 ●      Run the following commands to check BFD statistics:
                        –   Run the display bfd statistics session all [ for-ip | for-lsp | for-te ]
                            command to check statistics about all BFD sessions.
                        –   Run the display bfd statistics session static [ for-ip | for-lsp | for-te ]
                            command to check statistics about static BFD sessions.
                        –   Run the display bfd statistics session peer-ip peer-ip [ vpn-instance
                            vpn-name ] command to check statistics about BFD sessions with IP links.
                        –   Run the display bfd statistics session ldp-lsp peer-ip peer-ip [ nexthop
                            nexthop-ip [ interface interface-type interface-number ] ] command to
                            check statistics about BFD sessions with LDP LSPs.
                        –   Run the display bfd statistics session mpls-te interface interface-type
                            interface-number te-lsp command to check statistics about BFD sessions
                            with CR-LSPs.
                 ----End




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                406
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


4.24.6 Example for Configuring Static BFD for CR-LSP

Networking Requirements
                 On the MPLS network shown in Figure 4-37, establish a TE tunnel from LSR1 to
                 LSR3 and configure CR-LSP hot standby and best-effort path creation. Here:

                 ●      The path of the primary CR-LSP is LSR1 -> LSR2 -> LSR3.
                 ●      The path of the backup CR-LSP is LSR1 -> LSR4 -> LSR3.

