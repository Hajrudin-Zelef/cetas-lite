---
id: collect-260926-mikrotik/mikrotik/questions-10360-mikrotik-joining-two-subnets-b9279e02
title: "questions-10360-mikrotik-joining-two-subnets-b9279e02"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-10360-mikrotik-joining-two-subnets-b9279e02.md
source_anchor: ""
source_lines: [1, 24]
sha256: f33a47a8cc9d7432fe354ea7933cd2869d011ff4a0616790981af42165d2753d
---

# questions-10360-mikrotik-joining-two-subnets-b9279e02

I have two subnets which are geographically separated, connected by a wireless point to point link (5Mbps). Both networks are connected to internet through their own separate ILL WAN links which are configured on MikroTik routers. I want to connect the subnets through the routers. The scenario is :
Site 1:
192.168.10.0/24
Ether1 xxx.xxx.xxx.xxx/xx ---- connected to Internet
Ether2 192.168.10.1/24 ---- connected to local LAN
Site 2
10.10.10.0/24
Ether1 xxx.xxx.xxx.xxx/xx ---- connected to Internet
Ether2 10.10.10.1/24 ---- connected to local LAN
I want to connect the two LANs together so that there is a seamless traffic between the two.
NO FIREWALL rules on either routers. Only source NAT masquerade on WAN links on both routers. NO destination NAT
I have tried the below, which, obviously does not work.
Site 1:
192.168.10.0/24
Ether1 xxx.xxx.xxx.xxx/xx ---- connected to Internet
Ether2 192.168.10.1/24 ---- connected to local LAN
Ether3 10.10.10.10/24 ---- P2P between Site 1 and 2
Site 2
10.10.10.0/24
Ether1 xxx.xxx.xxx.xxx/xx ---- connected to Internet
Ether2 10.10.10.1/24 ---- connected to local LAN
Ether3 192.168.10.10/24 ---- P2P between Site 1 and 2
What am I missing? (Beside, looks like everything?)
Thanks!
