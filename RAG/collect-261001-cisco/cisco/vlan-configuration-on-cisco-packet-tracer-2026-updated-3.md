---
id: collect-261001-cisco/cisco/vlan-configuration-on-cisco-packet-tracer-2026-updated-3
title: "vlan-configuration-on-cisco-packet-tracer-2026-updated"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/vlan-configuration-on-cisco-packet-tracer-2026-updated.md
source_anchor: ""
source_lines: [436, 582]
sha256: e90b9a22efb933498e6df8ae51af414cccc0c5dbdbc78881208249681b6bdcdb
---

# vlan-configuration-on-cisco-packet-tracer-2026-updated

```
PC1:\>
```
**ping 10.0.0.8**
Pinging 10.0.0.8 with 32 bytes of data:
Request timed out.
Request timed out.
Request timed out.
Request timed out.
Ping statistics for 10.0.0.8:
Packets: Sent = 4, Received = 0, Lost = 4 (100% loss),

As you can see, ping is not successful because PC1 is in VLAN 1 while PC7 is in VLAN 3.


```
PC1:\>
```
**ping 10.0.0.5**
Pinging 10.0.0.5 with 32 bytes of data:
Request timed out.
Request timed out.
Request timed out.
Request timed out.
Ping statistics for 10.0.0.5:
Packets: Sent = 4, Received = 0, Lost = 4 (100% loss),

Again, ping from PC1 to PC4 is not successful. Because PC 1 is in VLAN 1 while PC4 is in VLAN 2. PCs in different VLANs can not ping each other even if there is **no VLAN routing**. This is the case of another lesson.



As you can see above, the **PC 1** can ping the **PCs** in the **same VLAN**, even if it is connected to a different switch. You can find the **packet tracer example (.pkt)**, switches’ and PCs’ configurations below.


You can download the **Cisco Packet Tracer** example with **.pkt** fortmat  here.

**You can download “Packet Tracer” in Tools section.**


In this **VLAN Configuration on Cisco Packet Tracer** example, we have learned the required commands for Virtual LAN Configuration on Cisco switches.


**VLAN Configurations** are very common in a companies switched network. Different teams, different departments are **divided** with **different VLANs** to create small **sub networks**. So, this configuration is very critical for a network engineer. Because, you will do this configurations a lot in your network engineering career. To learn this lesson of **Cisco switching**, you can do more examples on this lesson by yourself.


A **VLAN** (Virtual Local Area Network) is a **logical network** that divides a physical switch into **multiple separate broadcast domains.** VLANs improve network performance, security and management by grouping devices regardless of their physical location. Devices in different VLANs **cannot** communicate directly **unless inter-VLAN routing** is configured.


We can use “**vlan <vlan-id>**” command to **create VLANs** on Cisco switch. If we do not use this command, when we assign a switch port, VLAN can automatically created. To configure a VLAN on a Cisco switch, firstly we should **create the VLAN** and then optionally **assign a name** to this VLAN. After that we can **assign** **switch ports** to that VLAN. To check created VLANs, we can use “**show vlan brief**” command. Below we will create VLAN 10 with name **Engineering**.


```
Switch(config)#
```
 **vlan 10**
Switch(config-vlan)# **name Engineering**
Switch(config-vlan)# **exit**

An access port belongs to a **single VLAN** and is typically connected to end devices such as PCs, printers or IP phones. A trunk port carries traffic for **multiple VLANs** simultaneously and is commonly used between switches, routers or other networking devices.


An access port **connects** a **single end device**, such as a PC or printer, to a **specific VLAN**. To configure the interface as an access port we use “**switchport mode access**” command under the related interface. Below, we set fastEthernet 0/2 as VLAN access port.

```
Switch(config)#
```
 **interface fastEthernet 0/2**
Switch(config-if)# **switchport mode access**

To assign the desired VLAN to that interface, we use “**switchport access vlan <vlan-id>**” command after setting the interface as access port with “**switchport mode access**” command.


```
Switch(config)# 
```
**interface fastEthernet 0/2**
Switch(config-if)# **switchport access vlan 20**

A trunk port **carries** traffic for **multiple VLANs** between switches or other network devices. To configure the interface as a trunk we use “switchport mode trunk” command under the trunk interface.


```
Switch(config)# 
```
**interface gigabitEthernet 0/1**
Switch(config-if)# **switchport mode trunk**

**By default**, a Cisco trunk port allows traffic from **all VLAN**s. To improve security and reduce unnecessary traffic, you can limit the trunk to **carry only specific VLANs** using the “**switchport trunk allowed vlan <vlan-id>**” command under the trunk interface. We can use multiple VLANs separated by **comma (,)** or we can speciy a VLAN range with a **hyphen(-)** between starting VLAN and ending VLAN. Below, firstly, we will allow **VLANs 10,20** and **30**. Then, we will allow a range from **VLAN 5-9**.


```
Switch(config)# 
```
**interface gigabitEthernet 0/1**
Switch(config-if)# **switchport mode trunk**
Switch(config-if)# **switchport trunk allowed vlan 10,20,30**

```
Switch(config-if)#
```
 **switchport trunk allowed vlan 5-9**

The “**show vlan brief**” command displays all configured VLANs, their names, status and assigned access ports. It is the most commonly used command to verify VLAN configuration.


```
Switch#
```
 **show vlan brief**

After configuring VLANs, to verify VLAN configuration we can use “**show vlan brief**“, “**show interfaces trunk**“, ” **show running-config**” commands.


```
Switch# 
```
**show vlan brief**
Switch# **show interfaces trunk**
Switch# **show running-config**

By default, all switch ports belong to **VLAN 1** on Cisco switches. VLAN 1 **exists automatically** and **cannot be deleted**. Although it is the default VLAN, **using a separate management VLAN** is considered a **best practice** for improved network security.


**A Layer 2 VLAN** separates devices into different broadcast domains at the **Data Link layer**, while a **Layer 3 VLAN** includes routing capabilities that allow communication between different VLANs. Inter-VLAN communication requires a Layer 3 switch or a router.


By default, separate VLANs can not communicate each other because VLANs **separate domains** in **Layer 2** (Data-link layer). To overcome this, a **Layer 3** (Network layer) device is reuired. This can be a multi layer switch or a router. By using a layer 3 device, **routing between VLANs** is configured. This is Inter-VLAN Routing.  Inter-VLAN routing allows devices in different VLANs to communicate with each other. This can be implemented by using SVIs or a Router-on-a-Stick configuration.


You can remove an existing VLAN using the “**no vlan**” command in global configuration mode. After deleting the VLAN, you can verify the change with “**show vlan brief**” command.


```
Switch(config)# 
```
**no vlan 20**


Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

Am very interested this topic
