---
id: collect-260926-mikrotik/mikrotik/questions-1413176-nat-to-other-lan-using-mikrotik-routeros-2728a8e3
title: "questions-1413176-nat-to-other-lan-using-mikrotik-routeros-2728a8e3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1413176-nat-to-other-lan-using-mikrotik-routeros-2728a8e3.md
source_anchor: ""
source_lines: [1, 6]
sha256: 67c15528f3cbe522f3effdce284b425d3c9a6c944ef4b2e12da374ac7975cdfc
---

# questions-1413176-nat-to-other-lan-using-mikrotik-routeros-2728a8e3

I have a mikrotik rb4011igs with following connections:
- WAN on port1
- Bridge with DHCP for lan on ports2-9
- Other LAN on port10 (I cant make any changes to this networks router, so cant change the default gateway for this LANs devices)
Now I want to be able to access the other LAN devices with their addresses from my LAN(other LANs devices dont need to access the mikrotiks LAN).
As far as I understand, I need to create some NAT rules with some kind of masquerading, but unfortunately I'm not sure what to do exactly.
