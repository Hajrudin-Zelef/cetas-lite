---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-54
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [6887, 7046]
sha256: e22536f8b55f6103e4f7beae95e0c7d6f912deceb7a797f89f8de316be12b034
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                109
VPN Configuration
VPN Configuration                                                                  3 IPv4 L3VPN Configuration


                    local PE will generate a pair of primary and backup routes destined for the VPN
                    route prefix. If the link between the local PE and the CE is faulty, the local PE
                    rapidly switches traffic to the link between its standby PE and the CE for
                    transmission.

                    On the network shown in Figure 3-18, PE2 uses Link_A to forward traffic to the
                    CE and uses Link_B as the backup link in normal situations. When PE2 detects that
                    the route to the CE is unreachable, it immediately switches traffic to Link_B and
                    then performs operations such as VPN route convergence to minimize the impact
                    of the link failure on VPN services.


                    Figure 3-18 IP+VPNv4 hybrid FRR networking




                    IP+VPNv4 hybrid FRR can be classified into two types, which apply to different
                    networking environments and are configured differently.

                    ●   VPN IP FRR: This method applies to networks where a non-BGP protocol runs
                        between PEs and CEs.
                    ●   VPN BGP auto FRR: This method applies to networks where BGP runs
                        between PEs and CEs.


Procedure
                    ●   Configure IP FRR.
                              NOTE

                             This function is not supported by the S3710-H, S5735R-L-V2, S5735E-L-V2, S5735-L-V2,
                             and S5735I-L-V2.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the VPN instance view.
                             ip vpn-instance vpn-instance-name

                        c.   Enter the VPN instance IPv4 address family view.
                             ipv4-family

                        d.   Enable VPN IP FRR.
                             ip frr

                    ●   Configure VPN BGP auto FRR.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  110
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


                               NOTE

                              Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S,
                              S6750-S, S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and
                              S5735-S-V2 series support BGP.
                         a.   Enter the system view.
                              system-view
                         b.   Enter the BGP view.
                              bgp as-number
                         c.   Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name
                         d.   Enable VPN BGP auto FRR.
                              auto-frr
                         e.   Configure delayed route selection.
                              route-select delay delay-value

                    ----End

Verifying the Configuration
                    Run the display ip routing-table vpn-instance vpn-instance-name [ ipv4-
                    address ] verbose command to check the backup outbound interface and backup
                    next hop of a VPN IP route in the routing table.

3.10.6 Example for Configuring VPN FRR
Networking Requirements
                    On the network shown in Figure 3-19, it is required that a backup next hop be
                    configured on PE1 (with PE3 serving as a backup for PE2), so that when PE2 fails,
                    the traffic can be quickly switched to PE3.

                    Figure 3-19 VPN FRR networking
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     111
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration


Precautions
                    Note the following during the configuration:
                    ●    VPN instances with different RDs need to be configured on the PEs to which a
                         CE is dual-homed.
                    ●    In a VPN FRR scenario, traffic is switched back to the primary path after this
                         path recovers. Because the order in which nodes undergo IGP convergence
                         differs, packet loss may occur during the switchback. To resolve this problem,
                         run the route-select delay delay-value command to configure a route
                         selection delay so that traffic is switched back only after forwarding entries on
                         the devices along the primary path are updated. The delay specified using
                         delay-value depends on various factors, such as the number of routes on each
                         device. Configure an appropriate delay as needed.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure OSPF on each PE (PE1, PE2, and PE3) for them to communicate on
                         the MPLS backbone network.
                    2.   Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                         MPLS backbone network.
                    3.   Configure a VPN instance on PE1, PE2, and PE3. On PE2 and PE3, bind the
                         interfaces connected to CE1 to the corresponding VPN instances.
                    4.   Establish an MP-EBGP peer relationship between each PE and the CE to
                         import VPN routes. Establish an MP-IBGP peer relationship between PE1 and
                         PE2 and between PE1 and PE3.
                    5.   Configure VPN FRR on PE1.

Procedure
         Step 1 Configure IP addresses for involved interfaces on the VPN backbone network and
                at VPN sites. For detailed configurations, see Configuration Scripts.
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


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          112
VPN Configuration
VPN Configuration                                                                                     3 IPv4 L3VPN Configuration


