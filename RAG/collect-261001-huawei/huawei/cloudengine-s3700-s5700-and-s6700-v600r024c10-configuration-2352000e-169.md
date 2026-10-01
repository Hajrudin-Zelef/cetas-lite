---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-169
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [24423, 24591]
sha256: c56012378d4ad831700457ddd62a46a92390d878e91d857ff2a02d07fca8cd59
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    *>i Network : 2001:db8:1::                       PrefixLen : 64
                       NextHop : ::FFFF:1.1.1.1                     LocPrf : 100
                       MED      :0                              PrefVal : 0
                       Label : 21/23
                       Path/Ogn : 65001?
                    Route Distinguisher: 200:2

                    *> Network : 2001:db8:2::                        PrefixLen : 64
                       NextHop : ::FFFF:192.168.1.2                    LocPrf :
                       MED    :                                 PrefVal : 0
                       Label : 25/25
                       Path/Ogn : 200 65002?


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan batch 100
                          #
                          interface Vlanif100
                           ipv6 enable
                           ipv6 address 2001:db8:1::1/64
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 100
                          #
                          interface LoopBack1
                           ip address 10.10.10.10 255.255.255.255
                          #
                          bgp 65001
                           router-id 10.10.10.10
                           peer 2001:db8:1::2 as-number 100
                           #
                           ipv4-family unicast
                           #
                           ipv6-family unicast
                            import-route direct
                            peer 2001:db8:1::2 enable
                          #
                          return
                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 100 200
                          #
                          ip vpn-instance vpn1
                           ipv6-family
                            route-distinguisher 100:1


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                          386
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.16.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ipv6 enable
                         ipv6 address 2001:db8:1::2/64
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.2 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 2.2.2.2 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:db8:1::1 as-number 65001
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 172.16.1.0 0.0.0.255
                        #
                        return
                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.16.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 192.168.1.1 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         387
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         mpls
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        bgp 100
                         peer 192.168.1.2 as-number 200
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 192.168.1.2 enable
                          peer 1.1.1.1 enable
                         #
                         ipv6-family vpnv6
                          undo policy vpn-target
                          peer 1.1.1.1 enable
                          peer 192.168.1.2 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 172.16.1.0 0.0.0.255
                        #
                        return

