---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-146
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20783, 20927]
sha256: f84d36db581f975e3b87173e3277dc3e36cc308c374d103ae5d20596cc36938d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After IPv6+VPNv6 hybrid FRR is configured, if the next hop of the route from a PE
                    to a CE is unreachable, traffic can be rapidly switched to the link between the
                    standby PE and the CE for transmission.
                    If a local PE has learned an IPv6 VPN route from a CE and learned a VPNv6 route
                    with the same prefix as the learned IPv6 VPN route from a peer PE, you can
                    configure IPv6+VPNv6 hybrid FRR on the local PE. After IPv6+VPNv6 hybrid FRR is
                    configured, the local PE will generate a pair of primary and backup routes
                    destined for the VPN route prefix. If the link between the local PE and the CE is
                    faulty, the local PE rapidly switches traffic to the link between its standby PE and
                    the CE for transmission.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                329
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


                    IPv6+VPNv6 hybrid FRR can be classified into two types, which apply to different
                    networking environments and are configured differently.
                    ●   IPv6 FRR: applies to networks where a non-BGP routing protocol runs
                        between PEs and CEs.
                    ●   VPN BGP auto FRR: applies to networks where BGP runs between PEs and
                        CEs.

Prerequisites
                    Before configuring IPv6+VPNv6 hybrid FRR, you have completed the following
                    tasks:
                    ●   Configure IPv6 L3VPN.
                    ●   Ensure that a PE learns IPv6 routes with the same prefix from a CE and other
                        VPNv6 peers.

Procedure
                    ●   Configure IPv6 FRR.
                               NOTE

                              This function is not supported by the S3710-H, S5735R-L-V2, S5735E-L-V2, S5735-L-V2,
                              and S5735I-L-V2.
                        a.    Enter the system view.
                              system-view
                        b.    Enter the VPN instance view.
                              ip vpn-instance vpn-instance-name
                        c.    Enter the VPN instance IPv6 address family view.
                              ipv6-family
                        d.    Enable VPN IPv6 FRR.
                              ipv6 frr
                    ●   Configure VPN BGP auto FRR.
                        a.    Enter the system view.
                              system-view
                        b.    Enter the BGP view.
                              bgp as-number
                        c.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name
                        d.    Enable VPN BGP auto FRR.
                              auto-frr
                        e.    Configure a delay for route selection.
                              route-select delay delay-value

                    ----End

Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ ipv6-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IPv6 route in the routing table.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   330
VPN Configuration
VPN Configuration                                                                     4 IPv6 L3VPN Configuration


4.10.5 Example for Configuring IPv6 VPN FRR
Networking Requirements
                    If a CE is dual-homed to two PEs, you can configure IPv6 VPN FRR to ensure that
                    if the primary link between PEs fails, IPv6 VPN services are quickly switched to the
                    backup link.
                    On the network shown in Figure 4-4, PE1 learns two routes with the same prefix
                    to the CE from PE2 and PE3. It is required that PE3 be configured as a backup next
                    hop for the IPv6 VPN route on PE1. In this manner, VPN traffic can be quickly
                    switched to PE3 if PE2 becomes faulty.

                    Figure 4-4 Network diagram of IPv6 VPN FRR
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Precautions
                    ●    VPN instances with different RDs need to be configured on the PEs to which a
                         CE is dual-homed.
                    ●    In an IPv6 VPN FRR scenario, traffic is switched back to the primary path after
                         this path recovers. Because the order in which nodes undergo IGP
                         convergence differs, packet loss may occur during the switchback. To resolve
                         this problem, run the route-select delay delay-value command to configure a
                         route selection delay so that traffic is switched back only after forwarding
                         entries on the devices along the primary path are updated. The delay
                         specified using delay-value depends on various factors, such as the number of
                         routes on each device. Configure an appropriate delay as needed.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure OSPF on PE1, PE2, and PE3 for them to communicate on the MPLS
                         backbone network.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     331
VPN Configuration
VPN Configuration                                                              4 IPv6 L3VPN Configuration


                    2.   Configure basic MPLS capabilities and enable MPLS LDP to establish LDP LSPs
                         on the MPLS backbone network.
                    3.   Configure an IPv6-address-family-enabled VPN instance on PE1, PE2, and PE3.
                         On PE2 and PE3, bind the interfaces connected to the CE to the corresponding
                         VPN instances.
                    4.   Establish EBGP peer relationships between the PEs and CE to import IPv6 VPN
                         routes, and establish MP-IBGP peer relationships between the PEs.
                    5.   Configure static BFD for LDP LSPs on PE1 and PE2.
                    6.   Enable VPN BGP auto FRR on PE1.

Procedure
         Step 1 Configure IPv6 addresses for interfaces on the VPN backbone network and at VPN
                sites. For detailed configurations, see Configuration Scripts.
         Step 2 Configure OSPF on the MPLS backbone network for PEs to communicate. For
                detailed configurations, see Configuration Scripts.
         Step 3 Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                MPLS backbone network.
                    # Configure PE1.
                    <PE1> system-view
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif200
                    [PE1-Vlanif200] mpls
                    [PE1-Vlanif200] mpls ldp
                    [PE1-Vlanif200] quit
                    [PE1] interface Vlanif300
                    [PE1-Vlanif300] mpls
                    [PE1-Vlanif300] mpls ldp
                    [PE1-Vlanif300] quit

