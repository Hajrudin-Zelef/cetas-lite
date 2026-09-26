---
id: collect-260926-rattrapage/rattrapage/questions-5805-can-i-run-multiple-networks-on-this-mikrotik-l3-switch-7fede201
title: "questions-5805-can-i-run-multiple-networks-on-this-mikrotik-l3-switch-7fede201"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-5805-can-i-run-multiple-networks-on-this-mikrotik-l3-switch-7fede201.md
source_anchor: ""
source_lines: [1, 13]
sha256: 3c178c754f8ef9be0f2accb75806f8af282f49a9b8fc4d29f0d6a4a660b16812
---

# questions-5805-can-i-run-multiple-networks-on-this-mikrotik-l3-switch-7fede201

Assuming you don't specifically need VLANs you'd pick one port to be the master port for each group and set it to be the master port for every other port in the group. Ports with the same master port are isolated from ports with a different master port. Then you set a IP address on the master port for routing (if necessary).
For example, if you assume three groups of three ports (2-4, 5-7, 8-10) you'd do the following (note that I'm using CLI terminology so if I write /interface ethernet set master-port=ether2 where name=ether3 and you are using webfig or winbox you'd click on interface, go to the ethernet tab, click on ether3, and change the master-port option):
/interface ethernet set ether3 master-port=ether2
/interface ethernet set ether4 master-port=ether2
/interface ethernet set ether6 master-port=ether5 
/interface ethernet set ether7 master-port=ether5
/interface ethernet set ether8 master-port=ether8
/interface ethernet set ether9 master-port=ether8
/ip address add address=192.168.1.1/24 interface=ether2
/ip address add address=192.168.2.1/24 interface=ether5
/ip address add address=192.168.3.1/24 interface=ether8
Mikrotik's CRS example page has more examples along with a example if you need VLAN support.
If you need routing between switch groups or VLANs be aware that per the block diagram there is only a single 1G connection between the routing engine and the switch chip.
