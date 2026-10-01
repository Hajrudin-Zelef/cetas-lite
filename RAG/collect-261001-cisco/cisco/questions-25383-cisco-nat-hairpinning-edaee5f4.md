---
id: collect-261001-cisco/cisco/questions-25383-cisco-nat-hairpinning-edaee5f4
title: "questions-25383-cisco-nat-hairpinning-edaee5f4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-cisco/questions-25383-cisco-nat-hairpinning-edaee5f4.md
source_anchor: ""
source_lines: [1, 32]
sha256: a804fed5d13128e707e43b7ed33ae53531a447498071da20f2fcf50b87981657
---

# questions-25383-cisco-nat-hairpinning-edaee5f4

NVI NAT's already been brought up by Aaron D. 
Here's a the relevant config bits of a working example. It's been done on a CISCO881 with IOS 15.4(3)M6a
Outside network: 172.19.31.0 /24   on  FastEthernet4
Inside network:  172.19.140.0 /23  on VLAN141/SVI141
exposed host:    172.19.141.24
external port:   2222
internal port:   22
Interface configuration:
interface FastEthernet4
 ip address 172.19.31.2 255.255.255.0
 ip nat enable
interface Vlan141
 ip address 172.19.140.1 255.255.254.0
 !
 ! hairpinning did not work until ip redirects were disabled
 !
 no ip redirects
 ip nat enable
NAT ACL:
ip access-list standard ACLv4_SUBNET141
 permit 172.19.140.0 0.0.1.255
NAT rules:
ip nat source static tcp 172.19.141.24 22 interface FastEthernet4 2222
ip nat source list ACLv4_SUBNET141 interface FastEthernet4 overload
In a nutshell: 
- set the relevant interfaces to "ip nat enable" instead of "ip nat
in/outside", and slightly modify the NAT rules.
- make sure that there is an NVI NAT style outbound policy, or the hairpinnable host won't be able to connect outbound or hairpin to itself.   
- disable ip redirects on the "inside" interface, or hairpinning (at least not
from the host itself) will not work.
Caution: NVI NAT can be VERY taxing on the CPU of low-end routers like the 800 series. Where my old 881 used to be able to deliver 50-60Mbit/s with classic NAT, switching over to NVI caused the throughput to drop to 20-30Mbit/s and would have the CPU glowing red when under load. 
That was also the case when the to-be-hairpinned translation was not actually in use, just with traffic matching the normal "interface ... overload" outbound NAT rule.
