---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-231
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [34023, 34152]
sha256: 55906e6a780deefdd1b01d20f0f525ed212d0e98046bbfd41f742d493838d4d2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         If two PEs are directly connected, you do not need to configure a remote LDP session
                         between them.

                    # Configure PE1.
                    [PE1] mpls ldp remote-peer 4.4.4.4
                    [PE1-mpls-ldp-remote-4.4.4.4] remote-ip 4.4.4.4
                    [PE1-mpls-ldp-remote-4.4.4.4] quit
                    [PE1] mpls ldp remote-peer 5.5.5.5
                    [PE1-mpls-ldp-remote-5.5.5.5] remote-ip 5.5.5.5
                    [PE1-mpls-ldp-remote-5.5.5.5] quit

                    # Configure PE2.
                    [PE2] mpls ldp remote-peer 1.1.1.1
                    [PE2-mpls-ldp-remote-1.1.1.1] remote-ip 1.1.1.1
                    [PE2-mpls-ldp-remote-1.1.1.1] quit

                    # Configure PE3.
                    [PE3] mpls ldp remote-peer 1.1.1.1
                    [PE3-mpls-ldp-remote-1.1.1.1] remote-ip 1.1.1.1
                    [PE3-mpls-ldp-remote-1.1.1.1] quit

                    After the configurations are complete, run the display mpls ldp session command
                    on PEs. The command outputs show that the Status field displays Operational,
                    indicating that a remote LDP peer relationship has been established.
                    The following example uses the command output on PE1.
                    <PE1> display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     ------------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                     ------------------------------------------------------------------------------


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       542
VPN Configuration
VPN Configuration                                                                                    5 VPWS Configuration

                    2.2.2.2:0        Operational DU Passive 000:00:06 27/27
                    3.3.3.3:0        Operational DU Passive 000:00:05 24/24
                    4.4.4.4:0         Operational DU Passive 000:00:00 3/3
                    5.5.5.5:0        Operational DU Passive 000:00:00 2/2
                    ------------------------------------------------------------------------------
                    TOTAL: 4 session(s) Found.

         Step 5 Configure PWs on PEs using PW templates.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] pw-template 1to2
                    [PE1-pw-template-1to2] peer-address 4.4.4.4
                    [PE1-pw-template-1to2] control-word
                    [PE1-pw-template-1to2] quit
                    [PE1] pw-template 1to3
                    [PE1-pw-template-1to3] peer-address 5.5.5.5
                    [PE1-pw-template-1to3] control-word
                    [PE1-pw-template-1to3] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] mpls l2vc pw-template 1to2 100
                    [PE1-10GE1/0/1] mpls l2vc pw-template 1to3 200 secondary
                    [PE1-10GE1/0/1] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] pw-template 2to1
                    [PE2-pw-template-2to1] peer-address 1.1.1.1
                    [PE2-pw-template-2to1] control-word
                    [PE2-pw-template-2to1] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] mpls l2vc pw-template 2to1 100
                    [PE2-10GE1/0/1] quit

                    # Configure PE3.
                    [PE3] mpls l2vpn
                    [PE3-l2vpn] quit
                    [PE3] pw-template 3to1
                    [PE3-pw-template-3to1] peer-address 1.1.1.1
                    [PE3-pw-template-3to1] control-word
                    [PE3-pw-template-3to1] quit
                    [PE3] interface 10ge 1/0/2
                    [PE3-10GE1/0/2] mpls l2vc pw-template 3to1 200
                    [PE3-10GE1/0/2] quit

                    After the configurations are complete, run the display pw-template command on
                    PEs. The command outputs show the PW template configuration.

                    The following example uses the command output on PE1.
                    <PE1> display pw-template
                     Total PW template number : 2

                     PW Template Name : 1to2
                     PeerIP         : 4.4.4.4
                     Tnl Policy Name : --
                     CtrlWord         : Enable
                     MTU             : 1500
                     Seq-Number          : Disable
                     TDM Encapsulation Number: 32
                     Jitter-Buffer         : 20
                     Jitter-Buffer-Cep        : 1125 Payload-Compression DBA : UNEQ
                     Idle-Code              : ff
                     Rtp-Header               : Disable


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                       543
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration

                     VCCV Capability : cw alert lsp-ping bfd
                     Behavior Name : --
                     Total PW      : 1, Static PW : 0, LDP PW : 1

                     PW Template Name : 1to3
                     PeerIP         : 5.5.5.5
                     Tnl Policy Name : --
                     CtrlWord         : Enable
                     MTU             : 1500
                     Seq-Number          : Disable
                     TDM Encapsulation Number: 32
                     Jitter-Buffer         : 20
                     Jitter-Buffer-Cep        : 1125 Payload-Compression DBA : UNEQ
                     Idle-Code              : ff
                     Rtp-Header               : Disable
                     VCCV Capability : cw alert lsp-ping bfd
                     Behavior Name : --
                     Total PW         : 1, Static PW : 0, LDP PW : 1

