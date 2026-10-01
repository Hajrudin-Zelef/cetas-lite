---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-156
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22324, 22467]
sha256: 7632a68fb314d791910bdb3aead632c7159885e00d37b4c91e60d14d1a2add6c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Verifying the Configuration
                    After completing the configuration, run the display ipv6 routing-table vpn-
                    instance verbose command on PE2. The command output shows information
                    about the primary and backup routes destined for the CE's loopback interface in
                    the routing table of the VPN instance IPv6 address family. Because the EBGP route
                    takes precedence over the IBGP route, PE2 prefers the EBGP route sent from the
                    CE and uses the IBGP route sent from PE3 as the backup route. The fields in
                    boldface indicate the backup next hop, backup label, and backup tunnel ID. The
                    command output shows that a hybrid FRR entry has been generated.
                    <PE2> display ipv6 routing-table vpn-instance vpn1 2001:db8:0:1:2::1 verbose
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 1

                    Destination : 2001:db8:0:1:2::1           PrefixLength : 128
                    NextHop        : 2001:db8:1::1            Preference : 100
                    Neighbour : ::                       ProcessID : 0
                    Label      : NULL                     Protocol     : EBGP
                    State     : Active Adv Relied            Cost        :0
                    Entry ID    : 14                     EntryFlags : 0x00000000
                    Reference Cnt: 0                       Tag        :0
                    IndirectID : 0x8a9                     Age        : 3sec
                    RelayNextHop : 2001:db8:1::1                 TunnelID    : 0x0
                    Interface : Vlanif200          Flags       : RD
                    BkNextHop : ::                        BkInterface : LDP LSP
                    BkLabel       : 17                    BkTunnelID : 0x0
                    BkPETunnelID : 0x0000000001004c4b44               BkIndirectID : 0xae

                    Run the shutdown and then display ipv6 routing-table vpn-instance verbose
                    commands on VLANIF 200 of PE2. The command output shows that the next hop
                    to the loopback interface on the CE is changed to PE3.
                    <PE2> display ipv6 routing-table vpn-instance vpn1 2001:db8:0:1:2::1 verbose
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 1

                    Destination : 2001:db8::0:1:2::1          PrefixLength : 128
                    NextHop       : ::FFFF:3.3.3.3           Preference : 255
                    Neighbour : ::                       ProcessID : 0
                    Label      : 17                    Protocol     : EBGP
                    State     : Active Adv Relied           Cost         :0
                    Entry ID    :0                      EntryFlags : 0x00000000
                    Reference Cnt: 0                      Tag          :0
                    IndirectID : 0xa5                     Age         : 9sec
                    RelayNextHop : ::                      TunnelID      : 0x0000000001004c4b42
                    Interface : Vlanif200          Flags      : RD

                    IPv6+VPNv6 hybrid FRR has taken effect on PE2.

Configuration Scripts
                    ●     PE1
                          #
                          sysname PE1
                          #
                          vlan batch 200 300
                          #
                          ip vpn-instance vpn1
                           ipv6-family
                            route-distinguisher 100:1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         353
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.10.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip address 10.20.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        bgp 100
                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         peer 3.3.3.3 as-number 100
                         peer 3.3.3.3 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 2.2.2.2 enable
                          peer 3.3.3.3 enable
                        #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 2.2.2.2 enable
                          peer 3.3.3.3 enable
                        #
                         ipv6-family vpn-instance vpn1
                         import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                          network 10.20.1.0 0.0.0.3
                          network 1.1.1.1 0.0.0.0
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpn1
                         ipv6-family
                          route-distinguisher 100:2
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        mpls lsr-id 2.2.2.2
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         354
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

