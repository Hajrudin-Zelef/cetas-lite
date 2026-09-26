---
id: collect-260926-mikrotik/mikrotik/questions-848642-site-to-site-vpn-with-local-internet-gateways-on-mikrotik-9f334822
title: "questions-848642-site-to-site-vpn-with-local-internet-gateways-on-mikrotik-9f334822"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-848642-site-to-site-vpn-with-local-internet-gateways-on-mikrotik-9f334822.md
source_anchor: ""
source_lines: [1, 9]
sha256: 2edfb41d3a68b763cadd2ffec2bd3e9da1d20a6a7f7882c8757666745d37a926
---

# questions-848642-site-to-site-vpn-with-local-internet-gateways-on-mikrotik-9f334822

We have a corporate network in office 1 with Forefront TMG as a gateway. Internal office 1 network IP range is 192.168.0.0/24.
We have a branch office 2 with Mikrotik router (you can think of it as iptables powered Linux firewall if you are not aware of Mikrotik stuff in particular). IP range in office 2 is 192.168.88.0/24. 192.168.88.1 is an IP of the gateway.
I've setup a Site-to-site VPN connection, where only Mikrotik utilizes its PPTP Client to connect to office 1 VPN gateway on TMG.
TMG uses route relationship to communicate with computers in office 2 range.
If "Add Default Route" is checked on PPTP Client config on Mikrotik, all the traffic flows via TMG to both internet and office 1 network. Office 1 computers can access office 2 network as well, all works great.
But there is an overhead, where we don't want all the internet traffic from office 2 to flow through the TMG in office 1. We only need the office 1 IP's routed through VPN, while everything else goes through Internet uplink in office 2.
So I've disabled the "Add default route" checkmark in PPTP Client config and used Mingle setup in Firewall on Mikrotik to add a routing mark to all the traffic that targets office 1 network. In Routes table on Mikrotik I've basically added a route that says: all the traffic with office 1 mark goes through VPN gateway.
This almost gives me what I want. Office 2 connects to office 1 IP's over VPN gateway, other requests go through local Internet uplink. But the only thing that doesn't work in this setup, with "Add Default Route" unchecked, is that office 1 computers can't get neither to VPN'ed Mikrotik IP, or any IP in the office 2 network. Mikrotik basically doesn't route traffic that is coming from Office 1. I've tries several approaches (routes) based on routing mark as well, but neither allowed me to have office 1 access office 2. Only if "Add default route" is check, I can connect both ways.
Please assist in giving a hint of what's so special behind this "Add default route" in my setup, since I basically manually adding the same, which enables only half of communication setup to work.
