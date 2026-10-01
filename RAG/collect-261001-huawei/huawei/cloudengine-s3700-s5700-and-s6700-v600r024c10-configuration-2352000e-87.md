---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-87
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [12020, 12167]
sha256: 4a74fa7e13f8149b1b438697d1276279b726b30dd759227f19834ba7f4371b67
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ----End

Verifying the Configuration
                    Run the display ip routing-table command to check the IP routing table.

3.14.4 Example for Configuring HoVPN on an IP RAN
Networking Requirements
                    HVPN technology enables an IP RAN to provide fixed-mobile convergence (FMC)
                    and hierarchize the network between CSGs and RSGs. An IP RAN using HVPN
                    technology is easy to expand and flexible to deploy. HVPN has two networking
                    modes: HoVPN and H-VPN. The following example uses HoVPN, in which UPEs
                    only need to store specific routes to base stations and default routes to SPEs.
                    HoVPN lowers the routing and forwarding requirements for UPEs.
                    On the network shown in Figure 3-41, base stations connect to UPEs over a VPN.
                    HoVPN is deployed for base stations and base station controllers (EPC side) to
                    communicate. The link along UPE1 -> SPE1 -> NPE1 is the primary link. SPE2 is
                    the standby device for SPE1, and NPE2 is the standby device for NPE1.

                    Figure 3-41 HoVPN networking
                          NOTE

                         In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200,
                         and VLANIF 300, respectively.


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    193
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration




                    Table 3-5 IP addresses of physical interfaces
                     Device                      Interface                  IP Address

                     UPE1                        Loopback 1                 1.1.1.1/32

                                                 VLANIF 100                 172.16.3.1/24

                                                 VLANIF 200                 172.16.2.1/24

                                                 VLANIF 300                 10.1.1.2/24

                     UPE2                        Loopback 1                 2.2.2.2/32

                                                 VLANIF 100                 172.17.4.1/24

                                                 VLANIF 200                 172.16.2.2/24

                     SPE1                        Loopback 1                 3.3.3.3/32

                                                 VLANIF 100                 172.16.3.2/24

                                                 VLANIF 200                 172.18.4.1/24

                                                 VLANIF 300                 172.18.5.1/24

                     SPE2                        Loopback 1                 4.4.4.4/32

                                                 VLANIF 100                 172.17.4.2/24

                                                 VLANIF 200                 172.18.4.2/24

                                                 VLANIF 300                 172.19.6.1/24

                     NPE1                        Loopback 1                 5.5.5.5/32

                                                 VLANIF 100                 172.18.5.2/24

                                                 VLANIF 200                 172.20.6.1/24

                                                 VLANIF 300                 10.4.1.1/24

                     NPE2                        Loopback 1                 6.6.6.6/32



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         194
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                     Device                        Interface                  IP Address

                                                   VLANIF 100                 172.19.6.2/24

                                                   VLANIF 200                 172.20.6.2/24

                                                   VLANIF 300                 10.2.1.1/24

                     CE                            Loopback 1                 7.7.7.7/32

                                                   VLANIF 100                 10.4.1.2/24

                                                   VLANIF 200                 10.2.1.2/24

                                                   VLANIF 300                 10.3.1.1/24




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
                    5.    Create a VPN instance on SPEs and specify UPEs connected to SPEs as the
                          understratum PEs of SPEs.
                    6.    Configure static default routes.
                    7.    Configure a route-policy on each SPE and NPE to adjust the local preference
                          of the primary and backup routes. Configure VPN FRR on UPEs, SPEs, and
                          NPEs to enhance network reliability.

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
                    After the configuration is complete, LDP sessions can be established between UPEs
                    and SPEs and between SPEs and NPEs. Run the display mpls ldp session

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         195
VPN Configuration
VPN Configuration                                                                3 IPv4 L3VPN Configuration


                    command. The command output shows that the session status is Operational.
                    Then, run the display mpls ldp lsp command. The command output shows that
                    LDP LSPs have been established.

                    For detailed configurations, see Configuration Scripts.

         Step 3 Configure a VPN instance on UPEs and NPEs, establish EBGP peer relationships
                between NPEs and CEs, and import local direct routes to UPEs and NPEs.

