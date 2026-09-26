---
id: collect-260926-mikrotik/mikrotik/dual-wan-failover-port-forward-not-working-when-changing-route-distance
title: "dual-wan-failover-port-forward-not-working-when-changing-route-distance"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/troubleshooting/dual-wan-failover-port-forward-not-working-when-changing-route-distance.md
source_anchor: ""
source_lines: [1, 59]
sha256: 08dfd88bb3a7da9813750c815988019c7ea114abb4da6ec880b66076283a30d5
---

# dual-wan-failover-port-forward-not-working-when-changing-route-distance

(1) **Insecure and potentially dangerous**  to expose winbox port to the internet.  Suggest access router via Wireguard.

*add action=accept chain=input comment=“Remote access MEXUS” dst-port=8291 
protocol=tcp src-address=X.X>X>X*

(2)  You should only have four mangle rules.

FIXED

*/ip firewall mangle
add action=mark-connection chain=prerouting connection-mark=no-mark 
in-interface=ether1 new-connection-mark=ether1_mark passthrough=yes
add action=mark-routing chain=output comment=“Ether 1 output routing” 
connection-mark=ether1_mark disabled=no new-routing-mark=to_ether1 
passthrough=no
add action=mark-connection chain=prerouting  connection-mark=no-mark 
in-interface=ether2 new-connection-mark=ether2_mark passthrough=yes
add action=mark-routing chain=output comment=“Ether 2 output routing” 
connection-mark=ether2_mark disabled=no new-routing-mark=to_ether2 
passthrough=no*

(3) DST NAT has two possibilities.

a. Will work  - you do not have local users accessing servers by DNS or WANIP, and only by LANIP (directly).

b. Will NOT work - users are forced to access servers via DYNDNS type domain name/url or by WANIP.

(If b is the correct answer then we have to make changes due to hairpin nat).

(4) Do not use same DNS sites for Recursive as you do for DNS… so recommend change DNS

Keep 8.8.8.8 and 1.1.1.1  for DNS and  8.8.4.4 for recursive and 1.0.0.1 for recursive.

*/ip dns
set allow-remote-requests=yes servers=
**8.8.8.8,1.1.1.1,9.9.9.9,208.67.222.222***

AND REMOVE THIS DEFAULT static setting at IP DNS.

*/ip dns static
add address=192.168.0.1 comment=defconf name=router.lan*

(5) Routes are wrong…

FIXED

/ip route

add check-gateway=ping distance=2 dst-address=0.0.0.0/0 gateway=8.8.4.4  routing-table=main scope=10 target-scope=12

add check-gateway=ping distance=4 dst-address=0.0.0.0/0 gateway=1.0.0.1  routing-table=main scope=10 target-scope=12

add distance=2 dst-address=8.8.4.4/32 gateway=10.30.152.1 routing-table=main scope=10 target-scope=11

add distance=4 dst-address=1.0.0.1/32 gateway=10.10.10.112 routing-table=main scope=10 target-scope=11

add distance=1 dst-address=0.0.0.0/0 gateway=10.30.152.1 routing-table=to_ether1 scope=10 target-scope=30

add distance=1 dst-address=0.0.0.0/0 gateway=10.10.10.112 routing-table=to_ether2 scope=10 target-scope=30
