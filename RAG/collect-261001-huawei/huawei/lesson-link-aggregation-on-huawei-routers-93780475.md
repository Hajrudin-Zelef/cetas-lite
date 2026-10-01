---
id: collect-261001-huawei/huawei/lesson-link-aggregation-on-huawei-routers-93780475
title: "lesson-link-aggregation-on-huawei-routers-93780475"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-link-aggregation-on-huawei-routers-93780475.md
source_anchor: ""
source_lines: [1, 66]
sha256: 1f7cbc4dd1d787ed7ab68e57fd3ac5f6c83d975a771fc773c05a9a8919851c64
---

# lesson-link-aggregation-on-huawei-routers-93780475

Link Aggregation is one of the important lessons of networking. This is also used too much in network operations. In this lesson, we will talk about Huawei Link Aggregation Configuration. We will show Huawei Link Aggregation Configuration with two examples. One of these examples will be for Layer 2 and the other will be for Layer 3. You can also check Cisco LACP Configuration Example.
Before the configuration let’s remember some points about the Link Aggregation. As you know Link aggregation provide a logical aggregated single link. This logical link provides :
Test Yourself With Network Practice Questions!
And there are two link aggregation types. There two type link aggregations are :
In Manual mode, load balancing is done with all links. In LACP mode, there are active and backup links and backup links are used for redundancy.
Let’s start with Huawei Layer 2 Link Aggregation. For this example, we will use the below topology consist of two switches and three Layer 2 links.
Let’s start to configure Link Aggregation on Huawei Switches in the example topology.
Here, firstly we will enter the configuration mode with “system-view” command. After that we will create Logical Aggregated Link with “Eth-Trunk” command.
<Huawei-SwitchA> system-view
[Huawei-SwitchA] interface Eth-Trunk 10
[Huawei-SwitchA] interface Eth-Trunk 10
Then, we will enter each member candidate link and bind this links with Aggregated Logical Link with “eth-trunk” link.
[Huawei-SwitchA]interface GigabitEthernet1/1/1
[Huawei-SwitchA-GigabitEthernet1/1/1] eth-trunk 10
[Huawei-SwitchA-GigabitEthernet1/1/1] interface GigabitEthernet1/1/2
[Huawei-SwitchA-GigabitEthernet1/1/2] eth-trunk 10
[Huawei-SwitchA-GigabitEthernet1/1/2] interface GigabitEthernet1/1/3
[Huawei-SwitchA-GigabitEthernet1/1/3] eth-trunk 10
[Huawei-SwitchA-GigabitEthernet1/1/3] quit
Now, we will configure the second Huawei Switch.
<Huawei-SwitchB> system-view
[Huawei-SwitchB] interface Eth-Trunk 10
[Huawei-SwitchB] quit
[Huawei-SwitchB] interface GigabitEthernet1/1/1
[Huawei-SwitchB-GigabitEthernet1/1/1] eth-trunk 10
[Huawei-SwitchB-GigabitEthernet1/1/1] interface GigabitEthernet1/1/2
[Huawei-SwitchB-GigabitEthernet1/1/2] eth-trunk 10
[Huawei-SwitchB-GigabitEthernet1/1/2] interface GigabitEthernet1/1/3
[Huawei-SwitchB-GigabitEthernet1/1/3] eth-trunk 10
[Huawei-SwitchB-GigabitEthernet1/1/3] quit
It is time to verify our Huawei Layer 2 Link Aggregation Configuration. We can verify Aggregated Link with “display interface eth-trunk 10” command.
After first configuration example, now, let’s go to the second example and configure Huawei Layer 3 Link Aggregation.
In our topology, there are two routers that are connected eachother with 3 links. In this example, we will aggragete these layer 3 links and build a single logical layer 3 link.
Firstly, let’s start with Router A.Here, we will start with creating Eth-Trunk 20. We will configure this aggregated interface as layer 3 interface with “undo portswitch” command and then, we will add an IP address under this logical Eth-Trunk . Lastly, we enter the member candidate interfaces and bind them with this logical interface, Eth-Trunk.
<Huawei-SwitchA> system-view
[Huawei-RouterA] interface Eth-Trunk 20
[Huawei- RouterA-Eth-Trunk20] undo portswitch
[Huawei- RouterA-Eth-Trunk20] ip address 10.0.0.1 24
[Huawei- RouterA-Eth-Trunk20] interface GigabitEthernet1/1/1
[Huawei- RouterA-GigabitEthernet1/1/1] eth-trunk 20
[Huawei- RouterA-GigabitEthernet1/1/1] interface GigabitEthernet1/1/2
[Huawei- RouterA-GigabitEthernet1/1/2] eth-trunk 20
[Huawei- RouterA-GigabitEthernet1/1/2] interface GigabitEthernet1/1/3
[Huawei- RouterA-GigabitEthernet1/1/3] eth-trunk 20
[Huawei- RouterA-GigabitEthernet1/1/3] quit
Now, let’s configure Router B.
<Huawei-SwitchB> system-view
[Huawei-RouterB] interface Eth-Trunk 20
[Huawei- RouterB-Eth-Trunk20] undo portswitch
[Huawei- RouterB-Eth-Trunk20] ip address 10.0.0.2 24
[Huawei- RouterB-Eth-Trunk20] interface GigabitEthernet1/1/1
[Huawei- RouterB-GigabitEthernet1/1/1] eth-trunk 20
[Huawei- RouterB-GigabitEthernet1/1/1] interface GigabitEthernet1/1/2
[Huawei- RouterB-GigabitEthernet1/1/2] eth-trunk 20
[Huawei- RouterB-GigabitEthernet1/1/2] interface GigabitEthernet1/1/3
[Huawei- RouterB-GigabitEthernet1/1/3] eth-trunk 20
[Huawei- RouterB-GigabitEthernet1/1/3] quit
The last step is the verification. You can check Link Aggregation with “display interface eth-trunk 20” command.
[Huawei-RouterA] display interface eth-trunk 20
[Huawei-RouterB] display interface eth-trunk 20
You can also view the below Link Aggregation Configuration Lessons for different vendors:
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
