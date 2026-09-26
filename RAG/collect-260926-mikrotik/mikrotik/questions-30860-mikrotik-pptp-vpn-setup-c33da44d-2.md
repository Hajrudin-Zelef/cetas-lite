---
id: collect-260926-mikrotik/mikrotik/questions-30860-mikrotik-pptp-vpn-setup-c33da44d-2
title: "may/31/2016 23:02:50 by RouterOS 6.35.1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-30860-mikrotik-pptp-vpn-setup-c33da44d.md
source_anchor: ""
source_lines: [84, 129]
sha256: 3dd87103f4603fdaa69b483cc11ef1e68de393a533bfd4e8416f98286d66979f
---

# may/31/2016 23:02:50 by RouterOS 6.35.1

20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 54
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 56
20:20:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 62
20:20:45 pptp,ppp,info <pptp-USERNAME>: connected
20:20:45 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 50
20:20:45 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 334
20:20:48 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 334
20:20:51 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 334
20:20:54 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 334
20:20:57 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 334
20:21:04 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:21:04 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 40
20:21:24 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:21:25 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 40
20:21:44 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:21:45 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 40
20:22:04 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:22:06 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 40
20:22:24 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 48
20:22:27 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 40
20:22:32 pptp,ppp,info <pptp-USERNAME>: terminating...
20:22:32 pptp,ppp,info,account USERNAME logged out, 109 1568 98 14 8
20:22:32 pptp,ppp,info <pptp-USERNAME>: disconnected
20:22:32 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 57
20:22:32 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto 47, 82.132.216.62->82.XXX.XXX.177, len 57
20:22:32 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,FIN), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
20:22:32 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,FIN), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
20:22:32 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 82.132.216.62:14321->82.XXX.XXX.177:1723, len 52
22:00:44 firewall,info remote-access input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 93.174.93.94:47264->82.XXX.XXX.177:80, len 40
22:01:09 firewall,info remote-access input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 61.240.144.64:48406->82.XXX.XXX.177:80, len 40
22:09:42 firewall,info remote-access input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 94.102.49.54:22->82.XXX.XXX.177:80, len 40
22:57:26 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 123.151.149.222:22200->82.XXX.XXX.177:1723, len 40
22:57:26 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (RST), 123.151.149.222:22200->82.XXX.XXX.177:1723, len 40
22:57:26 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (RST), 123.151.149.222:22200->82.XXX.XXX.177:1723, len 40
22:57:26 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 123.151.42.61:17122->82.XXX.XXX.177:1723, len 48
22:57:27 pptp,info TCP connection established from 123.151.42.61
22:57:27 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 123.151.42.61:17122->82.XXX.XXX.177:1723, len 40
22:57:27 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,PSH), 123.151.42.61:17122->82.XXX.XXX.177:1723, len 196
22:57:27 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 123.151.42.61:17122->82.XXX.XXX.177:1723, len 40
22:57:27 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK,FIN), 123.151.42.61:17122->82.XXX.XXX.177:1723, len 40
22:57:27 firewall,info input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (ACK), 123.151.42.61:17122->82.XXX.XXX.177:1723, len 40
22:59:34 firewall,info remote-access input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 141.212.122.151:47113->82.XXX.XXX.177:80, len 40
22:59:34 firewall,info remote-access input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 141.212.122.152:38568->82.XXX.XXX.177:80, len 40
22:59:35 firewall,info remote-access input: in:ether1-gateway out:(none), src-mac 00:01:5c:82:ee:47, proto TCP (SYN), 141.212.122.145:45406->82.XXX.XXX.177:80, len 60
I have also enabled proxy-arp on the LAN connections.
I don't know why I cant get traffic to route from VPN > LAN.
