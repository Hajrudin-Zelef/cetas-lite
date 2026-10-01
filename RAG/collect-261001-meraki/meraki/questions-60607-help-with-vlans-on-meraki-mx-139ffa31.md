---
id: collect-261001-meraki/meraki/questions-60607-help-with-vlans-on-meraki-mx-139ffa31
title: "questions-60607-help-with-vlans-on-meraki-mx-139ffa31"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-60607-help-with-vlans-on-meraki-mx-139ffa31.md
source_anchor: ""
source_lines: [1, 14]
sha256: bc032fa42da0e06e2a6fc38554e369adae9ac26ec2d49d73412182f7e55bbba2
---

# questions-60607-help-with-vlans-on-meraki-mx-139ffa31

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Meraki MX84 in Routed mode with a Default 10.10.0.0/24 subnet and then VLAN 10.10.3.0/24 ID 3, VLAN 10.10.4.0/24 ID 4, VLAN 10.10.5.0/24 ID 5, and 10.10.8.0/24 with ID 8. All ports on it are set to Native VLAN 1 and Trunk
I have a Cisco SG200. I've added the VLANs 3-5 and 8. All ports are set to Trunk.
The MX is in port 1 of the SG200 and I set the Administrative VLANs to 1 and 5.
I have PCs on 10.10.0.0 and 10.10.8.0 and then a server uses the other subnets and need them to connect with each other. There is no reason to keep the subnets separated.
As it is right now they can not talk to each other. Is there a step I'm missing? Should all ports on the SG200 be added to all VLANs?
*Port 5 on the MX is set to VLAN 8 just for testing.
Trunks are not appropriate for your access interfaces where you connect PCs, printers, etc. Most end-devices do not understand VLAN tags, so they will drop traffic from trunk interfaces.
What you want to do is configure a trunk between the router and switch, and between the switch and any other switches. For all the other interfaces, you configure access interfaces. Each access interface is configured for the particular VLAN to which the end-device connects.
