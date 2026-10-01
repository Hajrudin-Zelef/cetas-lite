---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-317
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [47123, 47322]
sha256: 2f3d0fde8ebd4402346ad5ba76644a2fee15159da5d9ba190b0f588e16f36b1c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The command output shows that a PW to PE2 has been established for the VSI
                    named a2 and the VSI status is Up.
                    # Perform a ping test to check the connectivity.
                    CE1 can ping CE2, CE3, and CE4 successfully. The following shows the ping result
                    from CE1 to CE2.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=254 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=254 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=254 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=254 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=254 time=1 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 1/1/1 ms

                    CE2 can ping CE3 and CE4 successfully. The following shows the ping result from
                    CE2 to CE3.
                    [CE2] ping 10.1.1.3
                     PING 10.1.1.3: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.3: bytes=56 Sequence=1 ttl=254 time=1 ms
                      Reply from 10.1.1.3: bytes=56 Sequence=2 ttl=254 time=1 ms
                      Reply from 10.1.1.3: bytes=56 Sequence=3 ttl=254 time=1 ms
                      Reply from 10.1.1.3: bytes=56 Sequence=4 ttl=254 time=1 ms
                      Reply from 10.1.1.3: bytes=56 Sequence=5 ttl=254 time=1 ms

                     --- 10.1.1.3 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 1/1/1 ms

                    CE3 and CE4 cannot ping each other. The following shows the ping result from
                    CE3 to CE4.
                    [CE3] ping 10.1.1.4
                     PING 10.1.1.4: 56 data bytes, press CTRL_C to break
                      Request time out
                      Request time out
                      Request time out
                      Request time out
                      Request time out



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                               758
VPN Configuration
VPN Configuration                                                               6 VPLS Configuration

                    --- 10.1.1.4 ping statistics ---
                      5 packet(s) transmitted
                      0 packet(s) received
                      100.00% packet loss


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan batch 10
                         #
                         interface Vlanif10
                          ip address 10.1.1.1 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 10
                         #
                         return

                    ●    CE2
                         #
                         sysname CE2
                         #
                         vlan batch 100
                         #
                         interface Vlanif100
                          ip address 10.1.1.2 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 100
                         #
                         return

                    ●    CE3
                         #
                         sysname CE3
                         #
                         vlan batch 200
                         #
                         interface Vlanif200
                          ip address 10.1.1.3 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 200
                         #
                         return

                    ●    CE4
                         #
                         sysname CE4
                         #
                         vlan batch 300
                         #
                         interface Vlanif300
                          ip address 10.1.1.4 255.255.255.0
                         #
                         interface 10GE1/0/1
                          port link-type trunk
                          port trunk allow-pass vlan 300
                         #
                         return

                    ●    Switch

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                   759
VPN Configuration
VPN Configuration                                                             6 VPLS Configuration

                        #
                        sysname Switch
                        #
                        vlan batch 100 200 300
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100 200 300
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/5
                         port link-type trunk
                         port trunk allow-pass vlan 300
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
                         ip address 8.1.1.1 255.255.255.0
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
                          network 8.1.1.0 0.0.0.255
                        #
                        return



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   760

