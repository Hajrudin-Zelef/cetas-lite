---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-294
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [43438, 43616]
sha256: af1f10799d317309db818e69f8484c43ec8a13da1f87083b8dedd5fb41d5df61
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        isis 1
                         network-entity 10.0000.0000.0003.00
                        #
                        interface Vlanif30
                         l2 binding vsi v1
                        #
                        interface Vlanif40
                         ip address 100.3.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
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
                         ip address 3.3.3.3 255.255.255.255
                         isis enable 1
                        #
                        bgp 200
                         peer 4.4.4.4 as-number 200
                         peer 4.4.4.4 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 4.4.4.4 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 4.4.4.4 enable
                          peer 4.4.4.4 signaling vpls
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 40 50
                        #
                        mpls lsr-id 4.4.4.4
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi v1 auto
                         pwsignal bgp
                          route-distinguisher 200:2
                          vpn-target 1:1 import-extcommunity
                          vpn-target 1:1 export-extcommunity
                          site 2 range 5 default-offset 0
                        #
                        mpls ldp
                        #
                        isis 1
                         network-entity 10.0000.0000.0004.00
                        #
                        interface Vlanif40
                         ip address 100.3.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif50


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   697
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration

                         l2 binding vsi v1
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                         isis enable 1
                        #
                        bgp 200
                         peer 3.3.3.3 as-number 200
                         peer 3.3.3.3 connect-interface LoopBack1
                         #
                         ipv4-family unicast
                          undo synchronization
                          peer 3.3.3.3 enable
                         #
                         l2vpn-ad-family
                          policy vpn-target
                          peer 3.3.3.3 enable
                          peer 3.3.3.3 signaling vpls
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 50
                        #
                        interface Vlanif50
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        return



6.11 Configuring VPLS PW Redundancy
Prerequisites
                    Before configuring VPLS PW redundancy, you have completed the following tasks:

                    ●   Configure IP addresses and an IGP on PEs.
                    ●   Enable MPLS L2VPN on PEs.
                    ●   Establish tunnels between PEs to carry L2VPN services.

6.11.1 Understanding VPLS PW Redundancy

Context
                    During network deployment, a pair of devices is typically deployed to provide
                    redundancy protection for the same group of services to ensure reliability. In this
                    case, two PWs need to be deployed to provide service protection for VPWS or

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          698
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    VPLS services that access the two devices. This mechanism is also known as PW
                    redundancy. PW redundancy is a widely used technique and has become a
                    reliability standard. It improves service switching efficiency in case of device
                    failures and further minimizes the impact of such failures on services.
                    PW redundancy is best suited for point-to-point (P2P) services, such as VPWS.
                    VPLS, a point-to-multipoint (P2MP) service, can be viewed as a collection of P2P
                    services. Therefore, PW redundancy can also be used for VPLS.
                    PW redundancy in VPLS networks, also known as VPLS PW redundancy,
                    implements fast VPLS network convergence and shortens the service interruption
                    time.

Related Concepts
                    The following describes several key concepts of VPLS PW redundancy based on
                    service traffic protection between CE1 and CE2 on the VPLS network shown in
                    Figure 6-26.

                    Figure 6-26 VPLS PW redundancy networking




                    Currently, VPLS PW redundancy operates in either of the following modes (the
                    operating mode is specified on PE1):
                    ●   Master/Slave mode: PE1 determines whether a local PW is in the primary or
                        secondary state based on the configured PW forwarding priority.
                    ●   Independent mode: PE1 determines whether a local PW is in the primary or
                        secondary state based on the forwarding status notified by the remote PEs
                        (PE2 and PE5).
                    The PEs on both ends of the primary and secondary PWs in a PW protection group
                    must negotiate the PW status to ensure that they select the same PW to transmit
                    user packets. The following concepts are introduced for PW status negotiation
                    inside a PW protection group:
                    ●   Primary/Secondary: used to describe the forwarding priority of a PW and can
                        be configured. A smaller value indicates a higher priority, and a PW with the
                        highest priority is the primary PW.

