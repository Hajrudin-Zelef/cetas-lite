---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-85
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "embedding"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [11742, 11873]
sha256: 7a1425b2d5c493831c6ffc184dfeb9fee8d33d09418b05c5e778324d4cf86e75
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Related Functions
                    H-VPN supports HoPE embedding:
                    ●   You can connect an HoPE as a UPE to an existing SPE for them to form a new
                        HoPE.
                    ●   Similarly, you can connect an HoPE as an SPE to multiple UPEs for them to
                        form a new HoPE.
                    ●   HoPEs can be embedded repeatedly in the preceding two situations.
                    HoPE embedding can infinitely expand a VPN in theory.
                    Figure 3-40 shows a three-layer HoPE, and the PEs in the middle are referred to
                    as middle-level PEs (MPEs). MP-BGP runs between the SPE and MPEs, and
                    between the MPEs and UPEs.

                         NOTE

                        The MPE concept is introduced solely for descriptive purposes and does not actually exist in
                        an H-VPN model.

                    MP-BGP advertises all the VPN routes of UPEs to the SPE, but advertises only the
                    default routes of the VPN instances of the SPE to UPEs.
                    An SPE maintains the routes of all VPN sites connected to its understratum PEs,
                    whereas a UPE maintains only the routes of its directly connected VPN sites. The
                    numbers of routes maintained by an SPE, an MPE, and a UPE are in descending
                    order.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                      188
VPN Configuration
VPN Configuration                                                            3 IPv4 L3VPN Configuration


                    Figure 3-40 HoPE embedding




Benefits
                    HVPN networking offers the following benefits:

                    ●   Flexible expansibility If the performance of a UPE is insufficient, you can add
                        an SPE for the UPE. If the access capabilities of an SPE are insufficient, you
                        can add more UPEs to the SPE.
                    ●   Less interface resource consumption. Since a UPE and an SPE exchange
                        packets based on labels, they only need to be connected over a single link.
                    ●   Reduced pressure on UPEs A UPE needs to maintain only local VPN routes.
                        The remote VPN routes are represented by a default or summary route,
                        reducing the pressure on UPEs.
                    ●   Less configuration workload SPEs and UPEs use MP-BGP, a dynamic routing
                        protocol, to exchange routes and advertise labels. Each UPE only needs to
                        establish one MP-BGP peer relationship with an SPE.

3.14.2 Configuring HoVPN

Context
                    On an HoVPN, a UPE only needs to obtain a default route from an SPE. This
                    implementation mechanism reduces the route storage space required on a UPE.

                    This scenario requires the following configurations:
                    ●   Create VPN instances on UPEs, SPEs, and NPEs. For configuration details, see
                        Configuring an IPv4 VPN Instance on a PE.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             189
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


                              NOTE

                             According to related standards, the VPN instance status obtained through an NMS can
                             be up only when at least one interface bound to the VPN instance is up. In an HoVPN
                             scenario, an SPE does not require any interface to be bound to a VPN instance.
                             Consequently, the VPN instance status obtained by an NMS is down according to the
                             standards, which is opposite to the actual VPN instance status. To solve this problem,
                             run the transit-vpn command in the VPN instance view or VPN instance IPv4 address
                             family view of an SPE. Then, the VPN instance status obtained from the NMS is always
                             up, regardless of whether any interface is bound to the VPN instance.
                    ●   Configure MP-BGP peer relationships between SPEs and NPEs. This
                        configuration is similar to configuring MP-IBGP peer relationships between
                        PEs in an IPv4 VPN instance. For details, see Establishing an MP-IBGP Peer
                        Relationship Between PEs.
                    ●   Configure routing protocols for NPEs/UPEs to exchange routes with CEs. This
                        configuration is similar to configuring PEs and CEs to exchange routes in an
                        IPv4 VPN instance. For details, see Configuring an IPv4 VPN Instance.
                    ●   Configure MP-BGP peer relationships between UPEs and SPEs. An SPE needs
                        to advertise only a default or summary route to a UPE. You can configure an
                        SPE to advertise a default route to a UPE in either of the following modes:
                        –    Route filtering mode: Configure a default static route and a route-policy
                             to enable the SPE to send the default route to the UPE.
                        –    Command control mode: Run the peer default-originate vpn-instance
                             command to enable the SPE to automatically generate a default route
                             and send it to the UPE.
                        The default route generated using the peer default-originate vpn-instance
                        command cannot be associated with an interface. When the primary link
                        connecting the SPE to UPE fails, the SPE may send the default route to the
                        UPE over another interface, which affects network response to faults.
                        Therefore, using the route filtering mode is recommended.


Procedure
                    ●   Configure a UPE to establish an MP-BGP peer relationship with an SPE.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the BGP view.
                             bgp as-number

                        c.   Specify the SPE as a BGP peer.
                             peer { ipv4-address | group-name } as-number as-number

                        d.   Enter the BGP-VPNv4 address family view.
                             ipv4-family vpnv4

                        e.   Enable the function to exchange BGP-VPNv4 routes with the specified
                             peer.
                             peer { ipv4-address | group-name } enable

                    ●   Configure the SPE to send a default or summary route to the UPE.
                        –    Configure the SPE to send a default route to the UPE.
                             i.   Enter the system view.
                                  system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    190
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration


