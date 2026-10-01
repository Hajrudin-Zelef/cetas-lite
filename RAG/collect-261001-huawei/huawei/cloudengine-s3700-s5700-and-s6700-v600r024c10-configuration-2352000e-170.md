---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-170
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [24592, 24794]
sha256: 0d803c6b17fcdd61d4be7d54fff058af3cfadaeae306109944fab62719b6d830
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   ASBR2
                        #
                        sysname ASBR2
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 192.168.1.2 255.255.255.0
                         mpls
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
                         ip address 3.3.3.3 255.255.255.255
                        #
                        bgp 200
                         peer 192.168.1.1 as-number 100
                         peer 4.4.4.4 as-number 200
                         peer 4.4.4.4 connect-interface LoopBack1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         388
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                         #
                         ipv4-family unicast
                          peer 192.168.1.1 enable
                          peer 4.4.4.4 enable
                         #
                         ipv6-family vpnv6
                          undo policy vpn-target
                          peer 4.4.4.4 enable
                          peer 192.168.1.1 enable
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        ip vpn-instance vpn1
                         ipv6-family
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
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip binding vpn-instance vpn1
                         ipv6 enable
                         ipv6 address 2001:db8:2::2/64
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
                         ipv6-family vpnv6
                          policy vpn-target
                          peer 3.3.3.3 enable
                         #
                         ipv6-family vpn-instance vpn1
                          peer 2001:db8:2::1 as-number 65002
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         389
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration

                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ipv6 enable
                         ipv6 address 2001:db8:2::1/64
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        bgp 65002
                         router-id 11.11.11.11
                         peer 2001:db8:2::2 as-number 200
                         #
                          ipv4-family unicast
                         #
                         ipv6-family unicast
                          import-route direct
                          peer 2001:db8:2::2 enable
                        #
                        return



4.14 Configuring IPv6 L3VPN over MPLS Hub-Spoke

4.14.1 Understand IPv6 L3VPN over MPLS Hub-Spoke
                    The fundamentals of IPv6 L3VPN over MPLS hub-spoke are similar to those of
                    IPv4 L3VPN over MPLS hub-spoke. For details, see 3.15.1 Understanding IPv4
                    L3VPN over MPLS Hub-Spoke.

4.14.2 Configuring a VPN Instance

Context
                    A VPN instance can be configured to manage VPN routes.

                    In Hub-Spoke networking, the PE connected to a Hub site is called a Hub-PE and
                    that connected to a non-Hub site (Spoke site) is called a Spoke-PE.

                    You need to configure a VPN instance on each Spoke-PE and two VPN instances
                    (VPN-in and VPN-out) on each Hub-PE.
                    ●   VPN-in: receives and maintains all the VPNv6 routes advertised by all the
                        Spoke-PEs.
                    ●   VPN-out: maintains the routes of all Hub and Spoke sites and advertises those
                        routes to all the Spoke-PEs.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         390
VPN Configuration
VPN Configuration                                                                              4 IPv6 L3VPN Configuration


                          NOTE

                         Steps 1 to 8 are used to configure a VPN instance. Configurations of different VPN instances
                         are similar. If different VPN instances are configured on the same device, these VPN
                         instances must have different names, RDs, and descriptions.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Create a VPN instance and enter the VPN instance view.
                    ip vpn-instance vpn-instance-name

                          NOTE

                         The name of a VPN instance is case-sensitive. For example, vpn1 and VPN1 are considered
                         different VPN instances.

