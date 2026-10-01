---
id: collect-261001-general-networking/general-networking/questions-48887-trying-to-configure-stp-on-a-small-network-and-have-a-few-questi-c15cbff8
title: "questions-48887-trying-to-configure-stp-on-a-small-network-and-have-a-few-questi-c15cbff8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-48887-trying-to-configure-stp-on-a-small-network-and-have-a-few-questi-c15cbff8.md
source_anchor: ""
source_lines: [1, 13]
sha256: ffe1ab4a0af7d350df4cfe9d1ba73f95f15580623d9e981fd0127e5ab90674e1
---

# questions-48887-trying-to-configure-stp-on-a-small-network-and-have-a-few-questi-c15cbff8

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to configure the switches on a network to use STP or RSTP but I'm confused about how to do this. The network consists of 3 switches, all on the default subnet, so a fairly simple setup. However, all three switches are connected directly to a Meraki firewall. My concerns are as follows:
It seems like in this setup there would be no STP root, unless the meraki can act as root.
It seems like this network has no redundancy. Does this mean I should forgo the use of STP entirely? Should I add redundancy?
And finally would it be considered better practice to connect only one switch to the firewall and connect the other two switches to that one, making it the root?
Spanning tree is to prevent layer-2 forwarding loops. It only works if you have multiple bridges (switches are bridges) connected. It works by forwarding frames toward the root bridge, thus preventing loops.
If you connect each switch (bridge) to a separate layer-3 interface on your firewall/router, then you have no need for STP, but if you connect your switches together, then you need to use STP, and it will select a root bridge.
The optimal location of the root bridge depends on how the traffic should flow. If the traffic is primarily leaving your network through the firewall/router, then you want the bridge connected to the firewall/router to be the root bridge. If your traffic is primarily from host to host in the LAN, then you want the root bridge to be central to the LAN.
