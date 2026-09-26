---
id: collect-260926-mikrotik/mikrotik/questions-588036-mikrotik-download-limited-c6605f5f
title: "questions-588036-mikrotik-download-limited-c6605f5f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-588036-mikrotik-download-limited-c6605f5f.md
source_anchor: ""
source_lines: [1, 14]
sha256: 63893569eb4eef07d46ba18c56ac0ed49355ef1f706c5ac2a3d9e978cccd7b5a
---

# questions-588036-mikrotik-download-limited-c6605f5f

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Mikrotik routerboard RB2011UiAS.
I have an issue that when downloading we only get between 10mbit (1000pps) - 13mbit(1200pps) when connecting directly into the ISP provide with the Mikrotik.
If I put a hp procurve switch between the Mikrotik and ISP provider we get the full 100mbit download, how ever without we only get 10mbit - 13mbit.
Are they any settings or issue with the Mikrotik connecting directly with the ISP?
This might be a ROS bug (what version are you running). Have you tried creating a supout file and involving 'tik support?
As NickW points out this seems like a port negotiation issue.
Putting the Procurve between them certainly goes a long way towards pointing that.
Also, RB2011 has 5 Gbit and 5 10/100 ports. Any change in behavior if you move the WAN port to other ports either on the Gbit half (ie 2-5) or over to the 10/100 side (6-10)?
