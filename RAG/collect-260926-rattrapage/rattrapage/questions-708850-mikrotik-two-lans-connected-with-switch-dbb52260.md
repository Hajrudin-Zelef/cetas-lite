---
id: collect-260926-rattrapage/rattrapage/questions-708850-mikrotik-two-lans-connected-with-switch-dbb52260
title: "questions-708850-mikrotik-two-lans-connected-with-switch-dbb52260"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-rattrapage/ai-llm/questions-708850-mikrotik-two-lans-connected-with-switch-dbb52260.md
source_anchor: ""
source_lines: [1, 50]
sha256: 714762c73bd49c0c4eed0b3c231acba7487f29fe3327f41fdbfe859868d19d41
---

# questions-708850-mikrotik-two-lans-connected-with-switch-dbb52260

I have a very simple network. I have two LAN networks (192.168.2.0/24 and 192.168.3.0/24) connected via a router at each site and the routers are connected via a switch.
http://postimg.org/image/3y1uysszn/
The routers are MIKROTIK, the switch is some of the shelf equipment. I've set up all the routes, removed all the firewalls but I still can't ping from one PC to another. The strange thing is that when I use MIKROTIKs IP Scan tool, it finds all of the equipment, but when I try to ping lets say from PC at site 2, I can't get further than 172.30.2.222.
If I disable the bridge between LAN and WAN at site 2, I can (from the PC at site 2) ping to LAN IP: 192.168.3.50, which is at site 3. At the same time I can't ping to LAN IP: 192.168.2.1 from PC at site 3. If I reenable the bridge at site 2, I again can't get any further than 172.30.2.222 from site 2.
Does anyone have an idea what I am doing wrong? Is the PING somehow disabled in mikrotik routers?
Configuration:
[admin@ENG. SITE 3] >> /ip address export 
/ip address 
add address=192.168.3.1/24 comment="default configuration" interface=\ 
"ETH. 2 LAN" network=192.168.3.0 
add address=172.30.2.222/24 interface="ETH. 1 WAN" network=172.30.2.0 
[admin@ENG. SITE 3] >> ip route export 
/ip route 
add distance=1 gateway=172.30.2.221 add distance=1 dst-address=172.30.2.0/32 gateway="ETH. 1 WAN" 
add distance=1 dst-address=192.168.2.0/24 gateway="ETH. 1 WAN"
[admin@ENG. SITE 3] >> ip firewall export 
/ip firewall filter 
add chain=input comment="default configuration" disabled=yes protocol=icmp
add chain=input comment="default configuration" connection-state=established \ 
disabled=yes 
add chain=input comment="default configuration" connection-state=related \ 
disabled=yes 
add action=drop chain=input comment="default configuration" disabled=yes \ 
in-interface="ETH. 1 WAN"
add chain=forward comment="default configuration" connection-state=established \ 
disabled=yes 
add chain=forward comment="default configuration" connection-state=related \ 
disabled=yes 
add action=drop chain=forward comment="default configuration" connection-state=\ 
invalid disabled=yes 
/ip firewall nat 
add action=masquerade chain=srcnat comment="default configuration" \ 
out-interface="ETH. 1 WAN" 
[admin@ENG. SITE 2] > ip address export 
/ip address 
add address=192.168.2.1/24 comment="default configuration" interface "ETH. 2 LAN" network=192.168.2.0 
add address=172.30.2.221/24 interface="ETH. 1 WAN" network=172.30.221.0
[admin@ENG. SITE 2] > ip route export 
/ip route 
add disabled=yes distance=1 gateway=172.30.2.222 
add distance=1 dst-address=192.168.3.0/24 gateway="ETH. 1 WAN" 
[admin@ENG. SITE 2] > ip firewall export 
/ip firewall filter 
add chain=forward comment="default configuration" connection-state=e disabled=yes 
add chain=forward comment="default configuration" connection-state=r disabled=yes 
add action=drop chain=forward comment="default configuration" connec invalid disabled=yes 
/ip firewall nat 
add action=masquerade chain=srcnat comment="default configuration" \ 
out-interface="ETH. 1 WAN"
/ip address export/ip route export/ip firewall exportIf I disable the bridge between LAN and WAN at site 2? Do you have the routers in bridged mode? If so, why?
