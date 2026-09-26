---
id: collect-260926-mikrotik/mikrotik/questions-840730-script-mikrotik-for-search-contains-of-strings-17666c61
title: "questions-840730-script-mikrotik-for-search-contains-of-strings-17666c61"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-840730-script-mikrotik-for-search-contains-of-strings-17666c61.md
source_anchor: ""
source_lines: [1, 10]
sha256: 59c87fc3b1e8c534026ca8001596bed7a674977a22f94be18dde8c87187927f0
---

# questions-840730-script-mikrotik-for-search-contains-of-strings-17666c61

I have queue tree name int-up_bw_duyungoffice and int-down_bw_duyungoffice. Can i grep the words duyungoffice in MikroTik CLI? So the print result is information for int-up_bw_duyungoffice and int-down_bw_duyungoffice.
Default script that i know just like below:
> queue tree print where name=int-down_bw_duyungoffice
Flags: X - disabled, I - invalid     
 0   name="int-down_bw_duyungoffice" 
     parent=TOTAL INTL DOWN - 1Mbps - Blok 1 (1:16) 
     packet-mark=packet-down_bw_duyungoffice limit-at=64k 
     queue=pcq-down-client-1Mbps(1:16) priority=8 max-limit=1M burst-limit=0 
     burst-threshold=0 burst-time=0s
But that script only shows int-down_bw_duyungoffice only.
