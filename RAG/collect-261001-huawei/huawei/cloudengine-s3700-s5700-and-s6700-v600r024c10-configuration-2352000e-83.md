---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-83
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [11479, 11612]
sha256: 767cb790052eb9c0ea45f165555acca67c8b641f3f49bd3fc7d8dc3dc6445b11
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


                    To deploy an IPv4 L3VPN on a hierarchical network, you need to use a hierarchical
                    model instead of a plane model. This is where the concept of HVPN comes in.

Related Concepts
                    Figure 3-33 shows a basic HVPN architecture consisting of mainly UPEs, SPEs, and
                    NPEs:
                    ●   UPE: directly connects to a user and is referred to as an underlayer PE or user-
                        end PE. A UPE mainly provides user access.
                    ●   SPE: a type of PE connected to UPEs and located at the core of a network. An
                        SPE is also called a superstratum PE or service provider-end PE. An SPE
                        manages and advertises VPN routes.
                    ●   NPE: connects to SPEs and is located at the network side. An NPE is known as
                        a network provider-end PE.
                    A UPE and an SPE are connected by only one link and exchange packets based on
                    labels. An SPE does not need to provide a large number of interfaces for access
                    users. UPEs and SPEs can be connected by physical interfaces with physical links,
                    by sub-interfaces with VLANs or PVCs, or by tunnel interfaces with LSPs. If an IP or
                    MPLS network resides between a UPE and an SPE, the UPE and SPE can be
                    connected by tunnel interfaces to exchange labeled packets over a tunnel.
                    The capabilities of SPEs and UPEs differ according to the roles they play on a
                    network. SPEs require large-capacity routing tables and high forwarding
                    performance, but few interface resources. UPEs, on the other hand, require only
                    low-capacity routing tables and low forwarding performance, but high access
                    capabilities.

                         NOTE

                        The roles of UPEs and SPEs are relative. On an HVPN, a superstratum PE is the SPE of an
                        understratum PE, and an understratum PE is the UPE of a superstratum PE.
                        An HoPE is compatible with common PEs on an MPLS network.

                    If a UPE and an SPE belong to the same AS, they use MP-IBGP. If they belong to
                    different ASs, they use MP-EBGP.
                    If MP-IBGP is used, an SPE can function as the RR for multiple UPEs to advertise
                    routes between IBGP peers. To reduce the number of routes on UPEs, ensure that
                    an SPE that is already acting as the RR for UPEs is not used as the RR for other
                    PEs.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    182
VPN Configuration
VPN Configuration                                                          3 IPv4 L3VPN Configuration


                    Figure 3-33 HVPN architecture




                    HVPN can be classified into HoVPN and H-VPN.

                    Table 3-4 Comparison of HoVPN and H-VPN
                     HVPN Mode           Characteristic

                     HoVPN               An SPE advertises only a default or summary route to a
                                         UPE.
                                         ● An export policy must be configured on an SPE so that
                                           the SPE advertises only specific routes, such as the
                                           default routes, to UPEs.
                                         ● VPN instances must be configured on an SPE for the
                                           SPE to import default routes locally or summarize
                                           routes received from remote sites, so that the SPE
                                           advertises only a default or summary route to a UPE.

                     H-VPN               An SPE advertises all VPN routes to UPEs.
                                         ● VPN instances do not need to be configured on SPEs.
                                         ● MP-BGP peer relationships must be configured between
                                           SPEs and NPEs and between SPEs and UPEs. The NPEs
                                           and UPEs must be configured as the clients of SPEs that
                                           function as RRs and be configured to set the next hops
                                           of routes they receive as themselves.




                    The following describes the route exchange and packet forwarding processes on
                    an HoVPN or H-VPN. In the following figures, N indicates a next hop, and L
                    indicates a label.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          183
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


Route Advertisement from CE1 to Device1 on an HoVPN or H-VPN
                    Figure 3-34 shows route advertisement from CE1 to Device1 on an HoVPN or H-
                    VPN.
                    1.   CE1 advertises an IPv4 route to the UPE using the IP protocol.
                    2.   The UPE applies for label L1 for the received IPv4 route and converts it into a
                         VPNv4 route. Then, the UPE sets itself as the next hop of the route and
                         advertises it to the SPE.
                    3.   After receiving the VPNv4 route, the SPE saves label L1 locally and applies for
                         label L2 for the route. Then, the SPE sets itself as the next hop of the route
                         and advertises it to the NPE.
                    4.   After receiving the VPNv4 route, the NPE converts it into an IPv4 route and
                         imports the route to its VPN IPv4 routing table if the route's next hop is
                         reachable. The NPE retains VPN label L2 and recursion tunnel ID information
                         of the route for later packet forwarding.
                    5.   The NPE then advertises the IPv4 route to Device1 using the IP protocol.


                    Figure 3-34 Route advertisement from CE1 to Device1 on an HoVPN or H-VPN




Route Advertisement from Device1 to CE1 on an HoVPN
                    Figure 3-35 shows route advertisement from Device1 to CE1 on an HoVPN.
                    1.   Device1 advertises an IPv4 route to the NPE using the IP protocol.
                    2.   The NPE applies for label L3 for the received IPv4 route and converts it into a
                         VPNv4 route. Then, the NPE sets itself as the next hop of the route and
                         advertises it to the SPE.
                    3.   After receiving the VPNv4 route, the SPE saves label L3 locally, converts the
                         route into an IPv4 route, and imports the route into its VPN IPv4 routing table
                         if the route's next hop is reachable.
                    4.   The SPE imports a default route into its VPN IPv4 routing table or generates a
                         summary VPN route based on the received IPv4 route in its VPN IPv4 routing
                         table and applies for label L4 for the default route or summary VPN route.
                         Then, the SPE converts the default route or summary VPN route into a VPNv4
                         route, sets itself as the next hop of the route, and advertises the route to the
                         UPE.


