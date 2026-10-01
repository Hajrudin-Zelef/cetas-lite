---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-34
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [3866, 3998]
sha256: d49049b5a3b2b3efebdc102261b72f808137bc1554c4b51aba0a3bc6a212872b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                              This step is mandatory if the PE and CE are not directly connected. In
                              most cases, a directly connected physical link must be available between
                              EBGP peers. If you want to establish an EBGP peer relationship between
                              indirectly connected peers, run the peer ebgp-max-hop command to set
                              the maximum number of hops allowed for a TCP connection.
                        f.    (Optional) Run either of the following commands to import the direct
                              routes destined for the local CE into the corresponding VPN routing table:
                              import-route direct [ med med | route-policy route-policy-name ]
                              network ipv4-address [ mask | mask-length ] [ route-policy route-policy-name ]

                                    NOTE

                                   The PE can automatically learn the direct routes destined for the local CE. The
                                   learned routes take precedence over the direct routes advertised from the local
                                   CE using EBGP. If this step is not performed, the PE does not use MP-BGP to
                                   advertise the direct routes destined for the local CE to the remote PE.
                        g.    (Optional) Configure the device to allow routing loops.
                              peer ipv4-address allow-as-loop [ number ]

                    ●   Perform the following steps on the CE.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Configure a PE as a VPN peer.
                              peer ipv4-address as-number as-number

                        d.    (Optional) Configure the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

                              In most cases, a directly connected physical link must be available
                              between EBGP peers. If you want to establish an EBGP peer relationship
                              between indirectly connected peers, run the peer ebgp-max-hop
                              command to set the maximum number of hops allowed for a TCP
                              connection.
                        e.    Run either of the following commands to configure the CE to advertise its
                              VPN network segment address to the connected PE:
                              import-route { direct | static | rip process-id | ospf process-id | isis process-id } [ med med |
                              route-policy route-policy-name ] *
                              network ipv4-address [ mask | mask-length ]

                    ----End

3.6.6 Configuring IBGP Between the PE and CE
Procedure
                    ●   Perform the following steps on the PE.
                        a.    Enter the system view.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                       61
VPN Configuration
VPN Configuration                                                                              3 IPv4 L3VPN Configuration

                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        d.    Configure a CE as a VPN peer.
                              peer ipv4-address as-number as-number

                        e.    (Optional) Configure the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

                              This step is mandatory if the PE and CE are not directly connected. In
                              most cases, a directly connected physical link must be available between
                              EBGP peers. If you want to establish an EBGP peer relationship between
                              indirectly connected peers, run the peer ebgp-max-hop command to set
                              the maximum number of hops allowed for a TCP connection.
                        f.    (Optional) Run either of the following commands to import the direct
                              routes destined for the local CE into the corresponding VPN routing table:
                              import-route direct [ med med | route-policy route-policy-name ]
                              network ipv4-address [ mask | mask-length ] [ route-policy route-policy-name ]

                                    NOTE

                                   The PE can automatically learn the direct routes destined for the local CE. The
                                   learned routes take precedence over the direct routes advertised from the local
                                   CE using IBGP. If this step is not performed, the PE does not use MP-BGP to
                                   advertise the direct routes destined for the local CE to the remote PE.
                    ●   Perform the following steps on the CE.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Configure a PE as a VPN peer.
                              peer ipv4-address as-number as-number

                        d.    Run either of the following commands to configure the CE to advertise its
                              VPN network segment address to the connected PE:
                              import-route { direct | static | rip process-id | ospf process-id | isis process-id } [ med med |
                              route-policy route-policy-name ] *
                              network ipv4-address [ mask | mask-length ]

                    ----End

3.6.7 Configuring Static Routes Between the PE and CE

Procedure
                    ●   Perform the following steps on the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Configure a static route for a specified VPN instance IPv4 address family.
                              ip route-static vpn-instance vpn-source-name destination-address { mask | mask-length }
                              interface-type interface-number [ nexthop-address ] [ preference preference | tag tag ] *


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                       62
VPN Configuration
VPN Configuration                                                                            3 IPv4 L3VPN Configuration


                        c.    Enter the BGP view.
                              bgp as-number
                        d.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name
                        e.    Import the configured static route to the routing table of the BGP VPN
                              instance IPv4 address family.
                              import-route static [ med med | route-policy route-policy-name ] *

                                    NOTE

