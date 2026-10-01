---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-252
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [37038, 37164]
sha256: 18d90afa1ac7c85d1fe096c31b02a32c2ba53dc5561d07050b9aeb8aa60a5225
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     Tunnel              A tunnel can carry multiple PWs. A tunnel is a direct
                                         channel that transparently transmits data between the
                                         local and remote PEs. It can be either an MPLS or Generic
                                         Routing Encapsulation (GRE) tunnel.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           589
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                     Item                 Concept

                     Forwarder            After a PE receives packets from an AC, the forwarder of
                                          the PE selects a PW to forward these packets. It is similar to
                                          a VPLS forwarding table.




                    The following example uses the forwarding of a packet from CE1 to CE3 on VPN1,
                    as shown in Figure 6-2, to describe the basic data flow direction:

                    1.   PE1, PE2, and PE3 belong to the same VPLS domain. ACs connected to the
                         VPLS domain are mapped to PWs through a VSI to generate the forwarder of
                         the VSI.
                    2.   CE1 receives a Layer 2 packet from a user at Site1, and forwards the Layer 2
                         packet to PE1 through an AC.
                    3.   PE1 receives the packet, discovers that the packet needs to be forwarded
                         using VPLS, and selects a PW from the forwarder to forward the packet based
                         on the destination MAC address of the packet.
                    4.   PE1 then generates two labels according to the PW forwarding entry and
                         tunnel information. The inner private network label identifies the PW, and the
                         outer public network label identifies the tunnel between PE1 and PE2. PE1
                         then searches for the destination MAC address based on the destination MAC
                         address entry index, encapsulates the packet, and forwards the packet.
                    5.   The Layer 2 packet arrives at PE2 through the tunnel on the public network.
                         The private network label then becomes the outer label (the public network
                         label has been popped at the penultimate hop).
                    6.   PE2 receives the packet, selects a VSI for forwarding the packet according to
                         the private network label, removes the private network label, and selects the
                         forwarder of the VSI. The forwarder then forwards the packet to CE3
                         according to the destination MAC address.

VPLS Implementation Process
                    Transmission of packets between customer edges (CEs) relies on VSIs configured
                    on PEs and PWs established between the VSIs, as shown in Figure 6-2. Figure 6-3
                    shows the transmission of Ethernet frames over full-mesh PWs between PEs.

                    On a Layer 2 Ethernet network, the Spanning Tree Protocol (STP) is typically
                    enabled to prevent loops. However, VPLS users do not know the topology of the
                    ISP network. Therefore, enabling STP on the private network cannot prevent loops
                    on the ISP network. To prevent loops, VPLS uses full-mesh PWs and split horizon.

                    ●    PEs on a VPLS network must be fully meshed with PWs. In other words, a PE
                         must create a tree path to every other PE on the VPLS network.
                    ●    Each PE must support split horizon to prevent loops. Split horizon requires
                         that packets received from public network PWs are forwarded only to private
                         networks, not to other public network PWs. This is why full-mesh PWs are
                         required between PEs in a VSI.

                    The full-mesh PEs and split horizon ensure route reachability and prevent loops on
                    VPLS networks. When CEs are connected to PEs through multiple connections or

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           590
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    CEs on the same VPLS VPN are connected, VPLS cannot prevent loops. In such a
                    situation, other methods must be used to prevent loops.

                    STP can run on an L2VPN, and all STP bridge protocol data units (BPDUs) are
                    transparently transmitted over the ISP network.

                    Figure 6-3 VPLS packet forwarding model




                    A PE on a VPLS network consists of a control plane and a forwarding plane.

                    ●   The control plane of a VPLS PE is responsible for PW establishment, including:
                        –   Member discovery: a process in which a PE in a VSI discovers other PEs in
                            the same VSI. You can manually configure member information or use a
                            protocol to discover members automatically.
                        –   Signaling mechanism: PWs between PEs with the same VSI ID are
                            established, maintained, or torn down using signaling protocols, such as
                            LDP and BGP.
                    ●   The forwarding plane of a VPLS PE is responsible for data forwarding over the
                        PW, including:
                        –   Encapsulation: After receiving Ethernet frames from a CE, a PE
                            encapsulates the frames into packets and sends the packets to a PSN.
                        –   Forwarding: A PE determines how to forward a packet based on the
                            inbound interface and destination MAC address of the packet.
                        –   Decapsulation: After receiving packets from a PSN, a PE decapsulates
                            these packets into Ethernet frames and sends the frames to a CE.


VPLS Encapsulation Types
                    ●   Packet encapsulation types on AC interfaces

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          591
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                        Packet encapsulation on AC interfaces depends on the user access mode,
                        which can be VLAN or Ethernet access. The default user access mode is VLAN
                        access.

                        Table 6-2 Packet encapsulation types on AC interfaces
                         Packet          Description
                         Encapsulat
                         ion Type
                         on AC
                         Interfaces

                         VLAN            Ethernet frames transmitted between CEs and PEs carry a
                         access          VLAN tag called a Provider-tag (P-tag). This is a service
                                         delimiter identifying users on an ISP network.

