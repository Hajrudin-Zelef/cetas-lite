---
id: collect-261001-huawei/huawei/rstp-configuration-on-huawei-ec747969
title: "rstp-configuration-on-huawei-ec747969"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/rstp-configuration-on-huawei-ec747969.md
source_anchor: ""
source_lines: [1, 39]
sha256: 4e1a775f0e83955b3e5da08a4bf861b4d97996d487f0eb685aaae5adfdd87486
---

# rstp-configuration-on-huawei-ec747969

In this lessons, we will focus on RSTP (Rapid Spanning Tree Protocol) Configuration on Huawei Switches. For our Huawei RSTP Example, we will use the basic topology below.
Let’s start to configure RSTP on one of these Huawei Switches, Switch 1.
You can download this configuration on Huawei eNSP Labs Page.
Table of Contents
STP Mode Selection
Firstly, we will configure the mode of used STP like STP Configuration. Here, we will configure RSTP, so we will select RSTP.
system-view
[Huawei-Switch-1] stp mode ?
mstp Multiple Spanning Tree Protocol (MSTP) mode
rstp Rapid Spanning Tree Protocol (RSTP) mode
stp Spanning Tree Protocol (STP) mode
[Huawei-Switch-1] stp mode rstp
Switching RSTP Mode
In some situation we need to switch to RSTP mode. For example, if you connect an RSTP configured device to STP configured device and then disconnect them. Because of the fact that, during the configuration both mode will work as STP, the RSTP enabled device will need an RSTP switch. When we use “stp mcheck” command, then it will work with RSTP again.
[Huawei-Switch-1] stp mcheck
Bridge Priority Configuration
To manuplate Root Bridge Selection, we wil configure Bridge Priority value of the Switch 1. When we configure the lowest Priority value in this switch, it will be selected as Root Bridge automatically.
The default Bridge Priority is 32768. We will configure Switch 1 Bridge Priority as 4096. So, it will be the lowest one in the network and this will provide Switch 1 as Root Bridge.
[Huawei-Switch-1] stp priority 4096
If we do not manuplate this, the priority values will remain as default, 32768. Acocrding to MAC addresses, the Root Bridge will be selected. So, Switch 2 has the lowest MAC. It would be Root Bridge. Think about that, what if Switch 2 is a very low capacity switch? This can cause a bad network affect.
Edge Port Configuration
In RSTP, we can set Edge Ports for Host devices that is not need to receive and send BPDUs They will not participate in RSTP process.
By default all the ports are in Non Edge Mode. To configure all porst as Edge Port, we will use global “stp edged-port default” command.
[Huawei-Switch-1] stp edged-port default
We can also use “stp edged-port enabled” command under interfaces to set a specific Edge Port.Here, our PC is connected to the Ge 0/0/3, so our edge port is Ge 0/0/3 on Switch 3.
Loop Protection
It is used to prevent one node to receive any BPDU. To do this use “stp loop-protection” command under a specific interface. Here, we will also configure Ge 0/0/3 with this command. Normally a PC do not use BPDU, but a malicious atatck can come from a PC.
[Huawei-Switch-1] interface GigabitEthernet 0/0/3
[Huawei-Switch-1-GigabitEthernet0/0/3] stp loop-protection
Loop Protection is disabled on the interfaces by default.
Root Protection
As a protection mechanism for Root Bridge, we can configure the ports of Root Bridge with “stp root-protection”.
Root Protection is disabled on the interfaces by default.
AS an important note, we can not use both “Root Protection” and Loop Protection” on the same interface.
Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.
He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.
He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.
IPCisco.com | Best Route to Your Dreams
Leave a Reply
