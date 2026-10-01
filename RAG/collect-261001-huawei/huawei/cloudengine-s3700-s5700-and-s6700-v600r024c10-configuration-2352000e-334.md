---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-334
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [49944, 50079]
sha256: 9b6ce9f42d6aefe52602f4936ff2aaf1982bd15a1de9dd4e65b042c42fc0abad
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

7.1 Overview of Tunnel Management
                    Only the S6780-H, S6750-H, S6750E-S, S6750-S, S6730E-H-V2, S6730-H-V2,
                    S5755E-H, S5755-H, and S5732-H-V2 series support LDP tunnel management and
                    MPLS TE tunnel management.

Definition
                    The tunnel management (TNLM) is a module used to select a certain tunnel for
                    an application according to specific configurations and notifies the application of
                    the tunnel status.

Purpose
                    By managing tunnels, devices can better use the tunneling technology to establish
                    a dedicated data transmission channel on the backbone network to transparently
                    transmit packets.


7.2 Understanding Tunnel Management
                    Tunnel management involves the introduction to common tunnels, tunnel
                    configuration management, and tunnel policies.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            803
VPN Configuration
VPN Configuration                                                  7 Tunnel Management Configuration


Common Tunnels
                    Common tunnels are as follows:

                    ●   GRE
                        Generic Routing Encapsulation (GRE) is a tunneling protocol that
                        encapsulates the packets of a wide variety of network layer protocols, such as
                        Internetwork Packet Exchange (IPX) and AppleTalk, inside IP tunneling
                        packets. These packets can then be transmitted over an IPv4 network.
                        GRE provides a mechanism for transporting the packets of one protocol over
                        another protocol by means of encapsulation, enabling packets to be
                        transmitted over heterogeneous networks. The channel used for transmitting
                        these packets is called a tunnel.
                    ●   LDP
                        MPLS LDP is a label distribution protocol widely used for VPN services. MPLS
                        LDP features simple networking and configuration, allows routing topology-
                        driven LSP establishment, and supports large-capacity label switched paths
                        (LSPs).
                        LDP is an MPLS control protocol that provides functions similar to those of a
                        signaling protocol on a traditional network. LDP directs packets to forwarding
                        equivalence classes (FECs), assigns labels to packets, and establishes and
                        maintains LSPs. LDP defines the messages used in label distribution as well as
                        the message processing procedures.
                    ●   MPLS TE
                        With MPLS deployed, carriers are generally required to provide VPN users with
                        end-to-end QoS guarantees for various services, such as the audio, video,
                        mission-critical, and regular Internet access services. To meet user
                        requirements, MPLS TE tunnels can be used to optimize network resources
                        and offer users QoS-guaranteed services.


Tunnel Configuration and Management
                    The setup and management of tunnels vary according to tunnel types. This
                    chapter focuses on the following aspects:

                    ●   Tunnel interface configuration: You can specify different types of tunnels on a
                        tunnel interface. The configuration methods of different types of tunnels on
                        an interface are different.
                    ●   General tunnel management: notifies the tunnel status to the application that
                        uses the tunnel and provides the tunnel query policy to determine how to
                        select a tunnel. The tunnel policy function is commonly used.


7.3 Configuration Precautions for Tunnel Management

7.4 Configuring a Tunnel Policy

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           804
VPN Configuration
VPN Configuration                                                    7 Tunnel Management Configuration


7.4.1 Understanding Tunnel Policies

Tunnel Policy
                    A tunnel policy is used for an application module to select tunnels.

                    Tunnel policies can be categorized as either tunnel type prioritization policies or
                    tunnel binding policies. The two types of tunnel policies are mutually exclusive.

                    Tunnel Type Prioritizing Policy

                    In a tunnel type prioritizing policy, you can specify the sequence in which each
                    type of tunnel is selected and the maximum number of tunnels that can
                    participate in load balancing. This type of policy applies only to CR-LSPs, GRE
                    tunnels. Tunnels defined in a tunnel type prioritizing policy are selected in
                    sequence. The tunnels with the type specified first are selected as long as the
                    tunnels of this type are up, regardless of whether the tunnels are selected by other
                    services. Generally, the tunnels of a later specified type are not selected except
                    when load balancing is required or when the preceding tunnels are all down.

                    Tunnel Binding

                    In a tunnel binding policy, you can bind one destination address to a tunnel. Then,
                    VPN services to which the policy is applied will be transmitted over the bound
                    tunnel. The system does not check whether the bound tunnel is a TE tunnel, and
                    the tunnel binding policy takes effect only on TE tunnels. Therefore, ensure that
                    the tunnel binding policy is correctly configured. In Figure 7-1, two MPLS TE
                    tunnels (tunnel 1 and tunnel 2) are set up between PE1 and PE3.

                    Figure 7-1 Application of a tunnel binding policy




                    If you bind VPNA to tunnel 1 and VPNB to tunnel 2, VPNA and VPNB use separate
                    MPLS TE tunnels. This means that tunnel 1 serves only VPNA and tunnel 2 serves
                    only VPNB. In this manner, services of VPNA and VPNB are isolated from each

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             805
VPN Configuration
VPN Configuration                                                      7 Tunnel Management Configuration


                    other and also from other services. The bandwidth for VPNA and VPNB is
                    therefore ensured, which facilitates later QoS deployment.

                    In tunnel binding, you can bind one destination address to one or more TE tunnels
                    to load-balance services. In addition, you can configure the down-switch attribute
                    to enable other types of tunnels to be selected when the specified tunnels are
                    unavailable, ensuring traffic continuity.

                    A common tunnel binding policy selects common TE tunnels based on destination
                    addresses and tunnel interface indexes. A tunnel binding policy observes the
                    following tunnel selection rules:

