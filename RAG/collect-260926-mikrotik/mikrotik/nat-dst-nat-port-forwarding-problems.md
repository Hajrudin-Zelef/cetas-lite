---
id: collect-260926-mikrotik/mikrotik/nat-dst-nat-port-forwarding-problems
title: "nat-dst-nat-port-forwarding-problems"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/nat-dst-nat-port-forwarding-problems.md
source_anchor: ""
source_lines: [1, 59]
sha256: 6967cf7e7e8f3298f9f7c563e85048604e66186607b427fb57a2d104b1cc8066
---

# nat-dst-nat-port-forwarding-problems

Some additional information;

I added a log into my forward rule and I can see that the packet seems to arrive at the Routerboard

19:05:09 firewall,info dstnat: in:ether1 out:(none), src-mac 68:b6:fc:d6:bc:02, proto TCP (SYN), 1.2.3.4:59866->192.168.0.10:80, len 48

19:05:27 firewall,info dstnat: in:ether1 out:(none), src-mac 68:b6:fc:d6:bc:02, proto TCP (SYN), 1.2.3.4:59866->192.168.0.10:80, len 48

19:05:56 firewall,info dstnat: in:ether1 out:(none), src-mac 68:b6:fc:d6:bc:02, proto TCP (SYN), 1.2.3.4:59913->192.168.0.10:80, len 48

19:06:14 firewall,info dstnat: in:ether1 out:(none), src-mac 68:b6:fc:d6:bc:02, proto TCP (SYN), 1.2.3.4:59913->192.168.0.10:80, len 48

1.2.3.4 is the public IP from the ISP at another endpoint. Means, flow is ok ingoing, but outgoing there seems a hang.

Full list of my ruleset is here:

for NAT:

0    chain=srcnat action=masquerade out-interface=VPN log=no

1    ;;; defconf: masquerade

chain=srcnat action=masquerade out-interface=ether1

2    chain=dstnat action=dst-nat to-addresses=192.168.5.70 to-ports=80 protocol=tcp dst-address=192.168.0.10 in-interface=ether1 dst-port=80 log=yes

for all other (default in 6.38):

0  D ;;; special dummy rule to show fasttrack counters

chain=forward action=passthrough

1    ;;; defconf: accept ICMP

chain=input action=accept protocol=icmp

2    ;;; defconf: accept established,related

chain=input action=accept connection-state=established,related

3    ;;; defconf: drop all from WAN

chain=input action=drop in-interface=ether1 log=no

4    ;;; defconf: fasttrack

chain=forward action=fasttrack-connection connection-state=established,related

5    ;;; defconf: accept established,related

chain=forward action=accept connection-state=established,related

6    ;;; defconf: drop invalid

chain=forward action=drop connection-state=invalid

7    ;;; defconf:  drop all from WAN not DSTNATed

chain=forward action=drop connection-state=new connection-nat-state=!dstnat in-interface=ether1
