---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-28
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3045, 3190]
sha256: a289d4e2def1c4e973daacda616d727609cd7078433262cee897222ade29cf8d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    When an address family in a VPN instance is disabled, the configuration of this
                    address family on the interface is deleted; if no address family is configured for a
                    VPN instance, the interface is unbound from the VPN instance.

                          NOTE

                         Binding a VPN instance to a loopback interface is usually used to test whether VPNs can
                         communicate with each other. The prerequisite is that the VPN instance has been bound to
                         a VLANIF interface or physical interface. In actual service applications, a VPN instance must
                         be bound to a VLANIF interface, physical interface, or tunnel interface.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface to be bound to the VPN instance.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from
                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.

         Step 4 Bind the interface to a VPN instance.
                    ip binding vpn-instance vpn-instance-name

                    By default, an interface functions as a public network interface and is not bound
                    to any VPN instance.

                          NOTE

                         Running the ip binding vpn-instance command deletes Layer 3 (including IPv4 and IPv6)
                         configurations, such as IP address and routing protocol configurations, on the involved
                         interface. If needed, reconfigure them after running the command.

         Step 5 Configure an IP address for the interface.
                    ip address ip-address { mask | mask-length }

                    After an IP address is configured for the VPN interface, certain Layer 3 features,
                    such as route exchange, can be configured between the PE and CE.

                    ----End




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                       47
VPN Configuration
VPN Configuration                                                                    3 IPv4 L3VPN Configuration


3.5.3 Verifying the Configuration
Procedure
                    ●    Run the display ip vpn-instance vpn-instance-name command to check brief
                         information about a specified VPN instance.
                    ●    Run the display ip vpn-instance verbose vpn-instance-name command to
                         check detailed information about a specified VPN instance.
                    ●    Run the display ip vpn-instance import-vt ivt-value command to check
                         information about all VPN instances with the specified import VPN target.
                         This command is supported only by the S6780-H, S6750-H, S6730E-H-V2,
                         S6730-H-V2, S5732-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-S-V2, S5735I-S-
                         V2, S5735I-H-V2, S6750E-S, S6750-S, S5755-S, S5755E-H, and S5755-H.
                    ●    Run the display ip vpn-instance [ vpn-instance-name ] interface command
                         to check information about the interface bound to a specified VPN instance.
                    ●    Run the display ip routing-table vpn-instance vpn-instance-name command
                         on a PE to check information about routes in a specified VPN instance IPv4
                         address family.
                    ●    Run the display ip routing-table command on a CE to check routing
                         information.
                    ----End

3.5.4 Example for Configuring Mutual Access Between Local
IPv4 L3VPNs
Networking Requirements
                    As shown in Figure 3-8, CE1 and CE2 are connected to PE1. CE1 belongs to vpna,
                    and CE2 belongs to vpnb. It is required that Site 1 and Site 2 communicate with
                    each other. To meet this requirement, configure mutual access between local
                    VPNs.

                    Figure 3-8 Mutual access between local IPv4 VPNs
                         NOTE

                    In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      48
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPN instances on PE1 and different VPN targets for the instances to
                         isolate VPNs.
                    2.   On PE1, bind the interfaces connected to CEs to the VPN instances to provide
                         access for VPN users.
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
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [PE1-vpn-instance-vpna-af-ipv4] vpn-target 111:1 export-extcommunity
                    [PE1-vpn-instance-vpna-af-ipv4] vpn-target 111:1 222:2 import-extcommunity
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv4-family
                    [PE1-vpn-instance-vpnb-af-ipv4] route-distinguisher 100:2
                    [PE1-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 export-extcommunity
                    [PE1-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 111:1 import-extcommunity
                    [PE1-vpn-instance-vpnb-af-ipv4] quit
                    [PE1-vpn-instance-vpnb] quit
                    [PE1] vlan batch 100 200
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 100 200
                    [PE1-10GE1/0/1] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] ip binding vpn-instance vpna
                    [PE1-Vlanif100] ip address 10.1.1.2 24
                    [PE1-Vlanif100] quit
                    [PE1] interface 10ge1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 100 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpnb
                    [PE1-Vlanif200] ip address 10.2.1.2 24
                    [PE1-Vlanif200] quit

