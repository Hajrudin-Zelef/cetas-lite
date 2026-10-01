---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-243
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2008-07-24", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [35618, 35769]
sha256: 084e5d648c9b5b35e6c36cdcab104e15dbc23f79be5ff12021bea1a0e850c229
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     VC state               : down
                     VC ID                : 100
                     VC type               : PPP
                     destination             : 4.4.4.4
                     local group ID            :0              remote group ID     :0
                     local VC label           : 21504            remote VC label    :0
                     local AC OAM State               : up
                     local PSN State            : up
                     local forwarding state : not forwarding
                     local status code          : 0x0
                     BFD for PW                : unavailable
                     manual fault              : not set
                     active state           : inactive
                     forwarding entry             : not exist
                     link state           : down
                     local VC MTU                : 4470           remote VC MTU        : 4470
                     Local VCCV         : cw alert lsp-ping bfd
                     Remote VCCV           : none
                     local control word           : enable         remote control word : none
                     tunnel policy name              : --
                     traffic behavior name : --
                     PW template name                   : 1to2
                     primary or secondary : primary
                     VC tunnel/token info : 0 tunnels/tokens
                     create time             : 0 days, 0 hours, 30 minutes, 58 seconds
                     up time               : 0 days, 0 hours, 0 minutes, 0 seconds
                     last change time            : 0 days, 0 hours, 6 minutes, 46 seconds
                     VC last up time : 2008-07-24 12:31:31
                     VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds
                     CKey              : 16
                     NKey               : 15
                    *client interface         : 10GE1/0/1 is up
                     session state           : up
                     AC status                : up
                     VC state              : up
                     VC ID                : 200
                     VC type               : PPP
                     destination             : 5.5.5.5
                     local group ID            :0              remote group ID     :0
                     local VC label           : 21505            remote VC label    : 21504
                     local AC OAM state               : up
                     local PSN state           : up
                     local forwarding state : forwarding
                     local status code          : 0x0
                     remote AC OAM state : up
                     remote PSN state               : up
                     remote forwarding state: forwarding
                     remote statuscode               : 0x0
                     BFD for PW                : available
                       BFD sessionIndex            : 257          BFD state : up
                     manual fault              : not set
                     active state           : active
                     forwarding entry             : existent
                     link state           : up
                     local VC MTU                : 4470           remote VC MTU        : 4470
                     Local VCCV         : cw alert lsp-ping bfd
                     Remote VCCV           : cw alert lsp-ping bfd
                     local control word           : enable         remote control word : enable
                     tunnel policy name              : --
                     traffic behavior name : --
                     PW template name                   : 1to3
                     primary or secondary : secondary
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type : lsp , TNL ID : 0x1002008
                     create time             : 0 days, 0 hours, 30 minutes, 58 seconds
                     up time               : 0 days, 0 hours, 25 minutes, 12 seconds
                     last change time            : 0 days, 0 hours, 25 minutes, 12 seconds
                     VC last up time : 2008-07-24 12:31:31
                     VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds



Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                   565
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration

                     CKey             : 17
                     NKey              : 18
                    reroute policy        : delay 30 s, resume 10 s
                    reason of last reroute : --
                    time of last reroute : -- days, -- hours, -- minutes, -- seconds
                    delay timer ID         : --       residual time :--
                    resume timer ID          : --       residual time :--


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         interface 10GE1/0/1
                          undo portswitch
                          ip address 10.1.1.1 255.255.255.252
                          ip address 10.1.2.1 255.255.255.252 sub
                         #
                         return
                    ●    PE1
                         #
                         sysname PE1
                         #
                          bfd
                         #
                          mpls lsr-id 1.1.1.1
                          mpls
                         #
                          mpls l2vpn
                         #
                         pw-template 1to2
                          peer-address 4.4.4.4
                          control-word
                          bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                         #
                         pw-template 1to3
                          peer-address 5.5.5.5
                          control-word
                          bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                         #
                         mpls ldp
                         #
                          mpls ldp remote-peer 4.4.4.4
                          remote-ip 4.4.4.4
                         #
                          mpls ldp remote-peer 5.5.5.5
                          remote-ip 5.5.5.5
                         #
                         interface 10GE1/0/1
                          undo portswitch
                          mpls l2vc pw-template 1to2 100
                          mpls l2vc pw-template 1to3 200 secondary
                          mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                          mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4 secondary
                         #
                         interface 10GE1/0/2
                          undo portswitch
                          ip address 10.2.1.1 255.255.255.252
                          mpls
                          mpls ldp
                         #
                         interface 10GE1/0/3
                          undo portswitch
                          ip address 10.4.1.1 255.255.255.252
                          mpls
                          mpls ldp
                         #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        566
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

