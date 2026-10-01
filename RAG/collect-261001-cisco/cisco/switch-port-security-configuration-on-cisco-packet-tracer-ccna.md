---
id: collect-261001-cisco/cisco/switch-port-security-configuration-on-cisco-packet-tracer-ccna
title: "switch-port-security-configuration-on-cisco-packet-tracer-ccna"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/switch-port-security-configuration-on-cisco-packet-tracer-ccna.md
source_anchor: ""
source_lines: [1, 197]
sha256: 5a687f089f7a519781919ac590b45235cd3f7a7a8d78c16c25d1a635f3031f15
---

# switch-port-security-configuration-on-cisco-packet-tracer-ccna

Table of Contents

In this Cisco Switch Security configuration lesson, we will focus on **Switch Port Security Cisco Configuration** with **Cisco Packet Tracer**. In this **Packet Tracer Switch Port Security** Configuration example, you will learn **How to Configure Switch Port Security** on **Cisco IOS**. For our Cisco Port Security Configuration, we will use the below network topology.


**You can DOWNLOAD the Packet Tracer example with .pkt format HERE.**

**You can download all Cisco Packet Tracer Labs on Cisco Packet Tracer Labs** 


Here we will use **Four Switch Port Security Scenario** on four diffent switch port. According to these Switch Port Security Scenarios, we will do the below **Switch Port Security Configuration on Cisco Switch.**



In this first switch port security configuration example, we will do the below configuration on the first port of Cisco switch.

**Switch Port FastEthernet 0/1 Configuration**


```
Switch(config)# 
```
**interface fastEthernet 0/1**
Switch(config-if)# **switchport mode access**
Switch(config-if)# **switchport port-security**
Switch(config-if)# **switchport port-security maximum 2**
Switch(config-if)# **switchport port-security mac-address 00E0.B085.4951**
Switch(config-if)# **switchport port-security mac-address 0003.e445.1485**
Switch(config-if)# **switchport port-security violation shutdown**


In this second switch port security example, we will do the below configuration on the second port of Cisco switch.

**Switch Port FastEthernet 0/2 Configuration**


```
Switch(config)# 
```
**interface fastEthernet 0/2**
Switch(config-if)# **switchport mode access** 
Switch(config-if)# **switchport port-security** 
Switch(config-if)# **switchport port-security maximum 2**
Switch(config-if)# **switchport port-security mac-address sticky** 
Switch(config-if)# **switchport port-security violation restrict**


In third switch port security configuration example, we will do the below configuration on the third port.

**Switch Port FastEthernet 0/3 Configuration**


```
Switch(config)# 
```
**inter fastEthernet 0/3**
Switch(config-if)# **switchport mode access** 
Switch(config-if)# **switchport port-security maximum 1**
Switch(config-if)# **switchport port-security mac-address 1111.1111.1111**
Switch(config-if)# **switchport port-security violation shutdown**

 

In this last switch port security example, we will do the below configuration on the fourth port of Cisco switch.

**Switch Port FastEthernet 0/4 Configuration**


```
Switch(config)# 
```
**inter fastEthernet 0/4**
Switch(config-if)# **switchport mode access** 
Switch(config-if)# **switchport port-security** 
Switch(config-if)# **switchport port-security maximum 1**
Switch(config-if)# **switchport port-security mac-address sticky** 
Switch(config-if)# **switchport port-security violation protect**

After this **Port Security configuration**, connect the PC’s one by one. After this process, you will see that the **Port Security configuration** on switch, will avoid the unwanted devices and limit the max number of MAC addresses as configured.



You can also check the below **Switch Port Security verification commands** screenshots.


```
Switch# 
```
**show port-security** 
Secure Port MaxSecureAddr CurrentAddr SecurityViolation Security Action
(Count) (Count) (Count)
--------------------------------------------------------------------
Fa0/1 2 2 1 Shutdown
Fa0/2 2 2 1 Restrict
Fa0/3 1 1 1 Shutdown
Fa0/4 1 1 1 Protect
----------------------------------------------------------------------

```
Switch# 
```
**show port-security address** 
Secure Mac Address Table
-------------------------------------------------------------------------------
Vlan Mac Address Type Ports Remaining Age
(mins)
---- ----------- ---- ----- -------------
1 0003.E445.1485 SecureConfigured FastEthernet0/1 -
1 00E0.B085.4951 SecureConfigured FastEthernet0/1 -
1 0090.21B9.4D6D SecureSticky FastEthernet0/2 -
1 0009.7C63.A238 SecureSticky FastEthernet0/2 -
1 1111.1111.1111 SecureConfigured FastEthernet0/3 -
1 0009.7C63.A238 SecureSticky FastEthernet0/4 -
------------------------------------------------------------------------------

```
Switch# 
```
**show port-security interface fastEthernet 0/1**
Port Security : Enabled
Port Status : Secure-up
Violation Mode : Shutdown
Aging Time : 0 mins
Aging Type : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses : 2
Total MAC Addresses : 2
Configured MAC Addresses : 2
Sticky MAC Addresses : 0
Last Source Address:Vlan : 0000.0000.0000:0
Security Violation Count : 1

```
Switch# 
```
**show port-security interface fastEthernet 0/2**
Port Security : Enabled
Port Status : Secure-up
Violation Mode : Restrict
Aging Time : 0 mins
Aging Type : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses : 2
Total MAC Addresses : 2
Configured MAC Addresses : 0
Sticky MAC Addresses : 2
Last Source Address:Vlan : 0000.0000.0000:0
Security Violation Count : 1

```
Switch# 
```
**show port-security interface fastEthernet 0/3**
Port Security : Enabled
Port Status : Secure-up
Violation Mode : Shutdown
Aging Time : 0 mins
Aging Type : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses : 1
Total MAC Addresses : 1
Configured MAC Addresses : 1
Sticky MAC Addresses : 0
Last Source Address:Vlan : 0000.0000.0000:0
Security Violation Count : 1

```
Switch# 
```
**show port-security interface fastEthernet 0/4**
Port Security : Enabled
Port Status : Secure-up
Violation Mode : Protect
Aging Time : 0 mins
Aging Type : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses : 1
Total MAC Addresses : 1
Configured MAC Addresses : 0
Sticky MAC Addresses : 0
Last Source Address:Vlan : 0000.0000.0000:0
Security Violation Count : 1


In this example we have done **Port Security** configuration for our basic topology on Packet Tracer. After **Port Security** configuration, we have checked our **Port Security** configuration with **Port Security** verification commands.



Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
