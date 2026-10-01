---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-225
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-05-24", "2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [33051, 33184]
sha256: 212342c94b23528f39c707d45e9873c631af88b5b7d1d7ba299ade3e33a5de24
---

                    # Run the display mpls l2vc interface command on PE1. The command output
                    shows that the primary PW transits to the inactive state and the secondary PW
                    transits to the active state.
                    [PE1] display mpls l2vc interface 10ge 1/0/1
                     *client interface        : 10GE1/0/1 is up
                      Administrator PW               : no
                      session state          : down
                      AC status             : up
                      Ignore AC state            : disable
                      VC state             : down
                      Label state           :0
                      Token state             :0
                      VC ID              : 100
                      VC type              : Ethernet
                      destination            : 3.3.3.3
                      local group ID           :0                 remote group ID     :0
                      local VC label          : 21504               remote VC label    :0
                      local AC OAM State               : up
                      local PSN OAM State : up
                      local forwarding state : not forwarding
                      local status code          : 0x1
                      Dynamic BFD for PW                  : enable
                      Detect Multipier            :3
                      Min Transit Interval : 100
                      Max Receive Interval : 100
                      Dynamic BFD Session : not built
                      BFD for PW                : unavailable
                      VCCV State               : up
                      manual fault              : not set
                      active state          : inactive
                      forwarding entry             : not exist
                      link state         : down
                      local VC MTU                : 1500             remote VC MTU       :0
                      local VCCV              : cw alert ttl lsp-ping bfd
                      remote VCCV                  : none
                      local control word           : enable           remote control word : none
                      tunnel policy name             : p1
                      PW template name                   : 1to3
                      primary or secondary : primary
                      load balance type             : flow
                      Access-port             : false
                      Switchover Flag             : false
                      VC tunnel/token info : 0 tunnels/tokens
                        Backup TNL type              : lsp , TNL ID : 0x0
                      create time            : 0 days, 1 hours, 5 minutes, 19 seconds
                      up time              : 0 days, 0 hours, 43 minutes, 33 seconds
                      last change time            : 0 days, 0 hours, 43 minutes, 33 seconds
                      VC last up time            : 2024-05-24 12:31:31
                      VC total up time            : 0 days, 2 hours, 12 minutes, 51 seconds
                      CKey               : 16
                      NKey                : 15
                      PW redundancy mode                    : frr
                      AdminPw interface               : --
                      AdminPw link state              : --
                      Diffserv Mode              : uniform
                      Service Class          : be
                      Color             : --
                      DomainId                 : --
                      Domain Name                     : --

                    *client interface     : 10GE1/0/1 is up
                     Administrator PW        : no
                     session state       : up
                     AC status          : up
                     VC state          : up
                     Label state        :0
                     Token state         :0
                     VC ID            : 200


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                   528
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration

                     VC type              : Ethernet
                     destination           : 2.2.2.2
                     local group ID          :0                 remote group ID     :0
                     local VC label         : 21505               remote VC label    : 21505
                     local AC OAM State              : up
                     local PSN OAM State : up
                     local forwarding state : forwarding
                     local status code         : 0x0
                     remote AC OAM state : up
                     remote PSN OAM state : up
                     remote forwarding state: forwarding
                     remote status code            : 0x0
                     ignore standby state : no
                     Dynamic BFD for PW                 : enable
                     Detect Multipier           :3
                     Min Transit Interval : 100
                     Max Receive Interval : 100
                     Dynamic BFD Session : built
                     BFD for PW               : available
                       BFD sessionIndex           : 257            BFD state : up
                     VCCV State              : up
                     manual fault             : not set
                     active state         : active
                     forwarding entry            : exist
                     link state         : up
                     local VC MTU               : 1500             remote VC MTU        : 1500
                     local VCCV             : cw alert ttl lsp-ping bfd
                     remote VCCV                 : cw alert ttl lsp-ping bfd
                     local control word          : enable           remote control word : enable
                     tunnel policy name            : --
                     PW template name                  : 1to2
                     primary or secondary : secondary
                     load balance type            : flow
                     Access-port            : false
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type             : lsp , TNL ID : 0x2002001
                       Backup TNL type             : lsp , TNL ID : 0x0
                     create time           : 0 days, 1 hours, 4 minutes, 31 seconds
                     up time              : 0 days, 0 hours, 43 minutes, 44 seconds
                     last change time           : 0 days, 0 hours, 43 minutes, 44 seconds
                     VC last up time           : 2024-05-24 12:31:31
                     VC total up time           : 0 days, 2 hours, 12 minutes, 51 seconds
                     CKey               : 17
                     NKey                : 18
                     PW redundancy mode                   : frr
                     AdminPw interface              : --
                     AdminPw link state             : --
                     Diffserv Mode             : uniform
                     Service Class         : be
                     Color             : --
                     DomainId                : --
                     Domain Name                    : --

                    reroute policy       : delay 30 s, resume 10 s
                    reason of last reroute : LDP notification message was forwarded
                    time of last reroute : 0 days, 0 hours, 43 minutes, 2 seconds
                    delay timer ID        : --       residual time :--
                    resume timer ID         : --       residual time :--

