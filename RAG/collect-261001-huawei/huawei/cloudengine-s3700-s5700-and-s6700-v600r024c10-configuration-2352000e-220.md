---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-220
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [32383, 32563]
sha256: 4e1faaeb83112472aa5045eef8e3eae5eb8e4c1499b0508d5c3effd436979089
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The configurations of PE1, PE2, PE3, P, and CE2 are similar to the configuration of
                    CE1. For detailed configurations, see Configuration Scripts.
         Step 2 Configure an IGP on the MPLS backbone network to allow the PEs and P to
                communicate with each other.
                    # Configure PE1.
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.1 32
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 10.0.2.0 0.0.0.3
                    [PE1-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.3
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure the P.
                    [P] interface loopback 1
                    [P-LoopBack1] ip address 4.4.4.4 32
                    [P-LoopBack1] quit
                    [P] ospf 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                        518
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration

                    [P-ospf-1] area 0
                    [P-ospf-1-area-0.0.0.0] network 10.0.2.0 0.0.0.3
                    [P-ospf-1-area-0.0.0.0] network 10.0.3.0 0.0.0.3
                    [P-ospf-1-area-0.0.0.0] network 4.4.4.4 0.0.0.0
                    [P-ospf-1-area-0.0.0.0] quit
                    [P-ospf-1] quit

                    # Configure PE3.
                    [PE3] interface loopback 1
                    [PE3-LoopBack1] ip address 3.3.3.3 32
                    [PE3-LoopBack1] quit
                    [PE3] ospf 1
                    [PE3-ospf-1] area 0
                    [PE3-ospf-1-area-0.0.0.0] network 3.3.3.3 0.0.0.0
                    [PE3-ospf-1-area-0.0.0.0] network 10.0.3.0 0.0.0.3
                    [PE3-ospf-1-area-0.0.0.0] quit
                    [PE3-ospf-1] quit

                    # Configure PE2.
                    [PE2] interface loopback 1
                    [PE2-LoopBack1] ip address 2.2.2.2 32
                    [PE2-LoopBack1] quit
                    [PE2] ospf 1
                    [PE2-ospf-1] area 0
                    [PE2-ospf-1-area-0.0.0.0] network 2.2.2.2 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] network 10.2.1.2 0.0.0.3
                    [PE2-ospf-1-area-0.0.0.0] quit
                    [PE2-ospf-1] quit

                    After the configurations are complete, run the display ip routing-table command
                    on PEs. The command outputs show that PE1 and PE2, as well as PE1 and PE3
                    have learned the routes to each other's Loopback1 interface.
         Step 3 Configure basic MPLS functions on the MPLS backbone network.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] mpls
                    [PE1-Vlanif10] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls
                    [PE1-Vlanif20] quit

                    The configurations of PE2, PE3, and the P are similar to the configuration of PE1.
                    For detailed configurations, see Configuration Scripts.
         Step 4 Establish an MPLS TE tunnel between PE1 and PE3, and an LSP between PE1 and
                PE2.
                    # Configure PE1.
                    [PE1] mpls
                    [PE1-mpls] mpls te
                    [PE1-mpls] mpls rsvp-te
                    [PE1-mpls] mpls te cspf
                    [PE1-mpls] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls te
                    [PE1-Vlanif20] mpls rsvp-te
                    [PE1-Vlanif20] quit
                    [PE1] interface tunnel 2
                    [PE1-Tunnel2] ip address unnumbered interface loopback1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                       519
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

                    [PE1-Tunnel2] tunnel-protocol mpls te
                    [PE1-Tunnel2] destination 3.3.3.3
                    [PE1-Tunnel2] mpls te tunnel-id 13
                    [PE1-Tunnel2] mpls te
                    [PE1-Tunnel2] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] opaque-capability enable
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] mpls-te enable
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure the P.
                    [P] mpls
                    [P-mpls] mpls te
                    [P-mpls] mpls rsvp-te
                    [P-mpls] quit
                    [P] interface vlanif 10
                    [P-Vlanif10] mpls te
                    [P-Vlanif10] mpls rsvp-te
                    [P-Vlanif10] quit
                    [P] interface vlanif 20
                    [P-Vlanif20] mpls te
                    [P-Vlanif20] mpls rsvp-te
                    [P-Vlanif20] quit
                    [P] ospf 1
                    [P-ospf-1] opaque-capability enable
                    [P-ospf-1] area 0
                    [P-ospf-1-area-0.0.0.0] mpls-te enable
                    [P-ospf-1-area-0.0.0.0] quit
                    [P-ospf-1] quit

                    # Configure PE3.
                    [PE3] mpls
                    [PE3-mpls] mpls te
                    [PE3-mpls] mpls rsvp-te
                    [PE3-mpls] mpls te cspf
                    [PE3-mpls] quit
                    [PE3] interface vlanif 10
                    [PE3-Vlanif10] mpls te
                    [PE3-Vlanif10] mpls rsvp-te
                    [PE3-Vlanif10] quit
                    [PE3] interface tunnel 2
                    [PE3-Tunnel2] ip address unnumbered interface LoopBack1
                    [PE3-Tunnel2] tunnel-protocol mpls te
                    [PE3-Tunnel2] destination 1.1.1.1
                    [PE3-Tunnel2] mpls te tunnel-id 31
                    [PE3-Tunnel2] mpls te
                    [PE3-Tunnel2] quit
                    [PE3] ospf 1
                    [PE3-ospf-1] opaque-capability enable
                    [PE3-ospf-1] area 0
                    [PE3-ospf-1-area-0.0.0.0] mpls-te enable
                    [PE3-ospf-1-area-0.0.0.0] quit
                    [PE3-ospf-1] quit

                    # Configure PE1.
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface vlanif 30
                    [PE1-Vlanif30] mpls ldp
                    [PE1-Vlanif30] quit

                    # Configure PE2.
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   520
VPN Configuration
VPN Configuration                                                                                    5 VPWS Configuration

                    [PE2] interface vlanif 30
                    [PE2-Vlanif30] mpls ldp
                    [PE2-Vlanif30] quit

         Step 5 Establish a remote LDP session between the PEs.
                    Configure a remote LDP session, and use the loopback interface address of a
                    remote LDP peer as the remote peer IP address.

                          NOTE

                         In this example, PE1 and PE2 are directly connected, removing the need to configure a
                         remote LDP session between them.

