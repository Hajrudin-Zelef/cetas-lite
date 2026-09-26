---
id: collect-260926-mikrotik/mikrotik/questions-984184-mikrotik-lots-of-tcp-retransmission-packets-28ac6292
title: "DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-984184-mikrotik-lots-of-tcp-retransmission-packets-28ac6292.md
source_anchor: ""
source_lines: [1, 18]
sha256: 369818efafaffed6380534d050ad04f79bb11b7f7e75b7df62c53c381efd0ccb
---

# DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE

I have a ubuntu server with ip 192.168.10.144, in this server I have a docker network using ip range 10.0.0.0/24. I need connect my computer to some services running in docker, so I've added a route in Mikrotik:
#      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
0 ADS  0.0.0.0/0                          192.168.0.1               1
1 A S  10.0.0.0/24                        192.168.10.144            1
2 ADC  192.168.0.0/24     192.168.0.3     ether1-internet           0
3 ADC  192.168.10.0/24    192.168.10.1    ether2-proliant           0
However the connection between my computer and docker services are extremely slow. I used wireshark to check what was happening and found some of tcp-restransmission messages:
I've though a nat rule could solve the problem, however when using the following rules, no communication was made between my computer and docker:
0    ;;; default configuration
  chain=srcnat action=masquerade out-interface=ether1-internet log=no 
  log-prefix="" 
1    chain=dstnat action=dst-nat to-addresses=192.168.10.144 to-ports=80 
  protocol=tcp dst-port=8000 log=no log-prefix="" 
2 X  chain=dstnat action=dst-nat to-addresses=192.168.10.144 protocol=tcp 
  dst-address=10.0.0.0/24 log=no log-prefix="" 
3 X  chain=srcnat action=src-nat to-addresses=10.0.0.0/24 protocol=tcp 
  src-address=192.168.10.144 log=no log-prefix="" 
What am I doing wrong?
