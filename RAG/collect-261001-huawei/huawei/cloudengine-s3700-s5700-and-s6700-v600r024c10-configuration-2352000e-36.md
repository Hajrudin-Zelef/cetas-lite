---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-36
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [4116, 4266]
sha256: 49b22a83ab394f701ce27708530bf9a98e1c600517e59419675c8f6abda9c475
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                     A VPN that receives routes outside of it from devices other than PEs and
                                     advertises these routes to PEs is called a transit VPN. A VPN that receives only
                                     routes in it and routes advertised by PEs is called a stub VPN. Generally, a static
                                     route is only used for route exchange between the CE and PE in a stub VPN.
                    ●   Configure OSPF on the CE. The configuration details are not provided here.
                    ----End

3.6.9 Configuring IS-IS Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE.
                        a.    Enter the system view.
                              system-view
                        b.    Create an IS-IS instance between the PE and CE and enter the IS-IS view.
                              isis process-id vpn-instance vpn-instance-name

                              An IS-IS process can be bound to only one VPN instance. If an IS-IS
                              process is not bound to any VPN instance before it is started, this process
                              becomes a public network process and cannot be bound to a VPN
                              instance later.
                        c.    Set a NET.
                              network-entity net-addr

                              A NET specifies the current IS-IS area address and the system ID.
                        d.    (Optional) Configure the level of the device.
                              is-level { level-1 | level-1-2 | level-2 }
                        e.    Import BGP routes.
                              import-route bgp [ cost-type { external | internal } | cost cost | tag tag | route-policy route-
                              policy-name | [ level-1 | level-2 | level-1-2 ] ] *

                              If the IS-IS level is not specified in the command, BGP routes will be
                              imported into the Level-2 IS-IS routing table.
                        f.    Return to the system view.
                              quit
                        g.    Enter the BGP view.
                              bgp as-number
                        h.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                               65
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration


                        i.    Import IS-IS routes into the routing table of the BGP VPN instance IPv4
                              address family.
                              import-route isis process-id [ med med | route-policy route-policy-name ] *

                    ●   Configure IS-IS on the CE. The configuration details are not provided here.

                    ----End

3.6.10 Configuring RIP Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Create a RIP instance between the PE and CE and enter the RIP view.
                              rip process-id vpn-instance vpn-instance-name

                              A RIP process can be bound to only one VPN instance. If a RIP process is
                              not bound to any VPN instance before it is started, this process becomes
                              a public network process and cannot be bound to a VPN instance later.
                        c.    Enable RIP on the network segment of the interface to which the VPN
                              instance is bound.
                              network network-address

                        d.    Import BGP routes.
                              import-route bgp [ cost { cost | transparent } | route-policy route-policy-name ] *

                              After the import-route bgp command is run in the RIP view, the PE can
                              import the VPNv4 routes learned from the remote PE into the RIP routing
                              table and advertise them to the attached CE.
                        e.    Return to the system view.
                              quit

                        f.    Enter the BGP view.
                              bgp as-number

                        g.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        h.    Import RIP routes into the routing table of the BGP VPN instance IPv4
                              address family.
                              import-route rip process-id [ med med | route-policy route-policy-name ] *

                    ●   Configure RIPv2 on the CE. The configuration details are not provided here.

                    ----End

3.6.11 Verifying the Configuration

Procedure
                    ●   Run the display ip vpn-instance vpn-instance-name command to check brief
                        information about a specified VPN instance.
                    ●   Run the display ip vpn-instance verbose vpn-instance-name command to
                        check detailed information about a specified VPN instance.

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                      66
VPN Configuration
VPN Configuration                                                                     3 IPv4 L3VPN Configuration


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

3.6.12 Example for Configuring Basic IPv4 L3VPN over MPLS
Networking Requirements
                    On the network shown in Figure 3-11:
                    ●    CE1 and CE3 belong to vpna.
                    ●    CE2 and CE4 belong to vpnb.
                    ●    The VPN targets of vpna are 111:1, and those of vpnb are 222:2.
                    Users in the same VPN can communicate with each other, but users in different
                    VPNs cannot.

                    Figure 3-11 IPv4 L3VPN networking
                         NOTE

                    In this example, interface 1, interface 2, and interface 3 represent VLANIF 100, VLANIF 200, and
                    VLANIF 300, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       67
VPN Configuration
VPN Configuration                                                             3 IPv4 L3VPN Configuration




Precautions
                    Note the following during the configuration:

