---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-127
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [17970, 18107]
sha256: 563b5f9bda03de886b8d67ea23419c35f2a4bb9b293daaa0e3251ca8150e309a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        g.    Enable the function to exchange BGP routing information with a BGP
                              IPv6 peer.
                              peer ipv6-address enable

                        h.    Import the routes of the local site.
                              import-route { direct | static | ripng process-id | ospfv3 process-id | isis process-id } [ med
                              med | route-policy route-policy-name ]*

                              Configure the CE to advertise routes imported from the local site to the
                              connected PE. The PE then advertises these routes to the peer CE. The
                              type of route imported in this step may vary according to the networking
                              mode.
                    ----End

4.6.5 Configuring Static Routes Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                     285
VPN Configuration
VPN Configuration                                                                         4 IPv6 L3VPN Configuration


                        a.    Enter the system view.
                              system-view

                        b.    Configure a static route for the IPv6 address family of a specified VPN
                              instance.
                              ipv6 route-static vpn-instance vpn-instance-name dest-ipv6-address prefix-length { interface-
                              type interface-number [ nexthop-ipv6-address ] | vpn-instance vpn-destination-name nexthop-
                              ipv6-address | nexthop-ipv6-address [ public ] } [ preference preference | tag tag ]*
                              [ description text ]

                        c.    Enter the BGP view.
                              bgp as-number

                        d.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        e.    Import the configured static route into the routing table for the BGP VPN
                              instance IPv6 address family.
                              import-route static [ med med | route-policy route-policy-name ] *

                                    NOTE

                                  A VPN that receives routes outside of it from devices other than PEs and
                                  advertises these routes to PEs is called a transit VPN. A VPN that receives only
                                  routes in it and routes advertised by PEs is called a stub VPN. Generally, a static
                                  route is only used for route exchange between the CE and PE in a stub VPN.
                        f.    (Optional) Run either of the following commands to configure the device
                              to advertise specific routes in a BGP VPN routing table to a BGP VPNv6
                              routing table:

                              ▪    Configure the device to advertise only valid routes in a BGP VPN
                                   routing table to a BGP VPNv6 routing table.
                                   advertise valid-routes

                                   By default, the device advertises all routes in the BGP VPN routing
                                   table to the BGP VPNv6 routing table. The advertise valid-routes
                                   command allows the device to advertise only valid routes to the BGP
                                   VPNv6 routing table.
                    ●   Configure a static route on the CE. The configuration details are not provided
                        here. For details, see Configuring an IPv6 Static Route.

                    ----End

4.6.6 Configuring OSPFv3 Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE:
                        a.    Enter the system view.
                              system-view

                        b.    Start OSPFv3 multi-instance and enter the OSPFv3 multi-instance view.
                              ospfv3 [ process-id ] [ vpn-instance vpnname ]

                              An OSPFv3 process can be bound to only one VPN instance. If an OSPFv3
                              process is not bound to any VPN instance before being started, this
                              process becomes a public network process. The public network OSPFv3
                              process cannot be bound to a VPN instance.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                             286
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration


                                     NOTE

                                    Deleting a VPN instance or disabling the IPv6 address family of a VPN instance
                                    will also delete the related OSPFv3 processes.
                        c.   Configure a router ID.
                             router-id router-id

                             A router ID uniquely identifies an OSPFv3 process in an AS. If no router ID
                             is configured, the OSPFv3 process cannot run.
                        d.   (Optional) Configure a domain ID.
                             domain-id { domain-id-int | domain-id-ipaddr }

                             The domain ID can be either an integer or a dotted decimal number.

                             Generally, the routes imported from a PE are advertised as External-LSAs.
                             The routes that belong to different nodes of the same OSPFv3 domain
                             are advertised as Type-3 LSAs (intra-domain routes). This requires that
                             different nodes in the same OSPFv3 domain have the same domain ID.
                        e.   (Optional) Configure a VPN route tag.
                             route-tag tag

                        f.   (Optional) Disable the device from setting the DN bit in LSAs.
                             dn-bit-set disable { summary | ase | nssa }

                             To prevent routing loops, an OSPFv3 multi-instance process uses a bit as
                             a flag bit, which is called the DN bit. Perform this step if the device does
                             not need to set the DN bit. By default, a device sets the DN bit in LSAs.
                        g.   (Optional) Disable the device from checking the DN bit in LSAs.
                             dn-bit-check disable { ase | nssa | summary [ router-id router-id ] }

                             To prevent routing loops, an OSPFv3 multi-instance process uses a bit as
                             a flag bit, which is called the DN bit. Perform this step if the device does
                             not need to check the DN bit. By default, a device checks the DN bit in
                             LSAs.
                        h.   Import BGP routes.
                             import-route bgp [ cost cost | route-policy route-policy-name | tag tag | type type ] *

                        i.   Return to the system view.
                             quit

                        j.   Enter the view of the interface bound to the VPN instance.
                             interface interface-type interface-number

                        k.   Enable OSPFv3 on the interface.
                             ospfv3 process-id area area-id [ instance instance-id ]

                        l.   Return to the system view.
                             quit

