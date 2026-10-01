---
id: collect-261001-cisco/cisco/etherchannel-cisco-pagp-configuration-on-gns3-lag-with-pagp
title: "etherchannel-cisco-pagp-configuration-on-gns3-lag-with-pagp"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/etherchannel-cisco-pagp-configuration-on-gns3-lag-with-pagp.md
source_anchor: ""
source_lines: [1, 72]
sha256: 00ded2dcb2652479d540b44607edc41fcb40a76a3ed9530d03ba6d1867699858
---

# etherchannel-cisco-pagp-configuration-on-gns3-lag-with-pagp

**Link Aggregation (LAG)** is one of the most important mechanisms used for **bandwidth increase** and **link redundany** in network operation. Especially in **backbone networks** of Internet Service Providers, link aggregation is used for this purposes. In small networks, LAGs are also used for the same reason. In this **Cisco EtherChannel Configuration** lesson, we will focus on **Etherchannel Cisco PAgP Configuration**. Here, you will learn **How to Configure Etherchannel with PAgP on GNS3.** We will also remember some of the details of **PAgP (Port Aggregation Protocol)**. By the way, another network protool used for this purpuse is **LACP (Link Aggregation Control Protocol)**. We will talk **Cisco LACP Configuration** in anohter lesson.


For our **GNS3 PAgP Etherchannel Configuration**, we will use the below topology consist of two switches. We will configure the link between these switches as a **PAgP Link Aggregation** (LAG).



Here, we will aggregate, three ports of each switch and create one logical channel. To do this, there are two alternatives. One of them is **LACP (Link Aggregation Control Protocol)** and the other is PAgP (Port Aggregation Protocol). In this example, we will use PAgP with GNS3.


Before going to PAgP EtherChannel Setup Configuration steps, let’s remember the modes of **PAgP**. PAgP has **two modes**. These are given below:



**Desirable Mode** sends requests to form link aggregation **actively**. **Auto Mode** does not send requests but **waits for any requests**. So, if both ends are configured as **auto**, then link aggregation **cannot be formed**. But one of them is in desirable mode, then link aggregation is **formed**. Below, you can find the different cases of these PAgP modes.


Now, let’s configure each Cisco switch for **PAgP (Port Aggregation Protocol)**.


You can check also **another PAgP Example**


Table of Contents

To configure **Cisco PAgP** on Cisco swithes, we will start with the first switch, switch_1. On first switch, we will configure each port as the member of our etherchannel. And to do this we will add “**channel-group**” command with group number and the required mode. In the first switch we will use **PAgP Auto** mode.


```
Switch_1(config)# 
```
**interface gi0/0**
Switch_1(config-if)# **channel-group 1 mode auto** 
Switch_1(config)# **interface gi0/1**
Switch_1(config-if)# **channel-group 1 mode auto**
Switch_1(config)# **interface gi0/2**
Switch_1(config-if)# **channel-group 1 mode auto**


To download the GNS3 files of this lab, visit **GNS3 Lab Files Page**


On the second switch, switch_2, we will configure the ports similarly for **Cisco PAgP Configuration**. This time, our mode will be **PAgP** **desirable**.


```
Switch_2(config)# 
```
**interface gi0/0**
Switch_2(config-if)# **channel-group 1 mode desirable**
Switch_2(config)# **interface gi0/1**
Switch_2(config-if)# **channel-group 1 mode desirable**
Switch_2(config)# **interface gi0/2**
Switch_2(config-if)# **channel-group 1 mode desirable**


With this configuration, all the interfaces are bundled together. To verify our Etherchannel, we will use “**show etherchannel 1 summary**” command like below. As you can see below, our logical channel is there and it is configured with **Cisco PAgP**.



We can also use “**show etherchannel 1 detail**” to view the detailed information about **Etherchannel Cisco PAgP Configuration**.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
