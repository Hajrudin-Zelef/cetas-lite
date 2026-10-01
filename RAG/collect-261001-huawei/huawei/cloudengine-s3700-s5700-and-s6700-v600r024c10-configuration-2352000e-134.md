---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-134
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18999, 19177]
sha256: daf1874afe1035557dbd0768048c6b104dce2ef9e6110ea8091495ffbd3d5107
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   CE1
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
                         ipv6 enable
                         ipv6 address 2001:db8:8::1/64
                        #
                        bgp 65410
                         router-id 10.10.10.10
                         peer 2001:db8:1::2 as-number 100
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          network 2001:db8:8:: 64
                          import-route direct
                          peer 2001:db8:1::2 enable
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



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         301
VPN Configuration
VPN Configuration                                                              4 IPv6 L3VPN Configuration

                         ipv6 address 2001:db8:3::1/64
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:db8:8::1/64
                        #
                        ipv6 route-static :: 0 2001:db8:3::2
                        #
                        return

                    ●   CE3
                        #
                        sysname CE3
                        #
                        vlan batch 100
                        #
                        ospfv3 1
                         router-id 22.22.22.22
                         area 0.0.0.0
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:4::1/64
                         ospfv3 1 area 0.0.0.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:db8:9::1/64
                         ospfv3 1 area 0.0.0.0
                        #
                        return

                    ●   CE4
                        #
                        sysname CE4
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:5::1/64
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ipv6 enable
                         ipv6 address 2001:db8:9::1/64
                        #
                        bgp 65420
                         router-id 33.33.33.33
                         peer 2001:db8:5::2 as-number 100
                         #
                         ipv4-family unicast
                         #
                         ipv6-family unicast
                          import-route direct
                          peer 2001:db8:5::2 enable
                        #
                        return




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         302
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration




4.7 Configuring Route Import Between Instances

4.7.1 Understanding Route Import Between Different
Instances
                    In IPv6 L3VPN networking, IPv6 VPN users cannot communicate with IPv6 public
                    network users. In addition, IPv6 users in one VPN can communicate with those in
                    another VPN only if the two VPNs have matching VPN targets. To enable IPv6
                    VPN users to communicate with IPv6 public network users or IPv6 users of two
                    VPNs with unmatching VPN targets to communicate, configure route import
                    between instances.
                    Route import between instances is classified into the following types:
                    ●   Route import between the public network instance and a VPN instance
                        (including importing public network routes to a VPN instance and importing
                        VPN instance routes to the public network instance): IPv6 routes are first
                        imported to the VPN or public network instance's corresponding routing table.
                        For example, VPN OSPFv3 routes are first imported into the public network
                        instance's OSPFv3 routing table. If the imported routes are preferred in the
                        corresponding routing table, they will also be imported into the VPN or public
                        network instance's IPv6 routing table to guide traffic forwarding. In addition,
                        these routes can be advertised to other devices on the network.
                    ●   Route import between different VPN instances: For example, VPNA's OSPFv3
                        routes are first imported into VPNB's OSPFv3 routing table. If the imported
                        routes are preferred in the corresponding routing table, they will also be
                        imported into the VPN instance's IPv6 routing table to guide traffic
                        forwarding. In addition, these routes can be advertised to other devices on the
                        network.

4.7.2 Configuring Route Import from an IPv6 VPN Instance to
the Public Network Instance's Routing Tables

Prerequisites
                    Before configuring route import from an IPv6 VPN instance into the public
                    network instance's routing tables, you have completed the following task:
                    ●   4.5.1 Configuring an IPv6 VPN Instance on a PE
                    ●   4.5.2 Binding an Interface to the IPv6 VPN Instance

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Import different types of IPv6 routes from a VPN instance into the public network
                instance's corresponding routing tables.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         303
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


                         NOTE

