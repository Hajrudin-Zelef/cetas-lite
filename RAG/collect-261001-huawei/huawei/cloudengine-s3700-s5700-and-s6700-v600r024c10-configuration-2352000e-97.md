---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-97
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [13511, 13646]
sha256: 0db5a51728df0cdad9494d184ac6862f67c00df343744bbab5fccc02bb3883e1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The configuration of NPE2 is similar to that of NPE1. On NPE2, configure the local
                    preference as 180 for routes imported from SPE1 and 170 for routes imported
                    from SPE2.
         Step 7 Configure a route-policy to adjust the local preference of the primary and backup
                routes.
                    # Configure VPN FRR on UPEs and NPEs. The following uses UPE1 as an example.
                    [UPE1] bgp 100
                    [UPE1-bgp] ipv4-family vpn-instance vpna
                    [UPE1-bgp-vpna] auto-frr
                    [UPE1-bgp-vpna] route-select delay 300
                    [UPE1-bgp-vpna] quit
                    [UPE1-bgp] quit

                    Configure VPNv4 FRR on SPEs (VPN FRR cannot be configured on SPEs, because
                    SPEs do not have VPN instances). The following uses SPE1 as an example.
                    [SPE1] bgp 100
                    [SPE1-bgp] ipv4-family vpnv4
                    [SPE1-bgp-af-vpnv4] bestroute nexthop-resolved tunnel
                    [SPE1-bgp-af-vpnv4] auto-frr
                    [SPE1-bgp-af-vpnv4] route-select delay 300
                    [SPE1-bgp-af-vpnv4] quit
                    [SPE1-bgp] quit

                    ----End

Verifying the Configuration
                    After the configuration is complete, run the display ip routing-table vpn-
                    instance vpna command on UPE1 or NPE1. The command output shows specific
                    routes from UPE1 to NPE1 or from NPE1 to UPE1. UPE1 and NPE1 can ping each
                    other.
                    <UPE1> display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 8        Routes : 8

                    Destination/Mask      Proto Pre Cost      Flags NextHop         Interface

                         10.1.1.0/24 Direct 0 0           D 10.1.1.2     Vlanif300
                         10.1.1.2/32 Direct 0 0           D 127.0.0.1     Vlanif300
                        10.1.1.255/32 Direct 0 0           D 127.0.0.1     Vlanif300
                         10.4.1.0/24 IBGP 255 0            RD 3.3.3.3      Vlanif100
                         10.2.1.0/24 IBGP 255 0            RD 3.3.3.3      Vlanif100
                         10.3.1.0/24 IBGP 255 0            RD 3.3.3.3      Vlanif100
                          7.7.7.7/32 IBGP 255 0            RD 3.3.3.3     Vlanif100
                    255.255.255.255/32 Direct 0 0            D 127.0.0.1      InLoopBack0
                    <UPE1> ping -vpn-instance vpna 10.3.1.1
                     PING 10.3.1.1: 56 data bytes, press CTRL_C to break
                       Reply from 10.3.1.1: bytes=56 Sequence=1 ttl=251 time=5 ms
                       Reply from 10.3.1.1: bytes=56 Sequence=2 ttl=253 time=3 ms
                       Reply from 10.3.1.1: bytes=56 Sequence=3 ttl=251 time=3 ms
                       Reply from 10.3.1.1: bytes=56 Sequence=4 ttl=253 time=2 ms
                       Reply from 10.3.1.1: bytes=56 Sequence=5 ttl=251 time=2 ms

                     --- 10.3.1.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                        215
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                       0.00% packet loss
                       round-trip min/avg/max = 2/3/5 ms
                    <NPE1> display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 8        Routes : 8

                    Destination/Mask     Proto Pre Cost       Flags NextHop         Interface

                        10.1.1.0/24 IBGP 255 0              RD 3.3.3.3     Vlanif100
                        10.4.1.0/24 Direct 0 0             D 10.4.1.1     Vlanif300
                        10.4.1.1/32 Direct 0 0             D 127.0.0.1    Vlanif300
                       10.4.1.255/32 Direct 0 0             D 127.0.0.1    Vlanif300
                        10.2.1.0/24 EBGP 255 0              RD 10.4.1.2     Vlanif300
                        10.3.1.0/24 EBGP 255 0              RD 10.4.1.2     Vlanif300
                         7.7.7.7/32 EBGP 255 0              RD 10.4.1.2     Vlanif300
                    255.255.255.255/32 Direct 0 0             D 127.0.0.1     InLoopBack0

                    Run the display bgp vpnv4 vpn-instance vpna routing-table 10.3.1.1 command
                    on UPE1. The command output shows that the specified route is received from an
                    SPE functioning as an RR.
                    <UPE1> display bgp vpnv4 vpn-instance vpna routing-table 10.3.1.1
                     BGP local router ID : 1.1.1.1
                     Local AS number : 100

                    VPN-Instance vpna, router ID 1.1.1.1:
                    Paths: 2 available, 1 best, 1 select
                    BGP routing table entry information of 10.3.1.0/24:
                    Remote-Cross route
                    Label information (Received/Applied): 24/NULL
                    From: 3.3.3.3 (3.3.3.3)
                    Route Duration: 0d00h17m56s
                    Relay Tunnel Out-Interface: Vlanif100
                    Original nexthop: 3.3.3.3
                    Qos information : 0x0
                    Ext-Community: RT <1 : 1>
                    AS-path 65420, origin incomplete, MED 0, localpref 200, pref-val 0, valid, internal, best, select, pre 255
                    Originator: 5.5.5.5
                    Cluster list: 3.3.3.3
                    Not advertised to any peer yet

                     BGP routing table entry information of 10.3.1.0/24:
                     Remote-Cross route
                     Label information (Received/Applied): 22/NULL
                     From: 4.4.4.4 (4.4.4.4)
                     Route Duration: 0d00h00m34s
                     Relay Tunnel Out-Interface: Vlanif100
                     Original nexthop: 4.4.4.4
                     Qos information : 0x0
                     Ext-Community: RT <1 : 1>
                     AS-path 65420, origin incomplete, MED 0, localpref 180, pref-val 0, valid, internal, pre 255, not preferred
                    for Local_Pref
                     Originator: 5.5.5.5
                     Cluster list: 4.4.4.4
                     Not advertised to any peer yet

                    Run the display ip routing-table vpn-instance vpna 10.3.1.1 verbose command
                    on UPE1. The command output shows information about the backup label and
                    backup tunnel ID of the route to the EPC side.
                    <UPE1> display ip routing-table vpn-instance vpna 10.3.1.1 verbose
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                    Summary Count : 1


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                    216

