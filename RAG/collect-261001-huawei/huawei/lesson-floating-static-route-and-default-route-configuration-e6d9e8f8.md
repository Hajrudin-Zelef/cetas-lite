---
id: collect-261001-huawei/huawei/lesson-floating-static-route-and-default-route-configuration-e6d9e8f8
title: "lesson-floating-static-route-and-default-route-configuration-e6d9e8f8"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-floating-static-route-and-default-route-configuration-e6d9e8f8.md
source_anchor: ""
source_lines: [1, 25]
sha256: bce02793953fa4ed485c9a199f70df50f883fd12efcce0a0bb89372ac5dbd0b6
---

# lesson-floating-static-route-and-default-route-configuration-e6d9e8f8

In this HCIA (HCNA) lesson we will focus on Huawei Floating Static Route and Default Route on Huawe Routers. We will give examples for How to configure Huawei Floating Static Route Configuration and Default Route Configuration. Let’s start.
You can test youself with Huawei HCIA Questions Page.
Table of Contents
Floating Static Route is used for a Backup Static Route. A second route is defined with a high Preference and if the first link fails, then this floating static route is used as a Backup.
For our example, we will use the below topology. And as you can see, the Grey link will be configured as Backup to the destination (Floating Static Route).
You can download this configuration on Huawei eNSP Labs Page.
<Huawei-Router1> system-view
[Huawei-Router1] ip route-static 10.10.30.0 24 GigabitEthernet 1/1/1
[Huawei-Router1] ip route-static 10.10.30.0 24 GigabitEthernet 1/1/2 preference 110
<Huawei-Router1> system-view
[Huawei-Router2] ip route-static 10.10.30.0 24 GigabitEthernet 1/1/1
[Huawei-Router2] ip route-static 10.10.30.0 24 GigabitEthernet 1/1/2 preference 110
It is not necessary that to configure both end with the same Floating Static Route (Backup). Static Route is a Unidirectional configuration, so Backup Static Route is also Unidirectional.
With “display ip routing-table” command, we can see the Primary (Active) Static Route on Routing Table. Here, you can see the Default Preference value (Preference 60).
Default Route, is the route, that if no other destination, the routes are sent. In Huawei Routers, Default Route is also a specific Static Route like other platforms.
Here, we will use the below topology for our Default Route Configuration.
We will write a Default Route to Router 1 and we will say that, if there is no other route, go through Router 2 GigabitEthernet 1/1/1.We can do this with interface name or Next Hop Interface IP.
<Huawei-Router> system-view
[Huawei-Router] ip route-static 0.0.0.0 0.0.0.0 172.16.0.2
“Display ip routing-table” command will show us the default route on Router 1. As you can see, according to this route, always a destination is exist on the Routing Table.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
