---
id: collect-260926-mikrotik/mikrotik/questions-1008728-mikrotik-ban-host-if-there-are-too-many-http-requests-from-it-19b5fedb
title: "questions-1008728-mikrotik-ban-host-if-there-are-too-many-http-requests-from-it-19b5fedb"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-1008728-mikrotik-ban-host-if-there-are-too-many-http-requests-from-it-19b5fedb.md
source_anchor: ""
source_lines: [1, 23]
sha256: 708f7b68ee69057e5e73ac2abbeaa68e6a086e1066357c0ef569c3e2f13e445b
---

# questions-1008728-mikrotik-ban-host-if-there-are-too-many-http-requests-from-it-19b5fedb

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Yesterday I noticed a strange activity of my humble web server: it was moderately warm, was twitching by HDD heads and LAN activity was unusually high.
When I looked to logs, found that some host is scanning my web server for documents using file names brute force.
Is there any protection against such a brute force attack that I could implement in RouterOS?
Yes, there is a protection. Basically, you will have to add a firewall rule to detect such hosts (criteria: several tcp/port 80 connections from same host), and when you have one, add that source IP to an address-list.
The answer from Benoit allows for blocking hosts making simultaneous requests. But if they are closing the connections as fast as they are opening them you might not get much of them to filter by.
Another logic is to use dst-limit which have rate limit.
We add a rate-limit list:
/ip firewall address-list
add list=rate-limit
We then set the filtering rules:
Jump to rate-limit chain for all new connection on the WAN;
Check if there have been 10 connections for the last 1 minute with bursts of 5 based on the dst-address and resetting after 2 minutes of inactivity
If the dst-limit is reached, the return action is skipped and the next action in the rate-limit chain is executed - add-src-to-address-list to the rate-limit list with 10 minutes timeout.
Finally, we add the drop in the raw, so we save on resources:
/ip firewall raw
add action=drop chain=prerouting src-address-list=rate-limit
You can play with the connection count and the timings to fine tune based on specific needs. You can monitor in the Firewall/Conncetions tab with filter for your public IP to see
