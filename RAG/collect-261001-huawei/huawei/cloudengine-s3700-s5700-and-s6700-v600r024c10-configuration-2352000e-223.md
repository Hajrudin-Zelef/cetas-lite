---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-223
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-05-23", "2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32832, 32926]
sha256: 418962b94e75c6fcf31dd005777c4f0526b2123c517a41b32a92c5e13c229207
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100
                    [PE3-10GE1/0/1] mpls l2vpn trigger if-down
                    [PE3-10GE1/0/1] quit

                    After the configurations are complete, a BFD session has been established
                    between PE1 and PE2 and between PE1 and PE3. Run the display bfd session all
                    command. The command output shows that the State field displays Up.
                    The following example uses the command output on PE1.
                    [PE1] display bfd session all
                    --------------------------------------------------------------------------------
                    Local Remote PeerIpAddr             State     Type        InterfaceName
                    --------------------------------------------------------------------------------
                    8192 8192 --.--.--.--          Up       D_PW(M)          10GE1/0/1
                    8193 8192 --.--.--.--          Up       D_PW(S)         10GE1/0/1
                    --------------------------------------------------------------------------------
                        Total UP/DOWN Session Number : 2/0

                    ----End

Verifying the Configuration
                    # Check LDP VPWS connection information on PE1. The command output shows
                    that the primary PW is in the active state, the secondary PW is in the inactive
                    state, and the BFD for PW status of the primary and secondary PWs is available.
                    [PE1] display mpls l2vc interface 10ge 1/0/1
                     *client interface       : 10GE1/0/1 is up
                      Administrator PW              : no
                      session state         : up
                      AC status            : up
                      Ignore AC state           : disable
                      VC state            : up
                      Label state          :0
                      Token state            :0
                      VC ID             : 100
                      VC type             : Ethernet
                      destination           : 3.3.3.3
                      local group ID          :0             remote group ID      :0
                      local VC label         : 21504            remote VC label    : 21504
                      local AC OAM State             : up
                      local PSN OAM State : up
                      local forwarding state : forwarding
                      local status code         : 0x0
                      remote AC OAM state : up
                      remote PSN OAM state : up
                      remote forwarding state: forwarding
                      remote status code            : 0x0
                      ignore standby state : no
                      Dynamic BFD for PW               : enable
                      Detect Multipier           :3
                      Min Transit Interval : 100
                      Max Receive Interval : 100
                      Dynamic BFD Session : built
                      BFD for PW               : available
                        BFD sessionIndex           : 256         BFD state : up
                      VCCV State              : up
                      manual fault             : not set
                      active state         : active
                      forwarding entry            : exist
                      link state         : up
                      local VC MTU               : 1500          remote VC MTU        : 1500
                      local VCCV             : cw alert ttl lsp-ping bfd
                      remote VCCV                 : cw alert ttl lsp-ping bfd
                      local control word          : enable        remote control word : enable


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                        525
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration

                    tunnel policy name           : p1
                    PW template name                : 1to3
                    primary or secondary : primary
                    load balance type           : flow
                    Access-port           : false
                    Switchover Flag           : false
                    VC tunnel/token info : 1 tunnels/tokens
                      NO.0 TNL type            : cr lsp, TNL ID : 0x42002000
                      Backup TNL type            : lsp , TNL ID : 0x0
                    create time          : 0 days, 1 hours, 5 minutes, 19 seconds
                    up time             : 0 days, 0 hours, 43 minutes, 33 seconds
                    last change time          : 0 days, 0 hours, 43 minutes, 33 seconds
                    VC last up time         : 2024-05-23 12:31:31
                    VC total up time          : 0 days, 2 hours, 12 minutes, 51 seconds
                    CKey              : 16
                    NKey               : 15
                    PW redundancy mode                : frr
                    AdminPw interface             : --
                    AdminPw link state            : --
                    Diffserv Mode            : uniform
                    Service Class        : be
                    Color            : --
                    DomainId               : --
                    Domain Name                   : --

