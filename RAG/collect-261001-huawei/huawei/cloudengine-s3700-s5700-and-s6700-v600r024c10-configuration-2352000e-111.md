---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-111
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [15516, 15670]
sha256: c5ff13f76d527ebc2d071cd6d43618b2430e511d591f3bbcc15f26b3425a7fe6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Hub-PE
                        #
                        sysname Hub-PE
                        #
                        vlan batch 100 200 300 400
                        #
                        ip vpn-instance vpn_in
                         ipv4-family
                          route-distinguisher 100:21
                          vpn-target 100:1 import-extcommunity
                        #
                        ip vpn-instance vpn_out
                         ipv4-family
                          route-distinguisher 100:22
                          vpn-target 200:1 export-extcommunity
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 20.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 11.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpn_in
                         ip address 10.2.1.2 255.255.255.0
                        #
                        interface Vlanif400


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         246
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         ip binding vpn-instance vpn_out
                         ip address 10.3.1.2 255.255.255.0
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
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 3.3.3.9 enable
                        #
                         ipv4-family vpn-instance vpn_in
                          peer 10.2.1.1 as-number 65430
                        #
                         ipv4-family vpn-instance vpn_out
                          peer 10.3.1.1 as-number 65430
                          peer 10.3.1.1 allow-as-loop
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 20.1.1.0 0.0.0.255
                          network 11.1.1.0 0.0.0.255
                        #
                        return


3.15.8 Example 1 for Configuring Hub-Spoke (Single Link
Between the Hub-PE and Hub-CE)
Networking Requirements
                    On the network shown in Figure 3-49, the communication between the Spoke-
                    CEs is controlled by the Hub-CE at the central site. In other words, the traffic
                    between Spoke-CEs is forwarded also through the Hub-CE, not only through the
                    Hub-PE. The Hub-PE connects to the Hub-CE over a single link.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         247
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


                    Figure 3-49 Example 1 for configuring hub-spoke (single link between the Hub-PE
                    and Hub-CE)
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Precautions
                    Note the following during the configuration:
                    ●    The import and export VPN targets configured on a Spoke-PE are different.
                    ●    A VPN instance (vpnhub) is created on the Hub-PE. The VPN targets received
                         by vpnhub are the VPN targets advertised by the two Spoke-PEs; the VPN
                         targets advertised by vpnhub are the VPN targets received by the two Spoke-
                         PEs and are different from the VPN targets received by vpnhub.
                    ●    In this scenario, to avoid loops, ensure that connected interfaces have STP
                         disabled and are removed from VLAN 1. If STP is enabled and VLANIF
                         interfaces of devices are used to construct a Layer 3 ring network, an
                         interface on the network will be blocked. As a result, Layer 3 services on the
                         network cannot run normally.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     248
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Establish MP-IBGP peer relationships between the Hub-PE and Spoke-PEs.
                         (There is no need to establish an MP-IBGP peer relationship or exchange VPN
                         routing information between the two Spoke-PEs.)
                    2.   Create a VPN instance on each PE and configure pop-go on the Hub-PE.
                    3.   Configure EBGP connections between CEs and PEs. Configure the Hub-CE to
                         advertise the default route to the Hub-PE.

