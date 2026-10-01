---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-152
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [21690, 21820]
sha256: 3a683d4b817dbf1e7832b248eb9b569b4969fce0c0a9751249791563780338eb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Destination : 2001:DB8:0::                PrefixLength : 64
                    NextHop      : 2001:DB8:4::1               Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif100           Flags      :D

                    Destination : 2001:DB8:4::1               PrefixLength : 128
                    NextHop      : ::1                    Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif200           Flags      :D

                    Destination : 2001:DB8:1::                PrefixLength : 64
                    NextHop      : 2001:DB8:1::1               Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif200           Flags      :D



Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                          343
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    Destination : 2001:DB8:1::1               PrefixLength : 128
                    NextHop      : ::1                    Preference : 0
                    Cost      :0                        Protocol    : Direct
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif100           Flags      :D

                    Destination : 2001:DB8:2::                PrefixLength : 64
                    NextHop      : 2001:DB8:4::2               Preference : 255
                    Cost      : 100                      Protocol    : EBGP
                    RelayNextHop : ::                      TunnelID      : 0x0
                    Interface : Vlanif100           Flags      :D

                    Destination : 2001:DB8:3::                PrefixLength : 64
                    NextHop      : 2001:DB8:1::2               Preference : 255
                    Cost      :0                        Protocol    : EBGP
                    RelayNextHop : ::                      TunnelID     : 0x0
                    Interface : Vlanif200           Flags      :D

                    Destination : 2001:DB8:4::1                PrefixLength : 128
                    NextHop      : 2001:DB8:4::2               Preference : 255
                    Cost      : 100                      Protocol    : EBGP
                    RelayNextHop : ::                       TunnelID     : 0x0
                    Interface : Vlanif100            Flags      :D

                    Destination : FE80::                   PrefixLength : 10
                    NextHop      : ::                    Preference : 0
                    Cost      :0                        Protocol   : Direct
                    RelayNextHop : ::                      TunnelID     : 0x0
                    Interface : NULL0                       Flags     :D

         Step 6 Enable VPN IPv6 auto FRR on the PE.
                    # Configure the PE.
                    [PE] bgp 100
                    [PE-bgp] ipv6-family vpn-instance vpna
                    [PE-bgp6-vpna] auto-frr
                    [PE-bgp6-vpna] route-select delay 300
                    [PE-bgp6-vpna] quit
                    [PE-bgp] quit

                          NOTE

                    The auto-frr command configured in the BGP VPN instance IPv6 address family view applies
                    only to the network where BGP runs between the PE and CEs.

                    ----End

Verifying the Configuration
                    After completing the configuration, run the display ipv6 routing-table vpn-
                    instance command on the PE. The command output shows that the next hop to
                    2001:DB8:4::1/128 is 2001:DB8:4::2 and that there is a backup next hop and a
                    backup outbound interface.
                    <PE> display ipv6 routing-table vpn-instance vpna 2001:DB8:4::1 verbose
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                    Summary Count : 1

                    Destination : 2001:DB8:4::1                   PrefixLength : 128
                    NextHop       : 2001:DB8:4::2            Preference : 255
                    Neighbour : 2001:DB8:4::2                 ProcessID : 0
                    Label      : NULL                     Protocol    : EBGP
                    State     : Active Adv                 Cost      : 100
                    Entry ID    : 27                     EntryFlags : 0x80004100


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         344
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    Reference Cnt: 2                 Tag        :0
                    IndirectID : 0x6                Age       : 3sec
                    RelayNextHop : ::                TunnelID     : 0x0
                    Interface : Vlanif100     Flags      :D
                    BkNextHop : 2001:DB8:1::2              BkInterface : Vlanif200
                    BkLabel     : NULL                BkTunnelID : 0x0
                    BkPETunnelID : 0x0                  BkIndirectID : 0x5

                    Disable IPv6 on VLANIF 200 of CE1 so that IPv6 routes cannot be transmitted over
                    Link_A.
                    [CE1] interface Vlanif200
                    [CE1-Vlanif200] undo ipv6 enable
                    [CE1] quit

                    Run the display ipv6 routing-table vpn-instance command on the PE again. The
                    command output shows that the next hop to 2001:DB8:4::1/128 is 2001:DB8:1::2
                    and that there is no backup next hop or backup outbound interface.
                    <PE> display ipv6 routing-table vpn-instance vpna 2001:DB8:4::1 verbose
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                    Summary Count : 1

                    Destination : 2001:DB8:4::1                   PrefixLength : 128
                    NextHop        : 2001:DB8:1::2             Preference : 255
                    Neighbour : 2001:DB8:1::2                 ProcessID : 0
                    Label      : NULL                     Protocol      : EBGP
                    State     : Active Adv                 Cost        : 500
                    Entry ID    : 27                     EntryFlags : 0x80004100
                    Reference Cnt: 2                       Tag         :0
                    IndirectID : 0x6                     Age         : 3sec
                    RelayNextHop : ::                      TunnelID       : 0x0
                    Interface : Vlanif200          Flags      :D

                    VPN IPv6 auto FRR has taken effect.

