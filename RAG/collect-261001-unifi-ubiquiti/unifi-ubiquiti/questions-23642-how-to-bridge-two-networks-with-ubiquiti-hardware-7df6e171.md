---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-23642-how-to-bridge-two-networks-with-ubiquiti-hardware-7df6e171
title: "questions-23642-how-to-bridge-two-networks-with-ubiquiti-hardware-7df6e171"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-23642-how-to-bridge-two-networks-with-ubiquiti-hardware-7df6e171.md
source_anchor: ""
source_lines: [1, 10]
sha256: 2b590fbff113f1354ef4408470a0a0bfd0283323dd1067a3be4199c608dc091c
---

# questions-23642-how-to-bridge-two-networks-with-ubiquiti-hardware-7df6e171

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I want to connect two networks using a wireless bridge. My networks are in different subnets and each have their own router (and DHCP server, and independent internet connection).
I tried a basic layer 2 bridge, but I can't connect to devices on the other network. I can ping both wireless devices (Ubiquiti Nanostation) from network 1.
You need to provide a route, otherwise you have two networks with no reason to talk across the link, as they are different subnets, with different default gateways, and the nanostation at network 2 is on network 1. If you plug a single computer into the LAN port on that nanostation, does it connect to network 1? it should, if the link is working. But without a route, no traffic will flow.
So you need a router at each end that is capable of managing a connection to WAN and a connection to the other network. Not difficult, but not consumer-grade stuff.
