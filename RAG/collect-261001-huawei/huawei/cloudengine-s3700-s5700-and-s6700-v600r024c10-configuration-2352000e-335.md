---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-335
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [50080, 50220]
sha256: 0cbfaa6008fc083aa9d7d89fd20afc86f994c1bf278e2e19821fc65798eab69f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●   If the tunnel binding policy does not designate any TE tunnels for a
                        destination IP address, an available tunnel is selected based on the default
                        tunnel policy.
                    ●   If the tunnel binding policy designates several TE tunnels for a destination IP
                        address and more than one designated TE tunnel is available, one of the
                        available TE tunnels is selected.
                    ●   If the tunnel binding policy designates several TE tunnels for the destination
                        IP address but none of the designated TE tunnels is available, tunnel selection
                        is determined by the down-switch attribute. If the down-switch attribute is
                        not configured, no tunnels are selected. If the down-switch attribute is
                        configured, an available tunnel is selected based on the default tunnel policy.

                    Comparison of Tunnel Policies

                    Table 7-1 Comparison of tunnel policies

                     Policy                 Description

                     Tunnel type            Cannot ensure which tunnel is selected if there are several
                     prioritizing policy    tunnels of the same type.

                     Tunnel binding         Accurately defines which TE tunnel can be used, ensuring
                     policy                 QoS.




7.4.2 Configuring a Tunnel Policy

Prerequisites
                    Before configuring a tunnel type prioritizing policy, you need to set up a basic
                    network and create tunnels required by services.

Context
                         NOTE

                        Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2, S5755E-H,
                        S5755-H, and S5732-H-V2 series support this function.

                    By default, a VPN instance uses LSPs for service transmission on the backbone
                    network. To use other types of tunnels or configure load balancing for service

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 806
VPN Configuration
VPN Configuration                                                             7 Tunnel Management Configuration


                    transmission of the VPN instance, you need to apply a tunnel policy to the VPN
                    instance.


Procedure
                    ●   Configure a tunnel type prioritizing policy.
                        a.   Enter the system view.
                             system-view

                        b.   Create a tunnel policy and enter the tunnel policy view.
                             tunnel-policy policy-name

                        c.   (Optional) Configure a description for the tunnel policy.
                             description description-text

                             The tunnel policy description helps users memorize the tunnel policy.
                        d.   Configure the sequence in which types of tunnels are selected and the
                             maximum number of tunnels that can participate in load balancing.

                             ▪      For an IPv4 network, run the following command:
                                    tunnel select-seq { ldp | cr-lsp } * load-balance-number load-balance-number [ unmix ]

                                    After the command is run, the system selects tunnels of different
                                    types in the specified sequence. If tunnels that have higher priorities
                                    are unreachable, the system will continue to select tunnels that have
                                    lower priorities in the specified sequence.
                                    If unmix is configured, only one type of tunnel can be selected. For
                                    example, after the tunnel select-seq cr-lsp ldp load-balance-
                                    number 3 unmix command is run for the tunnel policy:
                                    ○    If three or more CR-LSPs are available, the system randomly
                                         selects three of them for service transmission.
                                    ○    If less than three CR-LSPs are available, the system selects only
                                         these available CR-LSPs, but does not select other types of
                                         tunnels.
                    ●   Configure a tunnel binding policy.
                        a.   Enter the system view.
                             system-view

                        b.   Create a tunnel interface and enter the tunnel interface view.
                             interface tunnel interface-number

                        c.   Enable the binding capability of TE tunnels.
                             mpls te reserved-for-binding

                        d.   Return to the system view.
                             quit

                        e.   Create a tunnel policy.
                             tunnel-policy policy-name

                        f.   Bind an MPLS TE tunnel to the tunnel binding policy.
                             tunnel binding destination dest-ip-address te { tunnel-name | tunnel-type tunnel-number }
                             &<1-32> [ ignore-destination-check ] [ down-switch | include-ldp ]
                             tunnel binding destination dest-ip-address auto-tunnel { auto-tunnel-name } &<1-32>
                             [ ignore-destination-check ] [ down-switch ]



Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         807
VPN Configuration
VPN Configuration                                                         7 Tunnel Management Configuration


                                    NOTE

                                  ● If the down-switch parameter is specified in the command, the system
                                    selects an LSP or CR-LSP in descending order of priority for VPN data
                                    transmission when the bound MPLS TE tunnel is unavailable.
                                  ● If the device has multiple peers, you can run the tunnel binding destination
                                    command several times with different destination addresses specified.
                                      If a destination IP address is not bound to a tunnel using the tunnel binding
                                      destination command, the tunnel management module searches for a tunnel
                                      based on the default tunnel policy. By default, routes can recurse only to an
                                      LSP. To change the default type of tunnel to which routes recurse, run the
                                      tunnel policy binding-default down-switch enable command to configure
                                      a tunnel policy to select a tunnel from available ones in descending order of
                                      priority: LSP > CR-LSP.

                    ----End

7.4.3 Applying a Tunnel Policy
Context
                    Only after a tunnel policy is applied, the system can select tunnels for services
                    based on the tunnel policy.
                    The tunnel policy application method varies according to the VPN type. Determine
                    which method to use based on the VPN type.

Procedure
                    ●   Apply a tunnel policy to an IPv4 L3VPN.
                        a.    Enter the system view.
                              system-view

