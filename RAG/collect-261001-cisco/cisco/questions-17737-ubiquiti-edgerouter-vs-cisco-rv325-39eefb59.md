---
id: collect-261001-cisco/cisco/questions-17737-ubiquiti-edgerouter-vs-cisco-rv325-39eefb59
title: "questions-17737-ubiquiti-edgerouter-vs-cisco-rv325-39eefb59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-17737-ubiquiti-edgerouter-vs-cisco-rv325-39eefb59.md
source_anchor: ""
source_lines: [1, 15]
sha256: eae9db2c8e71d33b14b18b1bf1799af08d893fe1a0eba76e82e608def1d97fc2
---

# questions-17737-ubiquiti-edgerouter-vs-cisco-rv325-39eefb59

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm building a bit more advanced network at home since I'm moving parts of my work to my residence and felt the need for better stuff. I have several servers, virtualized and physical. Ubiquiti Wireless AP's and Zyxel managed switches. What I need is a better router.
So far I've been looking in products at this class
Both systems seems to have all the VPN/IPSec technologies that one may need for coming years. I like the looks and abilities of EdgeOS but the capabilities of the Cisco Router are equally compentent by specification.
The Cisco has double wan ports with builtin failover to a 4G usb stick, it seems as if this is doable in the EdgeRouters as well but requires some extra work and the 4G option can only be obtained using a 4G router modem.
Does anyone have any of these devices and can give me some pros / cons?
I do not have the 8 port edgerouter, but I have used quite a few of the ERL and POE versions. I feel that they are very stable and consistent, and a great product. The Cisco "RV" series are from the Linksys "business class" hardware, and I have never felt that this line was a good product from a performance or stability POV.
Personally, I'd go with the edgerouter. The CLI commands are pretty easy once you get started and the Ubiquiti community is way better for knowledge than Linksys support techs.
Looking at the specs and knowing EdgeOS I would say EdgeMax would be more flexible choice if you do not mind using CLI. Personally I would go with EdgeMax for flexibility and Cisco for peace of mind when having less flexible needs.
I have both, from my experience the Edge Router has more of a learning curve, however it is a faster and more powerful ROUTER, don't use this for any switching though it will take a performance hit. The Cisco RV325 however isn't great for big networks but it is easy and quick to set up. This device as a router is slower than the EdgeRouter.
