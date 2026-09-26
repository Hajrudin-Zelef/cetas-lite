---
id: collect-260926-mikrotik/mikrotik/questions-643075-connect-two-mikrotik-lan-7ef0094f
title: "questions-643075-connect-two-mikrotik-lan-7ef0094f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "latency"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-643075-connect-two-mikrotik-lan-7ef0094f.md
source_anchor: ""
source_lines: [1, 8]
sha256: ae495ecb36faa6d2d9c02e110c8e6af499b52c71a14173b97812d78db92b2222
---

# questions-643075-connect-two-mikrotik-lan-7ef0094f

I got two RB951G-2HnD and two RBSXTG-5HPnD-HGr2.
  (clientsOfLan1) ))> RB1 <---> SXT1 ))> <(( SXT2 <---> RB2 <(( (clientsOfLan2)
I need to 'connect' two LAN's around 951.
I can't figure out how to 'bridge' those two networks that clientsOfLan1 could see clientsOfLan2 on L2 layer.
RB1/2 is in ap bridge
SXT1 is station pseudobridge
SXT2 is ap bridge
I am quite lost. Can anyone suggest how to achieve this? I need same L2 (ethernet) network on both sides best possible latency and bandwidth. I want use/setup SXT1 and SXT2 to simulate a 'wire' that I connect between those two routerboards.
