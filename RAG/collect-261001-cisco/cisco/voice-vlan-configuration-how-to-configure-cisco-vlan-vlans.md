---
id: collect-261001-cisco/cisco/voice-vlan-configuration-how-to-configure-cisco-vlan-vlans
title: "voice-vlan-configuration-how-to-configure-cisco-vlan-vlans"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/voice-vlan-configuration-how-to-configure-cisco-vlan-vlans.md
source_anchor: ""
source_lines: [1, 98]
sha256: 65abfe0ec07427550f77fafe5a999d2234ab1ec28b7419f102b17e511bdac329
---

# voice-vlan-configuration-how-to-configure-cisco-vlan-vlans

**IP Phones** are one of the most used network devices in modern offices today. **Voice over IP** **(VoIP)** is also a popular technology which is used in such devices. In an office LAN, beside normal **VLANs**, **Voice VLAN** is also used for these devices. Here, we will talk about these VLANs and **how to configure Voice VLANs** on Cisco switches. In other words we will configure **IP Phone VLAN setup**, we will do **Cisco VLAN configuration**.


**You can also check Packet Tracer VLAN Configuration Example.**


Table of Contents

Think about your desk in the office. Generally, only one network cable comes to the desk in the office. The data connection is done over this cable. What if you are using multiple devices and one of them is an **IP Phone**? You have one cable but you need to access two different services. Here, you need a solution.


For such a topology, to overcome this issue, **multiple ports** on the IP Phone are used beside a specific **Voice VLANs** configuration on the switch By using a separate Voice VLAN beside the Data VLAN, you can use both of these devices over a single link.


Here, a basic topology like below is used. As you can see, in this topology, IP Phone is connected to the switch and then PC is connected to that IP Phone.



The PC sends **untagged data** to IP Phone. Only, **untagged data** is passed over this link. On the other hand, on the link which connects IP Phone to switch, passes both **tagged voice** and **untagged data** traffic.



Now, let’s do a **Voice VLAN Configuration** example. Below example is typically an office desk connection. Here, we use two separate VLANs for data and voice. **Data** is used over **VLAN 20** and **Voice** is used over **VLAN 10**.



Below, you can find the configuration steps one by one.


Firstly, we will create **VLAN 10** and **VLAN 20**. We will also give the names of the VLANs.


```
Switch(config)# 
```
**vlan 10**
Switch(config-vlan)# **name VOICE**
SSwitch(config-vlan)# **vlan 20**
Switch(config-vlan)# **name DATA**

Then, we will go to the interface which **IP Phone** is connected to the switch. We will configure this port as **access port** and **two VLANs** can be used under this port.


```
Switch(config)# 
```
**interface fa0/1**
Switch(config-if)# **switchport mode access**
Switch(config-if)# **switchport access vlan 20**
Switch(config-if)# **switchport access voice vlan 10**

To verify **Cisco Voice VLAN Configuration**, we will use the “**show interfaces fa0/1 switchport**” command.


```
Switch#
```
 **show interfaces fa0/1 switchport**
Name: Fa0/1
Switchport: Enabled
Administrative Mode: static access
Operational Mode: static access
Administrative Trunking Encapsulation: dot1q
Operational Trunking Encapsulation: native
Negotiation of Trunking: Off
**Access Mode VLAN: 20 (VLAN00020)**
Trunking Native Mode VLAN: 1 (default)
**Voice VLAN: 10**
Administrative private-vlan host-association: none
Administrative private-vlan mapping: none

In this lesson, we have configured **Voice VLAN** basically. With this configuration, we will access both data and voice traffic over the same link. This configuration is especially widely used in offices which uses **IP Phones**.


**You can also check Full Packet Tracer Configuration Course FOr Cisco Hands on Experience.**


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

like it

Thank you Elie:)

Great help! Nice ty.

Thank you Tanos:)

VLAN

Merci

Welcome Sow
