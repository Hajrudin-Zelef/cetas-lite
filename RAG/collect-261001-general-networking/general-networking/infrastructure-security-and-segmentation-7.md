---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-7
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [473, 563]
sha256: 4bdc35edbc471cc73437aa7f430b71664a479d6f3dc3ea44a9c066f4e7c73b2e
---

# infrastructure-security-and-segmentation

Example 2-57 shows the ACL from Example 2-50 modified to allow RDP traffic every day between midnight and 2:00 a.m. It uses the recurring time range shown in Example 2-53 to accomplish this. Note how the second ACE in the example is changed from **deny** to **permit** with the time range, while a new line is added below it to deny RDP traffic during other times. If the new line is not added, RDP traffic is permitted by the **permit ip any any** ACE at all times.

#### **Example 2-57** *Using a Time Range in an ACL*

`Router(config)#**access-list 101 deny tcp any any eq 23**
Router(config)#**access-list 101 permit tcp any any eq 3389 time-range daily**
Router(config)#**access-list 101 deny tcp any any eq 3389**
Router(config)#**access-list 101 deny ip host 10.1.1.1 host 10.2.1.1**
Router(config)#**access-list 101 permit ip any any**
Router(config)#**interface Gi1**
Router(config-if)#**ip access-group 101 in**`

#### Reflexive ACLs

It is difficult to keep track of all traffic that should be allowed between various networks connected to a router. This is especially true when one of those segments is a public network, such as the Internet or an untrusted network. For example, when a router connects to the Internet, traffic in response to queries sent from inside is expected and should be allowed, whereas any other traffic not in response to an explicit query should not be allowed.

Extended ACLs in IOS allow tracking of sessions in one direction while allowing return traffic in the other. This is done by marking each ACE that needs session tracking in an outbound ACL with the **reflect** keyword. Every time a new session matching that ACE is encountered, a reflexive ACE entry is created in an inbound ACL that is marked with the **evaluate** keyword.

Each reflexive ACE is a temporary entry in the inbound ACL that matches the IP address pair, protocol, and port pair of the original outbound session that triggered the reflect. The temporary ACE is removed after the original session ends. The end of a session is determined by TCP packets with FIN or RST bits set or when a certain configurable time has passed without any packets seen for that session. The latter option is always used for non-TCP sessions.

To further understand how reflexive ACLs work, consider the network diagram shown in Figure 2-4. The figure shows router R1 connected to the internal network 10.1.1.0/24 on interface Gi1 and to the Internet on interface Gi3. On Interface Gi3, an outbound ACL called **allow-out** and an inbound ACL called **allow-in** are configured. The contents of the two ACLs are shown in Example 2-58.

#### **Example 2-58** *Initial Configuration for a Reflexive ACL*

`R1#**show running-config**
—removed for brevity—
ip access-list extended allow-in
 permit tcp any eq www any
 permit tcp any eq 443 any
 deny   ip any any
ip access-list extended allow-out
 permit tcp any any eq www
 permit tcp any any eq 443
 deny   ip any any
!
interface GigabitEthernet3
 ip address dhcp
 ip nat outside
 ip access-group allow-in in
 ip access-group allow-out out
 negotiation auto
—removed for brevity—`

Notice that the inbound ACL allows all traffic sourced from ports 80 and 443. While this permits responses from web servers, it also permits an attacker to reach any inside destination with a packet sourced from port 80 or 443. Hence, this is not a very secure or desirable configuration.

With reflexive ACL, the router in Figure 2-4 can be configured such that all web traffic going out of Gi3 will be evaluated, and a reflexive ACE will be created in the **allow-in** ACL to allow only return traffic to come in from the Internet.

To configure this, the two ACLs on R1 are changed as shown in Example 2-59. Notice that the **allow-in** ACL contains a single **evaluate** command pointing to the reflected ACL **racl**. IOS creates temporary ACEs in the **allow-in** ACL as sessions, matching the two permit statements in the **allow-out** ACL, and exits the Gi3 interface toward the Internet.

#### **Example 2-59** *Configuring Reflexive ACL*

`R1#**show running-config**
—removed for brevity—
ip access-list extended allow-in
 evaluate racl
ip access-list extended allow-out
 permit tcp any any eq www reflect racl
 permit tcp any any eq 443 reflect racl
!
interface GigabitEthernet3
 ip address dhcp
 ip nat outside
 ip access-group allow-in in
 ip access-group allow-out out
 negotiation auto
—removed for brevity—`

#### Unicast Reverse Path Forwarding

One of the key uses of traffic filtering is to prevent traffic with spoofed sources. While ACLs can be used to do this, creating and maintaining ACLs for this purpose can be cumbersome. IOS provides a nifty security feature called Unicast Reverse Path Forwarding (URPF) to simplify detection of spoofed IP packets.

When Unicast RPF is enabled on an interface, the router examines all packets received on that interface to see if the packet was received at an interface that is the best return path back to the source of the packet. In other words, it does a reverse lookup, or a “look back,” on each packet to see if the source of the packet appears in its forwarding table and is connected to the same interface where the packet was received.

If the router determines that the source of the packet cannot be reached through the interface where the packet was received, it drops the packet.

URPF can be used in any “single-homed” environment, where there is essentially only one access point out of the network (that is, one upstream connection). Networks having one access point offer the best example of symmetric routing, which means the interface where a packet enters the network is also the best return path to the source of the IP packet. URPF is best used at the network perimeter for Internet or other untrusted networks or in ISP environments for customer network terminations.

URPF can be configured on an interface with the **ip verify unicast reverse-path** command and can be verified with the **show ip interface** *interface* command.

#### Network Based Application Recognition (NBAR)

While filtering traffic with IP access lists is a good security practice, it has its shortcomings. Lot of applications and protocols, such as Real Time Protocol (RTP), use dynamic port ranges and are difficult to block with ACLs. For a better security posture, it is important to understand and limit traffic based on applications and protocols. This is where the Network Based Application Recognition (NBAR) feature of IOS helps.

NBAR is a classification engine that recognizes and classifies a wide variety of protocols and applications, including web-based and other difficult-to-classify applications and protocols that use dynamic TCP/UDP port assignments.

As traffic flows through interfaces configured for NBAR discovery, it recognizes and classifies protocols or applications that the packets belong to. The router can be configured to apply QoS policies based on the NBAR classifications, including dropping packets belonging to unwanted applications and protocols. The following are some of the key uses of NBAR:

- **Classification by HTTP header:** NBAR can be used to classify traffic based on HTTP header information, including URL, host, MIME, and user agent. This classification can be used to filter traffic going to known malicious sites or coming from unknown clients.
- **Classification of peer-to-peer (P2P) file-sharing traffic:** NBAR can be used to identify and block traffic belonging to P2P file-sharing applications such as BitTorrent to prevent data loss.
- **Classification of real-time traffic:** NBAR can be used to classify and filter or rate limit RTP or Real Time Streaming Protocol (RTSP) to prevent misuse of network resources.

