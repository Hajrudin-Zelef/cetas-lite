---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-47
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5763, 5895]
sha256: 4dae85365be31fe45b48b81020b7f9000bd7f79311ce0f44dca2510571873701
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                        d.   Enter the OSPF area view.
                             area area-id
                        e.   Enable OSPF on the network segment where an interface bound to the
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
                             On a loopback interface, by default, OSPF advertises its IP address in the
                             form of a 32-bit host route, independent of the mask length of the IP
                             address on the interface.
                        f.   Return to the OSPF view.
                             quit
                        g.   Import BGP routes.
                             import-route bgp [ cost cost | route-policy route-policy-name | tag tag | type type ] *


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          91
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration


                        h.   Return to the system view.
                             quit

                        i.   Enter the BGP view.
                             bgp as-number

                        j.   Enter the BGP VPN instance IPv4 address family view.
                             ipv4-family vpn-instance vpn-instance-name

                        k.   Import OSPF routes into the routing table of the BGP VPN instance IPv4
                             address family.
                             import-route ospf process-id [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.   Enter the system view.
                             system-view

                        b.   Create an OSPF process and enter the OSPF view.
                             ospf process-id [ router-id router-id ] vpn-instance vpnname

                             An OSPF process can be bound to only one VPN instance. The same OSPF
                             process must be configured on the MCE and its connected PE.

                             A router ID needs to be specified when an OSPF process is started after it
                             is bound to a VPN instance. The router ID must be different from the
                             public network router ID configured in the system view. If the router ID is
                             not specified, OSPF selects the IP address of one of the interfaces bound
                             to the VPN instance as the router ID based on a certain rule.
                        c.   (Optional) Configure a domain ID.
                             domain-id domain-idvalue [ secondary ]

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
                        d.   Disable routing loop detection.
                             vpn-instance-capability simple

                             If OSPF VPN multi-instance has been deployed on the MCE and PE, the
                             PE sends the MCE a link-state advertisement (LSA) with the Down (DN)
                             bit set to 1. Because VPN instances have been configured on the MCE,
                             the MCE has routing loop detection enabled. If the MCE detects that the
                             LSA contains the DN bit with value 1, this LSA cannot be used to
                             calculate routes. Run the vpn-instance-capability simple command to
                             disable OSPF routing loop detection. When OSPF routing loop detection
                             is disabled, the MCE calculates all OSPF routes without checking the DN
                             bit or route tag.
                        e.   Enter the OSPF area view.
                             area area-id


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      92
VPN Configuration
VPN Configuration                                                               3 IPv4 L3VPN Configuration


                        f.    Enable OSPF on the network segment where an interface bound to the
                              VPN instance resides.
                              network address wildcard-mask

                              A network segment can belong to only one area. In other words, you
                              need to specify an area for each interface running OSPF.

                              OSPF can properly run on an interface only when both of the following
                              conditions are met:

                              ▪    The IP address mask length of the interface is longer than or equal
                                   to that specified in the network command.

                              ▪    The interface's primary IP address belongs to the network segment
                                   specified in the network command.

                              On a loopback interface, by default, OSPF advertises its IP address in the
                              form of a 32-bit host route, independent of the mask length of the IP
                              address on the interface.

                    ----End

3.9.5 Configuring IS-IS Between an MCE and a PE

Prerequisites
                    Before configuring MCE, you have completed the following tasks:

