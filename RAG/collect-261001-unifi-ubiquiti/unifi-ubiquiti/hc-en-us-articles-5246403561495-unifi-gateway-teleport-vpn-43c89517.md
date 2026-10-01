---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-5246403561495-unifi-gateway-teleport-vpn-43c89517
title: "hc-en-us-articles-5246403561495-unifi-gateway-teleport-vpn-43c89517"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-5246403561495-unifi-gateway-teleport-vpn-43c89517.md
source_anchor: ""
source_lines: [1, 13]
sha256: d80eedb935abaf6714a2aa9c75c72cd1cc33432fa690f76b6d1ae7cea3a16a59
---

# hc-en-us-articles-5246403561495-unifi-gateway-teleport-vpn-43c89517

UniFi Gateway - Teleport VPN
Teleport is a zero-configuration VPN that allows you to instantly connect to your UniFi network from a remote location. Users with a Next-Gen gateway or UniFi Cloud Gateway running UniFi OS can access it from Network Settings > VPN.
How Does it Work?
After enabling Teleport, you can generate an invitation and share it with your desired recipient. If the recipient already has the WiFiman Mobile App (iOS / Android) or WiFiman Desktop installed, the invitation will automatically add the VPN to the app. If not, the invitation will prompt the user to install the app. Once installed, the invitation will add the Teleport VPN when it is clicked again.
Once the recipient has accepted the Teleport invitation, they can easily and securely access the UniFi network remotely, at any time.
Note: The invitation expires in 24 hours and can only be utilized by a single device at a time.
Note: Teleport requires an IPv6 connection on your gateway's WAN in some circumstances. It is not necessary to enable IPv6 on LAN.
Frequently Asked Questions
If the invitation is still pending, it can be revoked in Invitation History. If the invitation has already been accepted by the recipient, access can be revoked by viewing the Teleport client in the Client Devices section.
Teleport uses the Wireguard VPN to encrypt your traffic and secure remote access connections. Your data is not stored by Teleport.
Teleport’s Wireguard integration allows it to deliver higher throughput than traditional VPNs. Teleport doesn’t reserve addresses/ports and can be used alongside other VPNs. Teleport will update client DNS servers to use the UniFi Gateway's servers to which it is connected.
Yes. Unlike traditional VPNs such as L2TP, which encounter issues when behind NAT, Teleport can be used when both the UniFi gateway and client are behind NAT.
Yes, Teleport VPN is available through the WiFiman Desktop app.
