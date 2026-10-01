---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-212
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [31194, 31383]
sha256: c438f13f9a5ab4660542449e6b5f4fe27a2555dc79eb90eb0a2dc428ea7ff8f3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        mpls l2vpn vpn1 encapsulation ethernet
                         route-distinguisher 100:1
                         vpn-target 1:1 import-extcommunity
                         vpn-target 1:1 export-extcommunity
                         ce ce2 id 2 range 10 default-offset 0
                          connection ce-offset 1 interface 10GE1/0/2
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 1.1.1.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 1.1.1.9 signaling vpws
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.10.1.0 0.0.0.255
                        #
                        return
                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                        #
                        interface 10GE1/0/2
                         undo portswitch
                         ip address 10.20.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.0
                        #
                        mpls l2vpn vpn1 encapsulation ethernet
                         route-distinguisher 200:1
                         vpn-target 1:1 import-extcommunity
                         vpn-target 1:1 export-extcommunity
                         ce ce3 id 3 range 10 default-offset 0
                          connection ce-offset 4 interface 10GE1/0/1
                        #
                        bgp 200
                         peer 4.4.4.9 as-number 200
                         peer 4.4.4.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 4.4.4.9 enable
                         #
                         l2vpn-ad-family


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   499
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

                          policy vpn-target
                          peer 4.4.4.9 enable
                          peer 4.4.4.9 signaling vpws
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.20.1.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        mpls lsr-id 4.4.4.9
                        #
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
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.0
                        #
                        mpls l2vpn vpn1 encapsulation ethernet
                         route-distinguisher 200:1
                         vpn-target 1:1 import-extcommunity
                         vpn-target 1:1 export-extcommunity
                         ce ce4 id 4 range 10 default-offset 0
                          connection ce-offset 3 interface 10GE1/0/2
                        #
                        bgp 200
                         peer 3.3.3.9 as-number 200
                         peer 3.3.3.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 3.3.3.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 3.3.3.9 enable
                          peer 3.3.3.9 signaling vpws
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.20.1.0 0.0.0.255
                        #
                        return


5.9.6 Example for Configuring Inter-AS LDP VPWS Option A
Networking Requirements
                    In Figure 5-31, CE1 and CE2 access the backbone network through PE1 in AS100
                    and PE2 in AS200, respectively.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   500
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


                    Inter-AS LDP VPWS Option A needs to be configured for CE1 and CE2 to
                    communicate because the number of VPWS PWs is small and interfaces between
                    ASBRs need to be used as AC interfaces.

                    Figure 5-31 Network diagram of configuring inter-AS LDP VPWS Option A
                          NOTE

                         In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure an IGP on the backbone network to ensure IP connectivity within
                         each AS.
                    2.   Configure basic MPLS functions on the backbone network and establish a
                         dynamic LSP between the PE and ASBR in the same AS. If the PE and ASBR
                         are not directly connected, you also need to establish a remote LDP session
                         between them.
                    3.   Establish an LDP VPWS connection between the PE and ASBR in the same AS.

Procedure
         Step 1 Configure IP addresses for interfaces on each device.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] undo portswitch
                    [CE1-10GE1/0/1] ip address 10.1.1.1 24
                    [CE1-10GE1/0/1] quit

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] interface loopback1
                    [PE1-Loopback1] ip address 1.1.1.1 32


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    501

