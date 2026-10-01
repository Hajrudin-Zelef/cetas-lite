---
id: collect-261001-cisco/cisco/redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example-1
title: "redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example.md
source_anchor: ""
source_lines: [1, 179]
sha256: 951ca8804b7dcdaaf4e79c3cbb5de1b0b1d437a2dd58cd02f12ce19cb466cf9f
---

# redistribution-between-cisco-eigrp-into-ospf-and-vice-versa-example

The most common scenario for big enterprise networks with multiple routers is to have a single IGP routing protocol (IGP = Interior Gateway Protocol) running between the routers in order to distribute all the routing information of the network.

The most common IGP routing protocols used by today’s networks are OSPF, EIGRP (Cisco proprietary) and in some cases IS-IS. RIP also is an IGP but is not used anymore. It is found only in old legacy networks (or in lab environments for study purposes).

The above situation of having a single routing protocol is the most common case. However, there are networks that use two different IGP protocols for various reasons (such as merge of two companies that use different routing protocols, migration from one protocol to another, installation of router devices from a different vendor etc).

If you have two IGP routing protocols in the same network, then you must perform **Routing Redistribution** in order to exchange routing information between the two protocols.

Note that you can configure and enable two routing protocols on a router (e.g both OSPF and EIGRP). However, this is not enough to exchange information between the two routing protocols. You must configure Redistribution as well.

The tutorial below will focus on Routing Redistribution **between OSPF and EIGRP and vice-versa**. As the title says, this is an introductory tutorial in order to learn the basics of router redistribution. The network topology shown below is very basic but it will help you to expand the concepts in bigger networks.

From the topology above, assume we have a remote enterprise network (RED) running EIGRP only on router “**Ent_Remote**”. Also, we have another enterprise network (BLUE) with two routers (ENT1 and ENT2). ENT1 is running OSPF only and ENT2 is running both OSPF and EIGRP.

Our goal is to have all router devices learn all the routes from both networks. That is, we want to redistribute the routes of “Ent_Remote” into the routing table of router “ENT1” (which is running only OSPF) and vice-versa.

This scenario is often seen in real life when two companies merge and they have to converge without causing downtime or modifications to the networks.

Company BLUE is running OSPF and has the following networks: 10.10.0.0/24 & 10.10.100.0/24 behind router ENT2 and 172.16.1.0/24 and 172.16.100.0/24 behind ENT1.

Company RED (running EIGRP) has merged with BLUE and has networks 192.168.1.0/24 and 192.168.100.0/24 behind router “ENT_remote”

(The networks are emulated using loopbacks.)

NOTE: You can find the full router configurations at the end of this tutorial.

__First, let’s see the routing table on ENT1 before the redistribution:__

**ENT1#sh ip route**1.0.0.0/8 is variably subnetted, 2 subnets, 2 masks

C 1.1.1.0/30 is directly connected, FastEthernet4

L 1.1.1.1/32 is directly connected, FastEthernet4

10.0.0.0/32 is subnetted, 2 subnets

**O 10.10.0.1 [110/2] via 1.1.1.2, 02:25:53, FastEthernet4**

**O 10.10.100.1 [110/2] via 1.1.1.2, 02:25:53, FastEthernet4**

172.16.0.0/16 is variably subnetted, 4 subnets, 2 masksC 172.16.1.0/24 is directly connected, Loopback1

L 172.16.1.1/32 is directly connected, Loopback1

C 172.16.100.0/24 is directly connected, Loopback100

L 172.16.100.1/32 is directly connected, Loopback100

From the routing table above (before redistribution), currently there are routes for the local networks emulated by loopbacks 1&100 and also two routes learned via OSPF for the networks behind ENT2. (The blue color routes above denoted by “O” are learned via OSPF).

__Now, let’s see the routing table of ENT2 before the redistribution__

This router is the junction between the two networks and runs both OSPF and EIGRP. Therefore, its routing table contains the routes from both networks. However, we have not configured Redistribution yet.

**Codes: C – connected, D – EIGRP, O – OSPF,
ENT2#sh ip route**

ENT2#sh ip route

1.0.0.0/30 is subnetted, 1 subnets

C 1.1.1.0 is directly connected, FastEthernet1

2.0.0.0/30 is subnetted, 1 subnets

C 2.2.2.0 is directly connected, FastEthernet0

172.16.0.0/32 is subnetted, 2 subnets

**O 172.16.1.1 [110/2] via 1.1.1.1, 02:38:35, FastEthernet1**

**O 172.16.100.1 [110/2] via 1.1.1.1, 02:40:11, FastEthernet1**

10.0.0.0/24 is subnetted, 2 subnets

C 10.10.0.0 is directly connected, Loopback0

C 10.10.100.0 is directly connected, Loopback100

**D 192.168.1.0/24 [90/156160] via 2.2.2.2, 00:00:06, FastEthernet0**

**D 192.168.100.0/24 [90/156160] via 2.2.2.2, 00:00:06, FastEthernet0**

The blue color routes denoted by “**O**” are learned via OSPF (from the BLUE network) and are received from the neighbor router ENT1. The red color routes denoted by “**D**” are learned via EIGRP and are received from the RED network (from “Ent_Remote”).

__Let’s examine the “ENT_Remote” routing table:__ 

**ENT_Remote#sh ip route**2.0.0.0/8 is variably subnetted, 2 subnets, 2 masks

C 2.2.2.0/30 is directly connected, GigabitEthernet0

L 2.2.2.2/32 is directly connected, GigabitEthernet0

10.0.0.0/24 is subnetted, 2 subnets

**D 10.10.0.0 [90/156160] via 2.2.2.1, 00:06:11, GigabitEthernet0**

**D 10.10.100.0 [90/156160] via 2.2.2.1, 00:06:11, GigabitEthernet0**

192.168.1.0/24 is variably subnetted, 2 subnets, 2 masks

C 192.168.1.0/24 is directly connected, Loopback1

L 192.168.1.1/32 is directly connected, Loopback1

192.168.100.0/24 is variably subnetted, 2 subnets, 2 masks

C 192.168.100.0/24 is directly connected, Loopback100

L 192.168.100.1/32 is directly connected, Loopback100

The ENT_Remote device above has only EIGRP routes received from ENT2.

Our goal is to have a full converged network and every router needs to learn the routes to all the other networks. To achieve this, redistribution must be performed on **ENT2** router that runs both protocols.

When redistributing routes from other IGP’s into EIGRP you have to keep in mind that you **must specify the metric for the redistributed routes** which can be done in two ways:

1) OPTION1: Specify default metric for all the routes and issue redistribution command. All the routes that are redistributed will have that default metric.

__ENT2 Router__

**redistribute ospf 1**< —–Redistribute OSPF into EIGRP

network 2.2.2.0 0.0.0.3

network 10.10.0.0 0.0.0.255

network 10.10.100.0 0.0.0.255

**default-metric 1000 33 255 1 1500**< — Set the default metric for all redistributed routes

no auto-summary

eigrp router-id 2.2.2.1

2) OPTION2: Set the metric per each redistribute command

__ENT2 Router__

**redistribute ospf 1 metric 1000 33 255 1 1500**<- Set the metric for the specific OSPF 1 instance

network 2.2.2.0 0.0.0.3

network 10.10.0.0 0.0.0.255

network 10.10.100.0 0.0.0.255

no auto-summary

eigrp router-id 2.2.2.1

__NOTE:__

**metric 1000 33 255 1 1500 –** This command sets the K values that compose the metric

bandwidth to 1000 (Kbps), the delay to 33(tens-of-microseconds, or 330 microseconds), reliability to 255 (a value between 1–255 ,255 is best), load to 1 (a value between 1–255, 1 is best), and MTU of 1500.

The routing table on router “Ent_Remote” has the redistributed OSPF routes shown as

**D EX** (EIGRP external) as shown below.

**Ent_Remote#sh ip route eigrp**

1.0.0.0/30 is subnetted, 1 subnets

**D EX     1.1.1.0 [170/2571008] via 2.2.2.1, 00:01:29, GigabitEthernet0**

10.0.0.0/24 is subnetted, 2 subnets

D       10.10.0.0 [90/156160] via 2.2.2.1, 00:01:29, GigabitEthernet0

D       10.10.100.0 [90/156160] via 2.2.2.1, 00:01:29, GigabitEthernet0

172.16.0.0/32 is subnetted, 2 subnets

**D EX     172.16.1.1 [170/2571008] via 2.2.2.1, 00:01:29, GigabitEthernet0**

 **D EX     172.16.100.1 [170/2571008] via 2.2.2.1, 00:01:29, GigabitEthernet0**

