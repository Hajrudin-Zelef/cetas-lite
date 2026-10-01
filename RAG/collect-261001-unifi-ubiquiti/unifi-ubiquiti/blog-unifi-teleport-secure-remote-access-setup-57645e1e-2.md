---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-teleport-secure-remote-access-setup-57645e1e-2
title: "blog-unifi-teleport-secure-remote-access-setup-57645e1e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["SpaceX"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-teleport-secure-remote-access-setup-57645e1e.md
source_anchor: ""
source_lines: [66, 78]
sha256: 1bf2a5d0b53e357385510cd2fd6255e13a2548449dc4f5f4d2fcdf6478a2fd30
---

# blog-unifi-teleport-secure-remote-access-setup-57645e1e

Ubiquiti says Teleport works when both the gateway and the client are behind NAT, and it leaves Teleport out of the VPNs that need a public IP. The article doesn't mention CGNAT by name, and some setups need IPv6 on the gateway's WAN.
Is Teleport the same as the WireGuard VPN server?
No. Teleport uses WireGuard for encryption, but Teleport handles the connection setup through invitations and WiFiman. The WireGuard server uses configuration files and an open port on the gateway.
Can I remove a Teleport user?
Yes. Revoke a pending invitation in Invitation History, or remove an accepted one from the Teleport client in Client Devices.
Final thoughts
Teleport is the quickest way to give one person access to one UniFi network. It needs no public IP, no port forwarding and no config files, and WiFiman now covers Windows, Mac, Linux and phones.
It's also a per-network tool. When you look after many client networks, the daily work happens in Site Manager and the controller behind it. For that part, UniHosted runs UniFi OS Server for MSPs, with daily backups and tested updates on paid plans.
Related guides
Keep reading
- UniFi WireGuard VPN: how to set it up (2026)Set up the WireGuard VPN server on a UniFi gateway: port 51820, client files and QR codes, split tunneling, the WireGuard client, and when to use Teleport. Read guide
- How to optimize UniFi for remote work with Starlink and TeleportSet up a reliable remote work network with Starlink and UniFi by optimizing settings, enabling Teleport VPN, and managing remote access smoothly. Read guide
- How to run OpenVPN on your Ubiquiti networkSet up a secure OpenVPN server on your UniFi device and gain remote access to your home or office network from anywhere in the world. Read guide
