---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-95
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13226, 13370]
sha256: 6717af73b3582dc9476070465465c63ac6b2690d2796e9c966d8d68158454acb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                                     VLANIF 200                      172.18.4.2/24

                                                     VLANIF 300                      172.19.6.1/24

                     NPE1                            Loopback 1                      5.5.5.5/32

                                                     VLANIF 100                      172.18.5.2/24



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      210
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                     Device                      Interface                    IP Address

                                                 VLANIF 200                   172.20.6.1/24

                                                 VLANIF 300                   10.4.1.1/24

                     NPE2                        Loopback 1                   6.6.6.6/32

                                                 VLANIF 100                   172.19.6.2/24

                                                 VLANIF 200                   172.20.6.2/24

                                                 VLANIF 300                   10.2.1.1/24

                     CE                          Loopback 1                   7.7.7.7/32

                                                 VLANIF 100                   10.4.1.2/24

                                                 VLANIF 200                   10.2.1.2/24

                                                 VLANIF 300                   10.3.1.1/24




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.    Configure IGP for UPEs, SPEs, and NPEs to communicate and learn each
                          other's loopback addresses.
                    2.    Establish MPLS LSPs between UPEs and SPEs and between SPEs and NPEs.
                    3.    Configure a VPN instance on UPEs and NPEs, establish EBGP peer
                          relationships between NPEs and CEs, and import local direct routes to UPEs
                          and NPEs.
                    4.    Establish MP-IBGP peer relationships between UPEs and SPEs and between
                          SPEs and NPEs.
                    5.    Configure SPEs as RRs and specify UPEs and NPEs as RR clients.
                    6.    Configure a route-policy on each SPE and NPE to adjust the local preference
                          of the primary and backup routes. Configure VPN FRR on UPEs and NPEs and
                          VPNv4 FRR on SPEs to enhance network reliability.

Procedure
         Step 1 Configure OSPF for UPEs, SPEs, and NPEs to communicate.
                    After OSPF is configured, OSPF neighbor relationships can be established between
                    UPEs and SPEs and between SPEs and NPEs. Run the display ospf peer command.
                    The command output shows that the neighbor status is Full. Run the display ip
                    routing-table command. The command output shows that UPEs, SPEs, and NPEs
                    have learned the routes to each other's loopback interfaces.
                    For detailed configurations, see Configuration Scripts.
         Step 2 Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                MPLS backbone network.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          211
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                    After the configuration is complete, LDP sessions can be established between UPEs
                    and SPEs and between SPEs and NPEs. Run the display mpls ldp session
                    command. The command output shows that the session status is Operational.
                    Then, run the display mpls ldp lsp command. The command output shows that
                    LDP LSPs have been established.
                    For detailed configurations, see Configuration Scripts.
         Step 3 Configure a VPN instance on UPEs and NPEs, establish EBGP peer relationships
                between NPEs and CEs, and import local direct routes to UPEs and NPEs.
                    # Configure UPE1.
                    [UPE1] ip vpn-instance vpna
                    [UPE1-vpn-instance-vpna] ipv4-family
                    [UPE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [UPE1-vpn-instance-vpna-af-ipv4] vpn-target 1:1 both
                    [UPE1-vpn-instance-vpna-af-ipv4] quit
                    [UPE1-vpn-instance-vpna] quit
                    [UPE1] interface Vlanif300
                    [UPE1-Vlanif300] ip binding vpn-instance vpna
                    [UPE1-Vlanif300] ip address 10.1.1.2 24
                    [UPE1-Vlanif300] quit
                    [UPE1] bgp 100
                    [UPE1-bgp] ipv4-family vpn-instance vpna
                    [UPE1-bgp-vpna] peer 10.1.1.1 as-number 65410
                    [UPE1-bgp-vpna] import-route direct
                    [UPE1-bgp-vpna] quit
                    [UPE1-bgp] quit

                    # Configure NPE1.
                    [NPE1] ip vpn-instance vpna
                    [NPE1-vpn-instance-vpna] ipv4-family
                    [NPE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [NPE1-vpn-instance-vpna-af-ipv4] vpn-target 1:1 both
                    [NPE1-vpn-instance-vpna-af-ipv4] quit
                    [NPE1-vpn-instance-vpna] quit
                    [NPE1] interface Vlanif300
                    [NPE1-Vlanif300] ip binding vpn-instance vpna
                    [NPE1-Vlanif300] ip address 10.4.1.1 24
                    [NPE1-Vlanif300] quit
                    [NPE1] bgp 100
                    [NPE1-bgp] ipv4-family vpn-instance vpna
                    [NPE1-bgp-vpna] peer 10.4.1.2 as-number 65420
                    [NPE1-bgp-vpna] import-route direct
                    [NPE1-bgp-vpna] quit
                    [NPE1-bgp] quit

                    # Configure the CE.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE] interface Vlanif100
                    [CE-Vlanif100] ip address 10.4.1.2 24
                    [CE-Vlanif100] quit
                    [CE] interface Vlanif200
                    [CE-Vlanif200] ip address 10.2.1.2 24
                    [CE-Vlanif200] quit
                    [CE] interface Vlanif300
                    [CE-Vlanif300] ip address 10.3.1.1 24
                    [CE-Vlanif300] quit
                    [CE] interface loopback 1
                    [CE-Loopback1] ip address 7.7.7.7 32
                    [CE-Loopback1] quit
                    [CE] bgp 65420
                    [CE-bgp] peer 10.4.1.1 as-number 100
                    [CE-bgp] peer 10.2.1.1 as-number 100


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           212
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration

                    [CE-bgp] import-route direct
                    [CE-bgp] quit

