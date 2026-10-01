---
id: collect-261001-cisco/cisco/cisco-bgp-route-reflector-config-on-gns3-ipcisco
title: "cisco-bgp-route-reflector-config-on-gns3-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-bgp-route-reflector-config-on-gns3-ipcisco.md
source_anchor: ""
source_lines: [1, 72]
sha256: a277b15da4ae3f0cf1717b3b1c58c35d60eb26c5db5d0d94ccf9e15d5089e1a6
---

# cisco-bgp-route-reflector-config-on-gns3-ipcisco

**BGP Route Reflectors** are very important in a **BGP network**. With this mechanism, we **bypass Full Mesh iBGP** configuration requirement. So, route reflectors **reduce bandwidth and CPU usage** in a network. How can we configure a **Router reflector**? In this **Cisco BGP Route Reflector Configuration** lesson, you will learn **Cisco BGP Route  Reflector Configuration with GNS3**.


For our **BGP Route-Reflector** configuration example, we will use the below **GNS3 BGP topology**.



In this **GNS3 BGP Configuration example**, we will do **three** main configurations. These are:


Now, let’s start to configure our network according to the given details.


You can also view **Basic Cisco BGP Configuration Example on Packet Tracer**


Table of Contents

For **Cisco BGP Route Reflector Configuration**, we will follow the below step by step BGP Route Reflector configuration guide. After this lesson, you will be ready to configure Router Reflectors (RR) on Cisco routers.


We will start our **Cisco** **BGP** **Router Reflector configuration example** with Router 1. Router 1 will be in a different Autonomous System than other routers. On router 1, firstly we will create BGP process with the **Autonomous System Number 200**. Then, we will announce the networks 3.3.3.3 and 4.4.4.4. Lastly, we will configure EBGP peering with the remote BGP device.


```
R1(config)# 
```
**router bgp 200**                                                     //Creating AS
R1(config-router)# **network 3.3.3.3 mask 255.255.255.255**                        //Announcing networks
R1(config-router)# **network 4.4.4.4 mask 255.255.255.255**
R1(config-router)# **neighbor 10.0.0.1 remote-as 100**                             //Creating EBGP peering


Router 2 will be in **ASN 100**. This is our **Autonomous System** that we will **configure BGP router reflector**. On Router 2, firstly, we will create BGP process with **ASN 100**. After that we will announce the networks. Then, we will configure **EBGP peering** and **IBGP peering**.


```
R2(config)# 
```
**router bgp 100**                                                     //Creating AS
R2(config-router)# **network 8.8.8.8 mask 255.255.255.255**                        //Announcing networks
R2(config-router)# **network 11.11.11.11 mask 255.255.255.255**
R2(config-router)# **network 20.0.0.0 mask 255.255.255.0**
R2(config-router)# **network 30.0.0.0 mask 255.255.255.0**
R2(config-router)# **neighbor 10.0.0.2 remote-as 200**                             //Creating EBGP peering
R2(config-router)# **neighbor 20.0.0.2 remote-as 100**                             //Creating IBGP peering



As an IGP, we will use OSPF in our Autonomous System 100. So, we will configure OSPF on Router 2.


```
R2(config)# 
```
**router ospf 1**                                                             //Creating OSPF Process
R2(config-router)# **network 10.0.0.0 0.0.0.255 area 0**                                  //Adding networks
R2(config-router)# **network 20.0.0.0 0.0.0.255 area 0**


**Router 3** will be our **BGP Router Reflector**. To configure Router 3 as router reflector we will use “**route-reflector-client**” keyword. Firstly, we will create BGP process, we will do the configuration of the IBGP peers. In the last step, we will use “route-reflector-client” keyword with neighbor command towards each IBGP peer.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
