---
id: collect-261001-cisco/cisco/classification-and-marking-in-qos-ipcisco
title: "classification-and-marking-in-qos-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-cisco/classification-and-marking-in-qos-ipcisco.md
source_anchor: ""
source_lines: [1, 58]
sha256: 40b8fe694758fa413ed4697c711d1d461fd6056e7ed5021339ce55a5b020787a
---

# classification-and-marking-in-qos-ipcisco

To use **Quality of Service** for a traffic, firstly traffic need to be identified. With this identification, traffic types are classified and then they are marked with an understandable way by the network. This process is basically called “**QoS Classification and Marking**“. Another important mechanisms are Qos **Traffic Policing and QoS Traffic Shaping**. We will talk about this in the next lesson.


The identification mechanism used by **QoS** can be divide into **two** important process. These process are :


Now, let’s go a little deeply on **QoS Classification and Marking**, and talk about these two important QoS Process more.


You can download Various **Cheat Sheets** on Special Pages!


Table of Contents

There are carious of traffic types in a network. These traffic types can be **data, voice, video streaming** etc. **Without QoS**, all these traffic types are behaved **similarly**. But behaving similarly to all the traffic types is not a proper way. Because, different traffic types need different threatments.


For example, voice traffic must be **fast** but security is in the second plane for voice traffic. Beside, pure data, ftp can be slower than a voice traffic.


Because of the fact that different types of traffics need different threatments, first of all we need to categorize the type of our traffic. Identifying and categorizing the type of the traffic is called “**Classification**”. After this process, we know that, we have a voice traffic, or video, or what else.


**QoS Classification Process** can be done by checking the different fileds of a packet. There are fields that shows the traffic types in a packet like IP Precedence, DSCP. Beside, incoming interfaces, source and destination addresses can also be used for Classification.


QoS Classification is done **close to the source**. This is because, early determination of the type and threat as required through the network.


After classification, traffic type determination must be showed also in the packet. To do this, a field in a packet header is changed. This changes explains that, the packet is belong to a specific type of traffic. The name of this process is called “**Marking**”. It is also called “**Data Marking**“ or “**Coloring**”.


**QoS Classification Process** is a must before **Marking Process**. Because, you can not mark something about its characteristic without knowing what is it.


**Data Marking** can be done in different levels with different field changes. With this changes, the traffic is quickly recognized anywhere in the network.


In **Layer 2**, in Ethernet Header, **Class of Service** field is used for Data Marking (Data Coloring).

In **Layer 2.5**, in MPLS Header, **Type of Service (Experimental)** field is used for Data Marking (Data Coloring).

In **Layer 3**, in IP Header, Type of Service IP Precedence and **DSCP** fields are used for Data Marking (Data Coloring).

In **Upper Layers**, **NBAR and Deep Packet Inspection** are used for Data Marking (Data Coloring).

In **Cisco CCNA** and **CCNP ENCOR Certification Courses**, Traffic Classification and MArking is an important QoS Lesson.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
