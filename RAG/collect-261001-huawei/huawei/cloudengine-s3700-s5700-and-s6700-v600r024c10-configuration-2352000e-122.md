---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-122
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [17236, 17381]
sha256: 0d381db1b103015eebf596e75d4c0e789a5325a1d17c47c9f03552ad26a026ce
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                              4 IPv6 L3VPN Configuration

                    ipv6 enable

         Step 6 Configure an IPv6 address for the interface.
                    ipv6 address { ipv6-address prefix-length | ipv6-address/prefix-length }

                    Some Layer 3 features, such as route exchange between the PE and CE, can be
                    configured only after an IPv6 address is configured for the VPN interface on the
                    PE.

                    ----End

4.5.3 Verifying the Configuration

Procedure
                    ●    Run the display ip vpn-instance vpn-instance-name command to check brief
                         information about a specified VPN instance.
                    ●    Run the display ip vpn-instance verbose vpn-instance-name command to
                         check detailed information about a specified VPN instance.
                    ●    Run the display ip vpn-instance import-vt ivt-value command to check
                         information about VPN instances with the specified import VPN target.

                         This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                         S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-
                         V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
                    ●    Run the display ip vpn-instance [ vpn-instance-name ] interface command
                         to check information about the interface bound to a specified VPN instance.
                    ●    Run the display ipv6 routing-table vpn-instance vpn-instance-name
                         command on a PE to check route information in the VPN instance IPv6
                         address family.
                    ●    Run the display ipv6 routing-table command on a CE to check route
                         information.

                    ----End

4.5.4 Example for Configuring Mutual Access Between Local
IPv6 L3VPNs

Networking Requirements
                    On the network shown in Figure 4-1, CE1 and CE2 both connect to PE1. CE1
                    belongs to VPNA, and CE2 belongs to VPNB. It is required that Site 1 and Site 2
                    communicate with each other. To meet this requirement, configure mutual access
                    between local VPNs.


                    Figure 4-1 Mutual access between local IPv6 VPNs
                          NOTE

                    In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200, respectively.



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                        274
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPN instances on PE1 and different VPN targets for the VPN
                         instances to isolate VPNs.
                    2.   On PE1, bind the interfaces connected to CEs to the corresponding VPN
                         instances to provide access for VPN users.
                    3.   Import direct routes destined for local CEs into the VPN routing tables on PE1.
                         On each CE connected to PE1, configure a static route to the other local CE so
                         that both CEs can communicate with each other.

Procedure
         Step 1 Configure VPN instances on PE1 and bind PE1 interfaces connected to CEs to the
                corresponding VPN instances.
                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv6-family
                    [PE1-vpn-instance-vpna-af-ipv6] route-distinguisher 100:1
                    [PE1-vpn-instance-vpna-af-ipv6] vpn-target 111:1 export-extcommunity
                    [PE1-vpn-instance-vpna-af-ipv6] vpn-target 111:1 222:2 import-extcommunity
                    [PE1-vpn-instance-vpna-af-ipv6] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv6-family
                    [PE1-vpn-instance-vpnb-af-ipv6] route-distinguisher 100:2
                    [PE1-vpn-instance-vpnb-af-ipv6] vpn-target 222:2 export-extcommunity
                    [PE1-vpn-instance-vpnb-af-ipv6] vpn-target 222:2 111:1 import-extcommunity
                    [PE1-vpn-instance-vpnb-af-ipv6] quit
                    [PE1-vpn-instance-vpnb] quit
                    [PE1] vlan batch 100 200
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [PE1-10GE1/0/1] quit
                    [PE1] interface Vlanif100
                    [PE1-Vlanif100] ip binding vpn-instance vpna
                    [PE1-Vlanif100] ipv6 enable
                    [PE1-Vlanif100] ipv6 address 2001:DB8::11:2 112
                    [PE1-Vlanif100] quit


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  275
VPN Configuration
VPN Configuration                                                                 4 IPv6 L3VPN Configuration

                    [PE1] interface 10ge1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 100 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif200
                    [PE1-Vlanif200] ip binding vpn-instance vpnb
                    [PE1-Vlanif200] ipv6 enable
                    [PE1-Vlanif200] ipv6 address 2001:DB8::12:2 112
                    [PE1-Vlanif200] quit

                    # Configure an IP address for the interface used by CE1 to connect to PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 100
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE1-10GE1/0/1] quit
                    [CE1] interface Vlanif100
                    [CE1-Vlanif100] ipv6 enable
                    [CE1-Vlanif100] ipv6 address 2001:DB8::11:1 112

                    [CE1-Vlanif100] quit

                    # Configure an IP address for the interface used by CE2 to connect to PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 100
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] port link-type trunk
                    [CE2-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE2-10GE1/0/1] quit
                    [CE2] interface Vlanif100
                    [CE2-Vlanif100] ipv6 enable
                    [CE2-Vlanif100] ipv6 address 2001:DB8::12:1 112
                    [CE2-Vlanif100] quit

