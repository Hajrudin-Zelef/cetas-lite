---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-12
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [735, 874]
sha256: aa48fd182729893630ab9ffc904e483b898fd6945b019533888519935202667d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Benefits
                    GRE offers the following benefits:
                    ●   Enables packets to be transmitted between networks running different
                        protocols using a single network protocol.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              5
VPN Configuration
VPN Configuration                                                                  2 GRE Configuration


                    ●   Enlarges the scope of a hop-limited network.
                    ●   Connects discontiguous subnets for VPN setup.


2.2 Understanding GRE

2.2.1 GRE Fundamentals
Background
                    A single network protocol, such as IPv4, is often used to transmit packets on a
                    backbone network, whereas other network protocols, such as IPv6 and IPX, may
                    be used to transmit packets on non-backbone networks. If the backbone and non-
                    backbone networks use different protocols, packets cannot be transmitted
                    between them. GRE solves this issue by providing a mechanism for transporting
                    the packets of one protocol over another protocol through encapsulation.
                    In Figure 2-1, networks 1 and 2 are the non-backbone networks running Novell
                    IPX, and networks 3 and 4 are the non-backbone networks running IPv6. The
                    backbone network is an IPv4 network. To transmit packets between networks 1
                    and 2 and between networks 3 and 4 over the backbone network, GRE can be
                    used to establish a tunnel between DeviceA and DeviceB. When DeviceA receives a
                    packet from network 1 or network 3, it encapsulates the packet into a GRE packet,
                    which is then encapsulated into an IPv4 packet and forwarded through the
                    established tunnel to network 2 or network 4.

                    Figure 2-1 GRE networking




Related Concepts
                    GRE packet format
                    After receiving the packet of a specific network layer protocol (for example, IPX)
                    that needs to be encapsulated and routed, the system adds a GRE header to the
                    packet and encapsulates it into the packet of another network protocol, such as IP.
                    The packet can then be forwarded by the IP protocol. Figure 2-2 shows the format
                    of an encapsulated GRE packet.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            6
VPN Configuration
VPN Configuration                                                                    2 GRE Configuration


                    Figure 2-2 Format of an encapsulated GRE packet




                    ●   Payload: an inner packet that needs to be encapsulated and transmitted to a
                        destination network.
                    ●   Passenger protocol: the protocol that requires encapsulation.
                    ●   Encapsulation protocol: the protocol used to encapsulate the passenger
                        protocol, also known as the carrier protocol.
                    ●   Transport protocol: the protocol used to transmit encapsulated packets.
                    The following shows the format of an IPX packet encapsulated for transmission
                    over an IP tunnel.

                    Figure 2-3 Format of an IPX packet encapsulated for transmission over an IP
                    tunnel




                    GRE header
                    Figure 2-4 shows the format of a GRE header.

                    Figure 2-4 GRE header format




                    The following describes the fields in a GRE header:
                    ●   C: checksum present bit. If set to 0, the Checksum field is not present in the
                        GRE header.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               7
VPN Configuration
VPN Configuration                                                                              2 GRE Configuration


                    ●   K: key present bit. If set to 0, the Key field is not present in the GRE header.
                    ●   Recursion: number of GRE encapsulations. This field is incremented by 1 after
                        each GRE encapsulation. If the number of GRE encapsulations is greater than
                        3, the packet is discarded. This field prevents packets from being infinitely
                        encapsulated.
                             NOTE

                            The Recursion field defaults to 0.
                            If the field values on the transmit end and receive end are different, no error will
                            occur, and the receive end will ignore the field.
                            This field takes effect only during GRE encapsulation. When a device decapsulates a
                            GRE packet, it is unaware of this field.
                    ●   Flags: reserved field. Currently, the field value must be set to 0.
                    ●   Version: version field. The field value must be 0.
                    ●   Protocol Type: passenger protocol type.
                    ●   Checksum: checksum of the GRE header and the payload.
                    ●   Key: key field. It is used by the receive end to authenticate received packets.
                    Currently, a GRE header does not contain the source route field. As such, Bits 1, 3,
                    and 4 are all set to 0.

Packet Transmission over a GRE Tunnel
                    Packet transmission over a GRE tunnel consists of both encapsulation and
                    decapsulation. In Figure 2-5, a private network packet from the ingress PE to the
                    egress PE is encapsulated on the ingress PE and decapsulated on the egress PE.

                    Figure 2-5 Communication between private networks over a GRE tunnel




                    ●   Encapsulation
                        The ingress PE receives a private network packet from an interface connected
                        to a private network, and delivers the packet to the protocol module running
                        on the private network for processing. The VPN protocol module searches the
                        VPN routing or forwarding table for an outbound interface based on the
                        destination address carried in the VPN packet header. If the outbound
                        interface is a GRE tunnel interface, this module sends the packet to the tunnel
                        module.
                        The tunnel module processes the received packet as follows:

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                          8
VPN Configuration
VPN Configuration                                                                    2 GRE Configuration


