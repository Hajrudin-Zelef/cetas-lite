---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-76
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [10319, 10500]
sha256: 7aebf98aa7a3af13edbb44c23cbbdd0d5834e2180df7e1bf08bc6d28b033759e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv4-family
                          route-distinguisher 200:1
                          vpn-target 1:1 export-extcommunity
                          vpn-target 1:1 import-extcommunity
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.17.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         163
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ip address 10.2.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        bgp 200
                         peer 3.3.3.3 as-number 200
                         peer 3.3.3.3 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          peer 3.3.3.3 enable
                         #
                         ipv4-family vpnv4
                          policy vpn-target
                          peer 3.3.3.3 enable
                         #
                         ipv4-family vpn-instance vpn1
                          peer 10.2.1.1 as-number 65002
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.17.1.0 0.0.0.255
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.2.1.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface Loopback 1
                         ip address 6.6.6.6 255.255.255.255
                        #
                        bgp 65002
                         peer 10.2.1.2 as-number 200
                         network 6.6.6.6 255.255.255.255
                         #
                         ipv4-family unicast
                          peer 10.2.1.2 enable
                        #
                        return



3.13 Configuring IPv4 L3VPN over MPLS Inter-AS
Option B (ASBRs Also Functioning as PEs)


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         164
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


3.13.1 Understanding IPv4 L3VPN over MPLS Inter-AS Option
B (ASBRs Also Functioning as PEs)

Context
                    If ASBRs are capable of managing VPN routes and function as PEs for CE access,
                    inter-AS VPN Option B (ASBRs also functioning as PEs) can be used. This solution
                    requires ASBRs to maintain and advertise not only the VPNv4 routes of its own
                    VPN instances but also the VPNv4 routes of other VPN instances.
                    Inter-AS VPN Option B (ASBRs also functioning as PEs) is similar to inter-AS VPN
                    Option B (basic networking). For details, see 3.12.1 Understanding IPv4 L3VPN
                    over MPLS Inter-AS Option B (Basic Networking).

Prerequisites
                    Before configuring inter-AS VPN Option B (ASBRs also functioning as PEs), you
                    have completed the following tasks:
                    ●    Configure IGP for the MPLS backbone network in each AS to ensure IP
                         connectivity for the backbone network in each AS.
                    ●    Configure basic MPLS functions for the MPLS backbone network in each AS
                         and establish LDP LSPs or TE tunnels between MP-IBGP peers.
                    ●    Configure an IPv4 VPN instance and bind an interface to the IPv4 VPN
                         instance on each PE connected to CEs.
                    ●    Configure IP addresses on interfaces that connect CEs to PEs.

3.13.2 Configuring MP-IBGP Between the PE and ASBR in the
Same AS

Context
                    MP-IBGP, which introduces extended community attributes into BGP, can advertise
                    VPNv4 routes between PEs and ASBRs.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Establish an IBGP peer relationship between the PE and ASBR in the same AS.
                    peer peer-address as-number as-number

         Step 4 Configure a loopback interface as the outbound interface of the BGP session.
                    peer peer-address connect-interface loopback interface-number

         Step 5 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             165
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


         Step 6 Enable the function to exchange VPNv4 routes between the PE and ASBR in the
                same AS.
                    peer peer-address enable

                    ----End

3.13.3 Configuring MP-EBGP Between ASBRs in Different ASs

Context
                    After an MP-EBGP peer relationship is established between ASBRs, an ASBR can
                    advertise VPNv4 routes in the local AS to the other ASBR.
                    In inter-AS VPN Option B (ASBRs also functioning as PEs), VPN instances also
                    need to be created on ASBRs. An ASBR does not filter VPNv4 routes received from
                    a PE in the same AS based on VPN targets. Instead, it advertises the received
                    routes to the peer ASBR through MP-EBGP.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface connected to the peer ASBR.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

