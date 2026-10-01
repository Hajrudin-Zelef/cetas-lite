---
id: collect-261001-fortinet/fortinet/questions-1871-fortigate-ha-how-to-locate-switch-port-connected-to-passive-unit-076f6c2b
title: "questions-1871-fortigate-ha-how-to-locate-switch-port-connected-to-passive-unit-076f6c2b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-1871-fortigate-ha-how-to-locate-switch-port-connected-to-passive-unit-076f6c2b.md
source_anchor: ""
source_lines: [1, 17]
sha256: e9001c8a6b6427b096eccf928db00c418d6f26bb4bf8cee21c151d923956a1d0
---

# questions-1871-fortigate-ha-how-to-locate-switch-port-connected-to-passive-unit-076f6c2b

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
In a new customer's network I have found the following scenario:
Fortigate 1 and 2 form and HA cluster in active-passive mode. The HA link is just a cable connecting them directly.
Racks A and B are a several meters apart and cables between them run through the ceiling. It is not possible to follow them and they are not identified: no pannels, no labels, ...
Port 1 of each Fortigate is connected to a switch but customer doesn't know the switch ports.
All switch ports are access ports in the default VLAN. Fortigate port 1 is default gateway for the subnet associated to this VLAN
In order to locate the switch ports, I considered connecting my PC in one switch, setting up an address of the subnet and ping the default gateway. Then check ARP looking for the MAC address of the default gateway and finally search for this MAC address in the CAM of both switches. However, I realized this will make me discover only the switch port connected to the active Fortigate unit.
Is there a way to track down the switch port connected to the passive Fortigate unit, without disconnecting the cable on firewall side, neither forcing a transition of the cluster active unit?
There is no direct way to determine the slave's MAC of port1 as both HA cluster member units share identical virtual MACs for each port.
Disconnecting (or disabling) a port has a 50% chance of triggering a cluster failover, with subsequently dropping some traffic.
To minimize network interruption I'd disable port monitoring for port1 temporarily, in the HA config setup. Then pull one port1 line and observe which switch port signals a link down. This way, traffic will only be dropped for the duration of this action and not for the (longer) period of one or possibly two HA cluster failovers.
A different approach would be to observe how much traffic flows across each switch port. The slave's wan port will not carry much traffic so the traffic rate on both switchports should be different by a large extent.
