---
id: collect-261001-huawei/huawei/lesson-vlan-routing-with-layer-2-switch-and-router-on-huawei-a71e53db
title: "lesson-vlan-routing-with-layer-2-switch-and-router-on-huawei-a71e53db"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-vlan-routing-with-layer-2-switch-and-router-on-huawei-a71e53db.md
source_anchor: ""
source_lines: [1, 22]
sha256: dca4c00c94885171957d04c2e8cba82dac4f062b0908873a59ce85720b8a7c93
---

# lesson-vlan-routing-with-layer-2-switch-and-router-on-huawei-a71e53db

Huawei VLAN Routing with Layer 2 Switch is done with the help of a Layer 3 device, a router. Mainly, in the router, in layer 3, gateways are created for the VLANs. These gateways are the subinterfaces under the physical interface. For each subinterface, an IP address is assigned.
Our VLAN Routing Configuration example topology is below:
You can download this configuration on Huawei eNSP Labs Page.
On Huawei switch, firstly we will configure the VLAN interfaces. We will set link-types of this interfaces as access and we will set its default vlan.
[Huawei-Switch] interface GigabitEthernet0/0/1
[Huawei-Switch-GigabitEthernet0/0/1] port link-type access
[Huawei-Switch-GigabitEthernet0/0/1] port default vlan 10
[Huawei-Switch-GigabitEthernet0/0/1] quit
[Huawei-Switch] interface GigabitEthernet0/0/2
[Huawei-Switch-GigabitEthernet0/0/2] port link-type access
[Huawei-Switch-GigabitEthernet0/0/2] port default vlan 20
[Huawei-Switch-GigabitEthernet0/0/2] quit
[Huawei-Switch] interface GigabitEthernet0/0/3
[Huawei-Switch-GigabitEthernet0/0/3] port link-type access
[Huawei-Switch-GigabitEthernet0/0/3] port default vlan 30
[Huawei-Switch-GigabitEthernet0/0/3] quit
We will configure the router face interface of the switch as “trunk” and we will configure it to allow all vlans. If you want to allow only specific VLANs, it allows only these VLANs.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
