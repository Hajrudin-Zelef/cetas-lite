---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/en-wiki-opnsense-wireguard-vpn-site-to-site-configuration-f4c3c11e
title: "en-wiki-opnsense-wireguard-vpn-site-to-site-configuration-f4c3c11e"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/en-wiki-opnsense-wireguard-vpn-site-to-site-configuration-f4c3c11e.md
source_anchor: ""
source_lines: [1, 73]
sha256: 1a841ff50a07dc111ffa21870cba41ee36f3a9a98bf2f63cecf3d6f87e870d45
---

# en-wiki-opnsense-wireguard-vpn-site-to-site-configuration-f4c3c11e

OPNsense WireGuard VPN Site-to-Site configuration
| Please note that this article / this category refers either on older software / hardware components or is no longer maintained for other reasons. This page is no longer updated and is purely for reference purposes still here in the archive available. | 
|---|
The Open Source firewall OPNsense supports several technologies for setting up VPN (Virtual Private Network) connections. In addition to IPsec and OpenVPN, OPNsense version 19.7 offers the possibility to set up a VPN with WireGuard. In this article we show the configuration of the WireGuard VPN service to connect two OPNsense firewalls to a Site-to-Site VPN.
Note: This manual was created with an older version of OPNsense, it may be that some menus and settings are now done differently.
Prepare OPNsense for Wireguard Site-to-Site VPN
To set up WireGuard VPN on both firewalls, follow the steps described in the article OPNsense WireGuard VPN for Road Warrior configuration in the section Prepare OPNsense for Wireguard VPN.
'Additional firewall rules for both firewalls
- 
			
			On firewall Site A, configure a further rule on WireGuard analogous to the rule already created, which allows the LAN network of the remote site to access it.
- 
			
			On Firewall Site B, configure another rule that allows access to the LAN network of the remote site.
Configuration sequence of the two firewalls
The following section shows the configuration process for both firewalls of this WireGuard VPN Site-to-Site connection.
Configuration of the local endpoint on the firewall A
- Create Local Configuration:
  - Click on VPN ‣ WireGuard.
  - Click on the tab Local to configure the local WireGuard instance.
  - Click on the + symbol and fill in the following fields:
    - Name: ThomasKrennWGSitetoSiteA
    - Listen Port: 51820 (is alternatively randomly created and then also starts at 51820)
    - Tunnel Address: 10.11.0.1/24
  - Then click on Save.
    - The endpoint was created and Private Key and Public Key was created.
    - You can call up the menu again by clicking on the pen symbol, now the Private Key and the Public Key are displayed here. You will need the public key later for further configuration. Click on Cancel.
Configuration of the local endpoint on the firewall B
- Create Local Configuration:
  - Click on VPN ‣ WireGuard.
  - Click on the tab Local to configure the local WireGuard instance.
  - Click on the + symbol and fill in the following fields:
    - Name: ThomasKrennWGSitetoSiteB
    - Listen Port: 51820 (is alternatively randomly created and then also starts at 51820)
    - Tunnel Address: 10.11.0.2/24
  - Then click on Save.
    - The endpoint was created and Private Key and Public Key was created.
    - You can call up the menu again by clicking on the pen symbol, now the Private Key and the Public Key are displayed here. You will need the public key later for further configuration. Click on Cancel.
Configuration of the endpoint on firewall A
- Create endpoint for firewall B:
  - Click on VPN ‣ WireGuard
  - Then click on the tab Endpoints: Here you configure the remote WireGuard instance (firewall B).
  - Click on the + icon and fill in the following fields:
    - Name: TKLESnetworkplus
    - Public Key: Copy the public key of the local configuration of firewall B and paste it here.
    - Shared Secret: (Optional) Specify Shared Secret.
    - Allowed IPs: 10.11.0.2/32 and 192.168.2.0/24 (LAN address range of firewall B)
    - Endpoint Address: 10.1.102.252 (Publicly accessible IP address)
    - Endpoint Port: 51820
- Add to Local Configuration:
  - Switch to the Local tab.
  - Click on the button to edit the entry.
    - In the dropdown menu Peers you can now select the configured endpoint (TKLESnetworkplus).
    - Click on Save.
    - Click on Save again.
Configuration of the endpoint on firewall B
- Create endpoint for firewall A:
  - Click on VPN ‣ WireGuard
  - Then click on the tab Endpoints: Here you configure the remote WireGuard instance (firewall A).
  - Click on the + icon and fill in the following fields:
    - Name: TKX11SSHLN4F
    - Public Key: Copy the public key of the local configuration of firewall A and paste it here.
    - Shared Secret: (Optional) Specify Shared Secret.
    - Allowed IPs: 10.11.0.1/32 and 192.168.1.0/24 (LAN address range of firewall A)
    - Endpoint Address: 10.1.102.251 (Publicly accessible IP address)
    - Endpoint Port: 51820
- Add to Local Configuration:
  - Switch to the Local tab.
  - Click on the button to edit the entry.
    - In the dropdown menu Peers you can now select the configured endpoint (TKX11SSHLN4F).
    - Click on Save.
    - Click on Save again.
|  | Author: Thomas Niedermeier Thomas Niedermeier working in the product management team at Thomas-Krenn, completed his bachelor's degree in business informatics at the Deggendorf University of Applied Sciences. Since 2013 Thomas is employed at Thomas-Krenn and takes care of OPNsense firewalls, the Thomas-Krenn-Wiki and firmware security updates.  |
