---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-199
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28824, 28989]
sha256: acf8ac4b0fd5eadd9b8f176d68fab8e936871a10149c111ebf8562d5de50163a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      478
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 requires link protection. A bypass tunnel is established on a directly connected link
                 from LSR1's inbound interface to LSR3's outbound interface. This tunnel bypasses
                 the transit node LSR2 and provides node protection. Another bypass tunnel is
                 established from LSR1's outbound interface to LSR2's inbound interface. This
                 tunnel uses LSR4 as a transit node, bypasses the direct link from LSR1's outbound
                 interface to LSR2's inbound interface, and provides link protection.

                 Figure 4-51 Network diagram of MPLS TE Auto FRR (TE FRR link protection on
                 the ingress)




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure a primary tunnel, enable TE FRR in the tunnel interface view, and
                        enable MPLS Auto FRR in the MPLS view.
                 2.     Configure the tunnel ingress to set the Node protection flag to 0 in sent
                        messages, indicating that node protection is not desired by the tunnel.
                 3.     Configure the PLR to consider the Node protection flag sent by the tunnel
                        ingress during FRR protection.

Procedure
         Step 1 Configure interface IP addresses for the nodes.
                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100 200 400
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface vlanif 200
                 [LSR1-Vlanif200] ip address 10.21.1.1 24
                 [LSR1-Vlanif200] quit
                 [LSR1] interface vlanif 400
                 [LSR1-Vlanif400] ip address 10.41.1.2 24
                 [LSR1-Vlanif400] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           479
MPLS Configuration
MPLS Configuration                                                          4 MPLS TE Configuration

                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface 10ge 1/0/2
                 [LSR1-10GE1/0/2] port link-type trunk
                 [LSR1-10GE1/0/2] port trunk allow-pass vlan 200
                 [LSR1-10GE1/0/2] quit
                 [LSR1] interface 10ge 1/0/3
                 [LSR1-10GE1/0/3] port link-type trunk
                 [LSR1-10GE1/0/3] port trunk allow-pass vlan 400
                 [LSR1-10GE1/0/3] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.1 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

         Step 2 Configure OSPF to advertise routes.

                 # Configure LSR1.
                 [LSR1] ospf 1
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                 [LSR1-ospf-1-area-0.0.0.0] network 10.21.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] network 10.41.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

                 After the configuration is complete, run the display ip routing-table command on
                 each node to check whether the nodes have learned routes from each other.

         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.

                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along tunnels. Enable CSPF on the ingress of the primary tunnel and the PLR of
                 the bypass tunnel.

                 # Configure LSR1.
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
                 [LSR1] interface vlanif 200
                 [LSR1-Vlanif200] mpls
                 [LSR1-Vlanif200] mpls te
                 [LSR1-Vlanif200] mpls rsvp-te
                 [LSR1-Vlanif200] quit
                 [LSR1] interface vlanif 400
                 [LSR1-Vlanif400] mpls
                 [LSR1-Vlanif400] mpls te
                 [LSR1-Vlanif400] mpls rsvp-te
                 [LSR1-Vlanif400] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         480
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


         Step 4 Configure OSPF TE.
                 # Configure LSR1.
                 [LSR1] ospf
                 [LSR1-ospf-1] opaque-capability enable
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] mpls-te enable
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.
         Step 5 Configure an explicit path for the primary tunnel.
                 [LSR1] explicit-path master
                 [LSR1-explicit-path-pri-master] next hop 10.21.1.2
                 [LSR1-explicit-path-pri-master] next hop 10.31.1.2
                 [LSR1-explicit-path-pri-master] quit

         Step 6 Enable TE Auto FRR.
                 [LSR1] mpls
                 [LSR1-mpls] mpls te auto-frr

         Step 7 Configure the device to set the Node protection flag to 0 in sent messages,
                indicating that node protection is not desired by the tunnel.
                 [LSR1] mpls
                 [LSR1-mpls] mpls te frr no-node-protection

         Step 8 Configure the device to consider the Node protection flag sent by the tunnel
                ingress during FRR protection.
                 [LSR1] mpls
                 [LSR1-mpls] mpls te plr apply node-protection-flag

         Step 9 Configure the primary MPLS TE tunnel and bind it to the explicit path.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 3.3.3.3
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te record-route label
                 [LSR1-Tunnel1] mpls te path explicit-path master
                 [LSR1-Tunnel1] mpls te fast-reroute

                 ----End

