---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-214
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [31561, 31668]
sha256: 4529a63b1e08182cc7a37a6774d2be6f18e16d4ea54c40358b145b1f79520c4e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Verifying the Configuration
                    # Run the display mpls l2vc command on PEs. The command outputs show that
                    an L2VC has been established and is in the up state.
                    The following example uses the command output on PE1.
                    <PE1> display mpls l2vc interface 10ge 1/0/1
                    *client interface    : 10GE1/0/1 is up
                     Administrator PW               : no
                     session state             : up
                     AC status                : up
                     VC state                : up
                     Label state              :0
                     Token state                :0
                     VC ID                  : 100
                     VC type                 : VLAN
                     destination               : 2.2.2.2
                     local group ID              :0          remote group ID    :0
                     local VC label             : 18         remote VC label   : 18
                     local AC OAM State              : up
                     local PSN OAM State              : up
                     local forwarding state         : forwarding
                     local status code            : 0x0 (forwarding)
                     remote AC OAM State                : up
                     remote PSN OAM state                : up
                     remote forwarding state           : forwarding
                     remote status code             : 0x0 (forwarding)
                     ignore standby state           : no


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                       504
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration

                    Dynamic BFD for PW                   : enable
                    Detect Multipier            :3
                    Min Transit Interval          : 1000
                    Min Receive Interval            : 1000
                    Dynamic BFD Session                 : built
                    BFD for PW                 : available
                    BFD sessionIndex               : -- BFD state : up
                    VCCV State                : --
                    manual fault               : not set
                    active state           : active
                    forwarding entry              : exist
                    OAM Protocol                  : --
                    OAM Status                  : --
                    OAM Fault Type                  : --
                    PW APS ID                  : --
                    PW APS Status                 : --
                    TTL Value               :1
                    link state            : up
                    local VC MTU                 : 1500          remote VC MTU    : 1500
                    local VCCV               : alert ttl lsp-ping bfd
                    remote VCCV                   : alert ttl lsp-ping bfd
                    local control word            : disable remote control word : disable
                    tunnel policy name               : --
                    PW template name                    : --
                    primary or secondary              : primary
                    load balance type              : flow
                    Access-port              : false
                    Switchover Flag              : false
                    VC tunnel info             : 1 tunnels
                       NO.0 TNL type              : ldp, TNL ID : 0x0000000001004c4b43
                    create time             : 0 days, 1 hours, 2 minutes, 56 seconds
                    up time                : 0 days, 1 hours, 1 minutes, 48 seconds
                    last change time             : 0 days, 1 hours, 1 minutes, 48 seconds
                    VC last up time             : 2012/12/05 02:50:41
                    VC total up time             : 0 days, 1 hours, 1 minutes, 48 seconds
                    CKey                  :1
                    NKey                  : 1493172332
                    L2VPN QoS CIR value                 : 3888
                    L2VPN QoS PIR value                 : 3888
                    L2VPN QoS qos-profile name : --
                    PW redundancy mode                     : frr
                    AdminPw interface                 : --
                    AdminPw link state               : --
                    Forward state              : send inactive, receive inactive
                    Diffserv Mode               : uniform
                    Service Class           : --
                    Color                : --
                    DomainId                  : --
                    Domain Name                      : --

                    CE1 and CE2 can ping each other successfully. The following example uses the
                    command output on CE1.
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




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                               505
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration


