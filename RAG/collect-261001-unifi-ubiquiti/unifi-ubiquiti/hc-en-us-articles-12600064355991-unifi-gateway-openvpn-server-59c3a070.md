---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-12600064355991-unifi-gateway-openvpn-server-59c3a070
title: "hc-en-us-articles-12600064355991-unifi-gateway-openvpn-server-59c3a070"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-12600064355991-unifi-gateway-openvpn-server-59c3a070.md
source_anchor: ""
source_lines: [1, 20]
sha256: 7cee3b4074724293c4dfacedbd0a9b91eaf8442ec5af7e86405639b763b9d4e7
---

# hc-en-us-articles-12600064355991-unifi-gateway-openvpn-server-59c3a070

UniFi Gateway - OpenVPN Server
OpenVPN is a VPN server found in the VPN section of your Network application that allows you to connect to the UniFi network from a remote location.
Requirements
- A Next-Gen UniFi gateway or UniFi Cloud Gateway
- UniFi Network application version 7.4 or newer.
How does it work?
After enabling OpenVPN and specifying a port (default OpenVPN port is 1194), add a User and share the configuration file with your desired recipient. Once the recipient has installed the OpenVPN program or mobile app, they can import the configuration and easily remotely access the UniFi network at any time.
Frequently Asked Questions
1. Should I use Teleport, Wireguard or OpenVPN?
For mobile client devices, it is recommended to use Teleport.
For desktop or laptop clients, Teleport and Wireguard are recommended over OpenVPN.
2. Is OpenVPN secure?
OpenVPN encrypts your traffic and secures remote access connections. It also uses private and public keys.
3. How does OpenVPN compare with other VPNs, and can you use them simultaneously?
OpenVPN provides lower throughput than Wireguard. OpenVPN can be used alongside other VPNs.
4. Can OpenVPN be used when the UniFi gateway is behind NAT?
If the UniFi gateway is behind NAT, then the port used for OpenVPN needs to be forwarded by the upstream router.
We recommend using OpenVPN on a UniFi gateway that has access to a public IP address. Any performance or port forwarding issues on the upstream router can cause the VPN to disconnect.
5. Which clients support OpenVPN?
OpenVPN is supported by many different clients. See the OpenVPN installation page for more information.
