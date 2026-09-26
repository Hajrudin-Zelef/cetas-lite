---
id: collect-260926-mikrotik/mikrotik/questions-1153328-mikrotik-creating-letsencrypt-cert-fails-on-routeros-v7-51135b37
title: "questions-1153328-mikrotik-creating-letsencrypt-cert-fails-on-routeros-v7-51135b37"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1153328-mikrotik-creating-letsencrypt-cert-fails-on-routeros-v7-51135b37.md
source_anchor: ""
source_lines: [1, 6]
sha256: cfff081088a71b3b9a7d8b0912112d62add4efcb857294b2b11e9b3776cdba54
---

# questions-1153328-mikrotik-creating-letsencrypt-cert-fails-on-routeros-v7-51135b37

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
I was installing a LetsEncrypt cert on a MikroTik using only an IPv6 Global Unicast address which had me investigating IPv6 configuration. But this was a red-herring, and the issue I had could happen equally with an IPv4 address:
In addition to opening TCP/80 & TCP/443 in the FW for LetsEncrypt to reach the router, I had to also add "allow" addresses in ip/service for www(TCP/80) & www-ssl (TCP/443) which LetsEncrypt could reach to complete the cert registration process. Once I did this, the error cleared.
MikroTik does this to stop users from cutting their own heads off by mistakenly cracking-open TCP/80 and/or TCP/443 on the input chain and unwittingly exposing a management interface to an unintended audience. So you have to make an express decision to allow this connectivity. Same is also true for SSH connectivity
Note that although ip/services lives in the IPv4 part of the interface, you add IPv6 addresses here also to allow IPv6 HTTP/S connectivity.
Hopefully this will save other a lot of wasted time... D'Oh!
