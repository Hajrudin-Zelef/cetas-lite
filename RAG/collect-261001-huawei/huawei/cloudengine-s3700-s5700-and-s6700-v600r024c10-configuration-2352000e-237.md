---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-237
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [34860, 34994]
sha256: 8a9bed579fb78f71ad1130652367c61152b889ed7d8aed517bc8e380a9a3c45a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure MPLS functions.
                    2.   Establish PW1 between PE1 and PE2 as the primary PW and PW2 between
                         PE1 and PE3 as the secondary PW.
                    3.   Configure BFD to detect the connectivity of PW1 and PW2.

Procedure
         Step 1 Assign IP addresses to CE interfaces connected to PEs.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] undo portswitch
                    [CE1-10GE1/0/1] ip address 10.1.1.1 30
                    [CE1-10GE1/0/1] ip address 10.1.2.1 30 sub
                    [CE1-10GE1/0/1] quit

                    # Configure CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] undo portswitch
                    [CE2-10GE1/0/1] ip address 10.1.1.2 30
                    [CE2-10GE1/0/1] quit
                    [CE2] interface 10ge 1/0/2
                    [CE2-10GE1/0/2] undo portswitch
                    [CE2-10GE1/0/2] ip address 10.1.2.2 30
                    [CE2-10GE1/0/2] quit

         Step 2 Configure an IGP on the MPLS backbone network to allow the PEs and P to
                communicate with each other.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   554
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration


                    # Configure PE1.
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.1 32
                    [PE1-LoopBack1] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] undo portswitch
                    [PE1-10GE1/0/1] ip address 10.2.1.1 30
                    [PE1-10GE1/0/1] quit
                    [PE1] interface 10ge 1/0/3
                    [PE1-10GE1/0/3] undo portswitch
                    [PE1-10GE1/0/3] ip address 10.4.1.1 30
                    [PE1-10GE1/0/3] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.3
                    [PE1-ospf-1-area-0.0.0.0] network 10.4.1.0 0.0.0.3
                    [PE1-ospf-1-area-0.0.0.0] quit

                    # Configure P1.
                    [P1] interface loopback 1
                    [P1-LoopBack1] ip address 2.2.2.2 32
                    [P1-LoopBack1] quit
                    [P1] interface 10ge 1/0/1
                    [P1-10GE1/0/1] undo portswitch
                    [P1-10GE1/0/1] ip address 10.3.1.1 30
                    [P1-10GE1/0/1] quit
                    [P1] interface 10ge 1/0/2
                    [P1-10GE1/0/2] undo portswitch
                    [P1-10GE1/0/2] ip address 10.2.1.2 30
                    [P1-10GE1/0/2] quit
                    [P1] ospf 1
                    [P1-ospf-1] area 0
                    [P1-ospf-1-area-0.0.0.0] network 2.2.2.2 0.0.0.0
                    [P1-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.3
                    [P1-ospf-1-area-0.0.0.0] network 10.3.1.0 0.0.0.3
                    [P1-ospf-1-area-0.0.0.0] quit

                    # Configure P2.
                    [P2] interface loopback 1
                    [P2-LoopBack1] ip address 3.3.3.3 32
                    [P2-LoopBack1] quit
                    [P2] interface 10ge 1/0/1
                    [P2-10GE1/0/1] undo portswitch
                    [P2-10GE1/0/1] ip address 10.5.1.1 30
                    [P2-10GE1/0/1] quit
                    [P2] interface 10ge 1/0/2
                    [P2-10GE1/0/2] undo portswitch
                    [P2-10GE1/0/2] ip address 10.4.1.2 30
                    [P2-10GE1/0/2] quit
                    [P2] ospf 1
                    [P2-ospf-1] area 0
                    [P2-ospf-1-area-0.0.0.0] network 3.3.3.3 0.0.0.0
                    [P2-ospf-1-area-0.0.0.0] network 10.4.1.0 0.0.0.3
                    [P2-ospf-1-area-0.0.0.0] network 10.5.1.0 0.0.0.3
                    [P2-ospf-1-area-0.0.0.0] quit

                    # Configure PE2.
                    [PE2] interface loopback 1
                    [PE2-LoopBack1] ip address 4.4.4.4 32
                    [PE2-LoopBack1] quit
                    [PE2] interface 10ge 1/0/2
                    [PE2-10GE1/0/2] undo portswitch
                    [PE2-10GE1/0/2] ip address 10.3.1.2 30
                    [PE2-10GE1/0/2] quit
                    [PE2] ospf 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   555
VPN Configuration
VPN Configuration                                                                                      5 VPWS Configuration

                    [PE2-ospf-1] area 0
                    [PE2-ospf-1-area-0.0.0.0] network 4.4.4.4 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] network 10.3.1.0 0.0.0.3
                    [PE2-ospf-1-area-0.0.0.0] quit

                    # Configure PE3.
                    [PE3] interface loopback 1
                    [PE3-LoopBack1] ip address 5.5.5.5 32
                    [PE3-LoopBack1] quit
                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/1] undo portswitch
                    [PE3-10GE1/0/1] ip address 10.5.1.2 30
                    [PE3-10GE1/0/1] quit
                    [PE3] ospf 1
                    [PE3-ospf-1] area 0
                    [PE3-ospf-1-area-0.0.0.0] network 5.5.5.5 0.0.0.0
                    [PE3-ospf-1-area-0.0.0.0] network 10.5.1.0 0.0.0.3
                    [PE3-ospf-1-area-0.0.0.0] quit

                    After the configurations are complete, run the display ip routing-table command
                    on PEs. The command outputs show that PE1 and PE2, as well as PE1 and PE3
                    have learned the routes to each other's Loopback1 interface.

