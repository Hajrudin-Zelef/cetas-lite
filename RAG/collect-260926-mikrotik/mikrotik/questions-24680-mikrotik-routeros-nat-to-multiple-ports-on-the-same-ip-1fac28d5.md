---
id: collect-260926-mikrotik/mikrotik/questions-24680-mikrotik-routeros-nat-to-multiple-ports-on-the-same-ip-1fac28d5
title: "questions-24680-mikrotik-routeros-nat-to-multiple-ports-on-the-same-ip-1fac28d5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-24680-mikrotik-routeros-nat-to-multiple-ports-on-the-same-ip-1fac28d5.md
source_anchor: ""
source_lines: [1, 4]
sha256: 7503375cf02993bd655f2356c9496247b1492d940af59dd2bb3f47517404617e
---

# questions-24680-mikrotik-routeros-nat-to-multiple-ports-on-the-same-ip-1fac28d5

Some network services require a group of open ports, and those might not even be next to each other (by port number, I mean).
What I'm doing still is I create a dst-nat rule for each port that should be mapped to the server, then enable/disable the rule as needed.
In my specific case, that's a bit tedious because sometimes the ports should point to a different computer for a short period of time which means I have to manually change all the entries to the new IP each time I do this. Also, I wish I could just say "disable service x" instead of figuring out which port rules belong together.
Is there a better and more efficient way of doing this?
