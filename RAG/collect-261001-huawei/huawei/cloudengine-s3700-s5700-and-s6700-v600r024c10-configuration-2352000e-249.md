---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-249
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [36580, 36716]
sha256: 73cd4a807b95f447477dd87652d7c4d3aac05d2c043f627554178c28bcf12224
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The command output contains the following information:

                    ●   Response to each ping packet: If no response packet is received within the
                        timeout period, the message "Request time out" is displayed. If a response
                        packet is received, the number of data bytes, packet sequence number, TTL
                        value, and response time carried in the packet are displayed.
                    ●   Statistics about ping packets: These include the number of sent packets,
                        number of received packets, percentage of the packets that are not replied,
                        and the minimum, maximum and average response time.

5.13.5 Using Tracert to Test a PW Path on a VPWS Network

Prerequisites
                    Before testing PW connectivity using the tracert vc command, ensure that the
                    VPWS network has been configured correctly.

                    To use the control word channel, you need to run the control-word command in
                    the PW template view to enable the control word function.


Context
                    VPWS PW tracert can be used to test single-segment PWs.



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                            580
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


                    Figure 5-41 Typical single-segment VPWS network




                    On the network shown in Figure 5-41, UPE1 performs a single-segment VPWS PW
                    tracert operation on PW100. The tracert process is as follows:
                    1.   UPE1 checks whether PW100 exists. If not, an error message is displayed, and
                         the tracert process ends. If PW100 exists, the tracert process continues.
                    2.   UPE1 constructs an MPLS Echo Request message, with destination IP address
                         127.0.0.1/8 and TTL value 1 in the IP header. It then searches for the
                         corresponding LSP, adds an LSP label to the MPLS Echo Request message, and
                         sends the message to P.
                    3.   After receiving this message, P determines that the TTL of the LSP label has
                         timed out and replies with an MPLS Echo Reply message. In the MPLS Echo
                         Reply message, the destination UDP port number and destination IP address
                         are the source UDP port number and source IP address in the MPLS Echo
                         Request message, respectively. The TTL value is 255.
                    4.   After receiving the MPLS Echo Reply message, UPE1 sends an MPLS Echo
                         Request message with the TTL value of 2. P forwards the message using MPLS
                         to UPE2. After receiving this message, UPE2 determines itself as the egress
                         node and replies with an MPLS Echo Reply message.
                    Table 5-10 describes three test modes that can be used by VPWS PW tracert to
                    enable the destination device (UPE2) to forward packets to the CPU for
                    processing, instead of forwarding them to other devices.

Table 5-10 Three test modes
 Test Mode                        Description

 Control-word                     ● The request packet carries the control word between the PW
                                    label and the IP header. After receiving the packet with the
                                    control word, the destination device sends the packet to the CPU
                                    for processing.
                                  ● This test mode applies only when you enable the control word
                                    function on both ends of the PW.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          581
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration


 Test Mode                          Description

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
                        tracert vc vc-type pw-id [ peer-address ] [ -exp exp-value | -f first-ttl | -m max-ttl | -r reply-mode | -t
                        timeout-value | -g ] * control-word [ full-lsp-path ] [ pipe | uniform ] [ detail ]

                        BGP VPWS:
                        tracert vc -vpn-instance vpn-instance-name local-ce-id remote-ce-id [ -exp exp-value | -f firstTtl | -
                        m maxTtl | -r reply-mode | -t timeout | -g ] * control-word [ full-lsp-path ] [ detail ]

                    ●   Label-alert mode:
                        LDP VPWS:
                        tracert vc vc-type pw-id [ peer-address ] [ -exp exp-value | -f first-ttl | -m max-ttl | -r reply-mode | -t
                        timeout-value | -g ] * label-alert [ no-control-word ] [ full-lsp-path ] [ pipe | uniform ] [ detail ]

                        BGP VPWS:
                        tracert vc -vpn-instance vpn-instance-name local-ce-id remote-ce-id [ -exp exp-value | -f firstTtl | -
                        m maxTtl | -r reply-mode | -t timeout | -g ] * label-alert [ no-control-word ] [ full-lsp-path ]
                        [ detail ]

                    ●   TTL mode:
                        tracert vc vc-type pw-id [ peer-address ] [ -exp exp-value | -f first-ttl | -m max-ttl | -r reply-mode | -t
                        timeout-value | -g ] * normal [ no-control-word ] [ full-lsp-path ] [ pipe | uniform ] [ detail ]

                    ----End

Example
                    ●   Perform a tracert operation to test a PW path on an LDP VPWS network in
                        control-word scenarios.
                        <HUAWEI> tracert vc vlan 123 3.3.3.9 control-word
                        PW Trace Route : FEC 128 PSEUDOWIRE (NEW). Type = vlan, ID = 123, press CTRL_C to break
                        TTL Replier         Time Type      Downstream
                        0                     Ingress 192.168.1.2/[2140 ]
                        1   3.3.3.3        2 ms Egress

                    ●   Perform a tracert operation to test a PW path on a BGP VPWS network in
                        label-alert scenarios.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                     582
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration

