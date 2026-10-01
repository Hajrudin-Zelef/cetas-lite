---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-155
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22170, 22323]
sha256: 590f10e3fd7afbd0ad441866a8e95d40b98e4d0825fb930b6313929ecd23f926
---

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.1 as-number 100
                    [PE2-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE2-bgp] peer 3.3.3.3 as-number 100
                    [PE2-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [PE2-bgp] ipv6-family vpnv6
                    [PE2-bgp-af-vpnv6] peer 1.1.1.1 enable
                    [PE2-bgp-af-vpnv6] peer 3.3.3.3 enable


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                           350
VPN Configuration
VPN Configuration                                                                             4 IPv6 L3VPN Configuration

                    [PE2-bgp-af-vpnv6] quit
                    [PE2-bgp] quit

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] peer 1.1.1.1 as-number 100
                    [PE3-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE3-bgp] peer 2.2.2.2 as-number 100
                    [PE3-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE3-bgp] ipv6-family vpnv6
                    [PE3-bgp-af-vpnv6] peer 1.1.1.1 enable
                    [PE3-bgp-af-vpnv6] peer 2.2.2.2 enable
                    [PE3-bgp-af-vpnv6] quit
                    [PE3-bgp] quit

                    Run the display bgp vpnv6 all peer command on each PE. The command output
                    shows that the status of the MP-IBGP peer relationship between the PEs is
                    Established.
                    The following example uses the command output on PE1.
                    [PE1] display bgp vpnv6 all peer

                    BGP local router ID : 1.1.1.1
                    Local AS number : 100
                    Total number of peers : 2            Peers in established state : 2

                    Peer        V   AS MsgRcvd MsgSent        OutQ Up/Down          State PrefRcv

                    2.2.2.2     4 100         20    17    0 00:13:26 Established          0
                    3.3.3.3     4 100         24    19    0 00:17:18 Established          1

         Step 5 Configure an IPv6-address-family-enabled VPN instance on the PEs. On PE2 and
                PE3, bind the interfaces connected to the CE to the corresponding VPN instances.
                    # Configure PE1.
                    [PE1] ip vpn-instance vpn1
                    [PE1-vpn-instance-vpn1] ipv6-family
                    [PE1-vpn-instance-vpn1-af-ipv6] route-distinguisher 100:1
                    [PE1-vpn-instance-vpn1-af-ipv6] vpn-target 111:1
                    [PE1-vpn-instance-vpn1-af-ipv6] quit
                    [PE1-vpn-instance-vpn1] quit

                    # Configure PE2.
                    [PE2] ip vpn-instance vpn1
                    [PE2-vpn-instance-vpn1] ipv6-family
                    [PE2-vpn-instance-vpn1-af-ipv6] route-distinguisher 100:2
                    [PE2-vpn-instance-vpn1-af-ipv6] vpn-target 111:1
                    [PE2-vpn-instance-vpn1-af-ipv6] quit
                    [PE2-vpn-instance-vpn1] quit
                    [PE2] interface Vlanif200
                    [PE2-Vlanif200] ip binding vpn-instance vpn1
                    [PE2-Vlanif200] ipv6 enable
                    [PE2-Vlanif200] ipv6 address 2001:db8:1::2 64
                    [PE2-Vlanif200] quit

                    # Configure PE3.
                    [PE3] ip vpn-instance vpn1
                    [PE3-vpn-instance-vpn1] ipv6-family
                    [PE3-vpn-instance-vpn1-af-ipv6] route-distinguisher 100:3
                    [PE3-vpn-instance-vpn1-af-ipv6] vpn-target 111:1
                    [PE3-vpn-instance-vpn1-af-ipv6] quit
                    [PE3-vpn-instance-vpn1] quit
                    [PE3] interface Vlanif200
                    [PE3-Vlanif200] ip binding vpn-instance vpn1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       351
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    [PE3-Vlanif200] ipv6 enable
                    [PE3-Vlanif200] ipv6 address 2001:db8:3::2 64
                    [PE3-Vlanif200] quit

         Step 6 Establish an EBGP peer relationship between PE2 and the CE and between PE3
                and the CE, and import the routes destined for the CE's loopback interface into
                BGP.

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] ipv6-family vpn-instance vpn1
                    [PE2-bgp6-vpn1] peer 2001:db8:1::1 as-number 65410
                    [PE2-bgp6-vpn1] quit
                    [PE2-bgp] quit

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] ipv6-family vpn-instance vpn1
                    [PE3-bgp6-vpn1] peer 2001:db8:3::1 as-number 65410
                    [PE3-bgp6-vpn1] quit
                    [PE3-bgp] quit

                    # Configure the CE.
                    [CE] bgp 65410
                    [CE-bgp] router-id 10.10.10.10
                    [CE-bgp] peer 2001:db8:1::2 as-number 100
                    [CE-bgp] peer 2001:db8:3::2 as-number 100
                    [CE-bgp] ipv6-family unicast
                    [CE-bgp-af-ipv6] peer 2001:db8:1::2 enable
                    [CE-bgp-af-ipv6] peer 2001:db8:3::2 enable
                    [CE-bgp-af-ipv6] network 2001:db8:0:1:2::1 128
                    [CE-bgp-af-ipv6] quit
                    [CE-bgp] quit

                    After completing the configuration, run the display ipv6 routing-table vpn-
                    instance command on PE2. The command output shows the routes destined for
                    the CE's loopback interface.
                    <PE2> display ipv6 routing-table vpn-instance vpn1 2001:db8:0:1:2::1 128
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpn1
                    Summary Count : 1

                    Destination : 2001:db8:0:1:2::1          PrefixLength : 128
                    NextHop      : 2001:db8:1::1             Preference : 255
                    Cost      :0                       Protocol    : EBGP
                    RelayNextHop : 2001:db8:1::1                TunnelID   : 0x0
                    Interface : Vlanif200          Flags      : RD

         Step 7 Configure VPNv6 auto FRR on PE2, and adjust the precedence of EBGP routes for
                PE2 to prefer an EBGP route.

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] ipv6-family vpn-instance vpn1
                    [PE2-bgp6-vpn1] preference 100 255 255
                    [PE2-bgp6-vpn1] auto-frr
                    [PE2-bgp6-vpn1] route-select delay 300
                    [PE2-bgp6-vpn1] quit
                    [PE2-bgp] quit

                    ----End

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         352
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration


