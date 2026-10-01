---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-251
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution", "ethernet", "parameters", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [36887, 37037]
sha256: c8cf0f97fc76244188ed0ed9f1a0aeab77aa7a995dda55190dfd4fa4e6e29d24
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    6.1 Overview of VPLS
                    6.2 Understanding VPLS
                    6.3 Configuration Precautions for VPLS
                    6.4 Default Settings for VPLS
                    6.5 Configuring Static VPLS
                    6.6 Configuring LDP VPLS
                    6.7 Configuring BGP VPLS
                    6.8 Configuring BGP AD VPLS
                    6.9 Configuring LDP HVPLS
                    6.10 Configuring Inter-AS VPLS
                    6.11 Configuring VPLS PW Redundancy
                    6.12 Configuring VPLS Interworking
                    6.13 Configuring VPWS Accessing VPLS
                    6.14 Configuring MAC Address Learning
                    6.15 Configuring VPLS Service Isolation
                    6.16 Configuring Static BFD for VPLS PW
                    6.17 Configuring Common VPLS Parameters
                    6.18 Configuring ERPS over VPLS
                    6.19 Maintaining VPLS
                    6.20 Troubleshooting VPLS




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                   586
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration




6.1 Overview of VPLS
Definition
                    The virtual private LAN service (VPLS) is an MPLS-based Ethernet point-to-
                    multipoint (P2MP) L2VPN service provided over a public network. VPLS ensures
                    that geographically isolated user sites can communicate over MANs and WANs as
                    if they were on the same LAN. VPLS is also called transparent LAN service (TLS).

                    Figure 6-1 shows a typical VPLS network topology in which users located in
                    different geographical regions communicate with each other through different
                    provider edges (PEs). For users, a VPLS network is a Layer 2 switched network that
                    allows them to communicate with each other in a way similar to communication
                    over a LAN.

                    Figure 6-1 Typical VPLS network topology




Purpose
                    As enterprises set up more and more branches in different regions and demand
                    greater office flexibility, applications such as voice over IP (VoIP), instant
                    messaging, and teleconferencing are used more and more widely. This places high
                    requirements on end-to-end (E2E) data communications technologies. A network
                    capable of providing P2MP services is the key to E2E data communication.

                    Traditional asynchronous transfer mode (ATM) and frame relay (FR) technologies
                    provide only Layer 2 point-to-point (P2P) connections. In addition, those network
                    types have drawbacks, such as high construction costs, low speed, and complex
                    deployment. The development of IP has paved the way for MPLS VPN technology,

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         587
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    which can provide VPN services over an IP network and offer advantages such as
                    easy configuration and flexible bandwidth control. MPLS VPNs can be classified
                    into MPLS L2VPNs and MPLS L3VPNs.

                    ●   MPLS L2VPNs, such as virtual private wire service (VPWS) networks, can
                        provide P2P services but not P2MP services over a public network.
                    ●   MPLS L3VPNs can provide P2MP services on the precondition that PEs keep
                        routes destined for end users. This implementation requires PEs to have
                        powerful routing capabilities.

                    To solve the preceding problems, VPLS, an MPLS-based Ethernet technology, is
                    introduced.

                    ●   Like Ethernet, VPLS supports P2MP communication.
                    ●   MPLS is a Layer 2 label switching technology. For users, the entire MPLS IP
                        backbone network is a Layer 2 switching device. PEs do not need to keep
                        routes destined for end users.

                    VPLS provides a more comprehensive P2MP service solution for Internet service
                    providers (ISPs). It integrates the advantages of Ethernet and MPLS technologies.
                    By emulating traditional LAN functions, VPLS enables geographically isolated users
                    on different Ethernet LANs to communicate with each other over the IP/MPLS
                    network provided by ISPs as if they were on the same LAN.


Benefits
                    VPLS offers the following benefits:

                    ●   VPLS networks can be constructed based on ISPs' IP networks, reducing
                        construction costs.
                    ●   VPLS networks provide the same high speed as Ethernet networks.
                    ●   VPLS networks allow users to communicate over Ethernet links, regardless of
                        whether these links are on WANs or LANs. This feature allows services to be
                        rapidly and flexible deployed.
                    ●   VPLS networks free ISPs from configuring and maintaining routing policies,
                        reducing operational expenditure (OPEX).


6.2 Understanding VPLS

6.2.1 VPLS Description

Basic VPLS Transmission Structure
                    Figure 6-2 shows an example of a VPLS network. The entire VPLS network is
                    similar to a switch. Pseudo wires (PWs) are established over MPLS tunnels
                    between VPN sites to transparently transmit Layer 2 packets between sites. When
                    forwarding packets, PEs learn the source MAC addresses of these packets and
                    create MAC address entries to map MAC addresses to attachment circuits (ACs)
                    and PWs.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             588
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    Figure 6-2 Basic VPLS transmission structure




                    The following describes the concepts related to VPLS networks.

                    Table 6-1 VPLS concepts
                     Item                Concept

                     AC                  A link between a CE and a PE. An AC must be established
                                         using Ethernet interfaces.

                     PW                  A bidirectional virtual connection between two virtual
                                         switch instances (VSIs) residing on two PEs. A PW consists
                                         of a pair of unidirectional MPLS VCs in opposite directions.

                     VSI                 A type of instance used to map ACs to PWs. A VSI
                                         independently provides VPLS services and forwards Layer 2
                                         packets based on MAC addresses and VLAN tags. A VSI has
                                         the Ethernet bridge function and can terminate PWs.

                     PW signaling        A type of signaling used to create and maintain PWs. PW
                                         signaling is the foundation for VPLS implementation.
                                         Typically, Label Distribution Protocol (LDP) and BGP are
                                         used as the PW signaling protocols.

