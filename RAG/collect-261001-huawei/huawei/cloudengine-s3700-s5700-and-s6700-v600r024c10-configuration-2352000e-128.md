---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-128
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18108, 18252]
sha256: 6a6854bac83121efd15071ca40e9876ea3078b4ba042d9af7eaae8b82c6f94df
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        m. Enter the BGP view.
                             bgp as-number

                        n.   Enter the BGP-VPN instance IPv6 address family view.
                             ipv6-family vpn-instance vpn-instance-name

                        o.   Import OSPFv3 routes into the routing table for the BGP-VPN instance
                             IPv6 address family.
                             import-route ospfv3 process-id [ med med | route-policy route-policy-name ] *


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          287
VPN Configuration
VPN Configuration                                                                           4 IPv6 L3VPN Configuration


                        p.    (Optional) Run either of the following commands to configure the device
                              to advertise specific routes in a BGP-VPN routing table to a BGP-VPNv6
                              routing table:

                              ▪     Configure the device to advertise only valid routes in a BGP-VPN
                                    routing table to a BGP-VPNv6 routing table.
                                    advertise valid-routes

                                    By default, the device advertises all routes in the BGP-VPN routing
                                    table to the BGP-VPNv6 routing table. The advertise valid-routes
                                    command allows the device to advertise only valid routes to the
                                    BGP-VPNv6 routing table.
                    ●   Configure OSPFv3 on the CE. The configuration details are not provided here.
                    ----End

4.6.7 Configuring IS-ISv6 Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE:
                        a.    Enter the system view.
                              system-view

                        b.    Create an IS-IS instance between the PE and CE and enter the IS-IS view.
                              isis process-id vpn-instance vpn-instance-name

                              An IS-IS process can be bound to only one VPN instance. If an IS-IS
                              process is not bound to any VPN instance before it is started, this process
                              becomes a public network process and cannot be bound to a VPN
                              instance later.
                              If only one IS-IS process, either a public network process or a multi-
                              instance process, runs on the device, you do not need to specify process-
                              id in the command. The value of process-id is 1 by default.

                                     NOTE

                                   If an IS-IS multi-instance process is deleted, IS-IS will be disabled on all the
                                   interfaces in the process.
                                   Deleting a VPN instance or disabling a VPN instance IPv6 address family will
                                   delete all the IS-IS processes bound to the VPN instance or the VPN instance IPv6
                                   address family.
                        c.    Configure a network entity title (NET).
                              network-entity net-addr

                              A NET specifies the current IS-IS area address and the system ID of the
                              device.
                        d.    (Optional) Configure the level of the device.
                              is-level { level-1 | level-1-2 | level-2 }

                        e.    Enable IPv6 for the IS-IS process.
                              ipv6 enable [ topology { compatible [ enable-mt-spf ] | ipv6 | standard } ]

                        f.    Import BGP routes.
                              ipv6 import-route bgp inherit-cost [ tag tag | route-policy route-policy-name | [ level-1 |
                              level-2 | level-1-2 ] ] *


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                 288
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration


                        g.    Return to the system view.
                              quit

                        h.    Enter the interface view.
                              interface interface-type interface-number

                        i.    Enable IS-ISv6 on the interface.
                              isis ipv6 enable [ process-id ]

                        j.    Return to the system view.
                              quit

                        k.    Enter the BGP view.
                              bgp as-number

                        l.    Enter the BGP-VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        m. Import IS-IS routes into the routing table for the BGP-VPN instance IPv6
                           address family.
                              import-route isis process-id [ med med | route-policy route-policy-name ] *

                        n.    (Optional) Run either of the following commands to configure the device
                              to advertise specific routes in a BGP-VPN routing table to a BGP-VPNv6
                              routing table:

                              ▪      Configure the device to advertise only valid routes in a BGP-VPN
                                     routing table to a BGP-VPNv6 routing table.
                                     advertise valid-routes

                                     By default, the device advertises all routes in the BGP-VPN routing
                                     table to the BGP-VPNv6 routing table. The advertise valid-routes
                                     command allows the device to advertise only valid routes to the
                                     BGP-VPNv6 routing table.
                    ●   Configure IS-ISv6 on the CE. The configuration details are not provided here.

                    ----End

4.6.8 Configuring RIPng Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE:
                        a.    Enter the system view.
                              system-view

                        b.    Create a RIPng instance running between the PE and CE and enter the
                              RIPng view.
                              ripng [ process-id ] vpn-instance vpn-instance-name

                              A RIPng multi-instance process can be bound to only one VPN instance. If
                              a RIPng process is not bound to any VPN instance before it is started, this
                              process becomes a public network process.

                              If only one RIPng process, either a public network process or multi-
                              instance process, runs on the device, you do not need to specify process-
                              id in the command. The value of process-id is 1 by default.
                        c.    Import BGP routes.
                              import-route bgp [ cost cost | route-policy route-policy-name ] *


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                   289
VPN Configuration
VPN Configuration                                                                        4 IPv6 L3VPN Configuration


