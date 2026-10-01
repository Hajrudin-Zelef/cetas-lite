---
id: collect-261001-huawei/huawei/lesson-ripng-configuration-on-huawei-hcna-a3db61cd
title: "lesson-ripng-configuration-on-huawei-hcna-a3db61cd"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-ripng-configuration-on-huawei-hcna-a3db61cd.md
source_anchor: ""
source_lines: [1, 76]
sha256: 2c3f04299edcc337a7775b306bac906e80915b793a3ec5d33440cba1425ca3f3
---

# lesson-ripng-configuration-on-huawei-hcna-a3db61cd

In this lesson RIPng Configuration Huawei Example, we will focus on the configuration of the basic IPv6 Routing Protocol, RIPng. We will learn RIPng Configuration on Huawei routers. In another lesson, we will learn also OSPFv3 Configuration on Huawei routers.
For our RIPng Configuration Example, we will use the below RIPng Network Topology:
You can download this configuration on Huawei eNSP Labs Page.
Firstly, let’s enable IPv6 and create RIPng with Process ID “1” globally on Router 1.
<Huawei-Router-3> system-view
[Huawei-Router-1] ipv6
[Huawei-Router-1] ripng 1
Then, we will enable IPv6 on the interfaces that we would like to run RIPng and we will set the IPv6 address. We will also enable automatically generated link-local addresses.
Here, we will use the below IPv6 addresses for our interface configurations:
Router 1
————-
Ge0/0/0 : FC00:0:0:1000::1/64
Ge0/0/1 : FC00:0:0:1001::1/64
Router 2
————-
Ge0/0/0 : FC00:0:0:2000::1/64
Ge0/0/1 : FC00:0:0:2001::1/64
Router 3
————-
Ge0/0/0 : FC00:0:0:3000::1/64
Ge0/0/1 : FC00:0:0:3001::1/64
[Huawei-Router-1] interface GigabitEthernet 0/0/0
[Huawei-Router-1- GigabitEthernet0/0/0] ipv6 enable
[Huawei-Router-1- GigabitEthernet0/0/0] ipv6 address FC00:0:0:1000::1/24
[Huawei-Router-1- GigabitEthernet0/0/0] ipv6 address auto link-local
After that, we will configure RIPng under the interface.
[Huawei-Router-1- GigabitEthernet0/0/0] ripng 1 enable
[Huawei-Router-1- GigabitEthernet0/0/0] quit
We will also do the same configuration on the other interface that will run RIPng.
[Huawei-Router-1] interface GigabitEthernet 0/0/1
[Huawei-Router-1- GigabitEthernet0/0/1] ipv6 enable
[Huawei-Router-1- GigabitEthernet0/0/1] ipv6 address FC00:0:0:1001::1/24
[Huawei-Router-1- GigabitEthernet0/0/1] ipv6 address auto link-local
[Huawei-Router-1- GigabitEthernet0/0/1] ripng 1 enable
[Huawei-Router-1- GigabitEthernet0/0/1] quit
Lastly, we will save the configuration on Router 1.
[Huawei-Router-1] save
We will configure the other routers for RIPng like Router 1.
<Huawei-Router-3> system-view
[Huawei-Router-2] ipv6
[Huawei-Router-2] ripng 1
[Huawei-Router-2] interface GigabitEthernet 0/0/0
[Huawei-Router-2- GigabitEthernet0/0/0] ipv6 enable
[Huawei-Router-2- GigabitEthernet0/0/0] ipv6 address FC00:0:0:2000::1/24
[Huawei-Router-2- GigabitEthernet0/0/0] ipv6 address auto link-local
[Huawei-Router-2- GigabitEthernet0/0/0] ripng 1 enable
[Huawei-Router-2- GigabitEthernet0/0/0] quit
[Huawei-Router-2] interface GigabitEthernet 0/0/1
[Huawei-Router-2- GigabitEthernet0/0/1] ipv6 enable
[Huawei-Router-2- GigabitEthernet0/0/1] ipv6 address FC00:0:0:2001::1/24
[Huawei-Router-2- GigabitEthernet0/0/1] ipv6 address auto link-local
[Huawei-Router-2- GigabitEthernet0/0/1] ripng 1 enable
[Huawei-Router-2- GigabitEthernet0/0/1] quit
[Huawei-Router-2] save
<Huawei-Router-3> system-view
[Huawei-Router-3] ipv6
[Huawei-Router-3] ripng 1
[Huawei-Router-3] interface GigabitEthernet 0/0/0
[Huawei-Router-3- GigabitEthernet0/0/0] ipv6 enable
[Huawei-Router-3- GigabitEthernet0/0/0] ipv6 address FC00:0:0:3000::1/24
[Huawei-Router-3- GigabitEthernet0/0/0] ipv6 address auto link-local
[Huawei-Router-3- GigabitEthernet0/0/0] ripng 1 enable
[Huawei-Router-3- GigabitEthernet0/0/0] quit
[Huawei-Router-3] interface GigabitEthernet 0/0/1
[Huawei-Router-3- GigabitEthernet0/0/1] ipv6 enable
[Huawei-Router-3- GigabitEthernet0/0/1] ipv6 address FC00:0:0:3001::1/24
[Huawei-Router-3- GigabitEthernet0/0/1] ipv6 address auto link-local
[Huawei-Router-3- GigabitEthernet0/0/1] ripng 1 enable
[Huawei-Router-3- GigabitEthernet0/0/1] quit
[Huawei-Router-3] save
Now let’s check our RIPng Process with “display ripng 1” command.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
