---
id: collect-260926-mikrotik/mikrotik/questions-577894-routing-from-mikrotik-two-ip-addresses-to-same-gateway-a2a583ef
title: "questions-577894-routing-from-mikrotik-two-ip-addresses-to-same-gateway-a2a583ef"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-577894-routing-from-mikrotik-two-ip-addresses-to-same-gateway-a2a583ef.md
source_anchor: ""
source_lines: [1, 22]
sha256: 38357de68dc4390d0c8c79e47f0afd1d2a91882361278727dc3a8272bc30c6cb
---

# questions-577894-routing-from-mikrotik-two-ip-addresses-to-same-gateway-a2a583ef

I have 2 static addreses (a.b.c.72/24 and a.b.c.128/24) on gateway1 and gateway2 interfaces. ISP have one gateway  a.b.c.1 . I can use a.b.c.72 only on gateway1, a.b.c.128 only on gateway2.
How to get a.b.c.128 working? I want route some local machines via second interface. I know some ways to route via different gateways in mikrotik. In linux i can specify dev parameter to route.
I added routing tables and rules, but it still routed via gateway2.
 /ip dhcp-client print
Flags: X - disabled, I - invalid 
 #   INTERFACE           USE ADD-DEFAULT-ROUTE STATUS        ADDRESS           
 0   gateway1            yes no                bound         X.Y.164.72/24  
 1   gateway2            yes no                bound         X.Y.164.128/24 
/ip route
add distance=51 gateway=X.Y.164.1 pref-src=X.Y.164.128 routing-mark=gate2
add distance=1 dst-address=X.Y.164.0/24 gateway=gateway2 pref-src=X.Y.164.128 routing-mark=gate2
add distance=52 gateway=X.Y.164.1 pref-src=X.Y.164.72 routing-mark=gate1
add distance=1 dst-address=X.Y.164.0/24 gateway=gateway1 pref-src=X.Y.164.72 routing-mark=gate1
add distance=52 gateway=X.Y.164.1
add distance=10 dst-address=10.0.0.0/8 gateway=center
On selected line i need to have iproute equivalent to ip route add default via 109.60.164.1 dev gateway2 table gate2 but it routes via gateway1 interface
NAT
/ip firewall nat
add action=masquerade chain=srcnat out-interface=gateway1
add action=masquerade chain=srcnat out-interface=gateway2
Mangle is clean now. It had marks connections from and routes for selected clients. It is not a problem and workswith different gateways.
/ip address,/ip firewall filter,/ip firewall nat,/ip route.ip route add dst-address=0.0.0.0/0 gateway=1.1.1.1- you will need one of these for each interface, obviously.
