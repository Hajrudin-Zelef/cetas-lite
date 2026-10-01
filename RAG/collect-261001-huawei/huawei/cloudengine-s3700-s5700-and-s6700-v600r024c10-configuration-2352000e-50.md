---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-50
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [6211, 6376]
sha256: 0432e44075c5eb64bc9aac677259752ee4d9f401e7cdc96cdab529d58f3e60a8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           98
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration

                    [MCE-vpn-instance-vpnb] ipv4-family
                    [MCE-vpn-instance-vpnb-af-ipv4] route-distinguisher 100:2
                    [MCE-vpn-instance-vpnb-af-ipv4] quit
                    [MCE-vpn-instance-vpnb] quit
                    [MCE] vlan batch 100 200 300 400
                    [MCE] interface 10ge 1/0/1
                    [MCE-10GE1/0/1] port link-type trunk
                    [MCE-10GE1/0/1] port trunk allow-pass vlan 100
                    [MCE-10GE1/0/1] quit
                    [MCE] interface Vlanif 100
                    [MCE-Vlanif100] ip binding vpn-instance vpna
                    [MCE-Vlanif100] ip address 10.5.1.2 24
                    [MCE-Vlanif100] quit
                    [MCE] interface 10GE1/0/2
                    [MCE-10GE1/0/2] port link-type trunk
                    [MCE-10GE1/0/2] port trunk allow-pass vlan 200
                    [MCE-10GE1/0/2] quit
                    [MCE] interface Vlanif 200
                    [MCE-Vlanif200] ip binding vpn-instance vpnb
                    [MCE-Vlanif200] ip address 10.5.2.2 24
                    [MCE-Vlanif200] quit
                    [MCE] interface 10GE1/0/3
                    [MCE-10GE1/0/3] port link-type trunk
                    [MCE-10GE1/0/3] port trunk allow-pass vlan 300
                    [MCE-10GE1/0/3] quit
                    [MCE] interface Vlanif 300
                    [MCE-Vlanif300] ip binding vpn-instance vpna
                    [MCE-Vlanif300] ip address 10.3.1.2 24
                    [MCE-Vlanif300] quit
                    [MCE] interface 10GE1/0/4
                    [MCE-10GE1/0/4] port link-type trunk
                    [MCE-10GE1/0/4] port trunk allow-pass vlan 400
                    [MCE-10GE1/0/4] quit
                    [MCE] interface Vlanif 400
                    [MCE-Vlanif400] ip binding vpn-instance vpnb
                    [MCE-Vlanif400] ip address 10.4.1.2 24
                    [MCE-Vlanif400] quit

         Step 2 Configure OSPF multi-instance between PE1 and the MCE.

                    # Configure PE1.
                    [PE1] ospf 100 vpn-instance vpna
                    [PE1-ospf-100] area 0
                    [PE1-ospf-100-area-0.0.0.0] network 10.5.1.0 0.0.0.255
                    [PE1-ospf-100-area-0.0.0.0] quit
                    [PE1-ospf-100] quit
                    [PE1] ospf 200 vpn-instance vpnb
                    [PE1-ospf-200] area 0
                    [PE1-ospf-200-area-0.0.0.0] network 10.5.2.0 0.0.0.255
                    [PE1-ospf-200-area-0.0.0.0] quit
                    [PE1-ospf-200] quit

                    # Configure an MCE.
                    [MCE] ospf 100 vpn-instance vpna
                    [MCE-ospf-100] area 0
                    [MCE-ospf-100-area-0.0.0.0] network 10.5.1.0 0.0.0.255
                    [MCE-ospf-100-area-0.0.0.0] quit
                    [MCE-ospf-100] quit
                    [MCE] ospf 200 vpn-instance vpnb
                    [MCE-ospf-200] area 0
                    [MCE-ospf-200-area-0.0.0.0] network 10.5.2.0 0.0.0.255
                    [MCE-ospf-200-area-0.0.0.0] quit
                    [MCE-ospf-200] quit

         Step 3 Configure RIPv2 on the MCE to import VPN routes from Site 1 and Site 2.

                    # Configure an MCE.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           99
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration

                    [MCE] rip 100 vpn-instance vpna
                    [MCE-rip-100] version 2
                    [MCE-rip-100] network 10.0.0.0
                    [MCE-rip-100] import-route ospf 100
                    [MCE-rip-100] quit
                    [MCE] rip 200 vpn-instance vpnb
                    [MCE-rip-200] version 2
                    [MCE-rip-200] network 10.0.0.0
                    [MCE-rip-200] import-route ospf 200
                    [MCE-rip-200] quit

                    # Configure DeviceA.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceA
                    [DeviceA] vlan batch 100
                    [DeviceA] interface 10ge 1/0/1
                    [DeviceA-10GE1/0/1] port link-type trunk
                    [DeviceA-10GE1/0/1] port trunk allow-pass vlan 100
                    [DeviceA-10GE1/0/1] quit
                    [DeviceA] interface Vlanif 100
                    [DeviceA-Vlanif100] ip address 10.3.1.1 24
                    [DeviceA-Vlanif100] quit
                    [DeviceA] interface Loopback1
                    [DeviceA-Loopback1] ip address 3.3.3.3 32
                    [DeviceA-Loopback1] quit
                    [DeviceA] rip 100
                    [DeviceA-rip-100] version 2
                    [DeviceA-rip-100] network 10.0.0.0
                    [DeviceA-rip-100] network 3.0.0.0
                    [DeviceA-rip-100] quit

                    # Configure DeviceB.
                    <HUAWEI> system-view
                    [HUAWEI] sysname DeviceB
                    [DeviceB] vlan batch 100
                    [DeviceB] interface 10ge 1/0/1
                    [DeviceB-10GE1/0/1] port link-type trunk
                    [DeviceB-10GE1/0/1] port trunk allow-pass vlan 100
                    [DeviceB-10GE1/0/1] quit
                    [DeviceB] interface Vlanif 100
                    [DeviceB-Vlanif100] ip address 10.4.1.1 24
                    [DeviceB-Vlanif100] quit
                    [DeviceB] interface Loopback1
                    [DeviceB-Loopback1] ip address 4.4.4.4 32
                    [DeviceB-Loopback1] quit
                    [DeviceB] rip 200
                    [DeviceB-rip-200] version 2
                    [DeviceB-rip-200] network 10.0.0.0
                    [DeviceB-rip-200] network 4.0.0.0
                    [DeviceB-rip-200] quit

         Step 4 Disable routing loop detection on the MCE and import RIP routes destined for VPN
                sites.
                    [MCE] ospf 100 vpn-instance vpna
                    [MCE-ospf-100] vpn-instance-capability simple
                    [MCE-ospf-100] import-route rip 100
                    [MCE-ospf-100] quit
                    [MCE] ospf 200 vpn-instance vpnb
                    [MCE-ospf-200] vpn-instance-capability simple
                    [MCE-ospf-200] import-route rip 200
                    [MCE-ospf-200] quit

                    ----End




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                         100
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration


Verifying the Configuration
                    After the configuration is complete, run the display ip routing-table vpn-
                    instance command on the MCE to check the routing information of the VPN
                    instances.
                    The following uses the routing table of the VPN instance named vpna as an
                    example.
                    [MCE] display ip routing-table vpn-instance vpna
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table : vpna
                           Destinations : 9        Routes : 9

                    Destination/Mask     Proto Pre Cost       Flags NextHop         Interface

