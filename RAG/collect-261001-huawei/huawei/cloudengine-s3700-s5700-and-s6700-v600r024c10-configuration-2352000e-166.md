---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-166
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [23939, 24109]
sha256: 72e38a2b5c43ef565ecc9e721bea2f50793b6ae562c71fadd0266c62ee5e439a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Prerequisites
                    Before configuring inter-AS VPN Option B, you have completed the following
                    tasks:
                    ●   Configure IGP for the MPLS backbone network in each AS to ensure IP
                        connectivity for the backbone network in each AS.
                    ●   Configure basic MPLS functions for the MPLS backbone network in each AS
                        and establish LDP LSPs or TE tunnels between MP-IBGP peers.
                    ●   Configure an IPv6 VPN instance and bind an interface to the IPv6 VPN
                        instance on each PE connected to CEs.
                    ●   Configure IPv6 addresses on interfaces that connect CEs to PEs.

4.13.1 Understand IPv6 L3VPN over MPLS Inter-AS Option B
                    The fundamentals of inter-AS IPv6 VPN Option B are similar to those of inter-AS
                    VPN Option B. For details, see 3.12.1 Understanding IPv4 L3VPN over MPLS
                    Inter-AS Option B (Basic Networking).

Context
                    If ASBRs can manage VPN routes but there are insufficient interfaces available for
                    the dedicated use of all inter-AS VPNs, you can use inter-AS IPv6 VPN Option B.
                    This solution eliminates the need to create VPN instances on ASBRs, but requires
                    ASBRs to maintain and advertise VPNv6 routes. On the network shown in Figure
                    4-9, the connected interfaces between ASBRs do not need to be bound to the
                    VPN. A single-hop MP-EBGP peer relationship is set up between the ASBRs to
                    transmit all inter-AS VPN routes.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          377
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration


                    Figure 4-9 Inter-AS IPv6 VPN Option B




Prerequisites
                    Before configuring inter-AS VPN Option B, you have completed the following
                    tasks:
                    ●    Configure IGP for the MPLS backbone network in each AS to ensure IP
                         connectivity for the backbone network in each AS.
                    ●    Configure basic MPLS functions for the MPLS backbone network in each AS
                         and establish LDP LSPs or TE tunnels between MP-IBGP peers.
                    ●    Configure an IPv6 VPN instance and bind an interface to the IPv6 VPN
                         instance on each PE connected to CEs.
                    ●    Configure IPv6 addresses on interfaces that connect CEs to PEs.

4.13.2 Configuring MP-IBGP Between the PE and ASBR in the
Same AS

Context
                    MP-IBGP, which introduces extended community attributes into BGP, can advertise
                    VPNv6 routes between PEs and ASBRs.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Establish an IBGP peer relationship between the PE and ASBR in the same AS.
                    peer ipv4-address as-number as-number


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                         378
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


         Step 4 Configure a loopback interface as the outbound interface of the BGP session.
                    peer ipv4-address connect-interface loopback interface-number

         Step 5 Enter the BGP-VPNv6 address family view.
                    ipv6-family vpnv6 [ unicast ]

         Step 6 Enable the function to exchange VPNv4 routes between the PE and ASBR in the
                same AS.
                    peer ipv4-address enable

                    ----End

4.13.3 Configuring MP-EBGP Between ASBRs in Different ASs

Context
                    After an MP-EBGP peer relationship is established between ASBRs, an ASBR can
                    advertise VPNv6 routes in the local AS to the other ASBR.
                    In inter-AS IPv6 VPN Option B, you need not create VPN instances on ASBRs. An
                    ASBR does not filter VPNv6 routes received from the PE in the local AS based on
                    VPN targets. Instead, it advertises the received routes to the peer ASBR through
                    MP-EBGP.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface connected to the peer ASBR.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.
         Step 4 Configure IP addresses for interfaces.
                    ip address ip-address { mask | mask-length }

         Step 5 Enable MPLS.
                    mpls

         Step 6 Return to the system view.
                    quit

         Step 7 Enter the BGP view.
                    bgp as-number

         Step 8 Specify the peer ASBR as an EBGP peer.
                    peer ipv4-address as-number as-number

         Step 9 Configure a loopback interface as the outbound interface of the BGP session.
                    peer ipv4-address connect-interface loopback interface-number


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             379
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


        Step 10 Enter the BGP-VPNv6 address family view.
                    ipv6-family vpnv6 [ unicast ]

        Step 11 Enable the function to exchange VPNv6 routes between the PE and ASBR in the
                same AS.
                    peer ipv4-address enable

                    ----End

4.13.4 Controlling the Import and Export of VPN Routes on an
ASBR

Context
                    By default, an ASBR filters received VPNv6 routes based on VPN targets. Eligible
                    routes are imported into the routing table, and other routes are discarded.
                    Therefore, if no VPN instance is configured on the ASBR or no VPN target is
                    configured for the VPN instance, the ASBR discards all the received VPNv6 routes.
                    An ASBR can control the import and export of VPN routes using multiple methods.
                    For example:
                    ●    The ASBR can keep all VPN-IPv6 routes instead of filtering these routes based
                         on VPN targets.
                    ●    The ASBR can use a route-policy to filter VPNv6 routes based on VPN targets
                         and keeps only the eligible ones.
                    Use either of the preceding methods on each ASBR based on the actual situation.

Procedure
                    ●    No VPN target-based route filtering
                         a.    Enter the system view.
                               system-view

                         b.    Enter the BGP view.
                               bgp as-number

                         c.    Enter the BGP-VPNv6 address family view.
                               ipv6-family vpnv6 [ unicast ]

                         d.    Configure the device not to filter VPNv6 routes based on VPN targets.
                               undo policy vpn-target

