---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-58
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [7529, 7665]
sha256: 6809ddb181b559419be6e6a23aae0413933ed4858e622ffdd9157549317c4fcd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration


                    1.    Configure IP addresses for the interfaces on the devices at the VPN site.
                    2.    Configure IGP at the VPN site, so that the routes to the loopback interface of
                          DeviceA can be advertised to CE1.
                    3.    Configure a VPN instance on the PE and bind interfaces connecting the PE to
                          CEs to the VPN instance.
                    4.    Establish EBGP peer relationships between the PE and CEs.
                    5.    Configure OSPF and BGP to import routes from each other on CE1.
                    6.    Configure VPN static routes on the PE.
                    7.    Enable VPN IP FRR on the PE.

Procedure
         Step 1 Configure IP addresses for the interfaces on the devices at the VPN site.

                    For detailed configurations, see Configuration Scripts.

         Step 2 Configure IGP at the VPN site, so that the routes to the loopback interface of
                DeviceA can be advertised to CE1. The following example uses OSPF.

                    # Configure CE1.
                    [CE1] ospf 1
                    [CE1-ospf] area 0
                    [CE1-ospf-1-area-0.0.0.0] network 10.3.1.0 0.0.0.255
                    [CE1-ospf-1-area-0.0.0.0] quit
                    [CE1-ospf] quit

                    The configuration of DeviceA is similar to the configuration of CE1. For detailed
                    configurations, see Configuration Scripts.

                    After completing the configuration, run the display ip routing-table command on
                    the CEs. The command output shows that CE1 and CE2 have learned the route to
                    Loopback 1 on DeviceA. The following example uses the command output on CE1.
                    <CE1> display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : _public_
                           Destinations : 11        Routes : 11

                    Destination/Mask     Proto Pre Cost       Flags NextHop         Interface

                         10.1.1.0/24 Direct 0 0            D 10.1.1.2     Vlanif100
                         10.1.1.1/32 Direct 0 0            D 127.0.0.1    Vlanif100
                       10.1.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif100
                      11.11.11.11/32 OSPF 10 1                D 10.3.1.2      Vlanif200
                         10.3.1.0/24 Direct 0 0            D 10.3.1.1     Vlanif200
                         10.3.1.1/32 Direct 0 0            D 127.0.0.1    Vlanif200
                       10.3.1.255/32 Direct 0 0             D 127.0.0.1     Vlanif200
                        127.0.0.0/8 Direct 0 0             D 127.0.0.1    InLoopBack0
                        127.0.0.1/32 Direct 0 0            D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0             D 127.0.0.1      InLoopBack0

         Step 3 Configure a VPN instance on the PE and bind interfaces connecting the PE to CEs
                to the VPN instance.

                    # Configure a VPN instance (vpna) on the PE and bind VLANIF 100 and VLANIF
                    200 to vpna.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         120
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    <PE> system-view
                    [PE] ip vpn-instance vpna
                    [PE-vpn-instance-vpna] ipv4-family
                    [PE-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [PE-vpn-instance-vpna-af-ipv4] vpn-target 100:100
                    [PE-vpn-instance-vpna-af-ipv4] quit
                    [PE-vpn-instance-vpna] quit
                    [PE] interface Vlanif100
                    [PE-Vlanif100] ip binding vpn-instance vpna
                    [PE-Vlanif100] ip address 10.1.1.1 24
                    [PE-Vlanif100] quit
                    [PE] interface Vlanif200
                    [PE-Vlanif200] ip binding vpn-instance vpna
                    [PE-Vlanif200] ip address 10.2.1.1 24
                    [PE] quit

         Step 4 Establish EBGP peer relationships between the PE and CEs.
                    # Configure the PE.
                    [PE] bgp 100
                    [PE-bgp] ipv4-family vpn-instance vpna
                    [PE-bgp-vpna] peer 10.1.1.2 as-number 65410
                    [PE-bgp-vpna] quit
                    [PE-bgp] quit

                    # Configure CE1.
                    [CE1] bgp 65410
                    [CE1-bgp] peer 10.1.1.1 as-number 100
                    [CE1-bgp] quit

                    After completing the configuration, run the display bgp vpnv4 vpn-instance
                    vpna peer command on the PE. The command output shows that the status of
                    the EBGP peer relationships between the PE and CEs is Established, indicating
                    that EBGP peer relationships have been established between the PE and CEs.
                    <PE> display bgp vpnv4 vpn-instance vpna peer

                    BGP local router ID : 1.1.1.9
                    Local AS number : 100
                    Total number of peers : 1        Peers in established state : 1

                     Peer         V       AS MsgRcvd MsgSent OutQ Up/Down         State PrefRcv
                     10.1.1.2     4      65410   21   23   0 00:17:47 Established      1

         Step 5 Configure OSPF and BGP to import routes from each other on the CEs.
                    # Configure CE1.
                    [CE1] bgp 65410
                    [CE1-bgp] network 11.11.11.11 32
                    [CE1-bgp] quit
                    [CE1] ospf 1
                    [CE1-ospf-1] import-route bgp
                    [CE1-ospf-1] quit

                    After completing the configuration, run the display ip routing-table vpn-
                    instance command on the PE. The command output shows the route to the
                    loopback interface on DeviceA.
                    <PE> display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 8        Routes : 8



Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       121
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    Destination/Mask     Proto Pre Cost       Flags NextHop         Interface

