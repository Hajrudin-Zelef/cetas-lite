---
id: collect-260926-mikrotik/mikrotik/firewall-rule-ordering-best-practices-other-questions
title: "firewall-rule-ordering-best-practices-other-questions"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/firewall-rule-ordering-best-practices-other-questions.md
source_anchor: ""
source_lines: [1, 68]
sha256: dd4077bd925ca06ce3072545a8c11fa380de87887766eee33d0a4ed4c2b9aa65
---

# firewall-rule-ordering-best-practices-other-questions

Network Info / Goals

4 Segments, all with their own VLAN, DHCP, WiFI etc:

VLAN 99 - CORE - 192.168.1.0/24 (This is the original one)

VLAN 100 - HOME - 192.168.100.0/24

VLAN 110 - GUEST - 192.168.110.0/24

VLAN 120 - PRIV - 192.168.1.120/24 (this happens to have all traffic routed via a Wireguard tunnel (WG list))

```
add interface=CORE list=LAN
add interface=HOME list=LAN
add interface=PRIV list=LAN
add interface=GUEST list=LAN
add interface=PRIV list=SAFE-MGMT
add interface=ether5 list=SAFE-MGMT
```

1 interface = 1 VLAN = 1 subnet in this config, and is the approach I would always take. The only exception is that I also have an off-bridge port, ether5 - 192.168.250.0/24, this is purely emergency management access.

Simple goals:

- PRIV and ether5 to be able to reach the router for management access
- LAN interface list to have Internet access
- Allow specific services (only DNS and ICMP at the moment) on the Input chain from the LAN list
- VLANs within the LAN list should not be able to communicate with eachother by default (I may add individual exceptions at a later date)
- PRIV VLAN can talk to everyone, though
- Implicit drop rule on all chains to catch any mistakes or errors (I prefer default deny, probably because I cut my teeth on Cisco ASAs)

I am a but unsure about the ordering of the Forward chain rules in particular; does the position of the Fasttrack rule matter in this context? is there a general best practice here? And what about the position of the rule for ‘established,related,untracked’?

I try to keep things clean and go with the approach of:

- Use interface lists when I want all members of all those subnets to have the same setting (in my case this is mainly for Internet access)
- Specific Interface-In when it’s just for that segment (as opposed to using a subnet/IP list)
- Individual IP (or IP in general) only if it needs to be more granular

```
/ip firewall filter
;input chain
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="Safe Management Access" in-interface-list=SAFE-MGMT
add action=accept chain=input comment="Allowed UDP Services to Router" dst-port=53 in-interface-list=LAN protocol=udp
add action=accept chain=input comment="Allowed TCP Services to Router" dst-port=53 in-interface-list=LAN protocol=tcp
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=accept chain=input comment="Allow ICMP to Input chain from LAN Segments" in-interface-list=LAN protocol=icmp
add action=drop chain=input comment="Implicit Drop Input Chain" log-prefix=INPUT-NOT-LAN
;forward chain
add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=forward comment="PRIV to other Segments" in-interface=PRIV out-interface-list=LAN
add action=accept chain=forward comment="Internet Access for All" connection-state=new in-interface-list=LAN out-interface-list=WAN
add action=accept chain=forward comment="Applicable access for Wireguard Tunnel" in-interface=GUEST out-interface-list=WG
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related hw-offload=yes
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="Implicit Drop Forward Chain" log=yes log-prefix=IMPLICIT
```

Anything here I can tweak to make it more performant, logical or secure? I am all about least privilege.

As an extra question, what is the best practice approach for ip service and system users when it comes to the ‘available from’ setting? If you have already firewalled it off as above, do you leave that setting blank? I have still set it to match the subnets explicity, but not sure if it’s actually adding any value. I guess it is belt and braces.

With this part done, I have a nice, clean and segmented network on my hAP AX3 

Thank you!
