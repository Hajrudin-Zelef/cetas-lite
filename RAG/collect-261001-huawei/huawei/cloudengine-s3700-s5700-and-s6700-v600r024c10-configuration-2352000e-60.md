---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-60
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7804, 7961]
sha256: f5fe7826d63d2695abe251ff31a1204d5f3dc3455ebd59f6cee86bfa46261772
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 10.3.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        bgp 65410
                         peer 10.1.1.1 as-number 100
                         #
                         ipv4-family unicast
                          network 11.11.11.11 255.255.255.255
                          peer 10.1.1.1 enable
                        #
                        ospf 1
                         import-route bgp
                         area 0.0.0.0
                          network 10.3.1.0 0.0.0.255
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 10.4.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        return

                    ●   DeviceA
                        #
                        sysname DeviceA
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.3.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 10.4.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         124
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration

                          port trunk allow-pass vlan 200
                         #
                         interface LoopBack1
                          ip address 11.11.11.11 255.255.255.255
                         #
                         ospf 1
                          area 0.0.0.0
                           network 11.11.11.11 0.0.0.0
                           network 10.3.1.0 0.0.0.255
                         #
                         return


3.10.8 Example for Configuring IP+VPNv4 Hybrid FRR
Networking Requirements
                    A CE at a VPN site is dual-homed to two PEs, and a VPNv4 peer relationship is
                    established between the two PEs. To protect one of the links between the PEs and
                    CE, configure IP+VPNv4 hybrid FRR.
                    If one link fails, IP+VPNv4 hybrid FRR can quickly switch traffic destined for the CE
                    to the backup next hop (the other PE).
                    As shown in Figure 3-21, a CE is dual-homed to PE2 and PE3; an MPLS public
                    network tunnel and a VPNv4 peer relationship are established between PE2 and
                    PE3. OSPF is configured between PE2 and the CE and EBGP is configured between
                    PE3 and the CE to exchange routing information. PE3 learns from the CE a route
                    to Loopback 1 on the CE and advertises the route to its VPNv4 peer. PE2 then has
                    two BGP routes to Loopback 1 on the CE. One is learned using OSPF, and the
                    other is a VPNv4 route sent from PE3 through MP-IBGP. PE2 preferentially selects
                    the OSPF route sent from the CE because OSPF takes precedence over BGP. PE3
                    selects a route to the loopback interface on the CE from the routes sent from the
                    CE and PE2 in a similar manner.
                    It is required that:
                    ●    PE2 must be configured to use the VPNv4 route sent from PE3 as a backup
                         route for the OSPF route sent from the CE.
                    ●    PE3 must be configured to prefer the EBGP route received from the CE and
                         use the IBGP route learned from PE2 as a backup route.
                    Then, if the link between a PE and the CE fails, traffic can be switched to the other
                    PE for transmission.
                    To meet the requirements, enable VPN IP FRR on PE2 and BGP auto FRR on PE3 to
                    implement IP+VPNv4 hybrid FRR.

                    Figure 3-21 Configuring IP+VPNv4 hybrid FRR
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    125
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration




Precautions
                    In a VPN FRR scenario, traffic is switched back to the primary path after this path
                    recovers. Because the order in which nodes undergo IGP convergence differs,
                    packet loss may occur during the switchback. To resolve this problem, run the
                    route-select delay delay-value command to configure a route selection delay so
                    that traffic is switched back only after forwarding entries on the devices along the
                    primary path are updated. The delay specified using delay-value depends on
                    various factors, such as the number of routes on each device. Configure an
                    appropriate delay as needed.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure OSPF for PE1, PE2, and PE3 to communicate on the MPLS backbone
                         network.
                    2.   Configure basic MPLS capabilities and enable MPLS LDP to establish LDP LSPs
                         on the MPLS backbone network.
                    3.   Establish an MP-IBGP peer relationship between PEs, including between PE2
                         and PE3.
                    4.   Configure a VPN instance on each PE, and bind the interface connected to the
                         CE to the VPN instance on PE2 and PE3.
                    5.   Configure routing protocols between the PEs and CE.
                    6.   Configure VPN IP FRR on PE2 and BGP auto FRR on PE3.

