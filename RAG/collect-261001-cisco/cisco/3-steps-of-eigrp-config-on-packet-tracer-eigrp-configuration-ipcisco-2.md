---
id: collect-261001-cisco/cisco/3-steps-of-eigrp-config-on-packet-tracer-eigrp-configuration-ipcisco-2
title: "3-steps-of-eigrp-config-on-packet-tracer-eigrp-configuration-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/3-steps-of-eigrp-config-on-packet-tracer-eigrp-configuration-ipcisco.md
source_anchor: ""
source_lines: [216, 287]
sha256: 0d02b822ac087f8d0ed0d5e000ac9aae65bdf2e5ac1020f5bf86888b4d7dc971
---

# 3-steps-of-eigrp-config-on-packet-tracer-eigrp-configuration-ipcisco

We can see the whole Topology Table of Router0 and Router1 with **show ip eigrp topology**  command.


```
Router0# 
```
**show ip eigrp topology** 
IP-EIGRP Topology Table for AS 100/ID(192.168.1.2)
Codes: P - Passive, A - Active, U - Update, Q - Query, R - Reply,
r - Reply status
P 10.0.0.0/24, 1 successors, FD is 28160
via Connected, FastEthernet4/0
P 20.0.0.0/24, 1 successors, FD is 28160
via Connected, FastEthernet5/0
P 30.0.0.0/24, 1 successors, FD is 30720
via 10.0.0.2 (30720/28160), FastEthernet4/0
P 40.0.0.0/8, 1 successors, FD is 30720
via 20.0.0.2 (30720/28160), FastEthernet5/0
P 192.168.1.0/24, 1 successors, FD is 28160
via Connected, FastEthernet0/0
P 192.168.2.0/24, 1 successors, FD is 30720
via 10.0.0.2 (30720/28160), FastEthernet4/0
P 192.168.3.0/24, 1 successors, FD is 30720
via 20.0.0.2 (30720/28160), FastEthernet5/0
P 192.168.4.0/24, 1 successors, FD is 33280
via 10.0.0.2 (33280/30720), FastEthernet4/0
 
Router1# **show ip eigrp topology**
IP-EIGRP Topology Table for AS 100/ID(192.168.2.2)
Codes: P - Passive, A - Active, U - Update, Q - Query, R - Reply,
r - Reply status
P 10.0.0.0/24, 1 successors, FD is 28160
via Connected, FastEthernet4/0
P 20.0.0.0/24, 1 successors, FD is 30720
via 10.0.0.1 (30720/28160), FastEthernet4/0
P 30.0.0.0/24, 1 successors, FD is 28160
via Connected, FastEthernet5/0
P 40.0.0.0/8, 1 successors, FD is 30720
via 30.0.0.2 (30720/28160), FastEthernet5/0
P 192.168.1.0/24, 1 successors, FD is 30720
via 10.0.0.1 (30720/28160), FastEthernet4/0
P 192.168.2.0/24, 1 successors, FD is 28160
via Connected, FastEthernet0/0
P 192.168.3.0/24, 1 successors, FD is 33280
via 10.0.0.1 (33280/30720), FastEthernet4/0
P 192.168.4.0/24, 1 successors, FD is 30720
via 30.0.0.2 (30720/28160), FastEthernet5/0

We can also see the **EIGRP** traffic statistics with “**show ip eigrp traffic**” command.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

How can i configure OSPF and explaination about the ospf

Hi Capelo. You can follow OSPF Configuration on Cisco Packet Tracer lesson, to learn how to configure OSPF on Cisco routers.

When you use the network command without a wildcard mask, EIGRP assumes classful networks.

So this means:

192.168.1.0 → treated as 192.168.1.0 /24 ✔️

10.0.0.0 → treated as 10.0.0.0 /8 ⚠️

20.0.0.0 → treated as 20.0.0.0 /8 ⚠️
