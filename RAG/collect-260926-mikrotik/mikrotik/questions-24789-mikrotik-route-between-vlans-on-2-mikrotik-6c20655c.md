---
id: collect-260926-mikrotik/mikrotik/questions-24789-mikrotik-route-between-vlans-on-2-mikrotik-6c20655c
title: "questions-24789-mikrotik-route-between-vlans-on-2-mikrotik-6c20655c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-24789-mikrotik-route-between-vlans-on-2-mikrotik-6c20655c.md
source_anchor: ""
source_lines: [1, 57]
sha256: 99d5c002a4645580bd19e1b38458cb53407ff12e16221f4987713da440f9de3b
---

# questions-24789-mikrotik-route-between-vlans-on-2-mikrotik-6c20655c

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Here's the scenario:
There are 2 Mikrotiks, each are on different floor
Each ether 1 port connected to ISP through PPPoE
Both Mikrotiks connected to each other on port 3
There's a VOIP server on IP address 192.168.121.2 on Mikrotik 1
Mikrotik 1:
Ether 1 is out to internet
Ether 3 is connected to Mikrotik 2 and has an IP address of 192.168.30.1
Ether 4 is a master of Ether 5, and both ports connect to 2 different
switches which are used for IP phones on VLAN 103 (192.168.114.0/24) and
the local network on VLAN 102 (192.168.99.0/24)
Ether 6 (192.168.60.1) goes to the 6th floor which is only used for the Internet, which is connected to Ether 1
Ether 9 is the port to which the VoIP server (192.168.121.2) is connected, and Ether 9 has an IP address of 192.168.121.1/24
Mikrotik 2:
Ether 1 is out to another ISP for the Internet
Ether 3 is connected to Mikrotik 1 (192.168.30.2)
Ether 4 (192.168.110.0/24) is connected to a switch for use in this
floor and has 2 VLANs (103 and 104)
VLAN 103 = IP address 192.168.112.0/24 is used for the VoIP network
VLAN 104 = IP address 192.168.111.0/24 is used for local network and
Wi-Fi
My problems are:
VLAN 103 in Mikrotik 2 can't connect/route to VoIP server
(192.168.121.2) on Ether 9 in Mikrotik 1.
I need Ether 9 (VoIP server 192.168.121.2) in Mikrotik 1 to route to
VLAN 103 (192.168.112.0) in Mikrotik 2.
How do i route traffic from VLAN 102 in Mikrotik 1 to
Ether 3 (Mikrotik 1) and down to Ether 3 in Mikrotik 2 to the ISP which
is connected on Ether 1 in Mikrotik 2?
From the existing configuration, both Ether 1 on the Mikrotiks
connect to the ISPs using address acquisition of static instead of PPPoE
which is the IP address of the modem but both have PPPoE clients configured with
PPPoE usernames and passwords. Neither router can operate in
bridge mode, and configurations on both modems don't have PPPoE
usernames and passwords. How does it work?
First of all, a good network diagram will always make it easier for others to understand what you are trying to say. Even a quick little diagram is better than nothing.
I made this in an attempt to understand what you were talking about:
VLAN 103 in Mikrotik 2 can't connect/route to VoIP server (192.168.121.2) on Ether 9 in Mikrotik 1.
I need Ether 9 (VoIP server 192.168.121.2) in Mikrotik 1 to route to VLAN 103 (192.168.112.0) in Mikrotik 2.
How do i route traffic from VLAN 102 in Mikrotik 1 to Ether 3 (Mikrotik 1) and down to Ether 3 in Mikrotik 2 to the ISP which is connected on Ether 1 in Mikrotik 2?
Stop thinking about the VLANs when you are talking about routing. A router does not care what VLAN a subnet is on until it needs to exit a VLAN trucked interface. While the router is routing between subnets, it just wants to know the subnet IP addresses.
Without more information it is hard to know what the problem is, but it is most likely a routing problem, just like @Nick B. mentioned.
Each of the subnets on MikroTik 1 have to have a route on MikroTik 2, and each of the subnets on MikroTik 2 have to have a route on MikroTik 1.
You can make this easier by enabling a routing protocol on both routers.
Well, i would like to know do i need a vlan trunk here in both mikrotik? If so which port should i make it a trunk? A port that has vlan which connected to the switch or the port between mikrotik (port 3 on each mikrotik)?
Trunks are only used between two switches, or between a switch and a router to create a "router on a stick".
I really don't understand how mikrotik router route the traffic from one vlan suppose from vlan 103(192.168.114.0/24) in mikrotik1 to vlan103(192.168.112.0/24) in mikrotik2 as both vlan id is same but the IP is different?
Since the two VLANs are on different routers with different subnet addresses, they aren't related at all. The traffic needs to be routed between the two routers.
In mikrotik, is it possible that any ip address or vlan can route to another IP class/Vlan in the same router?
Routers by default always create a route to directly connected subnets. So nothing should prevent two VLANs on the same router from communicating, unless the VLANs are configured incorrectly or firewall rules are blocking it.
As @Nick B. started at, do the following on the designated routers.
