---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-90
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12556, 12687]
sha256: 64a46248f99f02500d78a8cced032ead3cd7618d3f8d16b2084356ab6b7087d6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Perform the following configuration on the ingress of an MPLS TE tunnel.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Create a tunnel interface and enter the tunnel interface view.
                 interface tunnel interface-number

         Step 3 Use one of the following methods to configure an IP address for the tunnel
                interface.
                 ●      Configure an IP address for the tunnel interface.
                        ip address ip-address { mask | mask-length } [ sub ]

                        Configure a primary IP address for the tunnel interface before assigning a
                        secondary IP address.
                 ●      Configure the tunnel interface to borrow the IP address of another interface.
                        ip address unnumbered interface interface-type interface-number

                              NOTE

                             A TE tunnel can be established on a tunnel interface without an IP address configured.
                             However, an IP address must be configured for the tunnel interface before it can
                             forward traffic over the TE tunnel. An MPLS TE tunnel is unidirectional and does not
                             involve peer address configuration. Therefore, you are advised to configure the tunnel
                             interface to borrow the ingress LSR ID as its IP address, rather than assigning a unique
                             IP address to the interface.

         Step 4 Configure MPLS TE as the tunneling protocol.
                 tunnel-protocol mpls te

         Step 5 Configure a destination address for the tunnel.
                 destination ip-address

                 Generally, the egress LSR ID is used as the tunnel destination address. Different
                 tunnel types require different destination addresses. Changing the tunneling
                 protocol to MPLS TE automatically removes the original destination address of the
                 tunnel. As a result, you need to reconfigure the destination address for the tunnel.

         Step 6 Configure a tunnel ID.
                 mpls te tunnel-id tunnel-id

         Step 7 Configure a static CR-LSP for the tunnel.
                 mpls te signal-protocol cr-static

         Step 8 (Optional) Configure a tunnel name.
                 mpls te signalled tunnel-name signalled-tunnel-name


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      214
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


                 Generally, a tunnel interface name is used as a tunnel name. For example, the
                 tunnel interface name Tunnel10 can also be used as the tunnel name. This step is
                 performed for the following purposes:
                 ●      Facilitate tunnel management.
                 ●      Allow the device to interwork with a non-Huawei device that does not uses a
                        tunnel interface name as the tunnel name.
                 ----End

4.5.3 Configuring Static CR-LSPs
Context
                 The configuration for a static CR-LSP varies according to the roles of devices on
                 the static CR-LSP.
                 ●      Ingress: Configure the ingress LSP forwarding entry and bind the LSP to a TE
                        tunnel interface.
                 ●      Transit node: Configure the transit LSP forwarding entry. This configuration is
                        not involved if the static CR-LSP does not have a transit node.
                 ●      Egress: Configure the egress LSP forwarding entry.
                 Perform the following configuration based on the roles of devices on an MPLS TE
                 tunnel.

Procedure
                 ●      Perform the following configuration on the ingress.
                        a.   Enter the system view.
                             system-view
                        b.   Configure the ingress for a static CR-LSP.
                             static-cr-lsp ingress { tunnel-interface tunnel interface-number | tunnel-name } destination
                             destination-address { nexthop next-hop-address | outgoing-interface interface-type interface-
                             number } * out-label out-label [ bandwidth [ ct0 ] bandwidth ]

                             To modify the destination destination-address, nexthop next-hop-
                             address, outgoing-interface interface-type interface-number, and out-
                             label out-label values, run the static-cr-lsp ingress command to set new
                             values. There is no need to run the undo static-cr-lsp ingress command
                             to delete the original settings.
                             Set tunnel interface-number to the MPLS TE tunnel interface number of
                             the static CR-LSP. The default bandwidth is 0. The bandwidth assigned to
                             the tunnel should not exceed the maximum reservable bandwidth of the
                             link.
                             A next hop or an outbound interface is determined based on the route
                             from the ingress to the egress. For the differences between a next hop
                             address and an outbound interface, see the section "Understanding IPv4
                             Static Routes" in CLI Configuration Guide - IPv4 Static Route
                             Configuration.
                             If an Ethernet interface is used as an LSP outbound interface, you must
                             specify the nexthop next-hop-address parameter to ensure normal traffic
                             forwarding on the LSP.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                             215
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


                 ●      Perform the following configuration on the transit node.
                        a.   Enter the system view.
                             system-view

                        b.   Configure the transit node for the static CR-LSP.
                             static-cr-lsp transit lsp-name incoming-interface interface-type interface-number in-label in-
                             label { nexthop next-hop-address | outgoing-interface interface-type interface-number } * out-
                             label out-label [ ingress-lsrid ingress-lsrid egress-lsrid egress-lsrid tunnel-id tunnel-id ]
                             [ bandwidth [ ct0 ] bandwidth ]

                             To modify all the parameters except lsp-name, run the static-cr-lsp
                             transit command to set new values. There is no need to run the undo
                             static-cr-lsp transit command to delete the original settings.
                             lsp-name must be unique on the transit node. To simply put, you can
                             specify the MPLS TE tunnel interface name of the static CR-LSP, for
                             example, Tunnel1, as the LSP name.
                             If an Ethernet interface is used as an LSP outbound interface, you must
                             specify the nexthop next-hop-address parameter to ensure normal traffic
                             forwarding on the LSP.
                 ●      Perform the following configuration on the egress.
                        a.   Enter the system view.
                             system-view

