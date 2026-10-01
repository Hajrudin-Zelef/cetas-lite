---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-98
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13647, 13800]
sha256: f9a3d3108c8005e9c179f0d64debeddeeddf55b1568085193ddae179bd2a935d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                                     3 IPv4 L3VPN Configuration


                    Destination: 10.3.1.0/24
                        Protocol: IBGP         Process ID: 0
                      Preference: 255               Cost: 0
                         NextHop: 3.3.3.3        Neighbour: 0.0.0.0
                          State: Active Adv Relied     Age: 00h02m08s
                           Tag: 0            Priority: low
                          Label: 24            QoSInfo: 0x0
                      IndirectID: 0x9D00006A
                     RelayNextHop: 3.3.3.3         Interface: Vlanif100
                        TunnelID: 0x0000000001004c4b44 Flags: RD
                       BkNextHop: 4.4.4.4       BkInterface: Vlanif200
                         BkLabel: 22          SecTunnelID: 0x0
                    BkPETunnelID: 0x0000000001004c4b82 BkPESecTunnelID: 0x0
                    BkIndirectID: 0x9D00006C

                    Run the display mpls lsp include ip-address mask-length verbose command on
                    SPE1. The command output shows information about the index and label of the
                    backup LSP to which the VPNv4 route recurses.
                    <SPE1> display mpls lsp include 10.3.1.1 24 verbose
                    -------------------------------------------------------------------------------
                                 LSP Information: L3VPN LSP
                    -------------------------------------------------------------------------------
                      No                : 0
                      VrfIndex            : ASBR LSP
                      RD Value               : 100:1
                      Fec              : 10.3.1.0/24
                      Nexthop                : 5.5.5.5
                      In-Label            : 24
                      Out-Label              : 23
                      In-Interface          : ------
                      Out-Interface           : ------
                      LspIndex             : 24
                      Type               : Primary
                      OutSegmentIndex             : ------
                      LsrType             : Transit
                      Outgoing TunnelID : 0x4c4b46
                      Label Operation           : SWAP
                      Mpls-Mtu                : ------
                      LspAge               : ------

                     No             : 1
                     VrfIndex         : ASBR LSP
                     RD Value           : 100:1
                     Fec           : 10.3.1.0/24
                     Nexthop            : 6.6.6.6
                     In-Label        : 24
                     Out-Label          : 23
                     In-Interface      : ------
                     Out-Interface       : ------
                     LspIndex          : 24
                     Type           : Backup
                     OutSegmentIndex         : ------
                     LsrType         : Transit
                     Outgoing TunnelID : 0x4c4b48
                     Label Operation       : SWAP
                     Mpls-Mtu            : ------
                     LspAge           : ------


Configuration Scripts
                    ●     UPE1
                          #
                          sysname UPE1
                          #
                          vlan batch 100 200 300
                          #


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                             217
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        ip vpn-instance vpna
                         ipv4-family
                          route-distinguisher 100:1
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.16.3.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.16.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ip address 10.1.1.2 255.255.255.0
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
                         ip address 1.1.1.1 255.255.255.255
                        #
                        bgp 100
                         router-id 1.1.1.1
                         peer 3.3.3.3 as-number 100
                         peer 3.3.3.3 connect-interface LoopBack1
                         peer 4.4.4.4 as-number 100
                         peer 4.4.4.4 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.3 enable
                          peer 4.4.4.4 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.3 enable
                          peer 4.4.4.4 enable
                         #
                         ipv4-family vpn-instance vpna
                          peer 10.1.1.1 as-number 65410
                          import-route direct
                          auto-frr
                          route-select delay 300
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 172.16.2.0 0.0.0.255
                          network 172.16.3.0 0.0.0.255




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         218
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

