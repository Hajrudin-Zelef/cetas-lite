---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-141
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [20011, 20131]
sha256: c653131f90eccb77fc14cc5f0fa9b1373339771bee4aa91521cfb924bf3d4be9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        e.    Enter the view of the interface connected to the MCE.
                              interface interface-type interface-number


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                            317
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


                        f.    Switch the interface working mode from Layer 2 to Layer 3.
                              undo portswitch

                              Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                              S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                              be switched from Layer 2 mode to Layer 3 mode using the undo
                              portswitch command. Determine whether to perform this step based on
                              the current interface working mode.
                        g.    Enable IPv6 on the interface.
                              ipv6 enable

                        h.    Enable RIPng on the interface.
                              ripng process-id enable

                        i.    Enter the BGP view.
                              bgp as-number

                        j.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        k.    Import RIP routes into the routing table for the BGP VPN instance IPv6
                              address family.
                              import-route rip process-id [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.    Enter the system view.
                              system-view

                        b.    Create a RIP process, bind it to a VPN instance, and enter the RIP view.
                              rip process-id vpn-instance vpn-instance-name

                              A RIP process can be bound to only one VPN instance.

                              If a RIP process is not bound to any VPN instance before it is started, this
                              process becomes a public network process and cannot be bound to a VPN
                              instance later.
                        c.    Return to the system view.
                              quit

                        d.    Enter the view of the interface connected to the PE.
                              interface interface-type interface-number

                        e.    Switch the interface working mode from Layer 2 to Layer 3.
                              undo portswitch

                              Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2,
                              S6750-S, S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can
                              be switched from Layer 2 mode to Layer 3 mode using the undo
                              portswitch command. Determine whether to perform this step based on
                              the current interface working mode.
                        f.    Enable IPv6 on the interface.
                              ipv6 enable

                        g.    Enable RIPng on the interface.
                              ripng process-id enable

                    ----End

Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                   318
VPN Configuration
VPN Configuration                                                                      4 IPv6 L3VPN Configuration


Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ verbose ]
                    command on the MCE to check the IPv6 routing table of the VPN instance.

4.9.6 Example for Configuring an IPv6 MCE

Networking Requirements
                    On the network shown in Figure 4-3, Site 1 and Site 2 both connect to PE1, but
                    the two sites need to be isolated from each other and use independent address
                    spaces. To meet the preceding requirements and reduce costs, use the MCE
                    solution.

                    Figure 4-3 Network diagram for IPv6 MCE
                          NOTE

                    In this example, interface 1, interface 2, interface 3, and interface 4 represent VLANIF 100,
                    VLANIF 200, VLANIF 300, and VLANIF 400, respectively.




Precautions
                    During the configuration, note the following:
                    ●    The MCE must have multiple VPN instances configured and have a different
                         interface bound to each VPN instance.
                    ●    Routing loop detection must be disabled on the MCE so that the MCE
                         exchanges routing information with the PE using the OSPFv3 multi-instance.
                    ●    RIPng must be configured on the MCE to import VPN routes from Site 1 and
                         Site 2.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         319
VPN Configuration
VPN Configuration                                                               4 IPv6 L3VPN Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPN instances on the MCE and PE1 and bind interfaces to the VPN
                         instances.
                    2.   Configure OSPFv3 multi-instance on the MCE and PE1 to exchange VPN
                         routing information.
                    3.   Configure RIPng on the MCE, DeviceA, and DeviceB to exchange VPN routing
                         information.
                    4.   Disable routing loop detection on the MCE and import RIPng routes destined
                         for VPN sites.

