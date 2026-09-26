---
id: collect-260926-mikrotik/mikrotik/questions-599086-openvpn-tun-routing-can-ping-tun-interfaces-with-mikrotik-and-n-8deaa86f
title: "questions-599086-openvpn-tun-routing-can-ping-tun-interfaces-with-mikrotik-and-n-8deaa86f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-599086-openvpn-tun-routing-can-ping-tun-interfaces-with-mikrotik-and-n-8deaa86f.md
source_anchor: ""
source_lines: [1, 71]
sha256: b91730e57387fe8e55f53c5edfc0686693e1f6def195cc13f2d7b4f512f0d16b
---

# questions-599086-openvpn-tun-routing-can-ping-tun-interfaces-with-mikrotik-and-n-8deaa86f

Problem
The Mikrotik site is unable to ping anything on the server site. The Server site is less important right now.
OpenVPN Server config is on LINUX
port 1194
proto tcp
dev tun
ca ca.crt
cert cert.crt
key cert.key
dh dh1024.pem
server 10.8.0.0 255.255.255.0
ifconfig-pool-persist ipp.txt
puch route 192.168.0.0 255.255.255.0
client-to-client
keepalive 10 120
persist-key
persist-tun
ver 3
mute 20
ipv4_fowarding is enabled
Open VPN Client is on Mikrotik Router.
connect to: server.public.ip
port: 1194
mode: ip
user: user
password: passwrd
profile: openvpn-out
cert: cert
auth: sha1
cipher: blowfish 128
**openvpn-out profile**
name: openvpn-out
change tcp mss: default
use mpls: default
use compression: default
use vj compression: default
use encryption: default
With those settings mikrotik give: Connection is established and everything that should work work on the mikrotik itself
1) From mikrotik router I can't ping the TUN interface (10.8.0.2) server but can ping own TUN (10.8.0.46) and subnet on server site (192.168.0.0)
2) The clients of subnet from Mikrotik site are unable to ping anything on server site, only TUN interface on mikrotik (10.8.0.46)
3) From Server I can ping TUN interface of Mikrotik (10.8.0.46) but can't subnet on that site (192.168.2.0)
4) The clients of subnet from Server sire are unable to ping anything on Mikrotik site, only TUN interface on both Mikrotik and Server
Mikrotik interface
ether1 - dynamic.ip.addres
OPVN - 10.8.0.46 network 10.8.0.45
bridge1 - 192.168.2.1/24
Mikrotik route:
   ds.address  gw 
    0.0.0.0/0   public.ip.address reachable ether1
    10.8.0.0/24 10.8.0.45 reachable OVPN
    10.8.0.45   OVPN reachable 
    public.ip   ether1 reachable 
    192.168.0.0/24 OVPN reachable 
    192.168.2.0/24 bridge1 reachable
Mikrotik ping 192.168.0.0 without problems Clients of mikrotik can't
Server interfaces
eth0 - public.ip.addre
eth1 - 192.168.0.9
tun0 - 10.8.0.1 p-t-p:10.8.0.2
Server route:
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         public.ip.addre 0.0.0.0         UG    100    0        0 eth0
10.8.0.0        10.8.0.2        255.255.255.0   UG    0      0        0 tun0
10.8.0.2        *               255.255.255.255 UH    0      0        0 tun0
localnet        *               255.255.255.252 U     0      0        0 eth0
192.168.0.0     *               255.255.255.0   U     0      0        0 eth1
192.168.2.0     10.8.0.1        255.255.255.0   UG    0      0        0 tun0
Server can't ping 192.168.2.0, ping 10.8.0.46. Server Clients ping only 10.8.0.46
I assume that iptables are set correctly in some way because connection is established and mikrotik can ping anything on there other site.
mikrotik nat rule
srcnat -o ether1 -acction masquerade
