---
id: collect-260926-mikrotik/mikrotik/questions-662077-mikrotik-add-http-header-parameter-to-requests-ae58d873
title: "questions-662077-mikrotik-add-http-header-parameter-to-requests-ae58d873"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-662077-mikrotik-add-http-header-parameter-to-requests-ae58d873.md
source_anchor: ""
source_lines: [1, 9]
sha256: a5fb4f01cb121bab08b7517cd183d46452d660c0aa66e90c9483af3022b2379a
---

# questions-662077-mikrotik-add-http-header-parameter-to-requests-ae58d873

This question already has an answer here:

I want to add a parameter of MikroTik to all HTTP header requests using MikroTik. So all client and servers that connected to the MikroTik can trace the parameter in their HTTP header.

How can I achive this?

|  | This question already has an answer here: I want to add a parameter of MikroTik to all HTTP header requests using MikroTik. So all client and servers that connected to the MikroTik can trace the parameter in their HTTP header. How can I achive this? | 

|  | You cannot manipulate the HTTP headers with Mikrotik. You need a reverse proxy like HAProxy or Nginx to achieve this. **Update:** Here I describe a solution on how you can add Nginx or HAProxy support on Mikrotik NAT to two different servers on the same port via hostname with Mikrotik RB2011 |
