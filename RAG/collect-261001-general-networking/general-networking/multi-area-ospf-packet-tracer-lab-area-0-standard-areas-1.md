---
id: collect-261001-general-networking/general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas-1
title: "multi-area-ospf-packet-tracer-lab-area-0-standard-areas"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas.md
source_anchor: ""
source_lines: [1, 228]
sha256: 8e5520b0c11bd7728a6b6bb7a8ecca9d6ac517bbb482d0ffe1abb30a558ef0b9
---

# multi-area-ospf-packet-tracer-lab-area-0-standard-areas

In this lesson we will focus Multi Area OSPF Configuration on Packet Tracer. We will work with **OSPF Standard Area** and **OSPF Backbone Area** together for our OSPF configuration example. In other words, we will do a **Multi Area OSPF Packet Tracer Configuration** with Backbone Area and OSPF Standard Areas. In other lesson, we will also give example for the other **OSPF Area Types** and their configuration examples on **Packet Tracer**.

You can **DOWNLOAD** the **Cisco Packet Tracer** example with **.pkt** format **at the end of this lesson**.

As you know from the previous lessons, **OSPF** **(Open Shortest Path First)** has **6** different **areas**. These OSPF Area Types are given below:


**You can check the below lessons for the configuration of different OSPF Area Types:**

In this OSPF Configuration lesson, we will focus on **Backbone Area** and **Standard (Normal) Area**. Beside this, we will see the configuration of **virtual-link** on OSPF. What is **virtual-link in OSPF?** It is the link that used to connect the normal areas to the backbone area, if they are not connected directly to the backbone area. Remember, in OSPF, there was a rule. All areas must be connected to the backbone area, **Area 0**. If they are not, they can temporarily connect to the backbone via **virtual-links**.

*OSPF Backbone Area and Standard Area with Accepted LSAs*


**You can also check Basic OSPF Confiuration**

**What is Packet Tracer?** | **Download Packet Tracer** 

**Packet Tracer Labs Coure** | **Download Packet Tracer Labs**


**Backbone Area** is also a Normal Area but it is **Area 0 OSPF**.

Normal Areas accept the **Summary LSAs** from other Areas **(Type 3 and Type 4 LSAs)**. They accept also the **External LSAs (Type 5 LSAs)**. Type 1 and Type 2 LSA are already accepted inside area.



Table of Contents

We talked about theorical too much. This post aims to show you the configuration of this areas and as you know, doing the configuration is the most effective way of learning network protocols.

Our topology will be like below fort he first example.

Firstly let’s configure the IP Addresses on all routers:

__Router1__

```
Router1>
```
**enable**
Router1# **configure terminal**
Router1(config)# **interface GigabitEthernet0/0**
Router1(config-if)# **ip address 10.1.0.1 255.255.255.0**
Router1(config-if)# **no shutdown**
Router1(config-if)# **exit**
Router1(config)# **interface GigabitEthernet0/1**
Router1(config-if)# **ip address 10.2.0.1 255.255.255.0**
Router1(config-if)# **no shutdown**
Router1(config-if)# **end**
Router1# **copy running-config startup-config**

__Router2__

```
Router2>
```
**enable**
Router2# **configure terminal**
Router2(config)# **interface GigabitEthernet0/0**
Router2(config-if)# **ip address 10.1.0.2 255.255.255.0**
Router2(config-if)# **no shutdown**
Router2(config-if)# **exit**
Router2(config)# **interface GigabitEthernet0/1**
Router2(config-if)# **ip address 10.3.0.1 255.255.255.0**
Router2(config-if)# **no shutdown**
Router2(config-if)# **end**
Router2# **copy running-config startup-config**

__Router3__

```
Router3>
```
**enable**
Router3# **configure terminal**
Router3(config)# **interface GigabitEthernet0/0**
Router3(config-if)# **ip address 10.2.0.2 255.255.255.0**
Router3(config-if)# **no shutdown**
Router3(config-if)# **exit**
Router3# **copy running-config startup-config**

__Router4__

```
Router4>
```
**enable**
Router4# **configure terminal**
Router4(config)# **interface GigabitEthernet0/0**
Router4(config-if)# **ip address 10.4.0.1 255.255.255.0**
Router4(config-if)# **no shutdown**
Router4(config-if)# **exit**
Router4(config)# **interface GigabitEthernet0/1**
Router4(config-if)# **ip address 10.3.0.2 255.255.255.0**
Router4(config-if)# **no shutdown**
Router4(config-if)# **end**
Router4# **copy running-config startup-config**

__Router5__

```
Router5>
```
**enable**
Router5# **configure terminal**
Router5(config)# **interface GigabitEthernet0/0**
Router5(config-if)# **ip address 10.4.0.2 255.255.255.0**
Router5(config-if)# **no shutdown**
Router5(config-if)# **exit**
Router5(config)# **interface GigabitEthernet0/1**
Router5(config-if)# **ip address 10.5.0.1 255.255.255.0**
Router5(config-if)# **no shutdown**
Router5(config-if)# **end**
Router5# **copy running-config startup-config**

__Router6__

```
Router6>
```
**enable**
Router6# **configure terminal**
Router6(config)# **interface GigabitEthernet0/1**
Router6(config-if)# **ip address 10.5.0.2 255.255.255.0**
Router6(config-if)# **no shutdown**
Router6(config-if)# **end**
Router6# **copy running-config startup-config**

After this basic interface configurations, now, we will focus on the main part of our **Multi Area OSPF Packet Tracer Configuraiton example.** Here, we will configure **OSPF (Open Shortest Path First)** on all routers. Here, OSPF process number will be 1 and the Routerx’s router ID will be x.x.x.x. Beside this, all the connected areas will be configured.

```
Router1>
```
**enable**
Router1# **configure terminal**
Router1(config)# **router ospf 1**
Router1(config-router)# **router-id 1.1.1.1**
Router1(config-router)# **network 10.1.0.0 0.0.0.255 area 0**
Router1(config-router)# **network 10.2.0.0 0.0.0.255 area 1**
Router1(config-router)# **end**
Router1# **copy running-config startup-config**

```
Router2>
```
**enable**
Router2# **configure terminal**
Router2(config)# **router ospf 1**
Router2(config-router)# **router-id 2.2.2.2**
Router2(config-router)# **network 10.1.0.0 0.0.0.255 area 0**
Router2(config-router)# **network 10.3.0.0 0.0.0.255 area 2**
Router2(config-router)# **end**
Router2# **copy running-config startup-config**

```
Router3>
```
**enable**
Router3# **configure terminal**
Router3(config)# **router ospf 1**
Router3(config-router)# **router-id 3.3.3.3**
Router3(config-router)# **network 10.2.0.0 0.0.0.255 area 1**
Router3(config-router)# **end**
Router3# **copy running-config startup-config**

```
Router4>
```
**enable**
Router4# **configure terminal**
Router4(config)# **router ospf 1**
Router4(config-router)# **router-id 4.4.4.4**
Router4(config-router)# **network 10.3.0.0 0.0.0.255 area 2**
Router4(config-router)# **network 10.4.0.0 0.0.0.255 area 2**
Router4(config-router)# **end**
Router4# **copy running-config startup-config**

```
Router5>
```
**enable**
Router5# **configure terminal**
Router5(config)# **router ospf 1**
Router5(config-router)# **router-id 5.5.5.5**
Router5(config-router)# **network 10.4.0.0 0.0.0.255 area 2**
Router5(config-router)# **network 10.5.0.0 0.0.0.255 area 3**
Router5(config-router)# **end**
Router5# **copy running-config startup-config**

```
Router6>
```
**enable**
Router6# **configure terminal**
Router6(config)# **router ospf 1**
Router6(config-router)# **router-id 6.6.6.6**
Router6(config-router)# **network 10.5.0.0 0.0.0.255 area 3**
Router6(config-router)# **end**
Router6# **copy running-config startup-config**

After doing this configuration we will see all the network on **Topology Table** except **Area 3**. Area 3 is not directly connected to the Backbone Area,Area 0. So in the routing table of the routers there will be no route to this Area 3.

To overcome this issue, on **Router 2** and **Router 5**, we will configure the **virtual-link**. After this configuration, we will see all the routes in all the routers’ Topology Tables.


In this step, we will configureOSPF Virtual-Link between Router 2 and Router 5. After this configuration, we will see all the routes in all the routers’ Topology Tables.

```
Router2>
```
**enable**
Router2# **configure terminal**
Router2(config)# **router ospf 1**
Router2(config-router)# **area 2 virtual-link 5.5.5.5**
Router2(config-router)# **end**
Router2# **copy running-config startup-config**

```
Router5>
```
**enable**
Router5# **configure terminal**
Router5(config)# **router ospf 1**
Router5(config-router)# **area 2 virtual-link 2.2.2.2**
Router5(config-router)# **end**
Router5# **copy running-config startup-config**

