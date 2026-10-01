---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-111
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15623, 15820]
sha256: 53650b1072aec919ce26af9b82b561bb2038da3b800de478dfb98abb334b377f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         ip address 10.1.1.2 255.255.255.0
                         ospf cost 15
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         ospf cost 10
                         mpls
                         mpls te
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
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200 300
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path pri-path
                         next hop 10.1.2.1
                         next hop 10.1.1.1
                         next hop 1.1.1.9
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         ospf cost 10
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      264
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 1.1.1.9
                         mpls te tunnel-id 2
                         mpls te path explicit-path pri-path
                         mpls te igp advertise
                         mpls te igp metric absolute 10
                        #
                        ospf 1
                         opaque-capability enable
                         enable traffic-adjustment advertise
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300 400 500
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         ospf cost 10
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
                         ospf cost 10
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                        #
                        return
                 ●      LSR5
                        #
                        sysname LSR5
                        #
                        vlan batch 500 600
                        #
                        interface Vlanif500
                         ip address 10.1.5.2 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      265
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration

                         ospf cost 10
                        #
                        interface Vlanif600
                         ip address 10.1.6.2 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.5.0 0.0.0.255
                          network 10.1.6.0 0.0.0.255
                        #
                        return



4.9 Adjusting RSVP-TE Signaling Parameters

4.9.1 Configuring an RSVP-TE Resource Reservation Style

Prerequisites
                 Before configuring an RSVP-TE resource reservation style, complete the following
                 task:

                 ●      Configure a dynamic MPLS TE tunnel.


Context
                 RSVP TE supports diversified signaling parameters, which meet requirements of
                 reliability and network resources and some MPLS TE advanced features. For details
                 about RSVP-TE, see 4.2.3 RSVP-TE Message Format and 4.2.4 RSVP-TE
                 Fundamentals.

                 A reservation style defines how a node reserves resource after receiving a request
                 sent by an upstream node. If paths of multiple CR-LSPs pass through the same
                 node, you need to configure a resource reservation style on the ingress of an
                 MPLS TE tunnel to manage the processing of resource reservation requests. Local
                 reserved resources can be assigned to multiple CR-LSPs either separately or in
                 shared mode.

