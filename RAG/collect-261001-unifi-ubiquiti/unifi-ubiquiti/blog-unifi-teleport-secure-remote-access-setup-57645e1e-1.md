---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-teleport-secure-remote-access-setup-57645e1e-1
title: "blog-unifi-teleport-secure-remote-access-setup-57645e1e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AMD", "Apple", "Intel", "SpaceX"]
dates: ["2022-04", "2023-05", "2023-12", "2024-10", "2025-12", "2026-03", "2026-04", "2026-09-24"]
keywords: ["amd", "intel"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-teleport-secure-remote-access-setup-57645e1e.md
source_anchor: ""
source_lines: [1, 65]
sha256: 026a89672e84729f9331fe040d65a664a094ea2f4fc7e1fef3c327bf2111ccf7
---

# blog-unifi-teleport-secure-remote-access-setup-57645e1e

UniFi Teleport VPN: how to set it up (Windows, Mac and mobile)
UniFi Teleport is a zero-configuration VPN built into UniFi gateways. You enable it in UniFi Network, send an invitation, and the other person connects with the WiFiman app. It runs on WireGuard, and it works without a public IP or port forwarding.
WiFiman runs on Windows, macOS, Linux, iOS and Android. This guide covers what Teleport needs, how to set it up, the client apps, its limits, and how it compares with Site Manager and the regular VPN servers.
What UniFi Teleport is
Ubiquiti's Teleport VPN article describes it like this:
- It's a remote access VPN into your UniFi network, encrypted with WireGuard.
- It's zero-configuration. You add users by sending an invitation.
- The recipient connects through the WiFiman app on a phone or a computer.
- It works when both the gateway and the client sit behind NAT.
- It can run next to the other UniFi VPN servers.
- Connected clients use the gateway's DNS servers.
Teleport reached an official UniFi Network release with version 7.1.61 in April 2022. Ubiquiti groups it with the VPN servers on the gateway, next to OpenVPN, WireGuard and L2TP (Introduction to VPNs).
What you need
- A UniFi gateway. The help article says Teleport is for "users with a Next-Gen gateway or UniFi Cloud Gateway running UniFi OS". Next-Gen gateway is Ubiquiti's older name for the standalone UniFi Gateways. Dream Machines and the other Cloud Gateways are included.
- Teleport enabled in UniFi Network.
- The WiFiman app on each device that will connect.
- In some cases, IPv6 on the gateway's internet connection. Ubiquiti notes that Teleport "requires an IPv6 connection on your gateway's WAN in some circumstances". You don't need IPv6 on your local networks.
Teleport doesn't need a public IP. Ubiquiti's remote access guide says a public IP is necessary for port forwarding and "most VPNs (excluding Teleport)".
A standalone gateway that's managed by a separate console, such as a Cloud Key, is a special case. UniFi Network 7.4.156 (May 2023) added Teleport for the UXG-Pro on the Default site only, with other sites "in a future release". We haven't found a later release note that adds the other sites.
How to set up UniFi Teleport
- In UniFi Network, open the VPN section of the settings and enable Teleport. Ubiquiti's help articles word this path slightly differently, and the labels change between versions.
- Generate an invitation and send the link to the person who needs access.
- The recipient opens the link. If WiFiman is installed, the invitation adds the Teleport VPN to the app. If it isn't, the link asks them to install WiFiman first. After installing, they open the link again.
- The recipient starts the connection in WiFiman. In the mobile app, Teleport has its own tab.
Two rules apply to every invitation. It expires after 24 hours, and only one device can use it at a time. Send a separate invitation for each laptop and phone.
Since UniFi Network 10.2.105 (March 2026), every admin on the console can start Teleport connections from WiFiman.
Revoking access
A pending invitation can be revoked in Invitation History. Once someone has accepted, you remove their access from the Teleport client in the Client Devices list.
Our own Fernando walks through the setup, including firewall rules, in the video at the top of this page. It was recorded on an earlier UniFi Network version, so some labels may differ from what you see.
UniFi Teleport on Windows, Mac and Linux
Teleport on a computer runs through WiFiman Desktop. You download it from Ubiquiti's WiFiman Desktop page.
| Platform | Minimum version | Builds | 
|---|---|---|
| Windows | Windows 10 | Intel/AMD and ARM | 
| macOS | macOS 13 | Intel and Apple Silicon | 
| Linux | Ubuntu 22 | .deb package | 
These minimums come from the WiFiman Desktop 1.3.0 release notes (September 24, 2026). Ubiquiti announced WiFiman Desktop for Windows, with Teleport, in version 0.3.1 (December 2023). ARM builds for Windows followed in version 1.1.0 (October 2024).
If Teleport drops or hangs on Windows, update WiFiman Desktop first. Version 1.3.0 fixed problems with rapid Teleport connects and disconnects on Windows. It also fixed Teleport getting stuck after a failed attempt or an app restart, and unexpected disconnects. Version 1.2.8 (December 2025) fixed the Windows WireGuard network interface getting stuck after a Teleport disconnect.
On phones, Teleport lives in the WiFiman app for iOS and Android. WiFiman for Android 2.11.1 added a Teleport tile to the Quick Settings menu.
Two UniFi devices can also connect over Teleport. UniFi Talk Touch phones scan a Teleport invitation QR code (UniFi Talk with Teleport). The UniFi mobile app release notes from April 2026 onward add Teleport features for the UniFi Travel Router.
Limiting what Teleport users can reach
On gateways that use the zone-based firewall, the default policy from the VPN zone to the Internal zone is Allow All (Zone-Based Firewalls in UniFi). So a VPN user can reach your internal networks until you add a policy.
To restrict a Teleport user, you add policies from the VPN zone to the networks they shouldn't reach. One caution: that article lists Identity One-Click VPN, WireGuard, L2TP and OpenVPN users in the VPN zone, and doesn't name Teleport. Check which zone your Teleport clients fall under before you rely on a policy.
Test each policy from a connected Teleport client. Try a ping, a web page and a file share on the network you blocked.
Limits and what Ubiquiti doesn't document
- Invitations last 24 hours and cover one device at a time.
- Teleport is a remote user VPN. For site-to-site links, UniFi uses Site Magic, IPsec or OpenVPN.
- Some setups need IPv6 on the gateway's internet connection.
- Ubiquiti's Teleport article doesn't state which subnet Teleport uses or whether you can change it. It also doesn't state which port it uses, whether it supports split tunneling (sending only the site's traffic through the VPN), or a maximum number of users.
If you connect from behind CGNAT (carrier-grade NAT, where your provider shares one public IP address across many customers) or over Starlink, our guide to UniFi remote work with Starlink and Teleport covers that setup.
Teleport vs Site Manager vs WireGuard and OpenVPN
UniFi has several ways to reach a site remotely. They solve different problems.
|  | Teleport | WireGuard server | OpenVPN server | Site Manager | 
|---|---|---|---|---|
| What you get | Access to the site's network | Access to the site's network | Access to the site's network | Management of the console | 
| How users are added | Invitation link | Config file or QR code | Config file | UI account with admin rights | 
| Client app | WiFiman | Any WireGuard client | Any OpenVPN client | Browser or UniFi app | 
| Public IP needed | No | Recommended; forward UDP 51820 if behind NAT | Recommended; forward port 1194 if behind NAT | No; outbound TCP 443 and 8883 | 
Sources: WireGuard VPN Server, OpenVPN Server and Enabling UniFi Remote Management. The ports are the defaults.
Ubiquiti recommends Teleport for phones. For laptops and desktops, it recommends Teleport or WireGuard over OpenVPN. Our UniFi WireGuard guide and OpenVPN guide cover those two servers.
Site Manager is a separate thing. Remote Management lets you run the console from unifi.ui.com, but it doesn't put your laptop on the site's network. To reach a file server or a printer at the site, Ubiquiti points you to a VPN or port forwarding. Our Site Manager guide explains how that side works.
Frequently asked questions
Does UniFi Teleport work on Windows?
Yes. Install WiFiman Desktop on Windows 10 or newer, open the invitation link, and connect in WiFiman. ARM laptops are supported too.
Does Teleport work behind CGNAT?
