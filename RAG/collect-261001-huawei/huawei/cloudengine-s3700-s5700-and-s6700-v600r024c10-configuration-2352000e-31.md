---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-31
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3459, 3576]
sha256: 7c73c0a254613dd7810a4bd9e2ba6ee30cf75e9dcfba1b1e0d1b424c8e5e47c0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         and PE. Routes advertised by the CE to the PE are standard IPv4 routes,
                         regardless of which routing protocol is used.
                         VPN instances on a PE are isolated from each other and independent of the
                         public routing and forwarding table, to prevent problems caused by VPN
                         address space overlapping. After learning routes from a CE, a PE needs to
                         determine to which routing and forwarding table the routes should be
                         installed. Common routing protocols, however, do not have this capability, and
                         manual configuration is required.
                    2.   Route advertisement from the ingress PE to the egress PE
                         Route advertisement from the ingress PE to the egress PE further consists of
                         the following phases:
                         –   After learning VPN routes from a CE, a PE stores these routes in the
                             corresponding VPN instances and adds RDs to these standard IPv4 routes.
                             VPN-IPv4 routes are then generated.
                         –   The ingress PE advertises VPN-IPv4 routes to the egress PE by sending
                             MP-BGP Update messages. The Update messages also contain VPN
                             targets and MPLS labels.
                         Before the next-hop PE receives the VPN-IPv4 routes, the routes are first
                         filtered by BGP route-policies, including the VRF export policy and peer export
                         policy.
                         After these routes arrive at the egress PE, if they match the BGP peer import
                         policy and their next hops are reachable or they can perform recursion, the
                         egress PE performs local route leaking and filters these routes based on a VRF
                         import policy. The egress PE then decides which routes are to be added to its
                         VPN routing tables. Routes received from other PEs are added to a VPN
                         routing table based on VPN targets. The egress PE stores the following
                         information for subsequent packet forwarding:
                         –   Values of MPLS labels contained in MP-BGP Update messages
                         –   Tunnel IDs generated after recursion to tunnels is successful
                    3.   Route advertisement from the egress PE to the remote CE
                         A remote CE can learn VPN routes from an egress PE over static routes or
                         routes established using RIP, OSPF, IS-IS, or BGP. Route advertisement from
                         the egress PE to a remote CE is similar to that from a local CE to the ingress
                         PE. The details are not described here. Note that the routes advertised by the
                         egress PE to the remote CE are standard IPv4 routes.
                    After a PE receives routes of different VPNs from a local CE, if the next hops of
                    these routes are reachable or these routes are recursive, the PE matches the
                    export VPN targets of these routes against the import VPN targets of its local VPN
                    instances. This process is called local route leaking. During local route leaking, the
                    PE filters these routes based on a VRF import policy and modifies the attributes of
                    eligible routes.

IPv4 L3VPN Packet Forwarding
                    On an IPv4 L3VPN backbone network, Ps are unaware of VPN routing information.
                    VPN packets are forwarded between PEs over tunnels. Figure 3-10 shows an
                    example of packet forwarding on an IPv4 L3VPN. Figure 3-10 shows packet
                    transmission from CE1 to CE2. I-L indicates an inner label, and O-L indicates an
                    outer label. The outer label directs the packet to the BGP next hop, and the inner

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              54
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


                    label identifies the outbound interface for the packet or the VPN to which the
                    packet belongs.

                    Figure 3-10 VPN packet forwarding




                    The specific packet forwarding process is as follows:
                    1.   CE1 sends an IP packet to the ingress PE.
                    2.   After receiving the packet from an interface bound to a VPN instance, the
                         ingress PE performs the following steps:
                         –   Searches the corresponding VPN forwarding table based on the RD of the
                             bound VPN instance.
                         –   Matches the destination IPv4 prefix against forwarding entries to search
                             for the corresponding tunnel ID.
                         –   Adds an I-L to the packet and finds the tunnel to be used based on the
                             tunnel ID.
                         –   Adds an outer label to the packet and sends the packet over the
                             corresponding tunnel. In this example, the tunnel is an LSP, and the outer
                             label is an MPLS label (O-L1).
                         –   Transmits the double-tagged packet over the backbone network. Each P
                             on the forwarding path swaps the outer label of the packet.
                    3.   After receiving the packet with double labels, the egress PE sends the packet
                         to MPLS for processing. MPLS removes the outer label.
                              NOTE

                             In this example, the final outer label of the packet is O-L2. If PHP is configured, O-L2
                             is removed on the penultimate hop, and the egress PE receives a packet with the inner
                             label only.
                    4.   The egress PE determines which VPN forwarding table to use based on the
                         inner label, searches this table for the outbound interface based on the IPv4
                         prefix, and then removes the inner label.
                    5.   The egress PE sends the packet through the corresponding outbound interface
                         to CE2. At this time, the packet is a native IP packet.
                    In this manner, the packet is sent from CE1 to CE2. CE2 forwards the packet to the
                    destination in the way it sends other IP packets.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        55
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


3.6.2 Configuring an IPv4 VPN Instance on a PE
Prerequisites
                    Before configuring an IPv4 VPN instance on a PE, you have completed the
                    following tasks:
                    ●    Configure link layer protocol parameters for interfaces to ensure that these
                         interfaces work properly.

Context
                    A VPN instance is also called a VRF table or a per-site forwarding table.
                    VPN instances are used to isolate VPN routes from public network routes. Routes
                    of different VPN instances are isolated from one another.

Procedure
         Step 1 Enter the system view.
                    system-view

