---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-82
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [11278, 11478]
sha256: 44432a2be4bff222029818c1f6611eece34fbfcb89ba7f70b7f469f99363fdf0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        interface Vlanif100
                         ip address 10.40.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.12.12.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ip address 10.4.1.2 255.255.255.0
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
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 200
                         peer 10.12.12.1 as-number 100
                         peer 4.4.4.9 as-number 200
                         peer 4.4.4.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 10.12.12.1 enable
                          peer 4.4.4.9 enable
                         #
                         ipv4-family vpnv4
                          undo policy vpn-target
                          peer 4.4.4.9 enable
                          peer 10.12.12.1 enable
                         #
                         ipv4-family vpn-instance vpna
                          peer 10.4.1.1 as-number 65004
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.40.1.0 0.0.0.255
                        #
                        return

                    ●   CE4
                        #
                        sysname CE4
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.4.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface Loopback1
                         ip address 44.44.44.44 255.255.255.255
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         179
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        bgp 65004
                         peer 10.4.1.2 as-number 200
                         #
                         ipv4-family unicast
                          network 44.44.44.44 255.255.255.255
                          peer 10.4.1.2 enable
                        return

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 200:1
                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
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
                         ip binding vpn-instance vpna
                         ip address 10.2.1.2 255.255.255.0
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
                         ip address 4.4.4.9 255.255.255.255
                        #
                        bgp 200
                         peer 3.3.3.9 as-number 200
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpn-instance vpna
                          peer 10.2.1.1 as-number 65002
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



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         180
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.2.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface Loopback1
                         ip address 22.22.22.22 255.255.255.255
                        #
                        bgp 65002
                         peer 10.2.1.2 as-number 200
                         #
                         ipv4-family unicast
                          network 22.22.22.22 255.255.255.255
                          peer 10.2.1.2 enable
                        #
                        return



3.14 Configuring IPv4 L3VPN over MPLS HVPN

3.14.1 Understanding IPv4 L3VPN over MPLS HVPN

Context
                    Currently, the hierarchical architecture is generally used in networking design. For
                    example, the typical architecture of a MAN consists of three layers: access layer,
                    aggregation layer, and core layer. On the network shown in Figure 3-32, all PEs
                    reside on the same plane and must provide the following functions:
                    ●   Provide access services for users. This function requires each PE to provide a
                        large number of interfaces.
                    ●   Manage and advertise VPN routes and process user packets. This function
                        requires each PE to have a high-capacity memory and strong forwarding
                        capabilities.

                    Figure 3-32 Basic IPv4 L3VPN network structure




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          181

