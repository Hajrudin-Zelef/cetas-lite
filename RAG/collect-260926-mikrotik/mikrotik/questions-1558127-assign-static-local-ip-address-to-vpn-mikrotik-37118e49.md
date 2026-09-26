---
id: collect-260926-mikrotik/mikrotik/questions-1558127-assign-static-local-ip-address-to-vpn-mikrotik-37118e49
title: "questions-1558127-assign-static-local-ip-address-to-vpn-mikrotik-37118e49"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1558127-assign-static-local-ip-address-to-vpn-mikrotik-37118e49.md
source_anchor: ""
source_lines: [1, 5]
sha256: 1be12db74e748ec5beff63dfdf4a71c10ccccfbc28eac3698da8c114663cb5ad
---

# questions-1558127-assign-static-local-ip-address-to-vpn-mikrotik-37118e49

I configure VPN on my router board and firewall NAT rules to pass all traffic through the VPN interface and all things work fine.
When VPN connection got connect to a remote VPN server, we have a local address which if we pass our traffic to it (Firewall => NAT => srcnat), our traffic passes through the VPN.
Now the problem is, every time the VPN got disconnected because of something, the VPN connection may get a different local IP address, and I should reconfigure firewall NAT rule to "srcant" the traffic through this IP.
as you can see in below image this local IP address is defined in "Address List" and is "D (Dynamic) which cause the problem, I can define static address here in "Address List" but I don't know how to force the VPN interface to use it instead of it own generated dynamically IP.
Thanks.
