---
id: collect-260926-mikrotik/mikrotik/questions-30860-mikrotik-pptp-vpn-setup-c33da44d-1
title: "may/31/2016 23:02:50 by RouterOS 6.35.1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-30860-mikrotik-pptp-vpn-setup-c33da44d.md
source_anchor: ""
source_lines: [1, 83]
sha256: f8eb28a5589129ed54afe49fb96b82750d5d180f6ae53a962b86a3373615635b
---

# may/31/2016 23:02:50 by RouterOS 6.35.1

I have attempted and followed many guides and am still having issues allowing traffic from my VPN to access devices on the LAN.
https://rbgeek.wordpress.com/2014/08/26/pptp-server-setup-on-mikrotik/ Is an example of one guide I have followed, and although I can connect, and ping from LAN > VPN I am unable to ping from VPN > LAN
I have the following setup:
Routerboard 750 WebFig v6.35.1 (stable)
LAN 192.168.88.0/24
PPTP Pool 192.168.200.10-192.168.200.20
PPTP Server: Enabled
PPTP Profile created using the PPTP IP Pool for both internal and external addresses
/ip pool
add name=dhcp ranges=192.168.88.100-192.168.88.254
add name=pptp-pool ranges=192.168.200.1-192.168.200.10<
# may/31/2016 23:02:50 by RouterOS 6.35.1
# software id = 8RIQ-2NZU
#
/ip firewall filter
add chain=input comment="default configuration" protocol=icmp
add chain=input comment="VPN PPTP ACCEPT" dst-port=1723 log=yes protocol=tcp
add chain=input comment="GRE ACCEPT" log=yes protocol=gre
add chain=input comment="default configuration" connection-state=established
add chain=input comment="default configuration" connection-state=related
add chain=input comment="allow l2tp" dst-port=1701 protocol=udp
add chain=input comment="allow sstp" dst-port=443 protocol=tcp
add chain=input comment="web access for config" dst-port=80 in-interface=ether1-gateway log=yes log-prefix=remote-access protocol=tcp
add action=drop chain=input comment="default configuration" in-interface=ether1-gateway
/ip firewall nat
add action=masquerade chain=srcnat comment="default configuration" out-interface=ether1-gateway to-addresses=0.0.0.0
# may/31/2016 23:03:52 by RouterOS 6.35.1
# software id = 8RIQ-2NZU
#
/ppp profile
add local-address=pptp-pool name=pptp-profile remote-address=pptp-pool
set *FFFFFFFE dns-server=0.0.0.0 use-compression=yes
/ppp secret
add name=USERNAME password=PASSWORD profile=pptp-profile service=pptp
From the logs
20:20:35 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 52
20:20:35 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,PSH), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 220
20:20:35 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 52
20:20:35 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,PSH), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 76
20:20:35 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 60
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 52
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 59
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 98
20:20:36 pptp,ppp,info,account USERNAME logged in, 192.168.200.10
20:20:36 pptp,ppp,info <pptp-USERNAME>: authenticated
20:20:36 pptp,ppp,info <pptp-USERNAME>: terminating...
20:20:36 pptp,ppp,info,account USERNAME logged out, 1 18 28 3 4
20:20:36 pptp,ppp,info <pptp-USERNAME>: disconnected
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 44
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 44
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 74
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,FIN), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 52
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,FIN), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 52
20:20:36 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14320->82.XXX.XXX.177:1723, len 52
20:20:41 system,info PPTP Server settings changed by admin
20:20:41 system,info PPTP Server settings changed by admin
20:20:43 pptp,info TCP connection established from 82.132.216.62
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 64
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,PSH), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 208
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,PSH), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 220
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,PSH), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 76
20:20:43 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 60
20:20:44 pptp,ppp,info,account USERNAME logged in, 192.168.200.10
20:20:44 pptp,ppp,info <pptp-USERNAME>: authenticated
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 59
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 98
20:20:44 pptp,ppp,info <pptp-USERNAME>: using encoding - MPPE128 stateless
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 44
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 62
