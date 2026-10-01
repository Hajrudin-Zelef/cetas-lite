---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47-1
title: "docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "cyber", "latency", "lean", "memory", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47.md
source_anchor: ""
source_lines: [1, 70]
sha256: c2ba43305df010a1f2b72882a1290efb0f09350760b8f54ee7afafc976992c8b
---

# docs-network-security-tutorials-how-to-setup-wireguard-on-opnsense-112d7c47

OPNsense Wireguard: Setup and Configuration
WireGuard is a simple and fast VPN protocol that employs modern cryptography. It aims to be faster and less complicated than IPsec, as well as a significantly more performant alternative to OpenVPN. It has a smaller codebase, which makes auditing and maintaining it easier. It was originally designed for the Linux kernel, but it is now cross-platform and widely deployed. We believe that the best VPN fo OPNsense is WireGuard.
In this guide, we walk you through the complete setup of WireGuard on OPNsense, a powerful open-source firewall and routing platform. Whether you're a network administrator or a home user, this tutorial will help you verify WireGuard availability, configure the VPN server, add peer devices like Windows or Android clients, and test your connection to ensure secure remote access. From initial setup to verifying tunnel handshakes, all critical steps are covered in a clear, step-by-step format.
Figure 1. WireGuard P2P VPN Topology (OPNsense WireGuard Server and Android/Windows Clients)
You can use this WireGuard setup to securely access your company or home network from all over the world.
You can quickly install Zenarmor NGFW Plug-in on your OPNsense VPN server to protect remote employees from cyber threats.
Zenarmor NGFW Plug-in for OPNsense is one of the most popular OPNsense plug-ins and allows you to easily upgrade your firewall to a Next Generation Firewall in seconds. NG Firewalls empower you to combat modern-day cyber attacks that are becoming more sophisticated every day.
Some of the capabilities are layer-7 application/user aware blocking, granular filtering policies, commercial-grade web filtering utilizing cloud-delivered AI-based Threat Intelligence, parental controls, and the industry's best network analytics and reporting.
Zenarmor Free Edition is available at no cost for all OPNsense users.
What Makes WireGuard a Secure and Fast VPN Option for OPNsense?
WireGuard has rapidly gained recognition as one of the most secure, lightweight, and high-performance VPN solutions available today. Its seamless integration with OPNsense, a widely used open-source firewall and routing platform, makes it a powerful choice for both home and enterprise networks. The following characteristics explain why WireGuard is particularly well-suited for OPNsense users.
- 
Modern and Robust Encryption Protocols: WireGuard utilizes modern cryptographic primitives such as ChaCha20 for encryption, Curve25519 for key exchange, Poly1305 for data authentication, and BLAKE2s for hashing. These algorithms are considered among the most secure and efficient in current use. Unlike older VPN protocols like OpenVPN or IPsec, WireGuard avoids supporting legacy cryptographic options, reducing the attack surface significantly. The use of ChaCha20 is especially beneficial for performance on low-powered devices, making it an excellent choice for embedded or resource-constrained environments running OPNsense.
- 
Lightweight and Auditable Codebase: WireGuard’s codebase is remarkably small, consisting of only about 4,000 lines of code. This minimalist design stands in stark contrast to older VPN protocols, which often contain hundreds of thousands of lines of code. A smaller codebase not only enhances performance but also facilitates comprehensive security audits, reducing the likelihood of vulnerabilities. OPNsense, being a security-focused platform, benefits greatly from WireGuard’s lean and clean implementation.
- 
Superior Performance and Faster Connections: WireGuard is designed to be connectionless and operates over UDP, enabling it to establish tunnels almost instantly. It employs a stateless handshake mechanism and supports roaming, meaning that clients can change IP addresses without interrupting the VPN connection. This leads to noticeably lower latency and faster reconnection times compared to traditional VPN protocols. OPNsense users experience quicker VPN sessions and smoother remote access when using WireGuard.
- 
Simplicity in Configuration: One of WireGuard’s key advantages is its configuration simplicity. It relies on static key pairs for authentication and uses clearly defined [Interface] and[Peer] configuration blocks. This straightforward structure eliminates the complexities found in OpenVPN setups, such as certificate authorities, TLS configurations, and Diffie-Hellman key exchanges. On OPNsense, administrators can easily configure WireGuard using the intuitive web interface, minimizing errors and accelerating deployment.
- 
Kernel-Level Integration for Enhanced Efficiency: WireGuard is implemented directly within the kernel on supported systems such as Linux and FreeBSD. This allows it to avoid the overhead associated with context switches between user space and kernel space, resulting in significantly better performance. Since OPNsense is based on FreeBSD, WireGuard’s kernel-level support ensures optimal throughput and lower CPU utilization, especially under high network loads.
- 
Seamless Integration with OPNsense GUI: OPNsense provides native support for WireGuard through a user-friendly graphical interface. Administrators can manage peers, generate QR codes for mobile devices, and monitor handshake status directly from the OPNsense dashboard. This integration streamlines the deployment process and allows for easier management and troubleshooting, reducing the operational burden on network administrators.
- 
Efficient Use of System Resources: Due to its optimized architecture and efficient cryptographic design, WireGuard consumes fewer CPU cycles, less memory, and overall lower system resources than traditional VPN solutions. This makes it particularly effective for use on embedded firewall appliances and small form-factor hardware, where OPNsense is commonly deployed. It ensures that other essential services, such as IDS/IPS, routing, and traffic shapin,g can run concurrently without degradation in VPN performance.
In summary, WireGuard’s combination of strong security practices, minimal code complexity, high performance, and ease of use makes it an ideal VPN solution for OPNsense environments. Whether deployed in home networks or enterprise settings, WireGuard provides a modern and reliable approach to secure communications with minimal overhead.
How to use WireGuard Server?
To successfully set up and use a WireGuard server on OPNsense, follow the step-by-step instructions below. These steps cover enabling the service, configuring the instance, and setting up client peers.
1. Enable the WireGuard Service:
Before you can start configuring the server, you need to activate the WireGuard service in the OPNsense interface with the following steps.
- 
Navigate to VPN → WireGuard → General .
- 
Check the box for Enable WireGuard.
- 
Click Save to apply the changes.
2. Create a Instance (Server Configuration):
Once the service is enabled, you need to create a WireGuard instance that will act as the server with the following steps.
- 
Go to VPN → WireGuard → Instances .
- 
Click on + Add to create a new instance.
- 
Fill in the required fields: 
  - 
Name: Give your instance a recognizable name (e.g., MyWireGuardServer).
  - 
Listen Port: Enter a UDP port number, such as 51820.
  - 
Private Key: Click the button to auto-generate a private key.
  - 
Public Key: This is automatically derived and shown after saving.
- 
- 
Click Save.
3. Add Peer (Client) Configuration:
After setting up the server instance, the next step is to add peers clients that will connect to your WireGuard server.
- 
Go to VPN → WireGuard → Peers .
- 
Click on + Add to create a new peer configuration.
- 
Fill in the following details: 
  - 
Name: Choose a name like MyAndroidClient or MyWindowsClient.
  - 
Public Key: Paste the public key generated on the client device.
  - 
Allowed IPs: Enter the client’s virtual IP address in CIDR format (e.g., 10.0.0.2/32).
  - 
