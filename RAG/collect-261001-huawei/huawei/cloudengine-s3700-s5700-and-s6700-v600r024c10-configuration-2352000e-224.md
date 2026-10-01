---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-224
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-05-23", "2025-03-03"]
keywords: ["copyright", "cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32927, 33050]
sha256: 5a0aaf35647526ca433618eb54e645aa3e53b2e193fb9f9cf5a6557846d9a964
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                     local group ID          :0              remote group ID      :0
                     local VC label         : 21505             remote VC label    : 21505
                     local AC OAM State             : up
                     local PSN OAM State : up
                     local forwarding state : forwarding
                     local status code         : 0x0
                     remote AC OAM state : up
                     remote PSN OAM state : up
                     remote forwarding state: forwarding
                     remote status code            : 0x0
                     ignore standby state : no
                     Dynamic BFD for PW                : enable
                     Detect Multipier           :3
                     Min Transit Interval : 100
                     Max Receive Interval : 100
                     Dynamic BFD Session : built
                     BFD for PW               : available
                       BFD sessionIndex           : 257          BFD state : up
                     VCCV State              : up
                     manual fault             : not set
                     active state         : inactive
                     forwarding entry            : exist
                     link state         : up
                     local VC MTU               : 1500           remote VC MTU       : 1500
                     local VCCV             : cw alert ttl lsp-ping bfd
                     remote VCCV                 : cw alert ttl lsp-ping bfd
                     local control word          : enable         remote control word : enable
                     tunnel policy name            : --
                     PW template name                 : 1to2
                     primary or secondary : secondary
                     load balance type            : flow
                     Access-port            : false
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type             : lsp , TNL ID : 0x2002001



Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                   526
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration

                       Backup TNL type          : lsp , TNL ID : 0x0
                     create time          : 0 days, 1 hours, 4 minutes, 31 seconds
                     up time             : 0 days, 0 hours, 43 minutes, 44 seconds
                     last change time         : 0 days, 0 hours, 43 minutes, 44 seconds
                     VC last up time        : 2024-05-23 12:31:31
                     VC total up time         : 0 days, 2 hours, 12 minutes, 51 seconds
                     CKey              : 17
                     NKey               : 18
                     PW redundancy mode              : frr
                     AdminPw interface           : --
                     AdminPw link state          : --
                     Diffserv Mode           : uniform
                     Service Class        : be
                     Color            : --
                     DomainId              : --
                     Domain Name                 : --

                    reroute policy       : delay 30 s, resume 10 s
                    reason of last reroute : LDP notification message was forwarded
                    time of last reroute : 0 days, 0 hours, 43 minutes, 2 seconds
                    delay timer ID        : --       residual time :--
                    resume timer ID         : --       residual time :--

                    # Check the IP routing table on CE2. The command output shows that the
                    outbound interface of the default route on CE2 is VLANIF 10. This indicates that
                    traffic is transmitted along the primary PW.
                    [CE2] display ip routing-table 0.0.0.0
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance
                    ------------------------------------------------------------------------------
                    Routing Table : Public
                    Summary Count : 1
                    Destination/Mask Proto Pre Cost                Flags NextHop           Interface

                          0.0.0.0/0 Static 60 0              D 10.1.1.1          Vlanif10

                    # You can view that CE2 can ping the address 10.1.3.1 of CE1 successfully.
                    [CE1] ping 10.1.3.1
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.3.1: bytes=56 Sequence=1 ttl=255 time=430 ms
                      Reply from 10.1.3.1: bytes=56 Sequence=2 ttl=255 time=220 ms
                      Reply from 10.1.3.1: bytes=56 Sequence=3 ttl=255 time=190 ms
                      Reply from 10.1.3.1: bytes=56 Sequence=4 ttl=255 time=190 ms
                      Reply from 10.1.3.1: bytes=56 Sequence=5 ttl=255 time=190 ms

                     --- 10.1.3.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 190/244/430 ms

                    # Manually simulate a failure on 10GE1/0/2 of PE3.
                    [PE3] interface 10ge 1/0/2
                    [PE3-10GE1/0/2] shutdown
                    [PE3-10GE1/0/2] quit

                    # Run the display bfd session all command on PE1. The command output shows
                    that the BFD session of the primary PW is down.
                    [PE1] display bfd session all
                    --------------------------------------------------------------------------------
                    Local Remote        PeerIpAddr       State     Type         InterfaceName
                    --------------------------------------------------------------------------------
                    8193 8192          --.--.--.--  Up        D_PW(S)        10GE1/0/1
                    --------------------------------------------------------------------------------
                    Total UP/DOWN Session Number : 1/0


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                        527
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration


