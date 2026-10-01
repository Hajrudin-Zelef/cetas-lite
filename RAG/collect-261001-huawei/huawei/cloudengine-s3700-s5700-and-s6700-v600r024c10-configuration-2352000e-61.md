---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-61
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7962, 8123]
sha256: b068085900917712358e09be806f25a403c566c8969a262eda588bb6e5c130ca
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Procedure
         Step 1 Configure IP addresses for involved interfaces on the VPN backbone network and
                at VPN sites. For detailed configurations, see Configuration Scripts.
         Step 2 Configure OSPF on the MPLS backbone network for PEs to communicate. For
                detailed configurations, see Configuration Scripts.
         Step 3 Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                MPLS backbone network.
                    # Configure PE1.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            126
VPN Configuration
VPN Configuration                                                                            3 IPv4 L3VPN Configuration

                    <PE1> system-view
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif200
                    [PE1-Vlanif200] mpls
                    [PE1-Vlanif200] mpls ldp
                    [PE1-Vlanif200] quit
                    [PE1] interface Vlanif300
                    [PE1-Vlanif300] mpls
                    [PE1-Vlanif300] mpls ldp
                    [PE1-Vlanif300] quit

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
                    [PE2] interface Vlanif300
                    [PE2-Vlanif300] mpls
                    [PE2-Vlanif300] mpls ldp
                    [PE2-Vlanif300] quit

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
                    [PE3] interface Vlanif300
                    [PE3-Vlanif300] mpls
                    [PE3-Vlanif300] mpls ldp
                    [PE3-Vlanif300] quit

                    Run the display mpls lsp command on the PEs. The command output shows that
                    LSPs have been established between PE1 and PE2 and between PE1 and PE3. The
                    following example uses the command output on PE1.
                    [PE1] display mpls lsp
                    ----------------------------------------------------------------------
                                 LSP Information: LDP LSP
                    ----------------------------------------------------------------------
                    FEC              In/Out Label In/Out IF                      Vrf Name
                    2.2.2.2/32         NULL/3         -/Vlanif200
                    2.2.2.2/32         1024/3        -/Vlanif200
                    3.3.3.3/32         NULL/3         -/Vlanif300
                    3.3.3.3/32         1025/3        -/Vlanif300

         Step 4 Establish MP-IBGP peer relationships between the PEs.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.2 as-number 100


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                    127
VPN Configuration
VPN Configuration                                                                             3 IPv4 L3VPN Configuration

                    [PE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE1-bgp] peer 3.3.3.3 as-number 100
                    [PE1-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [PE1-bgp] ipv4-family vpnv4
                    [PE1-bgp-af-vpnv4] peer 2.2.2.2 enable
                    [PE1-bgp-af-vpnv4] peer 3.3.3.3 enable
                    [PE1-bgp-af-vpnv4] quit
                    [PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.1 as-number 100
                    [PE2-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE2-bgp] peer 3.3.3.3 as-number 100
                    [PE2-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [PE2-bgp] ipv4-family vpnv4
                    [PE2-bgp-af-vpnv4] peer 1.1.1.1 enable
                    [PE2-bgp-af-vpnv4] peer 3.3.3.3 enable
                    [PE2-bgp-af-vpnv4] quit
                    [PE2-bgp] quit

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] peer 1.1.1.1 as-number 100
                    [PE3-bgp] peer 1.1.1.1 connect-interface loopback 1
                    [PE3-bgp] peer 2.2.2.2 as-number 100
                    [PE3-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE3-bgp] ipv4-family vpnv4
                    [PE3-bgp-af-vpnv4] peer 1.1.1.1 enable
                    [PE3-bgp-af-vpnv4] peer 2.2.2.2 enable
                    [PE3-bgp-af-vpnv4] quit
                    [PE3-bgp] quit

                    Run the display bgp vpnv4 all peer command on each PE. The command output
                    shows that the status of the MP-IBGP peer relationship between PEs is
                    Established.
                    The following example uses the command output on PE1.
                    [PE1] display bgp vpnv4 all peer

                    BGP local router ID : 1.1.1.1
                    Local AS number : 100
                    Total number of peers : 2            Peers in established state : 2

                    Peer        V   AS MsgRcvd MsgSent        OutQ Up/Down          State PrefRcv

                    2.2.2.2     4 100        20     17    0 00:13:26 Established          0
                    3.3.3.3     4 100        24     19    0 00:17:18 Established          1

         Step 5 Configure a VPN instance on each PE, and bind the interface connected to the CE
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


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       128
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    [PE2-vpn-instance-vpn1-af-ipv4] vpn-target 111:1
                    [PE2-vpn-instance-vpn1-af-ipv4] quit
                    [PE2-vpn-instance-vpn1] quit
                    [PE2] interface Vlanif200
                    [PE2-Vlanif200] ip binding vpn-instance vpn1
                    [PE2-Vlanif200] ip address 192.168.1.1 30
                    [PE2-Vlanif200] quit

