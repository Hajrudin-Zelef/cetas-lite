---
id: collect-261001-meraki/meraki/questions-39296-802-1q-vlan-tagging-over-wifi-meraki-3bccdbab
title: "questions-39296-802-1q-vlan-tagging-over-wifi-meraki-3bccdbab"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-39296-802-1q-vlan-tagging-over-wifi-meraki-3bccdbab.md
source_anchor: ""
source_lines: [1, 17]
sha256: 29293cf57e3c728d5061940686673176cd20f1102857740ff4453e7159ec07ab
---

# questions-39296-802-1q-vlan-tagging-over-wifi-meraki-3bccdbab

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've recently inherited an all wireless network that has every single client in the native VLAN, at multiple sites. The site in question currently has around 550 clients on a /22, which I feel is causing issues during peak times with an alarming number of broadcasts.
So my question is:
Will splitting this up into separate VLANs while allowing every VLAN to be passed along the trunk ports (All AP switchports are setup as Trunk ports), and placing certain clients on their preferred VLAN, improve performance?
With Cisco Meraki APs you can create group policies that will tag an individual client to a particular VLAN
I would segment the VLANs based on the local structure.
e.g.:
Floor Level1 => VLAN101
Floor Level2 => VLAN102
This is a quite static approach in segmenting your network and much more simple to implement than using group policys. In my opinion troubleshooting is also simplified in this approach.
Group policys are great for an organisation-based segmentation of your network, but in my opinon unnecessary if only the size of your broadcast-domain is your aim.
In general 550 is great number for single network; splitting will make it more healthy. Separating your network into VLANs will reduce number of broadcasts clients receive, but not necessarily will reduce broadcasts on AP itself.
