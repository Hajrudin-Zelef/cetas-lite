---
id: collect-260926-mikrotik/mikrotik/questions-716969-mikrotik-dst-nat-timeout-6cc2eeef
title: "questions-716969-mikrotik-dst-nat-timeout-6cc2eeef"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-716969-mikrotik-dst-nat-timeout-6cc2eeef.md
source_anchor: ""
source_lines: [1, 5]
sha256: 0068c16bc02e9a3a59669664f2c980a37e59de5bec1b8f9810fc8d88ce3c5498
---

# questions-716969-mikrotik-dst-nat-timeout-6cc2eeef

I have installed a brand new Mikrotik RB2011UiAS-2HnD-IN in our Office. Everything is working except port forwarding :( I have disabled all other rules and filters but nothing helped . I am able to ping the client from Mikrotik . When I am trying to connect the port from out of office I see the that Statistics of dst-nat rule changes packets and bytes , but after few seconds it timed outs
/ip firewall nat
add action=dst-nat chain=dstnat dst-port=3389 in-interface=pppoe-out1 \
    protocol=tcp to-addresses=192.168.1.101 to-ports=3389
Is there anything other that I have to check ?
