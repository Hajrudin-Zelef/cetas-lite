---
id: collect-261001-huawei/huawei/lesson-ripv2-configuration-on-huawei-ensp-2-f41bd32b
title: "lesson-ripv2-configuration-on-huawei-ensp-2-f41bd32b"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-ripv2-configuration-on-huawei-ensp-2-f41bd32b.md
source_anchor: ""
source_lines: [1, 51]
sha256: 5f1a15fda81702ef28bcb254df4070136672dd9380bfa9b052be45df438ae175
---

# lesson-ripv2-configuration-on-huawei-ensp-2-f41bd32b

RIPv2 is an old Routing Protocol but it is the first Routing Protocol that all Network Engineers meet first. In this lesson, we will configure RIPv2 on Huawei Routers.
For our Huawei RIPv2 Configuration Example, we will use the below topology consist of three routers. Here, indirectly connected networks will be learned via RIPv2.
We will follow the following steps for our Huawei RIPv2 Configuration Example:
1) Interface configurations of the routers.
2) RIPv2 Configuration on the routers.
3) RIPv2 Configuration Verification
Table of Contents
We will configure the Router interfaces as above IP addresses on the topology.
Huawei Router 1
< Huawei-Router-1> system-view
[Huawei-Router-1] interface GigabitEthernet 0/0/0
[Huawei-Router-1- GigabitEthernet0/0/0] ip address 10.0.0.1 ip address.1 24
[Huawei-Router-1- GigabitEthernet0/0/0] undo shutdown
[Huawei-Router-1- GigabitEthernet0/0/0] quit
[Huawei-Router-1] interface GigabitEthernet 0/0/1
[Huawei-Router-1- GigabitEthernet0/0/1] ip address 20.0.0.1 ip address.1 24
[Huawei-Router-1- GigabitEthernet0/0/1] undo shutdown
[Huawei-Router-1- GigabitEthernet0/0/1] quit
Huawei Router 2
< Huawei-Router-2> system-view
[Huawei-Router-2] interface GigabitEthernet 0/0/0
[Huawei-Router-2- GigabitEthernet0/0/0] ip address 10.0.0.2 ip address.1 24
[Huawei-Router-2- GigabitEthernet0/0/0] undo shutdown
[Huawei-Router-2- GigabitEthernet0/0/0] quit
[Huawei-Router-2] interface GigabitEthernet 0/0/1
[Huawei-Router-2- GigabitEthernet0/0/1] ip address 30.0.0.1 ip address.1 24
[Huawei-Router-2- GigabitEthernet0/0/1] undo shutdown
[Huawei-Router-2- GigabitEthernet0/0/1] quit
Huawei Router 3
< Huawei-Router-3> system-view
[Huawei-Router-3] interface GigabitEthernet 0/0/0
[Huawei-Router-3- GigabitEthernet0/0/0] ip address 20.0.0.2 ip address.1 24
[Huawei-Router-3 GigabitEthernet0/0/0] undo shutdown
[Huawei-Router-3- GigabitEthernet0/0/0] quit
[Huawei-Router-3] interface GigabitEthernet 0/0/1
[Huawei-Router-3- GigabitEthernet0/0/1] ip address 30.0.0.2 ip address.1 24
[Huawei-Router-3- GigabitEthernet0/0/1] undo shutdown
[Huawei-Router-3- GigabitEthernet0/0/1] quit
On each router, we will configure RIPv2 with the RIP process number “1”, we will set the RIP version as “version 2” and then, we will add the directly connected networks of each router to the RIP network. Latly, we will save our configuration on each Huawei router.
[Huawei-Router-1] rip 1
[Huawei-Router-1-rip-1] version 2
[Huawei-Router-1-rip-1] network 10.0.0.0
[Huawei-Router-1-rip-1] network 20.0.0.0
[Huawei-Router-1-rip-1] quit
[Huawei-Router-1] quit
<Huawei-Router-1> save
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
