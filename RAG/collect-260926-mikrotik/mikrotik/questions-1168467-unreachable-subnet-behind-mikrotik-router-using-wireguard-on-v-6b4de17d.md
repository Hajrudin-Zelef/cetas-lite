---
id: collect-260926-mikrotik/mikrotik/questions-1168467-unreachable-subnet-behind-mikrotik-router-using-wireguard-on-v-6b4de17d
title: "questions-1168467-unreachable-subnet-behind-mikrotik-router-using-wireguard-on-v-6b4de17d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/questions-1168467-unreachable-subnet-behind-mikrotik-router-using-wireguard-on-v-6b4de17d.md
source_anchor: ""
source_lines: [1, 44]
sha256: fdbf2b480d7a0a13854e26b681b1ee7a344ec9649bc960655b4d79772d2db630
---

# questions-1168467-unreachable-subnet-behind-mikrotik-router-using-wireguard-on-v-6b4de17d

I need set up a Wireguard connection from my home Microtik router to a VPS, and connect from outside through this connection to my home networks. Here, only one network is depicted.
The Microtik rules and configuration:
/interface/wireguard/export
/interface wireguard
add listen-port=36857 mtu=1420 name=wg-xxx-it
add disabled=yes listen-port=13231 mtu=1420 name=wireguard
/interface wireguard peers
add allowed-address=0.0.0.0/0 client-address=192.168.7.1/24 endpoint-address=wg.xxx-it.de endpoint-port=51820 interface=wg-xxx-it name=peer-wg-xxx-it persistent-keepalive=30m preshared-key="xxxx" public-key=\
    "xxxx"
/ip address
add address=192.168.0.1/24 comment="xxx-net" interface="bridge" network=192.168.0.0
add address=192.168.7.1/24 comment=VPN interface=wg-xxx-it network=192.168.7.0
The mobile client configuration:
[Interface]
PrivateKey = xxx
Address = 192.168.7.10/24
DNS = 192.168.7.1
[Peer]
PublicKey = xxx
PresharedKey = xxx
AllowedIPs = 192.168.7.0/24, 192.168.0.0/24
PersistentKeepalive = 0
Endpoint = wg.xxx-it.de:51820
After the client connects, it can ping the router and the router can ping the client. Router gets the IP 192.168.7.1/24, client 192.168.7.10/24.
However, the client is unable to connect to other parts of the home network. As for the Wireguard server, I'm using wg-easy with the following env-variables in docker-compose configuration:
environment:
  - WG_ALLOWED_IPS=192.168.7.0/24, 192.168.0.0/24
  - WG_DEFAULT_ADDRESS=192.168.7.x
  - WG_DEFAULT_DNS=192.168.7.1
  - WG_POST_UP=ip route add 192.168.0.0/24 via 192.168.7.1
  - WG_POST_DOWN=ip route rm 192.168.0.0/24 via 192.168.7.1
A traceroute from the client to a device on the home network results in endless hops (unreachable).
It looks to me like the route on the VPS is not working correctly, however, the route is there:
xxx:/app# route
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         docker-xx-xxxxx 0.0.0.0         UG    0      0        0 eth0
10.8.0.0        *               255.255.255.0   U     0      0        0 wg0
172.28.0.0      *               255.255.0.0     U     0      0        0 eth0
192.168.0.0     192.168.7.1     255.255.255.0   UG    0      0        0 wg0
192.168.7.1     *               255.255.255.255 UH    0      0        0 wg0
192.168.7.10    *               255.255.255.255 UH    0      0        0 wg0
Is there some configuration missing on router or on VPS?
/tool/sniffer/quickfor a specific IP address show packets arriving through the WG interface, and the same packets being forwarded out through the LAN bridge interface?tcpdump -n -i wg0on the VPS in that case? (See man.archlinux.org/man/pcap-filter.7 for filter syntax). I assume you meant no packets when connecting to addresses in the 192.168.0.x subnet?viaroutes, for the same reason it cannot be bridged.]
