---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-65
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [8554, 8727]
sha256: 7275ab4fe3b7dab68af613e3d9caf9f05a81fd48518cba44205f0b7f47d1de17
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         peer 2.2.2.2 as-number 100
                         peer 2.2.2.2 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 1.1.1.1 enable
                          peer 2.2.2.2 enable
                        #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 1.1.1.1 enable
                          peer 2.2.2.2 enable
                        #
                         ipv4-family vpn-instance vpn1
                          bestroute as-path-ignore
                          auto-frr
                          route-select delay 300
                          peer 192.168.2.2 as-number 65410
                          peer 192.168.2.2 preferred-value 600
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.20.1.0 0.0.0.3
                          network 10.11.1.0 0.0.0.3
                          network 3.3.3.3 0.0.0.0
                        #
                        return
                    ●   CE
                        #
                        sysname CE
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 192.168.1.2 255.255.255.252
                        #
                        interface Vlanif200
                         ip address 192.168.2.2 255.255.255.252
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 22.22.22.22 255.255.255.255
                        #
                        bgp 65410
                         peer 192.168.2.1 as-number 100
                         #
                         ipv4-family unicast
                          network 22.22.22.22 255.255.255.255
                          peer 192.168.2.1 enable
                        #
                        ospf 1
                         area 0.0.0.1
                          network 22.22.22.22 0.0.0.0
                          network 192.168.1.0 0.0.0.3
                        #
                        return



3.11 Configuring IPv4 L3VPN over MPLS Inter-AS
Option A

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         135
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration


3.11.1 Understanding IPv4 L3VPN over MPLS Inter-AS Option
A

Inter-AS VPN Option A Overview
                    As a basic IPv4 L3VPN application in the inter-AS scenario, Option A does not
                    need special configurations and MPLS does not need to run between ASBRs. In
                    this mode, ASBRs of two ASs directly connect to each other and function as PEs in
                    the ASs. Each ASBR views the peer ASBR as its CE, creates a VPN instance for each
                    VPN, and advertises IPv4 routes to the peer ASBR through EBGP.
                    On the network shown in Figure 3-22, ASBR1 in AS 100 views ASBR2 in AS 200 as
                    a CE. Similarly, ASBR2 also views ASBR1 as a CE. Here, a VPN LSP indicates a
                    private network tunnel, and an LSP indicates a public network tunnel.

                    Figure 3-22 Inter-AS VPN Option A networking




Route Advertisement in an Inter-AS VPN Option A Scenario
                    MP-IBGP runs between PEs and ASBRs to exchange VPN-IPv4 route information. A
                    common PE-CE routing protocol (BGP or IGP multi-instance) or static route can be
                    used between ASBRs for the exchange of VPN information. Because this involves
                    interaction between different ASs, using EBGP is recommended.
                    For example, CE1 advertises route 10.1.1.1/24 to CE2. Figure 3-23 shows the
                    process. D indicates the destination address, NH the next hop, and L1 and L2 the
                    VPN labels. This figure does not show the distribution of public network IGP
                    routes and labels.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         136
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration


                    Figure 3-23 Route advertisement in an inter-AS VPN Option A scenario




                    The specific process is as follows:

                    1.   CE1 uses BGP, OSPF, or RIP to advertise the route to PE1 in AS 100.
                    2.   PE1 in AS 100 uses MP-IBGP to advertise the labeled VPNv4 route to ASBR1 in
                         AS 100.
                    3.   After ASBR1 receives the VPNv4 route, if an export VPN target of the route
                         matches an import VPN target configured for the local VPN instance, ASBR1
                         adds the route to the VPN routing table, saves the VPN label carried in the
                         route, and advertises the unicast route to ASBR2 through the VPN instance
                         peer relationship.
                    4.   ASBR2 converts the received unicast route into a VPNv4 route and uses MP-
                         IBGP to advertise the route to PE2 in AS 200.
                    5.   PE2 in AS 200 uses BGP, OSPF, or RIP to advertise the route to CE2.



Packet Forwarding in an Inter-AS VPN Option A Scenario
                    Figure 3-24 shows packet forwarding over an LSP on the public network. L1 and
                    L2 indicate VPN labels, and Lx and Ly public network labels.


                    Figure 3-24 Packet forwarding in an inter-AS VPN Option A scenario




3.11.2 Configuring IPv4 L3VPN over MPLS Inter-AS Option A


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         137
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration


Prerequisites
                    Before configuring inter-AS VPN Option A, you have completed the following
                    tasks:
                    ●   Configure IGP for the MPLS backbone network in each AS to ensure IP
                        connectivity for the backbone network in each AS.
                    ●   Configure MPLS and MPLS LDP both globally and per interface on the PEs
                        and ASBRs.
                    ●   Establish a tunnel (LSP) between the PE and ASBR in the same AS.
                    ●   Configure IP addresses on interfaces that connect CEs to PEs.

Context
                    Inter-AS VPN Option A is a typical application of IPv4 L3VPN in an inter-AS
                    scenario and does not require special configurations. In this mode, an ASBR views
                    the peer ASBR as its CE and advertises IPv4 routes to the peer ASBR through EBGP.
                    On the network shown in 3.11.2 Configuring IPv4 L3VPN over MPLS Inter-AS
                    Option A, ASBR1 in AS 100 views ASBR2 in AS 200 as a CE. Similarly, ASBR2 also
                    views ASBR1 as a CE.

                    Figure 3-25 Inter-AS VPN Option A networking




