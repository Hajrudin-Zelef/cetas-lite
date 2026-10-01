---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-7983431932439-unifi-gateway-site-to-site-ipsec-vpn-with-third-cf47b643
title: "hc-en-us-articles-7983431932439-unifi-gateway-site-to-site-ipsec-vpn-with-third--cf47b643"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-7983431932439-unifi-gateway-site-to-site-ipsec-vpn-with-third--cf47b643.md
source_anchor: ""
source_lines: [1, 69]
sha256: 92e51a3a3e77dc9a6e221d4daa1c1544d1bf6065699a7a19832dff3efad81695
---

# hc-en-us-articles-7983431932439-unifi-gateway-site-to-site-ipsec-vpn-with-third--cf47b643

UniFi Gateway - Site-to-Site IPsec VPN with Third-Party Gateways (Advanced)
IPsec is a Site-to-Site VPN that allows you to connect a UniFi gateway to a remote location. A UniFi Gateway or UniFi Cloud Gateway is required.
How does it work?
IPsec Site-to-Site VPNs use a Pre-Shared Key for authentication. A unique key is automatically generated but a custom key can be used as well.
Additionally, the following information is required:
- Local IP: Use the IP address assigned to the WAN port or enter a manual address.
- Remote Networks: Network(s) used at the remote location.
- Remote IP / Hostname: Public IP address or hostname of the remote location.
Interoperability
In order to set up a successful VPN to a third-party gateway, the following information needs to match:
- Route-Based or Policy-Based VPN
- Remote and local subnets (when using a Policy-Based VPN)
- Key Exchange Version
- Encryption, Hash, and Lifetimes
- Diffie-Hellman (DH) Groups and Perfect Forward Secrecy (PFS)
- Pre-shared Key
- Remote and local IP address
- Remote or local authentication ID (when behind NAT)
Route-Based or Policy-Based VPN
- UniFi gateways use Route-Based VPNs by default. Switching to a Policy-Based VPN is possible.
- Route-Based VPNs use Virtual Tunnel Interfaces (VTIs) and automatically created static routes or exchange routes via OSPF.
  - When using OSPF, it is required to configure a Tunnel IP address to set up a neighbor connection.
- Policy-Based VPNs exchange the remote and local subnets. These need to match exactly between the two gateways.
- It is not possible to use a Route-Based VPN on one gateway and a Policy-Based on the other. The VPN type needs to match.
Note: If the third-party gateway doesn't provide an option to select a Route-Based or Policy-Based VPN, then it likely only supports Policy-Based.
Remote and Local Subnets
- When using Policy-Based VPNs, UniFi gateways automatically share all local networks over the Site-to-Site VPN.
  - It is possible to only use certain local networks.
- When using Policy-Based VPNs, ensure that the third-party gateway includes all the local networks used on the UniFi gateway.
- Route-Based VPNs use static routes or OSPF, and access is controlled with firewall rules.
  - When using OSPF, set the Remote Networks to None.
Note: It is a requirement for Policy-Based VPNs that the remote and local subnets match. If this is not the case, then the VPN may only partially establish or disconnect.
Key Exchange Version
- Both IKEv1 and IKEv2 are supported on UniFi gateways.
- IKEv2 on UniFi gateways use optimizations that some third-party gateways do not support.
- When using Policy-Based VPNs, if the VPN does not establish or disconnects when using IKEv2, then try separating the IKEv2 networks or switch to IKEv1.
Encryption, Hash, and Lifetimes
- Phase 1 (IKE) and Phase 2 (ESP) can use different encryption and hashing values. AES-128 and SHA1 is used by default for both phases but this can be changed.
- The IKE lifetime is set to 28800 seconds and the ESP lifetime is set to 3600 seconds by default but can be changed.
Diffie-Hellman (DH) Groups and Perfect Forward Secrecy (PFS)
- DH groups are referred to by the number, i.e. DH Group 14.
- Third-party gateways may use MODP notation instead , i.e. 2048-bit which equates to group 14.
- UniFi gateways support PFS and can use different DH groups for Phase 1 (IKE) and Phase 2 (ESP).
- PFS may not be supported on third-party gateways or the implementation is not compatible.
- If the VPN does not establish or disconnects when using PFS, then we recommend disabling this feature.
VPN Mode
- UniFi gateways only support Main Mode.
- Site-to-Site VPNs cannot be established if the third-party gateway is using Aggressive Mode.
Frequently Asked Questions
Route-Based VPNs are categorized by the usage of Virtual Tunnel Interfaces (VTIs). When using a Route-Based VPN, the Security Association (SA) will be set to 0.0.0.0/0 and routes for the remote subnet are used with the VTI as the next-hop.
Policy-Based VPNs are categorized by exchanging networks. When using a Policy-Based VPN, the Security Association (SA) will be set to the remote and local subnet (i.e. 192.168.1.0/24 and 172.16.1.0/24) and these need to match exactly between the gateways.
If the third-party gateway doesn't provide an option to select a Route-Based or Policy-Based VPN, then it likely only supports Policy-Based. Refer to the documentation of the third-party gateway for more information.
See the interoperability section for more information on Policy-Based and Route-Based VPNs.
Check if one of the gateways is assigned a private IP address and is behind another router.
If both gateways are using public IP addresses, then verify if the configuration matches.
See the interoperability section for more information.
To test connectivity over the VPN, try pinging between two clients instead of to or from the UniFi gateway itself.
If the ping is not working, then verify if there are any Firewall or Traffic Rules configured that may prevent connectivity.
When using Windows clients for testing, also ensure that the ping traffic is allowed through the Windows firewall.
We recommend to use IPsec Site-to-Site VPNs on a UniFi Gateway that has access to a public IP address. Any performance or port forwarding issues on the upstream router can cause the VPN to disconnect.
If this is not an option, then configure the authentication IDs. For example, an IPsec Site-to-Site VPN is set up between the below UniFi Gateways:
- UniFi Gateway Site A - WAN IP 192.168.5.1 (behind NAT)
- ISP modem/router Site A - WAN IP 203.0.113.1 (public IP)
- UniFi Gateway Site B - WAN IP IP 198.51.100.1 (public IP)
The VPN is set up between the public IP addresses 203.0.113.1 > 198.51.100.1.
When Site B receives the IPsec VPN peer request from Site A, it will contain both the 192.168.5.1 and 203.0.113.1 IP addresses. However, Site B is only configured to peer with 203.0.113.1 causing a mismatch. To resolve this, configure 203.0.113.1 as the Local Authentication ID on Site A.
The reverse is also possible. The 192.168.5.1 IP address can be configured as the Remote Authentication ID on site B.
Besides IP addresses, authentication IDs also support hostnames, email addresses and distinguished names.
No, these are automatically created.
