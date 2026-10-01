---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-vpn-asa-912-vpn-config-vpn-rem-aa176f3c
title: "c-en-us-td-docs-security-asa-asa912-configuration-vpn-asa-912-vpn-config-vpn-rem-aa176f3c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-vpn-asa-912-vpn-config-vpn-rem-aa176f3c.md
source_anchor: ""
source_lines: [1, 46]
sha256: c1eaabe6a7a0d75d52d747fc87e0e9d0b4ffaa05b3acacd826bed2317b9fd49f
---

# c-en-us-td-docs-security-asa-asa912-configuration-vpn-asa-912-vpn-config-vpn-rem-aa176f3c

Overview of Remote Access IPsec VPNs
Remote access VPNs allow users to connect to a central site through a secure connection over a TCP/IP network. The Internet Security Association and Key Management Protocol, also called IKE, is the negotiation protocol that lets the IPsec client on the remote PC and the ASA agree on how to build an IPsec Security Association. Each ISAKMP negotiation is divided into two sections called Phase1 and Phase2.
Phase 1 creates the first tunnel to protect later ISAKMP negotiation messages. Phase 2 creates the tunnel that protects data travelling across the secure connection.
To set the terms of the ISAKMP negotiations, you create an ISAKMP policy. It includes the following:
- 
                              		  
                              An authentication method, to ensure the identity of the peers.
- 
                              		  
                              An encryption method, to protect the data and ensure privacy.
- 
                              		  
                              A Hashed Message Authentication Codes (HMAC) method to ensure the identity of the sender and to ensure that the message has not been modified in transit.
- 
                              		  
                              A Diffie-Hellman group to set the size of the encryption key.
- 
                              		  
                              A time limit for how long the ASA uses an encryption key before replacing it.
A transform set combines an encryption method and an authentication method. During the IPsec security association negotiation with ISAKMP, the peers agree to use a particular transform set to protect a particular data flow. The transform set must be the same for both peers.
A transform set protects the data flows for the ACL specified in the associated crypto map entry. You can create transform sets in the ASA configuration, and then specify a maximum of 11 of them in a crypto map or dynamic crypto map entry. For more overview information, including a table that lists valid encryption and authentication methods, see Create an IKEv1 Transform Set or IKEv2 Proposal.
You can configure the ASA to assign an IPv4 address, an IPv6 address, or both an IPv4 and an IPv6 address to the AnyConnect Client by creating internal pools of addresses on the ASA or by assigning a dedicated address to a local user on the ASA.
The endpoint must have the dual-stack protocol implemented in its operating system to be assigned both types of addresses. In both scenarios, when no IPv6 address pools are left but IPv4 addresses are available or when no IPv4 address pools are left but IPv6 addresses are available, connection still occurs. The client is not notified; however, so the administrator must look through the ASA logs for the details.
Assigning an IPv6 address to the client is supported for the SSL protocol.
About Mobike and Remote Access VPNs
Mobile IKEv2 (mobike) extends ASA RA VPNs to support mobile device roaming. This support means the end-point IP address for a mobile device’s IKE/IPSEC security association (SA) can be updated rather than deleted when the device moves from its current connection point to another.
Mobike is available by default on ASAs since version 9.8(1), meaning Mobike is “always on.” Mobike is enabled for each SA only when the client proposes it and the ASA accepts it. This negotiation occurs as part of the IKE_AUTH exchange.
After the SA is established with mobike support as enabled, client can change its address anytime and notify the ASA using the INFORMATIONAL exchange with UPDATE_SA_ADDRESS payload indicating the new address. The ASA will process this message and update the SA with the new client IP address.
| Note | You can use the                                            		  show crypto ikev2 sa detail command to determine                                           		  whether mobike is enabled for all current SAs. | 
The current Mobike implementation supports the following:
-  
                                 		  
                                 IPv4 addresses only
-  
                                 		  
                                 Changes in NAT mappings
-  
                                 		  
                                 Path connectivity and outage detection, by means of optional Return Routability checking
-  
                                 		  
                                 Active/standby failover
-  
                                 		  
                                 VPN load balancing
If the Return Routability Check (RRC) feature is enabled, an RRC message is sent to the mobile client to confirm the new IP address before the SA is updated.
