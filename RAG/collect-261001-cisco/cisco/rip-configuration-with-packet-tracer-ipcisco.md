---
id: collect-261001-cisco/cisco/rip-configuration-with-packet-tracer-ipcisco
title: "rip-configuration-with-packet-tracer-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/rip-configuration-with-packet-tracer-ipcisco.md
source_anchor: ""
source_lines: [1, 62]
sha256: 042bd34b184af8807f279edf560fcdc0f3d49f7516cfec05dbe7612649818834
---

# rip-configuration-with-packet-tracer-ipcisco

In  this **RIP Cisco Configuration** lesson, we will focus on how to configure RIP (Routing Information Protocol) on Cisco Routers. Generally in all  network courses, RIP is used to explain routing protocol configuration basically. This is because of its basic configuration.

For our example, we will use the following basic topology:

With this topology, the IP addresses are configured. We will pass the configuration of IP addresses and continue only the routing protocol, **RIP Cisco Configuration**.

Before the configuration of RIP, let’s check the routing table of RouterA.

```
RouterA# 
```
**show ip route** 
Codes: C - connected, S - static, I - IGRP, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default, U - per-user static route, o - ODR
       P - periodic downloaded static route
Gateway of last resort is not set
     10.0.0.0/30 is subnetted, 1 subnets
C       10.10.10.0 is directly connected, FastEthernet0/0
     20.0.0.0/30 is subnetted, 1 subnets
C       20.20.20.0 is directly connected, FastEthernet0/1
```
RouterA# 
```
**show ip protocols**
```
RouterA# 
```
**show ip int brief** 
Interface              IP-Address      OK? Method Status                Protocol
 
FastEthernet0/0        10.10.10.1      YES manual up                    up
FastEthernet0/1        20.20.20.1      YES manual up                    up
 
Vlan1                  unassigned      YES unset  administratively down down
As you can see, there is only directly connected neighbours in the routing table of RouterA. RouterA do not know any other networks, also RouterD. There is no routing protocol configured on RouterA.


Firstly RouterA is configured like below with **network** commands under **rip process**.


```
RouterA(config)# 
```
**router rip** 
RouterA(config-router)# **version 2**
RouterA(config-router)# **network 10.10.10.0**
RouterA(config-router)# **network 20.20.20.0**
**You can download “Packet Tracer” in Tools section.**

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
