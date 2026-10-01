---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-133
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18816, 18998]
sha256: 969b23a18d54edda5c8b678913b07fd293b9bd030aece08bbebb4effe788814a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         ipv6 enable
                         ipv6 address 2001:db8:3::2/64
                        #
                        interface Vlanif300
                         ip address 10.11.11.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
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
                         ip address 1.1.1.9 255.255.255.255
                         isis enable 1
                        #
                        bgp 100
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 3.3.3.9 enable
                         #
                         ipv6-family vpn-instance vpna
                          peer 2001:db8:1::1 as-number 65410
                         #
                         ipv6-family vpn-instance vpnb
                          import-route static
                        #
                         ipv6 route-static vpn-instance vpnb 2001:db8:8:: 64 2001:db8:3::1
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 20.2222.2222.2222.00
                        #
                        interface Vlanif100
                         ip address 10.11.11.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.12.12.1 255.255.255.0
                         isis enable 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        299
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         mpls
                         mpls ldp
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
                         isis enable 1
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200 300
                        #
                        ip vpn-instance vpna
                         ipv6-family
                          route-distinguisher 100:2
                          vpn-target 33:33 export-extcommunity
                          vpn-target 22:22 import-extcommunity
                        #
                        ip vpn-instance vpnb
                         ipv6-family
                          route-distinguisher 100:4
                          vpn-target 55:55 export-extcommunity
                          vpn-target 44:44 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 30.3333.3333.3333.00
                        #
                        ospfv3 1 vpn-instance vpna
                         router-id 10.10.11.11
                         import-route bgp
                         area 0.0.0.0
                        #
                        interface Vlanif100
                         ip binding vpn-instance vpnb
                         ipv6 enable
                         ipv6 address 2001:db8:5::2/64
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpna
                         ipv6 enable
                         ipv6 address 2001:db8:4::2/64
                         ospfv3 1 area 0.0.0.0
                        #
                        interface Vlanif300
                         ip address 10.12.12.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         300
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                         isis enable 1
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.9 enable
                         #
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 1.1.1.9 enable
                         #
                         ipv6-family vpn-instance vpna
                          import-route ospfv3 1
                         #
                         ipv6-family vpn-instance vpnb
                          peer 2001:db8:5::1 as-number 65420
                        #
                        return

