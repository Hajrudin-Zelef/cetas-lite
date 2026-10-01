---
id: collect-261001-general-networking/general-networking/questions-2266-private-vlans-on-a-switch-that-doesnt-support-private-vlan-trunks-85fd22ea
title: "questions-2266-private-vlans-on-a-switch-that-doesnt-support-private-vlan-trunks-85fd22ea"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-2266-private-vlans-on-a-switch-that-doesnt-support-private-vlan-trunks-85fd22ea.md
source_anchor: ""
source_lines: [1, 3]
sha256: d892c8e189fd6f6a633df050a79558cae743721d77b8dc7f0c912343a72127b0
---

# questions-2266-private-vlans-on-a-switch-that-doesnt-support-private-vlan-trunks-85fd22ea

I have a Catalyst switch that doesn't support PVLAN trunks (Sup4, 4500, 12.2(54)S). I have multiple other Catalyst switches, 3750 metro, that do.
Am I correct in assuming that because the C4500 does not support PVLAN trunks with that Supervisor, that a normal trunk port would not work for the 3750s to share a private VLAN with the 4500?
The desired scenario is that the 4500 hosts primary VLAN 500 and isolated vlans 501,502 with 501,502 residing on other 3750s in the network. I assume that in order to achieve the desired config, I have to have the 3750 promiscuous ports connected to access ports on the 4500. Is this correct?
