---
id: collect-260926-mikrotik/mikrotik/questions-2583-bridging-vlan-trunks-on-routeros-5a0ae333
title: "questions-2583-bridging-vlan-trunks-on-routeros-5a0ae333"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "ethernet", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-2583-bridging-vlan-trunks-on-routeros-5a0ae333.md
source_anchor: ""
source_lines: [1, 27]
sha256: e38826c373402cfc6a1ad428c0963ac2b2816c1645ad8a9c4de941f66c85ac68
---

# questions-2583-bridging-vlan-trunks-on-routeros-5a0ae333

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Edit:
While Mikrotik is well known for their consumer-grade wireless
routers, this question is about their rack mounted,
carrier-grade, aggregation service routers.
I need to send a VLAN trunk over an EoIP tunnel through two RouterOS/Mikrotik routers. I have thus bridged an ethernet interface on each side of the link with the respective EoIP interfaces.
Seeing how different VLANs and trunks are handled on RouterOS makes me wonder how I should go about this.
The frames will be tagged on egress from Catalyst switches on either side of the routers. Can one simply skip any VLAN configuration on RouterOS in this case? Will it bridge the frames with 802.1q tags intact?
I assume that such configuration will cause the IP addresses of the routers to be accessible from any VLAN on the trunk, if any addresses are configured on the bridged ethernet interfaces.
Any other downside to this configuration? Is there a better way?
To be sure, create VLAN sub-interfaces on the physical interfaces and bridge the sub-interfaces. Any IP addresses are then configured on the respective bridges. VLANs are isolated from each other on layer 2 like on any switch.
Treat the EoIP tunnel interface like a physical interface for this exercise.
Oddly enough, the procedure is very similar to how VLAN trunks are configured on Juniper MX series routers :)
In the scenario above, where plain trunks are in their entirety passed through a RouterOS device without adding or removing tags or doing any selective bridging, you can get away with not configuring individual vlans at all.
By just bridging the ethernet ports, RouterOS will forward any tagged frames with the tags intact.
This type of configuration is illustrated in the following example in the docs:
AS soon as you need to do anything much with the VLANs, it's time to refer to the answer by @Aziraphale and create all of the individual vlan bridges and associated configurationse.
EoIP or Ether over IP tunnel is a tunnel protocol designed by
Mikrotik development team which allows network administrators to
easily connect private LANs located in different locations separated
by cities or countries. As long as the Mikrotik routers can ping each
other, we can create the EoIP tunnel among them.
