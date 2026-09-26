---
id: collect-260926-mikrotik/mikrotik/load-balancing-lb-pcc-2-line-isp-failover-mikrotik-script-routeros
title: "Load Balancing LB PCC 2 Line ISP Failover - Mikrotik Script RouterOS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2021-13-02"]
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/load-balancing-lb-pcc-2-line-isp-failover-mikrotik-script-routeros.md
source_anchor: ""
source_lines: [1, 33]
sha256: ac503d804ad8f24d8244489f4e72c9bff3eb06fa473b48adb7900d3fa1814241
---

# Load Balancing LB PCC 2 Line ISP Failover - Mikrotik Script RouterOS

Load balance PCC 2 ISP on mikrotik is a technique for distributing traffic load on two or more connection lines in a balanced way, so that traffic can run optimally, maximize throughput, reduce response time and avoid overloading one of the connection lines.
```
################################################
# LOAD BALANCING (LB) PCC SCRIPT GENERATOR
# Date/Time: 2/13/2021, 9:19:38 PM
# https://fb.me/buananet.pbun
# Load Balancing Metode -> PCC
################################################
/ip firewall address-list
add address=192.168.0.0/16 list=LOCAL-IP
add address=172.16.0.0/12 list=LOCAL-IP
add address=10.0.0.0/8 list=LOCAL-IP
/ip firewall nat
add chain=srcnat out-interface="ether1" action=masquerade
add chain=srcnat out-interface="ether2" action=masquerade
/ip route
add check-gateway=ping distance=1 gateway="192.168.1.1" routing-mark="to-ether1"
add check-gateway=ping distance=1 gateway="192.168.2.1" routing-mark="to-ether2"
add check-gateway=ping distance=1 gateway="192.168.1.1"
add check-gateway=ping distance=2 gateway="192.168.2.1"
/ip firewall mangle
add action=mark-connection chain=input in-interface="ether1" new-connection-mark="cm-ether1" passthrough=yes
add action=mark-connection chain=input in-interface="ether2" new-connection-mark="cm-ether2" passthrough=yes
add action=mark-routing chain=output connection-mark="cm-ether1" new-routing-mark="to-ether1" passthrough=yes
add action=mark-routing chain=output connection-mark="cm-ether2" new-routing-mark="to-ether2" passthrough=yes
add action=mark-connection chain=prerouting dst-address-list=!LOCAL-IP dst-address-type=!local new-connection-mark="cm-ether1" passthrough=yes per-connection-classifier=both-addresses-and-ports:2/0 src-address-list=LOCAL-IP
add action=mark-connection chain=prerouting dst-address-list=!LOCAL-IP dst-address-type=!local new-connection-mark="cm-ether2" passthrough=yes per-connection-classifier=both-addresses-and-ports:2/1 src-address-list=LOCAL-IP
add action=mark-routing chain=prerouting connection-mark="cm-ether1" dst-address-list=!LOCAL-IP new-routing-mark="to-ether1" passthrough=yes src-address-list=LOCAL-IP
add action=mark-routing chain=prerouting connection-mark="cm-ether2" dst-address-list=!LOCAL-IP new-routing-mark="to-ether2" passthrough=yes src-address-list=LOCAL-IP
```
Credit: https://www.o-om.com/2020/12/load-balancing-pcc-script-generator-for.html
