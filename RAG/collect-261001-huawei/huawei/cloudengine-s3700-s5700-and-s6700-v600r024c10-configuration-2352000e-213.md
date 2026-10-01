---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-213
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [31384, 31560]
sha256: 83cbabba16260a19fa1377ce7bd9a9588c20d235248956f314b8e062d3890697
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

                    [PE1-Loopback1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] undo portswitch
                    [PE1-10GE1/0/2] ip address 10.10.1.1 24
                    [PE1-10GE1/0/2] quit

                    # Configure ASBR1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR1
                    [ASBR1] interface loopback1
                    [ASBR1-Loopback1] ip address 2.2.2.2 32
                    [ASBR1-Loopback1] quit
                    [ASBR1] interface 10ge 1/0/1
                    [ASBR1-10GE1/0/1] undo portswitch
                    [ASBR1-10GE1/0/1] ip address 10.10.1.2 24
                    [ASBR1-10GE1/0/1] quit

                    # Configure ASBR2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname ASBR2
                    [ASBR2] interface loopback1
                    [ASBR2-Loopback1] ip address 3.3.3.3 32
                    [ASBR2-Loopback1] quit
                    [ASBR2] interface 10ge 1/0/2
                    [ASBR2-10GE1/0/2] undo portswitch
                    [ASBR2-10GE1/0/2] ip address 10.20.1.1 24
                    [ASBR2-10GE1/0/2] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] interface loopback1
                    [PE2-Loopback1] ip address 4.4.4.4 32
                    [PE2-Loopback1] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] undo portswitch
                    [PE2-10GE1/0/1] ip address 10.20.1.2 24
                    [PE2-10GE1/0/1] quit
                    [PE2] interface 10ge 1/0/2
                    [PE2-10GE1/0/2] undo portswitch
                    [PE2-10GE1/0/2] quit

                    # Configure CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] undo portswitch
                    [CE2-10GE1/0/1] ip address 10.1.1.2 24
                    [CE2-10GE1/0/1] quit

         Step 2 Configure an IGP on the backbone network.
                    # Configure PE1.
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 10.10.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure ASBR1.
                    [ASBR1] ospf 1
                    [ASBR1-ospf-1] area 0.0.0.0
                    [ASBR1-ospf-1-area-0.0.0.0] network 2.2.2.2 0.0.0.0


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   502
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

                    [ASBR1-ospf-1-area-0.0.0.0] network 10.10.1.0 0.0.0.255
                    [ASBR1-ospf-1-area-0.0.0.0] quit
                    [ASBR1-ospf-1] quit

                    # Configure ASBR2.
                    [ASBR2] ospf 1
                    [ASBR2-ospf-1] area 0.0.0.0
                    [ASBR2-ospf-1-area-0.0.0.0] network 3.3.3.3 0.0.0.0
                    [ASBR2-ospf-1-area-0.0.0.0] network 10.20.1.0 0.0.0.255
                    [ASBR2-ospf-1-area-0.0.0.0] quit
                    [ASBR2-ospf-1] quit

                    # Configure PE2.
                    [PE2] ospf 1
                    [PE2-ospf-1] area 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] network 4.4.4.4 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] network 10.20.1.0 0.0.0.255
                    [PE2-ospf-1-area-0.0.0.0] quit
                    [PE2-ospf-1] quit

         Step 3 Enable MPLS and establish LSPs.

                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] mpls
                    [PE1-10GE1/0/2] mpls ldp
                    [PE1-10GE1/0/2] quit

                    # Configure ASBR1.
                    [ASBR1] mpls lsr-id 2.2.2.2
                    [ASBR1] mpls
                    [ASBR1-mpls] quit
                    [ASBR1] mpls ldp
                    [ASBR1-mpls-ldp] quit
                    [ASBR1] interface 10ge 1/0/1
                    [ASBR1-10GE1/0/1] mpls
                    [ASBR1-10GE1/0/10] mpls ldp
                    [*ASBR1-10GE1/0/1] quit

                    # Configure ASBR2.
                    [ASBR2] mpls lsr-id 3.3.3.3
                    [ASBR2] mpls
                    [ASBR2-mpls] quit
                    [ASBR2] mpls ldp
                    [ASBR2-mpls-ldp] quit
                    [ASBR2] interface 10ge 1/0/2
                    [ASBR2-10GE1/0/2] mpls
                    [*ASBR2-10GE1/0/2] mpls ldp
                    [ASBR2-10GE1/0/2] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 4.4.4.4
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] mpls


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   503
VPN Configuration
VPN Configuration                                                                     5 VPWS Configuration

                    [PE2-10GE1/0/1] mpls ldp
                    [PE2-10GE1/0/1] quit

         Step 4 Configure an LDP VPWS connection.
                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] mpls l2vc 2.2.2.2 100
                    [PE1-10GE1/0/1] quit

                    # Configure ASBR1.
                    [ASBR1] mpls l2vpn
                    [ASBR1-l2vpn] quit
                    [ASBR1] interface 10ge 1/0/2
                    [ASBR1-10GE1/0/2] mpls l2vc 1.1.1.1 100
                    [ASBR1-10GE1/0/2] quit

                    # Configure ASBR2.
                    [ASBR2] mpls l2vpn
                    [ASBR2-l2vpn] quit
                    [ASBR2] interface 10ge 1/0/1
                    [ASBR2-10GE1/0/1] mpls l2vc 4.4.4.4 100
                    [ASBR2-10GE1/0/1] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] interface 10ge 1/0/2
                    [PE2-10GE1/0/2] mpls l2vc 3.3.3.3 100
                    [PE2-10GE1/0/2] quit

                    ----End

