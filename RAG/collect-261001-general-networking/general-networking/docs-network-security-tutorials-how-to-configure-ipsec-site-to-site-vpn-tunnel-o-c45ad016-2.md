---
id: collect-261001-general-networking/general-networking/docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016-2
title: "docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016.md
source_anchor: ""
source_lines: [184, 358]
sha256: 38e118205a8d91703d208babefdaf87edd2004e5499a57c794fb07eea088c91b
---

# docs-network-security-tutorials-how-to-configure-ipsec-site-to-site-vpn-tunnel-o-c45ad016

| DH key group | 14 (2048 bit) | 2048 bit should be sufficient | 
| Lifetime | 28800 sec | lifetime before renegotiation | 
Table 8. Phase 1 proposal (Algorithms) Phase-1 options for Site-A
Advanced Phase-1 options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Disable Rekey | Unchecked | Renegotiate when connection is about to expire | 
| Disable Reauth | Unchecked | For IKEv2 only re-authenticate peer on rekeying | 
| NAT Traversal | Disabled | For IKEv2 NAT traversal is always enabled | 
| Dead Peer Detection | Unchecked |  | 
Table 9. Advanced Phase-1 options for Site-A
You may easily configure IPSec Phase-1 on Site-A by following the next steps:
- 
Navigate to the VPN → IPSec → Tunnel Settings on Site-A OPNsense web UI.
- 
Click Add button with + at the right bottom of the Phase 1 pane.
- 
Enter the public IP address or hostname of the Remote Gateway, such as 11.11.11.2 .
- 
Enter a Description for your reference, such as Site B .
- 
You may leave other options as default in General information pane. Figure 8. General Information for Phase-1 on Site-A
- 
Enter your Pre-Shared Key string., such as MyS2SIPSecTunnel .
- 
You may leave other options as default in the Phase 1 proposal (Authentication) pane. Figure 9. Phase 1 proposal (Authentication) on Site-A
- 
Select Encryption algorithm, such as 256-bit AES-GCM with128-bit ICV .
- 
Select Hash algorithm, such as SHA512 .
- 
Select DH key group, such as 14 (2048) bits . This option must match the setting chosen on the remote side.Figure 10. Algorithms Phase 1 options on Site-A
- 
Set NAT Traversal to Disable in Advanced Options pane.
- 
Click Save. Figure 11. Advanced Phase 1 options on Site-A
3. Configuring Phase 2 on Site-A
General Information Phase-2 options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Mode | Tunnel IPv4 | Select Tunnel mode | 
| Description | Local LAN Site B | Freely chosen description | 
Table 10. General Information Phase-2 options on Site-A
Local Network Phase-2 options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Local Network | LAN subnet | Route the local LAN subnet | 
Table 11. Local Network Phase-2 options on Site-A
Remote Network Phase-2 options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Type | Network | Route a remote network | 
| Address | 10.10.11.0/24 | The remote LAN subnet | 
Table 12. Remote Network Phase-2 options on Site-A
Phase 2 proposal (SA/Key Exchange) options on Site-A are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Protocol | ESP | Choose ESP for encryption | 
| Encryption algorithms | AES256GCM16 | For the sample we use AES256GCM16 | 
| Hash algortihms | SHA512 | Choose a strong hash like SHA512 | 
| PFS Key group | 14 (2048 bit) | Not required but enhanced security | 
| Lifetime | 3600 sec |  | 
Table 13. Phase 2 proposal (SA/Key Exchange) on Site-A
You may easily configure IPSec Phase-1 on Site-A by following the next steps:
- 
Navigate to the VPN → IPSec → Tunnel Settings on Site-A OPNsense web UI.
- 
Click add phase 2 entry button with + at the Commands column of the recently added phase 1 entry.Figure 12. Adding Phase-2 Entry
- 
Add a Description, such as Local LAN Site B .
- 
Set Address option for Remote Network, such as 10.10.11.0/24 .Figure 13. General Information for Phase-2 on Site-A
- 
Select Encryption algorithms, such as AES256GCM16 .
- 
Select Hash algorithms, such as SHA512 .
- 
Select PFS key group, such as 14 (2048) bits .
- 
Set Lifetime, such as 3600 .
- 
You may leave other options as default. Figure 14. Algorithms for Phase-2 on Site-A
- 
Click Save.
- 
Click the checkbox at the beginning of the Phase 1 pane to view the Phase 2 settings.
4. Enabling IPsec on Site-A
You may quickly enable IPsec service on SIte-A by following the next steps:
- 
Navigate to the VPN → IPSec → Tunnel Settings on Site-A OPNsense web UI.
- 
Check Enable IPsec option at the bottom of the page.
- 
Click Apply Changes button at the top right corner of the page to activate the IPsec tunnel settings. Figure 15. Enabling IPsec on Site-A
5. Configuring Phase 1 on Site-B
General Phase-1 options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Connection method | default | default is "Start on traffic" | 
| Key Exchange version | V2 |  | 
| Internet Protocol | IPv4 |  | 
| Interface | WAN | choose the interface connected to the internet | 
| Remote gateway | 11.11.11.1 | the public IP address of your remote OPNsense | 
| Description | Site A | freely chosen description | 
Table 14. General Phase-1 options on Site-B
Authentication Phase-1 options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Authentication method | Mutual PSK | Using a Pre-shared Key | 
| My identifier | My IP address | Simple identification for fixed ip | 
| Peer identifier | Peer IP address | Simple identification for fixed ip | 
| Pre-Shared Key | MyS2SIPSecTunnel | Random key. You should create your own. | 
Table 15. Authentication Phase-1 options on Site-B
Phase 1 proposal (Algorithms) options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Encryption algorithm | 256-bit AES-GCM with128-bit ICV | For our sample we will Use AES/256 bits | 
| Hash algoritm | SHA512 | Use a strong hash like SHA512 | 
| DH key group | 14 (2048 bit) | 2048 bit should be sufficient | 
| Lifetime | 28800 sec | lifetime before renegotiation | 
Table 16. Algorithms Phase-1 options on Site-B
Advanced Phase-1 options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Disable Rekey | Unchecked | Renegotiate when connection is about to expire | 
| Disable Reauth | Unchecked | For IKEv2 only re-authenticate peer on rekeying | 
| NAT Traversal | Disabled | For IKEv2 NAT traversal is always enabled | 
| Dead Peer Detection | Unchecked |  | 
Table 17. Advanced Phase-1 options on Site-B
You may easily configure IPSec Phase-1 on Site-B by following the next steps:
- 
Navigate to the VPN → IPSec → Tunnel Settings on Site-A OPNsense web UI.
- 
Click Add button with + at the right bottom of the Phase 1 pane.
- 
Enter the public IP address or hostname of the Remote Gateway, such as 11.11.11.1 .
- 
Enter a Description for your reference, such as Site A .
- 
You may leave other options as default in General information pane. Figure 16. General information Phase-1 on Site-B
- 
Enter your Pre-Shared Key string., such as MyS2SIPSecTunnel .
- 
You may leave other options as default in the Phase 1 proposal (Authentication) pane.
- 
Select Encryption algorithm, such as 256-bit AES-GCM with128-bit ICV .
- 
Select Hash algorithm, such as SHA512 .
- 
Select DH key group, such as 14 (2048) bits . This option must match the setting chosen on the remote side.
- 
Set NAT Traversal to Disable in Advanced Options pane.
- 
Click Save.
6. Configuring Phase 2 on Site-B
General Information Phase-2 options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Mode | Tunnel IPv4 | Select Tunnel mode | 
| Description | Local LAN Site A | Freely chosen description | 
Table 18. General Information Phase-2 options on Site-B
Local Network Phase-2 options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Local Network | LAN subnet | Route the local LAN subnet | 
Table 19. Local Network Phase-2 options on Site-B
Remote Network Phase-2 options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
| Type | Network | Route a remote network | 
| Address | 10.10.10.0/24 | The remote LAN subnet | 
Table 20. Remote Network Phase-2 options on Site-B
Phase 2 proposal (SA/Key Exchange) options on Site-B are given in the next table.
| Option | Value | Description | 
|---|---|---|
