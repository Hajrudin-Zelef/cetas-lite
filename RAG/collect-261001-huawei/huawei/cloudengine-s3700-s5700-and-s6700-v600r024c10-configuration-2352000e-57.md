---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-57
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7326, 7528]
sha256: 53dba908a4ebaaa419fdcd68a4dc320e0406692c7caab21adba9477c74677bb5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         ipv4-family vpnv4
                          policy vpn-target
                          peer 2.2.2.2 enable
                          peer 3.3.3.3 enable
                         #
                         ipv4-family vpn-instance vpn1
                          auto-frr
                          route-select delay 300
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                          network 10.20.1.0 0.0.0.3
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 100:2
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.10.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ip address 10.1.1.2 255.255.255.252
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
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpn-instance vpn1
                          peer 10.1.1.1 as-number 65410
                        #
                        ospf 1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         117
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.10.1.0 0.0.0.3
                        #
                        return

                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 100:3
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.20.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ip address 10.2.1.2 255.255.255.252
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
                        bgp 100
                         peer 1.1.1.1 as-number 100
                         peer 1.1.1.1 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.1 enable
                         #
                         ipv4-family vpn-instance vpn1
                          peer 10.2.1.1 as-number 65410
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.20.1.0 0.0.0.3
                        #
                        return

                    ●   CE
                        #
                        sysname CE
                        #
                        vlan batch 100 200



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         118
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration

                         #
                         interface Vlanif100
                          ip address 10.1.1.1 255.255.255.252
                         #
                         interface Vlanif200
                          ip address 10.2.1.1 255.255.255.252
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
                          ip address 11.11.11.11 255.255.255.255
                         #
                         bgp 65410
                          peer 10.1.1.2 as-number 100
                          peer 10.2.1.2 as-number 100
                          #
                          ipv4-family unicast
                           network 11.11.11.11 255.255.255.255
                           peer 10.1.1.2 enable
                           peer 10.2.1.2 enable
                         #
                         return


3.10.7 Example for Configuring VPN IP FRR
Networking Requirements
                    As shown in Figure 3-20, it is required that a VPN backup outbound interface and
                    a backup next hop be configured on the PE so that Link_A functions as the backup
                    of Link_B. If Link_B fails, traffic can be fast switched to Link_A.

                    Figure 3-20 Configuring VPN IP FRR
                         NOTE

                    In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    119

