---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-165
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [23746, 23938]
sha256: 7d8c90f4e1e062e9008cb3245992227fe2e2ba8ded235a9cf49293d6426b7bc4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv6-family
                          route-distinguisher 200:2
                          vpn-target 2:2 export-extcommunity
                          vpn-target 2:2 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.40.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ipv6 enable
                         ipv6 address 2001:db8:3::2/64
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         374
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 200
                         peer 4.4.4.9 as-number 200
                         peer 4.4.4.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 4.4.4.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 4.4.4.9 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:db8:3::1 as-number 100
                          import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.40.1.0 0.0.0.255
                        #
                        return

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv6-family
                          route-distinguisher 200:1
                          vpn-target 2:2 export-extcommunity
                          vpn-target 2:2 import-extcommunity
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.40.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ipv6 enable
                         ipv6 address 2001:db8:2::2/64
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         375
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         ip address 4.4.4.9 255.255.255.255
                        #
                        bgp 200
                         peer 3.3.3.9 as-number 200
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 3.3.3.9 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:db8:2::1 as-number 65002
                          import-route direct
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.40.1.0 0.0.0.255
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:2::1/64
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface Loopback 1
                         ip address 11.11.11.11 255.255.255.255
                        #
                        bgp 65002
                         router-id 11.11.11.11
                         peer 2001:db8:2::2 as-number 200
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          import-route direct
                          peer 2001:db8:2::2 enable
                        #
                        return



4.13 Configuring IPv6 L3VPN over MPLS Inter-AS
Option B
Context
                    If ASBRs can manage VPN routes but there are insufficient interfaces available for
                    the dedicated use of all inter-AS VPNs, you can use inter-AS VPN Option B. This
                    solution eliminates the need to create VPN instances on ASBRs, but requires ASBRs
                    to maintain and advertise VPNv6 routes.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         376
VPN Configuration
VPN Configuration                                                           4 IPv6 L3VPN Configuration


                    On the network shown in Figure 4-8, the connected interfaces between ASBRs do
                    not need to be bound to the VPN. A single-hop MP-EBGP peer relationship is set
                    up between the ASBRs to transmit all inter-AS VPN routes.

                    Figure 4-8 Inter-AS IPv6 VPN Option B




