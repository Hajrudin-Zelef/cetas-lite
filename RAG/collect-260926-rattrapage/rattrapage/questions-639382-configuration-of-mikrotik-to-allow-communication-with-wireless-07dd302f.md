---
id: collect-260926-rattrapage/rattrapage/questions-639382-configuration-of-mikrotik-to-allow-communication-with-wireless-07dd302f
title: "questions-639382-configuration-of-mikrotik-to-allow-communication-with-wireless--07dd302f"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-rattrapage/servers-reviews/questions-639382-configuration-of-mikrotik-to-allow-communication-with-wireless--07dd302f.md
source_anchor: ""
source_lines: [1, 24]
sha256: a69515827c06d0a06aaf386d6e7f749d1f08b61f787f8f22869cb18ad15c2c18
---

# questions-639382-configuration-of-mikrotik-to-allow-communication-with-wireless--07dd302f

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm managing an relatively small office and my basic network setup is as follows:
All LAN computers are connected to a central switch which is connected with a Mikrotik Router. LAN subnet is 192.168.1.1/24
I have three different Internet lines also connected to Mikrotik
I have three wireless routers that are connected to LAN (their WAN ports). Wireless routers are configured with DHCP (range 192.168.0.1/24)
Wireless clients are able to connect with LAN devices, but from LAN I'm unable to connect with wireless clients (it makes sense since trying to connect with 192.168.0.1/24 it will go to the router and router will not know what to do about it).
What would be the right way to configure Mikrotik and/or Wireless router to allow connections from LAN to wireless clients?
If 2 client connect to wireless network using same Wireless routers there is no problem but when 2 client connect with separate Wireless routers can not see themselves. For fix this problem, there is 2 ways:
Connect Wireless routers to LAN using their LAN Ports (Not using wan port). In this case wireless client get IP address in range 192.168.1.0/24
Set each wireless network address with a specific range, for example
wireless A has IP address in range 192.168.2.0/24
wireless A has IP address in range 192.168.3.0/24
and set wireless WAN Ports IP address in mikrotik to make static for example
wireless A has IP address in range 192.168.1.10
wireless A has IP address in range 192.168.1.11
Now when a client attempt to connect other client if IP address not in same range forward IP address to mikrotik.
In mikrotik
Masqurade local IP address to local IP address destinations
Add route to routing other range to their wireless router
