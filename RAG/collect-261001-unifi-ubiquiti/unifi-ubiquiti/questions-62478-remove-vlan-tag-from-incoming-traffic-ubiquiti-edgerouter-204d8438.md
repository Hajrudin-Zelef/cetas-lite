---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-62478-remove-vlan-tag-from-incoming-traffic-ubiquiti-edgerouter-204d8438
title: "questions-62478-remove-vlan-tag-from-incoming-traffic-ubiquiti-edgerouter-204d8438"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-62478-remove-vlan-tag-from-incoming-traffic-ubiquiti-edgerouter-204d8438.md
source_anchor: ""
source_lines: [1, 10]
sha256: 89c14636e3a292c957d80c7008ff49e77078e21f5ba35904661e1c42668ccab7
---

# questions-62478-remove-vlan-tag-from-incoming-traffic-ubiquiti-edgerouter-204d8438

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
On eth4 I have incoming traffic from the WiFi AP. Some is tagged with VID 102 and some with VID 950. I want the Ubiquiti Edgerouter (which is currently acting as a managed switch) to take the 950 VLAN traffic and simply remove the tag, passing it off as now untagged traffic.
Does this question even make sense? If it does, how would I accomplish this?
The WAP tags frames on its Ethernet port depending on their source/destination SSID (SSID/VLAN association). That is required to keep SSIDs and VLANs separate.
Simply configure the switch port that is passing on the traffic as untagged port for VID 950. VLAN 950 frames will continue to be tagged between WAP and switch. When passing on those frames the switch removes the tag, according to the egress port's configuration.
