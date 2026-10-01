---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-211
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [31022, 31193]
sha256: 3b8802b15180dcf9b68fc9896c4ef5a5980b7a54df31e1d64fa32eaa74d61302
---

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] mpls l2vpn vpn1 encapsulation ethernet
                    [PE2-mpls-l2vpn-vpn1] route-distinguisher 200:1
                    [PE2-mpls-l2vpn-vpn1] vpn-target 1:1 both
                    [PE2-mpls-l2vpn-vpn1] ce ce4 id 4 range 10
                    [PE2-mpls-l2vpn-vpn1-ce-ce4] connection ce-offset 3 interface 10ge 1/0/2
                    [PE2-mpls-l2vpn-vpn1-ce-ce4] quit
                    [PE2-mpls-l2vpn-vpn1] quit

                    ----End

Verifying the Configuration
                    # Check VPWS connection information on PEs. The command outputs show that a
                    BGP VC has been established.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     496
VPN Configuration
VPN Configuration                                                                    5 VPWS Configuration


                    The following example uses the command output on PE1.
                    [PE1] display mpls l2vpn connection interface 10ge 1/0/1
                    conn-type: remote
                       local vc state:          up
                       remote vc state:           up
                       local ce-id:           1
                       local ce name:            ce1
                       remote ce-id:             2
                       intf(state,encap):        10GE1/0/1(up,ethernet)
                       peer id:               2.2.2.9
                       route-distinguisher:        100:1
                       local vc label:          294930
                       remote vc label:           294929
                       tunnel policy:           default
                       CKey:                  65
                       NKey:                  3841982617
                       primary or secondary:          primary
                       forward entry exist or not: true
                       forward entry active or not:true
                       manual fault set or not: not set
                       AC OAM state:               up
                       BFD for PW session index: --
                       BFD for PW state:            invalid
                       BFD for LSP state:          true
                       Local C bit is not set
                       Remote C bit is not set
                       tunnel type:             ldp
                       tunnel id:              0x0000000001004c6b42

                    # CE1 and CE2 can ping each other successfully.
                    The following example uses the command output on CE1.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=430 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=220 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=190 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=190 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=190 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 190/244/430 ms


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         interface 10GE1/0/1
                          undo portswitch
                          ip address 10.1.1.1 255.255.255.0
                         #
                         return
                    ●    CE2
                         #
                         sysname CE2
                         #
                         interface 10GE1/0/1
                          undo portswitch
                          ip address 10.1.1.2 255.255.255.0
                         #
                         return


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         497
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration


                    ●   PE1
                        #
                        sysname PE1
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.10.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        mpls l2vpn vpn1 encapsulation ethernet
                         route-distinguisher 100:1
                         vpn-target 1:1 import-extcommunity
                         vpn-target 1:1 export-extcommunity
                         ce ce1 id 1 range 10 default-offset 0
                          connection ce-offset 2 interface 10GE1/0/1
                        #
                        bgp 100
                         peer 2.2.2.9 as-number 100
                         peer 2.2.2.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 2.2.2.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 2.2.2.9 enable
                          peer 2.2.2.9 signaling vpws
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return

                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.10.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         undo portswitch
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   498
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

