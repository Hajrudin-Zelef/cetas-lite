---
id: collect-261001-cisco/cisco/vtp-configuration-with-packet-tracer-7-steps-cisco-vtp-config
title: "vtp-configuration-with-packet-tracer-7-steps-cisco-vtp-config"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["parameters", "pruning"]
source: docs/RAG/collect-261001-cisco/vtp-configuration-with-packet-tracer-7-steps-cisco-vtp-config.md
source_anchor: ""
source_lines: [1, 168]
sha256: c9d57119ec093f8e8faa001300cd23bf5ca67c858a21408127101379e0c872d7
---

# vtp-configuration-with-packet-tracer-7-steps-cisco-vtp-config

Table of Contents

In this lesson, we will focus on How to Configure **VTP** on Cisco switches with Packet Tracer.  For VLAN Trunking Protocol, **Cisco VTP Configuration**, we will use the topology below with two switches and two PCs. Beside VTP Configuration, you can also check **Cisco VLAN Configuration on Packet Tracer**.


You can **DOWNLOAD** the **Packet Tracer** example with **.pkt** format **HERE**.

For all Packet Tracer Examples and Files, you can check **Packet Tracer Labs** Page.


First of all we should configure the **trunk** between swithes. Because **VTP information** pass through only on trunk links on management **VLAN** (VLAN 1). To do this:


```
SwitchA>
```
 **enable**
SwitchA# **configure terminal**
SwitchA(config)# **interface fa 0/24**
SwitchA(config-if)#**switchport trunk encapsulation dot1q**
SwitchA(config-if)# **switchport mode trunk**
%LINEPROTO-5-UPDOWN: Line protocol on Interface FastEthernet0/24, changed state to down
%LINEPROTO-5-UPDOWN: Line protocol on Interface FastEthernet0/24, changed state to up
SwitchA(config-if)# **end**

Now the **trunk** is **activated**. For switchport mode trunk command it is not necessary to use the same command at the other side of the link, without this job it is automatically activated.



In this step, we will configure Switch A as a **VTP Server** and the Switch B as a **VTP Client**. The **default** switch VTP mode is **VTP Server**, so we will not change anything on SwitchA. But in SwitchB we will change vtp mode. Then in both switches we will configure **VTP domain, VTP password, vtp version** and **vtp prunning**.


Firstly, let’s start with switch B and set switch B as **vtp client mode.**


```
SwitchB> 
```
**enable**
SwitchB# **configure terminal**
SwitchB(config)# **vtp mode client**

Then, let’s change the default **domain name** from **NULL** to **cisco**.

```
SwitchB(config)# 
```
**vtp domain cisco**
After that, we will set **vtp password** as **abc123**, vtp **version 2**. And we will enable **vtp prunning**.

 

```
SwitchB(config)# 
```
**vtp password abc123**
SwitchB(config)# **vtp version 2**
SwitchB(config)# **vtp prunning enable**
SwitchB(config)# **end**


Now, it is time to configure switch A. We will configure the same parameters for the Switch A except **VTP Mode**. Switch A will be in **VTP Server Mode** by **default**. But to show you, we will use “**vtp server mode**” command..


```
SwitchA# 
```
**enable** 
SwitchA(config)# **vtp mode server**
 

Device mode already **VTP SERVER**.


```
SwitchA(config)# 
```
**vtp domain ciscoo**
Changing VTP domain name from NULL to ciscoo
04:50:49 %DTP-5-DOMAINMISMATCH: Unable to perform trunk negotiation
on port Fa0/24 because of VTP domain mismatch.

```
SwitchA(config)# 
```
**vtp domain cisco**
Changing VTP domain name from ciscoo to cisco
SwitchA(config)# **vtp password abc123**
SwitchA(config)# **vtp version 2**
SwitchA(config)# **vtp prunning enable**
SwitchA(config)# **end**

Here, during **VTP domain configuration**, an error occurred because of the **domain missmatch** and after that when we corrected the configuration then this error dissappeared.



To verify **Cisco VTP Configuration**, we will use **“show vtp status”** command on both switches.


```
SwithA# 
```
**show vtp status**
VTP Version : 2
Configuration Revision : 0
Maximum VLANs supported locally : 255
Number of existing VLANs : 5
VTP Operating Mode : Server
VTP Domain Name : cisco
VTP Pruning Mode : Disabled
VTP V2 Mode : Disabled
VTP Traps Generation : Disabled
MD5 digest : 0x12 0x34 0xBB 0x77 0x23 0xa4 0xA6 0x86
Configuration last modified by 0.0.0.0 at 4-6-23 07:15:24
Local updater ID is 0.0.0.0 (no valid interface found)

```
SwithB# 
```
**show vtp status**
VTP Version : 2
Configuration Revision : 0
Maximum VLANs supported locally : 255
Number of existing VLANs : 5
VTP Operating Mode : Client
VTP Domain Name : cisco
VTP Pruning Mode : Disabled
VTP V2 Mode : Disabled
VTP Traps Generation : Disabled
MD5 digest : 0x12 0x34 0xBB 0x77 0x23 0xa4 0xA6 0x86
Configuration last modified by 0.0.0.0 at 4-6-23 07:15:24
Local updater ID is 0.0.0.0 (no valid interface found)

To test our configuration, lets create three additional **VLANs** on Switch A and then check Switch B if the same VLANs created under it too. To do this:


```
SwitchA(config)# 
```
**vlan 10**
SwitchA(config)# **vlan 20**
SwitchA(config)# **vlan 30**
SwitchA(config)# **end**

When we use **“show vlan brief”** command on Switch A we will see the created VLANs like below:




Lastly, let ‘s use “**show vtp status**” command once more to check the status of vtp. Here, we will see that the total vlan number is 8 with additional 3 vlans.




I hope this article will be helpfull for you…


**You can download “Packet Tracer” in Tools section.**

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
