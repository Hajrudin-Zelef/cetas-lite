---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-47
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [6340, 6508]
sha256: f702103a3f0e4cd7f1a31de59795e11fc3cd3c4d68e6386434a55d9cdc59a5d3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         isis enable 1
                         isis circuit-level level-1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type access
                         port default vlan 300
                        #
                        interface LoopBack0
                         ip address 10.10.2.2 255.255.255.255
                         isis enable 1
                        #
                         ip ip-prefix permit-host index 10 permit 0.0.0.0 32
                        #
                        return
                 ●      LSRB
                        #
                         sysname LSRB
                        #
                        vlan batch 100
                        #
                         mpls lsr-id 10.10.3.1
                         mpls
                        #
                        mpls ldp
                        #
                        isis 1
                         is-level level-1
                         network-entity 10.0010.0300.0001.00
                        #
                        interface Vlanif100
                         ip address 10.1.2.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack0
                         ip address 10.10.3.1 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      LSRC
                        #
                         sysname LSRC
                        #
                        vlan batch 100
                        #
                         mpls lsr-id 10.10.3.2
                         mpls
                        #
                        mpls ldp
                        #


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       107
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        isis 1
                         is-level level-1
                         network-entity 10.0010.0300.0002.00
                        #
                        interface Vlanif100
                         ip address 10.1.3.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack0
                         ip address 10.10.3.2 255.255.255.255
                         isis enable 1
                        #
                        return



3.14 Configuring LDP Auto FRR

3.14.1 Understanding LDP Auto FRR
                 LDP auto fast reroute (FRR) backs up local interfaces to provide the fast reroute
                 function on an MPLS network.

Context
                 On an MPLS network with both primary and backup links, if the primary link fails,
                 IGP routes re-converge to the backup link, and then the LDP LSP is switched to the
                 backup link. As this process leads to a small amount of traffic loss, LDP Auto FRR
                 can be introduced to minimize this.
                 On the network enabled with LDP Auto FRR, if an interface failure (detected by
                 the interface itself or by an associated BFD session) or a primary LSP failure
                 (detected by an associated BFD session) occurs, LDP Auto FRR rapidly forwards
                 traffic to a backup LSP, protecting traffic on the primary LSP and minimizing the
                 traffic interruption time.

Implementation
                 LDP LFA FRR
                 LDP LFA FRR is implemented based on IGP LFA FRR's LDP Auto FRR. LDP LFA FRR
                 uses the liberal label retention mode to obtain a liberal label, applies for a
                 forwarding entry associated with the label, and forwards the forwarding entry to
                 the forwarding plane as a backup forwarding entry to be used by the primary LSP.
                 If an interface failure (detected by the interface itself or by an associated BFD
                 session) or a primary LSP failure (detected by an associated BFD session) occurs,
                 LDP LFA FRR rapidly forwards traffic to a backup LSP, thereby minimizing the
                 traffic interruption time.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       108
MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration


                 Figure 3-20 Typical application scenario of LDP Auto FRR - triangle topology




                 If a backup route corresponding to the source of the liberal label for LDP auto FRR
                 exists and the route's destination meets the policy for LDP to create a backup LSP,
                 LSR-A can apply a forwarding entry for the liberal label, establish a backup LSP,
                 and send the entries mapped to both the primary and backup LSPs to the
                 forwarding plane. In this way, the primary LSP is associated with the backup LSP.
                 LDP Auto FRR is triggered when the interface detects faults by itself, BFD detects
                 faults in the interface, or BFD detects a primary LSP failure. After LSP FRR is
                 complete, traffic is switched to the backup LSP based on the backup forwarding
                 entry. The route is then converged from the path LSR-A -> LSR-B to the path LSR-
                 A -> LSR-C -> LSR-B. A new LSP is established on the original backup path, and
                 the original primary LSP is torn down. Traffic is then forwarded along the new LSP
                 on the path LSR-A -> LSR-C -> LSR-B.
                 LDP Remote LFA FRR
                 LDP LFA FRR cannot calculate backup paths on some large-scale networks,
                 especially on ring networks, which fails to meet reliability requirements. As shown
                 in Figure 3-21, on a common ring network, traffic from PE1 to PE2 is transmitted
                 along the shortest path PE1 -> PE2 with the smallest cost. If a fault occurs on the
                 link between PE1 and PE2, PE1 first detects the fault and forwards traffic to P1. P1
                 is expected to forward traffic to P2 and finally to PE2. However, P1 does not detect
                 the fault immediately after the fault occurs. After the traffic forwarded by PE1
                 reaches P1, P1 sends the traffic back to PE1 as the path to PE1 has the smallest
                 cost. In this case, a routing loop occurs between PE1 and P1. A large number of
                 loop packets are transmitted on the link between PE1 and P1. As a result, some
                 normal packets from PE1 to P1 may be discarded due to congestion.




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                           109

