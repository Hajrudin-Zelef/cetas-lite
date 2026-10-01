---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-310
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45990, 46139]
sha256: ee8bb8fba0df9a9732620f210c8dfaaec02a6a6f7fee935e6605b567041c5ed2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    [SPE2] mpls ldp remote-peer 1.1.1.9
                    [SPE2-mpls-ldp-remote-1.1.1.9] remote-ip 1.1.1.9
                    [SPE2-mpls-ldp-remote-1.1.1.9] quit

         Step 5 Enable MPLS L2VPN and configure LDP VPWS on UPEs.
                    # Configure UPE1.
                    [UPE1] mpls l2vpn
                    [UPE1-l2vpn] quit
                    [UPE1] interface vlanif 50
                    [UPE1-Vlanif50] mpls l2vc 1.1.1.9 100
                    [UPE1-Vlanif50] quit

                    # Configure UPE2.
                    [UPE2] mpls l2vpn
                    [UPE2-l2vpn] quit
                    [UPE2] interface vlanif 60
                    [UPE2-Vlanif60] mpls l2vc 3.3.3.9 100
                    [UPE2-Vlanif60] quit

         Step 6 Enable MPLS L2VPN and configure a VSI on SPEs.
                    # Configure SPE1.
                    [SPE1] mpls l2vpn
                    [SPE1-l2vpn] quit
                    [SPE1] vsi v100 static
                    [SPE1-vsi-v100] pwsignal ldp
                    [SPE1-vsi-v100-ldp] vsi-id 100
                    [SPE1-vsi-v100-ldp] peer 3.3.3.9
                    [SPE1-vsi-v100-ldp] peer 4.4.4.9 upe
                    [SPE1-vsi-v100-ldp] quit
                    [SPE1-vsi-v100] quit

                    # Configure SPE2.
                    [SPE2] mpls l2vpn
                    [SPE2-l2vpn] quit
                    [SPE2] vsi v100 static
                    [SPE2-vsi-v100] pwsignal ldp
                    [SPE2-vsi-v100-ldp] vsi-id 100
                    [SPE2-vsi-v100-ldp] peer 1.1.1.9
                    [SPE2-vsi-v100-ldp] peer 5.5.5.9 upe
                    [SPE2-vsi-v100-ldp] quit
                    [SPE2-vsi-v100] quit

                    ----End

Verifying the Configuration
                    # Check information about LDP VPWS connections on a UPE.
                    [UPE1]display mpls l2vc
                     Total LDP VC : 1 1 up        0 down

                    *client interface     : Vlanif50 is up
                     Administrator PW        : no
                     session state       : up
                     AC status          : up
                     VC state          : up
                     Label state        :0
                     Token state          :0
                     VC ID            : 100
                     VC type           : VLAN
                     destination         : 1.1.1.9
                     local VC label       : 1024        remote VC label   : 1028


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                      739
VPN Configuration
VPN Configuration                                                                                 6 VPLS Configuration

                     control word            : disable
                     remote control word : disable
                     forwarding entry            : exist
                     local group ID          :0
                     remote group ID              :0
                     local AC OAM State              : up
                     local PSN OAM State : up
                     local forwarding state : forwarding
                     local status code         : 0x0 (forwarding)
                     remote AC OAM state : up
                     remote PSN OAM state : up
                     remote forwarding state: forwarding
                     remote status code            : 0x0 (forwarding)
                     remote interface           : unknown
                     ignore standby state : no
                     BFD for PW               : unavailable
                     VCCV State              : up
                     manual fault             : not set
                     active state         : active
                     TTL Value             :1
                     link state         : up
                     local VC MTU               : 1500         remote VC MTU        : 1500
                     local VCCV             : alert ttl lsp-ping bfd
                     remote VCCV                 : alert lsp-ping bfd
                     tunnel policy name            : --
                     PW template name                  : --
                     primary or secondary : primary
                     load balance type            : flow
                     Access-port            : false
                     Switchover Flag            : false
                     VC tunnel info           : 1 tunnels
                       NO.0 TNL type             : ldp          , TNL ID : 0x0000000001004c6b42
                     create time           : 0 days, 0 hours, 6 minutes, 2 seconds
                     up time              : 0 days, 0 hours, 4 minutes, 58 seconds
                     last change time           : 0 days, 0 hours, 4 minutes, 58 seconds
                     VC last up time           : 2024/08/06 17:22:00
                     VC total up time           : 0 days, 0 hours, 4 minutes, 58 seconds
                     CKey               : 65
                     NKey                : 16777404
                     PW redundancy mode                  : frr
                     AdminPw interface              : --
                     AdminPw link state             : --
                     Forward state            : send active, receive active
                     Diffserv Mode             : uniform
                     Service Class         : cs-
                     Color             : --
                     DomainId                :-
                     Domain Name                    :-

                    The command output shows that a dynamic VPWS PW has been established and
                    the VC status is up.
                    # Check detailed information about the VSI named v100 on an SPE.
                    [SPE1] display vsi name v100
                    Vsi                       Mem PW Mac                Encap      Mtu Vsi
                    Name                         Disc Type Learn         Type      Value State
                    --------------------------------------------------------------------------
                    v100                        static ldp unqualify vlan        1500 up

                    The command output shows that the VSI named v100 is up and the corresponding
                    PW is also up.
                    # Perform a ping test to check the connectivity.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=1 ms


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                   740
VPN Configuration
VPN Configuration                                                                    6 VPLS Configuration

                        Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=1 ms
                        Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=2 ms
                        Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=1 ms

                    --- 10.1.1.2 ping statistics ---
                      5 packet(s) transmitted
                      5 packet(s) received
                      0.00% packet loss
                      round-trip min/avg/max = 1/1/2 ms


