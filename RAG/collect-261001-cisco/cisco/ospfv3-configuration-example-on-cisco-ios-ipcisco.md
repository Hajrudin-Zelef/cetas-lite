---
id: collect-261001-cisco/cisco/ospfv3-configuration-example-on-cisco-ios-ipcisco
title: "ospfv3-configuration-example-on-cisco-ios-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/ospfv3-configuration-example-on-cisco-ios-ipcisco.md
source_anchor: ""
source_lines: [1, 30]
sha256: 08fceab05b2b52ad4ae0b890362b42c471040a91ba743ad2e764d50b6951fa1a
---

# ospfv3-configuration-example-on-cisco-ios-ipcisco

OSPFv3 is the IPv6 cabaple version of OSPF. The previous version of OSPF was OSPFv2. OSPFv2 is for IPv4.

Here, we will configure OSPFv3 for the below topology. As you can see, there are three areas and six routers in this topology.

For OSPFv3 Configuration, we will follow the below steps one by one:

1. Enabling Global IPv6 Routing
2. Enabling Interfaces for IPv6
3. IPv6 Address Configuration
4. Configuring OSPFv3 Process
5. Router ID Configuration
6. Adding Interfaces to OSPFv3 Process and Areas
7. Clear ipv6 ospf process
8. Verification

Now, let’s go to each COnfiguration steps one by one.

Firstly, we will enable global IPv6 Routing on each Router with “ipv6 unicast-routing” command.

In the third step, we will do the IPv6 address configuration/strong>. We will use “ipv6 address -ipv6address- eui-64“ commands. The Link-local addresses required for neighbourship, will produce Link-local addresses with this command.

Now, it is time to configure OSPFv3 process under each router. We will use the process number 1 here. Beside, we will configure a router ID in IPv4 format. Because for both OSPv2 and OSPv3, Router ID is in IPv4 format.

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

## Leave a Reply
