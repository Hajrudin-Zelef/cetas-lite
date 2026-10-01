---
id: collect-261001-general-networking/general-networking/questions-483-static-redistribution-into-eigrp-as-internal-failure-b53e74df
title: "questions-483-static-redistribution-into-eigrp-as-internal-failure-b53e74df"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-483-static-redistribution-into-eigrp-as-internal-failure-b53e74df.md
source_anchor: ""
source_lines: [1, 16]
sha256: 04565c8b09f731b0fcd877d037b5f6f98e24497b2681e4cc05e06b82be96c11e
---

# questions-483-static-redistribution-into-eigrp-as-internal-failure-b53e74df

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
9
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am redistributing a static route into EIGRP on a 6500. I had read here if you enter a network statement matching a static route that it would be redistributed as an internal route by EIGRP. However, after entering the configuration, neighbors indicate the route as external.
Does the redistribute static preempt and cause the route to be marked as external?
You only need one or the other; here you've used both. It seems that the redistribution command takes precedence over the network command so the route is appearing as static routes. Remove 'redistribute static' and the route should appear as internal.
Are you looking at two different routes? Maybe a /24 and a /23? In general a protocol can only redistribute what is already in the routing table. If you make a static route that matches the prefix of a connected interface the connected interface will be in the routing table and not the static. That would force the static to not placed into EIGRP (since it isn't in the routing table).
Now if you use "redistribute connected" that could change things, but I would still assume the connected interface would take precedence.
Just came across your question.
Network command actually takes precedence over Redistribution.
The reason why the static route appears to be an External EIGRP route is that in EIGRP, Network command can only advertise Static route that is pointing to EXIT Interface! Static route that is point to Next_Hop_Address would never be advertised via Network command.
Redistribution, however, doesn't have such restriction.
My experience of redistribution of statics with eigrp tells me that using the network statement will not inject the network, only the redistribute static works, however unless for also specify a metric be it default or on the redistribute statement it will not work. So use redistribution for statics with a metric and it will work
