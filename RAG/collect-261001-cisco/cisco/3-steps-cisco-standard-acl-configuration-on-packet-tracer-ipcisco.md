---
id: collect-261001-cisco/cisco/3-steps-cisco-standard-acl-configuration-on-packet-tracer-ipcisco
title: "3-steps-cisco-standard-acl-configuration-on-packet-tracer-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/3-steps-cisco-standard-acl-configuration-on-packet-tracer-ipcisco.md
source_anchor: ""
source_lines: [1, 148]
sha256: ba10e575355e981f03dcf5f25eac3db0aaaece5ae98b30eab28a7ba3802f0546
---

# 3-steps-cisco-standard-acl-configuration-on-packet-tracer-ipcisco

Table of Contents

In this **Cisco Packet Tracer Standard ACL Configuration** example you will learn **Cisco Standard ACL Configuration** and **How to Configure Standart Access-List** in Cisco IOS with Packet Tracer. Let’s start with what is Standard Access List? **Standard Access-Lists** are the simplest one. With Standard Access-List you can control network traffic by permitting or denying packets based only on the **Source IP Address**. The ranges of standard access lists are **1–99** for classic **Standard ACLs** and **1300–1999** for **Expanded Standard ACLs**.



By the way, there are **three types** **Access Contol Lists** in common. These access list types are :


**Extended Access-Lists**, you can check source, destination, specific port and protocols. Lastly, with **Named Access-Lists**, you can use **Names** instead of the numbers used in standard and extended ACLs. It do not have too much difference, but it is different with its named style.


You can **DOWNLOAD** the **Cisco Packet Tracer** example with **.pkt** format at the **End of This Lesson**.


In this lesson, we will focus on **Standart Access-List Configuration** with **Cisco Packet Tracer**. We will focus on the below topology.



Here, with our **Standard Access-List**, we will prohibit PC2 to access the server. But PC0 and PC1 can still access the server.


For our Standard Access-List, we can use the **ACL Number** **1 to 99**. These numbers can be **100 to 199**, if you use extended ACLs.


Let’s start to do **Cisco Standard ACL Configuration**. We will configure the **Standard Access-List** on router .


```
Router # 
```
**configure terminal**
Router (config)# **ip access-list standard 1** 
Router (config-std-nacl)# **permit 10.0.0.2 0.0.0.0**
Router (config-std-nacl)# **permit 10.0.0.3 0.0.0.0**

With this ACL configuration that we have written, we **permit** PC0 and PC1 to **access** the server. At the end of ACLs, there is an “**Implicit Deny**”. These **Implicit Deny**, prohibits the other IP addresses. Because of the fact that we did not, allow PC2’s IP address, it is **automatically denied** and can not access the server.


Here, there is no need to write but to show how to write deny, I will write the deny command also. As I said before, for this scenario, it is not necesary. But, you can write.


```
Router (config-std-nacl)# 
```
**deny 10.0.0.4 0.0.0.0**
Router (config-std-nacl)# **end**
Router # **copy run start**

After creating ACLs, we need to apply this ACL to the **interface**. For **Standard Access-List**, it is better to apply this ACL, close to the destination. So, for this configuration, we will apply our standard acceess list to the fastethernet 0/1 interface of the router.In other words, we will add ACL to the server face of the router.


**You can also DOWNLOAD all the Packet Tracer examples with .pkt format in Packet Tracer Labs**


```
Router (config)# 
```
**interface fastethernet 0/1** 
Router (config-if)# **ip access-group 1 out**
Router (config-if)# **end**
Router # **copy run start**

As you see above, to write a Standard Access-List, firstly we enter the **standard ACL configure mode**, then we write **permit/deny statement**. After that we **write the IP Address** that we would like to effect. Then, we write the **wildcard mask** for that subnet. Here, we only **deny** a specific **IP**, so our wildcard mask will be 0.0.0.0.


Lastly, we **apply** the standard access-list that we write, to the **interface close to the destination**.


Now, it is time to verify. Let’s verify our **Standard ACL Configuration** with **Cisco Packet Tracer**.Our aim was restricting PC2 to access the server. But PC0 and PC1 would still access the server.


Here we will **ping** the server ip address, 20.0.0.5 from each PC.

```
PC0> 
```
**ping 20.0.0.5** 
PC1> **ping 20.0.0.5**
PC2> **ping 20.0.0.5**


Here, the ping from PC0 and PC1 will be successfull. But, ping from PC2 will be unsuccessfull. The fastethernet 0/0 interface of router, willl send a “**destination host unreachable**” message to the PC2.



In this lesson, we have configured **Standard Access-List with Packet Tracer**. For **Extended** and **Named Access-list configurations**, you can check other ACL lessons.



The numbers **1–99** are used for Standard ACLs. This is the traditional range for standard ACL. On the other hand, Cisco later expanded this range to support more ACLs. This **extended standard ACL** range is **1300-1999.**


This is one of the confused questions about **access list placement**. Network engineers generally think that if they place standard access list close to the destination or close to source. The answer of this question is **Close to the Destination.**


A **Standard ACL** **filters** traffic only based on the **Source IP Address**. It does not check destination IP, protocol, or port numbers. If you place a Standard ACL **close to the source**, you might accidentally block traffic that should be allowed to reach other destinations.

Location of exteded ACL is different than Standard ACL. Extended ACLs are placed **Close to the Source**.


Extended ACL filters traffic based on **Source IP address,** **Destination IP address,** **Protocol (TCP, UDP, ICMP etc.),** **Port numbers (HTTP, SSH, FTP, etc.).** If unwanted traffic is blocked **near the source**, it **never travels through the network.** This reduces unnecessary bandwidth usage, reduces router/traffic processing and improves network efficiency.


At the end of this lesson, let’s briefly talk about Wildcard Masks. A **Wildcard Mask** is simply the **inverse of a subnet mask**. Let’s give an example.


Think about, we have a subnet mask 255.255.255.252. To find the wildcard mask of this subnet mask, firstly, we will convert the subnet mask into binary format. Then, to find the **Wildcard Mask**, simply we will change all **1s** to **0s** and all **0s** to **1s**. For 255.255.255.252 subnet mask, wildcard mask is 0.0.0.3.


Subnet Mask               : **255.255.255.252**

Binary                           : **11111111.11111111.11111111.11111100**

Wildcard Mask            : **00000000.00000000.00000000.00000011**

Decimal                        : **0.0.0.3**


Here are some examples:


| **Subnet Mask** | **Wildcard Mask** | 
| 255.255.255.0 | 0.0.0.255 | 
| 255.255.255.252 | 0.0.0.3 | 
| 255.255.255.248 | 0.0.0.7 | 
| 255.255.255.240 | 0.0.0.15 | 


Simply remember this:

**Wildcard Mask = 255.255.255.255 − Subnet Mask**

**Wildcard Mask as the inverse of a Subnet Mask**.


We will use wildcard masks frequently in **ACL**, **OSPF**, **EIGRP** and many other Cisco configurations.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
