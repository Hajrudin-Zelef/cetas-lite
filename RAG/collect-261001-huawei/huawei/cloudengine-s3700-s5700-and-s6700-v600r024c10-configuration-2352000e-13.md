---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-13
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [875, 983]
sha256: a4c9a189e3b47c17f81fd852d09671fd4e25129141cfd85f567aa9b025c11cc8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        a.   Adds a GRE header to the packet. Specifically, the tunnel module
                             encapsulates the packet according to the passenger protocol type of the
                             packet and the key parameter configured for the current GRE tunnel.
                        b.   Adds a transport protocol header (an IP header) to the packet based on
                             the configuration (with the transport protocol IP). The source and
                             destination addresses contained in the IP header are the tunnel's source
                             and destination addresses, respectively.
                        c.   Delivers the packet to the IP module. The IP module searches the public
                             network routing table for the outbound interface and forwards the
                             packet based on the destination address in the IP header. The
                             encapsulated packet is then transmitted on the IP public network.
                    ●   Decapsulation
                        The decapsulation process is opposite to the encapsulation process. After the
                        egress PE receives the packet, the egress PE analyzes the IP header. After
                        determining that the destination of the packet is itself and the Protocol Type
                        field is 47, which indicates that the protocol is GRE, the egress PE delivers the
                        packet to the GRE module for processing. The GRE module removes the IP
                        header and GRE header, determines that the passenger protocol is a private
                        network protocol based on the Protocol Type field in the GRE header, and
                        delivers the packet to the private network protocol. The private network
                        protocol then forwards the packet as an ordinary packet.

Automatic TCP MSS Adjustment
                    Devices support dynamic adjustment of the maximum segment size (MSS) for SYN
                    or SYN-ACK packets during TCP connection setup.
                    Background
                    A SYN or SYN-ACK packet transmitted during TCP connection setup may contain
                    an MSS option field that specifies the MSS value on the local device. After the
                    local and remote devices exchange and compare their MSS values, they forward
                    packets based on the smaller MSS value to prevent packets from being
                    fragmented on the network. Provided that packets remain unfragmented, a larger
                    MSS value allows a greater amount of data to be sent per segment, increasing
                    network efficiency. MSS adjustment can minimize the possibility of packet
                    fragmentation and increase reliability for transmission of large data packets,
                    improving end-to-end TCP transmission efficiency.
                    Implementation
                    ●   If a SYN or SYN-ACK packet does not contain an MSS field, the device
                        automatically inserts an appropriate MSS value:
                        MSS = MTU – 40 – APPENDLEN
                        In the MSS value, MSS indicates the automatically inserted MSS value, MTU
                        indicates the maximum transmission unit (MTU) of an interface, and
                        APPENDLEN indicates the packet length added during VPN encryption and
                        encapsulation.
                    ●   If a SYN or SYN-ACK packet contains an MSS field, the device compares MSS-
                        APPENDLEN with MTU-40-APPENDLEN and adjusts the MSS value
                        accordingly:
                        –    If MTU-40-APPENDLEN is greater than MSS-APPENDLEN, the device
                             retains the original MSS value.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               9
VPN Configuration
VPN Configuration                                                                             2 GRE Configuration


                         –    If MTU-40-APPENDLEN is less than MSS-APPENDLEN, the device uses
                              MTU-40-APPENDLEN as the new MSS value.

                    Restrictions

                    ●    The MTU values of the interfaces that VPN packets pass through must be the
                         same.
                    ●    Automatic TCP MSS adjustment is performed only when the MTU value of the
                         outbound interface ranges from 256 to 9600 bytes.

2.2.2 Keepalive Detection
Background
                    The GRE protocol cannot detect the link status. As a result, if a remote interface
                    becomes unreachable, the tunnel connection between the local and remote ends
                    cannot be efficiently terminated. The local end continues to send data, which is
                    then discarded by the remote end due to the unreachable tunnel. This leads to a
                    data black hole.

                    To solve this problem, devices support the link status detection function, also
                    known as the keepalive function, on GRE tunnels. The keepalive function detects
                    whether a tunnel is in keepalive state, that is whether the remote end is
                    reachable. Once the remote end is unreachable, the local device terminates the
                    tunnel connection, preventing data black holes.

Implementation
                    After the keepalive function is enabled on the source end of a GRE tunnel, the
                    source end periodically sends keepalive messages to the destination end. If the
                    destination end is reachable, the source end receives a response message from the
                    destination end. Otherwise, no response message is received. The keepalive
                    detection process is as follows:
                    1.   After the keepalive function is enabled on the source end of a GRE tunnel, the
                         source end creates a timer, periodically sends keepalive messages, and counts
                         the number of failures to receive keepalive response messages. The
                         unreachable counter increases by one each time a message is sent.
                    2.   The destination end sends a response message to the source end each time it
                         receives a keepalive message from the source end.
                    3.   If the source end receives a response message before the number of sent
                         keepalive messages reaches the retry-times value, the source end considers
                         the destination end reachable and resets the number of sent keepalive
                         messages. If the source does not receive any response message before the
                         counter reaches the preset value, specifically, the retry times, the source
                         considers the peer unreachable and resets the counter. Then, the source closes
                         the tunnel connection.

                          NOTE

                         The keepalive function takes effect on one end of a tunnel as long as it is configured on
                         this end, regardless of whether it is configured on the other end. Once the destination end
                         receives a keepalive message, it sends a response message to the source end, regardless of
                         whether the keepalive function is configured on the destination end.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       10

