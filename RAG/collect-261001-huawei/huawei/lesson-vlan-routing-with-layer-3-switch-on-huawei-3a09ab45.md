---
id: collect-261001-huawei/huawei/lesson-vlan-routing-with-layer-3-switch-on-huawei-3a09ab45
title: "lesson-vlan-routing-with-layer-3-switch-on-huawei-3a09ab45"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-vlan-routing-with-layer-3-switch-on-huawei-3a09ab45.md
source_anchor: ""
source_lines: [1, 18]
sha256: 8b82ad5f63749877ea673d0935a73f21a1a78c96a367cecff329c31c713f8e4f
---

# lesson-vlan-routing-with-layer-3-switch-on-huawei-3a09ab45

In this lesson we will learn Layer 3 VLAN Routing with VLAN Routing Huawei Configuration Example. We do not need any other Layer 3 device for VLAN Routing with Layer 3 Switch. Because Layer 3 switch can achieve this routing facility. Here, instead of subinterface, “VLANIF” is created for each VLAN.
Our example topology is below:
Let’s start the configuration.
You can download this configuration on Huawei eNSP Labs Page.
In our VLAN Routing Huawei Configuration Example, firstly we will configure, each interfaces of the L3 Switch. We will set these interfaces as access and we will give their default VLAN.
[Huawei-Switch] interface GigabitEthernet0/0/1
[Huawei-Switch-GigabitEthernet0/0/1] port link-type access
[Huawei-Switch-GigabitEthernet0/0/1] port default vlan 10
[Huawei-Switch-GigabitEthernet0/0/1] quit
[Huawei-Switch] interface GigabitEthernet0/0/2
[Huawei-Switch-GigabitEthernet0/0/2] port link-type access
[Huawei-Switch-GigabitEthernet0/0/2] port default vlan 20
[Huawei-Switch-GigabitEthernet0/0/2] quit
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
