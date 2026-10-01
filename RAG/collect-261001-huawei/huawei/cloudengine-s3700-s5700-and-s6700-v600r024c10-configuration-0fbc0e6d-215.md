---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-215
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [31249, 31420]
sha256: 7eda3b3b28369124d983b1067e1dab43c92d42a9700e65e8ccf3e47b6fd0eed9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      LSR5
                        #
                        sysname LSR5
                        #
                        vlan batch 400 500
                        #
                        mpls lsr-id 5.5.5.9
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 5.5.5.9 0.0.0.0
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR6


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      516
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration

                        #
                        sysname LSR6
                        #
                        vlan batch 600 700
                        #
                        mpls lsr-id 6.6.6.9
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif600
                         ip address 10.1.6.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif700
                         ip address 10.1.7.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 700
                        #
                        interface LoopBack1
                         ip address 6.6.6.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 6.6.6.9 0.0.0.0
                          network 10.1.6.0 0.0.0.255
                          network 10.1.7.0 0.0.0.255
                          mpls-te enable
                        #
                        return



4.31 Verifying TE Tunnels Using Ping and Tracert
4.31.1 Verifying the Connectivity of a TE Tunnel Using Ping
Prerequisites
                         NOTE

                        This feature is supported only by the following series: S6780-H, S6750-H, S6750E-S, S6750-
                        S, S5755E-H, S5755-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2
                 ●      The MPLS network has been correctly configured before MPLS ping is used to
                        check the connectivity of an LSP.
                 ●      The undo lspv mpls-lsp-ping echo disable command has been run to enable
                        a device to respond to MPLS Echo Request messages.
                 ●      Both the initiator and responder send LSP ping packets to the main control
                        unit for processing. If a large number of packets are sent to the main control
                        unit, the CPU usage of the main control unit surges, affecting normal device
                        running. To prevent this problem, you can run the lspv mpls-lsp-ping cpu-

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    517
MPLS Configuration
MPLS Configuration                                                                                4 MPLS TE Configuration


                        defend cpu-defend command to limit the rate at which MPLS Echo Request
                        messages are sent to the main control unit.

Context
                 After configuring an MPLS TE tunnel, you can run the ping lsp command on the
                 tunnel ingress to check whether the egress can be pinged.

Procedure
                 ●      Configure an LSP ping test to check the TE tunnel that carries IPv4 packets.
                        ping lsp [ -a source-ip | -c count | -exp exp-value | -h ttl-value | -m interval | -r reply-mode | -s
                        packet-size | -t time-out | -v | -g ] * te { { tunnelName | ifType ifNum } [ hot-standby | primary ]
                        [ compatible-mode ] }

                 ----End

Example
                 ●      Perform the LSP ping test to check the connectivity of the TE tunnel.
                        <HUAWEI> ping lsp te Tunnel 1
                         LSP PING FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 : 100 data bytes, press CTRL_C to break
                          Reply from 1.1.1.1: bytes=100 Sequence=1 time = 4 ms
                          Reply from 1.1.1.1: bytes=100 Sequence=2 time = 2 ms
                          Reply from 1.1.1.1: bytes=100 Sequence=3 time = 2 ms
                          Reply from 1.1.1.1: bytes=100 Sequence=4 time = 2 ms
                          Reply from 1.1.1.1: bytes=100 Sequence=5 time = 2 ms

                         --- FEC: RSVP IPV4 SESSION QUERY Tunnel1 ping statistics ---
                           5 packet(s) transmitted
                           5 packet(s) received
                           0.00% packet loss
                           round-trip min/avg/max = 2/2/4 ms

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

