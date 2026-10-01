---
id: collect-261001-cisco/cisco/bgp-configuration-packet-tracer-bgp-config-example-ipcisco
title: "bgp-configuration-packet-tracer-bgp-config-example-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/bgp-configuration-packet-tracer-bgp-config-example-ipcisco.md
source_anchor: ""
source_lines: [1, 119]
sha256: 50720e3beef2ddc2b23bf70516d5207ac793e313c121407567599cba089fd356
---

# bgp-configuration-packet-tracer-bgp-config-example-ipcisco

Table of Contents

To understand **BGP (Border Gateway Protocol)** better, we will make a basic **Packet Tracer BGP Configuration** example. Because of the limited numbers of commands available on **Packet Tracer**, we will practice a very basic configuration for our **BGP Config Example**. Beside this, I will add some additional cofiguration steps that is needed for IBGP but we can not config ure on Packet Tracer.

In the configuration we will use two **AS (Autonomous System)** with 3 routers for each. We will use the **Private AS block (64512 to 65535)** for this configuration, but in internet public AS numbers are used.

DOWNLOAD all the **Packet Tracer** examples with **.pkt** format in **Packet Tracer Labs** section.

The topology that we will be used for our **BGP Config** is below:

__Interfaces IP Addresses__

1.1.1.1

2.2.2.2

10.0.0.1

10.0.0.2

20.0.0.1

20.0.0.2

30.0.0.1

30.0.0.2

40.0.0.1

40.0.0.2

50.0.0.1

50.0.0.2

__RouterA1 Interface Configuration__

RouterA1(config)# **interface loopback 0**
RouterA1(config-if)# **ip address 1.1.1.1 255.255.255.255**
RouterA1(config-if)# **no shutdown**
RouterA1(config-if)# **exit**
RouterA1(config)# **interface gigabitEthernet 0/0**
RouterA1(config-if)# **ip address 10.0.0.1 255.255.255.0**
RouterA1(config-if)# **no shutdown**
RouterA1(config-if)# **exit**
RouterA1(config)# **interface gigabitEthernet 0/1**
RouterA1(config-if)# **ip address 20.0.0.1 255.255.255.0**
RouterA1(config-if)# **no shutdown**
RouterA1(config)# **interface gigabitEthernet 0/2**
RouterA1(config-if)# **ip address 30.0.0.1 255.255.255.0**
RouterA1(config-if)# **no shutdown**

__RouterB1 Interface Configuration__

RouterB1(config)# **interface loopback 0**
RouterB1(config-if)# **ip address 2.2.2.2 255.255.255.255**
RouterB1(config-if)# **no shutdown**
RouterB1(config-if)# **exit**
RouterB1(config)# **interface gigabitEthernet 0/0**
RouterB1(config-if)# **ip address 10.0.0.2 255.255.255.0**
RouterB1(config-if)# **no shutdown**
RouterB1(config-if)# **exit**
RouterB1(config)# **interface gigabitEthernet 0/1**
RouterB1(config-if)# **ip address 40.0.0.1 255.255.255.0**
RouterB1(config-if)# **no shutdown**
RouterB1(config-if)# **exit**
RouterB1(config)# **interface gigabitEthernet 0/2**
RouterB1(config-if)# **ip address 50.0.0.1 255.255.255.0**
RouterB1(config-if)# **no shutdown**

__RouterA2 Interface Configuration__

RouterA2(config)# **interface gigabitEthernet 0/1**
RouterA2(config-if)# **ip address 20.0.0.2 255.255.255.0**
RouterA2(config-if)# **no shutdown**

__RouterA3 Interface Configuration__

RouterA3(config)# **interface gigabitEthernet 0/2**
RouterA3(config-if)# **ip address 30.0.0.2 255.255.255.0**
RouterA3(config-if)# **no shutdown**

__RouterB2 Interface Configuration__

RouterB2(config)# **interface gigabitEthernet 0/1**
RouterB2(config-if)# **ip address 40.0.0.2 255.255.255.0**
RouterB2(config-if)# **no shutdown**

__RouterB3 Interface Configuration__

RouterB3(config)# **interface gigabitEthernet 0/2**
RouterB3(config-if)# **ip address 50.0.0.2 255.255.255.0**
RouterB3(config-if)# **no shutdown**

The exact important point of BGP Config is here. The configuration made in this part, is for the BGP.

As I said before, because of the Packet Tracer ‘s command limit, in the configuration file, th IBGP parts are not configured, but writen here (ibgp neighbourship and route reflector commands).

RouterA1(config)# **router bgp 64600**
RouterA1(config-router)# **neighbor 10.0.0.2 remote-as 64700**
RouterA1(config-router)# **neighbor 20.0.0.2 remote-as 64600**
RouterA1(config-router)# **neighbor 30.0.0.2 remote-as 64600**
RouterA1(config-router)# **neighbor 20.0.0.2 route-reflector-client**
RouterA1(config-router)# **neighbor 30.0.0.2 route-reflector-client**
RouterA1(config-router)# **network 20.0.0.0 mask 255.255.255.0**
RouterA1(config-router)# **network 30.0.0.0 mask 255.255.255.0**

**You can also test yourself with BGP Tests and Questions.**

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
