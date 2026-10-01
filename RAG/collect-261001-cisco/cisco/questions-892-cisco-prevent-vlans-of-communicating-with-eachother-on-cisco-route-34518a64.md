---
id: collect-261001-cisco/cisco/questions-892-cisco-prevent-vlans-of-communicating-with-eachother-on-cisco-route-34518a64
title: "questions-892-cisco-prevent-vlans-of-communicating-with-eachother-on-cisco-route-34518a64"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-892-cisco-prevent-vlans-of-communicating-with-eachother-on-cisco-route-34518a64.md
source_anchor: ""
source_lines: [1, 53]
sha256: 080272de64834fc6349901f6ad309e3fc49565350697b7cf0d704d7626a311f3
---

# questions-892-cisco-prevent-vlans-of-communicating-with-eachother-on-cisco-route-34518a64

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
10
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Setup: Cisco router with multiple VLAN's configured on it.
How can you prevent 2 VLAN's communicating with eachother? Normally I would do this with ACL's like this:
access-list 102 deny ip 1.1.1.0 0.0.0.255 2.2.2.0 0.0.0.255
access-list 102 deny ip 2.2.2.0 0.0.0.255 1.1.1.0 0.0.0.255
int vlan 1
ip address 1.1.1.1 255.255.255.0
access-group 102 in
int vlan 2
ip address 2.2.2.2 255.255.255.0
access-group 102 in
This however is not handy when dealing with lot's of VLAN's configured on a router. Any suggestions of tweaking this or using an alternative to improve scalability?
Fully agreed with Stefan. VRF is the way to go here. Quick example how to incorporate it to suggested config:
ip vrf VLAN1
rd 42:1
ip vrf VLAN2
rd 42:2
!
int vlan1
ip vrf forwarding VLAN1
ip address 1.1.1.1 255.255.255.0
int vlan2
ip vrf forwading VLAN2
ip address 2.2.2.2 255.255.255.0
!
Now vlan1 and vlan2 routing is separated.
To inspect routing tables, ping, traceroute you need to specify the vrf. e.g.:
ip route vrf VLAN1
traceroute vrf VLAN2 192.0.2.1
ping vrf VLAN2 192.0.2.1
Or same in new AFI aware, IPv6 supporting config:
vrf definition VLAN1
rd 42:1
address-family ipv4
vrf definition VLAN2
rd 42:2
address-family ipv4
!
int vlan1
vrf definition VLAN1
ip address 1.1.1.1 255.255.255.0
int vlan2
vrf definition VLAN2
ip address 2.2.2.2 255.255.255.0
!
While ACLs are a simple and safe way, it doesn't scale well indeed.
If your router provides VRF or at least the the VRF Lite feature you could group VLANs into VRFs. A VRF can be seen like a virtual router, VRF instances cannot talk to each other unless you explicitely define routing between them.
In a complex network, I group VLANs into several security domains done with VRFs, such as a VRF for office clients and server, a VRF for tech devices (door access control, lifts, cctv, ...), a VRF for guests and visitors.
