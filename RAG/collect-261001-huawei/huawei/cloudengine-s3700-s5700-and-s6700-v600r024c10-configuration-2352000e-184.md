---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-184
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [26749, 26911]
sha256: 49cfd70d10c19db28db354c207f5d65370d8e3ce8f87084d97a83f5b5a48ba25
---

                    # Configure the P.
                    <HUAWEI> system-view
                    [HUAWEI] sysname P
                    [P] vlan batch 10 20
                    [P] interface loopback 1
                    [P-LoopBack1] ip address 2.2.2.9 32
                    [P-LoopBack1] quit
                    [P] interface 10ge 1/0/1
                    [P-10GE1/0/1] port link-type trunk
                    [P-10GE1/0/1] port trunk allow-pass vlan 10
                    [P-10GE1/0/1] quit
                    [P] interface 10ge 1/0/2
                    [P-10GE1/0/2] port link-type trunk
                    [P-10GE1/0/2] port trunk allow-pass vlan 20
                    [P-10GE1/0/2] quit
                    [P] interface vlanif 10
                    [P-Vlanif10] ip address 10.2.2.2 24
                    [P-Vlanif10] quit
                    [P] interface vlanif 20
                    [P-Vlanif20] ip address 10.1.1.2 24
                    [P-Vlanif20] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] vlan batch 10 20
                    [PE2] interface loopback 1
                    [PE2-LoopBack1] ip address 3.3.3.9 32
                    [PE2-LoopBack1] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] port link-type trunk
                    [PE2-10GE1/0/1] port trunk allow-pass vlan 10
                    [PE2-10GE1/0/1] quit
                    [PE2] interface vlanif 10
                    [PE2-Vlanif10] ip address 10.2.2.1 24
                    [PE2-Vlanif10] quit

         Step 2 Configure basic MPLS TE functions on the MPLS backbone network.

                    # Configure PE1.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   425
VPN Configuration
VPN Configuration                                                                                 5 VPWS Configuration

                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] mpls te
                    [PE1-mpls] quit
                    [PE1] interface Vlanif 20
                    [PE1-Vlanif20] mpls
                    [PE1-Vlanif20] mpls te
                    [PE1-Vlanif20] quit

                    # Configure the P.
                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] mpls te
                    [P-mpls] quit
                    [P] interface Vlanif 10
                    [P-Vlanif10] mpls
                    [P-Vlanif10] mpls te
                    [P-Vlanif10] quit
                    [P] interface Vlanif 20
                    [P-Vlanif20] mpls
                    [P-Vlanif20] mpls te
                    [P-Vlanif20] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 3.3.3.9
                    [PE2] mpls
                    [PE2-mpls] mpls te
                    [PE2-mpls] quit
                    [PE2] interface Vlanif 10
                    [PE2-Vlanif10] mpls
                    [PE2-Vlanif10] mpls te
                    [PE2-Vlanif10] quit

         Step 3 Configure static CR-LSPs on the P.

                    # Configure the P. Specifically, configure a static CR-LSP for transmitting packets
                    from PE1 to PE2 and another one for transmitting packets from PE2 to PE1.
                    [P] static-cr-lsp transit PE1-PE2 incoming-interface vlanif 20 in-label 200 nexthop 10.2.2.1 out-label
                    201
                    [P] static-cr-lsp transit PE2-PE1 incoming-interface vlanif 10 in-label 101 nexthop 10.1.1.1 out-label
                    100

         Step 4 Create remote CCC connections on PEs.

                    # Configure PE1. Specifically, enable MPLS L2VPN globally, and create a CE1-to-
                    CE2 remote CCC connection, with the inbound interface connecting to CE1,
                    outbound interface connecting to the P, incoming label being 100, and outgoing
                    label being 200.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] quit
                    [PE1] ccc CE1-CE2 interface vlanif10 in-label 100 out-label 200 nexthop 10.1.1.2

                    # Configure PE2. Specifically, enable MPLS L2VPN globally; create a CE2-to-CE1
                    remote CCC connection, with the inbound interface connecting to CE2, outbound
                    interface connecting to the P, incoming label being 201, and outgoing label being
                    101.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] interface vlanif 20


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                             426
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration

                    [PE2-Vlanif20] quit
                    [PE2] ccc CE2-CE1 interface vlanif20 in-label 201 out-label 101 nexthop 10.2.2.2

                    ----End

Verifying the Configuration
                    # Check CCC connection information on PEs. The command outputs show that a
                    remote CCC connection has been established on each of PE1 and PE2 and is in the
                    up state.
                    <PE1> display vll ccc
                    total ccc vc : 1
                    local ccc vc : 0, 0 up
                    remote ccc vc : 1, 1 up


                    name: CE1-CE2, type: remote, state: up,
                    intf: Vlanif10 (up), in-label: 100 , out-label: 200 , nexthop : 10.1.1.2
                    VC last up time : 2024/03/02 08:17:36
                    VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds
                    <PE2> display vll ccc
                    total ccc vc : 1
                    local ccc vc : 0, 0 up
                    remote ccc vc : 1, 1 up


                    name: CE2-CE1, type: remote, state: up,
                    intf: Vlanif20 (up), in-label: 201 , out-label: 101 , nexthop : 10.2.2.2
                    VC last up time : 2024/03/02 08:17:50
                    VC total up time: 0 days, 2 hours, 12 minutes, 51 seconds

                    # Check information about the interfaces used for the CCC connections on PEs.
                    The command output shows that the VC type is CCC and the VC status is up. The
                    following example uses the command output on PE1.
                    <PE1> display l2vpn ccc-interface vc-type ccc
                    Total ccc-interface of CCC : 1
                    up (1), down (0)
                    Interface               Encap Type         State   VC Type
                    Vlanif10                ethernet         up      ccc

                    # Check LSP information on the P. The command output shows information about
                    the labels and interfaces of the two static CR-LSPs.
                    <P> display mpls lsp
                    ----------------------------------------------------------------------
                                 LSP Information: STATIC LSP
                    ----------------------------------------------------------------------
                    FEC              In/Out Label In/Out IF                      Vrf Name
                    -/-             200/201       Vlanif20/Vlanif10
                    -/-             101/100       Vlanif10/Vlanif20

