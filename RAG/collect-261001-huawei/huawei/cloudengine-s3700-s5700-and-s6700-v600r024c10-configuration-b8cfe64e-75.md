---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-75
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [10675, 10834]
sha256: 5e27bf90eab27e8f5957c3ca7bfded73367acbc6d022ba510d8c3521d5c5aa88
---

                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 3.3.3.9 as-number 100
                    [PE1-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [PE1-bgp] ipv4-family vpnv4
                    [PE1-bgp-af-vpnv4] peer 3.3.3.9 enable
                    [PE1-bgp-af-vpnv4] quit
                    [PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.9 as-number 100
                    [PE2-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [PE2-bgp] ipv4-family vpnv4
                    [PE2-bgp-af-vpnv4] peer 1.1.1.9 enable


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                   194
QoS Configuration
QoS Configuration                                                                                    12 MPLS QoS Configuration

                    [PE2-bgp-af-vpnv4] quit
                    [PE2-bgp] quit

                    After the configuration is complete, run the display bgp peer command on PEs.
                    The command output shows that a BGP peer relationship has been established
                    between PEs and is in Established state.
                    [PE1] display bgp peer

                    BGP local router ID : 1.1.1.9
                    Local AS number : 100
                    Total number of peers : 1                Peers in established state : 1

                     Peer         V    AS MsgRcvd MsgSent OutQ Up/Down                  State            PrefRcv

                     3.3.3.9      4 100       12     6       0 00:02:21       Established       0

         Step 5 Establish EBGP peer relationships between PEs and CEs and import VPN routes.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] bgp 65410
                    [CE1-bgp] peer 10.1.1.2 as-number 100
                    [CE1-bgp] import-route direct

                    The configurations of CE2, CE3, and CE4 are similar to the configuration of CE1,
                    and are not mentioned here.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpna
                    [PE1-bgp-vpna] peer 10.1.1.1 as-number 65410
                    [PE1-bgp-vpna] import-route direct
                    [PE1-bgp-vpna] quit
                    [PE1-bgp] ipv4-family vpn-instance vpnb
                    [PE1-bgp-vpnb] peer 10.2.1.1 as-number 65420
                    [PE1-bgp-vpnb] import-route direct
                    [PE1-bgp-vpnb] quit
                    [PE1-bgp] quit

                    The configuration of PE2 is similar to that of PE1, and is not mentioned here.
                    After the configuration is complete, run the display bgp vpnv4 vpn-instance peer
                    command on PEs. The command output shows that BGP peer relationships have
                    been established between PEs and CEs and are in Established state.
                    The following example uses the command output on PE1 to show that a peer
                    relationship has been established between PE1 and CE1.
                    [PE1] display bgp vpnv4 vpn-instance vpna peer

                    BGP local router ID : 1.1.1.9
                    Local AS number : 100
                    Total number of peers : 1                Peers in established state : 1

                     Peer         V    AS MsgRcvd MsgSent OutQ Up/Down                  State       PrefRcv

                     10.1.1.1     4 65410       11       9     0 00:07:25     Established       1

         Step 6 Configure a DiffServ mode.
                    # Configure PE1.
                    [PE1] mpls-qos ingress use vpn-label-exp
                    [PE1] ip vpn-instance vpna


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                           195
QoS Configuration
QoS Configuration                                                                   12 MPLS QoS Configuration

                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] diffserv-mode pipe mpls-exp 4
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv4-family
                    [PE1-vpn-instance-vpnb-af-ipv4] diffserv-mode pipe mpls-exp 3
                    [PE1-vpn-instance-vpnb-af-ipv4] quit
                    [PE1-vpn-instance-vpnb] quit

                    # Configure PE2.
                    [PE2] mpls-qos ingress use vpn-label-exp
                    [PE2] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE2-vpn-instance-vpna-af-ipv4] diffserv-mode pipe mpls-exp 4
                    [PE2-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE2] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv4-family
                    [PE2-vpn-instance-vpnb-af-ipv4] diffserv-mode pipe mpls-exp 3
                    [PE2-vpn-instance-vpnb-af-ipv4] quit
                    [PE1-vpn-instance-vpnb] quit

                    ----End

Configuration Scripts
                    ●    PE1
                         #
                         sysname PE1
                         #
                         mpls-qos ingress use vpn-label-exp
                         #
                         ip vpn-instance vpna
                          ipv4-family
                           route-distinguisher 100:1
                           vpn-target 111:1 export-extcommunity
                           vpn-target 111:1 import-extcommunity
                           diffserv-mode pipe mpls-exp 4
                         #
                         ip vpn-instance vpnb
                          ipv4-family
                           route-distinguisher 100:2
                           vpn-target 222:2 export-extcommunity
                           vpn-target 222:2 import-extcommunity
                           diffserv-mode pipe mpls-exp 3
                         #
                         mpls lsr-id 1.1.1.9
                         mpls
                         #
                         mpls ldp
                         #
                         interface 10ge1/0/1
                          undo portswitch
                          ip binding vpn-instance vpna
                          ip address 10.1.1.2 255.255.255.0
                         #
                         interface 10ge1/0/2
                          undo portswitch
                          ip binding vpn-instance vpnb
                          ip address 10.2.1.2 255.255.255.0
                         #
                         interface 10ge1/0/3
                          undo portswitch
                          ip address 172.16.1.1 255.255.255.0
                          mpls
                          mpls ldp


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             196
QoS Configuration
QoS Configuration                                                             12 MPLS QoS Configuration

