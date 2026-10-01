---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-144
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [20632, 20772]
sha256: 92cbfe86d0265d6ec4f5a90c66e1376ce2a8c51047f3f1a4c28519a22ccb3153
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Check the TEDB. You can view the bandwidth change of each link.
                 [LSR1] display mpls te cspf tedb node
                  Router ID: 1.1.1.1
                  IGP Type: OSPF      Process ID: 1     IGP Area: 0
                   MPLS-TE Link Count: 1
                   Link[1]:
                    OSPF Router ID: 10.1.1.1      Opaque LSA ID: 1.0.0.1
                    Interface IP Address: 10.1.1.1
                    DR Address: 10.1.1.2
                    IGP Area: 0
                    Link Type: Multi-access Link Status: Active
                    IGP Metric: 1           TE Metric: 1        Color: 0x10001
                    Bandwidth Allocation Model : -
                    Maximum Link-Bandwidth: 50000 (kbps)
                    Maximum Reservable Bandwidth: 50000 (kbps)
                    Operational Mode of Router: TE
                    Bandwidth Constraints:         Local Overbooking Multiplier:
                       BC[0]:      50000 (kbps)          LOM[0]:       1
                    BW Unreserved:
                       Class ID:
                       [0]:     50000 (kbps),          [1]:    50000 (kbps)
                       [2]:     50000 (kbps),          [3]:    50000 (kbps)
                       [4]:     50000 (kbps),          [5]:    50000 (kbps)
                       [6]:     50000 (kbps),          [7]:    10000 (kbps)
                  Router ID: 2.2.2.2
                  IGP Type: OSPF      Process ID: 1     IGP Area: 0
                   MPLS-TE Link Count: 3
                   Link[1]:
                    OSPF Router ID: 10.1.1.2      Opaque LSA ID: 1.0.0.1
                    Interface IP Address: 10.1.1.2
                    DR Address: 10.1.1.2
                    IGP Area: 0
                    Link Type: Multi-access Link Status: Active
                    IGP Metric: 1           TE Metric: 1        Color: 0x0
                    Bandwidth Allocation Model : -
                    Maximum Link-Bandwidth: 0 (kbps)
                    Maximum Reservable Bandwidth: 0 (kbps)
                    Operational Mode of Router: TE
                    Bandwidth Constraints:         Local Overbooking Multiplier:
                       BC[0]:        0 (kbps)         LOM[0]:       1
                    BW Unreserved:


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 345
MPLS Configuration
MPLS Configuration                                                                 4 MPLS TE Configuration

                       Class ID:
                       [0]:        0 (kbps),         [1]:       0 (kbps)
                       [2]:        0 (kbps),         [3]:       0 (kbps)
                       [4]:        0 (kbps),         [5]:       0 (kbps)
                       [6]:        0 (kbps),         [7]:       0 (kbps)
                   Link[2]:
                    OSPF Router ID: 10.1.1.2       Opaque LSA ID: 1.0.0.3
                    Interface IP Address: 10.1.2.1
                    DR Address: 10.1.2.1
                    IGP Area: 0
                    Link Type: Multi-access Link Status: Active
                    IGP Metric: 1            TE Metric: 1        Color: 0x10101
                    Bandwidth Allocation Model : -
                    Maximum Link-Bandwidth: 100000 (kbps)
                    Maximum Reservable Bandwidth: 100000 (kbps)
                    Operational Mode of Router: TE
                    Bandwidth Constraints:         Local Overbooking Multiplier:
                       BC[0]:     100000 (kbps)           LOM[0]:        1
                    BW Unreserved:
                       Class ID:
                       [0]:    100000 (kbps),           [1]:   100000 (kbps)
                       [2]:    100000 (kbps),           [3]:   100000 (kbps)
                       [4]:    100000 (kbps),           [5]:   100000 (kbps)
                       [6]:    100000 (kbps),           [7]:   60000 (kbps)
                   Link[3]:
                    OSPF Router ID: 10.1.1.2       Opaque LSA ID: 1.0.0.2
                    Interface IP Address: 10.1.3.1
                    DR Address: 10.1.3.1
                    IGP Area: 0
                    Link Type: Multi-access Link Status: Active
                    IGP Metric: 1            TE Metric: 1        Color: 0x10011
                    Bandwidth Allocation Model : -
                    Maximum Link-Bandwidth: 100000 (kbps)
                    Maximum Reservable Bandwidth: 100000 (kbps)
                    Operational Mode of Router: TE
                    Bandwidth Constraints:         Local Overbooking Multiplier:
                       BC[0]:     100000 (kbps)           LOM[0]:        1
                    BW Unreserved:
                       Class ID:
                       [0]:    100000 (kbps),           [1]:   100000 (kbps)
                       [2]:    100000 (kbps),           [3]:   100000 (kbps)
                       [4]:    100000 (kbps),           [5]:   100000 (kbps)
                       [6]:    100000 (kbps),           [7]:   100000 (kbps)
                  Router ID: 3.3.3.3
                  IGP Type: OSPF       Process ID: 1    IGP Area: 0
                   MPLS-TE Link Count: 2
                   Link[1]:
                    OSPF Router ID: 4.4.4.4      Opaque LSA ID: 1.0.0.2
                    Interface IP Address: 10.1.2.2
                    DR Address: 10.1.2.1
                    IGP Area: 0
                    Link Type: Multi-access Link Status: Active
                    IGP Metric: 1            TE Metric: 1        Color: 0x0
                    Bandwidth Allocation Model : -
                    Maximum Link-Bandwidth: 0 (kbps)
                    Maximum Reservable Bandwidth: 0 (kbps)
                    Operational Mode of Router: TE
                    Bandwidth Constraints:         Local Overbooking Multiplier:
                       BC[0]:         0 (kbps)        LOM[0]:        1
                    BW Unreserved:
                       Class ID:
                       [0]:        0 (kbps),         [1]:       0 (kbps)
                       [2]:        0 (kbps),         [3]:       0 (kbps)
                       [4]:        0 (kbps),         [5]:       0 (kbps)
                       [6]:        0 (kbps),         [7]:       0 (kbps)
                   Link[2]:
                    OSPF Router ID: 4.4.4.4      Opaque LSA ID: 1.0.0.1
                    Interface IP Address: 10.1.3.2
                    DR Address: 10.1.3.1



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            346
MPLS Configuration
MPLS Configuration                                                                                  4 MPLS TE Configuration

                     IGP Area: 0
                     Link Type: Multi-access Link Status: Active
                     IGP Metric: 1           TE Metric: 1       Color: 0x0
                     Bandwidth Allocation Model : -
                     Maximum Link-Bandwidth: 0 (kbps)
                     Maximum Reservable Bandwidth: 0 (kbps)
                     Operational Mode of Router: TE
                     Bandwidth Constraints:       Local Overbooking Multiplier:
                        BC[0]:        0 (kbps)        LOM[0]:       1
                     BW Unreserved:
                        Class ID:
                        [0]:       0 (kbps),        [1]:       0 (kbps)
                        [2]:       0 (kbps),        [3]:       0 (kbps)
                        [4]:       0 (kbps),        [5]:       0 (kbps)
                        [6]:       0 (kbps),        [7]:       0 (kbps)

