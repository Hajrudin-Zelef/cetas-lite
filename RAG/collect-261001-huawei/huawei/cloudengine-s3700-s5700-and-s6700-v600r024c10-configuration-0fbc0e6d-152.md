---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-152
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [21855, 22042]
sha256: e7a5b2cdac7a8f22a40866e1dd742c1ef46b47674ff105fdf366d33320a8795e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif700
                         ip address 10.7.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 700
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      P4
                        #
                        sysname P4
                        #
                        vlan batch 300 500 600
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0004.00
                         traffic-eng level-1-2
                        #
                        interface Vlanif300
                         ip address 10.3.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.5.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif600
                         ip address 10.6.1.1 255.255.255.252
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
                         port trunk allow-pass vlan 500
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      365
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                         isis enable 1
                        #
                        return

                 ●      PE2
                        #
                        sysname PE2
                        #
                        vlan batch 600 700 800
                        #
                        mpls lsr-id 6.6.6.6
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
                        interface Vlanif600
                         ip address 10.6.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif700
                         ip address 10.7.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif800
                         ip address 10.8.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 700
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 800
                        #
                        interface LoopBack1
                         ip address 6.6.6.6 255.255.255.255
                         isis enable 1
                        #
                        return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      366
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


4.19.4 Example for Configuring an SRLG in a TE Auto FRR
Scenario
Networking Requirements
                 On the network shown in Figure 4-32, an RSVP-TE tunnel is established from PE1
                 to PE2 over the path PE1 -> P1 -> PE2. The outbound interface of the primary
                 tunnel on P1 is VLANIF500. The links on network segments of 10.2.1.0/30 and
                 10.5.1.0/30 belong to SRLG1.
                 Configure TE Auto FRR on P1 to improve reliability. A link in an SRLG different
                 from those used by the primary tunnel is preferentially selected for the automatic
                 bypass tunnel.

                 Figure 4-32 Network diagram of an SRLG in a TE Auto FRR scenario




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure IP addresses for interfaces, including the loopback interfaces whose
                        addresses are to be used as MPLS LSR IDs.
                 2.     Configure IS-IS to ensure that nodes can reach each other over public network
                        routes.
                 3.     Enable MPLS, MPLS TE, and RSVP-TE on all nodes and interfaces involved.
                 4.     Configure IS-IS TE on all nodes and enable CSPF on PE1 and P1.
                 5.     Configure a dynamic primary MPLS TE tunnel from PE1 to PE2 over the
                        explicit path PE1 -> P1 -> PE2.
                 6.     Set the SRLG number on the member interfaces of the SRLG.
                 7.     Configure an SRLG-based path calculation mode on the PLR.
                 8.     Enable TE FRR in the tunnel interface view on the ingress. Enable TE Auto FRR
                        on the outbound interface of the primary tunnel on the PLR.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           367
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


