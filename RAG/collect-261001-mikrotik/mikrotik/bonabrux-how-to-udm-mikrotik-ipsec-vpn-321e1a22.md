---
id: collect-261001-mikrotik/mikrotik/bonabrux-how-to-udm-mikrotik-ipsec-vpn-321e1a22
title: "bonabrux-how-to-udm-mikrotik-ipsec-vpn-321e1a22"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/bonabrux-how-to-udm-mikrotik-ipsec-vpn-321e1a22.md
source_anchor: ""
source_lines: [1, 126]
sha256: 831070a379eb234bb6bd7e946f6a808b0c6502e34cae08636f382209fe8a8f5f
---

# bonabrux-how-to-udm-mikrotik-ipsec-vpn-321e1a22

How to create an IPsec VPN between Unifi UDM and Mikrotik firewalls
Hardware and software used
- UniFi OS 4.3.6
- Unifi Network 9.3.43
- Mikrotik RouterOS v7.17
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
Settings -> VPN -> Create Site-to-site VPN
| Configuration |  | 
| Name | Mikrotik-01 | 
| VPN Protocol | Manual IPsec | 
| Pre-shared Key | "YOUR SECRET KEY" | 
| UniFi Gateway IP | "WAN IP of UDM" | 
| Remote IP / Hostname | "WAN IP or FQDN of Mikrotik" | 
| VPN Method | Routed Based | 
| Tunnel IP | Unchecked | 
| Remote Network | Static | 
| Add your Remote Subnets | Mikrotik LAN subnet (e.g. 10.0.5.0/24) | 
| Advanced | Manual | 
| IPsec Profile | Customized | 
| Key Echange Version | IKEv2 | 
| Encryption | AES-128 | 
| Hash | SHA1 | 
| IKE DH Group | 14 | 
| ESP DH Group | 14 | 
| Perfect Forward Secrecy (PFS) | Enabled (checked) | 
| Local Authentication ID | if Auto your WAN IP will be used, check with what you configured as expected in Mikrotik IP -> IPsec -> Identities | 
| Remote Authentication ID | if Auto your Remote IP or FQDN will be used, check with what you configured as expected in Mikrotik IP -> IPsec -> Identities | 
| Route Distance | 30 | 
Mikrotik IPsec -> Installed SAs
- Something like this should show up when connection is up
|  | SPI | Src. Address | Dst. Address | Auth. Algorithm | Encr. Algorithm | Encr. Key Size | Current Bytes | 
| EH | 4cbfd50 | 62.x.y.z | 83.x.y.z | sha1 | aes cbc | 128 | 0 | 
| EH | c0a27199 | 83.x.y.z | 62.x.y.z | sha1 | aes cbc | 128 | 3440 | 
- You should be able to ping both ways from each location.
