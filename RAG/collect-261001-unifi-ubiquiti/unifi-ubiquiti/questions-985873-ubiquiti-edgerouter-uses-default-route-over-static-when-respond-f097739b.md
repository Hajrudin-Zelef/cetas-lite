---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-985873-ubiquiti-edgerouter-uses-default-route-over-static-when-respond-f097739b
title: "questions-985873-ubiquiti-edgerouter-uses-default-route-over-static-when-respond-f097739b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-985873-ubiquiti-edgerouter-uses-default-route-over-static-when-respond-f097739b.md
source_anchor: ""
source_lines: [1, 50]
sha256: 2ee92ae96072a52048b50fc8d217e2dfffc2456cb033adcf4f0c0ca89d554849
---

# questions-985873-ubiquiti-edgerouter-uses-default-route-over-static-when-respond-f097739b

I need to have a working connection between two router devices and the subnets behind them. Everything else in the setup seems to work except one static route. Here is the basic information about my setup:
Juniper:
Vlan95 10.250.0.1/24
Vlan1  192.168.1.1/21
Vlan1  192.168.4.1/21 (this interface has two default gateways due to ongoing development in the network)
fe-0/0/3 10.129.0.1/22
Ubiquiti:
eth2.620 10.250.0.3/24
eth2.100 192.168.8.1/24
routes - ubiquiti:
S    *> 0.0.0.0/0 [1/0] via xxx.xxx.xxx.xxx, eth0
C    *> 10.0.0.2/32 is directly connected, lo
C    *> 10.0.1.0/24 is directly connected, eth1
C    *> 10.0.2.0/24 is directly connected, eth2.200
S    *> 10.129.0.0/24 [1/0] via 10.250.0.1, eth2.620 (This works)
C    *> 10.250.0.0/24 is directly connected, eth2.620
C    *> 127.0.0.0/8 is directly connected, lo
C    *> 130.234.212.64/28 is directly connected, eth0
C       130.234.212.64/28 is directly connected, eth0
C    *> 192.168.0.0/24 is directly connected, eth1
S    *> 192.168.1.0/24 [1/0] via 10.250.0.1, eth2.620 (This doesn't seem to work)
S    *> 192.168.4.0/24 [1/0] via 10.250.0.1, eth2.620 (This does work)
S    *> 192.168.5.0/24 [1/0] via 10.250.0.1, eth2.620
C    *> 192.168.8.0/24 is directly connected, eth2.100
C    *> 192.168.24.0/22 is directly connected, eth2.14
In this setup networks 192.168.4.0/24 and 10.129.0.0/24 work. I can ping my Ubiquiti routers interface 10.250.0.3 from both networks and I can ping IP-addresses in both networks from that Ubiquiti IP-address.
The 192.168.1.0/24 route is different though. It looks like it doesn't exist at all, because when pinging from, for example 192.168.1.99, the traffic arrives correctly but instead of being sent back from 10.250.0.3 over the link on interface eth2.620 the Ubiquiti router tries to send them over eth0, which is my WAN address.
I verified this using tcpdump on ubiquiti
Here is a dump of successful ICMP request from 10.129.0.18 to 10.250.0.3:
user@device:~$ sudo tcpdump -i eth2.620 host 10.129.0.18
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on eth2.620, link-type EN10MB (Ethernet), capture size 262144 bytes
13:37:40.836129 IP 10.129.0.18 > 10.250.0.3: ICMP echo request, id 1, seq 103, length 40
13:37:40.836302 IP 10.250.0.3 > 10.129.0.18: ICMP echo reply, id 1, seq 103, length 40
13:37:41.841624 IP 10.129.0.18 > 10.250.0.3: ICMP echo request, id 1, seq 104, length 40
13:37:41.841819 IP 10.250.0.3 > 10.129.0.18: ICMP echo reply, id 1, seq 104, length 40
Here is a dump where I'm trying to ping from 192.168.1.99 to 10.250.0.3 (taken from eth2.620):
user@device:~$ sudo tcpdump -i eth2.620 host 192.168.1.99
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on eth2.620, link-type EN10MB (Ethernet), capture size 262144 bytes
13:38:45.563661 IP 192.168.1.99 > 10.250.0.3: ICMP echo request, id 1, seq 1364, length 40
13:38:50.564150 IP 192.168.1.99 > 10.250.0.3: ICMP echo request, id 1, seq 1365, length 40
And here is a dump taken from eth0 (WAN interface) when I'm trying to ping from 192.168.1.99 to 10.250.0.3:
user@device:~$ sudo tcpdump -i eth0 host 192.168.1.99
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
13:34:42.558520 IP 10.250.0.3 > 192.168.1.99: ICMP echo reply, id 1, seq 1353, length 40
13:34:47.558410 IP 10.250.0.3 > 192.168.1.99: ICMP echo reply, id 1, seq 1354, length 40
13:34:52.559685 IP 10.250.0.3 > 192.168.1.99: ICMP echo reply, id 1, seq 1355, length 40
Is there any easily recognizable problem in this setup? As far as I can see the static routes are all configured the same.
