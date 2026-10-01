---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-24
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [2528, 2625]
sha256: 292436671ce80767a3bf270d419a8ad74689d67d94fe84a5c1950db032361f3b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   Relationships Between the VPN, Site, and VPN Instance
                        The relationships between the VPN, site, and VPN instance are as follows:
                        –   A VPN consists of multiple sites. A site may belong to multiple VPNs.
                        –   Each site is associated with a VPN instance on a PE. A VPN instance is a
                            set of VPN members and routing rules of associated sites. Multiple sites
                            compose a VPN based on the rules of the VPN instance.
                    ●   RD and VPN-IPv4 Address
                        Conventional BGP cannot process the routes of VPNs with overlapped address
                        spaces. Assume that VPN 1 and VPN 2 use addresses on network segment
                        10.110.10.0/24, and each of them advertises a route destined for this network
                        segment. The local PE identifies the two VPN routes based on VPN instances
                        and sends them to the remote PE. Because routes from different VPNs cannot
                        work in load-balancing mode, the remote PE adds only one of the two routes
                        to its VRF table based on BGP route selection rules.
                        This is because BGP cannot distinguish VPN routes with the same IP address
                        prefix. To solve this problem, IPv4 L3VPN uses the VPN-IPv4 address family.
                        A VPN-IPv4 address consists of 12 bytes, where the left-most 8 bytes
                        represent an RD and the right-most 4 bytes represent an IPv4 address prefix.
                        Figure 3-5 shows the format of a VPN-IPv4 address.

                        Figure 3-5 VPN-IPv4 address format




                        RDs are used to distinguish IPv4 prefixes using the same address space. The
                        format of RDs enables SPs to allocate RDs independently. An RD must be
                        unique on the entire network to ensure correct routing in CE dual-homing
                        scenarios. IPv4 addresses with RDs are called VPN-IPv4 addresses. After

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              37
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


                        receiving IPv4 routes from a CE, a PE converts the routes to globally unique
                        VPN-IPv4 routes and advertises these routes over the public network.
                    ●   VPN Target
                        A VPN uses 64-bit BGP extended community attributes, namely, VPN targets
                        (also called route targets), to control the advertisement of VPN routes.
                        A VPN instance is associated with one or more VPN targets, which are of the
                        following types:
                        –   Export target (ERT): After learning an IPv4 route from a directly
                            connected site, a PE converts the route to a VPN-IPv4 route, sets the ERT
                            for the route, and advertises the route carrying the ERT attribute as a
                            BGP extended community attribute.
                        –   Import target (IRT): After receiving a VPN-IPv4 route distributed by
                            another PE, the PE checks the ERT attribute of the route. If the ERT is
                            identical with the IRT of a VPN instance on the PE, the PE adds the route
                            to the routing table of the VPN instance.
                        That is, the VPN target attribute defines the sites to which routes can be
                        advertised and the sites whose routes can be accepted.
                        After receiving a route from a directly connected CE, a PE adds one or more
                        ERTs to the route. The PE then uses BGP to advertise the route with ERTs to
                        related PEs. After receiving the route, the related PEs compare the received
                        ERTs with the local IRTs of all their VPN instances. If an ERT is identical with
                        an IRT, the route is installed into the routing table of the corresponding VPN
                        instance.
                        The reasons for using VPN targets instead of RDs as BGP extended
                        community attributes are as follows:
                        –   A VPN-IPv4 route has only one RD, but can be associated with multiple
                            VPN targets. With multiple extended community attributes, BGP can
                            greatly improve network flexibility and expansibility.
                        –   VPN targets are used to control route advertisement between different
                            VPNs on a PE using matching VPN targets.
                        On a PE, different VPNs have different RDs, but the extended community
                        attributes allowed by BGP are limited. Using RDs for route import limits
                        network scalability.
                        On a VPN, VPN targets are used to control the import and export of VPN
                        routes between sites. ERTs and IRTs are independent of each other and
                        multiple ERTs and IRTs can be set for one VPN instance, helping implement
                        flexible VPN access control and diversified VPN networking modes.
                    ●   MP-BGP
                        Traditional BGP-4 standards define how to manage only IPv4 routing
                        information, and do not define how to process VPN routes with overlapped
                        address spaces.
                        To correctly process VPN routes, VPNs use MP-BGP defined in the related
                        standard (Multiprotocol Extensions for BGP-4). MP-BGP supports multiple
                        network layer protocols. Network layer protocol information is contained in
                        the Network Layer Reachability Information (NLRI) field and the Next Hop
                        field of an MP-BGP Update message.
                        MP-BGP uses address families to differentiate network layer protocols. An
                        address family can be a traditional IPv4 address family or any other address
                        family, such as a VPN-IPv4 address family or an IPv6 address family.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              38
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration


3.2.2 VPN NSR
                    VPN non-stop routing (NSR) is a technique which ensures that if an active main
                    board (AMB) failure causes the control plane of a specific node to fail, the node
                    retains its control plane connections to neighboring nodes and ensures
                    uninterrupted traffic transmission on the forwarding plane.

