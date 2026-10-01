---
id: collect-261001-cisco/cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial-1
title: "cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial.md
source_anchor: ""
source_lines: [1, 153]
sha256: 9d092f8f4a8c58fc7b73c3b4b4efe39e687e505b0e9204eede2682c6a345df35
---

# cisco-asa-ipsec-site-to-site-vpn-ikev2-how-to-tutorial

In this tutorial I will explain the basic knowledge to understand IPsec VPN and this knowledge will be demonstrated on real use-case in the Testlab environment. I will configure two Cisco ASA Firewalls to demonstrate establishing IPsec connection using IKEv2 between these endpoints as well as packet capture for establishing the IPsec VPN connection.

1. Intro

**Internet Protocol Security** (**IPsec**) is a secure network protocol suite that authenticates and encrypts packets of data to provide secure encrypted communication between two computers over an Internet Protocol network. It is used in virtual private networks (VPNs).


When two computers establish a VPN connection, they must agree on a set of security protocols and encryption algorithms and exchange cryptographic keys to unlock and view the encrypted data. This is where IPsec enters the picture. IPsec works with VPN tunnels to establish a private bidirectional connection between devices. IPsec is not a single protocol, rather it is a complete set of protocols and standards that work together to help ensure the confidentiality, integrity, and authentication of Internet data packets flowing through a VPN tunnel. Here’s how IPsec creates a secure VPN tunnel:

- It authenticates data to ensure the integrity of the data packet in transit
- Encrypts Internet traffic through VPN tunnels so that the data cannot be viewed
- Protects against data replay attacks that can lead to unauthorized logins
- Enables secure exchange of cryptographic keys between computers
- Offers two security modes: tunnel and transport

IPsec VPN protects the flow of data from host to host, network to network, host to network, and gateway to gateway (called tunnel mode when the entire IP packet is encrypted and authenticated).


##### IPsec protocols and components

The IPsec standard is divided into several basic protocols and supporting components.

###### Basic IPsec protocols:

**Authentication Header (AH)**: This protocol protects the IP addresses of computers involved in a data exchange to ensure that bits of data are not lost, changed, or corrupted during transmission. AH also verifies that the person who sent the data actually sent it, protecting the tunnel from infiltration by unauthorized users.


**Encapsulating Security Payload (ESP)**: The ESP protocol provides the encryption portion of IPsec, which ensures the confidentiality of data traffic between devices. ESP encrypts data packets/payloads and authenticates the payload and its origin in the IPsec protocol suite. This protocol effectively encrypts Internet traffic so that no one looking at the tunnel can see what is there.


**Internet Key Exchange (IKE)**: For encryption to work, the computers involved in the exchange of private communications must share encryption keys. IKE allows two computers to securely exchange and share cryptographic keys when establishing a VPN connection. IKE establishes the SA between the communicating hosts, negotiating the cryptographic keys and algorithms that will be used in the course of the session.

There are two versions of IKE:

  - IKEv1
  - IKEv2

IKEv1 was introduced around 1998 and superseded by IKEv2 in 2005. There are some differences between the two versions:

- IKEv2 requires less bandwidth than IKEv1.
- IKEv2 supports EAP authentication (next to pre-shared keys and digital certificates).
- IKEv2 has built-in support for NAT traversal (required when your IPsec peer is behind a NAT router).
- IKEv2 has a built-in keepalive mechanism for tunnels.

IKE uses two phases:

- IKE Phase 1
- IKE Phase 2

ESP encrypts and authenticates data, while AH only authenticates data.


###### IPsec components:

**Security Associations (SA)**: defines some factors of communication peers like the protocols, operational modes, encryption algorithms (DES, 3DES, AES-128, AES-192 and AES-256), shared keys of data protection in particular flows and the life cycle of SA, etc. SA is used to process data flow in one direction. Therefore, in a bi-directional communication between two peers, you need at least two security associations to protect the data flow in both of the directions.


**Replay protection**: IPSec also includes standards that prevent replay of any data packets that are part of a successful login process. This standard prevents hackers from using the replayed information to replicate the login themselves.


**Encryption and hashing algorithms**: the cryptographic key works by using a hash value that is generated by a hashing algorithm. AH and ESP are general in that they do not specify a particular type of encryption. Encryption algorithms protect the data so it cannot be read by a third-party while in transit. Authentication algorithms verify the data integrity and authenticity of a message.


**IPsec encryption algorithms:**

- **AES** (Advanced Encryption Standard) — AES is the strongest encryption algorithm available. IPsec devices can use AES encryption keys of these lengths 128, 192, or 256 bits. AES is faster than 3DES.
- **3DES** (Triple-DES) — An encryption algorithm based on DES that uses the DES cipher algorithm three times to encrypt the data. The encryption key is 168-bit. 3DES is slower than AES.
- **DES** (Data Encryption Standard) — Uses an encryption key that is 56 bits long. DES is the weakest of the three algorithms, and it is considered to be insecure.

**IPsec authentication algorithms:**

- MD5 (Message Digest Algorithm 5) – MD5 produces a 128-bit (16 byte) message digest, which makes it faster than SHA1 or SHA2. This is the least secure algorithm.
- SHA1 (Secure Hash Algorithm 1) – SHA1 produces a 160-bit (20 byte) message digest. Although slower than MD5, this larger digest size makes it stronger against brute force attacks. SHA-1 is considered to be mostly insecure because of a vulnerability.
- SHA2 (Secure Hash Algorithm 2) – SHA2 is the most secure algorithm. There are three variants of SHA2 with different message digest lengths:
  - SHA2-256 — produces a 265-bit (32 byte) message digest
  - SHA2-384 — produces a 384-bit (48 byte) message digest
  - SHA2-512 — produces a 512-bit (64 byte) message digest

SHA2 is stronger than either SHA1 or MD5. We recommend that you specify a SHA2 variant.


**Tunneling modes:** tunnel and transport 

IPsec sends data using tunnel or transport mode. These modes are closely related to the type of protocols used, AH or ESP.

- Tunnel mode: In tunnel mode, the entire packet is protected. IPsec wraps the data packet in a new packet, encrypts it and adds a new IP header. It is commonly used in site-to-site VPN setups.
- Transport mode: In transport mode, the original IP header remains and is not encrypted. Only the payload and ESP trailer are encrypted. Transport mode is often used in client-to-web VPN setups.

When it comes to VPNs, the most common IPSec configuration you will probably see is ESP with tunnel mode authentication. This structure helps Internet traffic to move securely and anonymously inside the VPN tunnel over insecure networks.


(For a full technical explanation of IPsec, I highly recommend the excellent article on Networklessons)

2. Prerequisites

- Cisco ASA 9.8.1 (ASA1)
- Cisco ASA 9.8.1 (ASA2)
- Cisco IOS Router 15.9 (R1)
- PC1 (some sort of Endpoint behind ASA1)
- PC2 (some sort of Endpoint behind ASA2)

3. Lab setup

4. Cisco ASA1 configuration

**a)** Interface configuration:

!

interface GigabitEthernet0/0

nameif outside

security-level 0

ip address 192.0.2.6 255.255.255.252

!

interface GigabitEthernet0/1

nameif inside

security-level 100

ip address 10.0.0.254 255.255.255.0

!

**b)** routing configuration:

route outside 0.0.0.0 0.0.0.0 192.0.2.5 1

**c)** IPsec phase 2 configuration:

crypto ipsec ikev2 ipsec-proposal MY_PROPOSAL

protocol esp encryption aes-256

protocol esp integrity sha-1

**d)** IPsec phase 1 configuration:

crypto ikev2 policy 1

encryption aes-256

integrity sha

group 2

prf sha

