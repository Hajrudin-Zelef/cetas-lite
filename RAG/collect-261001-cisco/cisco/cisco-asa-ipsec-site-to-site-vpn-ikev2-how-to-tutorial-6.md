---
id: collect-261001-cisco/cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial-6
title: "cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial.md
source_anchor: ""
source_lines: [1755, 1767]
sha256: f5cf6f3366a51cd1dc89351c5930a0ce2ab473771f4f466e573661f087305de7
---

# cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial

1. IKEv2 uses only four messages for the initial exchange (packets 1-4)
2. IKE SA establishment and key generation (packets 1, 2) – IKE SA initial exchange
3. Identity authentication and the establishment of the first pair of IPsec SAs (packets 3, 4) – Authentication exchange
4. ESP packets n.5 – n.12 are our icmp encrypted traffic – the SPIs from source to destination and from destination to source are matching the SPIs in the ASA “show crypto ipsec details” details.
5. first packet (packet n.1) is carrying initial information such as: Encryption Algorithm, Integrity Algorithm, Diffie-Hellman group. This comes from the initiation ASA1 IPsec configuration. For the successful initiation of the IPsec Tunnel there should be a match between initiator proposal (in our case ASA1) and the responder proposal (in our case ASA2)

Let’s take a deeper look into the Payload:

The packet captures on both sides should be identical:

9. Conclusion

In this how to/tutorial I explained the basic theory to understand IPsec, how to configure Cisco ASA IPsec Tunnel with IKEv2, how to read and debug the IPsec parameter with help of tcpdump packet capture during tunnel initiation.
