---
id: collect-260926-mikrotik/mikrotik/questions-1300674-mikrotik-nat-halves-wan-speed-b01cd1d7
title: "questions-1300674-mikrotik-nat-halves-wan-speed-b01cd1d7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "research"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-1300674-mikrotik-nat-halves-wan-speed-b01cd1d7.md
source_anchor: ""
source_lines: [1, 17]
sha256: fba5eec16f586ea7ea3bbe2e15194955cbc15becbc9b153451f6b1ceb87c4b1b
---

# questions-1300674-mikrotik-nat-halves-wan-speed-b01cd1d7

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a 951G-2HnD as a router. All connections are via ethernet cables. I observe the following:
When I connect to the ISP via router, speed is about 270 Mbit/s, while ISP claims my plan has 500 Mbit/s speed limit.
When I connect cable from ISP directly to ethernet port on my PC, I indeed get about 500 Mbit/s download speed.
When I download files from another PC connected to the same router, the speed is about 900 Mbit/s which is correct for gigabit LAN.
Now because of 3, I think that the problem is not in router processing power. I suppose something is wrong with NAT processing. How do I improve speed so NAT connection is as fast as direct one? Or at least how do I debug the problem?
Just in case, I disabled all NAT rules except the main one:
David's answer is correct. To summarize, the hardware is capable of switching at the speed you got, but only routing at a fraction of that. Therefore, the problem is the routing, not just the NAT, although that probably doesn't help.
Take a look at these results for routers that should be capable of routing that much traffic:
Any of these three should be enough unless you're doing a large amount of very small packet routing.
Another possibility is the CHR product - you can run your router in a VM and give it as much CPU and RAM as you need. $30 covers the cost for a 1Gbps cap.
The 951G-2HnD has a rated peak routing speed of 250-300Mbps. It does its switching in hardware and should be able to switch at wire speed. So it sounds like you're getting roughly the performance the hardware is capable of.
