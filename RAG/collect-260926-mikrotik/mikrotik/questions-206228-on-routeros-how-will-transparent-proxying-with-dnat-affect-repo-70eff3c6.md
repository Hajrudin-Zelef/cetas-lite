---
id: collect-260926-mikrotik/mikrotik/questions-206228-on-routeros-how-will-transparent-proxying-with-dnat-affect-repo-70eff3c6
title: "questions-206228-on-routeros-how-will-transparent-proxying-with-dnat-affect-repo-70eff3c6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-206228-on-routeros-how-will-transparent-proxying-with-dnat-affect-repo-70eff3c6.md
source_anchor: ""
source_lines: [1, 8]
sha256: 90b93b71049a5d08fc90b15b6f5f1e28dc82f677ff1b8ab78b9519c52e5818d9
---

# questions-206228-on-routeros-how-will-transparent-proxying-with-dnat-affect-repo-70eff3c6

I have a box running Mikrotik RouterOS, which is set up to do transparent web proxying, as described here.
In short, this means that I have a firewall rule for destination NAT causing any port 80 traffic to get redirected to port 8080 on the router, which is received by the Mikrotik local web proxy. The local web proxy then makes the web request on the client's behalf, in this case to a parent web proxy server (which in turn does the real web request).
My question is, how will this two-part process get reported in the logging of traffic flow information (netflow)?
Looking at the logged information, what I seem to be seeing is this:
- One flow recorded from client machine (private IP address) to remote proxy (8080)
- Another flow recorded from router to remote proxy (8080)
The original request that the client made to port 80 isn't recorded.
I want to write code to analyse traffic usage, so I want to be sure I'm not losing information if I discard the latter of these.
