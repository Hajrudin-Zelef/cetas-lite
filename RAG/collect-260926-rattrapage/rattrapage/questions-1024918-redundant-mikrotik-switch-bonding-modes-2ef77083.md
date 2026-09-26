---
id: collect-260926-rattrapage/rattrapage/questions-1024918-redundant-mikrotik-switch-bonding-modes-2ef77083
title: "questions-1024918-redundant-mikrotik-switch-bonding-modes-2ef77083"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/ai-llm/questions-1024918-redundant-mikrotik-switch-bonding-modes-2ef77083.md
source_anchor: ""
source_lines: [1, 20]
sha256: 92e8f8230ecf53dbac67ca5091f076d9597a2da5f1dac3a2dfb6413c2b26efbb
---

# questions-1024918-redundant-mikrotik-switch-bonding-modes-2ef77083

All the posts I can find about redundant switches only address the servers connected to the switches. I need to confirm my configuration on both the servers and the router handling the traffic.
We're using Mikrotik routers and switches and I'm planning the following configuration.
                         +----------+
            +------------+ CCR 1072 +-------------+
            |       sfp1 +----------+ sfp2        |
            |         bond (balance-xor)          |
      +-----+----+                          +-----+----+
      |          |                          |          |
      | switch A +                          + switch B |
      |  CRS 317 |                          |  CRS 317 |
      +-----+----+                          +-----++---+
            |                                     |  
            |             +-------+               |
            +-------------+ host1 +---------------+
                     eth0 +-------+ eth1
                     bond (active/backup)
There's no link or stacking of the switches. I don't need to load balance I'm just looking for redundancy at this stage. I had considered active/backup on the router side too but if sfp1 fails on the router and eth0 on the host can still reach the switch I wondered if that would stay up and traffic would stop.
Am I on the right track here?
Is there anything I need to consider like STP for example?
Should I use 802.3ad instead?
