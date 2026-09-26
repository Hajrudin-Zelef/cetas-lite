---
id: collect-260926-mikrotik/mikrotik/load-balancing-mikrotik-pcc-dengan-pppoe
title: "Load Balancing Mikrotik PCC dengan PPPoe"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/load-balancing-mikrotik-pcc-dengan-pppoe.md
source_anchor: ""
source_lines: [1, 100]
sha256: 5ac9f760967cd494528d0ee6f373f418c34c9ec4d201f5508e1ac4b65635adae
---

# Load Balancing Mikrotik PCC dengan PPPoe

Mikrotik tutorial Load Balance Mikrotik with PCC using PPPoe-Client connection is used for bandwitdh optimalization,fail over, and bandwidth control usage.

First, make new mangle for mark connectio, packet and routing :

/ ip firewall mangle

add chain=input in-interface=internet1 action=mark-connection new-connection-mark=internet1_conn

add chain=input in-interface=internet2 action=mark-connection new-connection-mark=internet2_conn

add chain=output connection-mark=internet1_conn action=mark-routing new-routing-mark=to_internet1

add chain=output connection-mark=internet2_conn action=mark-routing new-routing-mark=to_internet2

add chain=prerouting dst-address-type=!local in-interface=Local per-connection-classifier=both-addresses:2/0 \

action=mark-connection new-connection-mark=internet1_conn passthrough=yes

add chain=prerouting dst-address-type=!local in-interface=Local per-connection-classifier=both-addresses:2/1 \

action=mark-connection new-connection-mark=internet2_conn passthrough=yes

add chain=prerouting connection-mark=internet1_conn in-interface=Local action=mark-routing new-routing-mark=to_internet1

add chain=prerouting connection-mark=internet2_conn in-interface=Local action=mark-routing new-routing-mark=to_internet2


Than set the default route and fail over gateway :

/ ip route

add dst-address=0.0.0.0/0 gateway=10.111.0.1 routing-mark=to_internet1 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.112.0.1 routing-mark=to_internet2 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.111.0.1 distance=1 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.112.0.1 distance=2 check-gateway=ping


Last, set the NAT for local connection:

/ ip firewall nat

add chain=srcnat out-interface=internet1 action=masquerade

add chain=srcnat out-interface=internet2 action=masquerade


That is simple script for Load Balance Mikrotik with PCC using PPPoe-Client connection.

First, make new mangle for mark connectio, packet and routing :

/ ip firewall mangle

add chain=input in-interface=internet1 action=mark-connection new-connection-mark=internet1_conn

add chain=input in-interface=internet2 action=mark-connection new-connection-mark=internet2_conn

add chain=output connection-mark=internet1_conn action=mark-routing new-routing-mark=to_internet1

add chain=output connection-mark=internet2_conn action=mark-routing new-routing-mark=to_internet2

add chain=prerouting dst-address-type=!local in-interface=Local per-connection-classifier=both-addresses:2/0 \

action=mark-connection new-connection-mark=internet1_conn passthrough=yes

add chain=prerouting dst-address-type=!local in-interface=Local per-connection-classifier=both-addresses:2/1 \

action=mark-connection new-connection-mark=internet2_conn passthrough=yes

add chain=prerouting connection-mark=internet1_conn in-interface=Local action=mark-routing new-routing-mark=to_internet1

add chain=prerouting connection-mark=internet2_conn in-interface=Local action=mark-routing new-routing-mark=to_internet2

Than set the default route and fail over gateway :

/ ip route

add dst-address=0.0.0.0/0 gateway=10.111.0.1 routing-mark=to_internet1 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.112.0.1 routing-mark=to_internet2 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.111.0.1 distance=1 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.112.0.1 distance=2 check-gateway=ping

Last, set the NAT for local connection:

/ ip firewall nat

add chain=srcnat out-interface=internet1 action=masquerade

add chain=srcnat out-interface=internet2 action=masquerade

That is simple script for Load Balance Mikrotik with PCC using PPPoe-Client connection.

## Diskusi
