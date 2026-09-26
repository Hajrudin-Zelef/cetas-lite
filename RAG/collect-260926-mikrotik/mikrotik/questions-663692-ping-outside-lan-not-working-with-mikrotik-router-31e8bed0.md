---
id: collect-260926-mikrotik/mikrotik/questions-663692-ping-outside-lan-not-working-with-mikrotik-router-31e8bed0
title: "questions-663692-ping-outside-lan-not-working-with-mikrotik-router-31e8bed0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-663692-ping-outside-lan-not-working-with-mikrotik-router-31e8bed0.md
source_anchor: ""
source_lines: [1, 17]
sha256: 1da014d9b85a0619e04f92ae5a12db489afa95953b043f5bda00035d8f637184
---

# questions-663692-ping-outside-lan-not-working-with-mikrotik-router-31e8bed0

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I try to ping outside internet from Mikrotik terminal but I can't. I have Mikrotik routerboard 750GL. I tried to ping gmail.com, It's IP is: 74.125.236.53
ping 74.125.236.53
I got output:
timout
But If I ping my local system from Mikrotik terminal I am able to do, my local system IP is:xxx.xxx.xxx.xxx
ping xxx.xxx.xxx.xxx
I got output:
connected
Why can not I able to ping outside my LAN connection?
Go to IP -> Firewall -> NAT, click the blue + to add a rule, Chain is srcnat, Out interface is ether1 - or whatever ether you get your internet from. Then go to the Action tab and select Masquerade.
ping 8.8.8.8?ping 8.8.8.8[myuser@MikroTik] > ping 8.8.8.8 Output:8.8.8.8 timeout 8.8.8.8 timeoutping 8.8.8.8from my local system, I mean I do ping through router but not from within router
