---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-180
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [26108, 26240]
sha256: d1eeedba7caa91dafef6b94d8c0c6fbee0039bd013f1c57211743c685420c5a3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Function Components
                    VPWS has the following function components:
                    ●    AC: An independent link or circuit connecting a CE and a PE. An AC interface
                         can be either physical or virtual. The AC attributes include the encapsulation
                         type, maximum transmission unit (MTU), and other interface parameters of
                         the specified link type.
                    ●    PW: A virtual link or path between two nodes on a network. In this document,
                         it is a virtual connection between two PEs.
                    ●    Tunnel: A connection used to transparently transmit service data.
                    ●    PW signaling: A type of signaling for PW negotiation.
                    The flow direction of VPN1 packets from CE1 to CE3 in Figure 5-1 is used as an
                    example to show data transmission.
                    1.   CE1 sends user packets to PE1 over an AC.
                    2.   PE1 receives the packets and selects a PW to forward the packets.
                    3.   PE1 generates double MPLS labels (one private network label and one public
                         network label) based on the PW forwarding entry. The private network label
                         identifies a PW, and the public network label identifies the tunnel between
                         PE1 and PE2.
                    4.   User packets are transmitted through the public network tunnel and arrive at
                         PE2, which then removes the private network label. The public network label
                         is removed using penultimate hop popping (PHP) on the P.
                    5.   PE2 selects an AC to forward these packets to CE3.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           413
VPN Configuration
VPN Configuration                                                                 5 VPWS Configuration


                    Figure 5-2 shows label changes in packet forwarding over a VPWS network.

                    Figure 5-2 VPWS label processing




                    In Figure 5-2:
                    ●   L2PDU: Layer 2 protocol data unit, a type of link layer packet.
                    ●   T: a tunnel label.
                    ●   V: a VC label.
                    ●   T': a substitute tunnel label during packet forwarding.

5.2.2 VPWS Classification
                    VPWS can be implemented in the following ways, depending on the scenario and
                    signaling protocols used:
                    ●   Circuit cross connect (CCC) VPWS
                        CCC VPWS needs to be manually configured by administrators and applies to
                        small MPLS networks with simple topologies. It does not require signaling
                        negotiation or exchange of control packets. It consumes few resources and is
                        easy to configure, but its maintenance is inconvenient — requiring manual
                        configuration by administrators — and its scalability is poor. For details, see
                        5.5.1 Understanding CCC VPWS.
                    ●   BGP VPWS
                        BGP VPWS uses BGP as the signaling protocol to transmit Layer 2 information
                        and VC labels between PEs. BGP VPWS uses VPN targets to control the
                        advertisement or acceptance of VPN routes, which improves networking
                        flexibility and scalability. For details, see 5.6.1 Understanding BGP VPWS.
                    ●   LDP VPWS
                        LDP VPWS establishes point-to-point links and uses LDP signaling to transmit
                        VC information. LDP VPWS has good scalability. It uses LDP signaling and
                        does not have to wait for PW status changes to periodically refresh, so LDP
                        VPWS fault detection is fast. For details, see 5.7.1 Understanding LDP VPWS.
                    ●   SVC VPWS
                        SVC VPWS uses inner labels manually configured on PEs for data
                        transmission, whereas LDP VPWS uses LDP to exchange VC labels. The SVC
                        mode can be regarded as the simplified LDP mode. For details, see 5.8.1
                        Understanding SVC VPWS.
                    Table 5-1 compares VPWS implementation modes.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           414
VPN Configuration
VPN Configuration                                                                  5 VPWS Configuration


                    Table 5-1 Comparison of VPWS implementation modes
                     Implem      Signalin   Tunnel                Applicati    Scalabi    Support for
                     entation    g                                on           lity       Local
                     Mode        Protocol                         Scenario                Connections

                     CCC         N/A        Local CCC: not        N/A          Compa      Yes
                                            requiring public                   ratively
                                            network tunnels                    poor
                                            Remote CCC: CR-
                                            LSP

                     LDP         LDP        Requiring a shared    Sparse       Poor       No
                                            LSP                   mode

                     SVC         N/A        Requiring a shared    N/A          Poor       No
                                            LSP

                     BGP         BGP        Requiring a shared    N/A          Good       Yes
                                            LSP




5.2.3 VCCV
                    Virtual circuit connectivity verification (VCCV) is an end-to-end fault detection and
                    diagnostic mechanism for PWs. VCCV is, in its simplest description, a control
                    channel between a PW's ingress and egress points over which connectivity
                    verification messages can be sent.
                    VCCV can detect and diagnose the connectivity of the forwarding paths on a PW.
                    VCCV uses the BFD CV field to indicate whether a PW supports BFD-based fault
                    detection. The BFD CV field value can be 0x04 or 0x08 but the default is 0x08. To
                    communicate with a peer that cannot identify the value 0x08, Huawei defines the
                    BFD CV field value carried in PW negotiation as 0x04 or 0x08.
                    ●   0x04: supports BFD IP/UDP encapsulation and is used for PW fault detection
                        only.
                    ●   0x08: supports BFD IP/UDP encapsulation and is used in PW fault detection
                        and AC/PW fault status signaling.
                    VCCV ping, an extension to LSP ping, is a tool used to manually verify the
                    connectivity of a VC. VCCV defines a series of messages exchanged between PEs to
                    verify PW connectivity. To ensure that VCCV packets and PW data packets are
                    transmitted along the same path, VCCV packets must be encapsulated in the same
                    way and transmitted over the same tunnel as PW data packets.


5.3 Configuration Precautions for VPWS

5.4 Default Settings for VPWS
                    The following table lists the default settings for VPWS.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            415
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration


