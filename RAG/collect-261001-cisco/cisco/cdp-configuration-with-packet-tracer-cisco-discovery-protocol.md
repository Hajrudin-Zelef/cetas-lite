---
id: collect-261001-cisco/cisco/cdp-configuration-with-packet-tracer-cisco-discovery-protocol
title: "cdp-configuration-with-packet-tracer-cisco-discovery-protocol"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cdp-configuration-with-packet-tracer-cisco-discovery-protocol.md
source_anchor: ""
source_lines: [1, 183]
sha256: f765facd8dcad13e0452ddad7b4cdbdb4cad0f02a2c943ffb2be58e8e6a7f535
---

# cdp-configuration-with-packet-tracer-cisco-discovery-protocol

As we have talked about in the previous lesson, **CDP** is a Cisco proprietary **Neighbor Discovery Protocol**. In this Cisco Discovery Protocol lesson, we will discuss **how to configure CDP** in Cisco IOS, we will learn **CDP Cisco Configuration**. LLDP is a **standard-based** neighbour discovery protocol. In another lesson we will also **configure LLDP**.


**You can DOWNLOAD the Packet Tracer example with .pkt format HERE.**

**For all Packet Tracer Examples and Files, you can check Packet Tracer Labs Page.**


For our example, the below topology is used.


We will not talk about the configuration of the central switch, but we will talk about the general **CDP configuration commands.**



To **enable CDP** **globally**, use **“cdp run”** on the router configuration mode. To disable **Cisco Discovery Protocol,** we can use **“no cdp run”** command.


```
Switch(config)# 
```
**cdp run**
Switch(config)# **no cdp run**

If this protocol is disabled globally and we  can enable it on a specific interface. To **enable CDP** on a specific interface, use **“cdp enable”** command on the interface configuration mode. By **default** CDP is already **enabled**. You can also disable CDP by using **“no cdp enable”** command.


```
Switch(config-if)# 
```
**cdp enable**
Switch(config-if)# **no cdp enable**


There are timers used in **Cisco Discovery Protocol.** The roles of these timers and default values are given below:


To configure cdp **Hello** time and **Hold** time, you can use the below commands. Time is mentioned as seconds.


```
Switch(config)# 
```
**cdp timer 50**
Switch(config)# **cdp holdtime 100**


To **clear** the CDP table, use the **“clear cdp table”** command.


```
Switch# 
```
**clear cdp table**


To **verify CDP**, you can use general CDP verification commands below:

```
Switch# 
```
**show cdp**
Switch# **show cdp interface**
Switch# **show cdp neighbors**
Switch# **show cdp entry**
Switch# **show cdp traffic**

In our topology, for the central switch, **“show cdp neighbours”** and **“show cdp interface”** command outputs are showed below.


```
Switch# 
```
**show cdp neighbors** 
Capability Codes: R - Router, T - Trans Bridge, B - Source Route Bridge
S - Switch, H - Host, I - IGMP, r - Repeater, P - Phone
Device ID Local Intrfce Holdtme Capability Platform Port ID
IP Phone Fas 0/4 124 H P 7960
Switch Fas 0/6 124 3560 Fas 0/1
Switch Fas 0/2 124 S 2960 Fas 0/2
Router Fas 0/5 124 R C2800 Fas 0/0
Switch Fas 0/1 124 S 2950 Fas 0/1


```
Switch# 
```
**show cdp interface**
FastEthernet0/1 is up, line protocol is up
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/2 is up, line protocol is up
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/3 is up, line protocol is up
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/4 is up, line protocol is up
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/5 is up, line protocol is up
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/6 is up, line protocol is up
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/7 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/8 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/9 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/10 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/11 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/12 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/13 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/14 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/15 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/16 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/17 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/18 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/19 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/20 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/21 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/22 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/23 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
FastEthernet0/24 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
GigabitEthernet0/1 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds
GigabitEthernet0/2 is down, line protocol is down
Sending CDP packets every 60 seconds
Holdtime is 180 seconds


In this article, we have talked about the configuration of **Cisco Discovery Protocol (CDP)** basically with Cisco Packet Tracer.


**You can download “Packet Tracer” in Tools section.**


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
