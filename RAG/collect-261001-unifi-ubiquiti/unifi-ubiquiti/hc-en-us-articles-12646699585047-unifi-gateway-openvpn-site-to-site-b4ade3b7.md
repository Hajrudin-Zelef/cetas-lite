---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-12646699585047-unifi-gateway-openvpn-site-to-site-b4ade3b7
title: "hc-en-us-articles-12646699585047-unifi-gateway-openvpn-site-to-site-b4ade3b7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-12646699585047-unifi-gateway-openvpn-site-to-site-b4ade3b7.md
source_anchor: ""
source_lines: [1, 31]
sha256: 061ca0dcf81fed099ca8c8003aa1e6a1c493cc1e708be8ee0bd7e7fcd83f8fa4
---

# hc-en-us-articles-12646699585047-unifi-gateway-openvpn-site-to-site-b4ade3b7

UniFi Gateway - OpenVPN Site-to-Site
OpenVPN is a Site-to-Site VPN that uses a 2048 bit static key for authentication. It can be configured in the VPN section of your Network application settings.
Requirements
- A UniFi Gateway or UniFi Cloud Gateway
How to Configure
- Navigate to the OpenVPN Site-to-Site settings in Network > Settings > VPN.
- SSH into your gateway using these instructions.
- Run the following command: 
 openvpn --genkey secret | egrep -o "[0-9a-z]{32}" | sed ':a; N; $!ba; s/\n//g'
- Copy the 512 character string and paste it in the Pre-shared Key field. Be sure to delete any spaces or line breaks.
- Configure the remaining required fields on the Site Manager:
  - 
Local and Remote Tunnel IP Address: IP addresses used inside the VPN tunnel.
Note: We recommend using private IP addresses that do not overlap with any other networks.
  - Local and Remote Port: This will generally be UDP port 1194.
  - Remote Networks: Network(s)/subnet(s) used at the remote location.
  - Remote IP Address: Public IP address or hostname of the remote location.
- 
Local and Remote Tunnel IP Address: IP addresses used inside the VPN tunnel.
Frequently Asked Questions
1. Should I use IPsec or OpenVPN Site-to-Site VPNs?
It is recommended to use IPsec as it provides higher throughput.
2. Is OpenVPN secure?
OpenVPN encrypts your traffic and secures the VPN connection. It also uses a 2048 bit static key for authentication.
3. How does OpenVPN compare with IPsec Site-to-Site VPNs, and can you use them simultaneously?
IPsec provides higher throughput than OpenVPN. Both VPNs can be used simultaneously.
4. Can OpenVPN be used when the UniFi gateway is behind NAT?
If the UniFi gateway is behind NAT, then the port used for OpenVPN needs to be forwarded by the upstream router.
We recommend using OpenVPN on a UniFi gateway that has access to a public IP address. Any performance or port forwarding issues on the upstream router can cause the VPN to disconnect.
5. How can I generate the OpenVPN key?
See the OpenVPN documentation page for more information.
