---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-240
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2008-07-24", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [35272, 35405]
sha256: 8f55a89906d9a13bbc956472003482f193537311b3bcdbc4d6f4bd41d8a7f073
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Run the display mpls l2vc interface command on PEs to check L2VPN connection
                    information. The command outputs show that PWs have been established and are
                    in the active state, and that BFD for PWs is not configured for the primary and
                    secondary PWs.
                    The following example uses the command output on PE1.
                    <PE1> display mpls l2vc interface 10ge 1/0/1
                    *client interface        : 10GE1/0/1 is up
                      session state           : up
                      AC status                : up
                      VC state              : up
                      VC ID                : 100
                      VC type               : PPP
                      destination             : 4.4.4.4
                      local group ID            :0             remote group ID     :0
                      local VC label           : 21504           remote VC label    : 21504
                      local AC OAM State              : up
                      local PSN State            : up
                      local forwarding state : forwarding
                      local status code          : 0x0
                      remote AC OAM state : up
                      remote PSN state              : up
                      remote forwarding state: forwarding
                      remote statuscode              : 0x0
                      BFD for PW                : unavailable
                      manual fault              : not set
                      active state           : active
                      forwarding entry             : exist
                      link state           : up
                      local VC MTU                : 4470          remote VC MTU        : 4470
                      Local VCCV         : cw alert lsp-ping bfd
                      Remote VCCV           : cw alert lsp-ping bfd
                      local control word           : enable        remote control word : enable
                      tunnel policy name             : --
                      traffic behavior name : --
                      PW template name                  : 1to2
                      primary or secondary : primary
                      VC tunnel/token info : 1 tunnels/tokens
                        NO.0 TNL type : lsp , TNL ID : 0x1002004
                      create time             : 0 days, 1 hours, 22 minutes, 22 seconds
                      up time               : 0 days, 1 hours, 21 minutes, 14 seconds
                      last change time            : 0 days, 1 hours, 21 minutes, 14 seconds
                      VC last up time : 2008-07-24 12:31:31
                      VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds
                      CKey              : 16
                      NKey               : 15
                     *client interface         : 10GE1/0/1 is up
                      session state           : up
                      AC status                : up


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                   560
VPN Configuration
VPN Configuration                                                                               5 VPWS Configuration

                     VC state             : up
                     VC ID               : 200
                     VC type              : PPP
                     destination            : 5.5.5.5
                     local group ID           :0             remote group ID     :0
                     local VC label          : 21505           remote VC label    : 21504
                     local AC OAM state             : up
                     local PSN state          : up
                     local forwarding state : forwarding
                     local status code         : 0x0
                     remote AC OAM state : up
                     remote PSN state             : up
                     remote forwarding state: forwarding
                     remote statuscode             : 0x0
                     BFD for PW               : unavailable
                     manual fault             : not set
                     active state          : inactive
                     forwarding entry            : existent
                     link state          : up
                     local VC MTU               : 4470          remote VC MTU        : 4470
                     Local VCCV        : cw alert lsp-ping bfd
                     Remote VCCV          : cw alert lsp-ping bfd
                     local control word          : enable        remote control word : enable
                     tunnel policy           : --
                     traffic behavior        : --
                     PW template name                 : 1to3
                     primary or secondary : secondary
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type : lsp , TNL ID : 0x1002006
                     create time            : 0 days, 1 hours, 22 minutes, 9 seconds
                     up time              : 0 days, 1 hours, 20 minutes, 22 seconds
                     last change time           : 0 days, 1 hours, 20 minutes, 22 seconds
                     VC last up time : 2008-07-24 12:31:31
                     VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds
                     CKey             : 17
                     NKey              : 18
                    reroute policy           : delay 30 s, resume 10 s
                    reason of last reroute : --
                    time of last reroute : -- days, -- hours, -- minutes, -- seconds
                    delay timer ID            : --           residual time :--
                    resume timer ID              : --         residual time :--

         Step 6 Configure dynamic BFD for PWs on PEs.
                    # Configure PE1.
                    [PE1] bfd
                    [PE1-bfd] quit
                    [PE1] pw-template 1to2
                    [PE1-pw-template-1to2] control-word
                    [PE1-pw-template-1to2] bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                    [PE1-pw-template-1to2] quit
                    [PE1] pw-template 1to3
                    [PE1-pw-template-1to3] control-word
                    [PE1-pw-template-1to3] bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                    [PE1-pw-template-1to3] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                    [PE1-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                    secondary
                    [PE1-10GE1/0/1] quit

                    # Configure PE2.
                    [PE2] bfd
                    [PE2-bfd] quit
                    [PE2] pw-template 2to1
                    [PE2-pw-template-2to1] control-word
                    [PE2-pw-template-2to1] bfd-detect min-rx-interval 100 min-tx-interval 100 detect-multiplier 4
                    [PE2-pw-template-2to1] quit


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                      561
VPN Configuration
VPN Configuration                                                                                       5 VPWS Configuration

                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 detect-multiplier 4 track-
                    interface
                    [PE2-10GE1/0/1] quit

