---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-159
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [22825, 22979]
sha256: 77cf350934adb9d1b19ea1fbc22443ee72546d25e3c1602fde18f6b780814a94
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    Some Layer 3 features, such as route exchange between the PE and CE, can be
                    configured only after an IPv6 address is configured for the VPN interface on the
                    PE.

                    ----End

4.11.3 (Optional) Configuring a Router ID for a BGP VPN
Instance

Context
                    No router ID is configured for a BGP VPN instance by default, so the BGP router ID
                    is used. As such, different BGP VPN instances on the same device have the same
                    router ID. In some cases, different router IDs need to be configured for different
                    BGP VPN instances. For example, BGP peer relationships need to be established
                    between different BGP VPN instances on the same PE.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Enter the BGP VPN instance view.
                    vpn-instance vpn-instance-name

         Step 4 Configure a router ID or enable the device to automatically select a router ID.
                    router-id { ipv4-address | auto-select }

                    ----End

4.11.4 Configuring MP-IBGP Between PEs

Context
                    If VPN sites in a basic IPv6 L3VPN need to communicate, PEs must use MP-IBGP to
                    advertise VPNv6 routes carrying the RD attribute to each other. Because all the
                    PEs reside in the same AS, MP-IBGP peer relationships can be set up between
                    them. In the current implementation, IPv4 BGP peer relationships are set up
                    between PEs.

                    Perform the following steps on each PE.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    360
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Configure the remote PE as a peer.
                    peer ipv4-address as-number as-number

         Step 4 Specify an interface for setting up a TCP connection with the BGP peer.
                    peer ipv4-address connect-interface loopback interface-number

                          NOTE

                        A PE must use a loopback interface address with a 32-bit mask to set up an MP-IBGP peer
                        relationship with the peer PE so that VPN routes can recurse to tunnels. The route to the
                        local loopback interface is advertised to the peer PE using IGP on the MPLS backbone
                        network.

         Step 5 Enter the BGP-VPNv6 address family view.
                    ipv6-family vpnv6

         Step 6 Enable the function to exchange VPN-IPv6 routes with the peer.
                    peer ipv4-address enable

                    ----End

4.11.5 Configure IBGP VPN Peer Relationships Between PEs
and CEs (IP Forwarding)
Context
                    In the 6VPE scenario, a routing protocol must be configured between each pair of
                    a PE and a CE to allow them to communicate and to allow the CE to obtain routes
                    from other CEs, and IPv4 as well as IPv6 addresses need to be configured for the
                    interfaces on the link between the PE and CE.
                    The routing protocol configurations on CEs and PEs are different:
                    ●    A CE is located on the client side and unaware of the VPN. This means that
                         you do not need to configure VPN parameters when configuring a routing
                         protocol on the CE.
                    ●    A PE is located at the carrier network edge and connects to CEs for route
                         exchange. If the CEs connecting to a PE belong to different VPNs, the PE must
                         maintain different VRF tables. When configuring a routing protocol on the PE,
                         you need to specify the name of the VPN instance to which the routing
                         protocol applies and configure the routing protocol and MP-BGP to import
                         routes from each other.

Procedure
                    ●    Perform the following steps on the PE.
                         a.   Enter the system view.
                              system-view


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                 361
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


                        b.   Create a route-policy with a node and enter the route-policy view.
                             route-policy route-policy-name { permit | deny } node node
                        c.   Change the next hop of the route.
                             apply ipv6 next-hop address
                        d.   Return to the system view.
                             quit
                        e.   Enter the BGP view.
                             bgp as-number
                        f.   Create a BGP VPN instance and enter the BGP VPN instance view.
                             vpn-instance vpn-instance-name
                        g.   Configure the CE as a VPN peer.
                             peer ipv4-address as-number as-number
                        h.   (Optional) Specify the source interface and source address used to
                             establish a TCP connection with the BGP peer.
                             peer ipv4-address connect-interface interface-type interface-number [ ipv4-source-address ]
                        i.   Return to the BGP view.
                             quit
                        j.   Enter the BGP VPN instance IPv6 address family view.
                             ipv6-family vpn-instance vpn-instance-name
                        k.   Enable the function of exchanging BGP routing information with a BGP
                             IPv4 peer.
                             peer ipv4-address enable
                        l.   Apply an export routing policy.
                             peer ipv4-address route-policy route-policy-name export
                    ●   Perform the following steps on the CE.
                        a.   Enter the system view.
                             system-view
                        b.   Create a route-policy with a node and enter the route-policy view.
                             route-policy route-policy-name { permit | deny } node node
                        c.   Change the next hop of the route.
                             apply ipv6 next-hop address
                        d.   Return to the system view.
                             quit
                        e.   Enter the BGP view.
                             bgp as-number
                        f.   Configure a BGP peer.
                             peer ipv4-address as-number as-number
                        g.   (Optional) Specify the source interface and source address used to
                             establish a TCP connection with the BGP peer.
                             peer ipv4-address connect-interface interface-type interface-number [ ipv4-source-address ]
                        h.   Return to the BGP view.
                             quit
                        i.   Enter the BGP-IPv6 unicast address family view.
                             ipv6-family unicast
                        j.   Enable the function of exchanging BGP routing information with a BGP
                             IPv4 peer.
                             peer ipv4-address enable


