---
id: collect-260926-mikrotik/mikrotik/questions-17681-hp-1810-24g-and-mikrotik-cloudcore1036-1gb-link-negotiation-prob-3d7633f8
title: "questions-17681-hp-1810-24g-and-mikrotik-cloudcore1036-1gb-link-negotiation-prob-3d7633f8"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["full-duplex"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-17681-hp-1810-24g-and-mikrotik-cloudcore1036-1gb-link-negotiation-prob-3d7633f8.md
source_anchor: ""
source_lines: [1, 8]
sha256: fd46a945f4bc98d194f2d3c5e05d6f627ee89680d74269a94bcddde53fc67ab8
---

# questions-17681-hp-1810-24g-and-mikrotik-cloudcore1036-1gb-link-negotiation-prob-3d7633f8

MY QUESTIONS ARE:
- What could be the cause of this nondeterministic HP switch behaviour ?!
- Why sometimes devices make 100Mb link and sometimes 1Gb ?!
At above picture You can see part of my network scheme. Main router based on Mikrotik CloudCore1036 within one bridge and HP 1810-24G switches within default VLAN.
I tried to make bonding between them using 802.3ad protocol because I know that my HP support this old one and I didn't have time to make additional experimental configuration (balanced-rr doesn't work because makes huge issue with ARP broadcasting along the whole network !!!). I noticed that bonding didn`t work properly because of links speeds between devices:
- CC1036 announce to the HP all full-duplex link speeds like 10/100/1000,
- HP announce only full-duplex 10/100 at E23 but at E24 also 1000 for example.
What is more strange sometimes HP gives to the LAN hosts 100Mb speed and sometimes 1Gb speed at the same ports - hosts always have 1Gb NICs. All HP ports are at admin mode and they have Autonegotiation enabled - CC1036 also. I cheked whole wiring between CC1036 and HP and all strands are ok. Top cable lenght is about 11m Cat5E.
