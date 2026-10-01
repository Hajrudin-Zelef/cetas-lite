---
id: collect-261001-huawei/huawei/lesson-huawei-router-interface-configuration-a36a7882
title: "lesson-huawei-router-interface-configuration-a36a7882"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/lesson-huawei-router-interface-configuration-a36a7882.md
source_anchor: ""
source_lines: [1, 37]
sha256: 89cf4c0dd3de207b2b14937ffe0b13648744853f5b44cd07ba50c80b5baeeb55
---

# lesson-huawei-router-interface-configuration-a36a7882

Huawei Router Interface Configuration is one of the first configuration that Junior Network Engineers learn on Huawei Routers. Here, we will show how to configure Huawei Router Interfaces with given IP addresses.
You can test youself with Huawei HCIA Questions Page.
For our example, we will use the below simple topology.
You can download this configuration on Huawei eNSP Labs Page.
As you can see above, there are two types interfaces. These are :
In this configuration example, we have 1 loopback and 1 physical interface per router. Physical interface configuration will be done for the communication of two router. Loopback interface configuration is done for other special activities.Here, we will also configure it only to show you that it is not different than a physical interface configuration.
Let’s start the configuration with Router A.
In Router A, firstly we will con figure Gigabit Ethernet 0/0/0 interface IP Address with its Subnet Mask. After this, we will open the port. Because by default router ports are shutdown. We will use “undo shutdown” command to open the port.
<Huawei-RouterA> system-view
[Huawei-RouterA] interface GigabitEthernet 0/0/0
[Huawei-RouterA-GigabitEthernet0/0/0] ip address 192.168.0.1  255.255.255.0
[Huawei-RouterA-GigabitEthernet0/0/0] undo shutdown
[Huawei-RouterA-GigabitEthernet0/0/0] quit
The second interface configuration on Router A will be Loopback 0 configuration. Again, we will configure IP Address and Subnet Mask. We can configure Subnet Mask with long versions like (255.255.255.255) or short versions (32).
[Huawei-RouterA] interface Loopback 0
[Huawei-RouterA-Loopback0] ip address 10.10.10.10 32
[Huawei-RouterA-Loopback0] quit
Now, it is time to configure Router B. We will do the same configuration steps for Router B too. Only the IP Addresses will change.
< Huawei-RouterB> system-view
[Huawei-RouterB] interface GigabitEthernet 0/0/0
[Huawei-RouterB-GigabitEthernet0/0/0] ip address 192.168.0.2  255.255.255.0
[Huawei-RouterB-GigabitEthernet0/0/0] undo shutdown
[Huawei-RouterB-GigabitEthernet0/0/0] quit
[Huawei-RouterB] interface Loopback 0
[Huawei-RouterB-Loopback0] ip address 20.20.20.20 32
[Huawei-RouterB-Loopback0] quit
After the configuraton of both Routers, we can check the connection with pinging Router B from Router A and vice versa.
[Huawei-RouterA] ping 192.168.0.2
[Huawei-RouterB] ping 192.168.0.1
To ping Loopback interfaces, routing is need on the routers. Because a router can only access to directly connected networks without routing. For routing, static routing or routing protocol is used. We will explain the configuration of static routing and routing protocols in the other lessons.
After Huawei interface configuration, you can focus on other Huawei HCIA configurations like Huawei NAT, DHCP, OSPF configurations. With more practice on Huawei CLI, you will be more familiar on Huawe commands. Because these commands are both improtant for your job and for HCIA Exam.
Cisco vs Juniper vs Huawei vs Nokia Configuration Commands
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
