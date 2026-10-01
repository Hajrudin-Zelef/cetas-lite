---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-18949-block-inter-subnet-routing-ubiquiti-edge-router-f4c63680
title: "questions-18949-block-inter-subnet-routing-ubiquiti-edge-router-f4c63680"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-18949-block-inter-subnet-routing-ubiquiti-edge-router-f4c63680.md
source_anchor: ""
source_lines: [1, 22]
sha256: a8d71cf9044502b210fc8341ff39ecd29204289349e456449e79e343476d25d9
---

# questions-18949-block-inter-subnet-routing-ubiquiti-edge-router-f4c63680

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I would like to know how to block routing between subnets on my Ubiquiti EdgeRouter. These subnets are on 6 different VLANs, and the layout goes like this:
VLAN 1 - 10.0.0.1/24 (management)
VLAN 100 - 10.1.1.1/24
VLAN 200 - 10.2.1.1/24
VLAN 300 - 10.3.1.1/24 (guest network)
VLAN 400 - 10.4.1.1/24
VLAN 450 - 10.4.2.1/24 (guest network)
I want VLAN 1 to have access into all the other networks, since it is the management network, and access to it is heavily secured, but I do not want any of the other networks to be able to talk to each other at all.
I would also like to allow a group of admin computers access to any subnet either from VLAN 200's subnet, or ideally, from any subnet (based on MAC address I suppose..)
I feel like this should be fairly simple to implement via some form of ACLs or Firewall rules, but I do not have a strong enough grasp of these concepts to implement this at this point in time.
Chapter 5 of the manual here should get you started.
What you want to do is:
Allow traffic from your admin computers to anywhere in 10.0.0.0/8
Deny all traffic from 10.0.0.0/8 to 10.0.0.0/8
Permit all traffic to 0.0.0.0/0 (assuming those other vlans are allowed to reach the Internet).
Rules are processed in order and the first rule they hit will be the effective one so your management computers will "hit" the first rule and be allowed. Other computers trying to reach other vlans will "hit" the second rule. Traffic that is not destined to other vlans (source and destination not both within 10.0.0.0/8) will be allowed (internet traffic).
