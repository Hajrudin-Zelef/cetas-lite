---
id: collect-260926-mikrotik/mikrotik/questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp-4100a377-2
title: "questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp--4100a377"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp--4100a377.md
source_anchor: ""
source_lines: [120, 120]
sha256: f353841f860eda87445b70969ba61320f9dd0dfa3642e42a7e141a9c974e613f
---

# questions-33391-why-central-mikrotik-router-not-routing-secondary-networks-pptp--4100a377

gateway=172.16.1.1? Your default gateway? Your 1.2.x.x NICs are on a different subnet, each needing their own default gateway for inter-LAN routing.172.16.1.1- vpn ip of main1;172.16.1.2- vpn ip of secondary2;172.16.1.3- vpn ip of secondary3. See the/ppp secretconfiguration section of main1. As @mmv-ru correctly answered, I missed the routeadd distance=1 dst-address=172.16.1.0/24 gateway=172.16.1.1. The problem is solved. Thanks to @mmv-ru!
