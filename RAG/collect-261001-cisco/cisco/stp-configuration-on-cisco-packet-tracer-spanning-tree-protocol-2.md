---
id: collect-261001-cisco/cisco/stp-configuration-on-cisco-packet-tracer-spanning-tree-protocol-2
title: "stp-configuration-on-cisco-packet-tracer-spanning-tree-protocol"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/stp-configuration-on-cisco-packet-tracer-spanning-tree-protocol.md
source_anchor: ""
source_lines: [220, 315]
sha256: 0200c50628636396eafcd24ab7c1702b6561d08bdf06c1f9297fb44c25273648
---

# stp-configuration-on-cisco-packet-tracer-spanning-tree-protocol

```
Switch3#
```
**show spanning-tree active** 
VLAN0001
Spanning tree enabled protocol ieee
Root ID Priority 32769
Address 0001.C90E.EDC0
Cost 38
Port 1(FastEthernet0/1)
Hello Time 2 sec Max Age 20 sec Forward Delay 15 sec
Bridge ID Priority 32769 (priority 32768 sys-id-ext 1)
Address 00D0.58E3.0126
Hello Time 2 sec Max Age 20 sec Forward Delay 15 sec
Aging Time 20
Interface Role Sts Cost Prio.Nbr Type
---------------- ---- --- --------- -------- --------------------------------
Fa0/1 Root FWD 19 128.1 P2p
Fa0/2 Altn BLK 19 128.2 P2p

As you can see above, the STP blocks one of the port of Switch3. This election is done according to the **cost to the root**. The **Designated Ports** are selected and the remainning **Non-Designated Port** on a segment is blocked. Remember, **only one** Designated Port can exist in a segment.


We can summarize the last logical network topology like below:


I hope this can be useful for you, to understand STP better. **STP(Spanning Tree Protocol)** is the first protocol for this mechanism. Beside STP, there are many protocols used today. **RSTP(Rapid Spanning Tree Protocol), PVRST (Per VLAN Rapid Spanning Tree), PVRST+** and **MST** are these protocol. In the following articles we will discuss these protocols one by one.


**You can download “Packet Tracer” in Tools section.**


As we have mentined before, **Spanning Tree Protocol (STP)** is **enabled** by **default** on Cisco switches. If it is **disabled**, you can enable it with the below configuration commands. 

```
Switch# 
```
**configure terminal**
Switch(config)# **spanning-tree**
There are different STP Modes. To select one of these Spanning Tree Protocol Modes, you an use the below STP commands.


```
Switch(config)# 
```
**spanning-tree mode ?**
mst Multiple spanning tree mode
pvst Per-Vlan spanning tree mode
rapid-pvst Per-Vlan rapid spanning tree mode
Switch(config)# **spanning-tree mode stp**

The meanings of these STP mdoes are given below:

**stp —** Classic STP.

**rstp —** Faster convergence of the spanning tree.

**mst —** MSTP is based on RSTP. 


To configure Bridge Priority for STP, you can use the below command. The range of this Bridge priority is between **0 and 61440**.And here this value must be in increments of **4096**. The defaulr STP Bridge Priority is **32768**.

```
Switch(config)# spanning-tree priority 32768
```

Here the range of **Hello Message Frequency** is between **1 and 10 seconds**. The **default** value is **2 seconds**.


```
Switch(config)# spanning-tree hello-time 2
```

Here the range of **STP Maximum Age** is between **6 and 40 seconds**. The **default** value is **20 seconds**.

 

```
Switch(config)# spanning-tree max-age 20
```

Forward Time is the time a port remains in the listening and learning states before entering the forwarding state. The range of **Bridge Forward Time** is between **4 and 30 seconds**. The **default** value is **15 seconds**.

```
Switch(config)# spanning-tree forward-time 15
```

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
