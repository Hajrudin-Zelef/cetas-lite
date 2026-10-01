---
id: collect-261001-cisco/cisco/3-steps-of-syslog-configuration-cisco-syslog-config-levels
title: "3-steps-of-syslog-configuration-cisco-syslog-config-levels"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/3-steps-of-syslog-configuration-cisco-syslog-config-levels.md
source_anchor: ""
source_lines: [1, 74]
sha256: ebb63a74c40f5fd628813c25b8547200dbeccfe4e97ed847374fc007567cb7f8
---

# 3-steps-of-syslog-configuration-cisco-syslog-config-levels

Table of Contents

In this **Syslog Configuration Cisco** example, we will learn How to do **Syslog Configuration on Cisco Routers**. Here, we will use the below simple topology consist of a Cisco Router and a **Syslog Server**.

Here, we will go through **3 steps**, these are:


So, let’s start **Syslog Configuration** step by step.


You can also check **Packet Tracer SSH Configuration**


Firsty, we will do the ip configurations of both router and server interfaces. This is a classical IP address configuraration. We will use the below ip addresses:


**Router Fa0/0 -> 10.0.0.1/24**

**Server Fa0 -> 10.0.0.2/24**


On Router we will do the below configuration:

```
Router # 
```
**configure terminal**
Router (config)# **interface fastEthernet 0/0**
Router (config-if)# **ip address 10.0.0.1 255.255.255.0**
Router (config-if)# **no shutdown**
Router (config-if)# **exit**
On Server we will do the below configuration:



You can also learn **Packet Tracer DHCP Configuration**


In this step, we will configure 3 sub steps. These are the main Syslog configuration steps in a router. Here, we will do the below jobs:


On Cisco routers, firstly we will **enable syslog** with “**logging on**” command.

```
Router (config)# 
```
**logging on**

After that we will enter the Log Server IP address. This IP address will be the interface IP address of our Syslog Server. It is 10.0.0.2.

```
Router (config)# 
```
**logging 10.0.0.2**

Lastly, we will set the trap level. As you can see in the first Picture, there are 8 trap levels. These are:


Here, if you select any of these levels, all the other levels lover than it, will be showed in the logs. For example if you select **Critical level (3),** then **Critical (3), Alert(2)** and **Emergency(1)** Levels will be in the logs. For thsi example, we will set tap level as “**debug**”. So, we will see all the logs after this configuration.

```
Router (config)# 
```
**logging trap debugging**

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
