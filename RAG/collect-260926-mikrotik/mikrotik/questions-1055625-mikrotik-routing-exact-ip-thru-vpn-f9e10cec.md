---
id: collect-260926-mikrotik/mikrotik/questions-1055625-mikrotik-routing-exact-ip-thru-vpn-f9e10cec
title: "questions-1055625-mikrotik-routing-exact-ip-thru-vpn-f9e10cec"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1055625-mikrotik-routing-exact-ip-thru-vpn-f9e10cec.md
source_anchor: ""
source_lines: [1, 10]
sha256: 1cd93a0255853b5c6676e6cf0a027fb12c072fe3af430f59ac6047ab4962de1d
---

# questions-1055625-mikrotik-routing-exact-ip-thru-vpn-f9e10cec

I have the following setup:
- Linux server running OpenVPN Server. It has full access to the Internet and public IP 1.2.3.4. OpenVPN uses tun0 and static IP 10.8.0.1
- Mikrotik hEX lite (i don't think the model really matter). It also has a full access to the Internet, but it's behind some other, uncontrolled router. It's connected to OpenVPN server via myvlan, IP 10.8.0.2.
- OpenVPN tunnel is active, i can ping each other, using 10.8.0.*
I'm trying to route traffic from linux server to 100.1.1.1 through VPN tullel to Mikrotik and then to the internet, like shown on the picture. So I have
route add 100.1.1.1 via 10.8.0.2 dev tun0 on linux server:
$ ip route get 100.1.1.1
100.1.1.1 via 10.8.0.2 dev tun0 src 10.8.0.1 uid 1000
    cache
But I have lack of knowledge, what should I do on Mikrotik side? I thought that default masquerade aka "scrnat everything to Out. Interface List = WAN" would be enough, but I was wrong.
