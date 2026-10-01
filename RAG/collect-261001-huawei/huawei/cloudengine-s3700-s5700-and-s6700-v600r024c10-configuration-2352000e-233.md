---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-233
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2008-07-24", "2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [34288, 34421]
sha256: 7fb5450e129af696a9b2e8844c7e71856529b553310f66c67af1748ff75975ec
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Verifying the Configuration
                    # After the configurations are complete, a BFD session has been established
                    between PE1 and PE2 and between PE1 and PE3. Run the display bfd session all
                    command. The command output shows that the State field displays Up.
                    The following example uses the command output on PE1.
                    <PE1> display bfd session all
                    S: Static session
                    D: Dynamic session
                    IP: IP session
                    IF: Single-hop session
                    PEER: Multi-hop session
                    AUTO: Automatically negotiated session
                    (w): State in WTR
                    (*): State is invalid

                    --------------------------------------------------------------------------------
                    Local Remote PeerIpAddr                        State     Type       InterfaceName
                    --------------------------------------------------------------------------------
                    12     21     --.--.--.--               Up        S_PW(M) 10GE1/0/1
                    13     31     --.--.--.--               Up        S_PW(S) 10GE1/0/1
                    --------------------------------------------------------------------------------
                        Total UP/DOWN Session Number : 2/0

                    # When the primary PW is working properly, CE1 can use the primary address to
                    ping 10.1.1.2 of CE2 successfully. Because the secondary PW is not transmitting
                    data, CE1 cannot use the secondary address to ping 10.1.2.2 of CE2.
                    <CE1> ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                       Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=140 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=90 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=120 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=120 ms
                       Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=130 ms
                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 90/120/140 ms
                    <CE1> ping 10.1.2.2
                     PING 10.1.2.2: 56 data bytes, press CTRL_C to break
                       Request time out
                       Request time out
                       Request time out
                       Request time out
                       Request time out
                     --- 10.1.2.2 ping statistics ---
                       5 packet(s) transmitted
                       0 packet(s) received
                       100.00% packet loss

                    # Run the display mpls l2vc interface command on PEs to check PW status. The
                    command outputs show that the BFD for PW field displays available, and the
                    BFD state field displays up for both the primary and secondary PWs.
                    <PE1> display mpls l2vc interface 10ge 1/0/1
                    *client interface  : 10GE1/0/1 is up


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                          546
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration

                     session state           : up
                     AC status                : up
                     VC state              : up
                     VC ID                : 100
                     VC type               : Ethernet
                     destination             : 4.4.4.4
                     local group ID            :0              remote group ID     :0
                     local VC label           : 21504            remote VC label    : 21504
                     local AC OAM State               : up
                     local PSN State            : up
                     local forwarding state : forwarding
                     local status code          : 0x0
                     remote AC OAM state : up
                     remote PSN state               : up
                     remote forwarding state: forwarding
                     remote statuscode               : 0x0
                     BFD for PW                : available
                       BFD sessionIndex            : 256          BFD state : up
                     manual fault              : not set
                     active state           : active
                     forwarding entry             : exist
                     link state           : up
                     local VC MTU                : 4470           remote VC MTU        : 4470
                     Local VCCV         : cw alert lsp-ping bfd
                     Remote VCCV           : cw alert lsp-ping bfd
                     local control word           : enable         remote control word : enable
                     tunnel policy name              : --
                     traffic behavior name : --
                     PW template name                   : 1to2
                     primary or secondary : primary
                     VC tunnel/token info : 1 tunnels/tokens
                       NO.0 TNL type : lsp , TNL ID : 0x1002004
                     create time             : 0 days, 1 hours, 17 minutes, 55 seconds
                     up time               : 0 days, 1 hours, 16 minutes, 47 seconds
                     last change time            : 0 days, 1 hours, 16 minutes, 47 seconds
                     VC last up time : 2008-07-24 12:31:31
                     VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds
                     CKey              : 16
                     NKey               : 15
                    *client interface         : 10GE1/0/1 is up
                     session state           : up
                     AC status                : up
                     VC state              : up
                     VC ID                : 200
                     VC type               : Ethernet
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
                     active state           : inactive
                     forwarding entry             : existent
                     link state           : up
                     local VC MTU                : 4470           remote VC MTU        : 4470
                     Local VCCV         : cw alert lsp-ping bfd
                     Remote VCCV           : cw alert lsp-ping bfd
                     local control word           : enable         remote control word : enable
                     tunnel policy name              : --
                     traffic behavior name : --
                     PW template name                   : 1to3



