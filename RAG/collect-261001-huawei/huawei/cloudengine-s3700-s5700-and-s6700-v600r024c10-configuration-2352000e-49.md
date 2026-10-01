---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-49
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [6042, 6210]
sha256: bdbd84c7c463c8dccbde4ad6f2753c471b9affac1d3b7e451650f5847b8061bc
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support RIP.


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                             95
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration


                    Deleting a VPN instance or disabling a VPN instance IPv4 address family will
                    delete all the RIP processes bound to the VPN instance and the VPN instance IPv4
                    address family.


Procedure
                    ●   Configure the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Create a RIP process and enter the RIP view.
                              rip process-id vpn-instance vpn-instance-name

                              A RIP process can be bound to only one VPN instance.
                        c.    Enable RIP on the network segment of the interface to which the VPN
                              instance is bound.
                              network network-address

                        d.    Import BGP routes.
                              import-route bgp [ cost { cost | transparent } | route-policy route-policy-name ] *

                        e.    Return to the system view.
                              quit

                        f.    Enter the BGP view.
                              bgp as-number

                        g.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        h.    Import RIP routes into the routing table of the BGP VPN instance IPv4
                              address family.
                              import-route rip process-id [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.    Enter the system view.
                              system-view

                        b.    Create a RIP process and enter the RIP view.
                              rip process-id vpn-instance vpn-instance-name

                              A RIP process can be bound to only one VPN instance.

                              If a RIP process is not bound to any VPN instance before it is started, this
                              process becomes a public network process and cannot be bound to a VPN
                              instance later.
                        c.    Enable RIP on the network segment of the interface to which the VPN
                              instance is bound.
                              network network-address

                    ----End

3.9.7 Verifying the Configuration



Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                      96
VPN Configuration
VPN Configuration                                                                      3 IPv4 L3VPN Configuration


Procedure
                    ●    Run the display ip routing-table vpn-instance vpn-instance-name
                         [ verbose ] command on the MCE to check the IP routing table of a VPN
                         instance.
                    ----End

3.9.8 Example for Configuring an MCE

Networking Requirements
                    On the network shown in Figure 3-14, Site 1 and Site 2 are both connected to
                    PE1, but the two sites need to be isolated from each other and use independent
                    address spaces. To meet the preceding requirements and reduce costs, an MCE
                    solution can be used.

                    Figure 3-14 MCE networking
                          NOTE

                    In this example, interface 1, interface 2, interface 3, and interface 4 represent VLANIF 100,
                    VLANIF 200, VLANIF 300, and VLANIF 400, respectively.




Precautions
                    Note the following during the configuration:
                    ●    The MCE must have multiple VPN instances configured and have a different
                         interface bound to each VPN instance.
                    ●    Routing loop detection must be disabled on the MCE so that the MCE
                         exchanges routing information with the PE using the OSPF multi-VPN-
                         instance.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         97
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration


                    ●    Configure RIPv2 on the MCE to import VPN routes from Site 1 and Site 2.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPN instances on the MCE and PE1 and bind related interfaces to
                         the VPN instances.
                    2.   Configure OSPF multi-instance on the MCE and PE1 to exchange VPN routing
                         information.
                    3.   Configure RIPv2 on the MCE, DeviceA, and DeviceB to exchange VPN routes.
                    4.   Disable routing loop detection on the MCE and import RIP routes destined for
                         VPN sites.

Procedure
         Step 1 Configure VPN instances on the MCE and PE1 and bind related interfaces to the
                VPN instances.
                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv4-family
                    [PE1-vpn-instance-vpna-af-ipv4] route-distinguisher 200:1
                    [PE1-vpn-instance-vpna-af-ipv4] vpn-target 111:1 both
                    [PE1-vpn-instance-vpna-af-ipv4] quit
                    [PE1-vpn-instance-vpna] quit
                    [PE1] ip vpn-instance vpnb
                    [PE1-vpn-instance-vpnb] ipv4-family
                    [PE1-vpn-instance-vpnb-af-ipv4] route-distinguisher 200:2
                    [PE1-vpn-instance-vpnb-af-ipv4] vpn-target 222:2 both
                    [PE1-vpn-instance-vpnb-af-ipv4] quit
                    [PE1-vpn-instance-vpnb] quit
                    [PE1] vlan batch 100 200
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [PE1-10GE1/0/1] quit
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] ip binding vpn-instance vpna
                    [PE1-Vlanif100] ip address 10.5.1.1 24
                    [PE1-Vlanif100] quit
                    [PE1] interface 10GE1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 200
                    [PE1-10GE1/0/2] quit
                    [PE1] interface Vlanif 200
                    [PE1-Vlanif200] ip binding vpn-instance vpnb
                    [PE1-Vlanif200] ip address 10.5.2.1 24
                    [PE1-Vlanif200] quit

                    # Configure an MCE.
                    <HUAWEI> system-view
                    [HUAWEI] sysname MCE
                    [MCE] ip vpn-instance vpna
                    [MCE-vpn-instance-vpna] ipv4-family
                    [MCE-vpn-instance-vpna-af-ipv4] route-distinguisher 100:1
                    [MCE-vpn-instance-vpna-af-ipv4] quit
                    [MCE-vpn-instance-vpna] quit
                    [MCE] ip vpn-instance vpnb


