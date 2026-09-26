---
id: collect-260926-mikrotik/mikrotik/2-wan-connections-mangle-rules-and-wireguard
title: "2-wan-connections-mangle-rules-and-wireguard"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/2-wan-connections-mangle-rules-and-wireguard.md
source_anchor: ""
source_lines: [1, 137]
sha256: 95e0c96a23f93932a278e70839c6b7cee53b60b759d6baa3867c6118eeee1d94
---

# 2-wan-connections-mangle-rules-and-wireguard

Okay here ya go!!

(1)  wireguard settings

*/interface wireguard
add comment=“WG Interface For A1/WAN1” listen-port=1**3**231 mtu=1420 name=
wireguard1
add comment=“WG Interface for EASYTV/WAN2” disabled=yes listen-port=1**4**23**2** 
mtu=1420 name=wireguard2*

(2) /interface wireguard peers

*add allowed-address=192.168.99.**3**/32 comment=“PC WG WAN1” interface=wireguard1 
public-key=“H4APrAYA7deOVfm2fETQybTL0aOEY23eo9s3kHSBCiE=”
add allowed-address=192.168.98.**2**/32 comment=“PC WG WAN2” interface=wireguard2 
public-key=“iowZhns7OHIH8Di+lZ4zpJMzM8Ts+GExi9dOb8CCPGk=”
add allowed-address=192.168.99.**5**/32 comment=“Phone WG WAN1” interface=
wireguard1 public-key=“RF/lwraS3xUzJUYA7De8PUFlMIDHxPH8OQat5LHTjVI=”
add allowed-address=192.168.98.**4**/24 comment=“Phone WG WAN2” interface=wireguard2 
public-key=“NyjZt96W9I79h5cTct6rEpWk5nYhhgDp7DaQIu9Kw0A=”*

**NOTE:** WIREGUARD**1** is odd  .9**9** and WIREGUARD**2** is even .9**8**, so I made clients the same  .3 and .5 for 99 and  .2 and .4 for 98,  I like patterns… Same with

listening ports  wan1 -  1**3**23**1,**   wan2  - 1**4**23**2**  

(3) /ip address

*add address=192.168.1.20/24 interface=WAN1 network=192.168.1.0
add address=192.168.100.1/24 interface=BridgeLAN network=192.168.100.0
add address=192.168.101.20/24 interface=WAN2 network=192.168.101.0
add address=192.168.**99**.**1/24** interface=wireguard1 network=192.168.99.0
add address=192.168.**98.****1/24** interface=wireguard2 network=
192.168.98.0*

(4) /ip dns

**allow-remote-servers**=yes****  set servers=1.1.1.1,8.8.8.8

(5) /ip firewall filter

…

*add action=accept chain=input comment=“allow WireGuard” dst-port=1**3**231 
protocol=udp
add action=accept chain=input comment=“allow WireGuard” dst-port=1**4**232 
protocol=udp*

…

(6)  ADD to both  entries…

/interface list

*add name=WAN
add name=LAN
**add name=WG***

/interface list member

…

add interface=wireguard1 list=WG

add interface=wireguard2 list=WG

…

(7) Modify from:

/ip firewall filter

*add action=accept chain=forward dst-address=192.168.100.0/24 in-interface=wireguard1*

TO:

*add action=accept chain=forward dst-address=192.168.100.0/24 in-interface**-list**=**WG***

(8) The fun part is the IP NAT and mangling.

We want to ensure a couple of things here.

A.   That incoming traffic to wireguard is not caught in the mangling.

B.   That reply traffic from LAN subnet to remote WG user is not caught in PCC mangling.

C.  Fixed your IP routes setup.

/ip firewall mangle

*add action=mark-connection chain=prerouting connection-mark=no-mark 
in-interface=WAN1 new-connection-mark=WAN1_conn passthrough=yes  **dst-port=!13231**  { excludes handshake from marking } ***
add action=mark-connection chain=prerouting connection-mark=no-mark 
in-interface=WAN2 new-connection-mark=WAN2_conn passthrough=yes **dst-port=!14232*** {excludes handshake from marking } ***

**add action=accept chain=prerouting connection-mark=no-mark in-interface=BridgeLAN  
dst-address-list=Wg-subnets**  { traffic heading back to WG remote users is not marked }

*add action=mark-connection chain=prerouting connection-mark=no-mark 
dst-address-type=!local in-interface=BridgeLAN new-connection-mark=
WAN1_conn passthrough=yes per-connection-classifier=both-addresses:2/0
add action=mark-connection chain=prerouting connection-mark=no-mark 
dst-address-type=!local in-interface=BridgeLAN new-connection-mark=
WAN2_conn passthrough=yes per-connection-classifier=both-addresses:2/1
add action=mark-routing chain=prerouting connection-mark=WAN1_conn 
in-interface=BridgeLAN new-routing-mark=to_WAN1 passthrough=yes
add action=mark-routing chain=prerouting connection-mark=WAN2_conn 
in-interface=BridgeLAN new-routing-mark=to_WAN2 passthrough=yes
add action=mark-routing chain=output connection-mark=WAN1_conn 
new-routing-mark=to_WAN1 passthrough=yes
add action=mark-routing chain=output connection-mark=WAN2_conn 
new-routing-mark=to_WAN2 passthrough=yes*

/ip firewall address

add address=192.168.99.0/24 list=WG-subnets

add address=192.168.98.0/24 list=WG-subnets

/routing table

add name=To-WAN1 fib

add name=To-WAN2 fib

/ip route

add distance=5 dst-address=0.0.0.0/0 gwy=ISP1  routing-table=main check-gateway=ping

add distance=10 dst-address=0.0.0.0/0 gwy=ISP2  routing-table=main

add dst-address=0.0.0.0/0 gwy=ISP1  routing-table=To-WAN1

add dst-address=0.0.0.0/0 gwy=ISP2  routing-table=To-WAN2

/routing rule add action=lookup-only-in-table src-address=IP-of-WAN-2  table=To-WAN2  { ensures if both WANs are up, that a handshake on WG2 will go back out WAN2 }

***  The reason for these two rules is to ensure the port forwarding coming on WAN1 and WAN2 go out WAN1 and WAN2 respectively during PCC.  IF you had no port forwarding they could be removed. I simply added the fact to not mark the wg incoming handshake connections so that they could follow routing separately and thus allow you to have both up at the same time if both WANs were available ( WAN1 on table main because its PRIMARY and WAN2 on routing rule because its secondary )
