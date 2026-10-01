---
id: collect-261001-huawei/huawei/lesson-huawei-static-routing-and-load-balancing-1dbea12d
title: "lesson-huawei-static-routing-and-load-balancing-1dbea12d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-huawei-static-routing-and-load-balancing-1dbea12d.md
source_anchor: ""
source_lines: [1, 16]
sha256: 69ba63027ebed83def077648a1e0e3ec929e84f296cdb97d4e4b419d14d70db0
---

# lesson-huawei-static-routing-and-load-balancing-1dbea12d

Table of Contents
Static Route is a manual simple route configuration. In Huawei Routers, Static Routing is similar to the other platforms like Cisco Static Routing, Nokia Static Routing etc.
The concept is simple:
You can download this configuration on Huawei eNSP Labs Page.
Here, we will show Huawei Static Route Configuration on the below topology.
On Router 1, we will do one of the below Huawei Static Routing Configurations. These three configuration is same. It is up to your configuration behaviour.
<Huawei-Router1> system-view
[Huawei-Router1] ip route-static 10.10.30.0 255.255.255.0 172.16.0.2
OR
<Huawei-Router1> system-view
[Huawei-Router1] ip route-static 10.10.30.0 255.255.255.0 GigabitEthernet 1/1/1
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
