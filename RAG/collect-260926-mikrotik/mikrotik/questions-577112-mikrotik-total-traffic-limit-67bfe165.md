---
id: collect-260926-mikrotik/mikrotik/questions-577112-mikrotik-total-traffic-limit-67bfe165
title: "questions-577112-mikrotik-total-traffic-limit-67bfe165"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/qos/questions-577112-mikrotik-total-traffic-limit-67bfe165.md
source_anchor: ""
source_lines: [1, 9]
sha256: 8888e7aef59aa75b08972031376d958851167a9d64149194c406e2ce4c4974e9
---

# questions-577112-mikrotik-total-traffic-limit-67bfe165

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am running PPPoe in mikrotik.... I can use speed rate limit per user but i want to limit of total traffic per month.... for example i want to limit a user for 10GB traffic at 1mbps speed after 10GB used client can not connect using pppoe please help me to do this
Also i want to track user traffic per login.... so that i can get details that how much traffic passed in a single login
That is exactly what Radius is designed to do. Authentication, Authorization, and Accounting. However, proper Radius usage is not trivial, and not really a concept that could be explained in one answer. I can highly recommend freeRADIUS, it is an excellent piece of software, and many people are using it to do what you are requesting.
