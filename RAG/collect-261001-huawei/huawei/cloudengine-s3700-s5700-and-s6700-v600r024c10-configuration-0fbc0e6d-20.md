---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-20
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [2137, 2273]
sha256: 29a1922b28e121e7df6996d1cb02657152ce68968c67191a2e30e5e64093e776
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                             To modify the incoming-interface in-interface-type in-interface-number,
                             in-label in-label, nexthop next-hop-address, outgoing-interface out-
                             interface-type outinterface-number, and out-label out-label parameter
                             settings, run the static-lsp transit command to set new values directly,
                             not requiring you to clear previous settings using the undo static-lsp
                             transit command.

                                      NOTE

                                     You are advised to specify a next hop for a static LSP. Ensure that the local
                                     routing table contains a routing entry that exactly matches the specified next-
                                     hop IP address.
                                     If an Ethernet interface is used as an outbound interface of an LSP, the nexthop
                                     next-hop-address parameter must be configured to ensure normal traffic
                                     forwarding on the LSP.
                 ●      Perform the following configurations on the egress.
                        a.   Enter the system view.
                             system-view

                        b.   Configure the local node as the egress of the LSP.
                             static-lsp egress lsp-name [ incoming-interface interface-type interface-number ] in-label in-
                             label

                             To modify the incoming-interface interface-type interface-number and
                             in-label in-label parameter settings, run the static-lsp egress command
                             to set new values directly, not requiring you to clear previous settings
                             using the undo static-lsp egress command.

                 ----End

Result
                 ●      Run the display mpls static-lsp [ lsp-name ] [ { include | exclude } ip-
                        address mask-length ] [ verbose ] command to check information about local
                        static LSPs.
                 ●      Run the display mpls lsp protocol static [{ include | exclude } destaddr
                        masklen ] [ incoming-interface in-port-type in-port-num ] [ outgoing-
                        interface out-port-type out-port-num ] [ in-label in-label-value ] [ out-label
                        out-label-value ] [ nexthop nexthopaddr ] [ lsr-role { ingress | transit |
                        egress }] [ verbose ] command to check information about static LSPs.


3.5.3 Example for Configuring a Static LSP
                 This section provides an example for configuring a static LSP.

Networking Requirements
                 On the network shown in Figure 3-5, all nodes support MPLS. OSPF is configured
                 on the MPLS backbone network. Configure a static LSP from LSRA to LSRC to
                 transmit services, such as L2VPN and L3VPN services, on the public network.

                         NOTE

                        Interfaces 1 and 2 in this example represent VLANIF100 and VLANIF200, respectively.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                               37
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


                 Figure 3-5 Configuring a static LSP




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface and configure a loopback address as an
                        LSR ID on each node. Configure OSPF to advertise the route to the network
                        segment connected to each interface and to advertise the host route to each
                        LSR ID.
                 2.     Enable MPLS globally on each node.
                 3.     Enable MPLS on each interface.
                 4.     On the ingress, specify the destination address, next hop, and outgoing label
                        value for the LSP.
                 5.     On each transit node, specify the incoming label value (same as the outgoing
                        label value on the previous node), outbound interface, next hop, and outgoing
                        label value for the LSP.
                 6.     On the egress, specify the inbound interface and incoming label value (same
                        as the outgoing label value on the previous node) for the LSP.

Procedure
         Step 1 Assign an IP address to each interface.
                 Assign an IP address and a mask to each interface (including loopback interfaces)
                 according to Configuration Scripts.
         Step 2 Configure OSPF to advertise the route to the network segment connected to each
                interface and to advertise the host route to each LSR ID.
                 # Configure LSRA.
                 [LSRA] ospf 1
                 [LSRA-ospf-1] area 0
                 [LSRA-ospf-1-area-0.0.0.0] network 192.168.1.9 0.0.0.0
                 [LSRA-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSRA-ospf-1-area-0.0.0.0] quit
                 [LSRA-ospf-1] quit

                 # Configure LSRB.
                 [LSRB] ospf 1
                 [LSRB-ospf-1] area 0
                 [LSRB-ospf-1-area-0.0.0.0] network 192.168.2.9 0.0.0.0
                 [LSRB-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSRB-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.255
                 [LSRB-ospf-1-area-0.0.0.0] quit
                 [LSRB-ospf-1] quit

                 # Configure LSRC.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           38
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration

                 [LSRC] ospf 1
                 [LSRC-ospf-1] area 0
                 [LSRC-ospf-1-area-0.0.0.0] network 192.168.3.9 0.0.0.0
                 [LSRC-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.255
                 [LSRC-ospf-1-area-0.0.0.0] quit
                 [LSRC-ospf-1] quit

                 After completing the configuration, run the display ip routing-table command on
                 each node to check whether all the nodes have learned routes from each other.
                 The following example uses the command output on LSRA.
                 [LSRA] display ip routing-table
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 11        Routes : 11

                 Destination/Mask     Proto Pre Cost        Flags NextHop         Interface

                       10.1.1.0/24 Direct 0 0           D 10.1.1.1        Vlanif100
                       10.1.1.1/32 Direct 0 0           D 127.0.0.1       Vlanif100
                     192.168.1.9/32 Direct 0 0           D 127.0.0.1        LoopBack1
                      10.1.1.255/32 Direct 0 0           D 127.0.0.1       Vlanif100
                       10.2.1.0/24 OSPF 10 2             D 10.1.1.2        Vlanif100
                     192.168.2.9/32 OSPF 10 1             D 10.1.1.2         Vlanif100
                     192.168.3.9/32 OSPF 10 2             D 10.1.1.2         Vlanif100

