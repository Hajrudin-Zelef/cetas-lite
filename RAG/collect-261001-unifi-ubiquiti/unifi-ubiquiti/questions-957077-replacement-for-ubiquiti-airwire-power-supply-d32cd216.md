---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-957077-replacement-for-ubiquiti-airwire-power-supply-d32cd216
title: "questions-957077-replacement-for-ubiquiti-airwire-power-supply-d32cd216"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-957077-replacement-for-ubiquiti-airwire-power-supply-d32cd216.md
source_anchor: ""
source_lines: [1, 11]
sha256: f782b307f9470252aae112ac8c168c585a47afbc9e481c910c008d6b65bbf677
---

# questions-957077-replacement-for-ubiquiti-airwire-power-supply-d32cd216

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have an "Ubiquiti AirWire Plug 'n Play Wireless Ethernet Cord" That has been working perfectly for years. The other night, there was a lightning strike that blew up everything connected to a specific outlet. Unfortunately, this device was one of them. However, luckily only the power supply was dead. The actual device still works. So my only problem now is how to find the very specialized replacement power cable. Is this a standard kind of cable? If so, what is it called?
here is a picture of the cable I want to replace:
aditional information: I'm not sure why the amazon page I linked to shows a USB type B input. Mine has an ethernet input on the back.
I can't find a datasheet that specifies, but given the evident era, the fact that it's 10/100 and Ubiquiti's design philosophy until recently, it's likely a 24V Passive PoE injector (PoE = Power over Ethernet.) But you could easily contact Ubiquiti support and ask. Or if you geek the right way you could measure voltage on (more like between) the blue and brown sets of wires on your working unit.
Current versions look a little different (two network ports that both use patch cables and an AC power cord) but should work the same if it's not an odd voltage. email support and inquire about that.
