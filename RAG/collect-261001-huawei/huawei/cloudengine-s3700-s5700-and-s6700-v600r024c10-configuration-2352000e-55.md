---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-55
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7047, 7206]
sha256: 79a5959cb8b5ca4aab11d81585e5a1e033b2522e2f99a8e74a1b5779462fbc6a
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

                    Run the display mpls lsp command on the PEs. The command output shows that
                    LSPs have been established between PE1 and PE2 and between PE1 and PE3. The
                    following example uses the command output on PE1.
                    <PE1> display mpls lsp
                    Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated
                    LSP
                    Flag after LDP FRR: (L) - Logic FRR LSP
                    -------------------------------------------------------------------------------
                                 LSP Information: LDP LSP
                    -------------------------------------------------------------------------------
                    FEC              In/Out Label In/Out IF                        Vrf Name
                    1.1.1.1/32         3/NULL          -/-
                    2.2.2.2/32         NULL/3          -/Vlanif200
                    2.2.2.2/32         4096/3         -/Vlanif200
                    3.3.3.3/32         NULL/3          -/Vlanif300
                    3.3.3.3/32         4097/3         -/Vlanif300

         Step 4 Configure a VPN instance on each PE, and bind the interface connected to the CE
                to the VPN instance on PE2 and PE3.

                    # Configure PE1.
                    [PE1] ip vpn-instance vpn1
                    [PE1-vpn-instance-vpn1] ipv4-family
                    [PE1-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:1
                    [PE1-vpn-instance-vpn1-af-ipv4] vpn-target 111:1
                    [PE1-vpn-instance-vpn1-af-ipv4] quit
                    [PE1-vpn-instance-vpn1] quit

                    # Configure PE2.
                    [PE2] ip vpn-instance vpn1
                    [PE2-vpn-instance-vpn1] ipv4-family
                    [PE2-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:2
                    [PE2-vpn-instance-vpn1-af-ipv4] vpn-target 111:1
                    [PE2-vpn-instance-vpn1-af-ipv4] quit
                    [PE2-vpn-instance-vpn1] quit
                    [PE2] interface Vlanif200
                    [PE2-Vlanif200] ip binding vpn-instance vpn1
                    [PE2-Vlanif200] ip address 10.1.1.2 30
                    [PE2-Vlanif200] quit


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                              113
VPN Configuration
VPN Configuration                                                                       3 IPv4 L3VPN Configuration


                    # Configure PE3.
                    [PE3] ip vpn-instance vpn1
                    [PE3-vpn-instance-vpn1] ipv4-family
                    [PE3-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:3
                    [PE3-vpn-instance-vpn1-af-ipv4] vpn-target 111:1
                    [PE3-vpn-instance-vpn1-af-ipv4] quit
                    [PE3-vpn-instance-vpn1] quit
                    [PE3] interface Vlanif200
                    [PE3-Vlanif200] ip binding vpn-instance vpn1
                    [PE3-Vlanif200] ip address 10.2.1.2 30
                    [PE3-Vlanif200] quit

         Step 5 Establish an EBGP peer relationship between PE2 and the CE and between PE3
                and the CE, and import the routes destined for the CE's loopback interface.
                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] ipv4-family vpn-instance vpn1
                    [PE2-bgp-vpn1] peer 10.1.1.1 as-number 65410

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] ipv4-family vpn-instance vpn1
                    [PE3-bgp-vpn1] peer 10.2.1.1 as-number 65410

                    # Configure the CE.
                    [CE] interface loopback 1
                    [CE-Loopback1] ip address 11.11.11.11 32
                    [CE-Loopback1] quit
                    [CE] bgp 65410
                    [CE-bgp] peer 10.1.1.2 as-number 100
                    [CE-bgp] peer 10.2.1.2 as-number 100
                    [CE-bgp] network 11.11.11.11 32
                    [CE-bgp] quit

                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] ipv4-family vpn-instance vpn1
                    [PE1-bgp-vpn1] quit

                    After completing the configuration, run the display bgp vpnv4 vpn-instance peer
                    command on PE2 and PE3. The command output shows that an EBGP peer
                    relationship has been established between PE2 and the CE and between PE3 and
                    the CE.
                    The following example uses the command output on PE2.
                    <PE2> display bgp vpnv4 vpn-instance vpn1 peer
                     BGP local router ID : 10.10.1.2
                     Local AS number : 100

                    VPN-Instance vpn1, Router ID 10.10.1.2:
                    Total number of peers : 1          Peers in established state : 1

                     Peer        V       AS MsgRcvd MsgSent OutQ Up/Down         State PrefRcv
                     10.1.1.1    4      65410   29   28   0 00:22:20 Established      1

         Step 6 Establish MP-IBGP peer relationships between the PEs.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.2 as-number 100


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                114
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    [PE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE1-bgp] peer 3.3.3.3 as-number 100
                    [PE1-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [PE1-bgp] ipv4-family vpnv4
                    [PE1-bgp-af-vpnv4] peer 2.2.2.2 enable
                    [PE1-bgp-af-vpnv4] peer 3.3.3.3 enable
                    [PE1-bgp-af-vpnv4] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.1 as-number 100
                    [PE2-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE2-bgp] ipv4-family vpnv4
                    [PE2-bgp-af-vpnv4] peer 1.1.1.1 enable
                    [PE2-bgp-af-vpnv4] quit

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] peer 1.1.1.1 as-number 100
                    [PE3-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE3-bgp] ipv4-family vpnv4
                    [PE3-bgp-af-vpnv4] peer 1.1.1.1 enable
                    [PE3-bgp-af-vpnv4] quit

