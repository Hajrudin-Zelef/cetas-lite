---
id: collect-260926-mikrotik/mikrotik/questions-678924-mikrotik-all-port-block-except-port-80-6fb5ce28
title: "questions-678924-mikrotik-all-port-block-except-port-80-6fb5ce28"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-678924-mikrotik-all-port-block-except-port-80-6fb5ce28.md
source_anchor: ""
source_lines: [1, 7]
sha256: 9243ca2add901b1cdb9ef55710a754d91baac50ca2e2092cc58385cd6e8c84cf
---

# questions-678924-mikrotik-all-port-block-except-port-80-6fb5ce28

I like to block all of my mikrotik router port to get rid of hackers and only open port 80. So please someone advise me, how I can block all of my incoming and outgoing mikrotik port, accept port 80.
1 Answer 1
If you have access to the device's web interface, you can go to the IP menu, then to Firewall. Click on the NAT tab. There you can see the ports it forwards to which devices. Simply disable all ports forwarded from your in interface into your LAN.
If you have ssh access, follow the same path:
/ip firewall nat
And print to see the existing rules. You can remove a specific (numbered) item to get rid of all the forwards except port 80 (if it exists) and add an entry for port 80 like so:
/ip firewall nat add chain=dstnat protocol=tcp dst-port=80 action=dst-nat to-addresses=<internal IP address> to-ports=80
