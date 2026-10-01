---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-306
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45333, 45485]
sha256: 621610f56fd8025cd6061d45d1410ab3300ff0b04325c4455804c9443519de58
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     Access-port        : false
                     VC tunnel info       : 1 tunnels
                     NO.0 TNL Type           : static-lsp , TNL ID : 0x2
                     Create time         : 0 days, 0 hours, 1 minutes, 4 seconds
                     UP time           : 0 days, 0 hours, 1 minutes, 4 seconds
                     Last change time        : 0 days, 0 hours, 1 minutes, 4 seconds
                     VC last up time       : 2014/11/13 14:07:19
                     VC total up time       : 0 days, 0 hours, 1 minutes, 4 seconds
                     CKey             :4
                     NKey             :3
                     BFD for PW           : unavailable

                    The command output shows that a static VPWS PW has been established and the
                    VC status is up. The command output on UPE1 is used as an example.
                    # Check detailed information about the VSI named V100 on an SPE.
                    [SPE1]display vsi name V100 verbose

                    ***VSI Name                : V100
                       Work Mode                 : normal
                       Administrator VSI          : no
                       Isolate Spoken          : disable
                       VSI Index            :4
                       PW Signaling             : ldp
                       Member Discovery Style : static
                       PW MAC Learn Style             : unqualify
                       Encapsulation Type           : vlan
                       MTU                 : 1500
                       Diffserv Mode            : uniform
                       Service Class         : cs-
                       Color             : --
                       DomainId               :-
                       Domain Name                  :-
                       Ignore AcState           : disable
                       P2P VSI             : disable
                       Multicast Fast Switch : disable
                       Create Time            : 0 days, 0 hours, 17 minutes, 39 seconds
                       VSI State           : up
                       Resource Status           : --

                      VSI ID            : 100
                      LDP MAC-WITHDRAW                : mac-withdraw Enable
                     *Peer Router ID          : 3.3.3.9
                      Negotiation-vc-id        : 100
                      Encapsulation Type         : vlan
                      primary or secondary : primary
                      ignore-standby-state : no
                      VC Label            : 1052
                      Peer Type            : dynamic
                      Session            : up
                      Tunnel ID            : 0x0000000001004c6b43
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      CKey              : 65
                      NKey               : 16777691
                      Stp Enable            :0
                      PwIndex              : 65
                      Control Word            : disable
                      BFD for PW             : unavailable
                     *Peer Router ID          : 4.4.4.9
                      Negotiation-vc-id        : 100
                      Encapsulation Type         : vlan
                      primary or secondary : primary
                      ignore-standby-state : no
                      VC Label            : 100
                      Peer Type            : static
                      Tunnel ID            : 0x000000000400895441
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             729
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration

                      CKey                : 66
                      NKey                 : 16777692
                      Stp Enable             :0
                      PwIndex                : 66
                      Control Word              : disable
                      BFD for PW               : unavailable

                     **PW Information:

                     *Peer Ip Address        : 3.3.3.9
                      PW State             : up
                      Local VC Label         : 1052
                      Remote VC Label           : 1181
                      Remote Control Word : disable
                      PW Type              : label
                      Local VCCV            : alert lsp-ping bfd
                      Remote VCCV              : alert lsp-ping bfd
                      Tunnel ID           : 0x0000000001004c6b43
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      Ckey             : 65
                      Nkey              : 16777691
                      Main PW Token             : 0x0
                      Slave PW Token           : 0x0
                      Tnl Type           : ldp
                      OutInterface          : --
                      Backup OutInterface : --
                      Stp Enable           :0
                      Mac Flapping            :0
                      PW Last Up Time           : 2024/08/05 18:12:29
                      PW Total Up Time           : 0 days, 0 hours, 9 minutes, 25 seconds
                     *Peer Ip Address        : 4.4.4.9
                      PW State             : up
                      Local VC Label         : 100
                      Remote VC Label           : 100
                      Remote Control Word : disable
                      PW Type              : MEHVPLS
                      Local VCCV            : alert lsp-ping bfd
                      Remote VCCV              : alert lsp-ping bfd
                      Tunnel ID           : 0x000000000400895441
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      Ckey             : 66
                      Nkey              : 16777692
                      Main PW Token             : 0x0
                      Slave PW Token           : 0x0
                      Tnl Type           : static-lsp
                      OutInterface          : --
                      Backup OutInterface : --
                      Stp Enable           :0
                      Mac Flapping            :0
                      PW Last Up Time           : 2024/08/05 18:04:28
                      PW Total Up Time           : 0 days, 0 hours, 17 minutes, 26 seconds

                    The command output shows that the VSI named V100 is up and the
                    corresponding PW is also up.
                    # Perform a ping test to check the connectivity.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=90 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=77 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=34 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=46 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=94 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                               730
VPN Configuration
VPN Configuration                                                                                     6 VPLS Configuration

                        0.00% packet loss
                        round-trip min/avg/max = 34/68/94 ms

