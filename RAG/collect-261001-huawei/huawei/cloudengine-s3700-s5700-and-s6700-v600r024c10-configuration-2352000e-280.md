---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-280
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41145, 41291]
sha256: a7707393321eee75cb2491313bbe804d596b39da23dae4037b23122086648d95
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         Broadcast Tunnel ID : 0x80003f
                         Broad BackupTunnel ID : 0x0
                         Ckey             : 0x2
                         Nkey              : 0x1
                         Main PW Token             : 0x80003f
                         Slave PW Token           : 0x0
                         Tnl Type           : ldp
                         OutInterface          : Vlanif20
                         Backup OutInterface :
                         Stp Enable           :0
                         PW Last Up Time           : 2012/07/06 16:18:23
                         PW Total Up Time           : 0 days, 1 hours, 22 minutes, 13 seconds
                        *Peer Ip Address        : 3.3.3.9
                         PW State             : up
                         Local VC Label         : 1025
                         Remote VC Label           : 1025
                         PW Type              : label
                         Local VCCV            : alert lsp-ping bfd
                         Remote VCCV              : alert lsp-ping bfd
                         Tunnel ID           : 0x800033
                         Broadcast Tunnel ID : 0x800033
                         Broad BackupTunnel ID : 0x0
                         Ckey             : 0x4
                         Nkey              : 0x3
                         Main PW Token             : 0x800033
                         Slave PW Token           : 0x0
                         Tnl Type           : ldp
                         OutInterface          : Vlanif30
                         Backup OutInterface :
                         Stp Enable           :0
                         PW Last Up Time           : 2012/07/06 15:56:46
                         PW Total Up Time           : 0 days, 1 hours, 0 minutes, 45 seconds

                    The command output shows that a PW to PE2 and a PW to PE3 have been
                    established for the VSI named vplsad1, and the VSI status is Up. A PW is also
                    established between PE2 and PE3, and the VSI status is Up.

                    # Perform a ping test to check the connectivity.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=140 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=140 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=140 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=190 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=110 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 110/144/190 ms


Configuration Scripts
                    ●       CE1
                            #
                            sysname CE1
                            #
                            vlan 10
                            #
                            interface Vlanif10
                             ip address 10.1.1.1 255.255.255.0
                            #
                            interface 10GE1/0/1
                             port link-type trunk
                             port trunk allow-pass vlan 10


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                661
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        #
                        return
                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan 50
                        #
                        interface Vlanif50
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        return
                    ●   CE3
                        #
                        sysname CE3
                        #
                        vlan 60
                        #
                        interface Vlanif60
                         ip address 10.1.1.3 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 60
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 10 20 30
                        #
                        mpls lsr-id 1.1.1.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi vplsad1
                         bgp-ad
                          vpls-id 172.16.1.1:1
                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         l2 binding vsi vplsad1
                        #
                        interface Vlanif20
                         ip address 172.16.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         ip address 172.16.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   662
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

