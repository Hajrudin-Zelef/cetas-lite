---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-151
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [21548, 21689]
sha256: 20bfab01570f99a97fc49d0f1ac177a682c51d2a4ed0673bd916a18b3d2a229d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Destination : 2001:DB8:4::2                 PrefixLength : 128
                    NextHop      : ::1                      Preference : 0
                    Cost      :0                          Protocol    : Direct
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Vlanif100             Flags      :D

                    Destination : 2001:DB8:2::                  PrefixLength : 64
                    NextHop      : 2001:DB8:2::1                 Preference : 0
                    Cost      :0                          Protocol    : Direct
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Vlanif200             Flags      :D

                    Destination : 2001:DB8:2::1                 PrefixLength : 128
                    NextHop      : ::1                      Preference : 0
                    Cost      :0                          Protocol    : Direct
                    RelayNextHop : ::                        TunnelID      : 0x0
                    Interface : Vlanif200             Flags      :D

                    Destination : 2001:DB8:3::                 PrefixLength : 64
                    NextHop     : FE80::5451:0:FAC1:1             Preference : 10


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                              341
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration

                    Cost      : 3124                   Protocol   : OSPFv3
                    RelayNextHop : ::                    TunnelID    : 0x0
                    Interface : Vlanif200        Flags      :D

                    Destination : 2001:DB8:4::1            PrefixLength : 128
                    NextHop      : FE80::5451:0:FAC1:1        Preference : 10
                    Cost      : 1562                   Protocol    : OSPFv3
                    RelayNextHop : ::                    TunnelID     : 0x0
                    Interface : Vlanif200          Flags     :D

                    Destination : FE80::                PrefixLength : 10
                    NextHop      : ::                 Preference : 0
                    Cost      :0                     Protocol   : Direct
                    RelayNextHop : ::                   TunnelID     : 0x0
                    Interface : NULL0                    Flags     :D

         Step 3 Configure an IPv6-address-family-enabled VPN instance on the PE and bind the
                interfaces that connect the PE to CEs to the VPN instance.

                    # Configure a VPN instance (vpna) on the PE and bind VLANIF 100 and VLANIF
                    200 to vpna.
                    <PE> system-view
                    [PE] ip vpn-instance vpna
                    [PE-vpn-instance-vpna] ipv6-family
                    [PE-vpn-instance-vpna-af-ipv6] route-distinguisher 100:1
                    [PE-vpn-instance-vpna-af-ipv6] vpn-target 100:100
                    [PE-vpn-instance-vpna-af-ipv6] quit
                    [PE-vpn-instance-vpna] quit
                    [PE] interface Vlanif100
                    [PE-Vlanif100] ip binding vpn-instance vpna
                    [PE-Vlanif100] ipv6 enable
                    [PE-Vlanif100] ipv6 address 2001:DB8:4::1 64
                    [PE-Vlanif100] quit
                    [PE] interface Vlanif200
                    [PE-Vlanif200] ip binding vpn-instance vpna
                    [PE-Vlanif200] ipv6 enable
                    [PE-Vlanif200] ipv6 address 2001:DB8:4::1 64
                    [PE] quit

         Step 4 Establish EBGP peer relationships between the PE and CEs.

                    # Configure the PE.
                    [PE] bgp 100
                    [PE-bgp] ipv6-family vpn-instance vpna
                    [PE-bgp6-vpna] peer 2001:DB8:4::2 as-number 65410
                    [PE-bgp6-vpna] peer 2001:DB8:1::2 as-number 65410
                    [PE-bgp-vpna] quit
                    [PE-bgp] quit

                    # Configure CE1.
                    [CE1] bgp 65410
                    [CE1-bgp] peer 2001:DB8:4::1 as-number 100
                    [CE1-bgp] ipv6-family unicast
                    [CE1-bgp-af-ipv6] peer 2001:DB8:4::1 enable
                    [CE1-bgp-af-ipv6] quit
                    [CE1-bgp] quit

                    The configuration of CE2 is similar to the configuration of CE1. For detailed
                    configurations, see Configuration Scripts.

                    After completing the configuration, run the display bgp vpnv6 vpn-instance
                    vpna peer command on the PE. The command output shows that the status of
                    the EBGP peer relationships between the PE and CEs is Established, indicating
                    that EBGP peer relationships have been established between the PE and CEs.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          342
VPN Configuration
VPN Configuration                                                                                  4 IPv6 L3VPN Configuration

                    <PE> display bgp vpnv6 vpn-instance vpna peer

                    BGP local router ID : 1.1.1.1
                    Local AS number : 100
                    Total number of peers : 2               Peers in established state : 2

                     Peer         V       AS MsgRcvd MsgSent OutQ Up/Down                    State PrefRcv

                     2001:DB8:4::2 4       65410       35      37    0 00:24:31 Established 3
                     2001:DB8:1::2 4       65410       41      43    0 00:24:03 Established 3

         Step 5 Configure route exchange between OSPFv3 and BGP on the CEs.
                    Configure OSPFv3 routes to be imported into BGP on the CEs. To enable the PE to
                    select the route along Link_A as the optimal route, ensure that the MED
                    configured for the OSPFv3 routes imported into BGP on CE1 is smaller than that
                    configured on CE2.
                    # Configure CE1.
                    [CE1] bgp 65410
                    [CE1-bgp] ipv6-family unicast
                    [CE1-bgp-af-ipv6] import-route ospfv3 1 med 100
                    [CE1-bgp-af-ipv6] quit
                    [CE1-bgp] quit
                    [CE1] ospfv3 1
                    [CE1-ospfv3-1] import-route bgp
                    [CE1-ospfv3-1] quit

                    # Configure CE2.
                    [CE2] bgp 65410
                    [CE2-bgp] ipv6-family unicast
                    [CE2-bgp-af-ipv6] import-route ospfv3 1 med 500
                    [CE2-bgp-af-ipv6] quit
                    [CE2-bgp] quit
                    [CE2] ospfv3 1
                    [CE2-ospfv3-1] import-route bgp
                    [CE2-ospfv3-1] quit

                    After completing the configuration, run the display ipv6 routing-table vpn-
                    instance command on the PE. The command output shows routes to the loopback
                    interface on DeviceA.
                    <PE> display ipv6 routing-table vpn-instance vpna
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 8        Routes : 8

