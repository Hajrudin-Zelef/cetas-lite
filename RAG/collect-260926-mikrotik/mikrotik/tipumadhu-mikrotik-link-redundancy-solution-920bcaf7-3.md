---
id: collect-260926-mikrotik/mikrotik/tipumadhu-mikrotik-link-redundancy-solution-920bcaf7-3
title: "tipumadhu-mikrotik-link-redundancy-solution-920bcaf7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/tipumadhu-mikrotik-link-redundancy-solution-920bcaf7.md
source_anchor: ""
source_lines: [74, 78]
sha256: 90b31d380fae8ac2c8890acc9f06b39596e4cb3f1bf8bf501d0d86eb9b450ca0
---

# tipumadhu-mikrotik-link-redundancy-solution-920bcaf7

- 23. Mikrotik PCC LoadBalancing Configuration. /ip address add address=192.168.0.1/24 network=192.168.0.0 broadcast=192.168.0.255 interface=Local add address=192.168.1.2/24 network=192.168.1.0 broadcast=192.168.1.255 interface=WAN1 add address=192.168.2.2/24 network=192.168.2.0 broadcast=192.168.2.255 interface=WAN2 /ip firewall mangle add chain=input in-interface=WAN1 action=mark-connection new-connection-mark=WAN1_conn add chain=input in-interface=WAN2 action=mark-connection new-connection-mark=WAN2_conn add chain=output connection-mark=WAN1_conn action=mark-routing new-routing-mark=to_WAN1 add chain=output connection-mark=WAN2_conn action=mark-routing new-routing-mark=to_WAN2 add chain=prerouting dst-address=192.168.1.0/24 action=accept in-interface=Local add chain=prerouting dst-address=192.168.2.0/24 action=accept in-interface=Local add chain=prerouting dst-address-type=!local in-interface=Local per-connection-classifier=both-addresses-and- ports:2/0 action=mark-connection new-connection-mark=WAN1_conn passthrough=yes add chain=prerouting dst-address-type=!local in-interface=Local per-connection-classifier=both-addresses-and- ports:2/1 action=mark-connection new-connection-mark=WAN2_conn passthrough=yes add chain=prerouting connection-mark=WAN1_conn in-interface=Local action=mark-routing new-routing- mark=to_WAN1 add chain=prerouting connection-mark=WAN2_conn in-interface=Local action=mark-routing new-routing- mark=to_WAN2
- 24. /ip route add dst-address=0.0.0.0/0gateway=192.168.1.1 routing- mark=to_WAN1 check-gateway=ping add dst-address=0.0.0.0/0 gateway=192.168.2.1 routing- mark=to_WAN2 check-gateway=ping add dst-address=0.0.0.0/0 gateway=192.168.1.1 distance=1 check-gateway=ping add dst-address=0.0.0.0/0 gateway=192.168.2.1 distance=2 check-gateway=ping /ip firewall nat add chain=srcnat out-interface=WAN1 action=masquerade add chain=srcnat out-interface=WAN2 action=masquerade
- 25. PCC WITH UN-EQUALWAN LINKS If you have Un-Equal WAN Links, for example WAN,1 is of 4MB and WAN,2 is of 8 Mb, and you want to force MT to use WAN42link more then other because of its capacity, Then you have to Add more PCC rules assigning the same two marks to a specific link i.e WAN2 , something like add chain=prerouting dst-address-type=!local in-interface=Local per- connection-classifier=both-addresses-and-ports:2/0 action=mark- connection new-connection-mark=WAN1_conn passthrough=yes add chain=prerouting dst-address-type=!local in-interface=Local per- connection-classifier=both-addresses-and-ports:2/1 action=mark- connection new-connection-mark=WAN2_conn passthrough=yes add chain=prerouting dst-address-type=!local in-interface=Local per- connection-classifier=both-addresses-and-ports:2/2 action=mark- connection new-connection-mark=WAN2_conn passthrough=yes
- 26.
- 27.
