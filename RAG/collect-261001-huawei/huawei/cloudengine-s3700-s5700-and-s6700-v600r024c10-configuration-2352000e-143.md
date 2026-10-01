---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-143
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20279, 20405]
sha256: dc7f9d7ec9f97f82a25156ee3e44929a5d20484f4f63d590449389abf39e1e90
---

                    # Configure DeviceA.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 100
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] port link-type trunk
                    [DeviceA-10GE1/0/1] port trunk allow-pass vlan 100
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface Vlanif100
                    [DeviceA-Vlanif100] ipv6 enable
                    [DeviceA-Vlanif100] ipv6 address 2001:DB8:3::1 64
                    [DeviceA-Vlanif100] quit
                    [DeviceA] interface Loopback 1
                    [DeviceA-Loopback1] ipv6 enable
                    [DeviceA-Loopback1] ipv6 address 2001:DB8:13::3 128
                    [DeviceA-Loopback1] quit
                    [DeviceA] ripng 100
                    [DeviceA-ripng-100] quit
                    [DeviceA] interface Vlanif100
                    [DeviceA-Vlanif100] ripng 100 enable
                    [DeviceA-Vlanif100] quit
                    [DeviceA] interface Loopback 1
                    [DeviceA-Loopback1] ripng 100 enable
                    [DeviceA-Loopback1] quit

                    # Configure DeviceB.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 100
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 100
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface Vlanif100
                    [DeviceB-Vlanif100] ipv6 enable
                    [DeviceB-Vlanif100] ipv6 address 2001:DB8:4::1 64
                    [DeviceB-Vlanif100] quit
                    [DeviceB] interface Loopback 1
                    [DeviceB-Loopback1] ipv6 enable
                    [DeviceB-Loopback1] ipv6 address 2001:DB8:14::4 128
                    [DeviceB-Loopback1] quit
                    [DeviceB] ripng 200
                    [DeviceB-ripng-200] quit
                    [DeviceB] interface Vlanif100
                    [DeviceB-Vlanif100] ripng 200 enable
                    [DeviceB-Vlanif100] quit
                    [DeviceB] interface Loopback 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         322
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    [DeviceB-Loopback1] ripng 200 enable
                    [DeviceB-Loopback1] quit

         Step 4 Disable routing loop detection on the MCE and import RIPng routes destined for
                VPN sites.
                    [MCE] ospfv3 100 vpn-instance vpna
                    [MCE-ospfv3-100] vpn-instance-capability simple
                    [MCE-ospfv3-100] import-route ripng 100
                    [MCE-ospfv3-100] quit
                    [MCE] ospfv3 200 vpn-instance vpnb
                    [MCE-ospfv3-200] vpn-instance-capability simple
                    [MCE-ospfv3-200] import-route ripng 200
                    [MCE-ospfv3-200] quit

                    ----End


Verifying the Configuration
                    After the configuration is complete, run the display ipv6 routing-table vpn-
                    instance command on the MCE to check IPv6 routing information of VPN
                    instances.

                    The following example uses the VPN instance vpna.
                    [MCE] display ipv6 routing-table vpn-instance vpna
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 7        Routes : 7

                    Destination : 2001:DB8:13::3                   PrefixLength : 128
                    NextHop      : FE80::2200:10FF:FE03:0             Preference : 100
                    Cost      :1                            Protocol    : RIPng
                    RelayNextHop : ::                          TunnelID      : 0x0
                    Interface : Vlanif300                       Flags       :D

                    Destination : 2001:DB8:8::                    PrefixLength : 64
                    NextHop      : 2001:DB8:8::2                   Preference : 0
                    Cost      :0                            Protocol    : Direct
                    RelayNextHop : ::                          TunnelID      : 0x0
                    Interface : Vlanif100                       Flags      :D

                    Destination : 2001:DB8:8::2                    PrefixLength : 128
                    NextHop      : ::1                        Preference : 0
                    Cost      :0                            Protocol     : Direct
                    RelayNextHop : ::                          TunnelID       : 0x0
                    Interface : Vlanif100                       Flags       :D

                    Destination : 2001:DB8:3::                    PrefixLength : 64
                    NextHop      : 2001:DB8:3::2                   Preference : 0
                    Cost      :0                            Protocol    : Direct
                    RelayNextHop : ::                          TunnelID      : 0x0
                    Interface : Vlanif300                       Flags      :D

                    Destination : 2001:DB8:3::2                    PrefixLength : 128
                    NextHop      : ::1                        Preference : 0
                    Cost      :0                            Protocol     : Direct
                    RelayNextHop : ::                          TunnelID       : 0x0
                    Interface : Vlanif300                       Flags       :D

                    Destination : FE80::                       PrefixLength : 10
                    NextHop      : ::                        Preference : 0
                    Cost      :0                            Protocol   : Direct
                    RelayNextHop : ::                          TunnelID     : 0x0
                    Interface : NULL0                           Flags     :D



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         323
VPN Configuration
VPN Configuration                                                             4 IPv6 L3VPN Configuration


