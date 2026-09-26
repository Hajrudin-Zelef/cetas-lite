---
id: collect-260926-mikrotik/mikrotik/questions-140518-why-use-mikrotik-hotspot-gateway-f344cf8c
title: "questions-140518-why-use-mikrotik-hotspot-gateway-f344cf8c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-140518-why-use-mikrotik-hotspot-gateway-f344cf8c.md
source_anchor: ""
source_lines: [1, 15]
sha256: 76c3f8c5a81532bccd14aa5f64250077f4eb3eff38b143d6270eb2b0229b0dc2
---

# questions-140518-why-use-mikrotik-hotspot-gateway-f344cf8c

Information Security is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
First of all Let me mention that this whole experiment is for learning purpose and not intended to cause harm or damage.
I am a novice networking enthusiast.Recently I tried to penetrate a wifi network which is open and has no security but authorization is done through a mikrotik hotspot gateway.When a user is connected to this network it redirects user to the gateway webpage on a server and asks to enter username and password for authentication.I could connect to the network after I managed to monitor devices that connect to this hotspot, record their MAC addresses and spoof my MAC address to impersonate one of the authorized devices.I have three questions;
Is it possible to access the routers web interface and modify setting if the administrator changed the default password/username of the router?
Why should a network administrator use this method to secure a network since this can be bypassed by mac spoofing and it is vulnerable compared to a WPA/WPA2 security?
Is it possible to fix this vulnerability of hotspot gateway?
Why should a network administrator use this method to secure a network since this can be bypassed by mac spoofing and it is vulnerable compared to a WPA/WPA2 security?
There is no such thing as a 100% secure solution and the more security you want the more expensive and/or unusable it usually gets. To make MAC spoofing impossible inside a WLAN you probably need some kind of authentication of the clients which is more robust than the initial username and password. 802.X provides this but then you would need to install authentication credentials (for example certificates) on each device and maybe special software too. You don't want this in public hotspots so you accept instead the risk of MAC spoofing.
Is it possible to fix this vulnerability of hotspot gateway?
There are solutions but they affect the usability too much if used in a public hotspot. If used inside a company where one has more control over the connecting devices one can detect and block MAC spoofing.
