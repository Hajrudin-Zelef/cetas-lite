---
id: collect-261001-cisco/cisco/questions-10158-asa-5520-physical-ports-b13c0e04
title: "questions-10158-asa-5520-physical-ports-b13c0e04"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-10158-asa-5520-physical-ports-b13c0e04.md
source_anchor: ""
source_lines: [1, 17]
sha256: 2cb966bdace811b5056c04793e109f3523de7fba0214eeeab1d8813fbf445811
---

# questions-10158-asa-5520-physical-ports-b13c0e04

I am Cisco certified and have worked for an ISP for the last 5 years. I have made the transition to the LAN side as a Network Administrator so I am a novice at certain things. With that being said, below is the configuration on one of my ASA5520s. I have changed addressing for security, but this is how it looks.
interface GigabitEthernet0/0.2
 nameif Inside
 security-level 100
 ip address 10.x.x.x 255.255.255.248
!
interface GigabitEthernet0/0.4
 nameif DMZ
 security-level 50
 ip address 10.x.x.x 255.255.255.248
!
interface GigabitEthernet0/1
 nameif Outside
 security-level 0
 ip address 159.x.x.x 255.255.255.0
My question is this: I have four physical ports on the back of my ASA, why don't they show up in the running config ,,,why?
And my second question: is it ok not to have a VLAN associated with my sub-interfaces?
