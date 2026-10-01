---
id: collect-261001-mikrotik/mikrotik/gobman-how-to-udm-mikrotik-ipsec-vpn-b2481cce
title: "gobman-how-to-udm-mikrotik-ipsec-vpn-b2481cce"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/gobman-how-to-udm-mikrotik-ipsec-vpn-b2481cce.md
source_anchor: ""
source_lines: [1, 139]
sha256: 79c251411a9cfe0ef2d60132c50f12d39e895e5add2b3e2b4bed1fb2632094ac
---

# gobman-how-to-udm-mikrotik-ipsec-vpn-b2481cce

How to create an IPsec VPN between Unifi UDM and Mikrotik firewalls
Hardware and software used
- UniFi OS 3.1.16 - 3.2.7
- Unifi Network 7.2.92 - 8.0.26
- Mikrotik RouterOS v7.4
You will find these placeholders in the configuration below and you need to replace the values
- "WAN IP of UDM" - The public IP of location with UDM
- "WAN IP of Mikrotik" - The public IP of location with Mikrotik
- "YOUR SECRET KEY" - a very long password atleast 64 characters long, best to use some password generator
mapping for Mikrotik to UDM (Google: Diffie-Hellman Groups)
- modp1024 = DH Group 2
- modp2048 = DH Group 14
Mikrotik configuration in WebFig interface
Select: IP -> IPsec -> Profiles
| New Profile |  | 
| Name | UDM-profile | 
| Hash Algorithms | sha1 | 
| PRF Algorithms | auto | 
| Encryption Algorithm | aes-128 | 
| DH Group | modp1024, modp2048 | 
| Proposal Check | obey | 
| Lifetime | 1d 00:00:00 | 
| Lifebytes | empty | 
| Nat Traversal | uncheck | 
| DPD Interval | 60 | 
| DPD Maximum Failures | 5 | 
Select: IP -> IPsec -> Peers
| Add New |  | 
| Name | UDM-01 | 
| Address | "WAN IP of UDM" | 
| Port | empty | 
| Local Address | empty (or LAN IP of Mikrotik if behind NAT) | 
| Profile | UDM-profile | 
| Exchange Mode | IKE2 | 
| Passive | uncheked | 
| Send INITIAL_CONTACT | checked | 
Select: IP -> IPsec -> Identities
| Add New |  | 
| Peer | UDM-01 | 
| Auth. Method | pre shared key | 
| Secret | "YOUR SECRET KEY" | 
| Policy Template Group | default | 
| Notrack Chain | empty | 
| My ID Type | auto (if router behind NAT use 'address') | 
| My ID | empty (if router behind NAT use "WAN IP of Mikrotik") | 
| Remote ID Type | auto | 
| Match By | remote id | 
| Mode Configuration | empty | 
| Generate Policy | no | 
Select: IP -> IPsec -> Proposals
| Add New |  | 
| Name | UDM proposal | 
| Auth. Algorithms | sha1 | 
| Encr. Algorithms | aes-128 cbc | 
| Lifetime | 00:30:00 | 
| PFS Group | modp2048 | 
Select: IP -> IPsec -> Policies
| Add New IPsec Policy |  | 
| General Tab |  | 
| Peer | UDM-01 | 
| Tunnel | checked | 
| Src. Address | Mikrotik internal LAN network address (the whole network e.g. 10.0.5.0/24) | 
| Src. Port | empty | 
| Dst. Address | UDM internal LAN network address (the whole network e.g. 192.168.5.0/24) | 
| Dst. Port | empty | 
| Template | unchecked | 
| Action tab |  | 
| Action | encrypt | 
| Level | unique | 
| IPsec Protocols | esp | 
| Proposal | UDM proposal | 
| Status tab | (this will be populated once the connection is established, fields are read only) | 
| PH2 Count | 1 | 
| PH2 State | established | 
| SA Src. Address | WAN IP of Mikrotik location (or LAN IP of router if behind NAT) | 
| SA Dst. Address | WAN IP of UDM location | 
Select: IP -> Firewall -> Filter Rules
| New Firewall Rule |  | 
| Chain | input | 
| Protocol | 50 (ipsec-esp) | 
| In. Interface | unchecked, ether1 (could be different for you, this is the port on the router where the internet connection comes in) | 
| Action | accept | 
| Comment | allow L2TP VPN (ipsec-esp) | 
| New Firewall Rule |  | 
| Chain | input | 
| Protocol | 17 (udp) | 
| Dst. Port | 500, 1701, 4500 | 
| In. Interface | unchecked, ether1 (could be different for you, this is the port on the router where the internet connection comes in) | 
| Action | accept | 
| Comment | allow L2TP VPN (500,4500,1701/udp) | 
- Move the rule to the top of the firewall Filter rules after the "defconf: accept established,related,untracked" rule.
Select: IP -> Firewall -> NAT
| New NAT Rule |  | 
| Chain | srcnat | 
| Src. Address | LAN network of Mikrotik (e.g. 10.0.5.0/24) | 
| Dst. Address | LAN network of UDM (e.g. 192.168.5.0/24) | 
| Action | accept | 
- Move the rule to the top of the firewall NAT rules.
Settings -> Teleport & VPN -> VPN Server
| Configuration |  | 
| VPN Server | Enabled (checked) | 
| VPN Protocol | L2TP | 
| Pre-shared Key | "YOUR SECRET KEY for UDM" (not the same as for Mikrotik) | 
| UniFi Gateway IP | "WAN IP of UDM" | 
| If you want to also connect with VPN client to your UDM add a user for (Windows VPN clients enable MSCHAPv2 on network adapter). |  | 
| User Authentication | Create a new user, enter username and password for user (make it complex) | 
| Advanced Configuration | Manual | 
| Network name | VPN | 
| User Access List (RADIUS Profiler) | Default | 
| Gateway/Subnet | (e.g. 192.168.6.1/24) (Can't be the same as local IP range of UDM) | 
| Name Server | unchecked | 
| WINS Server | unchecked | 
| Site-to-Site VPN | Enabled (checked) | 
| Require Strong Authentication | Enabled (checked) - this is where you need MSCHAPv2 | 
| Allow weak ciphers | Enabled (checked) | 
Settings -> Teleport & VPN -> Create Site-to-site VPN
| Configuration |  | 
| Name | Mikrotik-01 | 
| VPN Protocol | Manual IPsec | 
| Pre-shared Key | "YOUR SECRET KEY" | 
| UniFi Gateway IP | "WAN IP of UDM" | 
| Shared Remote Subnets | Mikrotik LAN subnet (e.g. 10.0.5.0/24) | 
| Remote IP | "WAN IP of Mikrotik" | 
| Advanced | Manual | 
| IPsec Profile | Customized | 
| Key Echange Version | IKEv2 | 
| Encryption | AES-128 | 
| Hash | SHA1 | 
| IKE DH Group | 14 | 
| ESP DH Group | 14 | 
| Perfect Forward Secrecy (PFS) | Enabled (checked) | 
| Route-Based VPN | Enabled (checked) | 
| Route Distance | 30 | 
Mikrotik IPsec -> Installed SAs
- Something like this should show up when connection is up
|  | SPI | Src. Address | Dst. Address | Auth. Algorithm | Encr. Algorithm | Encr. Key Size | Current Bytes | 
| EH | 4cbfd50 | 62.x.y.z | 83.x.y.z | sha1 | aes cbc | 128 | 0 | 
| EH | c0a27199 | 83.x.y.z | 62.x.y.z | sha1 | aes cbc | 128 | 3440 | 
- You should be able to ping both ways from each location.
