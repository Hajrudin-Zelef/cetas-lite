---
id: collect-261001-general-networking/general-networking/questions-51567-direction-incoming-vs-outgoing-in-logs-1c9f7902
title: "questions-51567-direction-incoming-vs-outgoing-in-logs-1c9f7902"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-51567-direction-incoming-vs-outgoing-in-logs-1c9f7902.md
source_anchor: ""
source_lines: [1, 13]
sha256: 4fe7864889b7393a01d561bc892da70f3189b09a47aed32b612a94b8fa679649
---

# questions-51567-direction-incoming-vs-outgoing-in-logs-1c9f7902

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm trying to understand some Fortinet firewall logs but I'm not sure I fully understand what is being logged by the firewall when it comes to direction (Incoming vs Outgoing)
I'm not sure why that would be categorized as outgoing instead of incoming. Since the traffic appears to be initiated from an external source. Would it not be incoming traffic? Or is it because the last interface is it hits is the LAN interface and it's "outgoing" from there?
Fortigate firewalls log each connection with the initiator as source, ie. for normal Internet access the LAN client (initiator) is always the source. For virtual IPs/port forwarding the WAN client is the source. Often, this does not reflect the data flow direction.
Your example is a virtual IP/forwarded port from a WAN client (srcintfrole) to the LAN server (dstintfrole).
I'd guess the "outgoing" indicates the actual data flow direction for the log item - strangely, I don't see this column on our 100D (v5.6.5)...
The connection initiated from external networks which hits egress interfàce of firewall and transverse to Lan interfàce of firewall .. considered as incoming traffic or inbound traffic ..
The connection initiated from internal network which hits on ingress interface of firewall and transverse to egress interfàce where ISP link is connected .. consider as outbound traffic or outgoing traffic ..
