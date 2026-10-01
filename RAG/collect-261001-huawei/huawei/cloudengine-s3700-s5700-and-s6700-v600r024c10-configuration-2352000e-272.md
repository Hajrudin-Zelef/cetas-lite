---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-272
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["agi", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40002, 40170]
sha256: 2b63a8084e2e458efbb45d664f6d3dd5a53c5e9766f6e4f499c911a677a74776
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                          network 192.168.2.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 30 40
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi bgp1 auto
                         pwsignal bgp
                          route-distinguisher 9.9.9.2:1
                          vpn-target 100:1 import-extcommunity
                          vpn-target 100:1 export-extcommunity
                          site 2 range 5 default-offset 0
                        #
                        mpls ldp
                        #
                        interface Vlanif30
                         ip address 192.168.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         l2 binding vsi bgp1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        bgp 100
                         peer 1.1.1.9 as-number 100
                         peer 1.1.1.9 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 1.1.1.9 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 1.1.1.9 enable
                          peer 1.1.1.9 signaling vpls
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   641
VPN Configuration
VPN Configuration                                                                      6 VPLS Configuration

                          network 192.168.2.0 0.0.0.255
                        #
                        return



6.8 Configuring BGP AD VPLS
Prerequisites
                    Before configuring BGP AD VPLS, you have completed the following tasks:

                    ●   Configure IP addresses and an IGP on PEs and Ps.
                    ●   Configure LSR IDs and enable basic MPLS functions on PEs and Ps.
                    ●   Enable MPLS L2VPN on PEs.
                    ●   Establish tunnels between PEs to carry L2VPN services.

6.8.1 Understanding BGP AD VPLS

Definition
                    BGP AD VPLS, short for Border Gateway Protocol Auto-Discovery virtual private
                    LAN service, is a new technology for automatically deploying VPLS services.

                    BGP AD VPLS uses extended BGP Update messages to automatically discover
                    members in a VPLS domain, and uses LDP FEC 129 for local and remote VSIs to
                    automatically negotiate and establish VPLS PWs. In addition, BGP AD also
                    supports hierarchical virtual private LAN service (HVPLS). You can disable split
                    horizon to allow a peer to be used as a user-side device on an HVPLS network.


Purpose
                    As VPLS technologies are used more widely and VPLS networks grow in scale,
                    VPLS configurations on networks increase accordingly. BGP AD VPLS is introduced
                    to simplify network configurations, enable automatic service deployment, and
                    reduce OPEX.

                    BGP AD VPLS combines the advantages of both BGP VPLS and LDP VPLS. BGP AD
                    VPLS-enabled devices use extended BGP Update messages to automatically
                    discover members in a VPLS domain, and use LDP FEC 129 to negotiate and
                    establish PWs and automatically deploy VPLS services.

                    With automatic VPLS member discovery and automatic PW deployment, BGP AD
                    VPLS reduces the VPLS network configuration workload, implements automatic
                    service deployment, and reduces customers' OPEX.


Related Concepts
                     Acronyms or           Full Spelling                    Function
                     Abbreviations

                     VPLS ID               virtual private LAN service ID   Identifier of a VPLS domain


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                               642
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                     Acronyms or         Full Spelling                   Function
                     Abbreviations

                     VSI ID              virtual switching instance ID   Identifier of a VSI in a VPLS
                                                                         domain

                     RD                  route distinguisher             Route distinguisher in a BGP
                                                                         message which carries VSI
                                                                         information

                     RT                  route target                    Route attribute carried in a
                                                                         BGP message used to
                                                                         advertise VSI information

                     AGI                 attachment group identifier     Domain identifier used for
                                                                         negotiation between VSIs in
                                                                         the same VPLS domain

                     AII                 attachment individual           VSI identifier used for
                                         identifier                      negotiation between VSIs in
                                                                         the same VPLS domain

                     SAII                source attachment individual    IP address used by a local
                                         identifier                      device to negotiate a PW
                                                                         with its peer in a BGP-AD VSI.

                     TAII                target attachment individual    IP address used by a remote
                                         identifier                      device to negotiate a PW
                                                                         with its peer in a BGP-AD VSI.

                     FEC 129             forwarding equivalence class    New type of FEC used by LDP
                                         129                             signaling




