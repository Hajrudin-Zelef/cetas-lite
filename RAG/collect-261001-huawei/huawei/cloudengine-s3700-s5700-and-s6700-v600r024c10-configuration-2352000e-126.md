---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-126
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [17821, 17969]
sha256: 2a47d31dd2950c4e7927123fdda78e16cd51977706cdbfb2bcbd54d67f79f2e4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Context
                    If VPN sites in a basic IPv6 L3VPN need to communicate, PEs must use MP-IBGP to
                    advertise VPNv6 routes carrying the RD attribute to each other. Because all the
                    PEs reside in the same AS, MP-IBGP peer relationships can be set up between
                    them. In the current implementation, IPv4 BGP peer relationships are set up
                    between PEs.
                    Perform the following steps on each PE.

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

4.6.4 Configuring BGP4+ Between the PE and CE

Context
                          NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP4+.


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                 283
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration


                    On an L3VPN, a routing protocol must be configured between PEs and CEs to
                    allow them to communicate and allow one CE to obtain routes to other CEs.
                    The routing protocol configurations on CEs and PEs are different:
                    ●   A CE is located on the client side and unaware of the VPN. This means that
                        you do not need to configure VPN parameters when configuring a routing
                        protocol on the CE.
                    ●   A PE is located at the carrier network edge and connects to CEs for route
                        exchange. If the CEs connecting to a PE belong to different VPNs, the PE must
                        maintain different VRF tables. When configuring a routing protocol on the PE,
                        you need to specify the name of the VPN instance to which the routing
                        protocol applies and configure the routing protocol and MP-BGP to import
                        routes from each other.

Procedure
                    ●   Perform the following steps on the PE:
                        a.   Enter the system view.
                             system-view

                        b.   Enter the BGP view.
                             bgp as-number

                        c.   Enter the BGP VPN instance IPv6 address family view.
                             ipv6-family vpn-instance vpn-instance-name

                        d.   Configure a CE as an IPv6 VPN peer.
                             peer ipv6-address as-number as-number

                        e.   (Optional) Configure the maximum number of hops allowed for an EBGP
                             connection.
                             peer { ipv6-address | group-name } ebgp-max-hop [ hop-count ]

                             This step is mandatory if the PE and CE are not directly connected. In
                             most cases, a directly connected physical link must be available between
                             EBGP peers. If you want to establish an EBGP peer relationship between
                             indirectly connected peers, run the peer ebgp-max-hop command to
                             configure the maximum number of hops allowed for an EBGP connection.
                        f.   (Optional) Configure the Site-of-Origin (SoO) attribute.
                             peer { ipv6-address | group-name } soo site-of-origin

                             Several CEs at a VPN site may establish BGP connections with different
                             PEs. The VPN routes advertised from the CEs to the PEs may be re-
                             advertised to the same VPN site after the routes traverse the backbone
                             network. And this may cause routing loops at the VPN site.
                             After the SoO attribute is configured, the PE adds the attribute to a route
                             received from a CE before advertising the route to the peer PE. If the SoO
                             attribute is the same as the local SoO attribute on the peer PE, the peer
                             PE does not send the route to its attached CEs.
                        g.   (Optional) Configure the local device to allow routing loops.
                             peer { ipv6-address | group-name } allow-as-loop [ number ]

                             This step applies to Hub-Spoke networking.
                        h.   (Optional) Enable BGP AS number substitution.
                             peer { ipv6-address | group-name } substitute-as


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                       284
VPN Configuration
VPN Configuration                                                                             4 IPv6 L3VPN Configuration


                              This step needs to be performed on PEs in a scenario where CEs at
                              different sites use the same AS number.


                                   NOTICE

                              Enabling BGP AS number substitution may cause routing loops on a CE
                              multi-homing network.

                    ●   Perform the following steps on the CE:
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    (Optional) Configure the ID of the local device.
                              router-id ipv4-address

                              If no interface on the local CE is configured with an IPv4 address, you
                              need to configure a router ID for the local CE.
                        d.    Configure a PE as a VPN peer.
                              peer ipv6-address as-number as-number

                        e.    (Optional) Configure the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv6-address | group-name } ebgp-max-hop [ hop-count ]

                              In most cases, a directly connected physical link must be available
                              between EBGP peers. If you want to establish an EBGP peer relationship
                              between indirectly connected peers, run the peer ebgp-max-hop
                              command to configure the maximum number of hops allowed for an
                              EBGP connection.
                        f.    Enter the BGP-IPv6 unicast address family view.
                              ipv6-family unicast

