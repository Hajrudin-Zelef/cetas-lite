---
id: collect-260926-mikrotik/mikrotik/questions-1722877-bandwidth-control-with-queue-in-mikrotik-f7d4aaf6
title: "questions-1722877-bandwidth-control-with-queue-in-mikrotik-f7d4aaf6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/questions-1722877-bandwidth-control-with-queue-in-mikrotik-f7d4aaf6.md
source_anchor: ""
source_lines: [1, 3]
sha256: a7b7243a5b2033d93b3fc89d535d6f42dd0108f1467139b8e6a67dc558044a1b
---

# questions-1722877-bandwidth-control-with-queue-in-mikrotik-f7d4aaf6

I set a bandwidth limit in Mikrotik (simple queues) for each IP, and now I need to set a general bandwidth limit for each of my networks. My question is how to set a 10M limit for a range like 192.168.102.0/24 and a 3M limit for each user on this network like 192.168.102.1 and 192.168.102.2 and etc. So that even if all network users start downloading, they will not be able to download more than their own network ceiling of 10M? In other words, suppose like next picture:
I set 10M bandwidth for a group of users and now I want to consider 3M bandwidth for per users in this group but the total consumption of this group should not exceed10M
Update: Thanks, but this doesn't solve my main problem. The main problem is not exceeding the total bandwidth of users. Let me explain my question again with the help of mathematics We have 10 users and we want to give each of them a maximum speed of 3 megabytes. So, normally we need 30 megabytes, but we don't want their total speed to exceed 10 megabytes. The main problem is limiting the total bandwidth of users to 10 megabytes
