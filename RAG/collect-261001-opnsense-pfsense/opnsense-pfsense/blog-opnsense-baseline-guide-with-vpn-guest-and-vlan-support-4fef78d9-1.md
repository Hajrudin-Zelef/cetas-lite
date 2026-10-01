---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9-1
title: "blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "research"]
source: docs/RAG/collect-261001-opnsense-pfsense/blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9.md
source_anchor: ""
source_lines: [1, 104]
sha256: 677fc298c3446819b287317934d6a684154a453e464e8658fbeb0714daf76583
---

# blog-opnsense-baseline-guide-with-vpn-guest-and-vlan-support-4fef78d9

OPNsense Baseline Guide with Mullvad VPN Multi-WAN, Guest, and VLAN Support
This beginner-friendly, step-by-step guide walks you through the initial configuration of your OPNsense firewall. The title of this guide is an homage to the pfSense baseline guide with VPN, Guest, and VLAN support that some of you guys might know, and this is an OPNsense migration of it. I found that guide two years ago and immediately fell in love with the network setup. After researching for weeks, I decided to use OPNsense instead of pfSense. I bit the bullet and bought the Deciso DEC630 appliance. Albeit expensive and possibly overkill for my needs, I’m happy to support the open-source mission of Deciso, the maintainers of OPNsense. The only thing I regret about the purchase is that I now can’t afford the sexier-looking successor model, the DEC690.
To configure OPNsense, I followed the instructions of the pfSense guide, taking notes on the differences. Some options moved to different menus or changed. As my notes grew, I decided to publish them as a guide on my website.
My goal was to create a comprehensive guide that’s easy to follow. But I tried to strike a different balance regarding the brevity of the instructions compared to the pfSense guide. It’s a matter of personal taste, but I find the instructions in that guide too verbose. I intentionally omit most of the repetitive “click save and apply” instructions and only list configuration changes deviating from defaults, making exceptions for important settings. I consider the OPNsense defaults stable enough for this approach in the hope of keeping the effort required to maintain this guide to a minimum.
I’m a homelab hobbyist, so be warned that this guide likely contains errors. Please, verify the steps yourself and do your research. I hope this guide is as helpful and inspiring to you as the pfSense guide was to me. Your feedback is always welcome and very much appreciated.
Overview
WAN
- DHCP WAN from a single Internet Service Provider (ISP)
- Mullvad VPN multi-WAN with gateway groups
LAN
We segregate the local network into several areas with different requirements.
Management Network (VLAN 10)
The Management network connects native management interfaces like WiFi access points and IPMI interfaces.
VPN Network (VLAN 20)
The primary LAN network uses the WireGuard VPN tunnels for outbound connections, maximizing privacy and security. If the VPN tunnels fail, outbound connections won’t be possible. Exceptions to selectively route traffic through the ISP WAN gateway are possible.
“Clear” Network (VLAN 30)
General-purpose web access network that doesn’t use VPN tunnels. All outgoing connections leave through the ISP WAN gateway. It serves as a backup network in case the VPN tunnels fail.
Guest Network (VLAN 40)
The network that visitors use. It allows unrestricted internet access. Local networks aren’t accessible.
LAN Network
“Native” VLAN, used to debug and test new configurations.
DNS Services
We’ll configure a DNS resolver (Unbound), as well as a DNS forwarder (Dnsmasq) in OPNsense. Management and VPN networks will use the resolver, the Clear network will use the forwarder, and the Guest network will use Cloudflare as an external DNS resolver. We’ll dig into the details later.
Hardware Selection and Installation
The original pfSense guide features a large section of hardware recommendations and installation instructions.
As mentioned earlier, I bought the Deciso DEC630 appliance, which is why I’m not advising on hardware choices. Have a look at the official hardware sizing & setup guidelines for more information. See also Initial Installation & Configuration.
I verified this guide with a clean install of OPNsense version 21.7.5.
Wizard
Navigate to 192.168.1.1 in your browser and login with default credentials:
- Username: root
- Password: opnsense
Click Next to leave the welcome screen and get started with the initial wizard
configuration.
General Information
I prefer using the DNS servers of Quad9 over the ones of my ISP. Only the Clear network will use these anyway, as secured networks use Unbound instead. The Guest network will use Cloudflare DNS servers.
For the domain, I prefer to use a subdomain of a domain name I own, like
corp.example.com. I only use this subdomain internally. I consider the
local.lan pattern a relic of the past. To prevent our local network structure
from leaking to the outside world, we’ll later configure Unbound and Dnsmasq to
treat the domain as private.
| Domain | corp.example.com | 
| Primary DNS Server | 9.9.9.9 | 
| Secondary DNS Server | 149.112.112.112 | 
| Override DNS | unchecked | 
| Enable DNSSEC Support | checked | 
| Harden DNSSEC data | checked | 
If you prefer using your ISP’s DNS servers, leave the Override DNS option checked.
Time Server Information
Choose the NTP servers geographically closest to your location. I live in
Switzerland, which makes the
servers from the ch.pool.ntp.org pool the
natural choice.
| Time server hostname | 0.ch.pool.ntp.org 1.ch.pool.ntp.org 2.ch.pool.ntp.org 3.ch.pool.ntp.org | 
| Timezone | Europe/Zurich | 
Configure Interfaces
By default, the WAN interface obtains an IP address from your ISP via DHCP. DHCP
is also configured for the LAN interface by default and has the IP
192.168.1.1. It works for most people, so we just keep the defaults.
Set Root Password
Choose a strong root password and complete the wizard.
General Settings
Access
Navigate to System → Settings → Administration.
| HTTP Redirect |  | 
|---|---|
| Disable web GUI redirect rule | checked | 
Permitting root user login and password login is a quick and dirty way of enabling SSH access, but I strongly discourage you from doing it. They are disabled for security reasons. I highly recommend using certificate- or key-based authentication. If your device has a serial console port, like the Deciso DEC630, enabling SSH is not required.
| Secure Shell |  |  | 
|---|---|---|
| Secure Shell Server | checked |  | 
| Authentication |  |  | 
|---|---|---|
| Sudo | Ask password | Permit sudo usage for administrators with shell access. | 
Navigate to System → Access → Users and add a new user.
| Username | <choose a username> | 
| Password | <choose a secure password> | 
| Login shell | /bin/csh | 
| Group Memberships | admins | 
| Authorized keys | <valid SSH public key> | 
Configuring the SSH client and generating keys is out of scope for this guide, so I’ll just recommend this DigitalOcean tutorial covering SSH essentials.
Miscellaneous
Navigate to System → Settings → Miscellaneous.
| Power Savings |  | 
|---|---|
| Use PowerD | checked | 
| Power Mode | Hiadaptive | 
Choose Cryptography settings and Thermal Sensors settings compatible with your hardware.
Firewall Settings
Navigate to Firewall → Settings → Advanced.
Although IPv6 is something I want to use, it’s out of scope for this guide, so we uncheck the following.
| Allow IPv6 | unchecked | 
When a rule uses a specific gateway and goes down, a rule gets created, sending traffic to the default gateway. Checking this option skips the creation of this rule.
| Gateway Monitoring |  | 
|---|---|
| Skip rules | checked | 
Depending on your hardware, you might want to tweak the following settings to improve performance.
| Miscellaneous |  |  | 
|---|---|---|
| Firewall Optimization | conservative | Tries to avoid dropping any legitimate idle connections at the expense of increased memory usage and CPU utilization. | 
| Firewall Maximum Table Entries | 2000000 | default is 1'000'000 | 
We disable the auto-generated anti-lockout rule because we’ll define it manually later.
| Disable anti-lockout | checked | 
Checksum Offloading
For some hardware, checksum offloading doesn’t work, particularly some Realtek cards. Rarely, drivers may have problems with checksum offloading and some specific NICs. If your hardware is incompatible with checksum offloading, disable it.
