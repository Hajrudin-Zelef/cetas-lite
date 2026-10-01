---
id: collect-261001-cisco/cisco/enterprise-intl-en-thread-common-ospf-operations-on-s-series-switches-6672186617-b1897602
title: "enterprise-intl-en-thread-common-ospf-operations-on-s-series-switches-6672186617-b1897602"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/enterprise-intl-en-thread-common-ospf-operations-on-s-series-switches-6672186617-b1897602.md
source_anchor: ""
source_lines: [1, 77]
sha256: 618523b2e0ce12dd63501ba317977a7c92ff09ac6b1afddc26ece63b03d6a97c
---

# enterprise-intl-en-thread-common-ospf-operations-on-s-series-switches-6672186617-b1897602

· General packet errors: indicates the statistics about general packet errors.
o IP: received my own packet: indicates that a device receives a packet with the source address being the IP address of the device.
Cause: A Layer 2 loop occurs.
o Bad packet: indicates that an error packet is received.
The causes are listed as follows:
o Bad version: indicates that the version number is incorrect.
Cause: The OSPF version number contained in the received packet is not 2.
o Bad area id: indicates that a packet is received from an area that the interface does not belong to.
Cause: The area whose ID is changed is restarted and a retransmitted packet is received.
o Drop on unnumbered interface: indicates the number of packets discarded by an unnumbered P2P interface.
Cause: The OSPF network type configured on the unnumbered P2P interface is not P2P.
o Bad virtual link: indicates the number of error packets received on the Vlink.
o Bad authentication type: indicates that the authentication type is incorrect.
o Bad authentication key: indicates that the authentication key is incorrect.
o Packet too small: indicates that a packet of an improper length is received.
Cause: The sum of the IP header length and packet length is different from the length of the packet received by socket.
o Transmit error: indicates that the packet is incorrectly transmitted.
Cause: The packet fails to be transmitted.
o Interface down: indicates the number of OSPF interfaces in the Down state.
Cause: The count is increased by 1 each time an OSPF interface becomes Down.
o Unknown neighbor: indicates that there is no corresponding neighbor.
o Netmask mismatch: indicates that the network mask is unmatched.
Cause: The mask contained in the header of the Hello packet received in a non-P2P network is inconsistent with the mask of the interface that receives the packet.
· HELLO packet errors: indicates the statistics about Hello packet errors.
o Hello timer mismatch: The values of the Hello timers on the two ends of a link are inconsistent.
Cause: The set values of the Hello timers on the two ends of a link are different.
o Dead timer mismatch: indicates the values of the dead timers on the two ends of a link are inconsistent.
Cause: The values of the dead timers on the two ends of a link are different.
o Extern option mismatch: indicates that the Option bits in Hello packets mismatch.
o Router id confusion: indicates that router IDs conflict.
Cause: The router IDs of two neighboring devices conflict (note: this count is applicable to only the case that the router IDs of two neighboring devices conflict, and is inapplicable to router ID conflict on the entire network).
o Virtual neighbor unknown: indicates an unknown neighbor on a Vlink.
Cause: The router ID in the received Hello packet is different from that of the neighbor on the Vlink.
o NBMA neighbor unknown: indicates an unknown NBMA neighbor.
Cause: The network type of the neighbor contained in the received Hello packet is NBMA, but the NBMA neighbor does not exist.
o Invalid Source Address: indicates that the source address of the packet is invalid.
Cause: The Hello packet is received but there is no corresponding neighbor.
· DD packet errors: indicates the statistics about DD packet errors.
o Neighbor state low: indicates that status of the neighbor that receives the DD packet is low.
o MTU option mismatch: indicates that the MTU in the received DD packet is greater than that configured on the interface.
Cause: The MTU in the received DD packet is greater than the OSPF MTU configured on the interface.
· LS ACK packet errors: indicates the statistics about LSAck packet errors.
o Neighbor state low: is the same as that in a DD packet.
o Bad ack: indicates that an ACK packet in which the LSA is unmatched with that in the retransmission list in terms of the contents and age is received.
Cause: The LSA contained in the ACK packet sent by the neighbor is unmatched with that in the retransmission list in terms of the contents and age.
o Duplicate ack: indicates duplicate ACK packets.
Cause: An Ack packet is replied several times for an LSA. This count is not considered as an error on a broadcast network.
· LS REQ packet errors: indicates the statistics about LSRequest packet errors.
o Bad request: indicates that an incorrect request packet is received.
· LS UPD packet errors: indicates the statistics about Update packet errors.
o Received less recent LSA: indicates that an LSA that is older than that in the LSDB is received.
Cause: An older LSA is received, that is, the device has a new LSA.
· Opaque errors: indicates statistics about Opaque LSA errors.
o 9-out of flooding scope: indicates that the flooding of Type 9 LSAs is out of the flooding range.
Cause: Type 9 Opaque LSAs are received but the Opaque capability is not enabled.
o 10-out of flooding scope: indicates that the flooding of Type 10 LSAs is out of the flooding range.
Cause: Type 10 Opaque LSAs are received but the Opaque capability is not enabled.
o 11-out of flooding scope: indicates that the flooding of Type 11 LSAs is out of the flooding range.
Cause: Opaque LSAs are received when the Opaque capability is not enabled, or Type 11 LSAs are received when the current area does not support the As-external-lsa capability.
· Retransmission for packet over Limitation errors: indicates that the number of packet retransmission times exceeds the set limit.
o Number for DD Packet: indicates that the number of DD packet retransmission times exceeds the set limit.
Cause: After the limit on the number of packet retransmission times is set, the neighbor becomes Down and the count is increased by 1, each time the number of DD packet retransmission times exceeds the set limit.
o Number for Update Packet: indicates that the number of Update packet retransmission times exceeds the set limit.
Cause: After the limit on the number of packet retransmission times is set, the neighbor becomes Down and the count is increased by 1, each time the number of Update packet retransmission times exceeds the set limit.
o Number for Request Packet: indicates that the number of request packet retransmission times exceeds the set limit.
Cause: After the limit on the number of packet retransmission times is set, the neighbor becomes Down and the count is increased by 1, each time the number of LS Request packet retransmission times exceeds the set limit.
· Receive Grace LSA errors: indicates the statistics about received Grace LSA errors.
o Number of invalid LSAs: indicates the number of invalid Grace LSAs.
Cause: The received Grace LSA is incorrectly parsed.
o Number of policy failed LSAs: indicates the number of policy failures.
Cause: The GR helper policy fails and the device fails to enter the helper state.
o Number of wrong period LSAs: indicates that the value of the GR timer contained in the received Grace LSA is incorrect.
· Configuration errors: indicates the statistics about configuration errors.
o Tunnel cost mistake: indicates that the cost of the tunnel is incorrect.
Cause: The cost of the tunnel is not greater than 0.
o The network type of the neighboring interface is not consistent: indicates that the network type of the local interface is inconsistent with that of the neighboring interface.
Cause: The network types configured on the neighboring interfaces may be inconsistent.
