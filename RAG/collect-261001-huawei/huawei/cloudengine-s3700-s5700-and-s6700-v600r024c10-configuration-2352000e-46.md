---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-46
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [5626, 5762]
sha256: aed5f97f2e9f4a151b7f153ebd9c5af4c52413128b73efce64b553e7afd41383
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        c.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        d.    Configure the MCE as a VPN BGP peer for the PE.
                              peer ipv4-address as-number as-number

                        e.    (Optional) Configure the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

                              This step is mandatory if the PE is not directly connected to the MCE but
                              an EBGP peer relationship needs to be established between them.

                              In most cases, a directly connected physical link must be available
                              between EBGP peers. If you want to establish an EBGP peer relationship
                              between indirectly connected peers, run the peer ebgp-max-hop
                              command to set the maximum number of hops allowed for a TCP
                              connection.

                              The default value of hop-count is 255. If the maximum number of hops is
                              set to 1, the MCE can establish an EBGP connection with only a directly
                              connected peer.
                        f.    (Optional) Configure the MCE to import the direct routes in the site
                              where it resides into its IPv4 VPN instance routing table.

                              ▪    Use the import-route command to import routes.
                                   import-route direct [ med med-value | route-policy route-policy-name ] *

                              ▪    Use the network command to import routes.
                                   network ipv4-address [ mask | mask-length ] [ route-policy route-policy-name ]

                    ----End

3.9.3 Configuring Static Routes Between an MCE and a PE

Prerequisites
                    Before configuring MCE, you have completed the following tasks:

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 3.5.1 Configuring an IPv4 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind each MCE interface and the PE interface connecting to the MCE to the
                        VPN instance, and configure IP addresses for the interfaces. For detailed
                        configurations, see 3.5.2 Binding an Interface to an IPv4 VPN Instance.


Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         89
VPN Configuration
VPN Configuration                                                                          3 IPv4 L3VPN Configuration


Procedure
                    ●   Configure the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Configure a static route for a specified VPN instance IPv4 address family.
                              ip route-static vpn-instance vpn-source-name destination-address { mask | mask-length }
                              interface-type interface-number [ nexthop-address ] [ preference preference | tag tag ] *

                        c.    Enter the BGP view.
                              bgp as-number

                        d.    Enter the BGP VPN instance IPv4 address family view.
                              ipv4-family vpn-instance vpn-instance-name

                        e.    Import the configured static route to the routing table of the BGP VPN
                              instance IPv4 address family.
                              import-route static [ med med-value | route-policy route-policy-name ] *

                    ●   Configure the MCE.
                        a.    Enter the system view.
                              system-view

                        b.    Configure a static route for a specified VPN instance IPv4 address family.
                              ip route-static vpn-instance vpn-source-name destination-address { mask | mask-length }
                              interface-type interface-number [ nexthop-address ] [ preference preference | tag tag ] *

                    ----End

3.9.4 Configuring OSPF Between an MCE and a PE

Prerequisites
                    Before configuring MCE, you have completed the following tasks:

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 3.5.1 Configuring an IPv4 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind each MCE interface and the PE interface connecting to the MCE to the
                        VPN instance, and configure IP addresses for the interfaces. For detailed
                        configurations, see 3.5.2 Binding an Interface to an IPv4 VPN Instance.


Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, S5735E-S-V2, S5735-
                        S-V2, S5735R-L-V2, S5735E-L-V2, S5735-L-V2, and S5735I-L-V2 series support OSPF.

                    Deleting a VPN instance or disabling a VPN instance IPv4 address family will
                    delete all the OSPF processes bound to the VPN instance and the VPN instance
                    IPv4 address family.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               90
VPN Configuration
VPN Configuration                                                                           3 IPv4 L3VPN Configuration


Procedure
                    ●   Configure the PE.
                        a.   Enter the system view.
                             system-view
                        b.   Create an OSPF process and enter the OSPF view.
                             ospf process-id [ router-id router-id ] vpn-instance vpnname

                             An OSPF process can be bound to only one VPN instance.
                             A router ID needs to be specified when an OSPF process is started after it
                             is bound to a VPN instance. The router ID must be different from the
                             public network router ID configured in the system view. If the router ID is
                             not specified, OSPF selects the IP address of one of the interfaces bound
                             to the VPN instance as the router ID based on a certain rule.
                        c.   (Optional) Configure a domain ID.
                             domain-id domain-id [ secondary ]

