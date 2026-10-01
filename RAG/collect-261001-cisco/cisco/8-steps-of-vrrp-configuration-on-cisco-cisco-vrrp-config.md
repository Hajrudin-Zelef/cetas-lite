---
id: collect-261001-cisco/cisco/8-steps-of-vrrp-configuration-on-cisco-cisco-vrrp-config
title: "8-steps-of-vrrp-configuration-on-cisco-cisco-vrrp-config"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-cisco/8-steps-of-vrrp-configuration-on-cisco-cisco-vrrp-config.md
source_anchor: ""
source_lines: [1, 128]
sha256: 21c4363d75d41554f5c71c5931f610931935453776a36de268ec50dd12a10443
---

# 8-steps-of-vrrp-configuration-on-cisco-cisco-vrrp-config

Table of Contents

In this **VRRP** Cisco Configuration Example, we will work on **Cisco IOS** and configure **VRRP (Virtual Router Redundancy Protocol)** on **Cisco Routers.** **Configuring VRRP Protocol** is like configuring Cisco Proprietary protocol **HSRP**. There are only small command differences. Here, you will learn **How to Configure VRRP on Cisco Routers** detailly. If you would like to check the **comparison of VRRP and HSRP**, you can check the related lesson.


For our **VRRP Cisco Configuration Example**, we will use the below topology:


Now, let’s configure **VRRP Protocol** on **Cisco Routers** and learn VRRP better.


The first step of our **VRRP Cisco Configuration** is the IP address configuration on interfaces.


```
Switch A (config)# 
```
**int fa0/1**
Switch A (config-if)# **no switchport**
Switch A (config-if)# **ip address 172.16.0.1 255.255.255.0**
Switch A (config-if)# **no shutdown**

```
Switch B (config)# 
```
**int fa0/1**
Switch B (config-if)# **no switchport**
Switch B (config-if)# **ip address 172.16.0.2 255.255.255.0**
Switch B (config-if)# **no shutdown**

**VRRP** process is created with the **VRRP Group ID** and the **Virtual IP Address**.Here, our **VRRP Group ID** is 1 and our **Virtual IP** address is 172.16.0.5.


```
Switch A (config-if)# 
```
**vrrp 1 ip 172.16.0.5**

```
Switch B (config-if)# 
```
**vrrp 1 ip 172.16.0.5**

Like all descriptions, in **VRRP**, we can use description to give our VRRP group a memorable description.


```
Switch A (config-if)# 
```
**vrrp 1 description Our-vrrp**

```
Switch B (config-if)# 
```
**vrrp 1 description Our-vrrp**

For **Master/Backup** selection, **VRRP** priorities are configured. To manuplate VRRP selection and determine a Master manually, **VRRP Priority** values are important. The default one is 100 and the highest one is elected as VRRP Master.


```
Switch A (config-if)# 
```
**vrrp 1 priority 200**

```
Switch B (config-if)# 
```
**vrrp 1 priority 100**

It is not required to configure **VRRP Priority** for switch B, because it is the default value. Here, we write this command to show you only.


To configure **VRRP authentication,** we will use the below command on switches. Our password is CISCO.


```
Switch A (config-if)# 
```
**vrrp 1 authentication CISCO** 

```
Switch B (config-if)# 
```
**vrrp 1 authentication CISCO**

As you can see, **VRRP** and **HSRP** configuration is very similar. Only HSRP “standby” word is leaving us. Instead, for VRRP, “**vrrp**” is coming.


The **Preempt** mechanims is enabled on VRRP by default. To disable it, we can use “**no vrrp preempt**” command.


To verify the **VRRP** configuration we can use **“show vrrp**” and “**show vrrp brief all**” VRRP verification commands.


```
Switch A # 
```
**show vrrp**
Ethernet1/0 - Group 1
State is Master
Virtual IP address is 172.16.0.5
Virtual MAC address is 0000.5e00.0101
Advertisement interval is 1.000 sec
Preemption is enabled
min delay is 0.000 sec
Priority is 200
Track object 1 state Down decrement 15
Master Router is 172.16.0.1(local), priority is 200
Master Advertisement interval is 1.000 sec
Master Down interval is 3.531 sec
To disable VRRP, we can use “**no vrrp 1 enable**” command.


This is basically, **VRRP Cisco Configuration**. You can alşo check **HSRP** and **GLBP** Configuration Lessons to learn more on **FHRPs (First Hop Redundancy Protocols)**.


**All First Hop Redundancy Lessons:**


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
