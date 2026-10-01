---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-141
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [20224, 20379]
sha256: 130693760b33c6409fc9e330ff0aea4490d5b36441ff4ba48b0ca10e186c6830
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Verifying the Configuration
                 ●      Run the display interface tunnel [ interface-number ] command to check the
                        operating status of tunnel interfaces.
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the local node.
                 ●      Run the display mpls lsp command to check CR-LSP information.
                 ●      Run the display mpls te tunnel [ verbose ] command to check MPLS TE
                        tunnel information.
                 ●      Run the display mpls te tunnel path command to check the path attributes
                        of tunnels on the local node.

4.16.3 Example for Configuring the Link Administrative Group
and Affinity Attributes of MPLS TE

Networking Requirements
                 On the network shown in Figure 4-27, the bandwidth of the shared link LSR1 ->
                 LSR2 is 50 Mbit/s, the maximum reservable bandwidth of the other links is 100
                 Mbit/s, and the BC0 bandwidth is 100 Mbit/s.

                 LSR1 has two dynamic MPLS TE tunnels to LSR3 (Tunnel1 and Tunnel2). Both
                 tunnels require 40 Mbit/s of bandwidth. The total bandwidth of these two tunnels
                 is 80 Mbit/s, higher than the bandwidth (50 Mbit/s) of the shared link LSR1 ->
                 LSR2. In addition, Tunnel2 has a higher priority than Tunnel1, and preemption is
                 enabled.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                       339
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


                 Configure the link administrative group and affinity attributes and masks to allow
                 Tunnel1 and Tunnel2 on LSR1 to use separate physical links on the path LSR2 ->
                 LSR3.

                 Figure 4-27 Networking diagram of the link administrative group and affinity
                 attributes of MPLS TE




Procedure
         Step 1 Configure interface IP addresses for the nodes.

                 # Configure LSR1.
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
                 [LSR1-loopback1] ip address 1.1.1.1 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.

         Step 2 Configure OSPF to advertise routes.

                 # Configure LSR1.
                 [LSR1] ospf 1
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.

                 After the configuration is complete, run the display ip routing-table command on
                 each node to check whether the nodes have learned routes from each other.

         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.

                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the ingress.

                 # Configure LSR1.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                       340
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration

                 [LSR1] mpls lsr-id 1.1.1.1
                 [LSR1] mpls
                 [LSR1-mpls] mpls te
                 [LSR1-mpls] mpls rsvp-te
                 [LSR1-mpls] mpls te cspf
                 [LSR1-mpls] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls
                 [LSR1-Vlanif100] mpls te
                 [LSR1-Vlanif100] mpls rsvp-te
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 4 Configure OSPF TE.
                 # Configure LSR1.
                 [LSR1] ospf
                 [LSR1-ospf-1] opaque-capability enable
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] mpls-te enable
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 5 Configure MPLS TE attributes on the outbound interface of each node.
                 # Set the maximum reservable link bandwidth and BC0 bandwidth to 50 Mbit/s
                 on LSR1.
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls te bandwidth max-reservable-bandwidth 50000
                 [LSR1-Vlanif100] mpls te bandwidth bc0 50000

                 # Set the link administrative group attribute to 0x10001 on LSR1.
                 [LSR1-Vlanif100] mpls te link administrative group 10001
                 [LSR1-Vlanif100] quit

                 # Configure MPLS TE attributes on LSR2.
                 [LSR2] interface vlanif 200
                 [LSR2-Vlanif200] mpls te bandwidth max-reservable-bandwidth 100000
                 [LSR2-Vlanif200] mpls te bandwidth bc0 100000
                 [LSR2-Vlanif200] mpls te link administrative group 10101
                 [LSR2-Vlanif200] quit
                 [LSR2] interface vlanif 300
                 [LSR2-Vlanif300] mpls te bandwidth max-reservable-bandwidth 100000
                 [LSR2-Vlanif300] mpls te bandwidth bc0 100000
                 [LSR2-Vlanif300] mpls te link administrative group 10011
                 [LSR2-Vlanif300] quit

                 # After the configuration is complete, check the TEDB on LSR1. The TEDB contains
                 the maximum available bandwidth and maximum reservable bandwidth of each
                 link, and the Color field, which indicates the administrative group attribute of
                 each link.
                 [LSR1] display mpls te cspf tedb node
                  Router ID: 1.1.1.1
                  IGP Type: OSPF     Process ID: 1    IGP Area: 0
                   MPLS-TE Link Count: 1
                   Link[1]:
                    OSPF Router ID: 10.1.1.1      Opaque LSA ID: 1.0.0.1
                    Interface IP Address: 10.1.1.1


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               341
MPLS Configuration
MPLS Configuration                                                                 4 MPLS TE Configuration

