---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-93
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [12977, 13154]
sha256: 4dffcab9930b2336a091596708786040aba04f5afb23d53a656a9df4d9bfa794
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

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



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                                  221
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration

                         ip address 3.3.3.9 255.255.255.255
                        #
                        interface Tunnel20
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 1.1.1.9
                         mpls te signal-protocol cr-static
                         mpls te tunnel-id 200
                        #
                        return



4.6 Configuring Static Bidirectional Associated LSPs

4.6.1 Understanding Static Bidirectional Associated LSPs
                 Static bidirectional associated LSPs provide bandwidth protection for bidirectional
                 services on a network and can quickly trigger service switching on both ends.

Purpose
                 An MPLS network faces the following challenges:
                 ●      A static MPLS TE tunnel is unidirectional. The ingress forwards services to the
                        egress along the MPLS TE tunnel. If the egress forwards services to the ingress
                        over an IP route, the services may be congested as bandwidth cannot be
                        reserved on the IP link.
                 ●      If a static MPLS TE tunnel is also established from the egress to the ingress
                        and a tunnel switchover occurs on one of the tunnels, a quick switchover
                        cannot be executed on the other tunnel, resulting in service interruptions.

                 Bidirectional associated LSPs can address the preceding issues. These LSPs are
                 established by binding the ingresses and egresses of two tunnels. They help avoid
                 service congestions and interruptions by allowing bandwidth reservation in both
                 traffic directions and ensuring that either end can promptly respond to a traffic
                 switchover on the other end.

Fundamentals
                 Figure 4-6 shows the setup of static bidirectional associated LSPs. There are two
                 tunnels (Tunnel1 and Tunnel2) in opposite directions. The setup configuration
                 requires the following requirements to be met:
                 ●      Tunnel1 and Tunnel2 must be manually configured static MPLS TE tunnels.
                 ●      The remote LSP is configured as the reverse LSP of the local one by specifying
                        the remote LSP's tunnel ID and ingress LSR ID on the local tunnel interface.
                        This configuration needs to be performed on both ends. A reverse LSP must
                        be specified for both the ingress and egress of a tunnel. In addition, the
                        binding relationships must match each other. On the network shown in
                        Figure 4-6, this means that the reverse LSP's tunnel ID 200 and ingress LSR ID
                        4.4.4.4 must be set on the interface of Tunnel1.
                              NOTE

                            The ingress LSR ID of the reverse LSP is the same as the egress LSR ID of the forward
                            LSP.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    222
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration


                 ●      Penultimate hop popping (PHP) is not configured, as bidirectional associated
                        LSPs do not support this function.
                            NOTE

                           The forward and reverse LSPs can be established over the same or different paths.
                           Using the same path is recommended to ensure consistent delay.


                 Figure 4-6 Setup of static bidirectional associated LSPs




4.6.2 Configuring Static Bidirectional Associated LSPs
Prerequisites
                 Before configuring static bidirectional associated LSPs, complete the following
                 task:
                 Configure forward and reverse static MPLS TE tunnels. For details, see 4.5
                 Configuring Static MPLS TE Tunnels.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the tunnel interface view.
                 interface tunnel interface-number

         Step 3 Configure a reverse CR-LSP for the tunnel.
                 mpls te reverse-lsp protocol static lsp-name lsp-name

                 ----End

Verifying the Configuration
                 ●      Run the display mpls te reverse-lsp verbose command to check detailed
                        information about the reverse LSP.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                      223
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


4.6.3 Example for Configuring Static Bidirectional Associated
LSPs

