---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-222
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32686, 32831]
sha256: 82acb47743f7963dc6c419785da2041f7c155cbb96a54d9922fde4f42dcbb3b2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    remote forwarding state: forwarding
                    remote status code           : 0x0
                    ignore standby state : no
                    BFD for PW              : unavailable
                    VCCV State             : up
                    manual fault            : not set
                    active state        : active
                    forwarding entry           : exist
                    link state         : up
                    local VC MTU              : 1500         remote VC MTU       : 1500
                    local VCCV            : cw alert ttl lsp-ping bfd
                    remote VCCV                : cw alert ttl lsp-ping bfd
                    local control word         : enable       remote control word : enable
                    tunnel policy name           : p1
                    PW template name                : 1to3
                    primary or secondary : primary
                    load balance type           : flow
                    Access-port           : false
                    Switchover Flag           : false
                    VC tunnel/token info : 1 tunnels/tokens
                      NO.0 TNL type            : cr lsp, TNL ID : 0x15
                      Backup TNL type            : lsp , TNL ID : 0x0
                    create time          : 0 days, 0 hours, 1 minutes, 4 seconds
                    up time             : 0 days, 0 hours, 0 minutes, 57 seconds
                    last change time          : 0 days, 0 hours, 0 minutes, 57 seconds
                    VC last up time          : 2024/05/23 18:49:30
                    VC total up time          : 0 days, 0 hours, 0 minutes, 57 seconds
                    CKey               :2
                    NKey               :1
                    PW redundancy mode                : frr
                    AdminPw interface             : --
                    AdminPw link state            : --
                    Diffserv Mode            : uniform
                    Service Class        : be
                    Color             : --
                    DomainId               : --
                    Domain Name                   : --

                    *client interface       : 10GE1/0/1 is up
                     Administrator PW              : no
                     session state         : up
                     AC status            : up
                     VC state            : up
                     Label state          :0
                     Token state            :0
                     VC ID             : 200
                     VC type             : Ethernet
                     destination           : 2.2.2.2
                     local group ID          :0            remote group ID      :0
                     local VC label         : 4098           remote VC label     : 4097
                     local AC OAM State             : up
                     local PSN OAM State : up
                     local forwarding state : forwarding
                     local status code         : 0x0
                     remote AC OAM state : up
                     remote PSN OAM state : up
                     remote forwarding state: forwarding
                     remote status code            : 0x0
                     ignore standby state : no
                     BFD for PW               : unavailable
                     VCCV State              : up
                     manual fault             : not set
                     active state         : inactive
                     forwarding entry            : exist
                     link state         : up
                     local VC MTU               : 1500         remote VC MTU        : 1500
                     local VCCV             : cw alert ttl lsp-ping bfd
                     remote VCCV                 : cw alert ttl lsp-ping bfd
                     local control word          : enable       remote control word : enable



Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                 523
VPN Configuration
VPN Configuration                                                                               5 VPWS Configuration

                     tunnel policy name          : --
                     PW template name               : 1to2
                     primary or secondary : secondary
                     load balance type          : flow
                     Access-port          : false
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type           : lsp , TNL ID : 0x16
                       Backup TNL type           : lsp , TNL ID : 0x0
                     create time         : 0 days, 0 hours, 1 minutes, 4 seconds
                     up time            : 0 days, 0 hours, 0 minutes, 59 seconds
                     last change time         : 0 days, 0 hours, 0 minutes, 59 seconds
                     VC last up time        : 2024/05/23 18:49:28
                     VC total up time         : 0 days, 0 hours, 0 minutes, 59 seconds
                     CKey              :4
                     NKey              :3
                     PW redundancy mode               : frr
                     AdminPw interface            : --
                     AdminPw link state           : --
                     Diffserv Mode           : uniform
                     Service Class       : be
                     Color            : --
                     DomainId              : --
                     Domain Name                  : --

                    reroute policy       : delay 30 s, resume 10 s
                    reason of last reroute : LDP notification message was forwarded
                    time of last reroute : 0 days, 0 hours, 0 minutes, 27 seconds
                    delay timer ID        : --       residual time :--
                    resume timer ID         : --       residual time :--

         Step 8 Configure CEs to communicate with each other.
                    Configure two default routes on CE2, and ensure that data is preferentially
                    forwarded through VLANIF 10.
                    # Configure CE2.
                    [CE2] ip route-static 0.0.0.0 0.0.0.0 vlanif10 10.1.1.1
                    [CE2] ip route-static 0.0.0.0 0.0.0.0 vlanif20 10.1.2.1 preference 100

         Step 9 Configure BFD for PW on PEs, and enable physical-layer fault notification on PE2
                and PE3.
                          NOTE

                         In this example, dynamic BFD for PW is used.

                    # Configure PE1.
                    [PE1] bfd
                    [PE1-bfd] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100
                    [PE1-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100 secondary
                    [PE1-10GE1/0/1] quit

                    # Configure PE2.
                    [PE2] bfd
                    [PE2-bfd] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100
                    [PE2-10GE1/0/1] mpls l2vpn trigger if-down
                    [PE2-10GE1/0/1] quit

                    # Configure PE3.
                    [PE3] bfd
                    [PE3-bfd] quit


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                  524
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration

