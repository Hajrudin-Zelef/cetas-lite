---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-45
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5476, 5625]
sha256: b36f9b4b81586fafe81ec806c9d2d9ab2c1bdbb7f9611cf9007f0cf4a68e230a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                   Rules for automatically selecting a router ID for a BGP VPN instance IPv4 address
                                   family are as follows:
                                   ● If the loopback interfaces configured with IP addresses are bound to the VPN
                                     instance enabled with the IPv4 address family, the largest IP address among
                                     them is selected as the router ID.
                                   ● If no loopback interfaces configured with IP addresses are bound to the VPN
                                     instance enabled with the IPv4 address family, the largest IP address among
                                     those of other interfaces bound to the VPN instance is selected as the router
                                     ID, regardless of whether the interface is up or down.
                    ●   Set a router ID for a specified BGP VPN instance IPv4 address family.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        d.    Manually set a router ID or enable the device to automatically select a
                              router ID for the BGP VPN instance IPv4 address family.
                              router-id { ipv4-address | auto-select }

                    ----End


3.9 Configuring the MCE Function

3.9.1 Understanding MCE
                    Multi-VPN-instance customer edge (MCE) technology provides logically
                    independent VPN instances and address spaces on a CE, allowing multiple VPN
                    users to share the same CE. The MCE technology provides an economical and
                    easy-to-use solution to solve problems concerned with VPN service isolation and
                    security.
                    VPN services are becoming increasingly refined and the demand for VPN service
                    security is growing. As such, carriers must isolate different types of VPN services
                    on networks to meet this demand. Deploying a separate CE for each VPN to
                    connect to upper-layer devices results in high cost and complex deployment.
                    Deploying the same CE for multiple VPNs to connect to upper-layer devices poses
                    data security risks, since these VPNs use the same routing and forwarding table.
                    To tackle issues regarding network cost and data security when multiple VPNs are
                    deployed, MCE technology can be used.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      86
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration


                    Figure 3-12 Networking diagram of using basic VPN technologies to isolate
                    services




                    The MCE technology creates a VPN instance for each VPN service to be isolated.
                    Each VPN uses an independent routing protocol to communicate with the MCE to
                    which these VPNs are connected. A VPN instance is bound to each link between
                    the MCE (interface or sub-interface) and the PE (interface or sub-interface) to
                    which the MCE is bound. As a result, a separate channel is established for each
                    VPN service, and different VPN services are isolated.

                    As shown in Figure 3-13, three VPN instances are configured on the MCE: VPN 1,
                    VPN 2, and VPN 3. Three separate VPN instance routing and forwarding tables are
                    created on the MCE accordingly. On the MCE, the interface connected to site 1 is
                    bound to VPN 1, the interface connected to site 2 is bound to VPN 2, and the
                    interface connected to site 3 is bound to VPN 3. In addition, the three interfaces
                    connected to the PE are bound to corresponding VPN instances. These
                    configurations allow VPN services to be isolated using only one MCE.


                    Figure 3-13 MCE networking




3.9.2 Configuring BGP Between an MCE and a PE

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           87
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


Prerequisites
                    Before configuring MCE, you have completed the following tasks:

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 3.5.1 Configuring an IPv4 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind each MCE interface and the PE interface connecting to the MCE to the
                        VPN instance, and configure IP addresses for the interfaces. For detailed
                        configurations, see 3.5.2 Binding an Interface to an IPv4 VPN Instance.

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP.


Procedure
                    ●   Configure the PE.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the BGP view.
                             bgp as-number

                        c.   Enter the BGP VPN instance IPv4 address family view.
                             ipv4-family vpn-instance vpn-instance-name

                        d.   Configure the MCE as a VPN BGP peer for the PE.
                             peer ipv4-address as-number as-number

                        e.   (Optional) Configure the maximum number of hops allowed for an EBGP
                             connection.
                             peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

                             This step is mandatory if the PE is not directly connected to the MCE but
                             an EBGP peer relationship needs to be established between them.

                             In most cases, a directly connected physical link must be available
                             between EBGP peers. If you want to establish an EBGP peer relationship
                             between indirectly connected peers, run the peer ebgp-max-hop
                             command to set the maximum number of hops allowed for a TCP
                             connection.

                             The default value of hop-count is 255. If the maximum number of hops is
                             set to 1, the MCE can establish an EBGP connection with only a directly
                             connected peer.
                    ●   Configure the MCE.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the BGP view.
                             bgp as-number


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    88
VPN Configuration
VPN Configuration                                                                        3 IPv4 L3VPN Configuration


