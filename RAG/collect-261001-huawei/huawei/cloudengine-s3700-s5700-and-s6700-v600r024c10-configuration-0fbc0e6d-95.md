---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-95
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [13295, 13489]
sha256: ec6ead9e6b2e26b9f6c9426b53cb84aca4fab9a8bcab168306a6b32fb4e4bd56
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                        #
                        static-cr-lsp ingress tunnel-interface Tunnel10 destination 3.3.3.9 nexthop 10.1.1.2 out-label 20
                        bandwidth ct0 0
                        #
                        static-cr-lsp egress Tunnel20 incoming-interface Vlanif100 in-label 130
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel10
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te signal-protocol cr-static
                         mpls te reverse-lsp protocol static lsp-name Tunnel20
                         mpls te tunnel-id 100
                        #
                        return


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                 226
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         mpls
                         mpls te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                         mpls te
                        #
                        static-cr-lsp transit Tunnel10 incoming-interface Vlanif100 in-label 20 nexthop 10.1.2.2 out-label 30
                        bandwidth ct0 0
                        #
                        static-cr-lsp transit Tunnel20 incoming-interface Vlanif200 in-label 120 nexthop 10.1.1.1 out-label 130
                        bandwidth ct0 0
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        return
                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                        #
                        static-cr-lsp ingress tunnel-interface Tunnel20 destination 1.1.1.9 nexthop 10.1.2.1 out-label 120
                        bandwidth ct0 0
                        #
                        static-cr-lsp egress Tunnel10 incoming-interface Vlanif200 in-label 30
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        interface Tunnel20
                         ip address unnumbered interface LoopBack1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                  227
MPLS Configuration
MPLS Configuration                                                                         4 MPLS TE Configuration

                         tunnel-protocol mpls te
                         destination 1.1.1.9
                         mpls te signal-protocol cr-static
                         mpls te reverse-lsp protocol static lsp-name Tunnel10
                         mpls te tunnel-id 200
                        #
                        return



4.7 Configuring Dynamic MPLS TE Tunnels

4.7.1 Enabling MPLS TE and RSVP-TE

Prerequisites
                 Before configuring dynamic MPLS TE tunnels, complete the following task:

                 ●      Configure OSPF or IS-IS to implement device connectivity at the network
                        layer.


Context
                 Dynamic MPLS TE tunnels are established through dynamic CR-LSPs, with
                 resources and labels dynamically reserved and allocated using RSVP-TE. Dynamic
                 MPLS TE tunnels can dynamically adapt to network changes without necessitating
                 hop-by-hop configuration, making them well suited for large-scale networks.

                 Enabling MPLS TE and RSVP-TE is a prerequisite for configuring a dynamic MPLS
                 TE tunnel. Perform the following configuration on each node of an MPLS TE
                 tunnel.

                         NOTE

                        ● If MPLS TE is disabled in the MPLS view, it is also disabled in all interface views. As a
                          result, all CR-LSPs go down, and all configurations related to the CR-LSPs are deleted.
                        ● If MPLS TE is disabled in an interface view, all CR-LSPs on the interface go down.
                        ● If RSVP-TE is disabled on a device, it is also disabled on all interfaces of the device.


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Configure an LSR ID for the local node.
                 mpls lsr-id lsr-id

                 When configuring an LSR ID, note the following:
                 ●      LSR IDs must be set before you run other MPLS commands.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      Using the address of a loopback interface as the LSR ID is recommended.

         Step 3 Enable MPLS and enter the MPLS view.
                 mpls


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          228
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


         Step 4 Enable MPLS TE globally on the local node.
                 mpls te

                 To enable MPLS TE on an interface, you must first enable MPLS TE globally in the
                 MPLS view.

         Step 5 Enable RSVP-TE on the node.
                 mpls rsvp-te

         Step 6 Return to the system view.
                 quit

