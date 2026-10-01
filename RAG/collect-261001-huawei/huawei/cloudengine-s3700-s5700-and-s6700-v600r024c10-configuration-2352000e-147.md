---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-147
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2021-04-09", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20928, 21080]
sha256: b79abd514b1c8af5596844ebf3bd586b54d9253f2f52abd98bd8a736822d8239
---

                    # Configure PE2.
                    <PE2> system-view
                    [PE2] mpls lsr-id 2.2.2.2
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface Vlanif100
                    [PE2-Vlanif100] mpls
                    [PE2-Vlanif100] mpls ldp
                    [PE2-Vlanif100] quit

                    # Configure PE3.
                    <PE3> system-view
                    [PE3] mpls lsr-id 3.3.3.3
                    [PE3] mpls
                    [PE3-mpls] quit
                    [PE3] mpls ldp
                    [PE3-mpls-ldp] quit
                    [PE3] interface Vlanif100
                    [PE3-Vlanif100] mpls
                    [PE3-Vlanif100] mpls ldp
                    [PE3-Vlanif100] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         332
VPN Configuration
VPN Configuration                                                                                     4 IPv6 L3VPN Configuration


                    Run the display mpls lsp command on the PEs. The command output shows that
                    LSPs have been established between PE1 and PE2 and between PE1 and PE3. The
                    following example uses the command output on PE1.
                    <PE1> display mpls lsp
                    2021-04-09 01:14:17.813
                    Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated
                    LSP
                    Flag after LDP FRR: (L) - Logic FRR LSP
                    -------------------------------------------------------------------------------
                                 LSP Information: LDP LSP
                    -------------------------------------------------------------------------------
                    FEC              In/Out Label In/Out IF                        Vrf Name
                    1.1.1.1/32         3/NULL          -/-
                    2.2.2.2/32         NULL/3          -/Vlanif200
                    2.2.2.2/32         1025/3         -/Vlanif200
                    3.3.3.3/32         NULL/3          -/Vlanif300
                    3.3.3.3/32         1024/3         -/Vlanif300

         Step 4 Configure an IPv6-address-family-enabled VPN instance on the PEs. On PE2 and
                PE3, bind the interfaces connected to the CE to the corresponding VPN instances.
                    # Configure PE1.
                    [PE1] ip vpn-instance vpn1
                    [PE1-vpn-instance-vpn1] ipv6-family
                    [PE1-vpn-instance-vpn1-af-ipv6] route-distinguisher 100:1
                    [PE1-vpn-instance-vpn1-af-ipv6] vpn-target 111:1
                    [PE1-vpn-instance-vpn1-af-ipv6] quit
                    [PE1-vpn-instance-vpn1] quit
                    [PE1] interface loopback2
                    [PE1-Loopback2] ip binding vpn-instance vpn1
                    [PE1-Loopback2] ipv6 enable
                    [PE1-Loopback2] ipv6 address 2001:DB8:8::1 64
                    [PE1-Loobpack2] quit

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
                    [PE2-Vlanif200] ipv6 address 2001:DB8:1::2 64
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
                    [PE3-Vlanif200] ipv6 enable
                    [PE3-Vlanif200] ipv6 address 2001:DB8:3::2 64
                    [PE3-Vlanif200] quit

         Step 5 Establish an EBGP peer relationship between PE2 and the CE and between PE3
                and the CE, and import the routes destined for the CE's loopback interface.
                    # Configure PE2.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                              333
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration

                    [PE2] bgp 100
                    [PE2-bgp] ipv6-family vpn-instance vpn1
                    [PE2-bgp6-vpn1] peer 2001:DB8:1::1 as-number 65410

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] ipv6-family vpn-instance vpn1
                    [PE3-bgp6-vpn1] peer 2001:DB8:3::1 as-number 65410

                    # Configure the CE.
                    <CE> system-view
                    [CE] bgp 65410
                    [CE-bgp] router-id 10.10.10.10
                    [CE-bgp] peer 2001:DB8:1::2 as-number 100
                    [CE-bgp] peer 2001:DB8:3::2 as-number 100
                    [CE-bgp] ipv6-family unicast
                    [CE-bgp-af-ipv6] peer 2001:DB8:1::2 enable
                    [CE-bgp-af-ipv6] peer 2001:DB8:3::2 enable
                    [CE-bgp-af-ipv6] network 2001:DB8:0:1:2::1 128
                    [CE-bgp-af-ipv6] quit
                    [CE-bgp] quit

                    After completing the configuration, run the display bgp vpnv6 vpn-instance
                    vpn1 peer command on PE2 and PE3. The command output shows that the status
                    of the EBGP peer relationship between PE2 and the CE and between PE3 and the
                    CE is Established.
                    The following example uses the command output on PE2.
                    <PE2> display bgp vpnv6 vpn-instance vpn1 peer
                     BGP local router ID : 2.2.2.2
                     Local AS number : 100
                     Total number of peers : 1        Peers in established state : 1
                      Peer        V AS MsgRcvd MsgSent OutQ Up/Down              State PrefRcv
                      2001:DB8:1::1 4 65001        46 46   0 00:08:36 Established      5

         Step 6 Establish MP-IBGP peer relationships between the PEs.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.2 as-number 100
                    [PE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE1-bgp] peer 3.3.3.3 as-number 100
                    [PE1-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [PE1-bgp] ipv6-family vpnv6
                    [PE1-bgp-af-vpnv6] peer 2.2.2.2 enable
                    [PE1-bgp-af-vpnv6] peer 3.3.3.3 enable
                    [PE1-bgp-af-vpnv6] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.1 as-number 100
                    [PE2-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE2-bgp] ipv6-family vpnv6
                    [PE2-bgp-af-vpnv6] peer 1.1.1.1 enable
                    [PE2-bgp-af-vpnv6] quit

