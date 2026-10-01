---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-248
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [36443, 36579]
sha256: 57328b6d970df621420eb6ccea0c52874d04ba4a5c27f0d64a262e106ec7eccc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

5.13.4 Using Ping to Test PW Connectivity on a VPWS
Network
Prerequisites
                    Before testing pseudo wire (PW) connectivity using the ping vc command, ensure
                    that the virtual private wire service (VPWS) network has been configured correctly.
                    To use the control word channel, you need to run the control-word command in
                    the PW template view to enable the control word function.

Context
                    VPWS PW ping can be used to test single-segment PWs.

                    Figure 5-40 Typical single-segment VPWS network




                    On the network shown in Figure 5-40, UPE1 performs a single-segment VPWS PW
                    ping operation on PW100. The ping process is as follows:
                    1.   UPE1 checks whether PW100 exists. If not, an error message is displayed, and
                         the ping process ends. If PW100 exists, the ping process continues.
                    2.   UPE1 constructs an MPLS Echo Request message, with destination IP address
                         127.0.0.1/8 and TTL value 1 in the IP header. It then searches for the
                         corresponding LSP, adds an LSP label to the MPLS Echo Request message, and
                         sends the message to P.
                    3.   The transit node P forwards the MPLS Echo Request message using MPLS to
                         UPE2.
                    4.   If the forwarding is successful, the MPLS Echo Request message reaches the
                         egress node UPE2, which replies with an MPLS Echo Reply message.
                    Table 5-9 describes three test modes that can be used by VPWS PW ping to
                    enable the destination device (UPE2) to forward packets to the CPU for
                    processing, instead of forwarding them to other devices.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          578
VPN Configuration
VPN Configuration                                                                                     5 VPWS Configuration


Table 5-9 Three test modes

 Test Mode                          Description

 Control-word                       ● The request packet carries the control word between the PW
                                      label and the IP header. After receiving the packet with the
                                      control word, the destination device sends the packet to the CPU
                                      for processing.
                                    ● This test mode applies only when you enable the control word
                                      function on both ends of the PW.

 Label-alert                        ● The request packet carries the label-alert tag between the PW
                                      label and the LSP label. After receiving the packet with the label-
                                      alert tag, the destination device sends the packet to the CPU for
                                      processing.
                                    ● This test mode applies only when you enable the label alert
                                      function on both ends of the PW.
                                    ● In multi-segment PW scenarios, this test mode cannot be used
                                      because SPEs do not support packet forwarding through
                                      software.

 TTL                                You must specify a TTL value to ensure that the destination device
                                    sends the packet to the CPU for processing when the TTL expires.




Procedure
                    ●   Control-word mode:

                        LDP VPWS:
                        ping vc vc-type pw-id [ peer-address ] [ -c echo-number | -m time-value | -s data-bytes | -t timeout-
                        value | -exp exp-value | -r reply-mode | -v | -g ] * control-word [ ttl ttl-value ] [ pipe | uniform ]

                        BGP VPWS:
                        ping vc vpn-instance vpn-instance-name local-ce-id remote-ce-id [ -c echo-number | -m time-value |
                        -s data-bytes | -t timeout-value | -exp exp-value | -r reply-mode | -v | -g ] * control-word

                    ●   Label-alert mode:

                        LDP VPWS:
                        ping vc vc-type pw-id [ peer-address ] [ -c echo-number | -m time-value | -s data-bytes | -t timeout-
                        value | -exp exp-value | -r reply-mode | -v | -g ] * label-alert [ no-control-word ]

                        BGP VPWS:
                        ping vc vpn-instance vpn-instance-name local-ce-id remote-ce-id [ -c echo-number | -m time-value |
                        -s data-bytes | -t timeout-value | -exp exp-value | -r reply-mode | -v | -g ] * label-alert [ no-control-
                        word ]

                    ●   TTL mode:
                        ping vc vc-type pw-id [ peer-address ] [ -c echo-number | -m time-value | -s data-bytes | -t timeout-
                        value | -exp exp-value | -r reply-mode | -v | -g ] * normal [ no-control-word ] [ ttl ttl-value ] [ pipe |
                        uniform ]

                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                   579
VPN Configuration
VPN Configuration                                                                               5 VPWS Configuration


Example
                    ●   Perform a ping operation to test LDP VPWS PW connectivity in control-word
                        scenarios.
                        <HUAWEI> ping vc ethernet 100 control-word
                           PW PING : FEC 128 PSEUDOWIRE (NEW). Type = ethernet, ID = 100 : 100 data bytes, press CTRL_C
                        to break
                            Reply from 10.10.10.10: bytes=100 Sequence=1 time = 140 ms
                            Reply from 10.10.10.10: bytes=100 Sequence=2 time = 40 ms
                            Reply from 10.10.10.10: bytes=100 Sequence=3 time = 30 ms
                            Reply from 10.10.10.10: bytes=100 Sequence=4 time = 50 ms
                            Reply from 10.10.10.10: bytes=100 Sequence=5 time = 50 ms
                        - -- FEC: FEC 128 PSEUDOWIRE (NEW). Type = ethernet, ID = 100 ping statistics ---
                            5 packet(s) transmitted
                            5 packet(s) received
                            0.00% packet loss
                            round-trip min/avg/max = 30/62/140 ms

                    ●   Perform a ping operation to test BGP VPWS PW connectivity in label-alert
                        scenario.
                        <HUAWEI> ping vc vpn-instance vpn1 1 2 -v label-alert
                        PW PING : FEC 128 PSEUDOWIRE (NEW). Type = vlan, ID = 100 : 100 data bytes, press CTRL_C to
                        break
                           Reply from 4.4.4.4: bytes=100 Sequence=1 time = 110 ms Return Code 3, Subcode 1
                           Reply from 4.4.4.4: bytes=100 Sequence=2 time = 90 ms Return Code 3, Subcode 1
                           Reply from 4.4.4.4: bytes=100 Sequence=3 time = 60 ms Return Code 3, Subcode 1
                           Reply from 4.4.4.4: bytes=100 Sequence=4 time = 60 ms Return Code 3, Subcode 1
                           Reply from 4.4.4.4: bytes=100 Sequence=5 time = 90 ms Return Code 3, Subcode 1

                         --- FEC: L2 VPN ENDPOINT. Sender VEID = 1, Remote VEID = 2 ping statistics ---
                           5 packet(s) transmitted
                           5 packet(s) received
                           0.00% packet loss
                           round-trip min/avg/max = 60/82/110 ms

