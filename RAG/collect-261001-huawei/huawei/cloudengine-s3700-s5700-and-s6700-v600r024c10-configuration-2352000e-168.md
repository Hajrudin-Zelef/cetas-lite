---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-168
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [24285, 24422]
sha256: b76a2cb0e984eec545d7d19fa822a13baebf373a3d17ca52ba3889ace0ceff5a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After completing the configuration, run the display ospf peer command on an
                    ASBR and PE. The command output shows that the OSPF neighbor relationship is
                    in the Full state, indicating that the OSPF neighbor relationship has been
                    established between the ASBR and PE in the same AS.
                    The ASBR and PE in the same AS can learn and successfully ping the address of
                    each other's loopback interface.
         Step 2 Configure basic MPLS capabilities and MPLS LDP on the MPLS backbone networks
                in AS100 and AS200 and establish LDP LSPs.
                    For configuration details, see 4.12.3 Example for Configuring IPv6 L3VPN over
                    MPLS Inter-AS Option A.
         Step 3 Configure basic IPv6 L3VPN functions on PE1 and PE2.
                          NOTE

                        PE1's import and export VPN targets must match PE2's export and import VPN targets,
                        respectively.

                    For detailed configurations, see Configuration Scripts.
         Step 4 Configure inter-AS VPN Option B.
                    # On ASBR1, enable MPLS on VLANIF 200 connected to ASBR2.
                    [ASBR1] interface Vlanif200
                    [ASBR1-Vlanif200] ip address 192.168.1.1 24
                    [ASBR1-Vlanif200] mpls
                    [ASBR1-Vlanif200] quit

                    # Configure ASBR1 to establish an MP-EBGP peer relationship with ASBR2 and not
                    to filter received VPN IPv6 routes based on VPN targets. Then, enable one-label-
                    per-next-hop label distribution on ASBR1.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 192.168.1.2 as-number 200
                    [ASBR1-bgp] ipv6-family vpnv6
                    [ASBR1-bgp-af-vpnv6] peer 192.168.1.2 enable
                    [ASBR1-bgp-af-vpnv6] undo policy vpn-target

                          NOTE

                        The configuration of ASBR2 is similar to that of ASBR1.

                    ----End

Verifying the Configuration
                    After the configuration is complete, CEs can learn routes to the interface of each
                    other, and CE1 and CE2 can successfully ping each other.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       384
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration


                    The following example uses the command output on CE1.
                    <CE1> display ipv6 routing-table

                    Routing Table : _public_
                          Destinations : 7      Routes : 7

                    Destination : ::1                          PrefixLength : 128
                    NextHop      : ::1                         Preference : 0
                    Cost      :0                             Protocol    : Direct
                    RelayNextHop : ::                           TunnelID      : 0x0
                    Interface : InLoopBack0                        Flags       :D

                    Destination : ::FFFF:127.0.0.0                 PrefixLength : 104
                    NextHop      : ::FFFF:127.0.0.1                Preference : 0
                    Cost      :0                             Protocol    : Direct
                    RelayNextHop : ::                           TunnelID      : 0x0
                    Interface : InLoopBack0                       Flags        :D

                    Destination : ::FFFF:127.0.0.1                 PrefixLength : 128
                    NextHop      : ::1                         Preference : 0
                    Cost      :0                             Protocol    : Direct
                    RelayNextHop : ::                           TunnelID      : 0x0
                    Interface : InLoopBack0                        Flags       :D

                    Destination : 2001:db8:1::                     PrefixLength : 64
                    NextHop      : 2001:db8:1::1                    Preference : 0
                    Cost      :0                             Protocol     : Direct
                    RelayNextHop : ::                           TunnelID       : 0x0
                    Interface : Vlanif100                        Flags       :D

                    Destination : 2001:db8:1::1                    PrefixLength : 128
                    NextHop      : ::1                         Preference : 0
                    Cost      :0                             Protocol    : Direct
                    RelayNextHop : ::                           TunnelID      : 0x0
                    Interface : Vlanif100                        Flags      :D

                    Destination : 2001:db8:2::                     PrefixLength : 64
                    NextHop      : 2001:db8:1::2                    Preference : 255
                    Cost      :0                             Protocol     : EBGP
                    RelayNextHop : 2001:db8:1::2                      TunnelID     : 0x0
                    Interface : Vlanif100                        Flags       :D

                    Destination : FE80::                       PrefixLength : 10
                    NextHop       : ::                       Preference : 0
                    Cost       :0                           Protocol   : Direct
                    RelayNextHop : ::                          TunnelID     : 0x0
                    Interface : NULL0                           Flags     :D
                    <CE1> ping ipv6 2001:db8:2::1
                     PING 2001:db8:2::1 : 56 data bytes, press CTRL_C to break
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=1 hop limit=62 time = 125 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=2 hop limit=62 time = 109 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=3 hop limit=62 time = 109 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=4 hop limit=62 time = 109 ms
                       Reply from 2001:db8:2::1
                       bytes=56 Sequence=5 hop limit=62 time = 110 ms


                     --- 2001:db8:2::1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 109/112/125 ms




Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                   385
VPN Configuration
VPN Configuration                                                                                 4 IPv6 L3VPN Configuration


                    Run the display bgp vpnv6 all routing-table command on an ASBR. The
                    command output shows the VPNv6 routes on the ASBR.
                    The following example uses the command output on ASBR1.
                    <ASBR1> display bgp vpnv6 all routing-table
                     BGP Local router ID is 192.168.1.1
                     Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                              h - history, i - internal, s - suppressed, S - Stale
                              Origin : i - IGP, e - EGP, ? - incomplete


                    Total number of routes from all PE: 2
                    Route Distinguisher: 100:1

