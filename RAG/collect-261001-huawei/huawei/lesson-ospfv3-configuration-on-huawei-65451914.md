---
id: collect-261001-huawei/huawei/lesson-ospfv3-configuration-on-huawei-65451914
title: "lesson-ospfv3-configuration-on-huawei-65451914"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-ospfv3-configuration-on-huawei-65451914.md
source_anchor: ""
source_lines: [1, 33]
sha256: d41e62ba9e740c162563d47f3592fd4f92b65a0a5da8b1057b6d71025ea30beb
---

# lesson-ospfv3-configuration-on-huawei-65451914

OSPFv3 is the IPv6 capable version of OSPFv2. In this Huawei OSPFv3 Configuration lesson, we will configure OSPFv3 on Huawei routers. For our OSPFv3 Configuration Example, we will use the below OSPFv3 Network Topology:
Let’s start our Huawei OSPFv3 Configuration on these three Huawei Routers.
You can download this configuration on Huawei eNSP Labs Page.
Firstly, we will enable IPv6 and we will create OSPFv3 Process globally on Router 1. To enable IPv6 glovbally, we will use “ipv6” command and to enable OSPFv3, we will use “ospfv3 1” command. Here, “1” is the OSPFv3 Process Number.
system-view
[Huawei-Router-1] ipv6
[Huawei-Router-1] ospfv3 1
After that, we will set OSPF Router ID that makes our router unique in the OSPF network.
[Huawei-Router-1-ospfv3-1] router-id 10.10.10.10
Then, we will enable IPv6 on the interfaces that we would like to run OSPFv3 and we will enable IPv6 auto addressesing (Link Local Addresses). We will also give IPv6 addresses manually. These IPv6 addresses will be like below:
Router 1
——–
Ge0/0/0 : FC00:0:0:1000::1/64
Ge0/0/1 : FC00:0:0:1001::1/64
Router 2
——–
Ge0/0/0 : FC00:0:0:2000::1/64
Ge0/0/1 : FC00:0:0:2001::1/64
Router 3
——–
Ge0/0/0 : FC00:0:0:3000::1/64
Ge0/0/1 : FC00:0:0:3001::1/64
[Huawei-Router-1] interface GigabitEthernet 0/0/0
[Huawei-Router-1-GigabitEthernet0/0/0] ipv6 enable
[Huawei-Router-1-GigabitEthernet0/0/0] ipv6 address auto link-local
[Huawei-Router-1-GigabitEthernet0/0/0] ipv6 address FC00:0:0:1000::1/64
After IPv6 configurations, we will add our interface to the OSPFv3 Area 0. Here, we will use area 0 as 0.0.0.0.
[Huawei-Router-1- GigabitEthernet0/0/0] ospfv3 1 area 0.0.0.0
[Huawei-Router-1- GigabitEthernet0/0/0] quit
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
