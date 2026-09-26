---
id: collect-260926-mikrotik/mikrotik/questions-493351-mikrotik-not-accepting-vpn-connections-over-wifi-24320e6f
title: "questions-493351-mikrotik-not-accepting-vpn-connections-over-wifi-24320e6f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-493351-mikrotik-not-accepting-vpn-connections-over-wifi-24320e6f.md
source_anchor: ""
source_lines: [1, 10]
sha256: 364c9d93e496a8605a63987df10bc773119d764655d47f56f67a59659d9ddf10
---

# questions-493351-mikrotik-not-accepting-vpn-connections-over-wifi-24320e6f

I'm using MikroTik RouterOS v5.18 x86 on a Dell PC. I have done the following:
- Set Addresses
- Set PPPoE Client & Server
- Set PPTP/L2TP Server
- Make Profiles and Secrets
- SRCNAT MASQUERADE the IPs of PPPoE and PPTP POOL
Now the problem is this:
I have the Setup as shown below. 4 PCs Connected to RouterOS via LAN Switch. 1 PC and Android devices via Tenda W311R+ WIFI Router. I don't want to give anonymous access to internet via WIFI hence I set it to DHCP Client not PPPoE.
The problem is that when I connect to VPN via the PC or the Android via WIFI through Tenda router, VPN does not connect. It connects over the Ethernet (wired) connection to Tenda. I have to work with what I have got, so please do not ask me to buy a WIFI Card for the PC as I know it would make it much easier!
PS This is for a Non-Profit Organization, so you get the idea why I cant buy new stuff!
