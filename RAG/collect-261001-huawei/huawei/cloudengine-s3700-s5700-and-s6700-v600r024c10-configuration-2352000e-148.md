---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-148
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [21081, 21236]
sha256: 3fec7e12e75fd10251276460110644ebc07129a93789b81c2070d4194d54b79a
---

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] peer 1.1.1.1 as-number 100
                    [PE3-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE3-bgp] ipv6-family vpnv6
                    [PE3-bgp-af-vpnv6] peer 1.1.1.1 enable
                    [PE3-bgp-af-vpnv6] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     334
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration


                    Run the display bgp vpnv6 all peer command on each PE. The command output
                    shows that the status of the MP-IBGP peer relationship between the PEs is
                    Established.
                    The following example uses the command output on PE1.
                    <PE1> display bgp vpnv6 all peer
                     BGP local router ID : 10.10.1.1
                     Local AS number : 100
                     Total number of peers : 2            Peers in established state : 2

                     Peer         V        AS MsgRcvd MsgSent OutQ Up/Down        State PrefRcv
                     2.2.2.2      4       100    29   26   0 00:20:08 Established    2
                     3.3.3.3      4       100    18   17   0 00:11:32 Established    1

         Step 7 Configure static BFD for LDP LSPs.
                    # Configure static BFD for LDP LSPs on PE1.
                    [PE1] bfd
                    [PE1-bfd] quit
                    [PE1] bfd for_ldp_lsp bind peer-ip 2.2.2.2 interface Vlanif200
                    [PE1-bfd-session-for_ldp_lsp] discriminator local 10
                    [PE1-bfd-session-for_ldp_lsp] discriminator remote 20
                    [PE1-bfd-session-for_ldp_lsp] process-pst
                    [PE1-bfd-session-for_ldp_lsp] quit

                    # Configure static BFD for LDP LSPs on PE2.
                    [PE2] bfd
                    [PE2-bfd] quit
                    [PE2] bfd for_ldp_lsp bind peer-ip 1.1.1.1 interface Vlanif100
                    [PE2-bfd-session-for_ldp_lsp] discriminator local 20
                    [PE2-bfd-session-for_ldp_lsp] discriminator remote 10
                    [PE2-bfd-session-for_ldp_lsp] quit

                    # After completing the configuration, run the display bfd session all verbose
                    command on PE1 and PE2. The command output shows that the State field
                    displays Up, and the BFD Bind Type field displays LDP_LSP.
         Step 8 Enable VPN BGP auto FRR.
                    [PE1] bgp 100
                    [PE1-bgp] ipv6-family vpn-instance vpn1
                    [PE1-bgp6-vpn1] auto-frr
                    [PE1-bgp6-vpn1] route-select delay 300
                    [PE1-bgp6-vpn1] quit
                    [PE1-bgp] quit

                    ----End

Verifying the Configuration
                    After completing the configuration, run the display ipv6 routing-table vpn-
                    instance verbose command on PE1. The command output shows the backup next
                    hop, backup label, and backup tunnel ID of the IPv6 VPN route.
                    <PE1> display ipv6 routing-table vpn-instance vpn1 2001:DB8:2::1 128 verbose
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 1

                    Destination : 2001:DB8:2::1                    PrefixLength : 128
                    NextHop       : 2.2.2.2                     Preference : 255
                    Neighbour : ::                            ProcessID : 0
                    Label      : 4099                         Protocol    : IBGP


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       335
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration

                    State     : Active Adv Relied          Cost         :0
                    Entry ID    :0                     EntryFlags : 0x00000000
                    Reference Cnt: 0                     Tag         :0
                    Priority  : low                    Age        : 450sec
                    IndirectID : 0x5A00006E
                    RelayNextHop : ::                    TunnelID       : 0x0000000001004c4b42
                    Interface : LDP LSP                   Flags       : RD
                    BkNextHop : 3.3.3.3                    BkInterface : LDP LSP
                    BkLabel     : 4098                   BkTunnelID : 0x0
                    BkPETunnelID: 0x0000000001004c4b43               BkIndirectID : 0x5A000070


Configuration Scripts
                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 200 300
                         #
                         ip vpn-instance vpn1
                          ipv6-family
                           route-distinguisher 100:1
                           vpn-target 111:1 export-extcommunity
                           vpn-target 111:1 import-extcommunity
                         #
                         bfd
                         #
                         mpls lsr-id 1.1.1.1
                         #
                         mpls
                          mpls bfd enable
                         #
                         mpls ldp
                         #
                         interface Vlanif200
                          ip address 10.10.1.1 255.255.255.252
                          mpls
                          mpls ldp
                         #
                         interface Vlanif300
                          ip address 10.20.1.1 255.255.255.252
                          mpls
                          mpls ldp
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
                         interface LoopBack2
                          ip binding vpn-instance vpn1
                          ipv6 enable
                          ipv6 address 2001:DB8:8::1/64
                         #
                         bgp 100
                          peer 2.2.2.2 as-number 100
                          peer 2.2.2.2 connect-interface LoopBack1
                          peer 3.3.3.3 as-number 100
                          peer 3.3.3.3 connect-interface LoopBack1
                          #
                          ipv4-family unicast
                           peer 2.2.2.2 enable
                           peer 3.3.3.3 enable


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 336
VPN Configuration
VPN Configuration                                                                  4 IPv6 L3VPN Configuration

