---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-155
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [22383, 22575]
sha256: fdb1b453799520e85d93a3e812fcd5f4578114aa1db091b601f6d15741f38fa7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         traffic-eng level-1-2
                        #
                        interface vlanif 100
                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface vlanif 200
                         ip address 10.2.1.1 255.255.255.252
                         mpls
                         mpls te
                         mpls te auto-frr link
                         mpls te srlg 1
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface vlanif 500
                         ip address 10.5.1.1 255.255.255.252
                         mpls
                         mpls te
                         mpls te srlg 1
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface vlanif 300
                         ip address 10.3.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      P2
                        #
                        sysname P2
                        #
                        vlan batch 300 400
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0002.00


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      373
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         traffic-eng level-1-2
                        #
                        interface vlanif 300
                         ip address 10.3.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface vlanif 400
                         ip address 10.4.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      PE2
                        #
                        sysname PE2
                        #
                        vlan batch 200 400 500
                        #
                        mpls lsr-id 5.5.5.5
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0006.00
                         traffic-eng level-1-2
                        #
                        interface vlanif 200
                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface vlanif 500
                         ip address 10.5.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface vlanif 400
                         ip address 10.4.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      374
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration

                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 5.5.5.5 255.255.255.255
                         isis enable 1
                        #
                        return



4.20 Configuring MPLS TE Tunnel Re-optimization

4.20.1 Understanding MPLS TE Tunnel Re-optimization
                 MPLS TE tunnel re-optimization enables a TE tunnel to be automatically
                 reestablished over new optimal paths when the MPLS network topology changes.

Context
                 A main function of MPLS TE tunnels is to optimize traffic distribution over a
                 network. Generally, an MPLS TE tunnel is configured using the initial bandwidth
                 required for services and initial network topology. However, a network topology
                 and link attributes change very often, and the initial path of a tunnel may not be
                 the optimal path. As a result, network bandwidth resources may be wasted or
                 traffic distribution needs to be further optimized. This deviates from the primary
                 goal of MPLS TE to a certain extent. As such, MPLS TE tunnel re-optimization is
                 required.

Implementation
                 A specific event that occurs on the ingress node can trigger optimization of a CR-
                 LSP. The optimization enables the CR-LSP to be reestablished over the optimal
                 path with the smallest metric value.

                         NOTE

