---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-23
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [2421, 2527]
sha256: 3e17df71100623b8f80d1e0a2fc890b8357dbe46c980fa891996b52575104ccb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        geographically adjacent to one another. If two geographically separated IP
                        systems are connected over a private line instead of an SP network, the two
                        systems compose a site.
                        As shown in Figure 3-2, if the branch network in City B connects to the HQ
                        network in City A over a private line instead of the SP network, the branch
                        network and HQ network form a site.
                        Devices at a site can belong to multiple VPNs. In other words, a site can
                        belong to more than one VPN.
                        As shown in Figure 3-3, in Company X, the decision-making department in
                        City A (Site A) needs to be configured to communicate with the R&D
                        department in City B (Site B) and the financial department in City C (Site C).
                        Site B and Site C, however, are not allowed to communicate. In this scenario,
                        two VPNs can be created: VPN 1 for sites A and B and VPN 2 for sites A and
                        C. Site A is configured to belong to multiple VPNs.

                        Figure 3-3 One site belonging to more than one VPN




                        A site connects to an SP network through a CE and may contain multiple CEs,
                        but one CE can belong to only one site. Determine the devices to be used as
                        CEs based on site conditions:
                        –   If the site contains a single host, use the host as the CE.
                        –   If the site contains a subnet, use switches as CEs.
                        –   If the site contains multiple subnets, use high-performance devices as
                            CEs.
                        Sites connecting to the same SP network can be categorized into different
                        sets based on configured policies. Sites can access one another only when
                        they belong to the same set. Such a set is a VPN.
                    ●   Address Space Overlapping
                        As a private network, a VPN independently manages an address space.
                        Address spaces of different VPNs may overlap. For example, address space
                        overlapping occurs if both VPN 1 and VPN 2 use addresses on network
                        segment 10.110.10.0/24.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             35
VPN Configuration
VPN Configuration                                                                 3 IPv4 L3VPN Configuration


                             NOTE

                            Two VPNs can use overlapped address spaces in the following situations:
                            ● The two VPNs do not cover the same site.
                            ● The two VPNs cover the same site, but the devices at such a site do not
                              communicate with the devices using overlapped address spaces in the VPNs.
                    ●   VPN Instance
                        CEs are user-side devices and need to send only local VPN routes to PEs,
                        irrespective of whether the PEs connect to the public network or other VPNs.
                        Generally, a PE of an SP network connects to multiple CEs of different VPNs.
                        Therefore, the PE receives routes from different VPNs. Because address spaces
                        used by these VPNs may overlap, routes sent from these VPNs may carry the
                        same destination address. If a PE maintains only one routing and forwarding
                        table, this table will retain only one of the routes from different VPNs and
                        with the same destination address, causing route loss. This is where VPN
                        instances come into play.
                        A VPN instance is also called a VPN routing and forwarding (VRF) table. A PE
                        has multiple routing and forwarding tables, including a public routing and
                        forwarding table and one or more VRF tables. In other words, a PE has
                        multiple instances, including a public network instance and one or more VPN
                        instances, as shown in Figure 3-4. Each VPN instance maintains routes from
                        the corresponding VPN, and the public network instance maintains public
                        network routes. This enables a PE to keep all routes from VPNs, preventing
                        route loss stemming from destination address overlapping.

                        Figure 3-4 VPN instances




                        The differences between a public routing and forwarding table and a VRF
                        table are as follows:
                        –   A public routing table contains the IPv4 routes advertised by all backbone
                            network devices. These IPv4 routes include configured static routes or
                            routes generated by routing protocols on the backbone network.
                        –   A VPN routing table contains routes of all sites in a VPN instance, and
                            these routes are obtained through the exchange of VPN routing
                            information between CEs and PEs, or between PEs.
                        –   Based on route management policies, a public forwarding table contains
                            the minimum forwarding information extracted from the corresponding

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                 36
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration


                            public routing table, whereas a VPN forwarding table contains the
                            minimum forwarding information extracted from the corresponding VPN
                            routing table.
                            The VPN instances on a PE are independent of each other and are
                            independent of the public routing and forwarding table.
                            Each VPN instance can be considered as a virtual device, which maintains
                            a separate address space and has one or more interfaces connected to
                            the device.
                            In related standards, VPN instances are called per-site forwarding tables.
                            As the name implies, a VPN instance corresponds to sites. To be specific,
                            every connection between a CE and PE corresponds to a VPN instance,
                            but this is not a one-to-one mapping. The VPN instance is manually
                            bound to the PE interface that directly connects to the CE.
                            A VPN instance uses a route distinguisher (RD) to identify an
                            independent address space and uses VPN targets to manage VPN
                            memberships and routing rules of directly connected sites and remote
                            sites.

