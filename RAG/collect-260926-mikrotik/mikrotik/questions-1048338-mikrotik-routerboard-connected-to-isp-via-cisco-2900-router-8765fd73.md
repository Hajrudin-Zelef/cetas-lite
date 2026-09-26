---
id: collect-260926-mikrotik/mikrotik/questions-1048338-mikrotik-routerboard-connected-to-isp-via-cisco-2900-router-8765fd73
title: "questions-1048338-mikrotik-routerboard-connected-to-isp-via-cisco-2900-router-8765fd73"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1048338-mikrotik-routerboard-connected-to-isp-via-cisco-2900-router-8765fd73.md
source_anchor: ""
source_lines: [1, 46]
sha256: 9d190c8b0b0564945df0a436c2617d06980d56cd9d9b2f5e63c5d27ae4b41cf4
---

# questions-1048338-mikrotik-routerboard-connected-to-isp-via-cisco-2900-router-8765fd73

I'm sorry it's a bit of a long one and I had to mask some information.
With that said, I am in a conundrum and need help. Our network guy left and I've been tasked with getting the house in order. I do have a bit of background in networking but that's well over a decade ago.
Key Details:
About 2,000 workforce mostly connected via omnidirectional & sectoral WLAN
A Cisco 2900 router and a Mikrotik RouterBoard were purchased
DIA via microwave
Connection should be ISP IDU to Cisco Router to Mikrotik RouterBoard to end-users
ISP given addresses:
IP 1xx.xx.x.162
NM 255.255.255.252
GW 1xx.xx.x.161
ISP given DNS:
xxx.xxx.xxx.1
xxx.xxx.xxx.2
Summarily, I configured the Cisco:
#GE0/0 - 1xx.xx.x.162 255.255.255.252
#GE0/1 - 192.168.xx.1 255.255.255.0
#DHCP pool 192.168.xx.0 255.255.255.0
#Default Route 1xx.xx.x.162
#dns-server xxx.xxx.xxx.1 xxx.xxx.xxx.2
#excluded 192.168.xx.241 192.168.xx.254
#GE0/0 nat outside
#GE0/1 nat inside (also source list & overloaded)
#ip route 0.0.0.0 0.0.0.0 1xx.xx.x.161
It's been a while since I did this but I believe it's ok (corrections will be appreciated if otherwise)
My challenge is Mikrotik. I've gotten a refresher and I believe setting up hotspot, profiles and users won't be an issue rather IP addressing, DHCP and NAT as there may be a double/conflict with Cisco.
So, I configured the Mikrotik:
**eth1:** 192.168.xx.1/24 (IP of Cisco GE0/1)
**eth2:** 192.168.1.0/24
**DHCP client:** eth1
Routes:
Destination 0.0.0.0/0
 
Gateway 1xx.xx.x.161 (or should I use IP of GE0/1 i.e., 192.168.xx.1?)
**NAT:**
Source 192.168.1.0/24
Destination 0.0.0.0/0, Action Masquerade
DHCP Server Using DHCP setup:
Int 2
Address 192.168.1.0/24
Gateway 192.168.1.1(IP of Mikrotik)
Address to give out (192.168.1.2, 192.168.1.254)
DNS servers (192,168.1.1, xxx.xxx.xxx.1, xxx,xxx,xxx,2)
NAT:
port-forwarding using netmap on port 80.
That's it with Mikrotik and like I said, hotspot, profiles and users won't be an issue But I really need help on the others.
