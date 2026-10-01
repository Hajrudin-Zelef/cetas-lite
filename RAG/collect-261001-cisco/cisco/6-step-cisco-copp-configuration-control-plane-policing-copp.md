---
id: collect-261001-cisco/cisco/6-step-cisco-copp-configuration-control-plane-policing-copp
title: "6-step-cisco-copp-configuration-control-plane-policing-copp"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/6-step-cisco-copp-configuration-control-plane-policing-copp.md
source_anchor: ""
source_lines: [1, 210]
sha256: 6e97353272b72c7985f95fdb5e4ff3ee52566494898d77a4181243c9dc84e40e
---

# 6-step-cisco-copp-configuration-control-plane-policing-copp

**Control Plane Policing (CoPP)** is one of the mechanisms used for **Control Plane defence** against the **Network Attacks** that is done to the **Control Plane**. Here, we will focus on what is **Control Plane Policing** and **How to use CoPP**. We will also learn **How to Configure CoPP**.



Table of Contents

As you know we have **two main planes** in the network equipment like routers and switches. These are “**Data Plane**” and “**Control Plane**”. **Data Plane** is used to forward data and the **Control Plane** is used to learn how to send this data to the neighbors, in other words, it learns about the protocols.


There are various attacks types towards these two systems. Especially, attacks towards the control plane are too much. So, how can we protect our control plane against these network attacks?


To protect control plane against network attacks, there are two way. The first way is using the security mechanism of the protocols used in control plane. The other way is using policies. In other words, using **Control Plane Policing**. Here, we will learn what is **Control Plane Policing (CoPP).**


**Control Plane Policing (CoPP)** is the mechanism that allows us to **policy the coming traffic** to the Control Plane. Here, CoPP defines **traffic classification, queue mapping** and **queue shaping** for control plane packets. This **protects** **Control Plane** from malicious network attacks like DoS Attacks.


**Control Plane Policing (CoPP)** uses a Control Plane configuration through **Quality of Service (QoS)** module of **CoS (Class of Service).** In other words, a dedicated configuration is used through **QoS CLI (MQC)** for **CoPP** that filters and limits the traffic coming to the Control Plane. This policy is independent policy that can be used on the device for filtering or any other policing activities. By protecting control plane, **Control Plane Policing** ensures routing stability, reachability and packet delivery.


**You can DOWNLOAD Cisco Packet Tracer Configuration Examples!**


There are four main processes in **CoPP.** These CoPP Processes are given below:

In **Classification**, traffic flow is classified according to the defined traffic classes.

In **Queuing Mapping**, different packets types are sent to the different CPU queues. Each of these packets has different scheduling priorities

In **Scheduling**, one of the queue is selected with a scheduling algorithm and processed. Weighted Round Robin (WRR) scheduling is used in CoPP.

In **Queue Shaping**, the limit of the CPU queues are determined and with this process CPU do not have an excessive laod.



**To configure CoPP,** there are six basic steps. Then, how can we configure **Control Plane Policing?** Let’s learn each of these steps with an example.



First of all we should enable **MLS QoS.** To enable MLS QoS globally, we will use “**mls qos**” command under global configuration mode.

Router(config)# **mls qos**




The second step is defining Access List. Here, we will define an Extended Access List as below:

Router(config)# **ip access-list extended** *access-list-name*


 

*After creating **Extended Access List**, we will enter the **Access List** lines with “**permit**” and “**deny**” lines. According to our need, we can use protocol numbers, tos, time-ranges etc.*


Router(config-ext-nacl)# {**permit** | **deny**} *protocol* **source** *source-wildcard* **destination** *destination-wildcard* [**precedence** *precedence*] [**tos** *tos*] [**established**] [**log** | **log-input**] [**time-range** *time-range-name*] [**fragments**]




After defining Access List, the next step is defining **Packet Classification**. We will identify the traffic belong to a specific class with “**match**” command. To define traffic class, we will use “**class-map**” keyword.


Router(config)# **class-map** *traffic-class-name*

Router(config-cmap)# **match** {**ip precedence**} |{**ip dscp**} | *access-group*




The next step is defining a **Service Policy Map** with the command “**policy-map**”. Beside this, in this step, we will **associate class** with the **Service Policy** with the command “**class**” under policy-map configuration. And we will **associate the actions** to the **Service Policy** map with “**police**” command like below:


Router(config)# **policy-map** service-policy-name

Router(config-pmap)# **class** traffic-class-name



Router(config-pmap-c)# **police** {bits-per-second [normal-burst-bytes] [maximum-burst-bytes] [pir peak-rate-bps]} | [conform-action action] [exceed-action action] [violate-action action]




To enter Control Plane Config mode, we will use “**control-plane**” command.


Router(config)# **control-plane**




Lastly, we will **apply the QoS service policy** to the control plane with “**service-policy**” command under control plane config mode.


Router(config-cp)# **service-policy input *service-policy-name***




To verify Control Plane Policing, we can use different commands. These commands are given below:


“**show policy-map control-plane**” command is used **to display monitoring statistics** about the policy.


Router# **show policy-map control-plane**

Control Plane Interface

Service policy CoPP-normal

Hardware Counters:

class-map: CoPP-normal (match-all)

Match: access-group 111

police :

96000 bps 3000 limit 3000 extended limit

Earl in slot 1 :

0 bytes

5 minute offered rate 0 bps

aggregate-forwarded 0 bytes action: transmit

exceeded 0 bytes action: drop

aggregate-forward 0 bps exceed 0 bps

Earl in slot 3 :

0 bytes

5 minute offered rate 0 bps

aggregate-forwarded 0 bytes action: transmit

exceeded 0 bytes action: drop

aggregate-forward 0 bps exceed 0 bps

Software Counters:

Class-map: CoPP-normal (match-all) 0 packets, 0 bytes

5 minute offered rate 0 bps, drop rate 0 bps

Match: access-group 111

police:

96000 bps, 1250 limit, 1250 extended limit

conformed 0 packets, 0 bytes; action: transmit

exceeded 0 packets, 0 bytes; action: drop

conformed 0 bps, exceed 0 bps, violate 0 bps



“**show platform qos ip**” command is used **to display the hardware counters** that shows the dropped and forwarded packets by policy.


Router# **show platform qos ip**

QoS Summary [IP]: (* – shared aggregates, Mod – switch module)

Int Mod Dir Class-map DSCP Agg Trust Fl AgForward-By AgPoliced-By

Id Id

——————————————————————————-

CPP 5 In CoPP-normal 0 1 dscp 0 303485 58615147

CPP 9 In CoPP-normal 0 4 dscp 0 0 0


“**show access-lists coppacl-bgp**” command is used **to display CoPP ACL Information**.

Router# **show access-lists coppacl-bgp**

Extended IP access list coppacl-bgp

10 permit tcp host 5.5.5.5 host 10.1.1.1 eq bgp (2 matches)

20 permit tcp host 5.5.5.5 eq bgp host 10.1.1.1



This is basically how to configure **Control Plane Policing (CoPP)** on Cisco devices.


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
