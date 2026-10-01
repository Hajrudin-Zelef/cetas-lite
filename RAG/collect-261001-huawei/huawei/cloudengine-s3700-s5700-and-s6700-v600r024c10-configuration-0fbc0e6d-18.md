---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-18
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1871, 1985]
sha256: 5face0d354771cd78b7b8db8d48245556d9af49851e7cd0720a0ad421d4f5b99
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                               31
MPLS Configuration
MPLS Configuration                                                          3 MPLS LDP Configuration


3.2.3 LDP MTU
                 The maximum transmission unit (MTU) size defines the maximum number of
                 bytes that a device can transmit in a packet. It plays an important role when two
                 devices communicate on a network. If the MTU value of a packet exceeds the
                 maximum size supported by a receiver or transit device, the packet is fragmented
                 during transmission, increasing the workload of the network. The packet may even
                 be discarded during transmission, affecting services. To prevent such issues,
                 configure a device to calculate the MTU size before sending packets for
                 communication.

Fundamentals
                 LDP LSP forwarding and common IP forwarding differ greatly in terms of
                 implementation mechanism but share a large number of similar aspects about the
                 MTU. If MTU values are correctly negotiated before packet transmission, packets
                 can successfully reach the receiver without packet fragmentation and reassembly.
                 For a FEC, an LSR calculates the smallest value among all MTU values advertised
                 by preferred next-hop LSRs as well as the MTU value of the local outbound
                 interface. The LSR then adds the calculated MTU value into the MTU TLV of Label
                 Mapping messages to be sent to the upstream device. When the MTU value
                 changes (for example, when the local outbound interface or the MTU
                 configuration changes), the LSR recalculates the MTU value and advertises the
                 new value to its upstream device through Label Mapping messages.

3.2.4 Coexistent Local and Remote LDP Session
Fundamentals
                 For a coexistent local and remote LDP session, the local and remote LDP
                 adjacencies are connected to the same peers so that both the adjacencies are used
                 to maintain the peers.
                 On the network shown in Figure 3-3, when the local LDP adjacency is deleted due
                 to a failure in the link to which the adjacency is connected, the peer type may
                 change, without affecting the peer's presence or status. (The peer type is
                 determined by the adjacency type, which can be local, remote, or coexistent local
                 and remote.)
                 When a link becomes faulty or is recovering from a fault, the peer type may
                 change. If this is the case, the type of the session associated with the peer changes
                 accordingly. However, the session is not deleted or set to down.




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                             32
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


Application Scenario

                 Figure 3-3 Coexistent local and remote LDP session




                 Coexistent local and remote LDP sessions are typically used in L2VPN scenarios.
                 On the network shown in Figure 3-3, L2VPN services are deployed between PE1
                 and PE2. The following describes how to configure a coexistent local and remote
                 LDP session and how the services are processed if the directly connected link
                 between PE1 and PE2 becomes faulty and then recovers.

                 1.     Enable MPLS LDP on PE1 and PE2, which are directly connected, to set up a
                        local session between them. Configure PE1 and PE2 as each other's remote
                        peer to set up a remote session between them. In this case, both a local
                        adjacency and a remote adjacency are set up between PE1 and PE2. The
                        session between them is a coexistent local and remote LDP session. L2VPN
                        signaling messages can then be transmitted through the session.
                 2.     When the physical link between PE1 and PE2 goes down, so too does the
                        local LDP adjacency. The route between PE1 and PE2 is still reachable through
                        the P, which means that the remote LDP adjacency remains up. The session
                        changes to a remote session so that it can remain up. The L2VPN does not
                        detect the session change or delete the session. As such, the L2VPN does not
                        need to disconnect and then recover services, preventing service interruptions
                        in this process.
                 3.     When the fault is rectified, the link between PE1 and PE2 and the local LDP
                        adjacency go up again. The session then reverts to a coexistent local and
                        remote LDP session and remains in the up state. Again, the L2VPN does not
                        detect the session change or delete the session, preventing service
                        interruptions in this process.

3.2.5 Distributing Labels to All Peers
                 Distributing labels to all peers helps to prevent slow convergence in the case that
                 a link becomes faulty.

                 If labels are distributed only to upstream peers, the upstream and downstream
                 relationships of a session need to be confirmed based on routing information
                 when Label Mapping messages are sent. An upstream node cannot send Label
                 Mapping messages to its downstream node along a route. If the route changes
                 and the upstream/downstream relationship is switched, the new downstream
                 node resends Label Mapping messages. This results in slow convergence.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              33
MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration


                 If distributing labels to all peers is enabled, each node can send Label Mapping
                 messages to all peers, regardless of the upstream and downstream relationships.

                 On the network shown in Figure 3-4, the original routes from P2 to PE3 are P2 ->
                 P1 -> P3 -> PE3 and P2 -> P4-> PE4 ->PE3. For the loopback interface route on
                 PE3, P1 is the next hop of P2. When labels can be distributed only to upstream
                 nodes and P2 receives a Label Mapping message from P1, P2 does not send the
                 Label Mapping message associated with the route to P1. If the link between P1
                 and P3 is faulty, the route from PE1 to PE3 is switched from PE1 -> P1 -> P3 ->
                 PE3 to PE1 -> P1 -> P2 -> P4 -> P3 -> PE3, and P2 becomes the downstream node
                 of P1. The LSP can be set up only after P2 resends a Label Mapping message.
                 However, P2 does not send a Label Mapping message to P1, resulting in slow LSP
                 re-convergence.

