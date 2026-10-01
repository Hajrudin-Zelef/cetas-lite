---
id: collect-261001-huawei/huawei/lesson-dhcp-configuration-on-huawei-routers-7ffa10c1
title: "lesson-dhcp-configuration-on-huawei-routers-7ffa10c1"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/lesson-dhcp-configuration-on-huawei-routers-7ffa10c1.md
source_anchor: ""
source_lines: [1, 33]
sha256: 4b60efb8d373055ed0c7376cc00faff9aa5fb0f7199bab97c02c0d347eb58ac2
---

# lesson-dhcp-configuration-on-huawei-routers-7ffa10c1

Table of Contents
In this DHCP example we will configure, Huawei DHCP, a Huawei Router as a DHCP Server. After the configuration, a host device, here a PC, will receive its IP information from DHCP Server Huawei Router.
For our DHCP configuration example, we will use the topology below:
Let’s firstly talk about the DHCP Configuration steps on the Huawei Router. We will do the below steps one by one:
1.Enabling DHCP
2.DHCP IP Pool Configuration
3.DNS Configuration
4.Interface DHCP Configuration
5.DHCP Verification
6.DHCP Client Configuration
You can download this configuration on Huawei eNSP Labs Page.
By default, DHCP is not enabled on Huawei Routers. So, the first step of configuring a Huawei Router as DHCP Server is enabling DHCP on the router.
<Huawei-Router>system-view
[Huawei-Router] dhcp enable
Here, we will configure the IP Pool on Huawei Router. Our IP allocations will be done from this IP Pool.
Two different IP Pool can be defined :
In Global IP Pool configuration, firstly, we will configure the IP Pool range. After this we will configure a gateway address. This gateway address is the address of Huawei Router for local network devices. Lastly, we will give a lease time. This lease time is the time limit that DHCP Client can use this IP address.
[Huawei-Router] ip pool OurPool
[Huawei-Router-ip-pool-OurPool] network 192.168.0.0 mask 24
[Huawei-Router-ip-pool-OurPool] gateway-list 192.168.0.1
[Huawei-Router-ip-pool-OurPool] lease day 2 day hour 18
Other DHCP Articles:
DHCP Overview
DHCP IP Allocation Operation
Cisco DHCP Configuration With Packet Tracer
Huawei DHCP Configuration
DNS information is also provided from the DHCP Server. To get this information from DHCP Server, we will configure Huawei Router, with DNS commands.
Here, firstly, we will configure the DNS IP address and then we will exclude this address from the IP Pool.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
