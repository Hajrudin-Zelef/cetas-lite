---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-216
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [31421, 31585]
sha256: 1dba462be69e82d4a5a90050fcb7a081e5f9526dbcdd929a32acdeab0d41a279
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        After the test is complete, you are advised to run this command to disable
                        devices from responding to MPLS Echo Request messages to prevent system
                        resource consumption.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                     518
MPLS Configuration
MPLS Configuration                                                                                   4 MPLS TE Configuration


4.31.2 Verifying the Path of a TE Tunnel Using Tracert
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
                 If a ping failure occurs on a TE tunnel, run the tracert lsp command to locate the
                 failure point.

Procedure
                 ●      Configure an LSP trace test to check the path over which a TE tunnel that
                        carries IPv4 packets is established or locate the failure point on the path.
                        tracert lsp [ -a source-ip | -exp exp-value | -h ttl-value | -r reply-mode | -t time-out | -s size | -g ] * te
                        { tunnelName | ifType ifNum } [ hot-standby | primary ] [ compatible-mode ] [ detail ]

                 ----End

Example
                 ●      Perform the LSP trace test to check the path over which a TE tunnel is
                        established or locate the failure point on the path.
                        <HUAWEI> tracert lsp te Tunnel 1
                         LSP Trace Route FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 , press CTRL_C to break.
                         TTL Replier          Time Type      Downstream
                         0                      Ingress 10.1.1.1/[3 ]
                         1    1.1.1.1       4     Egress


Follow-up Procedure
                 ●      Check LSPV packet statistics.
                        display lspv statistics

                        If the test performed using the tracert lsp command fails, you can run this
                        command to check whether the fault occurs on the LSP or the device.
                 ●      Clear LSPV packet statistics.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                                      519
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration

                        reset lspv statistics
                 ●      Disable devices from responding to MPLS Echo Request messages.
                        lspv mpls-lsp-ping echo disable
                        After the test is complete, you are advised to run this command to disable
                        devices from responding to MPLS Echo Request messages to prevent system
                        resource consumption.


4.32 Maintaining MPLS TE

4.32.1 Resetting MPLS TE
Context
                 During routine maintenance, you can reset MPLS TE, which includes resetting
                 tunnel interfaces and the RSVP-TE process.


                        NOTICE

                 Resetting the RSVP-TE process causes all RSVP CR-LSPs to be torn down and re-
                 established.


Procedure
                 ●      Run the reset mpls te tunnel-interface tunnel interface-number command
                        to reset a specified MPLS TE tunnel interface.
                 ●      Run the reset mpls rsvp-te command to reset the RSVP-TE process.
                        To re-establish all RSVP-TE CR-LSPs or verify the RSVP-TE working process, you
                        can reset the RSVP-TE process.
                 ----End

4.32.2 Clearing RSVP-TE Statistics
Context
                 During routine maintenance, you can clear RSVP-TE statistics.


                        NOTICE

                 RSVP-TE statistics cannot be restored after they are cleared. Exercise caution when
                 performing this operation.


Procedure
                 ●      Run the reset mpls rsvp-te statistics { global | interface { interface-type
                        interface-number | interface-name} }command to clear RSVP-TE statistics.
                 ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           520
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


4.32.3 Monitoring TE
Context
                 In routine maintenance, you can run the following commands in any view to learn
                 the TE operating status.

Procedure
                 ●      Run the display default-parameter mpls te management command to
                        check the default settings of MPLS TE management parameters.
                 ●      Run the display mpls te tunnel statistics command to check tunnel
                        statistics.
                 ●      Run the display mpls te tunnel-interface last-error [ ] command to check
                        the latest errors that occur on a tunnel interface of an ingress.
                 ●      Run the display mpls te tunnel-interface failed command to check the
                        MPLS TE tunnels that fail to be established or are being established.
                 ●      Run the display mpls rsvp-te statistics { global | interface { interface-type
                        interface-number | interface-name } } command to check RSVP-TE statistics.
                 ----End

4.32.4 Deleting and Re-establishing Automatically Generated
Bypass Tunnels
Context
                 If MPLS TE Auto FRR is enabled, you can run the related command to delete and
                 re-establish a bypass tunnel.

Procedure
                 ●      Run the reset mpls te auto-frr { lsp-id ingress-lsr-id tunnel-id | name
                        bypass-tunnel-name } command to delete and re-establish an automatically
                        generated bypass tunnel.

                 ----End


4.33 Troubleshooting MPLS TE

4.33.1 An MPLS TE Tunnel's State Is Down
Fault Symptom
                 After an MPLS TE tunnel is established, its state is down.

Possible Causes
                 ●      The tunnel configuration is incomplete.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           521
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


