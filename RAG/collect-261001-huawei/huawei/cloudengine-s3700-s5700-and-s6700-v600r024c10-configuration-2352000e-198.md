---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-198
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [28929, 29082]
sha256: 355528ff3fe9e5838d2f0c7fdf42c9e3f658f2cf6b197f151f4a686e19b6478e
---

                    # Configure PE2.
                    [PE2] mpls lsr-id 192.168.3.3
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface vlanif 20
                    [PE2-Vlanif20] mpls
                    [PE2-Vlanif20] mpls ldp
                    [PE2-Vlanif20] quit
                    [PE2] mpls ldp remote-peer 192.168.2.2
                    [PE2-mpls-ldp-remote-192.168.2.2] remote-ip 192.168.2.2
                    [PE2-mpls-ldp-remote-192.168.2.2] quit

                    # Configure the P.
                    [P] mpls lsr-id 192.168.4.4
                    [P] mpls
                    [P-mpls] quit
                    [P] mpls ldp
                    [P-mpls-ldp] quit
                    [P] interface vlanif 10
                    [P-Vlanif10] mpls
                    [P-Vlanif10] mpls ldp
                    [P-Vlanif10] quit
                    [P] interface vlanif 20
                    [P-Vlanif20] mpls
                    [P-Vlanif20] mpls ldp
                    [P-Vlanif20] quit

                    After completing the configurations, run the display mpls ldp session command.
                    The command output shows that an LDP session has been established between
                    the PEs and between each PE and the P. The status of these LDP sessions is
                    Operational.

                    The following example uses the command output on PE1.
                    <PE1> display mpls ldp session

                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                    --------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge              KASent/Rcv
                    --------------------------------------------------------------------------
                     192.168.3.3:0       Operational DU Passive 000:00:00 4/5
                     192.168.4.4:0       Operational DU Passive 000:00:02 10/10
                    --------------------------------------------------------------------------
                    TOTAL: 2 Session(s) Found.

         Step 4 Enable MPLS L2VPN and create a VPWS connection on PEs.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] interface vlanif 20


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    464
VPN Configuration
VPN Configuration                                                                                    5 VPWS Configuration

                    [PE1-Vlanif20] mpls l2vc 192.168.3.3 100
                    [PE1-Vlanif20] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] interface vlanif 10
                    [PE2-Vlanif10] mpls l2vc 192.168.2.2 100
                    [PE2-Vlanif10] quit

                    ----End

Verifying the Configuration
                    # Check VPWS connection information on PEs. The command outputs show that a
                    VC has been established and is in the up state.
                    The following example uses the command output on PE1.
                    <PE1> display mpls l2vc interface vlanif 20
                    *client interface      : Vlanif20 is up
                      Administrator PW              : no
                      session state         : up
                      AC status            : up
                      VC state            : up
                      Label state          :0
                      Token state            :0
                      VC ID             : 100
                      VC type             : VLAN
                      destination           : 192.168.3.3
                      local group ID          :0              remote group ID       :0
                      local VC label         : 1035             remote VC label      : 1040
                      local AC OAM State             : up
                      local PSN OAM State : up
                      local forwarding state : forwarding
                      local status code         : 0x0 (forwarding)
                      remote AC OAM state : up
                      remote PSN OAM state : up
                      remote forwarding state: forwarding
                      remote status code            : 0x0 (forwarding)
                      remote interface           : Vlanif10
                      ignore standby state : no
                      BFD for PW               : unavailable
                      VCCV State              : up
                      manual fault             : not set
                      active state         : active
                      forwarding entry            : exist
                      TTL Value             :1
                      link state         : up
                      local VC MTU               : 1500           remote VC MTU         : 1500
                      local VCCV             : alert ttl lsp-ping bfd
                      remote VCCV                 : alert ttl lsp-ping bfd
                      local control word          : disable       remote control word : disable
                      tunnel policy name            : --
                      PW template name                 : --
                      primary or secondary : primary
                      load balance type            : flow
                      Access-port            : false
                      Switchover Flag            : false
                      VC tunnel info           : 1 tunnels
                        NO.0 TNL type             : ldp            , TNL ID : 0x0000000001004c6b43
                      create time           : 0 days, 0 hours, 1 minutes, 50 seconds
                      up time             : 0 days, 0 hours, 0 minutes, 20 seconds
                      last change time           : 0 days, 0 hours, 0 minutes, 20 seconds
                      VC last up time           : 2024/08/07 14:15:06
                      VC total up time           : 0 days, 0 hours, 0 minutes, 20 seconds
                      CKey              : 385


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                      465
VPN Configuration
VPN Configuration                                                                   5 VPWS Configuration

                     NKey             : 16777476
                     PW redundancy mode           : frr
                     AdminPw interface        : --
                     AdminPw link state       : --
                     Forward state        : send active, receive active
                     Diffserv Mode         : uniform
                     Service Class      : cs-
                     Color           : --
                     DomainId            :-
                     Domain Name              :-

                    # CE1 and CE2 can ping each other successfully.

                    The following example uses the command output on CE1.
                    <CE1> ping 10.10.1.2
                     PING 10.10.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.10.1.2: bytes=56 Sequence=1 ttl=255 time=6 ms
                      Reply from 10.10.1.2: bytes=56 Sequence=2 ttl=255 time=2 ms
                      Reply from 10.10.1.2: bytes=56 Sequence=3 ttl=255 time=2 ms
                      Reply from 10.10.1.2: bytes=56 Sequence=4 ttl=255 time=3 ms
                      Reply from 10.10.1.2: bytes=56 Sequence=5 ttl=255 time=2 ms

