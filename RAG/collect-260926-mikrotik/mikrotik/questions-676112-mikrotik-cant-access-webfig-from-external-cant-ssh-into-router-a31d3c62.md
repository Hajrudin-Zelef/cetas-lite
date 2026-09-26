---
id: collect-260926-mikrotik/mikrotik/questions-676112-mikrotik-cant-access-webfig-from-external-cant-ssh-into-router-a31d3c62
title: "questions-676112-mikrotik-cant-access-webfig-from-external-cant-ssh-into-router--a31d3c62"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-676112-mikrotik-cant-access-webfig-from-external-cant-ssh-into-router--a31d3c62.md
source_anchor: ""
source_lines: [1, 30]
sha256: 4d1493677fd8d86f839bd187a33bc5c809a8e0af6f5281b3cb2c38f458880acb
---

# questions-676112-mikrotik-cant-access-webfig-from-external-cant-ssh-into-router--a31d3c62

Ok,
here's the deal.
Let's say that my public facing IP is 10.0.01. I can't webfig into 10.0.01 from external and can't SSH into mikrotik router from external IP.
I can do it if I am physically connected to the router (on the same lan). It allows me to webfig and SSH both either by using 10.0.0.1 or using 192.168.88.1.
However, if I'm on diferent LAN, can't connect.
BTW, I set my IP > Services > ports for webfig is 64291 and SSH is 23.
Here's my firewall rules and NAT.
RULES
0 ;;; ALLOW ALL TO LAN
chain=input action=accept connection-state=established,related,new in-interface=bridge-local log=no log-prefix=""
1 ;;; ALLOW ICMP (Ping) ON ALL
chain=input action=accept protocol=icmp log=no log-prefix=""
2 ;;; Drop Everything Else
chain=input action=drop log=no log-prefix=""
3 ;;; default configuration
chain=input action=drop in-interface=ether1-gateway log=no log-prefix=""
4 ;;; default configuration
chain=forward action=accept connection-state=established,related,new in-interface=bridge-local log=no log-prefix=""
5 ;;; default configuration
chain=forward action=drop connection-state=invalid log=no log-prefix=""
6 ;;; default configuration
chain=forward action=drop connection-state=new connection-nat-state=!dstnat in-interface=ether1-gateway log=no log-prefix=""
NAT
0 ;;; default configuration
chain=srcnat action=masquerade to-addresses=X.X.X.X out-interface=ether1-gateway     log=no log-prefix=""
1 chain=dstnat action=dst-nat to-addresses=192.168.88.200 protocol=tcp dst-address=X.X.X.X dst-port=80 log=no log-prefix=""
2 chain=srcnat action=src-nat to-addresses=X.X.X.X protocol=tcp src-address=192.168.88.0/24 log=no log-prefix=""
3 chain=dstnat action=dst-nat to-addresses=192.168.88.200 to-ports=22 protocol=tcp dst-address=X.X.X.X dst-port=22 log=no log-prefix=""
4 chain=srcnat action=src-nat to-addresses=192.168.88.200 to-ports=22 protocol=tcp src-address=192.168.88.0/24 log=no log-prefix=""
5 chain=dstnat action=dst-nat to-addresses=192.168.88.1 protocol=tcp dst-address=X.X.X.X dst-port=23 log=no log-prefix=""
