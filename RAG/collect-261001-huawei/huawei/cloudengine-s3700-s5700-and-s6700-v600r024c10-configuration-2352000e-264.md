---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-264
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [38715, 38863]
sha256: 158335185f64b03716ffd558cae702c61573157290e53e35db440027d1167fd9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        State              : up
                        Ac Block State          : unblocked
                        Access Port           : false
                        Last Up Time             : 2014/11/10 16:37:47
                        Total Up Time            : 0 days, 0 hours, 1 minutes, 3 seconds

                     **PW Information:

                        *Peer Ip Address        : 3.3.3.9
                         PW State             : up
                         Local VC Label         : 4096
                         Remote VC Label           : 4096
                         Remote Control Word : disable
                         PW Type              : label
                         Local VCCV            : alert lsp-ping bfd
                         Remote VCCV              : alert lsp-ping bfd
                         Tunnel ID           : 0x0000000001004cab43
                         Broadcast Tunnel ID : --
                         Broad BackupTunnel ID : --
                         Ckey             : 0x6
                         Nkey              : 0x5
                         Main PW Token             : 0x1a
                         Slave PW Token           : 0x0
                         Tnl Type           : ldp
                         OutInterface          : Vlanif20
                         Backup OutInterface :
                         Stp Enable           :0
                         Mac Flapping            :0
                         PW Last Up Time           : 2014/11/10 16:38:47
                         PW Total Up Time           : 0 days, 0 hours, 0 minutes, 3 seconds

                    The command output shows that a PW to PE2 has been established for the VSI
                    named a2 and the VSI status is Up.

                    # Perform a ping test to check the connectivity.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=1 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 1/1/1 ms




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
                            #
                            return


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                              620
VPN Configuration
VPN Configuration                                                              6 VPLS Configuration


                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan 40
                        #
                        interface Vlanif40
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 1.1.1.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi a2 static
                         pwsignal ldp
                          vsi-id 2
                          peer 3.3.3.9
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 3.3.3.9
                         remote-ip 3.3.3.9
                        #
                        interface Vlanif10
                         l2 binding vsi a2
                        #
                        interface Vlanif20
                         ip address 192.168.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 20 30
                        #
                        mpls lsr-id 2.2.2.9


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   621
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

