---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-265
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [38864, 39034]
sha256: 1c86b3bb05b387bc509b8cfbf8b379fb23d5d4fa757a68fa8761295fb90224ba
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif20
                         ip address 192.168.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         ip address 192.168.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
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
                        vsi a2 static
                         pwsignal ldp
                          vsi-id 2
                          peer 1.1.1.9
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.9
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif30
                         ip address 192.168.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         l2 binding vsi a2
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   622
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration

                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 192.168.2.0 0.0.0.255
                        #
                        return



6.7 Configuring BGP VPLS
Prerequisites
                    Before configuring BGP VPLS, you have completed the following tasks:
                    ●   Configure IP addresses and an IGP on PEs and Ps.
                    ●   Configure LSR IDs and enable basic MPLS functions on PEs and Ps.
                    ●   Establish tunnels between PEs to carry L2VPN services.

6.7.1 Understanding BGP VPLS
Context
                    BGP VPLS uses a dynamic discovery mechanism to discover VPLS members and
                    uses BGP as the signaling protocol. BGP VPLS uses Multiprotocol Extensions for
                    BGP (MP-BGP) messages to transmit VPLS member information. The MP-REACH
                    and MP-UNREACH attributes carry VPLS label information; the extended
                    community attributes carry interface parameters, route distinguishers (RDs), and
                    VPN targets; the RDs and VPN targets identify VPN member relationships.

Related Concepts
                    BGP VPLS involves the following concepts:
                    ●   MP-BGP: supports multiple network layer protocols and identifies the
                        protocols based on address families. MP-BGP transmits VPN composition
                        information and VPN-IPv4 routes between PEs. MP-BGP introduces two path
                        attributes: MP_REACH and MP_UNREACH.
                    ●   MP_REACH: advertises reachable routes and their next hops.
                    ●   MP_UNREACH: instructs a peer to delete unreachable routes.

Implementation
                    Figure 6-11 shows the process for establishing a PW using BGP signaling.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       623
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    Figure 6-11 Process for establishing a PW using BGP signaling




                    1.   After PE1 has a local label block configured and a BGP session is established
                         between PE1 and PE2, PE1 sends an Update message carrying the MP-REACH
                         attribute, site ID, and label block information to PE2.
                    2.   After receiving the Update message, PE2 calculates a unique label as the VC
                         label based on its own site ID and label block carried in the message, and
                         then establishes a unidirectional VC VC1. In addition, PE2 calculates the VC
                         label of PE1 based on the site ID and label block carried in the Update
                         message sent by PE1. Based on this information, PE2 sends an Update
                         message to PE1. After receiving the Update message, PE1 takes similar
                         actions as PE2 and then establishes a unidirectional VC VC2.

                    Figure 6-12 shows the process for tearing down a PW using BGP signaling.

                    Figure 6-12 Process for tearing down a PW using BGP signaling




                    1.   After the local label block is deleted from PE1, PE1 withdraws its local VC
                         label and tears down VC2, and sends an Update message carrying the MP-
                         UNREACH attribute to PE2.
                    2.   After receiving the message, PE2 withdraws its local VC label and tears down
                         VC1.

Application Scenario
                    BGP VPLS applies to the core layer of large networks where PEs run BGP and
                    inter-AS communication is required.
                    If PEs support BGP as the VPLS signaling protocol, you can configure BGP VPLS. In
                    Figure 6-13, PE1, PE2, and PE3 are in the same VPLS network.
                    ●    To enable CEs connected to PE1, PE2, and PE3 to communicate with each
                         other by constructing a full-mesh VPLS network, ensure that PE1, PE2, and
                         PE3 have matching VPN targets.
                    ●    To enable CE1 to communicate with both CE2 and CE3 and prevent CE2 and
                         CE3 from communicating with each other, ensure that the import VPN target

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            624
VPN Configuration
VPN Configuration                                                                6 VPLS Configuration


