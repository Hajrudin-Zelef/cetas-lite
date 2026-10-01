---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-42
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5023, 5205]
sha256: 9c4bbdf679c91e7c28471f2887346eb14c0e724bbbb31ee9c268e408548ec7e3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                          peer 1.1.1.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.9 enable
                         #
                         ipv4-family vpn-instance vpna
                          import-route direct
                          peer 10.3.1.1 as-number 65430
                         #
                         ipv4-family vpn-instance vpnb
                          import-route direct
                          peer 10.4.1.1 as-number 65440
                        #
                        ospf 1
                         area 0.0.0.0
                          network 12.12.12.0 0.0.0.255
                          network 3.3.3.9 0.0.0.0
                        #
                        return

                    ●   CE1
                        #
                        sysname CE1
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 11.11.11.11 255.255.255.255
                        #
                        bgp 65410
                         peer 10.1.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 11.11.11.11 255.255.255.255
                          peer 10.1.1.2 enable
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 200
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.0
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 22.22.22.22 255.255.255.255
                        #
                        bgp 65420
                         peer 10.2.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 22.22.22.22 255.255.255.255
                          peer 10.2.1.2 enable
                        #
                        return



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          79
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                    ●   CE3
                        #
                        sysname CE3
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.3.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 33.33.33.33 255.255.255.255
                        #
                        bgp 65430
                         peer 10.3.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 33.33.33.33 255.255.255.255
                          peer 10.3.1.2 enable
                        #
                        return

                    ●   CE4
                        #
                        sysname CE4
                        #
                        vlan batch 300
                        #
                        interface Vlanif300
                         ip address 10.4.1.1 255.255.255.0
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 44.44.44.44 255.255.255.255
                        #
                        bgp 65440
                         peer 10.4.1.2 as-number 100
                         #
                         ipv4-family unicast
                          network 44.44.44.44 255.255.255.255
                          peer 10.4.1.2 enable
                        #
                        return



3.7 Importing Routes Between Instances

3.7.1 Understanding Route Import Between Different
Instances

                    In IPv4 L3VPN networking, IPv4 VPN users cannot communicate with IPv4 public
                    network users. In addition, IPv4 users in one VPN can communicate with those in
                    another VPN only if the two VPNs have matching VPN targets. To address this
                    issue, route import between different instances can be configured.

                    Route import between instances is classified into the following types:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          80
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


                    ●   Route import between a public network and a VPN instance: involves
                        importing public network routes into the routing table of a VPN instance and
                        importing VPN instance routes into the routing table of the public network.
                        IPv4 routes are first imported to the VPN or public network instance's
                        corresponding routing table. For example, VPN OSPF routes are first imported
                        into the public network instance's OSPF routing table. If the imported routes
                        are preferred in the corresponding routing table, they will also be imported
                        into the VPN or public network instance's IPv4 routing table to guide traffic
                        forwarding. In addition, these routes can be advertised to other devices on the
                        network.
                    ●   Route import between different VPN instances: For example, VPN A's OSPF
                        routes are first imported into VPN B's OSPF routing table. If the imported
                        routes are preferred in the corresponding routing table, they will also be
                        imported into the VPN instance's IPv4 routing table to guide traffic
                        forwarding. In addition, these routes can be advertised to other devices on the
                        network.

3.7.2 Importing the IPv4 VPN Instance's Routes to the Routing
Tables of a Public Network Instance

Prerequisites
                    Before importing an IPv4 VPN instance's routes to the routing tables of a public
                    network instance, you have completed the following tasks:
                    ●   3.5.1 Configuring an IPv4 VPN Instance on a PE
                    ●   3.5.2 Binding an Interface to an IPv4 VPN Instance

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Import different types of routes from a VPN instance into the public network
                instance's corresponding routing tables.
                         NOTE

