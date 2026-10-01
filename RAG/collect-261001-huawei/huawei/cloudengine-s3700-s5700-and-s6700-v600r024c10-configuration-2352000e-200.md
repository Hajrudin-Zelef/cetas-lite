---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-200
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [29290, 29450]
sha256: 89db16818dee92568b611586e7d072f573aef9031418925e7afe616ece7444db
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure an IGP on the backbone network to enable communication
                         between devices on the backbone network.
                    2.   Configure basic MPLS functions on the backbone network, establish an MPLS
                         TE tunnel, and configure a tunnel policy.
                    3.   Enable L2VPN and create a VPWS connection on PEs.
                               NOTE

                              By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device.
                              If a VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts
                              with LNP. In this case, run the lnp disable command in the system view to disable
                              LNP.


Procedure
         Step 1 Assign an IP address to each interface and configure an IGP to implement
                interworking.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 20
                    [CE1] interface 10ge 1/0/2
                    [CE1-10GE1/0/2] port link-type trunk
                    [CE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE1-10GE1/0/2] quit
                    [CE1] interface vlanif 20
                    [CE1-Vlanif20] ip address 10.10.1.1 255.255.255.0
                    [CE1-Vlanif20] quit

                    The configuration of CE2 is similar to that of CE1. For detailed configurations, see
                    Configuration Scripts.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     469
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] vlan batch 10 20
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [PE1-10GE1/0/1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [PE1-10GE1/0/2] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] ip address 10.1.1.1 255.255.255.0
                    [PE1-Vlanif10] quit
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
         Step 2 Enable MPLS, MPLS TE, MPLS RSVP-TE, and MPLS TE Constrained Shortest Path
                First (CSPF).
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] mpls te
                    [PE1-mpls] mpls rsvp-te
                    [PE1-mpls] mpls te cspf
                    [PE1-mpls] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] mpls
                    [PE1-Vlanif10] mpls te
                    [PE1-Vlanif10] mpls rsvp-te
                    [PE1-Vlanif10] quit

                    # Configure the P.
                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] mpls te
                    [P-mpls] mpls rsvp-te
                    [P-mpls] quit
                    [P] interface vlanif 10
                    [P-Vlanif10] mpls
                    [P-Vlanif10] mpls te
                    [P-Vlanif10] mpls rsvp-te
                    [P-Vlanif10] quit
                    [P] interface vlanif 20
                    [P-Vlanif20] mpls
                    [P-Vlanif20] mpls te
                    [P-Vlanif20] mpls rsvp-te
                    [P-Vlanif20] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 3.3.3.9
                    [PE2] mpls
                    [PE2-mpls] mpls te
                    [PE2-mpls] mpls rsvp-te


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                          470
VPN Configuration
VPN Configuration                                                                     5 VPWS Configuration

                    [PE2-mpls] mpls te cspf
                    [PE2-mpls] quit
                    [PE2] interface vlanif 20
                    [PE2-Vlanif20] mpls
                    [PE2-Vlanif20] mpls te
                    [PE2-Vlanif20] mpls rsvp-te
                    [PE2-Vlanif20] quit

         Step 3 Configure OSPF TE on the backbone network to advertise TE information.
                    # Configure PE1.
                    [PE1] ospf 1
                    [PE1-ospf-1] opaque-capability enable
                    [PE1-ospf-1] area 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] mpls-te enable
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
         Step 4 Configure MPLS TE attributes for links and set the maximum bandwidth and
                maximum reservable bandwidth for links on each interface along the tunnel.
                    # Configure PE1.
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] mpls te bandwidth max-reservable-bandwidth 10000
                    [PE1-Vlanif10] mpls te bandwidth bc0 5000
                    [PE1-Vlanif10] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
         Step 5 Configure an MPLS TE tunnel interface, specify the tunnel protocol as MPLS TE
                and the signaling protocol as RSVP-TE, and specify the bandwidth.
                    # Configure PE1.
                    [PE1] interface tunnel 10
                    [PE1-Tunnel10] ip address unnumbered interface loopback 1
                    [PE1-Tunnel10] tunnel-protocol mpls te
                    [PE1-Tunnel10] destination 3.3.3.9
                    [PE1-Tunnel10] mpls te tunnel-id 10
                    [PE1-Tunnel10] mpls te signal-protocol rsvp-te
                    [PE1-Tunnel10] mpls te bandwidth ct0 2000
                    [PE1-Tunnel10] quit

                    # Configure PE2.
                    [PE2] interface tunnel 10
                    [PE2-Tunnel10] ip address unnumbered interface loopback 1
                    [PE2-Tunnel10] tunnel-protocol mpls te
                    [PE2-Tunnel10] destination 1.1.1.9
                    [PE2-Tunnel10] mpls te tunnel-id 10
                    [PE2-Tunnel10] mpls te signal-protocol rsvp-te
                    [PE2-Tunnel10] mpls te bandwidth ct0 2000
                    [PE2-Tunnel10] quit

