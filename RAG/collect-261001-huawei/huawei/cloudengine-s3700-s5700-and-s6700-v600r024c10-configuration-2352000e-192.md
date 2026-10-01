---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-192
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27994, 28179]
sha256: 3ce38b52ff62d7ecd827d1612dd11f792d3916e5adda38b451e904471eae6db2
---

                    # Configure PE2.
                    [PE2] mpls lsr-id 3.3.3.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface vlanif 20
                    [PE2-Vlanif20] mpls
                    [PE2-Vlanif20] mpls ldp
                    [PE2-Vlanif20] quit

         Step 4 Configure BGP peers to exchange VPWS information.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 3.3.3.9 as-number 100
                    [PE1-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [PE1-bgp] l2vpn-ad-family
                    [PE1-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                    [PE1-bgp-af-l2vpn-ad] peer 3.3.3.9 signaling vpws
                    [PE1-bgp-af-l2vpn-ad] quit
                    [PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.9 as-number 100
                    [PE2-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [PE2-bgp] l2vpn-ad-family
                    [PE2-bgp-af-l2vpn-ad] peer 1.1.1.9 enable
                    [PE2-bgp-af-l2vpn-ad] peer 1.1.1.9 signaling vpws
                    [PE2-bgp-af-l2vpn-ad] quit
                    [PE2-bgp] quit

         Step 5 Configure remote BGP VPWS connections.
                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] mpls l2vpn vpn1 encapsulation vlan
                    [PE1-mpls-l2vpn-vpn1] route-distinguisher 100:1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                   447
VPN Configuration
VPN Configuration                                                                                  5 VPWS Configuration

                    [PE1-mpls-l2vpn-vpn1] vpn-target 200:1 both
                    [PE1-mpls-l2vpn-vpn1] ce ce1 id 1 range 10
                    [PE1-mpls-l2vpn-vpn1-ce-ce1] connection ce-offset 2 interface vlanif 20
                    [PE1-mpls-l2vpn-vpn1-ce-ce1] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] mpls l2vpn vpn1 encapsulation vlan
                    [PE2-mpls-l2vpn-vpn1] route-distinguisher 100:1
                    [PE2-mpls-l2vpn-vpn1] vpn-target 200:1 both
                    [PE2-mpls-l2vpn-vpn1] ce ce2 id 2 range 10
                    [PE2-mpls-l2vpn-vpn1-ce-ce2] connection ce-offset 1 interface vlanif 10
                    [PE2-mpls-l2vpn-vpn1-ce-ce2] quit
                    [PE2-mpls-l2vpn-vpn1] quit

                    ----End

Verifying the Configuration
                    # Check VPWS connection information on PEs. The command outputs show that a
                    remote VPWS connection has been established and is in the up state.
                    The following example uses the command output on PE1.
                    [PE1] display mpls l2vpn connection
                    1 total connections,
                    connections: 1 up, 0 down, 0 local, 1 remote, 0 unknown

                    VPN name: vpn1,
                    1 total connections,
                    connections: 1 up, 0 down , 0 local, 1 remote, 0 unknown

                      CE name: ce1, id: 1,
                      Rid type status peer-id          route-distinguisher interface
                      primary or not
                    ----------------------------------------------------------------------------
                      2 rmt up         3.3.3.9       100:1             Vlanif20
                      primary

                    # CE1 and CE2 can ping each other successfully.
                    [CE1] ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=7 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=3 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=4 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=2 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=3 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 2/3/7 ms


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan batch 20
                          #
                          interface Vlanif20
                           ip address 10.1.1.1 255.255.255.0


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                    448
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        return
                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 10
                        #
                        interface Vlanif10
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface Vlanif20
                        #
                        interface Vlanif10
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
                        mpls l2vpn vpn1 encapsulation vlan
                         route-distinguisher 100:1
                         vpn-target 200:1 import-extcommunity
                         vpn-target 200:1 export-extcommunity
                         ce ce1 id 1 range 10 default-offset 0
                          connection ce-offset 2 interface Vlanif20
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   449

