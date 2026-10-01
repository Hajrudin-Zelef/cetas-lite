---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-210
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [30846, 31021]
sha256: 31820b7ceae08887b29e1b15d6c1bad99464a34a1a3eb848524f6868a1c60ba8
---

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
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 10.10.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure ASBR1.
                    [ASBR1] ospf 1
                    [ASBR1-ospf-1] area 0.0.0.0
                    [ASBR1-ospf-1-area-0.0.0.0] network 2.2.2.9 0.0.0.0
                    [ASBR1-ospf-1-area-0.0.0.0] network 10.10.1.0 0.0.0.255
                    [ASBR1-ospf-1-area-0.0.0.0] quit
                    [ASBR1-ospf-1] quit

                    # Configure ASBR2.
                    [ASBR2] ospf 1
                    [ASBR2-ospf-1] area 0.0.0.0
                    [ASBR2-ospf-1-area-0.0.0.0] network 3.3.3.9 0.0.0.0
                    [ASBR2-ospf-1-area-0.0.0.0] network 10.20.1.0 0.0.0.255
                    [ASBR2-ospf-1-area-0.0.0.0] quit
                    [ASBR2-ospf-1] quit

                    # Configure PE2.
                    [PE2] ospf 1
                    [PE2-ospf-1] area 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] network 4.4.4.9 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] network 10.20.1.0 0.0.0.255
                    [PE2-ospf-1-area-0.0.0.0] quit
                    [PE2-ospf-1] quit

         Step 3 Enable MPLS and establish LSPs.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] mpls
                    [PE1-10GE1/0/2] mpls ldp
                    [PE1-10GE1/0/2] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   494
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration


                    # Configure ASBR1.
                    [ASBR1] mpls lsr-id 2.2.2.9
                    [ASBR1] mpls
                    [ASBR1-mpls] quit
                    [ASBR1] mpls ldp
                    [ASBR1-mpls-ldp] quit
                    [ASBR1] interface 10ge 1/0/1
                    [ASBR1-10GE1/0/1] mpls
                    [ASBR1-10GE1/0/1] mpls ldp
                    [ASBR1-10GE1/0/1] quit

                    # Configure ASBR2.
                    [ASBR2] mpls lsr-id 3.3.3.9
                    [ASBR2] mpls
                    [ASBR2-mpls] quit
                    [ASBR2] mpls ldp
                    [ASBR2-mpls-ldp] quit
                    [ASBR2] interface 10ge 1/0/2
                    [ASBR2-10GE1/0/2] mpls
                    [ASBR2-10GE1/0/2] mpls ldp
                    [ASBR2-10GE1/0/2] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 4.4.4.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] mpls
                    [PE2-10GE1/0/1] mpls ldp
                    [PE2-10GE1/0/1] quit

         Step 4 Configure BGP peers to exchange VPWS information.
                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.9 as-number 100
                    [PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [PE1-bgp] l2vpn-ad-family
                    [PE1-bgp-af-l2vpn-ad] peer 2.2.2.9 enable
                    [PE1-bgp-af-l2vpn-ad] peer 2.2.2.9 signaling vpws
                    [PE1-bgp-af-l2vpn-ad] quit
                    [PE1-bgp] quit

                    # Configure ASBR1.
                    [ASBR1] bgp 100
                    [ASBR1-bgp] peer 1.1.1.9 as-number 100
                    [ASBR1-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [ASBR1-bgp] l2vpn-ad-family
                    [ASBR1-bgp-af-l2vpn-ad] peer 1.1.1.9 enable
                    [ASBR1-bgp-af-l2vpn-ad] peer 1.1.1.9 signaling vpws
                    [ASBR1-bgp-af-l2vpn-ad] quit
                    [ASBR1-bgp] quit

                    # Configure ASBR2.
                    [ASBR2] bgp 200
                    [ASBR2-bgp] peer 4.4.4.9 as-number 200
                    [ASBR2-bgp] peer 4.4.4.9 connect-interface loopback 1
                    [ASBR2-bgp] l2vpn-ad-family
                    [ASBR2-bgp-af-l2vpn-ad] peer 4.4.4.9 enable
                    [ASBR2-bgp-af-l2vpn-ad] peer 4.4.4.9 signaling vpws
                    [ASBR2-bgp-af-l2vpn-ad] quit
                    [ASBR2-bgp] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   495
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration


                    # Configure PE2.
                    [PE2] bgp 200
                    [PE2-bgp] peer 3.3.3.9 as-number 200
                    [PE2-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [PE2-bgp] l2vpn-ad-family
                    [PE2-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                    [PE2-bgp-af-l2vpn-ad] peer 3.3.3.9 signaling vpws
                    [PE2-bgp-af-l2vpn-ad] quit
                    [PE2-bgp] quit

         Step 5 Configure a BGP connection.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] mpls l2vpn vpn1 encapsulation ethernet
                    [PE1-mpls-l2vpn-vpn1] route-distinguisher 100:1
                    [PE1-mpls-l2vpn-vpn1] vpn-target 1:1 both
                    [PE1-mpls-l2vpn-vpn1] ce ce1 id 1 range 10
                    [PE1-mpls-l2vpn-vpn1-ce-ce1] connection ce-offset 2 interface 10ge 1/0/1
                    [PE1-mpls-l2vpn-vpn1-ce-ce1] quit
                    [PE1-mpls-l2vpn-vpn1] quit

                    # Configure ASBR1.
                    [ASBR1] mpls l2vpn
                    [ASBR1-l2vpn] quit
                    [ASBR1] mpls l2vpn vpn1 encapsulation ethernet
                    [ASBR1-mpls-l2vpn-vpn1] route-distinguisher 100:1
                    [ASBR1-mpls-l2vpn-vpn1] vpn-target 1:1 both
                    [ASBR1-mpls-l2vpn-vpn1] ce ce2 id 2 range 10
                    [ASBR1-mpls-l2vpn-vpn1-ce-ce2] connection ce-offset 1 interface 10ge 1/0/2
                    [ASBR1-mpls-l2vpn-vpn1-ce-ce2] quit
                    [ASBR1-mpls-l2vpn-vpn1] quit

                    # Configure ASBR2.
                    [ASBR2] mpls l2vpn
                    [ASBR2-l2vpn] quit
                    [ASBR2] mpls l2vpn vpn1 encapsulation ethernet
                    [ASBR2-mpls-l2vpn-vpn1] route-distinguisher 200:1
                    [ASBR2-mpls-l2vpn-vpn1] vpn-target 1:1 both
                    [ASBR2-mpls-l2vpn-vpn1] ce ce3 id 3 range 10
                    [ASBR2-mpls-l2vpn-vpn1-ce-ce3] connection ce-offset 4 interface 10ge 1/0/1
                    [ASBR2-mpls-l2vpn-vpn1-ce-ce3] quit
                    [ASBR2-mpls-l2vpn-vpn1] quit

