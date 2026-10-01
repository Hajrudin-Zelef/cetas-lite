---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-271
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [39818, 40001]
sha256: 80cdac352c31ed341bab7533131af34f21e73a2e500dfab58f861f5fb15952cb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                      Interface Name         : Vlanif10
                      State          : up
                      Ac Block State      : unblocked
                      Access Port       : false
                      Last Up Time         : 2024/08/03 15:22:53
                      Total Up Time        : 0 days, 0 hours, 2 minutes, 32 seconds

                     **PW Information:

                     *Peer Ip Address        : 3.3.3.9
                      PW State             : up
                      Local VC Label         : 1016578
                      Remote VC Label           : 1016577
                      Remote Control Word : disable
                      Negotiated Control Word: disable
                      PW Type              : label
                      Tunnel ID           : 0x0000000001004cab43
                      Broadcast Tunnel ID : --
                      Broad BackupTunnel ID : --
                      Ckey             : 513
                      Nkey              : 16777524
                      Main PW Token             : 0x0
                      Slave PW Token           : 0x0
                      Tnl Type           : ldp
                      OutInterface          : --
                      Backup OutInterface : --
                      Stp Enable           :0
                      Mac Flapping            :0


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            638
VPN Configuration
VPN Configuration                                                                        6 VPLS Configuration

                        PW Last Up Time       : 2024/08/03 15:24:29
                        PW Total Up Time      : 0 days, 0 hours, 0 minutes, 56 seconds

                    The command output shows that a PW to PE2 has been established for the VSI
                    named bgp1, and the VSI status is Up.
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
                    ●     CE1
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

                    ●     CE2
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

                    ●     PE1
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
                          vsi bgp1 auto
                           pwsignal bgp
                            route-distinguisher 8.8.8.1:1


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                           639
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                          site 1 range 5 default-offset 0
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         l2 binding vsi bgp1
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
                        bgp 100
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 3.3.3.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 3.3.3.9 enable
                          peer 3.3.3.9 signaling vpls
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
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif20
                         ip address 192.168.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         ip address 192.168.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   640
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

