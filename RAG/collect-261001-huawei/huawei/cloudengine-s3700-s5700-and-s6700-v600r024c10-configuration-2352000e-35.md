---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-35
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3999, 4115]
sha256: 4a090b3f26fe868bcb7ddbc2f21f26065cf2b5ac3da6cb9b933cd5069471c229
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                  A VPN that receives routes outside of it from devices other than PEs and
                                  advertises these routes to PEs is called a transit VPN. A VPN that receives only
                                  routes in it and routes advertised by PEs is called a stub VPN. Generally, a static
                                  route is only used for route exchange between the CE and PE in a stub VPN.
                    ●   Configure a static route on the CE. The configuration details are not provided
                        here. For details, see Configuring an IPv4 Static Route.
                    ----End

3.6.8 Configuring OSPF Between the PE and CE
Procedure
                    ●   Perform the following steps on the PE.
                        a.    Enter the system view.
                              system-view
                        b.    Create an OSPF instance between the PE and CE and enter the OSPF
                              view.
                              ospf process-id [ router-id router-id ] vpn-instance vpnname

                              An OSPF process can be bound to only one VPN instance. If an OSPF
                              process is not bound to any VPN instance before it is started, this process
                              becomes a public network process and cannot be bound to a VPN
                              instance later.
                              A router ID needs to be specified when an OSPF process is started after it
                              is bound to a VPN instance. The router ID must be different from the
                              public network router ID configured in the system view. If the router ID is
                              not specified, OSPF selects the IP address of one of the interfaces bound
                              to the VPN instance as the router ID based on a certain rule.
                        c.    (Optional) Configure a domain ID.
                              domain-id domain-id [ secondary ]

                              The domain ID can be an integer or in dotted decimal notation.
                              Two domain IDs can be configured for each OSPF process. Different
                              processes can have the same domain ID. There are no restrictions on the
                              domain IDs of the OSPF processes for different VPNs on a PE. The OSPF
                              processes of the same VPN must be configured with the same domain ID
                              to ensure correct route advertisement.
                              The domain ID of an OSPF process is contained in the routes generated
                              by the process. When OSPF routes are imported into the BGP routing
                              table, the domain ID is added to BGP VPN routes and advertised as a BGP
                              extended community attribute.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         63
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration


                        d.   (Optional) Configure a VPN route tag.
                             route-tag tag
                        e.   (Optional) Disable the device from setting the DN bit in LSAs.
                             dn-bit-set disable { summary | ase | nssa }

                             To prevent routing loops, an OSPF multi-instance process uses a bit as a
                             flag bit, which is called the DN bit. Perform this step if the device does
                             not need to set the DN bit. By default, a device sets the DN bit in LSAs.
                        f.   (Optional) Disable the device from checking the DN bit in LSAs.
                             dn-bit-check disable { ase | nssa | summary [ router-id router-id ] }

                             To prevent routing loops, an OSPF multi-instance process uses a bit as a
                             flag bit, which is called the DN bit. Perform this step if the device does
                             not need to check the DN bit. By default, a device checks the DN bit in
                             LSAs.
                        g.   (Optional) Set the route type of the extended community attribute of
                             OSPF VPN to 0x8000.
                             eca-route-type compatible

                             In a VPN scenario with devices that support standard protocols, you can
                             set the route type of the extended community attribute of OSPF VPN to
                             0x0306 and configure the devices to identify both 0x0306 and 0x8000
                             route types. If there is a device that does not support standard protocols,
                             you can set the route type of the extended community attribute of OSPF
                             VPN to 0x8000 and configure the device to identify only the 0x8000 route
                             type.
                             Perform this step if the device needs to correctly identify the route type
                             of the extended community attribute of OSPF VPN so that it can
                             communicate with other devices.
                             By default, the route type of the extended community attribute of OSPF
                             VPN is 0x0306.
                        h.   Import BGP routes.
                             import-route bgp [ cost cost | route-policy route-policy-name | tag tag | type type ] *
                        i.   Enter the OSPF area view.
                             area area-id
                        j.   Enable OSPF on the network segment where an interface bound to the
                             VPN instance resides.
                             network address wildcard-mask

                             A network segment can belong to only one area. In other words, you
                             need to specify an area for each interface running OSPF.
                             OSPF can properly run on an interface only when both of the following
                             conditions are met:

                             ▪      The IP address mask length of the interface is longer than or equal
                                    to that specified in the network command.

                             ▪      The interface's primary IP address belongs to the network segment
                                    specified in the network command.
                        k.   Return to the OSPF view.
                             quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          64
VPN Configuration
VPN Configuration                                                                            3 IPv4 L3VPN Configuration


                        l.    Return to the system view.
                              quit
                        m. Enter the BGP view.
                              bgp as-number
                        n.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name
                        o.    Import OSPF routes into the routing table of the BGP VPN instance IPv4
                              address family.
                              import-route ospf process-id [ med med | route-policy route-policy-name ] *

                                      NOTE

