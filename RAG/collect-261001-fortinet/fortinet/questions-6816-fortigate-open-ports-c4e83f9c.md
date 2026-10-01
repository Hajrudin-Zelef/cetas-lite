---
id: collect-261001-fortinet/fortinet/questions-6816-fortigate-open-ports-c4e83f9c
title: "Fortigate Open Ports"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-6816-fortigate-open-ports-c4e83f9c.md
source_anchor: ""
source_lines: [1, 15]
sha256: 869253d80fac9d93854b487ec4ba860cb17c7872be410d08ab6fbb6034ce0e96
---

# Fortigate Open Ports

*Score : 3 | Source : https://networkengineering.stackexchange.com/questions/6816/fortigate-open-ports*

I need to open this port in Fortigate:
  443 PSOM/TLS Outbound Data sharing sessions
But i dont know where I can open. I created a Custom Service, chose the option of TCP, but do not know if it's the right thing.
Does anyone know how can I open this port?
Thanks

---

### Reponse (acceptee) — score 3

The FortiGate doesn't care which protocol is running over the port 443, so you just need to create a policy and select the corresponding interfaces/addresses and as service you can select HTTPS. If it's a policy from internal network to WAN, be sure to select NAT also
