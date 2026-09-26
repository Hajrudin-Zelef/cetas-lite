---
id: collect-260926-mikrotik/mikrotik/load-balancing-mikrotik-pcc-3
title: "load-balancing-mikrotik-pcc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/load-balancing-mikrotik-pcc.md
source_anchor: ""
source_lines: [109, 159]
sha256: 9d358f698df293a0f55e4943ac836e17ee950f9d7973a0d7d29b37636ed2e607
---

# load-balancing-mikrotik-pcc

|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:4/2 \ action=mark-connection new-connection-mark=wlan3_conn passthrough=yes | 
|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:4/3 \ action=mark-connection new-connection-mark=wlan4_conn passthrough=yes | 
|  | add chain=prerouting connection-mark=wlan1_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan1 | 
|  | add chain=prerouting connection-mark=wlan2_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan2 | 
|  | add chain=prerouting connection-mark=wlan3_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan3 | 
|  | add chain=prerouting connection-mark=wlan4_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan4 | 
|  | / ip route | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.1.254 routing-mark=to_wlan1 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.2.254 routing-mark=to_wlan2 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.3.254 routing-mark=to_wlan3 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.4.254 routing-mark=to_wlan4 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.1.254 distance=1 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.2.254 distance=2 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.3.254 distance=3 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.4.254 distance=4 check-gateway=ping | 
|  | / ip firewall nat | 
|  | add chain=srcnat out-interface=wlan1 action=masquerade | 
|  | add chain=srcnat out-interface=wlan2 action=masquerade | 
|  | add chain=srcnat out-interface=wlan3 action=masquerade | 
|  | add chain=srcnat out-interface=wlan4 action=masquerade | 
|  | /ip firewall mangle | 
|  | add action=mark-connection chain=prerouting comment=LB connection-state=new \ | 
|  | in-interface="Eth 5 Distribusi" new-connection-mark=LB1 nth=4,1 protocol=\ | 
|  | tcp | 
|  | add action=mark-routing chain=prerouting connection-mark=LB1 in-interface=\ | 
|  | "Eth 5 Distribusi" new-routing-mark=route-lb1 passthrough=no | 
|  | add action=mark-connection chain=prerouting connection-state=new in-interface=\ | 
|  | "Eth 5 Distribusi" new-connection-mark=LB2 nth=4,2 protocol=tcp | 
|  | add action=mark-routing chain=prerouting connection-mark=LB2 in-interface=\ | 
|  | "Eth 5 Distribusi" new-routing-mark=route-lb2 passthrough=no | 
|  | add action=mark-connection chain=prerouting connection-state=new in-interface=\ | 
|  | "Eth 5 Distribusi" new-connection-mark=LB3 nth=4,3 protocol=tcp | 
|  | add action=mark-routing chain=prerouting connection-mark=LB3 in-interface=\ | 
|  | "Eth 5 Distribusi" new-routing-mark=route-lb3 passthrough=no | 
|  | add action=mark-connection chain=prerouting connection-state=new in-interface=\ | 
|  | "Eth 5 Distribusi" new-connection-mark=LB4 nth=4,4 protocol=tcp | 
|  | add action=mark-routing chain=prerouting connection-mark=LB4 in-interface=\ | 
|  | "Eth 5 Distribusi" new-routing-mark=route-lb4 passthrough=no | 
|  | /ip route | 
|  | add distance=1 gateway=###IP Static### routing-mark=to-publik | 
|  | add distance=1 gateway="Dial Speedy 1" routing-mark=route-lb1 | 
|  | add distance=1 gateway="Dial Speedy 2" routing-mark=route-lb2 | 
|  | add distance=1 gateway="Dial Speedy 3" routing-mark=route-lb3 | 
|  | add distance=1 gateway="Dial Speedy 4" routing-mark=route-lb4 | 
|  | add distance=1 gateway="Dial Speedy 1" | 
|  | add distance=2 gateway="Dial Speedy 2" | 
|  | add distance=3 gateway="Dial Speedy 3" | 
|  | add distance=4 gateway="Dial Speedy 4" | 
|  | add distance=1 dst-address=####/32 gateway=###iP Static### | 

gak pakai fail over ?
