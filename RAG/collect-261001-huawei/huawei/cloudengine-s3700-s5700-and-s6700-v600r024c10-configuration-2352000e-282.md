---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-282
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41491, 41586]
sha256: 797dbbaef4780768b1ab2a6e442cb93706cc2b137a9da745d000f4a9ecbcc770
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        BGP VPLS uses BGP signaling. Configuring a route reflector can solve the problem of
                        excessive connections caused by a full-mesh VPLS network. Therefore, it is meaningless to
                        use the HVPLS networking scheme for BGP. Only LDP HVPLS can be configured.
                        If UPEs and SPEs or SPEs are indirectly connected and LDP LSPs are used on the public
                        network, remote LDP sessions need to be established.


6.9.1 Understanding HVPLS
Definition
                    To prevent loops, BGP VPLS or LDP VPLS requires a full mesh of LSPs between all
                    PEs that provide VPLS services. As required by split horizon, packets received

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     665
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    through a PW are not forwarded to other PWs during data forwarding. If a VPLS
                    network has N PEs, the number of PWs to be established between the PEs must
                    be N x (N – 1)/2. For example, if there are 100 PEs, 4950 PWs need to be
                    established between the PEs. When a PE is added to a VPLS network, certain PEs
                    need to re-designated and PWs increase exponentially. As a result, many system
                    resources are consumed and system performance is severely diminished, limiting
                    VPLS networking deployment and application.

                    The PWs, LDP sessions, and LSPs are established using a signaling protocol. The
                    full-mesh VPLS solution cannot be applied on a large scale because bandwidth is
                    wasted when the PEs providing VCs copy data packets. To be specific, after a PE
                    receives the first unknown unicast, broadcast, or multicast packet, the PE
                    broadcasts the packet to all its peers, wasting the bandwidth.

                    Against this backdrop, HVPLS is introduced to overcome the disadvantages of a
                    full-mesh VPLS network. The core of HVPLS is network hierarchy. The network of
                    each level is fully meshed. Devices of different levels are connected through PWs
                    and forward data to each other without following the split horizon rule. This
                    reduces the burden of signaling protocols as well as that brought by data packet
                    replication, so that VPLS can be applied on a large scale.

Purpose
                    HVPLS is introduced to overcome the disadvantages of a full-mesh VPLS network
                    and enhance network scalability.

HVPLS Model
                    Figure 6-20 shows the basic HVPLS model.

                    Figure 6-20 HVPLS model




                    In the basic HVPLS model, PEs can be classified into the following types:

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          666
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    ●   UPE: a user aggregation device that is directly connected to a CE. A UPE needs
                        to be connected to only one PE on a full-mesh VPLS network. A UPE supports
                        routing and MPLS encapsulation. If a UPE connects to multiple CEs and can
                        provide the basic bridging function, frame forwarding needs to be performed
                        only on the UPE. This reduces the burden on the SPE.
                    ●   SPE: a device that is connected to a UPE and located in the core of a full-
                        mesh VPLS network. An SPE is connected to all other devices on a full-mesh
                        VPLS network.

                    For an SPE, a UPE functions like a CE. In data forwarding, an SPE uses the PW
                    established with a UPE as an AC. The UPE adds two MPLS labels to packets sent
                    by CEs. The outer label is an LSP label that is switched when a packet passes
                    through devices on the access network. The inner label is a VC label that identifies
                    a VC, and remains unchanged when a packet is transmitted along an LSP. When
                    receiving double-tagged packets, the SPE directly removes the outer label if it is a
                    public network label. The SPE determines which VSI the AC accesses based on the
                    inner label.


HVPLS Access Mode
                    The device supports only LDP HVPLS. A UPE is connected to an SPE through an
                    LSP.

                    Figure 6-21 HVPLS access through an LSP




                    In Figure 6-21, UPE1 functions as an aggregation device. It establishes a VC only
                    with SPE1 to access a PW and does not establish any VCs with other peers. The

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           667
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


