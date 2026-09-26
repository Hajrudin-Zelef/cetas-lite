---
id: collect-260926-rattrapage/rattrapage/questions-26591-how-to-configure-802-1q-trunk-on-mikrotik-router-switch-405a775d
title: "questions-26591-how-to-configure-802-1q-trunk-on-mikrotik-router-switch-405a775d"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/ai-llm/questions-26591-how-to-configure-802-1q-trunk-on-mikrotik-router-switch-405a775d.md
source_anchor: ""
source_lines: [1, 6]
sha256: ea3bea0b9d7d423899c56173961a89c904b0b81fd92f6f1b23115eccbafc04b8
---

# questions-26591-how-to-configure-802-1q-trunk-on-mikrotik-router-switch-405a775d

Cisco L2 switch <--trunk--> CRS125 <--trunk--> RB750G
- RB750G is the only device that handles routing
- CRS125 and RB750G has interfaces with untagged traffic
I am familiar with the config on Cisco side, but had no success on Mikrotik side (keep getting locked out and have to reset the device). Mikrotik has really poor documentation and the config seems inconsistent on different models.
With RB750G, I managed to get VLAN to work by setting master port to none for each port, creating a bridge for each VLAN, then creating a bridge for each trunk with allowed vlans added to the bridge following MUM tutorial for VLAN in MikroTik. Isn't this effectively soft switching everything? Not to mention this get complicated really quick as number of vlan grows.
I had no success with the switch chip config. "Management IP Configuration" section made no sense to me and following the wiki guarantees a lockout. Do I have to create the VLANs in both switch config vlan database and router interface vlan database?
