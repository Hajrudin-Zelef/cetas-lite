---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-191
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [27812, 27993]
sha256: e7df0b67d8d9f0982fb20e78c27e1fbc7cb6fe43db4139515dc0ed5cbc834414
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   PE
                        #
                        sysname PE
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        interface Vlanif10
                        #
                        interface Vlanif20
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        mpls l2vpn vpn1 encapsulation vlan
                         route-distinguisher 100:1
                         ce ce1 id 1 range 10 default-offset 0
                          connection ce-offset 2 interface Vlanif10
                         ce ce2 id 2 range 10 default-offset 0
                          connection ce-offset 1 interface Vlanif20
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   444
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


5.6.6 Example for Configuring a Remote BGP VPWS
Connection
Networking Requirements
                    In Figure 5-18, CE1 and CE2 are connected to different PEs. A remote BGP VPWS
                    connection needs to be established between PEs for the two CEs to communicate.

                    Figure 5-18 Network diagram of configuring a remote BGP VPWS connection
                          NOTE

                         In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure a routing protocol on the PEs and P of the backbone network to
                         ensure IP connectivity and configure basic MPLS functions and LDP.
                    2.   Configure BGP peers to exchange VPWS information.
                    3.   Enable MPLS L2VPN on PEs and create a remote BGP VPWS connection
                         between CE1 and CE2.
                               NOTE

                              By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device.
                              If a VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts
                              with LNP. In this case, run the lnp disable command in the system view to disable
                              LNP.


Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      445
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration

                    [CE1] vlan batch 20
                    [CE1] interface 10ge 1/0/2
                    [CE1-10GE1/0/2] port link-type trunk
                    [CE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE1-10GE1/0/2] quit
                    [CE1] interface vlanif 20
                    [CE1-Vlanif20] ip address 10.1.1.1 255.255.255.0
                    [CE1-Vlanif20] quit

                    The configuration of CE2 is similar to that of CE1. For detailed configurations, see
                    Configuration Scripts.

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] vlan batch 10 20
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [PE1-10GE1/0/1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [PE1-10GE1/0/2] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] ip address 192.168.1.1 255.255.255.0
                    [PE1-Vlanif10] quit
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.

         Step 2 Configure an IGP on the backbone network. In this example, OSPF is used.

                    # Configure PE1.
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 192.168.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure the P.
                    [P] ospf 1
                    [P-ospf-1] area 0
                    [P-ospf-1-area-0.0.0.0] network 192.168.1.0 0.0.0.255
                    [P-ospf-1-area-0.0.0.0] network 2.2.2.9 0.0.0.0
                    [P-ospf-1-area-0.0.0.0] network 192.168.10.0 0.0.0.255
                    [P-ospf-1-area-0.0.0.0] quit
                    [P-ospf-1] quit

                    # Configure PE2.
                    [PE2] ospf 1
                    [PE2-ospf-1] area 0
                    [PE2-ospf-1-area-0.0.0.0] network 192.168.10.0 0.0.0.255
                    [PE2-ospf-1-area-0.0.0.0] network 3.3.3.9 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] quit
                    [PE2-ospf-1] quit

         Step 3 Enable MPLS and establish LSPs.

                    # Configure PE1.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           446
VPN Configuration
VPN Configuration                                                               5 VPWS Configuration

                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif20] mpls
                    [PE1-Vlanif20] mpls ldp
                    [PE1-Vlanif20] quit

                    # Configure the P.
                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] quit
                    [P] mpls ldp
                    [P-mpls-ldp] quit
                    [P] interface vlanif 10
                    [P-Vlanif10] mpls
                    [P-Vlanif10] mpls ldp
                    [P-Vlanif10] quit
                    [P] interface vlanif 20
                    [P-Vlanif20] mpls
                    [P-Vlanif20] mpls ldp
                    [P-Vlanif20] quit

