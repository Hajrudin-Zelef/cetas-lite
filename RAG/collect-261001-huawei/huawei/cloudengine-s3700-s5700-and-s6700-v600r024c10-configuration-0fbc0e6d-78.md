---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-78
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [10966, 11114]
sha256: e136215c32e71cdc15c2c2fcc69c0102104d912b61a486554c8740961888687c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 2.     Device A constructs an MPLS Echo Request message, with destination IP
                        address 127.0.0.0/8 and TTL value 1 in the IP header. Then, Device A adds an
                        LSP label (TTL value of the label is 1) to the MPLS Echo Request message and
                        sends the message to Device B. Upon receiving this message, Device B
                        determines that the TTL of the LSP label has timed out and replies with an
                        MPLS Echo Reply message. In the MPLS Echo Reply message, the destination
                        UDP port number and destination IP address are the source UDP port number
                        and source IP address in the MPLS Echo Request message, respectively. The
                        TTL value is 255.
                 3.     Upon receiving the MPLS Echo Reply message sent by Device B, Device A
                        sends an MPLS Echo Request message with the TTL value of 2. Device B then
                        forwards the message using MPLS. After receiving this message, Device C
                        determines that the TTL of the LSP label has timed out and replies with an
                        MPLS Echo Reply message.
                 4.     Upon receiving the MPLS Echo Reply message sent by Device C, Device A
                        sends an MPLS Echo Request message with the TTL value of 3. Device B and
                        Device C then forward the MPLS Echo Request message using MPLS. After
                        receiving this message, Device D determines itself as the egress node and
                        replies with an MPLS Echo Reply message.

                 Figure 3-37 MPLS network




Procedure
                 ●      Configure an LSP trace test to check the path over which an LDP tunnel that
                        carries IPv4 packets is established or locate the failure point on the path.
                        tracert lsp [ -a source-ip | -exp exp-value | -h ttl-value | -r reply-mode | -t time-out | -s size | -g ] * ip
                        destination-iphost mask-length [ ip-address ] [ nexthop nexthopAddr ] [ detail ]

                               NOTE

                             If FEC is not required for LDP services, you can run the lspv echo-reply fec-validation
                             ldp disable command to disable FEC.

                 ----End

Example
                 ●      Perform the LSP trace test to check the path over which an LDP tunnel is
                        established or locate the failure point on the path.
                        <HUAWEI> tracert lsp ip 1.1.1.1 32
                         LSP Trace Route FEC: IPV4 PREFIX 1.1.1.1/32 , press CTRL_C to break.
                         TTL Replier          Time Type       Downstream
                         0                      Ingress 10.1.1.1/[3 ]
                         1    1.1.1.1       5     Egress




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                      185
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


Follow-up Procedure
                 ●      Check LSPV packet statistics.
                        display lspv statistics

                        If the test performed using the tracert lsp command fails, you can run this
                        command to check whether the fault occurs on the LSP or the device.
                 ●      Clear LSPV packet statistics.
                        reset lspv statistics

                 ●      Disable devices from responding to MPLS Echo Request messages.
                        lspv mpls-lsp-ping echo disable

                        After the test is complete, you are advised to run this command to disable
                        devices from responding to MPLS Echo Request messages to prevent system
                        resource consumption.


3.25 Maintaining MPLS LDP

3.25.1 Restarting LDP

Context


                        NOTICE

                 ● Restarting LDP affects the establishment of an LSP. Exercise caution when
                   performing this operation.
                 ● Restarting LDP is prohibited during the LDP GR.



Procedure
                 ●      To restart the global LDP instance, run the reset mpls ldp command in the
                        user view, allowing new configurations to take effect
                 ●      To restart all LDP instances, run the reset mpls ldp all command, making
                        new configurations take effect.
                 ●      To restart a specified LDP peer, run the reset mpls ldp peer peer-id command
                        in the user view, allowing new configurations to take effect.
                 ●      When the global LDP instance on a GR-capable device needs to be restarted
                        and service interruption should be prevented, run the reset mpls ldp graceful
                        command in the user view to restart all peers in the global LDP instance,
                        allowing new configurations to take effect.
                 ●      When a specified LDP peer needs to be restarted and service interruption
                        should be prevented, run the reset mpls ldp peer peer-id graceful command
                        in the user view to restart the specified LDP peer, allowing new configurations
                        to take effect.

                 ----End



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           186
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


3.25.2 Monitoring the LDP Running Status
Context
                 In routine maintenance, you can run the following commands in any view to view
                 the LDP running status.

Procedure
                 ●      Run the display mpls ldp error packet { tcp [ peer peer-id ] | udp
                        [ interface interface-type interface-num ]} command to check statistics and
                        content of error TCP and UDP packets received by LDP.
                 ●      Run the display mpls ldp event adjacency-down [ interface interface-type
                        interface-number | remote ] [ peer peer-id ] [ verbose ] command to check
                        information about LDP peers in the down state.
                 ●      Run the display mpls ldp event session-down [ peer-id ] [ verbose ]
                        command to check the time and reason when an LDP session goes down.
                 ●      Run the display default-parameter mpls management command to check
                        the default configurations of MPLS management.
                 ●      Run the display mpls interface command to check information about MPLS-
                        enabled interfaces.
                 ●      Run the display mpls label all summary command to check the allocation
                        information of all MPLS labels.
                 ●      Run the display mpls label dynamic available command to check the labels
                        that can be used by dynamic services.
                 ●      Run the display mpls lsp protocol ldp command to check LDP LSP
                        information.
                 ●      Run the display mpls lsp statistics command to check the number of LSPs in
                        the Up state and the number of activated LSPs on the ingress, transit node,
                        and egress.
                 ----End

3.25.3 Clearing LDP Statistics
Context


                        NOTICE

