---
id: collect-261001-cisco/cisco/vlan-configuration-on-cisco-packet-tracer-2026-updated-1
title: "vlan-configuration-on-cisco-packet-tracer-2026-updated"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/vlan-configuration-on-cisco-packet-tracer-2026-updated.md
source_anchor: ""
source_lines: [1, 192]
sha256: d7b7fd80af6af711c35730cc7aea3332a0ca5554dd0c3c192708bab323fa8f7f
---

# vlan-configuration-on-cisco-packet-tracer-2026-updated

Table of Contents

In this **VLAN Cisco Packet Tracer** Example, we will learn **how to configure** **VLANs on Cisco** switches. For this lesson, we will use the **VLAN topology** below. In this topology, there are 2 **Cisco Catalyst 2960-24** switches and 8 **PCs**. After this **VLAN Packet Tracer Example**, you can configure **VLANs** on your network easily. VLANs are also very important for Cisco **CCNA**, **CCNP** and **CCIE** Certifications.


*Packet Tracer VLAN Topology Example*

You can **DOWNLOAD** the **Cisco Packet Tracer** example with **.pkt** format **at the end of the lesson**.

**For all Packet Tracer Examples and Files, you can check Packet Tracer Labs Page.**

In this Cisco VLAN configuration example, we will follow the below configuration steps one by one:



For our **VLAN Configuration example**, we will set our PC IPaddresses as below. These ip addresses will be required at the end of this configuration example to test our configuration for Virtual LAN verification.

**PC 0 –>** 10.0.0.1 *VLAN 2 (HR)*

**PC 1 –>** 10.0.0.2 *VLAN 1 (Accounting)*

**PC 2 –>** 10.0.0.3 *VLAN 1 (Accounting)*

**PC 3 –>** 10.0.0.4 *VLAN 3 (R&D)*

**PC 4 –>** 10.0.0.5 *VLAN 2 (HR)*

**PC 5 –>** 10.0.0.6 *VLAN 1 (Accounting)*

**PC 6 –>** 10.0.0.7 *VLAN 1 (Accounting)*

**PC 7 –>** 10.0.0.8 *VLAN 3 (R&D)*


After PC IP configurations, now, we can start our **VLAN Packet Tracer Configuration** steps.


To create a VLAN on a Cisco switch, we use “**vlan vlan-id**” command under global configuration mode. Here, we will create **VLAN 2** and **VLAN 3**. By **default** **VLAN 1** has already created. We will also give a name to these vlans with “**name *vlan-name***” command.


```
Switch 1# 
```
**configure terminal**
Switch 1(config)# **vlan 1**
Switch 1(config-vlan)# **name Accounting**
Switch 1(config-vlan)# **vlan 2**
Switch 1(config-vlan)# **name HR**
Switch 1(config-vlan)# **vlan 3**
Switch 1(config-vlan)# **name R&D**

We will do the similar configuration on switch 2 also.


```
Switch 2# 
```
**configure terminal**
Switch 2(config)# **vlan 1**
Switch 2(config-vlan)# **name Accounting**
Switch 2(config-vlan)# **vlan 2**
Switch 2(config-vlan)# **name HR**
Switch 2(config-vlan)# **vlan 3**
Switch 2(config-vlan)# **name R&D**


Cisco switches use two primary port types: **Access ports** and **Trunk ports.** Access ports connect end devices to a single VLAN, while trunk ports carry traffic for multiple VLANs between network devices using IEEE 802.1Q tagging. Below we will configure both of them in this Virtual LAN configuration example.



Access ports are the ports that can be a member of a single VLAN. To configure a port as access port to assign a specific VLAN, we will use two commands. We use “**switchport mode access**” command to set the interface as access port. Then we use “**switchport access vlan *vlan-id***” command to assign the related VLAN to that interface.


Here, we will configure **fastEthernet** **0/2** as the member of **VLAN 2**, **fastEthernet** **0/3** and **fastEthernet 0/4** as the member of **VLAN 1** and  **fastEthernet** **0/5** as the member of **VLAN 3**.


```
Switch 1(config)# 
```
**interface fastEthernet 0/2**
Switch 1(config-if)# **switchport mode access**
Switch 1(config-if)# **switchport access vlan 2**
Switch 1(config-if)# 
Switch 1(config-if)# **interface fastEthernet 0/3**
Switch 1(config-if)# **switchport mode access**
Switch 1(config-if)# **switchport access vlan 1**
Switch 1(config-if)# 
Switch 1(config-if)# **interface fastEthernet 0/4**
Switch 1(config-if)# **switchport mode access**
Switch 1(config-if)# **switchport access vlan 1**
Switch 1(config-if)# 
Switch 1(config)# **interface fastEthernet 0/5**
Switch 1(config-if)# **switchport mode access**
Switch 1(config-if)# **switchport access vlan 3**
We will do the similar configuration on the second switch with some differences. Here, we will configure **fastEthernet** **0/2** as the member of VLAN 3. while we will configure **fastEthernet** **0/3** and **fastEthernet** **0/4** as the member of VLAN 2.


```
Switch 2(config)# 
```
**interface fastEthernet 0/2**
Switch 2(config-if)# **switchport mode access**
Switch 2(config-if)# **switchport access vlan 2**
Switch 2(config-if)# 
Switch 2(config-if)# **interface fastEthernet 0/3**
Switch 2(config-if)# **switchport mode access**
Switch 2(config-if)# **switchport access vlan 1**
Switch 2(config-if)# 
Switch 2(config-if)# **interface fastEthernet 0/4**
Switch 2(config-if)# **switchport mode access**
Switch 2(config-if)# **switchport access vlan 1**
Switch 2(config-if)# 
Switch 2(config-if)# **interface fastEthernet 0/5**
Switch 2(config-if)# **switchport mode access**
Switch 2(config-if)# **switchport access vlan 3**

After creating VLANs, setting access ports and assigning VLANs, now it is time to create Trunk ports that will allow all our VLANs. To set a port as trunk port we will use “**switchport mode trunk**” command. Then we will use “**no negotiate**” command to prevent negotiation about the port role. Lastly, we will add the allowed VLANs on this Trunk port with “**switchport trunk allowed vlan *vlan-range***” command.


Here, our trunk ports will be  **fastEthernet 0/1.**


```
Switch 1(config)# 
```
**interface fastEthernet 0/1**
Switch 1(config-if)# **switchport mode trunk**
Switch 1(config-if)# **switchport nonegotiate**
Switch 1(config-if)# **switchport trunk allowed vlan 1-3**

We will do the same configuration on the second Cisco switch:


```
Switch 2(config)# 
```
**interface fastEthernet 0/1**
Switch 2(config-if)# **switchport mode trunk**
Switch 2(config-if)# **switchport nonegotiate**
Switch 2(config-if)# **switchport trunk allowed vlan 1-3**

When traffic crosses a trunk link, Cisco switches add an **IEEE 802.1Q tag** to Ethernet frames. This tag identifies the **VLAN ID** so multiple VLANs can share the same physical connection without mixing traffic.


Unlike other VLANs, the native VLAN is **transmitted** across a trunk link **without an IEEE 802.1Q tag**. Both ends of the trunk should use the **same native VLAN** to avoid connectivity and security issues.



After our configuration steps, we will save our configuration with “**copy running-config startup-config**” command on both Cisco switches in packet tracer.


```
Switch 1(config-if)# 
```
**end**
Switch 1# **copy running-config startup-config**

```
Switch 2(config-if)# 
```
**end**
Switch 2# **copy running-config startup-config**

Our last step of **VLAN Packet Tracer Example** is configuration **verification**. to verify our VLAN Packet Tracer Configuration, we will use verification commands like “**show vlan brief**“, “**show interfaces**“, “**show interfaces trunk**” etc.


To verify Cisco Packet Tracer VLAN configuration, firstly we will use “**show vlan brief**” command. With this command, we can see the creatd VLANs and the ports that are assigned to these VLANs. Here, the port that is not appear is Trunk port, fastEthernet 0/1.


```
Switch1#
```
**show vlan brief**
VLAN Name Status Ports
---- -------------------------------- --------- -------------------------------
1 default active Fa0/3, Fa0/4, Fa0/6, Fa0/7
Fa0/8, Fa0/9, Fa0/10, Fa0/11
Fa0/12, Fa0/13, Fa0/14, Fa0/15
Fa0/16, Fa0/17, Fa0/18, Fa0/19
Fa0/20, Fa0/21, Fa0/22, Fa0/23
Fa0/24, Gig0/1, Gig0/2
2 HR active Fa0/2
3 R&D active Fa0/5
1002 fddi-default active
1003 token-ring-default active
1004 fddinet-default active
1005 trnet-default active


Another important command on VLAN verification is using “**show interfaces *interface* switchport**” command. With this command, we can see a lot of details about our VLAN configuration.


