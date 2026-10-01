---
id: collect-261001-huawei/huawei/lesson-gvrp-configuration-on-huawei-4b639967
title: "lesson-gvrp-configuration-on-huawei-4b639967"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-gvrp-configuration-on-huawei-4b639967.md
source_anchor: ""
source_lines: [1, 72]
sha256: 0c5b55b1b58c354747ee9d463e180503745184805d2bc84068d275e4fa1bef53
---

# lesson-gvrp-configuration-on-huawei-4b639967

GVRP is the standard based counterpart of Cisco VTP. In this configuration example, we will focus on how to configure GVRP on Huawei switches. WE will learn Huawei GVRP Configuration.
For our Huawei GVRP Configuration example, we will use the topology below.
On important reminding. Do not think that there can be a loop on this topology. STP mechanisms prohibit this.
Let’s start our GVRP Configuration on each device.
You can download thic configuration on Huawei eNSP Labs Page.
Table of Contents
Firstly, we will enable GVRP Globally on the switch.
< Huawei-Switch1 > system-view
[Huawei-Switch1] gvrp
Secondly, we will configure the first interface as trunk and we will allow all the VLANs to pass through this interfaces.
[Huawei-Switch1] interface GigabitEthernet 1/1/1
[Huawei-Switch1-GigabitEthernet1/1/1] port link type trunk
[Huawei-Switch1-GigabitEthernet1/1/1] port trunk allow-pass vlan all
Then, we will enable GVRP on the interface and we will configure its registration type. There are three types of registration as we talked about before. “Normal”, “Fixed” and “Forbidden”.
Here we will use “fixed” type.
[Huawei-Switch1-GigabitEthernet1/1/1] gvrp
[Huawei-Switch1-GigabitEthernet1/1/1] gvrp registration fixed
[Huawei-Switch1-GigabitEthernet1/1/1] quit
We will configure the second interface similar.
[Huawei-Switch1] interface GigabitEthernet 1/1/2
[Huawei-Switch1-GigabitEthernet1/1/2] port link type trunk
[Huawei-Switch1-GigabitEthernet1/1/2] port trunk allow-pass vlan all
[Huawei-Switch1-GigabitEthernet1/1/2] gvrp
[Huawei-Switch1-GigabitEthernet1/1/2] gvrp registration fixed
On Switch 2, we will do the similar configuration of Switch 1.
< Huawei-Switch2 > system-view
[Huawei-Switch2] gvrp
[Huawei-Switch2] interface GigabitEthernet 1/1/1
[Huawei-Switch2-GigabitEthernet1/1/1] port link type trunk
[Huawei-Switch2-GigabitEthernet1/1/1] port trunk allow-pass vlan all
[Huawei-Switch2-GigabitEthernet1/1/1] gvrp
[Huawei-Switch2-GigabitEthernet1/1/1] gvrp registration fixed
[Huawei-Switch2-GigabitEthernet1/1/1] quit
[Huawei-Switch2] interface GigabitEthernet 1/1/2
[Huawei-Switch2-GigabitEthernet1/1/2] port link type trunk
[Huawei-Switch2-GigabitEthernet1/1/2] port trunk allow-pass vlan all
[Huawei-Switch2-GigabitEthernet1/1/2] gvrp
[Huawei-Switch2-GigabitEthernet1/1/2] gvrp registration fixed
On Switch 3, we will do the similar configuration of Switch 1 and Switch 2.
< Huawei-Switch2 > system-view
[Huawei-Switch3] gvrp
[Huawei-Switch3] interface GigabitEthernet 1/1/1
[Huawei-Switch3-GigabitEthernet1/1/1] port link type trunk
[Huawei-Switch3-GigabitEthernet1/1/1] port trunk allow-pass vlan all
[Huawei-Switch3-GigabitEthernet1/1/1] gvrp
[Huawei-Switch3-GigabitEthernet1/1/1] gvrp registration fixed
[Huawei-Switch3-GigabitEthernet1/1/1] quit
[Huawei-Switch3] interface GigabitEthernet 1/1/2
[Huawei-Switch3-GigabitEthernet1/1/2] port link type trunk
[Huawei-Switch3-GigabitEthernet1/1/2] port trunk allow-pass vlan all
[Huawei-Switch3-GigabitEthernet1/1/2] gvrp
[Huawei-Switch3-GigabitEthernet1/1/2] gvrp registration fixed
[Huawei-Switch3-GigabitEthernet1/1/2] quit
To verify our GVRP configuration on Huawei Switches, we will use “display gvrp status” and “display gvrp statistics” commands.Here, wem will show the results of switch 1. The other switches will also give the similar results.
[Huawei-Switch1] display gvrp status
GVRP is enabled
[Huawei-Switch1] display gvrp statistics
GVRP statistics on port GigabitEthernet1/1/1
GVRP status : Enabled
GVRP registrations failed : 0
GVRP last PDU origin : 0000-0000-0000
GVRP registration type : Fixed
GVRP statistics on port GigabitEthernet1/1/2
GVRP status : Enabled
GVRP registrations failed : 0
GVRP last PDU origin : 0000-0000-0000
GVRP registration type : Fixed
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
