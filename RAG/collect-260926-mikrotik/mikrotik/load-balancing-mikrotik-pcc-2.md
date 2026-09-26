---
id: collect-260926-mikrotik/mikrotik/load-balancing-mikrotik-pcc-2
title: "load-balancing-mikrotik-pcc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/load-balancing-mikrotik-pcc.md
source_anchor: ""
source_lines: [19, 108]
sha256: a1da912cce9d541d17497a1beb805febc7e96d51bce3690b7e6372e02653d1c1
---

# load-balancing-mikrotik-pcc

  |  | #loadbalance 2 interface | 
|  | / ip address | 
|  | add address=192.168.10.1/24 network=192.168.10.0 broadcast=192.168.10.255 interface=ether3 | 
|  | add address=192.168.9.9/24 network=192.168.9.0 broadcast=192.168.9.255 interface=ether1 | 
|  | add address=192.168.8.2/24 network=192.168.8.0 broadcast=192.168.8.255 interface=ether2 | 
|  | /ip firewall address-list | 
|  | add address=192.168.0.0/16 list=local | 
|  | add address=172.16.0.0/12 list=local | 
|  | add address=10.0.0.0/8 list=local | 
|  | /ip dns set allow-remote-requests=yes cache-max-ttl=1d cache-size=5000KiB servers=8.8.8.8,8.8.4.4 | 
|  | /ip firewall nat | 
|  | add action=redirect chain=dstnat dst-port=53 protocol=tcp to-ports=53 | 
|  | add action=redirect chain=dstnat dst-port=53 protocol=udp to-ports=53 | 
|  | /ip firewall filter | 
|  | add action=drop chain=input dst-port=53 in-interface=ether2 protocol=tcp | 
|  | add action=drop chain=input dst-port=53 in-interface=ether2 protocol=udp | 
|  | / ip firewall mangle | 
|  | add chain=input in-interface=ether1 action=mark-connection new-connection-mark=ether1_conn | 
|  | add chain=input in-interface=ether2 action=mark-connection new-connection-mark=ether2_conn | 
|  | add chain=output connection-mark=ether1_conn action=mark-routing new-routing-mark=to_ether1 | 
|  | add chain=output connection-mark=ether2_conn action=mark-routing new-routing-mark=to_ether2 | 
|  | add chain=prerouting dst-address=192.168.9.0/24 action=accept in-interface=ether3 | 
|  | add chain=prerouting dst-address=192.168.8.0/24 action=accept in-interface=ether3 | 
|  | add chain=prerouting dst-address-type=!local in-interface=ether3 per-connection-classifier=both-addresses:2/0 \ action=mark-connection new-connection-mark=ether1_conn passthrough=yes | 
|  | add chain=prerouting dst-address-type=!local in-interface=ether3 per-connection-classifier=both-addresses:2/1 \ action=mark-connection new-connection-mark=ether2_conn passthrough=yes | 
|  | add chain=prerouting dst-address-type=!local connection-mark=ether1_conn in-interface=ether3 action=mark-routing new-routing-mark=to_ether1 | 
|  | add chain=prerouting dst-address-type=!local connection-mark=ether2_conn in-interface=ether3 action=mark-routing new-routing-mark=to_ether2 | 
|  | / ip route | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.9.1 routing-mark=to_ether1 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.8.1 routing-mark=to_ether2 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.9.1 distance=1 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.8.1 distance=2 check-gateway=ping | 
|  | / ip firewall nat | 
|  | add chain=srcnat out-interface=ether1 action=masquerade | 
|  | add chain=srcnat out-interface=ether2 action=masquerade | 
|  | #loadbalance 3 interface | 
|  | / ip address | 
|  | add address=192.168.5.254/24 network=192.168.5.0 broadcast=192.168.5.255 interface=lan | 
|  | add address=192.168.1.1/24 network=192.168.1.0 broadcast=192.168.1.255 interface=wlan1 | 
|  | add address=192.168.2.1/24 network=192.168.2.0 broadcast=192.168.2.255 interface=wlan2 | 
|  | add address=192.168.3.1/24 network=192.168.3.0 broadcast=192.168.3.255 interface=wlan3 | 
|  | / ip firewall mangle | 
|  | add chain=input in-interface=wlan1 action=mark-connection new-connection-mark=wlan1_conn | 
|  | add chain=input in-interface=wlan2 action=mark-connection new-connection-mark=wlan2_conn | 
|  | add chain=input in-interface=wlan3 action=mark-connection new-connection-mark=wlan3_conn | 
|  | add chain=output connection-mark=wlan1_conn action=mark-routing new-routing-mark=to_wlan1 | 
|  | add chain=output connection-mark=wlan2_conn action=mark-routing new-routing-mark=to_wlan2 | 
|  | add chain=output connection-mark=wlan3_conn action=mark-routing new-routing-mark=to_wlan3 | 
|  | add chain=prerouting dst-address=192.168.1.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address=192.168.2.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address=192.168.3.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:3/0 \ action=mark-connection new-connection-mark=wlan1_conn passthrough=yes | 
|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:3/1 \ action=mark-connection new-connection-mark=wlan2_conn passthrough=yes | 
|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:3/2 \ action=mark-connection new-connection-mark=wlan3_conn passthrough=yes | 
|  | add chain=prerouting connection-mark=wlan1_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan1 | 
|  | add chain=prerouting connection-mark=wlan2_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan2 | 
|  | add chain=prerouting connection-mark=wlan3_conn in-interface=lan action=mark-routing new-routing-mark=to_wlan3 | 
|  | / ip route | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.1.254 routing-mark=to_wlan1 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.2.254 routing-mark=to_wlan2 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.3.254 routing-mark=to_wlan3 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.1.254 distance=1 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.2.254 distance=2 check-gateway=ping | 
|  | add dst-address=0.0.0.0/0 gateway=192.168.3.254 distance=3 check-gateway=ping | 
|  | / ip firewall nat | 
|  | add chain=srcnat out-interface=wlan1 action=masquerade | 
|  | add chain=srcnat out-interface=wlan2 action=masquerade | 
|  | add chain=srcnat out-interface=wlan3 action=masquerade | 
|  | #loadbalance 4 interface | 
|  | / ip address | 
|  | add address=192.168.5.254/24 network=192.168.5.0 broadcast=192.168.5.255 interface=lan | 
|  | add address=192.168.1.1/24 network=192.168.1.0 broadcast=192.168.1.255 interface=wlan1 | 
|  | add address=192.168.2.1/24 network=192.168.2.0 broadcast=192.168.2.255 interface=wlan2 | 
|  | add address=192.168.3.1/24 network=192.168.3.0 broadcast=192.168.3.255 interface=wlan3 | 
|  | add address=192.168.4.1/24 network=192.168.4.0 broadcast=192.168.4.255 interface=wlan4 | 
|  | / ip firewall mangle | 
|  | add chain=input in-interface=wlan1 action=mark-connection new-connection-mark=wlan1_conn | 
|  | add chain=input in-interface=wlan2 action=mark-connection new-connection-mark=wlan2_conn | 
|  | add chain=input in-interface=wlan3 action=mark-connection new-connection-mark=wlan3_conn | 
|  | add chain=input in-interface=wlan4 action=mark-connection new-connection-mark=wlan4_conn | 
|  | add chain=output connection-mark=wlan1_conn action=mark-routing new-routing-mark=to_wlan1 | 
|  | add chain=output connection-mark=wlan2_conn action=mark-routing new-routing-mark=to_wlan2 | 
|  | add chain=output connection-mark=wlan3_conn action=mark-routing new-routing-mark=to_wlan3 | 
|  | add chain=output connection-mark=wlan4_conn action=mark-routing new-routing-mark=to_wlan4 | 
|  | add chain=prerouting dst-address=192.168.1.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address=192.168.2.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address=192.168.3.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address=192.168.4.0/24 action=accept in-interface=lan | 
|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:4/0 \ action=mark-connection new-connection-mark=wlan1_conn passthrough=yes | 
|  | add chain=prerouting dst-address-type=!local in-interface=lan per-connection-classifier=both-addresses:4/1 \ action=mark-connection new-connection-mark=wlan2_conn passthrough=yes | 
