---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-68
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [9010, 9146]
sha256: 853b856b90bd33fbd826800d6fb012d82fd591899503dabecee567dae56443b0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 4 Configure inter-AS VPN in VRF-to-VRF mode.
                    # Create a VPN instance on ASBR1 and bind the interface that connects ASBR1 to
                    ASBR2 (viewed as a CE by ASBR1) to the VPN instance.
                    [ASBR1] ip vpn-instance vpn1
                    [ASBR1-vpn-instance-vpn1] ipv4-family
                    [ASBR1-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:2
                    [ASBR1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 both
                    [ASBR1-vpn-instance-vpn1-af-ipv4] quit
                    [ASBR1-vpn-instance-vpn1] quit
                    [ASBR1] interface 10GE1/0/2
                    [ASBR1-10GE1/0/2] port link-type trunk
                    [ASBR1-10GE1/0/2] port trunk allow-pass vlan 200
                    [ASBR1-10GE1/0/2] quit
                    [ASBR1] interface Vlanif 200
                    [ASBR1-Vlanif200] ip binding vpn-instance vpn1
                    [ASBR1-Vlanif200] ip address 12.12.12.1 24
                    [ASBR1-Vlanif200] quit

                    # Create a VPN instance on ASBR2 and bind the interface that connects ASBR2 to
                    ASBR1 (viewed as a CE by ASBR2) to the VPN instance.
                    [ASBR2] ip vpn-instance vpn1
                    [ASBR2-vpn-instance-vpn1] ipv4-family
                    [ASBR2-vpn-instance-vpn1-af-ipv4] route-distinguisher 200:2


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                  143
VPN Configuration
VPN Configuration                                                                                3 IPv4 L3VPN Configuration

                    [ASBR2-vpn-instance-vpn1-af-ipv4] vpn-target 2:2 both
                    [ASBR2-vpn-instance-vpn1-af-ipv4] quit
                    [ASBR2-vpn-instance-vpn1] quit
                    [ASBR2] interface 10GE1/0/2
                    [ASBR2-10GE1/0/2] port link-type trunk
                    [ASBR2-10GE1/0/2] port trunk allow-pass vlan 200
                    [ASBR2-10GE1/0/2] quit
                    [ASBR2] interface Vlanif 200
                    [ASBR2-Vlanif200] ip binding vpn-instance vpn1
                    [ASBR2-Vlanif200] ip address 12.12.12.2 24
                    [ASBR2-Vlanif200] quit

                    # Configure ASBR1 to set up an EBGP peer relationship with ASBR2.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] ipv4-family vpn-instance vpn1
                    [ASBR1-bgp-vpn1] peer 12.12.12.2 as-number 200

                    # Configure ASBR2 to set up an EBGP peer relationship with ASBR1.
                    [ASBR2] bgp 200
                    [ASBR2-bgp] ipv4-family vpn-instance vpn1
                    [ASBR2-bgp-vpn1] peer 12.12.12.1 as-number 100

                    After completing the configuration, run the display bgp vpnv4 vpn-instance vpn-
                    instance-name peer command on an ASBR. The command output shows that a
                    BGP peer relationship has been established between ASBRs and is in the
                    Established state.

                    ----End

Verifying the Configuration
                    After the configuration is complete, the CEs can learn routes to interfaces of each
                    other and can ping each other successfully.
                    The following example uses the command output on CE1.
                    <CE1> display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 9        Routes : 9

                    Destination/Mask     Proto Pre Cost        Flags NextHop         Interface

                         10.1.1.0/24 Direct 0 0          D 10.1.1.1      Vlanif100
                         10.1.1.1/32 Direct 0 0          D 127.0.0.1     Vlanif100
                       10.1.1.255/32 Direct 0 0           D 127.0.0.1      Vlanif100
                      11.11.11.11/32 Direct 0 0            D 127.0.0.1     LoopBack1
                      22.22.22.22/32 EBGP 255 0              RD 10.1.1.2     Vlanif100
                        127.0.0.0/8 Direct 0 0           D 127.0.0.1     InLoopBack0
                        127.0.0.1/32 Direct 0 0           D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                    <CE1> ping -a 11.11.11.11 22.22.22.22
                     PING 22.22.22.22: 56 data bytes, press CTRL_C to break
                      Reply from 22.22.22.22: bytes=56 Sequence=1 ttl=251 time=46 ms
                      Reply from 22.22.22.22: bytes=56 Sequence=2 ttl=251 time=4 ms
                      Reply from 22.22.22.22: bytes=56 Sequence=3 ttl=251 time=4 ms
                      Reply from 22.22.22.22: bytes=56 Sequence=4 ttl=251 time=4 ms
                      Reply from 22.22.22.22: bytes=56 Sequence=5 ttl=251 time=4 ms

                     --- 22.22.22.22 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                         144
VPN Configuration
VPN Configuration                                                                                     3 IPv4 L3VPN Configuration

                         0.00% packet loss
                         round-trip min/avg/max = 4/12/46 ms

                    Run the display ip routing-table vpn-instance command on an ASBR. The
                    command output shows its VPN routing table.
                    <ASBR1> display ip routing-table vpn-instance vpn1
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                           Destinations : 7        Routes : 7

                    Destination/Mask     Proto Pre Cost           Flags NextHop          Interface

                       11.11.11.11/32 IBGP 255 0                 RD 1.1.1.9       Vlanif100
                        12.12.12.0/24 Direct 0 0                D 12.12.12.1     Vlanif200
                        12.12.12.1/32 Direct 0 0                D 127.0.0.1     Vlanif200
                      12.12.12.255/32 Direct 0 0                 D 127.0.0.1     Vlanif200
                       22.22.22.22/32 EBGP 255 0                  RD 12.12.12.2     Vlanif200
                         127.0.0.0/8 Direct 0 0                D 127.0.0.1     InLoopBack0
                    255.255.255.255/32 Direct 0 0                 D 127.0.0.1      InLoopBack0

                    Run the display bgp vpnv4 all routing-table command on an ASBR. The
                    command output shows the VPNv4 routes on the ASBR.
                    <ASBR1> display bgp vpnv4 all routing-table
                    BGP Local router ID is 10.10.1.1
                     Status codes: * - valid, > - best, d - damped, x - best external, a - add path,
                               h - history, i - internal, s - suppressed, S - Stale
                               Origin : i - IGP, e - EGP, ? - incomplete
                     RPKI validation codes: V - valid, I - invalid, N - not-found


                    Total number of routes from all PE: 2
                    Route Distinguisher: 100:1


                           Network           NextHop       MED        LocPrf       PrefVal Path/Ogn

