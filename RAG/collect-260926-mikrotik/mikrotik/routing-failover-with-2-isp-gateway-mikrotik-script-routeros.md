---
id: collect-260926-mikrotik/mikrotik/routing-failover-with-2-isp-gateway-mikrotik-script-routeros
title: "Routing FailOver with 2 ISP (gateway) - MikroTik Script RouterOS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/routing-failover-with-2-isp-gateway-mikrotik-script-routeros.md
source_anchor: ""
source_lines: [1, 11]
sha256: 207e594dba46bf31323df8c974fd4f3ecf3c2770296188a977c24e1ec0942fda
---

# Routing FailOver with 2 ISP (gateway) - MikroTik Script RouterOS

Using more than one internet gateway (ISP) allows us to fail over where one link can be used as the main gateway and the other becomes the backup link. For these needs, the usual configuration is to define a check-gateway and differentiate the distance value for each routing rule.
WAN-1 = IP 192.168.2.2/24
WAN-2 = IP 192.168.3.2/24
```
/ip route
add check-gateway=ping distance=1 gateway=192.168.2.1 target-scope=10
add check-gateway=ping distance=2 gateway=192.168.3.1 target-scope=10
```
Credit: www.o-om.com
