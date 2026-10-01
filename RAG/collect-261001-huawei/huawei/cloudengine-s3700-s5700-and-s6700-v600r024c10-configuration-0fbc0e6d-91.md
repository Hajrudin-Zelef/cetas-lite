---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-91
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12688, 12831]
sha256: a913158df7a99831c46a0b5ab1c6824b4aebdf8897b69fb75f60a0e7c3fe1639
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        b.   Configure the egress for the static CR-LSP.
                             static-cr-lsp egress lsp-name-val incoming-interface interface-type interface-number in-label
                             in-label [ lsrid ingress-lsr-id tunnel-id tunnel-id ]

                             To modify the incoming-interface interface-type interface-number and
                             in-label in-label-value values, run the static-cr-lsp egress command to
                             set new values. There is no need to run the undo static-cr-lsp egress
                             command to delete the original settings.
                             lsp-name must be unique on the transit node. To simply put, you can
                             specify the MPLS TE tunnel interface name of the static CR-LSP, for
                             example, Tunnel1, as the LSP name.
                 ----End

4.5.4 Verifying the Configuration
Procedure
                 ●      Run the display interface tunnel [ interface-number ] command to check the
                        operating status of tunnel interfaces.
                 ●      Run the display mpls static-cr-lsp [ lsp-name ] [ verbose ] command to
                        check static CR-LSP information.
                 ●      Run the display mpls te tunnel [ verbose ] command to check MPLS TE
                        tunnel information.
                 ●      Run the display mpls te tunnel statistics or display mpls lsp statistics
                        command to check the MPLS TE tunnel statistics.
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the ingress of a tunnel.
                 ----End

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                             216
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration


4.5.5 Example for Configuring Static MPLS TE Tunnels
Networking Requirements
                 On the network shown in Figure 4-5, where all the nodes support MPLS,
                 configure two static MPLS TE tunnels: one from LSR1 to LSR3 and one from LSR3
                 to LSR1.

                 Figure 4-5 Network diagram of static MPLS TE tunnels




Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Configure IP addresses for interfaces, including the loopback interfaces whose
                        addresses are to be used as MPLS LSR IDs.
                 2.     Configure LSR IDs and enable MPLS and MPLS TE globally on the nodes and
                        on interfaces.
                 3.     Create a tunnel interface on the ingress of each tunnel. Set a tunnel IP
                        address, tunneling protocol, destination address, tunnel ID, and signaling
                        protocol used to establish the tunnel.
                 4.     Configure a static CR-LSP for each tunnel. Configure the next-hop address and
                        outgoing label on the ingress, the incoming interface, next-hop address, and
                        outgoing label on the transit node, and the incoming label and inbound
                        interface on the egress to set up the LSP.

                         NOTE

                        ● The outgoing label value on a node should be equal to the incoming label value on its
                          next hop.
                        ● The tunnel-name value specified in the static-cr-lsp ingress { tunnel-interface tunnel
                          interface-number | tunnel-name } destination destination-address { nexthop next-hop-
                          address | outgoing-interface interface-type interface-number } * out-label out-
                          labeltunnel-name command for configuring a CR-LSP ingress should be the same as the
                          tunnel interface name configured using the interface tunnel interface-number
                          command. The tunnel-name value is case sensitive and does not support spaces. For
                          example, if a tunnel interface is created using the interface tunnel 1 command, the
                          tunnel name is Tunnel1. In this case, you must set tunnel-name to Tunnel1 in the CR-
                          LSP ingress configuration command. Otherwise, the tunnel fails to be set up. This
                          restriction does not apply to transit nodes or egresses.


Procedure
         Step 1 Configure interface IP addresses for the nodes.

                 # Configure LSR1.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   217
MPLS Configuration
MPLS Configuration                                                                          4 MPLS TE Configuration

                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 2 Configure LSR IDs and enable MPLS and MPLS TE globally on the nodes and on
                interfaces.
                 # Configure LSR1.
                 [LSR1] mpls lsr-id 1.1.1.9
                 [LSR1] mpls
                 [LSR1-mpls] mpls te
                 [LSR1-mpls] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls
                 [LSR1-Vlanif100] mpls te
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 3 Configure MPLS TE tunnel interfaces.
                 # On LSR1, configure an MPLS TE tunnel to LSR3.
                 [LSR1] interface Tunnel 10
                 [LSR1-Tunnel10] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel10] tunnel-protocol mpls te
                 [LSR1-Tunnel10] destination 3.3.3.9
                 [LSR1-Tunnel10] mpls te tunnel-id 100
                 [LSR1-Tunnel10] mpls te signal-protocol cr-static
                 [LSR1-Tunnel10] quit

                 # On LSR3, configure an MPLS TE tunnel to LSR1.
                 [LSR3] interface Tunnel 20
                 [LSR3-Tunnel20] ip address unnumbered interface loopback 1
                 [LSR3-Tunnel20] tunnel-protocol mpls te
                 [LSR3-Tunnel20] destination 1.1.1.9
                 [LSR3-Tunnel20] mpls te tunnel-id 200
                 [LSR3-Tunnel20] mpls te signal-protocol cr-static
                 [LSR3-Tunnel20] quit

         Step 4 Configure a static CR-LSP from LSR1 to LSR3.
                 # Configure LSR1 as the ingress of the static CR-LSP.
                 [LSR1] static-cr-lsp ingress tunnel-interface Tunnel10 destination 3.3.3.9 nexthop 10.1.1.2 out-label 20

                 # Configure LSR2 as the transit node of the static CR-LSP.
                 [LSR2] static-cr-lsp transit Tunnel10 incoming-interface vlanif 100 in-label 20 nexthop 10.1.2.2 out-
                 label 30

                 # Configure LSR3 as the egress of the static CR-LSP.

