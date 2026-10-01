---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-70
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [9335, 9525]
sha256: 16ad5a19f74ca3e86f64e52e8279168aeef395ec33c11ba27aae854543547241
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        ip vpn-instance vpn1
                         ipv4-family
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
                         ip address 12.12.12.2 255.255.255.0
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
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 200
                         peer 4.4.4.9 as-number 200
                         peer 4.4.4.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 4.4.4.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 4.4.4.9 enable
                         #
                         ipv4-family vpn-instance vpn1
                          peer 12.12.12.1 as-number 100
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
                         ipv4-family
                          route-distinguisher 200:1
                          vpn-target 2:2 export-extcommunity
                          vpn-target 2:2 import-extcommunity
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         148
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.40.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
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
                         ipv4-family vpn-instance vpn1
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
                        interface Loopback 1
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




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         149
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration




3.12 Configuring IPv4 L3VPN over MPLS Inter-AS
Option B (Basic Networking)

3.12.1 Understanding IPv4 L3VPN over MPLS Inter-AS Option
B (Basic Networking)
                    ●   Inter-AS VPN Option B Overview
                        On the inter-AS VPN Option B network shown in Figure 3-27, two ASBRs use
                        MP-EBGP to exchange labeled VPN-IPv4 routes received from local PEs in
                        their respective ASs. A VPN LSP indicates a private network tunnel, and an
                        LSP indicates a public network tunnel.

                        Figure 3-27 Inter-AS VPN Option B networking




                        In inter-AS VPN Option B, ASBRs receive all inter-AS VPN-IPv4 routes from the
                        local and external ASs and then advertise these routes. In basic MPLS VPN
                        implementation, a PE stores only the VPN routes that match the VPN targets
                        of its local VPN instances. The ASBRs are configured to store all the received
                        VPN routes, regardless of whether these routes match the VPN targets of its
                        local VPN instances.
                        The advantage of this solution is that all traffic is forwarded by ASBRs. In this
                        way, traffic is controllable, but the loads on the ASBRs are heavy. BGP routing
                        policies, such as VPN target-based filtering policies, can be configured on
                        ASBRs, so that ASBRs only save some of VPN-IPv4 routes.
                    ●   Route Advertisement in an Inter-AS VPN Option B Scenario
                        Figure 3-28 shows a route advertisement example. In this example, CE1
                        advertises route 10.1.1.1/24 to CE2. NH indicates the next hop, and L1, L2,
                        and L3 the VPN labels. This figure does not show the distribution of public
                        network IGP routes and labels.




