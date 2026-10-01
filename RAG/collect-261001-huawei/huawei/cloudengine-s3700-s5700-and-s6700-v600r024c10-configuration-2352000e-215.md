---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-215
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [31669, 31865]
sha256: 14f5b7c3aa9fe70538108870b4cd27467dfc927129214edd142530defbf83251
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Scripts
                    ●   CE1
                        #
                        sysname CE1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.1.1.1 255.255.255.0
                        #
                        return
                    ●   PE1
                        #
                        sysname PE1
                        #
                        mpls lsr-id 1.1.1.1
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         mpls l2vc 2.2.2.2 100
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.10.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return
                    ●   ASBR1
                        #
                        sysname ASBR1
                        #
                        mpls lsr-id 2.2.2.2
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.10.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         mpls l2vc 1.1.1.1 100
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   506
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

                          network 2.2.2.2 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return
                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        mpls lsr-id 3.3.3.3
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         mpls l2vc 4.4.4.4 100
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.20.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.20.1.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        mpls lsr-id 4.4.4.4
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.20.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         mpls l2vc 3.3.3.3 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.20.1.0 0.0.0.255
                        #
                        return
                    ●   CE2
                        #
                        sysname CE2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   507
VPN Configuration
VPN Configuration                                                                     5 VPWS Configuration

                        #
                        interface 10GE1/0/1
                         undo portswitch
                         ip address 10.1.1.2 255.255.255.0
                        #
                        return




5.10 Configuring VPWS FRR

5.10.1 Understanding VPWS FRR
                    Widespread MPLS L2VPN adoption has raised reliability requirements for L2VPNs,
                    especially for L2VPNs that carry real-time services such as VoIP and IPTV.
                    VPWS fast reroute (FRR) uses redundant networking to improve MPLS L2VPN
                    reliability. When a PW or PE fails, VPWS FRR quickly switches traffic to a backup
                    link. This mechanism implements end-to-end fault detection on PWs and provides
                    PW backup, significantly improving link-layer reliability for MPLS L2VPNs.
                    VPWS FRR mainly applies to the following scenarios:
                    ●   Symmetric access of CEs to PEs: The two CEs on both ends of a VC
                        communicate over two paths, one as the primary path and the other the
                        backup path.
                    ●   Asymmetric access of CEs to PEs: The CE on one end of a VC is connected to a
                        PE over a high-reliability link, while the CE on the other end is dual-homed to
                        lower-reliability PEs. The two CEs can then communicate over two paths. The
                        higher-reliability path serves as the primary path, and the lower-reliability
                        path serves as the backup path.

Implementation
                    VPWS FRR is mainly used on a network where the CE on one end is single-homed
                    to a PE and the CE on the other end is dual-homed to PEs, as shown in Figure
                    5-32.

                    Figure 5-32 Asymmetric access of CEs to PEs




                    In this networking, PE1 and CE2 terminate fault notification. When the primary
                    link fails, PE1 detects the fault, triggers traffic switching, and does not notify the
                    fault to CE1. In this case, CE2 receives the fault notification from the PE on the
                    primary link and switches traffic to the backup link.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             508
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


