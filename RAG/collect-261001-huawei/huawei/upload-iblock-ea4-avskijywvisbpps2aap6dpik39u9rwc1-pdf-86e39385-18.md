---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-18
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [1918, 2054]
sha256: e062aba4289055e5abf9257ac1ac74ff2a1afa65e31b9689903c09d667de7fbb
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

Multi-ISP Networking Adaptability 
The USG6000 series delivers features such as PBR and multi-interface NAT to 
improve the multi-IPS networking solution. Users can configure PBR to 
specify two interfaces to share traffic. If one interface is faulty, all the traffic 
fails over to the other interface by the USG6000 series. 
3.12 Excellent VPN Functions 
The USG6000 series provides IPSec mechanisms based on software or 
hardware encryption (DES, 3DES, AH, and ESP) to offer services such as 
access control, connectionless integrity, data source authentication, anti-replay, 
encryption, and data flow classification and encryption to both parties of the 
communications. Through Authentication Header (AH) and Encapsulating 
Security Payload (ESP), data transmitted at the IP layer or upper layers are 
protected, and the tunnel encapsulation mode is supported. 
In addition to supporting IPSec VPN application and providing highly reliable 
security transport channels, the USG6000 series can incorporate Layer 2 
Tunneling Protocol (L2TP) and Generic Routing Encapsulation (GRE) to 
provide diversified VPN applications: 
 L2TP VPN 
 IPSec VPN 
 GRE VPN 
 SSL VPN 
 L2TP over IPSec VPN 
 GRE over IPSec VPN 
GRE VPN 
GRE, a Layer 3 tunneling protocol of VPNs, can add an IP header on the IP 
packet. In other words, GRE adds a "coat" on private data for secure 
transmission. 
The USG6000 series not only supports the GRE VPN function, which sets up a 
GRE tunnel between two gateways to provide secure transmission, but also 
incorporates IPSec to provide diversified VPN applications. 
L2TP VPN 
The USG6000 series supports L2TP that implements the transparent 
transmission of PPP packets between users and enterprise servers, which is 
widely applied to access VPNs. Layer-2 data packets are encapsulated in a 
tunnel. For example, PPP packets are encapsulated in the L2TP tunnel. 
When serving as an LNS, the USG6000 series allows mobile users to initiate 
L2TP tunnel connections and requires mobile users to install VPN Client and 
know the IP address of the LNS. After receiving the requests of mobile users, 
the USG6000 series authenticates the mobile users based on the user name and 
password, allocates private addresses for mobile users, and establishes tunnels.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  44 
   
 
The USG6000 series, serving as a LAC, initiates L2TP tunnel connections for 
users when they access the Internet. Users can access the Internet using PPP or 
PPPoE. When the user name and password are authenticated on the LAC, you 
can identify L2TP tunnel users by user name. The LAC automatically initiates 
connections to the LNS, and the user then can access the enterprise VPN. 
Mobile users can use L2TP client software to connect to the LNS and access 
the headquarters intranet, but the IP address of the LNS, which is a private IP 
address, must be translated by the NAT server. 
IPSec VPN 
Using the IPSec mechanism, the USG6000 series provides security services 
such as access control, connectionless integrity, data source authentication, 
anti-replay, encryption, and data flow classification and encryption for 
communications parties. Data transmitted at the IP layer or upper layers is 
protected using AH and ESP , and the data can be encapsulated in tunnels. 
IPSec provides the following types of network security services: 
1. Privacy: Before transmitting packets, IPSec encrypts packets to ensure the 
data privacy. 
2. Integrity: IPSec verifies packets at the destination to ensure that the 
packets are not modified during the transmission. 
3. Authenticity: IPSec authenticates all the protected packets. 
4. Anti-replay: IPSec prevents packet retransmission. That is, the earlier or 
the repeated packets are denied at the destination. The packets are denied 
by packet sequence number. 
The USG6000 series uses the IPSec VPN to establish tunnels between the 
headquarters VPN gateway and branch VPN gateways and to obtain private 
addresses, securing the transmission and information. IPSec provides data 
protection between two hosts, two security gateways, or a host and a security 
gateway. Multiple security associations (SAs) can be established between two 
ends. By using ACLs and SAs, IPSec can apply different protection policies to 
data flows, to provide varied protection. IPSec SAs can be manually 
established. When the nodes on the network increase, it is difficult to configure 
SAs and ensure security. In this case, IKE is required to automatically establish 
SAs and implement key exchange. The IPSec VPN function of the USG6000 
series provides the certificate authentication mechanism based on the PKI 
framework. This mechanism supports certificate application, storage, and 
authentication, but not certificate generation. In addition, this mechanism 
supports digital envelop-based IKE negotiation. That is, certificate 
authentication is used during IKE negotiation. 
The USG6000 series supports IPSec VPN on IPv6, which provides VPN 
connectivity in IPv6 environment. 
BGP/MPLS VPN 
As a Layer 3 Virtual Private Network (L3VPN), BGP/MPLS IP VPN employs 
BGP to advertise VPN routes and MPLS to forward VPN packets on the 
backbone networks of ISPs. "IP" indicates that the VPN carries IP packets.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  45 
   
 
The Multiprotocol Label Switching (MPLS) technology combines the flexible 
IP routing and convenient asynchronous transfer mode (ATM) label switching. 
MPLS incorporates the connection-oriented control plane into the 
connectionless IP network to facilitate network management and operation. 
Therefore, an MPLS VPN that uses the MPLS-based IP network as the 
backbone network has become an important method for IP network carriers to 
provide value-added services and attracts more carriers. 
Unlike IGP , BGP focuses on controlling route advertisement and choosing the 
optimal route instead of finding and computing routes. VPN uses the public 
network to transmit data, where IGP route discovery and calculation have been 
applied. The primary concerns for constructing a VPN are controlling the 
spread of VPN routes and choosing the best route between two PEs. 
BGP uses TCP (port 179) as the transport protocol to improve reliability. Two 
USG6000-connected PE devices can run the BGP protocol to exchange VPN 
routes. 
BGP carries any information attached to routes as optional BGP attributes. The 
USG6000 series directly forwards the routes with any unknown attributes. 
Such processing facilitates the spread of VPN routes between PEs. 
BGP sends only the updated routes, instead of all, to reduce the bandwidth for 
route transmission, making it possible to transmit a large number of VPN 
routes on the public network. 
As an Exterior Gateway Protocol (EGP), BGP better applies to the VPN across 
carriers' networks. 
DSVPN 
Dynamic Smart VPN (DSVPN) is a technology that dynamically sets up a data 
forwarding tunnel between branches in the Hub-Spoke network model. 
On a traditional Hub-Spoke network, data mainly flows between the Spokes 
and the Hub. If data exchange is required between the Spokes and the IPSec 
technology has been applied, the Hub decrypts data over the branch tunnel for 
receiving data and re-encrypts data over the branch tunnel for sending data. 
The data exchanged between Spokes passes through the Hub, consumes Hub 
resources, and brings about delays. The DSVPN technology enables dynamic 
establishment of a data forwarding tunnel between Spokes, which resolves the 
previous issue.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  46 
   
 
 
