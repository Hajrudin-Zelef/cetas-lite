---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-164
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [23558, 23745]
sha256: d83ad7e8f21cc1804abdc8ce664f972309466a834b4bb91e3afbad3f2a37480d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Total Number of Routes: 4
                    *>i Network : 2001:db8:1::                    PrefixLen : 64
                       NextHop : ::FFFF:1.1.1.9                  LocPrf : 100
                       MED     :0                            PrefVal : 0
                       Label : 105472
                       Path/Ogn : ?
                    *> Network : 2001:db8:2::                     PrefixLen : 64
                       NextHop : 2001:db8:3::2                    LocPrf :
                       MED     :                             PrefVal : 0
                       Label : NULL
                       Path/Ogn : 200 ?
                    *> Network : 2001:db8:3::                     PrefixLen : 64
                       NextHop : ::                          LocPrf :
                       MED     :0                            PrefVal : 0
                       Label : NULL
                       Path/Ogn : ?
                    *
                       NextHop : 2001:db8:3::2                    LocPrf   :
                       MED     :0                            PrefVal : 0
                       Label : NULL
                       Path/Ogn : 200 ?


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
                            vpn-target 1:1 export-extcommunity
                            vpn-target 1:1 import-extcommunity


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                            372
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.10.1.2 255.255.255.0
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
                         ip address 1.1.1.9 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.9 as-number 100
                         peer 2.2.2.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 2.2.2.9 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:db8:1::1 as-number 65001
                          import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return
                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv6-family
                          route-distinguisher 100:2
                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.10.1.1 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         373
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ipv6 enable
                         ipv6 address 2001:db8:3::1/64
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
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          import-route direct
                          peer 1.1.1.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.9 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:db8:3::2 as-number 200
                          import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return

