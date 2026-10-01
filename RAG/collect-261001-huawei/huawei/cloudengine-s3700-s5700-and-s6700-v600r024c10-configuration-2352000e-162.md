---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-162
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [23297, 23430]
sha256: 35271ee01daf2a93ea1702a35c6ec9f058a005a78155e381e99b2c912b94bfd5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After completing the configuration, run the display bgp vpnv6 vpn-instance peer
                    command on the PEs. The command output shows that BGP peer relationships
                    have been established between the PEs and CEs and are in Established state. Run
                    the display bgp vpnv6 all peer command on a PE or ASBR. The command output
                    shows that BGP peer relationships have been established between PEs and CEs,
                    and between PEs and ASBRs.
                    The following example uses the command output on PE1.
                    <PE1> display bgp vpnv6 vpn-instance vpn1 peer
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 1        Peers in established state : 1
                      Peer        V AS MsgRcvd MsgSent OutQ Up/Down              State PrefRcv

                      2001:db8:1::1 4 65001        14 12    0 00:08:36 Established        1
                    <PE1> display bgp vpnv6 all peer
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 2         Peers in established state : 2

                     Peer          V    AS MsgRcvd MsgSent OutQ Up/Down            State PrefRcv

                     2.2.2.9       4 100       13   12    0 00:09:10 Established      0

                      Peer of vpn instance :

                     VPN-Instance vpn1, router ID 1.1.1.9:
                     2001:db8:1::1 4 65001      17      14   0 00:11:09 Established       1

         Step 4 Configure inter-AS VPN in VRF-to-VRF mode.
                    # Create an IPv6-address-family-enabled VPN instance on ASBR1 and bind the
                    interface that connects ASBR1 to ASBR2 (viewed as a CE by ASBR1) to the VPN
                    instance.

Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                    368
VPN Configuration
VPN Configuration                                                                 4 IPv6 L3VPN Configuration

                    [ASBR1] ip vpn-instance vpn1
                    [ASBR1-vpn-instance-vpn1] ipv6-family
                    [ASBR1-vpn-instance-vpn1-af-ipv6] route-distinguisher 100:2
                    [ASBR1-vpn-instance-vpn1-af-ipv6] vpn-target 1:1 both
                    [ASBR1-vpn-instance-vpn1-af-ipv6] quit
                    [ASBR1-vpn-instance-vpn1] quit
                    [ASBR1] interface Vlanif200
                    [ASBR1-Vlanif200] ip binding vpn-instance vpn1
                    [ASBR1-Vlanif200] ipv6 enable
                    [ASBR1-Vlanif200] ipv6 address 2001:db8:3::1 64
                    [ASBR1-Vlanif200] quit

                    # Create a VPN instance on ASBR2 and bind the interface that connects ASBR2 to
                    ASBR1 (viewed as a CE by ASBR2) to the VPN instance.
                    [ASBR2] ip vpn-instance vpn1
                    [ASBR2-vpn-instance-vpn1] ipv6-family
                    [ASBR2-vpn-instance-vpn1-af-ipv6] route-distinguisher 200:2
                    [ASBR2-vpn-instance-vpn1-af-ipv6] vpn-target 2:2 both
                    [ASBR2-vpn-instance-vpn1-af-ipv6] quit
                    [ASBR2-vpn-instance-vpn1] quit
                    [ASBR2] interface Vlanif200
                    [ASBR2-Vlanif200] ip binding vpn-instance vpn1
                    [ASBR2-Vlanif200] ipv6 enable
                    [ASBR2-Vlanif200] ipv6 address 2001:db8:3::2 64
                    [ASBR2-Vlanif200] quit

                    # Configure ASBR1 to set up an EBGP peer relationship with ASBR2.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] ipv6-family vpn-instance vpn1
                    [ASBR1-bgp6-vpn1] peer 2001:db8:3::2 as-number 200
                    [ASBR1-bgp6-vpn1] import-route direct

                    # Configure ASBR2 to set up an EBGP peer relationship with ASBR1.
                    [ASBR2] bgp 200
                    [ASBR2-bgp] ipv6-family vpn-instance vpn1
                    [ASBR2-bgp6-vpn1] peer 2001:db8:3::1 as-number 100
                    [ASBR2-bgp6-vpn1] import-route direct

                    After completing the configuration, run the display bgp vpnv6 vpn-instance peer
                    command. The command output shows that the BGP peer relationship between
                    the ASBRs is in the Established state.

                    ----End

Verifying the Configuration
                    After the configuration is complete, CE1 and CE2 can learn routes to interfaces on
                    each other and ping each other successfully. The following example uses the
                    command output on CE1.
                    The following example uses the command output on CE1.
                    <CE1> display ipv6 routing-table
                    Routing Table : _public_
                          Destinations : 8   Routes : 8

                    Destination : ::1                  PrefixLength : 128
                    NextHop      : ::1                 Preference : 0
                    Cost      :0                     Protocol    : Direct
                    RelayNextHop : ::                   TunnelID      : 0x0
                    Interface : InLoopBack0                Flags       :D

                    Destination : ::FFFF:127.0.0.0         PrefixLength : 104
                    NextHop     : ::FFFF:127.0.0.1         Preference : 0


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                            369
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    Cost      :0                        Protocol   : Direct
                    RelayNextHop : ::                      TunnelID     : 0x0
                    Interface : InLoopBack0                  Flags       :D

                    Destination : ::FFFF:127.0.0.1            PrefixLength : 128
                    NextHop      : ::1                    Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : InLoopBack0                   Flags       :D

                    Destination : 2001:db8:1::                PrefixLength : 64
                    NextHop      : 2001:db8:1::1               Preference : 0
                    Cost      :0                        Protocol     : Direct
                    RelayNextHop : ::                      TunnelID       : 0x0
                    Interface : Vlanif100                   Flags       :D

                    Destination : 2001:db8:1::1               PrefixLength : 128
                    NextHop      : ::1                    Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif100                   Flags      :D

                    Destination : 2001:db8:2::                PrefixLength : 64
                    NextHop      : 2001:db8:1::2               Preference : 255
                    Cost      :0                        Protocol     : EBGP
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif100                   Flags       :D

