---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-76
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [10835, 11033]
sha256: 8c193309c5d7c4d3c9243dd3bce692d2a9252fd3d177fae6d36c92208ee6a6e0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        bgp 100
                         peer 3.3.3.9 as-number 100
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.9 enable
                         #
                         ipv4-family vpn-instance vpna
                          import-route direct
                          peer 10.1.1.1 as-number 65410
                         #
                         ipv4-family vpn-instance vpnb
                          import-route direct
                          peer 10.2.1.1 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 172.16.1.0 0.0.0.255
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls ldp
                        #
                        interface 10ge1/0/1
                         undo portswitch
                         ip address 172.16.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10ge1/0/2
                         undo portswitch
                         ip address 172.17.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 172.16.1.0 0.0.0.255
                          network 172.17.1.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        mpls-qos ingress use vpn-label-exp
                        #
                        ip vpn-instance vpna


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        197
QoS Configuration
QoS Configuration                                                             12 MPLS QoS Configuration

                         ipv4-family
                          route-distinguisher 200:1
                          vpn-target 111:1 export-extcommunity
                          vpn-target 111:1 import-extcommunity
                          diffserv-mode pipe mpls-exp 4
                        #
                        ip vpn-instance vpnb
                         ipv4-family
                          route-distinguisher 200:2
                          vpn-target 222:2 export-extcommunity
                          vpn-target 222:2 import-extcommunity
                          diffserv-mode pipe mpls-exp 3
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls ldp
                        #
                        interface 10ge1/0/1
                         undo portswitch
                         ip binding vpn-instance vpna
                         ip address 10.3.1.2 255.255.255.0
                        #
                        interface 10ge1/0/2
                         undo portswitch
                         ip binding vpn-instance vpnb
                         ip address 10.4.1.2 255.255.255.0
                        #
                        interface 10ge1/0/3
                         undo portswitch
                         ip address 172.17.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
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
                          network 3.3.3.9 0.0.0.0
                          network 172.17.1.0 0.0.0.255
                        #
                        return
                    ●   CE1 at the headquarters egress of enterprise A
                        #
                        sysname CE1
                        #
                        interface 10ge1/0/1
                         undo portswitch


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        198
QoS Configuration
QoS Configuration                                                             12 MPLS QoS Configuration

                         ip address 10.1.1.1 255.255.255.0
                        #
                        bgp 65410
                         peer 10.1.1.2 as-number 100
                         #
                         ipv4-family unicast
                          undo synchronization
                          import-route direct
                          peer 10.1.1.2 enable
                        #
                        return

                    ●   CE2 at the headquarters egress of enterprise B
                        #
                        sysname CE2
                        #
                        interface 10ge1/0/1
                         undo portswitch
                         ip address 10.2.1.1 255.255.255.0
                        #
                        bgp 65420
                         peer 10.2.1.2 as-number 100
                         #
                         ipv4-family unicast
                          undo synchronization
                          import-route direct
                          peer 10.2.1.2 enable
                        #
                        return

                    ●   CE3 at the branch egress of enterprise A
                        #
                        sysname CE3
                        #
                        interface 10ge1/0/1
                         undo portswitch
                         ip address 10.3.1.1 255.255.255.0
                        #
                        bgp 65430
                         peer 10.3.1.2 as-number 100
                         #
                         ipv4-family unicast
                          undo synchronization
                          import-route direct
                          peer 10.3.1.2 enable
                        #
                        return

