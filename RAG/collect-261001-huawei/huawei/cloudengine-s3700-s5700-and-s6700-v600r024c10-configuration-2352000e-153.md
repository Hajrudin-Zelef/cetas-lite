---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-153
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [21821, 22022]
sha256: a410e7043ff8c2ffafb312eefdeca284c958d4c2fe88330737eb0474f3646440
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●     PE
                          #
                          sysname PE
                          #
                          vlan batch 100 200
                          #
                          ip vpn-instance vpna
                           ipv6-family
                            route-distinguisher 100:1
                            vpn-target 100:100 export-extcommunity
                            vpn-target 100:100 import-extcommunity
                          #
                          interface Vlanif100
                           ip binding vpn-instance vpna
                           ipv6 enable
                           ipv6 address 2001:DB8:4::1/64
                          #
                          interface Vlanif200
                           ip binding vpn-instance vpna
                           ipv6 enable
                           ipv6 address 2001:DB8:1::1/64
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 100
                          #
                          interface 10GE1/0/2
                           port link-type trunk


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         345
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        bgp 100
                         router-id 1.1.1.1
                         #
                         ipv4-family unicast
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                         #
                         ipv6-family vpn-instance vpna
                          auto-frr
                          route-select delay 300
                          peer 2001:DB8:4::2 as-number 65410
                          peer 2001:DB8:1::2 as-number 65410
                        #
                        return

                    ●   CE1
                        #
                        sysname CE1
                        #
                        vlan batch 100 200
                        #
                        ospfv3 1
                         router-id 2.2.2.2
                         import-route bgp
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:DB8:4::2/64
                        #
                        interface Vlanif200
                         ipv6 enable
                         ipv6 address 2001:DB8:2::1/64
                         ospfv3 1 area 0.0.0.0
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
                        bgp 65410
                         router-id 2.2.2.2
                         peer 2001:DB8:4::1 as-number 100
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          import-route ospfv3 1 med 100
                          peer 2001:DB8:4::1 enable
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 100 200
                        #
                        ospfv3 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         346
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         router-id 3.3.3.3
                         import-route bgp
                         area 0.0.0.0
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:DB8:1::2/64
                        #
                        interface Vlanif200
                         ipv6 enable
                         ipv6 address 2001:DB8:3::1/64
                         ospfv3 1 area 0.0.0.0
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
                         ip address 3.3.3.3 255.255.255.255
                        #
                        bgp 65410
                         router-id 3.3.3.3
                         peer 2001:DB8:1::1 as-number 100
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          import-route ospfv3 1 med 500
                          peer 2001:DB8:1::1 enable
                        #
                        return

                    ●   DeviceA
                        #
                        sysname DeviceA
                        #
                        vlan batch 100 200
                        #
                        ospfv3 1
                         router-id 4.4.4.4
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:DB8:2::2/64
                         ospfv3 1 area 0.0.0.0
                        #
                        interface Vlanif200
                         ipv6 enable
                         ipv6 address 2001:DB8:3::2/64
                         ospfv3 1 area 0.0.0.0
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
                         ipv6 enable
                         ipv6 address 2001:DB8:4::1/128
                         ospfv3 1 area 0.0.0.0
                        #
                        return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         347
VPN Configuration
VPN Configuration                                                                     4 IPv6 L3VPN Configuration


4.10.7 Example for Configuring IPv6+VPNv6 Hybrid FRR
Networking Requirements
                    A CE at an IPv6 VPN site is dual-homed to two PEs, and a VPNv6 peer relationship
                    is established between the two PEs. To protect one of the links between the PEs
                    and CE, IPv6+VPNv6 hybrid FRR can be configured.
                    If one link fails, IPv6+VPNv6 hybrid FRR can quickly switch traffic destined for the
                    CE to the backup next hop (the other PE).

                         NOTE

