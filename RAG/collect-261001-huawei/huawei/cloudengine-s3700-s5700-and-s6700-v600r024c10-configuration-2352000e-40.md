---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-40
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [4701, 4826]
sha256: 2e3b77fc38da63ba9fc41d1d24336f3ff53107204a6597c4c8585f3dd016068b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 5 Establish MP-IBGP peer relationships between PEs.
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
                    [PE2-bgp-af-vpnv4] quit
                    [PE2-bgp] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       74
VPN Configuration
VPN Configuration                                                                                   3 IPv4 L3VPN Configuration


                    After the configuration is complete, run the display bgp peer or display bgp
                    vpnv4 all peer command on PEs. The command output shows that BGP peer
                    relationships have been established between the PEs and are in the Established
                    state.
                    [PE1] display bgp peer
                     Status codes: * - Dynamic
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 1
                     Peers in established state : 1
                     Total number of dynamic peers : 0

                      Peer         V AS MsgRcvd MsgSent OutQ Up/Down                    State         PrefRcv
                      3.3.3.9      4 100          2   6 0 00:00:12 Established                  0
                    [PE1] display bgp vpnv4 all peer
                     Status codes: * - Dynamic
                     BGP local router ID : 1.1.1.9
                     Local AS number : 100
                     Total number of peers : 3
                     Peers in established state : 3
                     Total number of dynamic peers : 0

                     Peer         V AS MsgRcvd MsgSent OutQ Up/Down State                       PrefRcv
                     3.3.3.9      4 100 12      18        0   00:09:38 Established 0
                     Peer of vpn instance:
                     VPN-Instance vpna, router ID 1.1.1.9:
                     10.1.1.1     4 65410 25      25        0  00:17:57 Established 1
                     VPN-Instance vpnb, router ID 1.1.1.9:
                     10.2.1.1     4 65420 21      22        0  00:17:10 Established 1

                    ----End

Verifying the Configuration
                    Run the display ip routing-table vpn-instance command on PEs to check the
                    routes to CEs' loopback interfaces.
                    The following example uses the command output on PE1.
                    [PE1] display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: vpna
                           Destinations : 6        Routes : 6

                    Destination/Mask Proto Pre Cost   Flags NextHop         Interface
                       10.1.1.0/24 Direct 0 0     D   10.1.1.2      Vlanif100
                       10.1.1.2/32 Direct 0 0     D   127.0.0.1     Vlanif100
                      10.1.1.255/32 Direct 0 0     D   127.0.0.1     Vlanif100
                     11.11.11.11/32 EBGP 255 0       RD 10.1.1.1         Vlanif100
                     33.33.33.33/32 IBGP 255 0       RD 3.3.3.9        Vlanif300
                    255.255.255.255/32 Direct 0 0    D    127.0.0.1     InLoopBack0

                    [PE1] display ip routing-table vpn-instance vpnb
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: vpnb
                           Destinations : 6        Routes : 6

                    Destination/Mask Proto Pre Cost  Flags NextHop       Interface
                       10.2.1.0/24 Direct 0 0    D   10.2.1.2    Vlanif200
                       10.2.1.2/32 Direct 0 0    D   127.0.0.1   Vlanif200
                     10.2.1.255/32 Direct 0 0     D   127.0.0.1    Vlanif200
                    22.22.22.22/32 EBGP 255 0       RD 10.2.1.1      Vlanif200


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              75
VPN Configuration
VPN Configuration                                                                       3 IPv4 L3VPN Configuration

                     44.44.44.44/32 IBGP 255 0         RD   3.3.3.9     Vlanif300
                    255.255.255.255/32 Direct 0 0       D   127.0.0.1    InLoopBack0

                    CEs on the same VPN can ping each other, but CEs on different VPNs cannot.
                    For example, CE1 can ping CE3 at 10.3.1.1, but cannot ping CE4 at 10.4.1.1.
                    [CE1] ping -a 11.11.11.11 33.33.33.33
                     PING 33.33.33.33: 56 data bytes, press CTRL_C to break
                       Reply from 33.33.33.33: bytes=56 Sequence=1 ttl=251 time=72 ms
                       Reply from 33.33.33.33: bytes=56 Sequence=2 ttl=251 time=34 ms
                       Reply from 33.33.33.33: bytes=56 Sequence=3 ttl=251 time=50 ms
                       Reply from 33.33.33.33: bytes=56 Sequence=4 ttl=251 time=50 ms
                       Reply from 33.33.33.33: bytes=56 Sequence=5 ttl=251 time=34 ms
                     --- 33.33.33.33 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 34/48/72 ms
                    [CE1] ping -a 11.11.11.11 44.44.44.44
                     PING 44.44.44.44: 56 data bytes, press CTRL_C to break
                       Request time out
                       Request time out
                       Request time out
                       Request time out
                       Request time out
                     --- 44.44.44.44 ping statistics ---
                       5 packet(s) transmitted
                       0 packet(s) received
                       100.00% packet loss


