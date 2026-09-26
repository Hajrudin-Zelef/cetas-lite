---
id: collect-260926-mikrotik/mikrotik/questions-1130906-port-forwarding-between-tp-link-and-mikrotik-to-server-cant-ac-7cb18150
title: "questions-1130906-port-forwarding-between-tp-link-and-mikrotik-to-server-cant-ac-7cb18150"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1130906-port-forwarding-between-tp-link-and-mikrotik-to-server-cant-ac-7cb18150.md
source_anchor: ""
source_lines: [1, 12]
sha256: a06feff368421a629b48bffab1373fab827184369fc8eb2b884f846d71aed3bf
---

# questions-1130906-port-forwarding-between-tp-link-and-mikrotik-to-server-cant-ac-7cb18150

My customer has this setup TP-Link TD-8840T (ISP connection) -> RB2011UiAS-2HnD-IN (the main router that divedes all the traffic).
The RB2011 has a ip address 192.168.1.100 (connected to tp-link device) 192.168.88.1 (gateway) and the server 192.168.88.7
I've setup a port formwarting for ssh on the RB2011 according to this tutorial http://www.icafemenu.com/how-to-port-forward-in-mikrotik-router.htm
Then when trying to connect to the router address on 192.168.88.1:22 the ssh console shows up.
I've setup a NAT on the tp link which should work the same (forwarding 22 port to 192.168.1.100:22). But cant get it to work and I don't know if this is due to the RouterOS lacking in some configuration or the tp-link device doe's not forward it further.
The RouterOS is very intimidating to me at this I've fond something called Harpin NAT tried to make the configuration but with no luck.
partial solution
I've made a workaround which right now is sufficient. I've setup the TP-Link device to work as a bridge.
the above solution doesn't work after all
It seems the problem is the Mikrotik (Routerboard) router - which can be accessed now via ssh but the port forwarding setup I've made doesn't work.
It should now forward the connection on the 2323 port to the computer with address 192.168.88.7 on port 22 (ssh port).
The forwarding is blocked for some reason. I also tried to disable all other rules that could be blocking external connections with no luck as well.
