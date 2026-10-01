---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-137
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [19562, 19706]
sha256: 51a289d987cae22b3f3208ef99f1764f53df62e2ce3efd0c7a7059a0c27f70cd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                     Group Status          : Up           Oam Status : None
                     IPTN InLabel         :-            Tunnel BFD Status : -
                     BackUp LSP Type          : None         BestEffort : Disabled
                     Secondary HopLimit : -
                     BestEffort HopLimit : -
                     Secondary Explicit Path Name: -
                     Secondary Affinity Prop/Mask: 0x0/0x0
                     BestEffort Affinity Prop/Mask: 0x0/0x0
                     IsConfigLspConstraint: -
                     Hot-Standby Revertive Mode: Revertive
                     Hot-Standby Overlap-path: Disabled
                     Hot-Standby Switch State: CLEAR
                     Bit Error Detection: Disabled
                     Bit Error Detection Switch Threshold: -
                     Bit Error Detection Resume Threshold: -
                     Ip-Prefix Name : -
                     P2p-Template Name : -
                     PCE Delegate        : No            LSP Control Status : Local control
                     Path Verification : -
                     Entropy Label      : None
                     Associated Tunnel Group ID: -           Associated Tunnel Group Type: -
                     Auto BW Remain Time : -                 Reopt Remain Time      :-
                     Metric Inherit IGP : None
                     Binding Sid       :-             Reverse Binding Sid : -
                     Self-Ping       : Disable         Self-Ping Duration : 1800 sec
                     FRR Attr Source : -               Is FRR degrade down : -

                     Primary LSP ID        : 1.1.1.9:40
                     LSP State         : UP             LSP Type     : Primary
                     Configured Attribute Information:
                     Setup Priority     :7              Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Affinity Prop/Mask : 0x0/0x0           Resv Style : SE
                     Actual Attribute Information:
                     Setup Priority     :7              Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Metric Type         : TE           Metric      : 30
                     Configured Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 10000          CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0            CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0            CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0            CT7 Bandwidth(Kbit/sec): 0
                     Actual Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 10000          CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0            CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0            CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0            CT7 Bandwidth(Kbit/sec): 0
                     Explicit Path Name : -                       Hop Limit: -
                     Record Route          : Enabled       Record Label : Disabled
                     Route Pinning         : Disabled
                     FRR Flag          : Disabled
                     IdleTime Remain         :-
                     BFD Status          :-
                     Soft Preemption        : Disabled
                     Reroute Flag         : Enabled
                     Pce Flag         : Normal
                     Path Setup Type        : CSPF
                     Create Modify LSP Reason: -
                     Self-Ping Status : -

                 # Display TEDB information on LSR1.
                 [LSR1] display mpls te cspf tedb interface 10.1.1.1
                 Router ID: 1.1.1.9
                  IGP Type: ISIS    Process Id: 1
                  Link[1]:


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       329
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration

                     ISIS System ID: 0000.0000.0001.00       Opaque LSA ID: 0000.0000.0001.00:00
                     Interface IP Address: 10.1.1.1
                     DR Address: 10.1.1.1
                     DR ISIS System ID: 0000.0000.0001.01
                     IGP Area: Level-2
                     Link Type: Multi-access Link Status: Active
                     IGP Metric: 10    TE Metric: 10 Color: 0x0
                     Bandwidth Allocation Model : -
                     Maximum Link-Bandwidth: 100000 (kbps)
                     Maximum Reservable Bandwidth: 100000 (kbps)
                     Operational Mode of Router : TE
                     Bandwidth Constraints:         Local Overbooking Multiplier:
                       BC[0]:     100000 (kbps)          LOM[0]:       1
                     BW Unreserved:
                         Class ID:
                         [0]:    100000 (kbps),         [1]:   100000 (kbps)
                         [2]:    100000 (kbps),         [3]:   100000 (kbps)
                         [4]:    100000 (kbps),         [5]:   100000 (kbps)
                         [6]:    100000 (kbps),         [7]:   100000 (kbps)

                 ----End

Verifying the Configuration
                 # Change the tunnel bandwidth to 20000 kbit/s.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] mpls te bandwidth ct0 20000
                 [LSR1-Tunnel1] quit

                 # Display TEDB information on LSR1 again.
                 [LSR1] display mpls te cspf tedb interface 10.1.1.1
                 Router ID: 1.1.1.9
                  IGP Type: ISIS     Process Id: 1
                  Link[1]:
                   ISIS System ID: 0000.0000.0001.00       Opaque LSA ID: 0000.0000.0001.00:00
                   Interface IP Address: 10.1.1.1
                   DR Address: 10.1.1.1
                   DR ISIS System ID: 0000.0000.0001.01
                   IGP Area: Level-2
                   Link Type: Multi-access Link Status: Active
                   IGP Metric: 10    TE Metric: 10 Color: 0x0
                   Bandwidth Allocation Model : -
                   Maximum Link-Bandwidth: 100000 (kbps)
                   Maximum Reservable Bandwidth: 100000 (kbps)
                   Operational Mode of Router : TE
                   Bandwidth Constraints:         Local Overbooking Multiplier:
                     BC[0]:     100000 (kbps)          LOM[0]:       1
                   BW Unreserved:
                       Class ID:
                       [0]:    100000 (kbps),         [1]:   100000 (kbps)
                       [2]:    100000 (kbps),         [3]:   100000 (kbps)
                       [4]:    100000 (kbps),         [5]:   100000 (kbps)
                       [6]:    100000 (kbps),         [7]:    80000 (kbps)

                 The command output shows that the tunnel named Tunnel1 has been
                 reestablished successfully. Its bandwidth is 20000 kbit/s, reaching 20% (the
                 configured bandwidth flooding threshold). Therefore, CSPF TEDB information has
                 been updated.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       330
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

