---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-77
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10841, 10965]
sha256: f9b9acb295efd5310641be888ef14d9751c427e4ce542f4c8c17a9991dec80d2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
                 ●      Configure an LSP ping test to check the LDP tunnel that carries IPv4 packets.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             182
MPLS Configuration
MPLS Configuration                                                                                3 MPLS LDP Configuration

                        ping lsp [ -a source-ip | -c count | -exp exp-value | -h ttl-value | -m interval | -r reply-mode | -s
                        packet-size | -t time-out | -v | -g ] * ip destination-iphost mask-length [ ip-address ] [ nexthop
                        nexthop-address ]

                               NOTE

                             If FEC is not required for LDP services, you can run the lspv echo-reply fec-validation
                             ldp disable command to disable FEC.

                 ----End

Example
                 ●      Perform the LSP ping test to check the connectivity of the LDP tunnel.
                        <HUAWEI> ping lsp -v ip 3.3.3.3 32
                         LSP PING FEC: IPV4 PREFIX 3.3.3.3/32 : 100 data bytes, press CTRL_C to break
                          Reply from 3.3.3.3: bytes=100 Sequence=1 time = 4 ms Return Code 3, Subcode 1
                          Reply from 3.3.3.3: bytes=100 Sequence=2 time = 4 ms Return Code 3, Subcode 1
                          Reply from 3.3.3.3: bytes=100 Sequence=3 time = 4 ms Return Code 3, Subcode 1
                          Reply from 3.3.3.3: bytes=100 Sequence=4 time = 4 ms Return Code 3, Subcode 1
                          Reply from 3.3.3.3: bytes=100 Sequence=5 time = 5 ms Return Code 3, Subcode 1

                         --- FEC: IPV4 PREFIX 3.3.3.3/32 ping statistics ---
                           5 packet(s) transmitted
                           5 packet(s) received
                           0.00% packet loss
                           round-trip min/avg/max = 4/4/5 ms

                 In the command output:
                 ●      If still no response message is received on the source after the timer expires,
                        the message "Request time out" is displayed. If a response message is
                        received before the timer expires, information such as the number of data
                        bytes, packet sequence number, and response time is displayed.
                 ●      The statistics include the numbers of sent and received packets, the
                        percentage of packets that are not replied to, and the minimum, maximum,
                        and average response time.

Follow-up Procedure
                 ●      Check LSPV packet statistics.
                        display lspv statistics

                        If the test performed using the ping lsp command fails, you can run this
                        command to check whether the fault occurs on the LSP or the device.
                 ●      Clear LSPV packet statistics.
                        reset lspv statistics

                 ●      Disable devices from responding to MPLS Echo Request messages.
                        lspv mpls-lsp-ping echo disable

                        After the test is complete, you are advised to run this command to disable
                        devices from responding to MPLS Echo Request messages to prevent system
                        resource consumption.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                    183
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


3.24.2 Checking MPLS Network Path Information Using
Tracert
Prerequisites
                         NOTE

                        This feature is supported only by the following series: S6780-H, S6750-H, S6750E-S, S6750-
                        S, S5755E-H, S5755-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2
                 ●      The MPLS network has been correctly configured before MPLS tracert is used
                        to check the LSP information or locate an LSP fault.
                 ●      The undo lspv mpls-lsp-ping echo disable command has been run to enable
                        a device to respond to MPLS Echo Request messages.
                 ●      If the device interworks with a non-Huawei device, you need to run the lspv
                        echo-reply compitable fec enable command to enable the device to respond
                        to MPLS Echo Request messages with MPLS Echo Reply messages that do not
                        carry FEC information.
                 ●      Both the initiator and responder send LSP trace packets to the main control
                        unit for processing. If a large number of packets are sent to the main control
                        unit, the CPU usage of the main control unit surges, affecting normal device
                        running. To prevent this problem, you can run the lspv mpls-lsp-ping cpu-
                        defend cpu-defend command to limit the rate at which MPLS Echo Request
                        messages are sent to the main control unit.

Context
                 On an MPLS network, the MPLS control plane for establishing label switched
                 paths (LSPs) cannot detect data forwarding failures in LSPs, complicating network
                 maintenance. To address this, MPLS ping and tracert can be used to detect LSP
                 faults and quickly locate faulty nodes. MPLS tracert is mainly used to check
                 network connectivity and locate faults on the network.

                 MPLS tracert checks the LSP status by sending MPLS Echo Request and Reply
                 messages. The two types of messages are encapsulated into UDP packets and
                 transmitted using port 3503. The receiver identifies the MPLS Echo Request and
                 Reply messages based on the UDP port number.

                 An MPLS Echo Request message carries information about the forwarding
                 equivalence class (FEC) for an LSP to be checked. The MPLS Echo Request
                 message is forwarded along the LSP with other packets belonging to this FEC. This
                 procedure enables the LSP connectivity to be checked. MPLS Echo Request
                 messages are forwarded to the destination using MPLS, whereas MPLS Echo Reply
                 messages are forwarded to the source using IP.

                 To prevent the egress from forwarding a received MPLS Echo Request message to
                 other nodes, you can set the destination IP address to 127.0.0.0/8 (the local
                 loopback address) and the TTL value to 1 in the IP header of the message.

                 In Figure 3-37, the MPLS tracert process of 4.4.4.4/32 on Device A is as follows:
                 1.     Device A checks whether the LSP is established. If not, an error code is
                        returned and the MPLS tracert process ends. If the LSP is established, the
                        MPLS tracert process continues.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     184
MPLS Configuration
MPLS Configuration                                                                                 3 MPLS LDP Configuration


