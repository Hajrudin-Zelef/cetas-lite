---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-154
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2021-04-09", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22023, 22169]
sha256: 89c140cf7b75f5234270bc72d85ab40122cd240f7124c92803e76254a4d4cedb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        IPv6+VPNv6 hybrid FRR applies only to the networking where CEs establish BGP peer
                        relationships with PEs.

                    As shown in Figure 4-6, a CE is dual-homed to PE2 and PE3; an MPLS public
                    network tunnel and a VPNv6 peer relationship are established between PE2 and
                    PE3. PE2 and PE3 use EBGP to exchange routing information with the CE. PE3
                    learns from the CE a route to Loopback 1 on the CE and advertises the route to its
                    VPNv6 peer. PE2 then has two BGP routes to Loopback 1 on the CE. One is sent
                    from the CE using EBGP, and the other is sent from PE3 using MP-IBGP.
                    It is required that PE2 be configured to prefer the EBGP route sent from the CE
                    and use the VPNv6 route sent from PE3 as a backup route. If the link between PE2
                    and the CE is faulty, traffic destined for the CE can be switched to PE3 that serves
                    as the backup next hop.

                    Figure 4-6 Network diagram of IPv6+VPNv6 hybrid FRR
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Precautions
                    In a VPN FRR scenario, traffic is switched back to the primary path after this path
                    recovers. Because the order in which nodes undergo IGP convergence differs,
                    packet loss may occur during the switchback. To resolve this problem, run the

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     348
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration


                    route-select delay delay-value command to configure a route selection delay so
                    that traffic is switched back only after forwarding entries on the devices along the
                    primary path are updated. The delay specified using delay-value depends on
                    various factors, such as the number of routes on each device. Configure an
                    appropriate delay as needed.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure OSPF on PE1, PE2, and PE3 for them to communicate on the MPLS
                         backbone network.
                    2.   Configure basic MPLS capabilities and enable MPLS LDP to establish LDP LSPs
                         on the MPLS backbone network.
                    3.   Establish an MP-IBGP peer relationship between PE1, PE2, and PE3.
                    4.   Configure an IPv6-address-family-enabled VPN instance on the PEs. On PE2
                         and PE3, bind the interfaces connected to the CE to the corresponding VPN
                         instances.
                    5.   Establish EBGP peer relationships between the PEs and the CE, and import the
                         routes destined for the CE's loopback interface into BGP.
                    6.   Configure auto FRR for the BGP VPN instance IPv6 address family on PE2 so
                         that the VPNv6 route sent from PE3 can serve as a backup route.

Procedure
         Step 1 Configure IPv6 addresses for interfaces on the VPN backbone network and at VPN
                sites. For detailed configurations, see Configuration Scripts.
         Step 2 Configure OSPF on the MPLS backbone network for PEs to communicate. For
                detailed configurations, see Configuration Scripts.
         Step 3 Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                MPLS backbone network.
                    # Configure PE1.
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


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         349
VPN Configuration
VPN Configuration                                                                                 4 IPv6 L3VPN Configuration

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
                    2021-04-09 01:14:17.813
                    Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated
                    LSP
                    Flag after LDP FRR: (L) - Logic FRR LSP
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
                    [PE1-bgp] peer 2.2.2.2 connect-interface loopback 1
                    [PE1-bgp] peer 3.3.3.3 as-number 100
                    [PE1-bgp] peer 3.3.3.3 connect-interface loopback 1
                    [PE1-bgp] ipv6-family vpnv6
                    [PE1-bgp-af-vpnv6] peer 2.2.2.2 enable
                    [PE1-bgp-af-vpnv6] peer 3.3.3.3 enable
                    [PE1-bgp-af-vpnv6] quit
                    [PE1-bgp] quit

