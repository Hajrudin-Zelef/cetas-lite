---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-206
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [29845, 29997]
sha256: 2fda70ff243a70018c3a249fa5ba145183d3b2c3d5fefa2428c8c50b359f9a94
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                 The primary path of VPN FRR is PE1 -> Switch -> PE2, and the backup path is PE1
                 -> PE3. In a normal situation, VPN traffic is transmitted over the primary path. If
                 the primary path fails, VPN traffic is switched to the backup path. Configure static
                 BFD for TE tunnel to monitor the tunnel of the primary path and enable VPN to
                 rapidly detect tunnel faults. Traffic is rapidly switched between the primary and
                 backup paths, and fault recovery is sped up.

                 Figure 4-53 Network diagram of BFD for TE tunnel
                         NOTE

                        For simplicity, the IP addresses of the interfaces connected the PEs and the CEs are not
                        shown in the diagram.




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure IP addresses for interfaces, including the loopback interfaces whose
                        addresses are to be used as MPLS LSR IDs.
                 2.     Configure OSPF to ensure that nodes can reach each other over public
                        network routes.
                 3.     Configure an MPLS network and establish bidirectional TE tunnels between
                        PE1 and PE2, and between PE1 and PE3.
                 4.     Configure VPN FRR on PE1.
                 5.     Enable BFD on PE1, PE2, and PE3 globally.
                 6.     Establish a BFD session on PE1 to monitor the TE tunnel of the primary path.
                 7.     Configure BFD sessions on PE2 and PE3 and specify the TE tunnel as the BFD
                        reverse tunnel.

Procedure
         Step 1 Configure interface IP addresses for the devices.
                 # Configure PE1.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                         495
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration

                 <HUAWEI> system-view
                 [HUAWEI] sysname PE1
                 [PE1] vlan batch 100 200
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] ip address 10.1.1.1 30
                 [PE1-Vlanif100] quit
                 [PE1] interface vlanif 200
                 [PE1-Vlanif200] ip address 10.2.1.1 30
                 [PE1-Vlanif200] quit
                 [PE1] interface 10ge 1/0/1
                 [PE1-10GE1/0/1] port link-type trunk
                 [PE1-10GE1/0/1] port trunk allow-pass vlan 100
                 [PE1-10GE1/0/1] quit
                 [PE1] interface 10ge 1/0/2
                 [PE1-10GE1/0/2] port link-type trunk
                 [PE1-10GE1/0/2] port trunk allow-pass vlan 200
                 [PE1-10GE1/0/2] quit
                 [PE1] interface loopback 1
                 [PE1-loopback1] ip address 1.1.1.1 32
                 [PE1-loopback1] quit

                 The configurations of PE2 and PE3 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.
         Step 2 Configure OSPF to advertise routes.
                 # Configure PE1.
                 [PE1] ospf 1
                 [PE1-ospf-1] area 0
                 [PE1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                 [PE1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.3
                 [PE1-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.255
                 [PE1-ospf-1-area-0.0.0.0] quit
                 [PE1-ospf-1] quit

                 The configurations of PE2 and PE3 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.
                 After the configuration is complete, run the display ip routing-table command on
                 each node to check whether the nodes have learned routes from each other.
         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the ingress.
                 # Configure PE1.
                 [PE1] mpls lsr-id 1.1.1.1
                 [PE1] mpls
                 [PE1-mpls] mpls te
                 [PE1-mpls] mpls rsvp-te
                 [PE1-mpls] mpls te cspf
                 [PE1-mpls] quit
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] mpls
                 [PE1-Vlanif100] mpls te
                 [PE1-Vlanif100] mpls rsvp-te
                 [PE1-Vlanif100] quit
                 [PE1] interface vlanif 200
                 [PE1-Vlanif200] mpls
                 [PE1-Vlanif200] mpls te
                 [PE1-Vlanif200] mpls rsvp-te
                 [PE1-Vlanif200] quit

                 The configurations of PE2 and PE3 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         496
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration


         Step 4 Configure OSPF TE.

                 # Configure PE1.
                 [PE1] ospf
                 [PE1-ospf-1] opaque-capability enable
                 [PE1-ospf-1] area 0
                 [PE1-ospf-1-area-0.0.0.0] mpls-te enable
                 [PE1-ospf-1-area-0.0.0.0] quit
                 [PE1-ospf-1] quit

                 The configurations of PE2 and PE3 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.

         Step 5 Configure an explicit path for a tunnel.

                 # Configure explicit paths on PE1, PE2, and PE3. Two explicit paths need to be
                 created on PE1.
                 [PE1] explicit-path tope2
                 [PE1-explicit-path-tope2] next hop 10.2.1.2
                 [PE1-explicit-path-tope2] next hop 3.3.3.3
                 [PE1-explicit-path-tope2] quit
                 [PE1] explicit-path tope3
                 [PE1-explicit-path-tope3] next hop 10.1.1.2
                 [PE1-explicit-path-tope3] next hop 2.2.2.2
                 [PE1-explicit-path-tope3] quit

                 # Configure an explicit path from PE2 to PE1.
                 [PE2] explicit-path tope1
                 [PE2-explicit-path-tope1] next hop 10.2.1.1
                 [PE2-explicit-path-tope1] next hop 1.1.1.1
                 [PE2-explicit-path-tope1] quit

                 # Configure an explicit path from PE3 to PE1.
                 [PE3] explicit-path tope1
                 [PE3-explicit-path-tope1] next hop 10.1.1.1
                 [PE3-explicit-path-tope1] next hop 1.1.1.1
                 [PE3-explicit-path-tope1] quit

         Step 6 Configure MPLS TE tunnel interfaces.

                 Create tunnel interfaces and specify explicit paths on PE1, PE2, and PE3. Reserve
                 the tunnel for VPN binding Configure two tunnel interfaces on PE1.

