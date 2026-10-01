---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-202
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [29604, 29752]
sha256: 098080bd557c178db167619701dfadefe3d02fffa6df928ac96b6726f96ac2d4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     LspAge           : 1499 sec
                     Entropy Label Flag : None

                    -------------------------------------------------------------------------------
                                 LSP Information: LDP LSP
                    -------------------------------------------------------------------------------
                      No                : 3
                      VrfIndex            :
                      Fec              : 1.1.1.9/32
                      Nexthop                : 127.0.0.1
                      In-Label            : 3
                      Out-Label              : NULL
                      In-Interface          : ----------
                      Out-Interface           : ----------
                      LspIndex             : 5000001
                      Type               : Primary
                      OutSegmentIndex             : 4294967295
                      LsrType             : Egress
                      Outgoing TunnelType : ------
                      Outgoing TunnelID : 0x0
                      Label Operation           : POP
                      Mpls-Mtu                : ------
                      LspAge               : 1243 sec
                      Ingress-ELC            : ------

                    # Check L2VPN connection information on PEs. The command outputs show that
                    an L2VC has been established and is in the up state.
                    The following example uses the command output on PE1.
                    [PE1] display mpls l2vc interface vlanif 20
                     *client interface       : Vlanif20 is up
                      Administrator PW              : no
                      session state         : up
                      AC status            : up
                      Ignore AC state           : disable
                      Ignore AC state           : disable
                      VC state            : up
                      Label state          :0
                      Token state            :0
                      VC ID             : 10
                      VC type             : VLAN
                      destination           : 3.3.3.9
                      local group ID          :0               remote group ID    :0
                      local VC label         : 1026             remote VC label    : 1032
                      local AC OAM State             : up
                      local PSN OAM State : up
                      local forwarding state : forwarding
                      local status code         : 0x0 (forwarding)
                      remote AC OAM state : up
                      remote PSN OAM state : up
                      remote forwarding state: forwarding
                      remote status code            : 0x0 (forwarding)
                      remote interface           : Vlanif20
                      ignore standby state : no
                      Dynamic BFD for PW                : disable
                      VCCV State              : up
                      manual fault             : not set
                      active state         : active
                      forwarding entry            : exist
                      TTL Value             :1
                      link state         : up
                      local VC MTU               : 1500           remote VC MTU       : 1500
                      local VCCV             : alert ttl lsp-ping bfd
                      remote VCCV                 : alert ttl lsp-ping bfd
                      local control word          : disable       remote control word : disable
                      tunnel policy name            :1
                      PW template name                 : --
                      primary or secondary : primary
                      load balance type            : flow


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       474
VPN Configuration
VPN Configuration                                                                       5 VPWS Configuration

                     Access-port         : false
                     Switchover Flag        : false
                     VC tunnel info       : 1 tunnels
                       NO.0 TNL type         : te, TNL ID : 0x48000002
                     create time        : 0 days, 4 hours, 16 minutes, 25 seconds
                     up time           : 0 days, 4 hours, 15 minutes, 58 seconds
                     last change time       : 0 days, 4 hours, 15 minutes, 58 seconds
                     VC last up time       : 2024/05/09 22:57:04
                     VC total up time       : 0 days, 4 hours, 15 minutes, 58 seconds
                     CKey             :4
                     NKey             :3
                     PW redundancy mode                   : frr
                     AdminPw interface                : --
                     AdminPw link state               : --
                     Forward state                : send active, receive active
                     Diffserv Mode                 : uniform
                     Service Class              : cs-
                     Color                   : --
                     DomainId                    : --
                     Domain Name                      : --

                    # CE1 and CE2 can ping each other successfully.
                    <CE1> ping 10.10.1.2
                     PING 10.10.1.2: 56 data bytes, press CTRL_C to break
                        Reply from 10.10.1.2: bytes=56 Sequence=1 ttl=255 time=125 ms
                        Reply from 10.10.1.2: bytes=56 Sequence=2 ttl=255 time=125 ms
                        Reply from 10.10.1.2: bytes=56 Sequence=3 ttl=255 time=94 ms
                        Reply from 10.10.1.2: bytes=56 Sequence=4 ttl=255 time=125 ms
                        Reply from 10.10.1.2: bytes=56 Sequence=5 ttl=255 time=125 ms
                      --- 10.10.1.2 ping statistics ---
                        5 packet(s) transmitted
                        5 packet(s) received
                        0.00% packet loss
                        round-trip min/avg/max = 94/118/125 ms


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan batch 20
                         #
                         interface Vlanif20
                          ip address 10.10.1.1 255.255.255.0
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 20
                         #
                         return
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 10 20
                         #
                         mpls lsr-id 1.1.1.9
                         #
                         mpls
                          mpls te
                          mpls rsvp-te
                          mpls te cspf
                         #
                         mpls l2vpn
                         #
                         mpls ldp


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                           475
VPN Configuration
VPN Configuration                                                            5 VPWS Configuration

