---
id: collect-261001-cisco/cisco/questions-20473-how-to-view-bgp-community-attached-to-prefix-in-cisco-asr-router-6b1fa6d7
title: "questions-20473-how-to-view-bgp-community-attached-to-prefix-in-cisco-asr-router-6b1fa6d7"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-20473-how-to-view-bgp-community-attached-to-prefix-in-cisco-asr-router-6b1fa6d7.md
source_anchor: ""
source_lines: [1, 17]
sha256: d3899ea6a10655fc9e8e742fb3e2d44b6586a83fd6e606e2a9109acb5aab9661
---

# questions-20473-how-to-view-bgp-community-attached-to-prefix-in-cisco-asr-router-6b1fa6d7

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Cisco ASR routers use IOS XE software. One can see the received IPv4 unicast prefixes from BGP neighbor before any inbound policies with sh bgp ipv4 unic neigh <neighbor-address> received routes command. Output of this command is following:
Status codes: s suppressed, d damped, h history, * valid, > best
i - internal, r RIB-failure, S stale
Origin codes: i - IGP, e - EGP, ? - incomplete
Network Next Hop Metric LocPrf Weight Path
* 10.10.10.0/24 192.168.4.9 0 65001 65005 i
Is it possible to see communities attached to 10.10.10.0/24 prefix before any inbound policies?
I believe this is not possible as @jwbensley said. Though it's also unlikely that cisco.com explicitly says that this functionality isn't implemented.
When you're viewing attributes such as community you're looking at BGP table. Route is getting into BGP table after policy has been applied. So if you change attributes with policy, you can only see modified attributes.
If you turn on inbound soft-reconfiguration on the neighbor in question, then you will see another entry in the list with the pre-policy communities on it.
If soft reconfiguration inbound is configured, the path was received but dropped by inbound policy, or was accepted and modified. In either event, the received-only value is a copy of the original, unmodified path.
