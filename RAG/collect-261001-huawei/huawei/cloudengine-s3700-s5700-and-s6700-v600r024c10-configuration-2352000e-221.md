---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-221
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32564, 32685]
sha256: 144a80407495f2c0da3bd86b9b082a7f1bec09dcc53d6d0db3da7e0bc4ef548f
---

                    # Configure PE1.
                    [PE1] mpls ldp remote-peer 3.3.3.3
                    [PE1-mpls-ldp-remote-3.3.3.3] remote-ip 3.3.3.3
                    [PE1-mpls-ldp-remote-3.3.3.3] quit

                    # Configure PE3.
                    [PE3] mpls ldp
                    [PE3-mpls-ldp] quit
                    [PE3] mpls ldp remote-peer 1.1.1.1
                    [PE3-mpls-ldp-remote-1.1.1.1] remote-ip 1.1.1.1
                    [PE3-mpls-ldp-remote-1.1.1.1] quit

                    After the configurations are complete, run the display mpls ldp session command
                    on PEs. The command outputs show that the Status field displays Operational,
                    indicating that a remote LDP peer relationship has been established.
                    The following example uses the command output on PE1.
                    [PE1] display mpls ldp session

                    LDP Session(s) in Public Network
                    Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                    A '*' before a session means the session is being deleted.
                    ------------------------------------------------------------------------------
                    PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    ------------------------------------------------------------------------------
                    2.2.2.2:0        Operational DU Passive 000:00:03 16/16
                    3.3.3.3:0        Operational DU Passive 000:00:00 1/1
                    ------------------------------------------------------------------------------
                    TOTAL: 2 session(s) Found.

         Step 6 Configure tunnel policies on PEs.
                    # Configure PE1.
                    [PE1] tunnel-policy p1
                    [PE1-tunnel-policy-p1] tunnel select-seq cr-lsp load-balance-number 1
                    [PE1-tunnel-policy-p1] quit

                    # Configure PE3.
                    [PE3] tunnel-policy p1
                    [PE3-tunnel-policy-p1] tunnel select-seq cr-lsp load-balance-number 1
                    [PE3-tunnel-policy-p1] quit

         Step 7 Configure PWs on PEs using PW templates.
                    # Configure the primary and secondary PWs on PE1. Configure a PW on PE2 and
                    PE3. PE2 and PE3 each have only one PW.
                    # Configure PE1.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                       521
VPN Configuration
VPN Configuration                                                                     5 VPWS Configuration

                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] pw-template 1to2
                    [PE1-pw-template-1to2] peer-address 2.2.2.2
                    [PE1-pw-template-1to2] control-word
                    [PE1-pw-template-1to2] quit
                    [PE1] pw-template 1to3
                    [PE1-pw-template-1to3] peer-address 3.3.3.3
                    [PE1-pw-template-1to3] control-word
                    [PE1-pw-template-1to3] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] undo portswitch
                    [PE1-10GE1/0/1] mpls l2vc pw-template 1to3 100 tunnel-policy p1
                    [PE1-10GE1/0/1] mpls l2vc pw-template 1to2 200 secondary
                    [PE1-10GE1/0/1] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] pw-template 2to1
                    [PE2-pw-template-2to1] peer-address 1.1.1.1
                    [PE2-pw-template-2to1] control-word
                    [PE2-pw-template-2to1] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] undo portswitch
                    [PE2-10GE1/0/1] mpls l2vc pw-template 2to1 200
                    [PE2-10GE1/0/1] quit

                    # Configure PE3.
                    [PE3] mpls l2vpn
                    [PE3-l2vpn] quit
                    [PE3] pw-template 3to1
                    [PE3-pw-template-3to1] peer-address 1.1.1.1
                    [PE3-pw-template-3to1] control-word
                    [PE3-pw-template-3to1] quit
                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/1] undo portswitch
                    [PE3-10GE1/0/1] mpls l2vc pw-template 3to1 100 tunnel-policy p1
                    [PE3-10GE1/0/1] quit

                    After the configurations are complete, run the display mpls l2vc command on PEs
                    to check L2VPN connection information. The command outputs show that the
                    primary and secondary PWs have been established and are up. The primary PW is
                    in the active state, and the secondary PW is in the inactive state.
                    The following example uses the command output on PE1.
                    [PE1] display mpls l2vc interface 10ge 1/0/1
                     *client interface      : 10GE1/0/1 is up
                      Administrator PW          : no
                      session state        : up
                      AC status           : up
                      Ignore AC state         : disable
                      VC state           : up
                      Label state         :0
                      Token state           :0
                      VC ID             : 100
                      VC type            : Ethernet
                      destination          : 3.3.3.3
                      local group ID         :0         remote group ID    :0
                      local VC label        : 4097       remote VC label    : 4096
                      local AC OAM State         : up
                      local PSN OAM State : up
                      local forwarding state : forwarding
                      local status code       : 0x0
                      remote AC OAM state : up
                      remote PSN OAM state : up


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                        522
VPN Configuration
VPN Configuration                                                                              5 VPWS Configuration

