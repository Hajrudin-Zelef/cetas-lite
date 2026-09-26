---
id: collect-260926-mikrotik/mikrotik/questions-1294865-how-to-configure-hotspot-on-a-mikrotik-router-to-provide-user-4f84e487
title: "questions-1294865-how-to-configure-hotspot-on-a-mikrotik-router-to-provide-user--4f84e487"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1294865-how-to-configure-hotspot-on-a-mikrotik-router-to-provide-user--4f84e487.md
source_anchor: ""
source_lines: [1, 14]
sha256: e3928a3ccea36f068d9a8fad73093d0c1e692460874da1211ccda7264bd7d1e9
---

# questions-1294865-how-to-configure-hotspot-on-a-mikrotik-router-to-provide-user--4f84e487

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have recently purchased a Mikrotik router (RB-951G). There are multiple wireless routers available in my network. They are located in multiple floors. I have connected them to a Cisco small business switch. Further, I have connected the switch to the 2nd port of the Mikrotik router. Mikrotik is then connected to our main firewall. Also it is worth mentioning that the wireless routers are D-Link and TP-Link manufactured.
By configuring a Radius server on the Mikrotik router, is it possible to
provide username and password to every individual that is connecting
to all wireless devices with username and password?
I believe it's possible to install both captive portal and radius on microtik routers. That being said it really depends on the amount of users you expect to simultaneously connect to your authentication router, and how much horsepower your microtik has under the hood. It may support 10, 20, or 200 depending on model.
You may have to customize these instructions to your needs:
Just so you are aware, this is a bad idea if you are running any kind of business communication over the wireless. Anyone would be able to capture your network traffic in real time unadulterated.
If you find your router doesn't have the guts to support the simultaneous users you want connecting. You can go the PFsense route, which also has radius and captive portal. And is pretty easy to get going.
