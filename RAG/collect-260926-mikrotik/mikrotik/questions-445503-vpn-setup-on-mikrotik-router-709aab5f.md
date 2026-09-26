---
id: collect-260926-mikrotik/mikrotik/questions-445503-vpn-setup-on-mikrotik-router-709aab5f
title: "questions-445503-vpn-setup-on-mikrotik-router-709aab5f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-445503-vpn-setup-on-mikrotik-router-709aab5f.md
source_anchor: ""
source_lines: [1, 11]
sha256: eb83103a08a3549fda20a3cb69c322035736ecdc914ed42e179d8983db16d0f2
---

# questions-445503-vpn-setup-on-mikrotik-router-709aab5f

I have successfully setup a pptp vpn server on a Mikrotik routerboard. I am able to connect to the VPN and can successfully access a certain IP range.
This is my setup:
ADSL modem/router - mikrotik router -> private community network (CTWUG)
IP range at home is a subset of the larger CTWUG network range : 172.18.107.224/28 The CTWUG network rage : 172.18.0.0/16
- In my home setup the mikrotik router acts as the gateway and router. IP: 172.18.107.238
- The ADSL router acts as the DHCP server. IP : 172.18.107.237
- There is a route on the mikrotik for access to the CTWUG range: 172.18.0.0/16 -> 172.18.107.252
- And a route for internet access 0.0.0.0/0 -> 172.18.107.237 (ADSL modem)
- For the VPN the local ip issued is 172.18.107.236 and remote ip issued is 172.18.107.235
- On my ADSL modem I have setup port forwarding on port 1723 TCP to 172.18.107.238 (mikrotik)
The problem and question is: I am able to access the complete CTWUG range 172.16.0.0/16 except my home range 172.18.107.224/28. I have torched the pptp(VPN) interface and I can see the packets for my local range being forwarded to the device on my local range. There is just no packets returning over the VPN. Any help please.
