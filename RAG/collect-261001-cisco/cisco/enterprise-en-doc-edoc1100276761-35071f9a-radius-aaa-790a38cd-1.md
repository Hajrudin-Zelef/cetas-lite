---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-1
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [1, 42]
sha256: 1f7b21853b98f0dc414a187eefd8dc64036eef57e54d095ad37290eb0b3d9ff3
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

Enterprise
AAA can be implemented using multiple protocols. RADIUS is most frequently used in actual scenarios.
RADIUS is a protocol that uses the client/server model in distributed mode and protects a network from unauthorized access. It is often used on networks that require high security and control remote user access. It defines the UDP-based RADIUS packet format and transmission mechanism, and specifies destination UDP ports 1812 and 1813 as the default authentication and accounting ports respectively.
At the very beginning, RADIUS was only the AAA protocol used for dial-up users. As the user access mode diversifies, such as Ethernet access, RADIUS can also be applied to these access modes. RADIUS provides the access service through authentication and authorization and records the network resource usage of users through accounting.
RADIUS has the following characteristics:
Client/Server model
Secure message exchange mechanism
Fine scalability
RADIUS client
RADIUS clients run on the NAS to transmit user information to a specified RADIUS server and process requests (for example, permit or reject user access requests) based on the responses from the server. RADIUS clients can locate at any node on a network.
As a RADIUS client, a device supports:
standard RADIUS protocol and its extensions, including RFC 2865 and RFC 2866
RADIUS server status detection
retransmission of Accounting-Request(Stop) packets in the local buffer
active/standby and load balancing functions between RADIUS servers
RADIUS server
RADIUS servers typically run on central computers and workstations to maintain user authentication and network service access information. The servers receive connection requests from users, authenticate the users, and send all required information (such as permitting or rejecting authentication requests) to the clients. A RADIUS server generally needs to maintain three databases, as shown in Figure 1-6.
Authentication messages between a RADIUS server and RADIUS clients are exchanged using a shared key. The shared key is a character string that is transmitted in out-of-band mode, is known to both clients and the server, and does not need to be transmitted independently on the network.
A RADIUS packet has a 16-octet Authenticator field that contains the digital signature data of the whole packet. The signature data is calculated using the MD5 algorithm and shared key. The RADIUS packet receiver needs to verify whether the signature is correct and discards the packet if the signature is incorrect.
This mechanism improves security of message exchange between RADIUS clients and the RADIUS server. In addition, user passwords contained in RADIUS packets are encrypted using shared keys before the packets are transmitted to prevent the user passwords from being stolen during transmission on an insecure network.
A RADIUS packet consists of a packet header and a certain number of attributes. The protocol implementation remains unchanged even if new attributes are added to a RADIUS packet.
RADIUS is based on the UDP protocol. Figure 1-7 shows the RADIUS packet format.
Attribute: This field is variable in length. RADIUS attributes carry the specific authentication, authorization, accounting information and configuration details for the request and reply packets. The Attribute field may contain multiple attributes, each of which consists of Type, Length, and Value. For details, see RADIUS Attributes.
RADIUS defines 16 types of packets. Table 1-5 describes types of the authentication packets, Table 1-6 describes types of the accounting packets. For RADIUS CoA/DM packets, see RADIUS CoA/DM.
| Packet Name | Description | 
|---|---|
| Access-Request | Access-Request packets are sent from a client to a RADIUS server and is the first packet transmitted in a RADIUS packet exchange process. This packet conveys information (such as the user name and password) used to determine whether a user is allowed access to a specific NAS and any special services requested for that user. | 
| Access-Accept | After a RADIUS server receives an Access-Request packet, it must send an Access-Accept packet if all attribute values in the Access-Request packet are acceptable (authentication success). The user is allowed access to requested services only after the RADIUS client receives this packet. | 
| Access-Reject | After a RADIUS server receives an Access-Request packet, it must send an Access-Reject packet if any of the attribute values are not acceptable (authentication failure). | 
| Access-Challenge | During an EAP relay authentication, when a RADIUS server receives an Access-Request packet carrying the user name from a client, it generates a random MD5 challenge and sends the MD5 challenge to the client through an Access-Challenge packet. The client encrypts the user password using the MD5 challenge, and then sends the encrypted password in an Access-Request packet to the RADIUS server. The RADIUS server compares the encrypted password received from the client with the locally encrypted password. If they are the same, the server determines the user is valid. | 
| Packet Name | Description | 
|---|---|
| Accounting-Request(Start) | If a RADIUS client uses RADIUS accounting, the client sends this packet to a RADIUS server before accessing network resources. | 
| Accounting-Response(Start) | The RADIUS server must send an Accounting-Response(Start) packet after the server successfully receives and records an Accounting-Request(Start) packet. | 
| Accounting-Request(Interim-update) | You can configure the real-time accounting function on a RADIUS client to prevent the RADIUS server from continuing user accounting if it fails to receive the Accounting-Request(Stop) packet. The client then periodically sends Accounting-Request(Interim-update) packets to the server, reducing accounting deviation. | 
| Accounting-Response(Interim-update) | The RADIUS server must send an Accounting-Response(Interim-update) packet after the server successfully receives and records an Accounting-Request(Interim-update) packet. | 
| Accounting-Request(Stop) | When a user goes offline proactively or is forcibly disconnected by the NAS, the RADIUS client sends this packet carrying the network resource usage information (including the online duration and number of incoming/outgoing bytes) to the RADIUS server, requesting the server to stop accounting. | 
| Accounting-Response(Stop) | The RADIUS server must send an Accounting-Response(Stop) packet after receiving an Accounting-Request(Stop) packet. | 
A device that functions as a RADIUS client collects user information, including the user name and password, and sends the information to the RADIUS server. The RADIUS server then authenticates users according to the information, after which it performs authorization and accounting for the users. Figure 1-8 shows the information exchange process between a user, a RADIUS client, and a RADIUS server.
The RADIUS server verifies the user identity:
When a user is authenticated, a device sends an Access-Request packet to the RADIUS server. To ensure that the device can receive a response packet from the server even if a network fault or delay occurs, a retransmission upon timeout mechanism is used. The retransmission times and retransmission interval are controlled using timers.
As shown in Figure 1-9, 802.1X authentication and client-initiated authentication are used as an example. After receiving an EAP packet (EAP-Response/Identity) containing the user name of the client, the device encapsulates the packet into a RADIUS Access-Request packet and sends the packet to the RADIUS server. The retransmission timer is enabled at the same time. The retransmission timer is composed of the retransmission interval and retransmission times. If the device does not receive any response packet from the RADIUS server when the retransmission interval expires, it sends a RADIUS Access-Request packet again.
