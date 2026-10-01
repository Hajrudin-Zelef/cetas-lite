---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-33
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3715, 3865]
sha256: 00ce4ff3e6bfde31a98c24e03b5fb44e282673bc9ef5cc8fb5efe32b100a2dce
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    When an address family in a VPN instance is disabled, the configuration of this
                    address family on the interface is deleted; if no address family is configured for a
                    VPN instance, the interface is unbound from the VPN instance.

                          NOTE

                         Binding a VPN instance to a loopback interface is usually used to test whether VPNs can
                         communicate with each other. The prerequisite is that the VPN instance has been bound to
                         a VLANIF interface or physical interface. In actual service applications, a VPN instance must
                         be bound to a VLANIF interface, physical interface, or tunnel interface.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface to be bound to the VPN instance.
                    interface interface-type interface-number


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       58
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.

         Step 4 Bind the interface to a VPN instance.
                    ip binding vpn-instance vpn-instance-name

                    By default, an interface functions as a public network interface and is not bound
                    to any VPN instance.

                          NOTE

                         Running the ip binding vpn-instance command deletes Layer 3 (including IPv4 and IPv6)
                         configurations, such as IP address and routing protocol configurations, on the involved
                         interface. If needed, reconfigure them after running the command.

         Step 5 Configure an IP address for the interface.
                    ip address ip-address { mask | mask-length }

                    After an IP address is configured for the VPN interface, certain Layer 3 features,
                    such as route exchange, can be configured between the PE and CE.

                    ----End

3.6.4 Establishing MP-IBGP Peer Relationships Between PEs

Context
                    If VPN sites in a basic IPv4 L3VPN need to communicate, PEs must use MP-IBGP to
                    advertise VPNv4 routes carrying the RD attribute to each other. Since all the PEs
                    reside in the same AS, MP-IBGP peer relationships can be established between
                    them.

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




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                  59
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


                          NOTE

                        A PE must use a loopback interface address with a 32-bit mask to set up an MP-IBGP peer
                        relationship with the peer PE so that VPN routes can recurse to tunnels. The route to the
                        local loopback interface is advertised to the peer PE using IGP on the MPLS backbone
                        network.

         Step 5 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4

         Step 6 (Optional) Enable the VPNv4 route matching rule.
                    activate-route-tag compatible

                    By default, the tag value of VPNv4 routes is 0, and the VPNv4 routes that carry
                    tag 0 can be matched against the rule configured using the if-match tag
                    command. If the tag value of a VPNv4 route is not 0, you need to run the
                    activate-route-tag compatible command to enable the if-match tag command,
                    so that the VPNv4 route can be matched against the rule configured using the if-
                    match tag command.
         Step 7 Enable the ability to exchange VPN-IPv4 routes with a specified peer.
                    peer ipv4-address enable

                    ----End

3.6.5 Configuring EBGP Between the PE and CE

Context
                    On an IPv4 L3VPN, a routing protocol must be configured between PEs and CEs to
                    allow them to communicate and allow one CE to obtain routes to other CEs.
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

                         b.   Enter the BGP view.
                              bgp as-number

                         c.   Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                         d.   Configure a CE as a VPN peer.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                    60
VPN Configuration
VPN Configuration                                                                              3 IPv4 L3VPN Configuration

                              peer ipv4-address as-number as-number

                        e.    (Optional) Configure the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

