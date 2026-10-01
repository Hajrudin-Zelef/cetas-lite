---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-193
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [28180, 28364]
sha256: c053be022178a81122b167427eda0a9858bb50f834ba9f9cde430994af359a4d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

                          policy vpn-target
                          peer 3.3.3.9 enable
                          peer 3.3.3.9 signaling vpws
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         ip address 192.168.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif20
                         ip address 192.168.10.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                          network 192.168.10.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                        #
                        interface Vlanif20
                         ip address 192.168.10.2 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   450
VPN Configuration
VPN Configuration                                                                5 VPWS Configuration

                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        mpls l2vpn vpn1 encapsulation vlan
                         route-distinguisher 100:1
                         vpn-target 200:1 import-extcommunity
                         vpn-target 200:1 export-extcommunity
                         ce ce2 id 2 range 10 default-offset 0
                          connection ce-offset 1 interface Vlanif10
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
                          network 3.3.3.9 0.0.0.0
                          network 192.168.10.0 0.0.0.255
                        #
                        return



5.7 Configuring LDP VPWS

5.7.1 Understanding LDP VPWS
Definition
                    LDP VPWS is an MPLS L2VPN technology that establishes point-to-point links to
                    implement L2VPN and uses the LDP signaling to transmit VC information.
                    LDP VPWS uses double labels for traffic transmission. The inner label is exchanged
                    using extended LDP, and the outer label is a tunnel label.
                    On an LDP VPWS network, multiple VCs can be established over one LSP between
                    two PEs. Only PEs store a small amount of L2VPN information, such as mappings
                    between VC labels and LSPs. Ps do not store any L2VPN information. Therefore,
                    LDP VPWS has excellent scalability. To add a VC, you just need to configure a
                    unidirectional VC on each PE at both ends. This operation does not affect network
                    operations.
                    LDP VPWS uses LDP signaling. It does not have to wait for PW status changes to
                    periodically refresh, so LDP VPWS fault detection is fast.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        451
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


Basic Concepts
                    In LDP VPWS, the VC type and VC ID together uniquely identify a VC between two
                    CEs.

                    ●   The VC type indicates the encapsulation type of a VC, such as Ethernet or
                        VLAN.
                    ●   The VC ID identifies a VC. VCs of the same type must have unique IDs on a
                        PE.

                    The PEs connected to two CEs exchange VC labels using LDP and are bound to the
                    corresponding CEs based on the VC ID. A VC that transmits Layer 2 data can be
                    successfully established if the following conditions are all met:

                    ●   The physical status of AC interfaces is up.
                    ●   The tunnel between the PEs has been established.
                    ●   Labels have been exchanged and CE binding have been completed.

                    In LDP VPWS, the outer label is used to transmit the data of each VC over an ISP
                    network, and the inner VC label is used to identify service data. As such, an LSP on
                    the ISP network can be shared by multiple VCs.

                    To support LDP VPWS, an ISP network must be able to automatically establish
                    LSPs. This means that the ISP network must support MPLS forwarding and MPLS
                    LDP.


LDP VPWS Topology
                    Figure 5-19 shows the LDP VPWS topology.

                    Figure 5-19 LDP VPWS topology




                    In Figure 5-19, Site1 and Site2 of VPN1 are connected over a remote LDP
                    connection. Site1 and Site2 of VPN2 are also connected over a remote LDP
                    connection. Either one or two LSPs can be established within the ISP network for
                    communication between VPN1 and VPN2.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           452

