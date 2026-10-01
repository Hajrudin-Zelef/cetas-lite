---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-94
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13043, 13225]
sha256: d37d8f9126068a921c7378919b776a6d56b4f2eaf8b7a18294077befbb070407
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 172.19.6.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 172.20.6.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip binding vpn-instance vpna
                         ip address 10.2.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 6.6.6.6 255.255.255.255
                        #
                        interface LoopBack2
                         ip binding vpn-instance vpna
                         ip address 66.66.66.66 255.255.255.255
                        #
                        bgp 100
                         peer 3.3.3.3 as-number 100
                         peer 3.3.3.3 connect-interface LoopBack1
                         peer 4.4.4.4 as-number 100
                         peer 4.4.4.4 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.3 enable
                          peer 4.4.4.4 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.3 enable
                          peer 3.3.3.3 route-policy SPE1 import
                          peer 4.4.4.4 enable
                          peer 4.4.4.4 route-policy SPE2 import
                         #
                         ipv4-family vpn-instance vpna
                          import-route direct
                          auto-frr
                          route-select delay 300
                          peer 10.2.1.2 as-number 65420
                        #
                        ospf 1
                         area 0.0.0.0
                          network 6.6.6.6 0.0.0.0
                          network 172.19.6.0 0.0.0.255
                          network 172.20.6.0 0.0.0.255
                        #
                        route-policy SPE1 permit node 10
                         apply local-preference 180
                        #
                        route-policy SPE2 permit node 10



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         208
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                         apply local-preference 170
                        #
                        return

                    ●   CE
                        #
                        sysname CE
                        #
                        vlan batch 100 200 300
                        #
                        interface Vlanif100
                         ip address 10.4.1.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 10.2.1.2 255.255.255.0
                        #
                        interface Vlanif300
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
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 7.7.7.7 255.255.255.255
                        #
                        bgp 65420
                         peer 10.4.1.1 as-number 100
                         peer 10.2.1.1 as-number 100
                         #
                         ipv4-family unicast
                          import-route direct
                          peer 10.4.1.1 enable
                          peer 10.2.1.1 enable
                        #
                        return


3.14.5 Example for Configuring H-VPN on an IP RAN
Networking Requirements
                    HVPN technology enables an IP RAN to provide fixed-mobile convergence (FMC)
                    and hierarchize the network between CSGs and RSGs. An IP RAN using HVPN
                    technology is easy to expand and flexible to deploy. HVPN has two networking
                    modes: HoVPN and H-VPN. The following example uses H-VPN, in which UPEs
                    store specific routes to SPEs. In H-VPN networking, SPEs function as RRs and UPEs
                    function as RR clients.
                    On the network shown in Figure 3-42, base stations connect to UPEs over a VPN.
                    H-VPN is deployed for base stations and base station controllers (EPC side) to
                    communicate. The link along UPE1 -> SPE1 -> NPE1 is the primary link. SPE2 is
                    the standby device for SPE1, and NPE2 is the standby device for NPE1.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         209
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


                    Figure 3-42 H-VPN networking
                         NOTE

                        In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200,
                        and VLANIF 300, respectively.




                    Table 3-6 IP addresses of physical interfaces
                     Device                          Interface                       IP Address

                     UPE1                            Loopback 1                      1.1.1.1/32

                                                     VLANIF 100                      172.16.3.1/24

                                                     VLANIF 200                      172.16.2.1/24

                                                     VLANIF 300                      10.1.1.2/24

                     UPE2                            Loopback 1                      2.2.2.2/32

                                                     VLANIF 100                      172.17.4.1/24

                                                     VLANIF 200                      172.16.2.2/24

                     SPE1                            Loopback 1                      3.3.3.3/32

                                                     VLANIF 100                      172.16.3.2/24

                                                     VLANIF 200                      172.18.4.1/24

                                                     VLANIF 300                      172.18.5.1/24

                     SPE2                            Loopback 1                      4.4.4.4/32

                                                     VLANIF 100                      172.17.4.2/24

