---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-137
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [19457, 19600]
sha256: 3cdf960b90c30242b13b0de178726085fe77c1d290478533dccf8db259c08695
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                 Rules for automatically selecting a router ID for a BGP VPN instance IPv6 address
                                 family are as follows:
                                 ● If loopback interfaces configured with IP addresses are bound to a VPN
                                   instance enabled with the IPv6 address family, the largest IP address among
                                   them is selected as the router ID.
                                 ● If no loopback interfaces configured with IP addresses are bound to a VPN
                                   instance enabled with the IPv6 address family, the largest IP address among
                                   those of other interfaces bound to the VPN instance is selected as the router
                                   ID, regardless of whether the interface is up or down.
                    ●   Configure a router ID for a specified BGP VPN instance IPv6 address family.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the BGP view.
                             bgp as-number

                        c.   Enter the BGP VPN instance IPv6 address family view.
                             ipv6-family vpn-instance vpn-instance-name


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    308
VPN Configuration
VPN Configuration                                                                       4 IPv6 L3VPN Configuration


                        d.    Manually configure or automatically select a router ID for the BGP VPN
                              instance IPv6 address family.
                              router-id { ipv4-address | auto-select }

                    ----End


4.9 Configuring IPv6 MCE

4.9.1 Configuring BGP4+ Between an MCE and a PE

Prerequisites
                    Before configuring MCE, complete the following tasks:
                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 4.5.1 Configuring an IPv6 VPN Instance on a PE.
                    ●   Configure link and network layer protocols for LAN interfaces, and connect
                        the LAN interface for each type of service to the MCE.
                    ●   Bind the MCE's interfaces and the PE's interface connected to the MCE to VPN
                        instances and configure IP addresses for these interfaces. For details, see 4.5.2
                        Binding an Interface to the IPv6 VPN Instance.

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6730E-H-V2, S6730-H-V2, S5732-H-V2, S6750E-S, S6750-S,
                        S5755-S, S5755E-H, S5755-H, S5735I-S-V2, S5735I-H-V2, S5735R-S-V2, and S5735-S-V2
                        series support BGP4+.


Procedure
                    ●   Configure the PE.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        d.    Configure the MCE as a VPN peer for the PE.
                              peer ipv4-address as-number as-number

                        e.    (Optional) Set the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

                              This step is mandatory if the PE is not directly connected to the MCE but
                              an EBGP peer relationship needs to be established between them.
                              In most cases, a directly connected physical link must be available
                              between EBGP peers. If you want to establish an EBGP peer relationship

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   309
VPN Configuration
VPN Configuration                                                                        4 IPv6 L3VPN Configuration


                              between indirectly connected peers, run the peer ebgp-max-hop
                              command to set the maximum number of hops allowed for the TCP
                              connection.

                              The default hop-count value is 255. If the maximum number of hops is
                              set to 1, the PE can establish an EBGP connection with only a directly
                              connected peer.
                    ●   Configure the MCE.
                        a.    Enter the system view.
                              system-view

                        b.    Enter the BGP view.
                              bgp as-number

                        c.    Enter the BGP VPN instance IPv6 address family view.
                              ipv6-family vpn-instance vpn-instance-name

                        d.    Configure the PE as a VPN peer for the MCE.
                              peer ipv4-address as-number as-number

                        e.    (Optional) Set the maximum number of hops allowed for an EBGP
                              connection.
                              peer { ipv4-address | group-name } ebgp-max-hop [ hop-count ]

                              This step is mandatory if the PE is not directly connected to the MCE but
                              an EBGP peer relationship needs to be established between them.

                              In most cases, a directly connected physical link must be available
                              between EBGP peers. If you want to establish an EBGP peer relationship
                              between indirectly connected peers, run the peer ebgp-max-hop
                              command to set the maximum number of hops allowed for the TCP
                              connection.

                              The default hop-count value is 255. If the maximum number of hops is
                              set to 1, the MCE can establish an EBGP connection with only a directly
                              connected peer.
                        f.    (Optional) Configure the MCE to import the direct routes in the site
                              where it resides into its IPv6 VPN instance routing table.
                              import-route direct [ med med-value | route-policy route-policy-name ] *

                    ----End

Verifying the Configuration
                    Run the display ipv6 routing-table vpn-instance vpn-instance-name [ verbose ]
                    command on the MCE to check the IPv6 routing table of the VPN instance.

4.9.2 Configuring Static Routes Between an MCE and a PE

Prerequisites
                    Before configuring MCE, complete the following tasks:

                    ●   Configure a VPN instance for each service on the MCE and its connected PE.
                        For details, see 4.5.1 Configuring an IPv6 VPN Instance on a PE.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    310
VPN Configuration
VPN Configuration                                                                           4 IPv6 L3VPN Configuration


