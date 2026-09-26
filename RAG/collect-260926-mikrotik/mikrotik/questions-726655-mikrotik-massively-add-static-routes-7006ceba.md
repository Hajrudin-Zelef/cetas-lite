---
id: collect-260926-mikrotik/mikrotik/questions-726655-mikrotik-massively-add-static-routes-7006ceba
title: "questions-726655-mikrotik-massively-add-static-routes-7006ceba"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-726655-mikrotik-massively-add-static-routes-7006ceba.md
source_anchor: ""
source_lines: [1, 10]
sha256: 70968a0b746e1c23143f61bbbe97ea0c0b7aae950c3fe7e9501e6750c500a8cd
---

# questions-726655-mikrotik-massively-add-static-routes-7006ceba

I'm wondering of it is possible to add static routes to Mikrotik massively. For example, I have list of Facebook servers which I need to route and how can I add them in terminal by one or 2 commands rather than entering manually in WinBox? Thank you in advance.
2 Answers 2
You can add static routes (or anything else you can do in Winbox really) via the terminal.
Just log in to the terminal (either via winbox, SSH or Telnet) and run the following command to add a static route.
/ip route add dst-address=DESTINATION_IP_ADDRESS_HERE gateway=YOUR_GATEWAY_ADDRESS_HERE
Repeat as many times as you need to add all the static routes you want.
Here's the official documentation on IP > Routes
Yes, it is possible. But I wouldn't do that through WinBox. As Cha0s mentioned, it's better to do that kind of thing through the terminal instead.
Your question wasn't quite clear to me. Do you want to force traffic to Facebook to go out through an specific interface using a route that is already there? If that's the case, the better solution is to create a routing mark in that route, create an address list and a mangle rule to mark packets that match the IP addresses stored in the list. Much more organized and easy to manage, and you can use MikroTik's script language to populate the list.
Hope this helps!
