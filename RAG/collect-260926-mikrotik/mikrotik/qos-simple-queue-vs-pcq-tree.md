---
id: collect-260926-mikrotik/mikrotik/qos-simple-queue-vs-pcq-tree
title: "qos-simple-queue-vs-pcq-tree"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/qos-simple-queue-vs-pcq-tree.md
source_anchor: ""
source_lines: [1, 11]
sha256: 6459cbc7b8c34a496b78fb2579dfc86e2e4fd7f53ff12e3e7944a24a80d82ac5
---

# qos-simple-queue-vs-pcq-tree

Hey, so we have about 700 customers who need to be queued. up until now we have been using simple queues in their respective access point.

i based a new QOS Model on janis megis’s “Qos Best Practices” marking the connection, then the packet then setting up a PCQ tree. but i just read that in v6 there were considerable improvements to simple queues. I would still like to prioritize data like i do with the tree still. (download over upload. higher Package Higher priority)

which is the more effective way to do this? with the simple queues or the PCQ tree?

On another Note: im trying to do traffic prioritization based off of different services (also based off of “Qos Best Practices”) however i cant mark the p2p in the method she described (says i must use layer7 protocol) and for all of the communication services there is no port, or address-list or anything i can see to mark it with.

after having it all marked i need to queue it with different priorities and no limit with higher priority than my upload/download queues (Or on a seperate router that is in front of the primary queueing router)

correct? or is this no longer a good way of prioritizing traffic and it should all be down in the layer 7 protocol (like it told me to do for p2p)
