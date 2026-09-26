---
id: collect-260926-mikrotik/mikrotik/questions-1619525-setup-mikrotik-hap-ac-to-work-with-fiber-internet-c07bd339
title: "questions-1619525-setup-mikrotik-hap-ac-to-work-with-fiber-internet-c07bd339"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1619525-setup-mikrotik-hap-ac-to-work-with-fiber-internet-c07bd339.md
source_anchor: ""
source_lines: [1, 31]
sha256: f45f0616b139629eb4cf5ac8db11c0a81b698b17897fac0c8a9a8d3882b2c0cf
---

# questions-1619525-setup-mikrotik-hap-ac-to-work-with-fiber-internet-c07bd339

I got new ISP few days ago, and it's fiber line that according to ISP I can use with my MikroTik router.
In mail I got the following (made up IP, but same pattern):
WAN block 93.83.73.68/30
ISP port: 93.83.73.69
Your port: 93.83.73.70
Subnet mask 255.255.255.252
LAN block: 77.67.57.215/32
Subnet: 255.255.255.255
WAN block is used for interconnection between our network and the device at your location and can’t be used for Internet access, LAN block is given to you, registered in your name and is used for Internet access.
Up until now I used WAN IP + private block 192.168.88.0/24, but now my LAN block is a single public IP?
If I understood correctly it should be setup like this:
- Set ISP facing port IP 93.83.73.70/30
- Set Gateway IP 93.83.73.69
- Set DNS IP 8.8.8.8, 8.8.4.4
- Set LAN network and DHCP 192.168.88.0/24
- Add Firewall NAT rule that like this:
add chain=srcnat src-address=192.168.88.0/24 action=src-nat to-addresses=77.67.57.215 out-interface=ether1
Am I missing something, it won't work still and ISP phone support wants to sell me their crappy router?
IP > Routes shows 93.83.73.69 as reachable via ether1.
//Updates after following instructions from the comments:
/ping 93.83.73.69 src-address=93.83.73.70
gets timeout and unreachable
/ping 93.83.73.69 src-address=93.83.73.70 interface=ether1 arp-ping=yes
gets timeout only
/ip arp print
shows gateway (93.83.73.69) ip address in the list but MAC is empty
During all this, sniffer is only showing:
SRC-MAC   DST-MAC            SRC-ADDRESS                     DST-ADDRESS
Ether1MAC FF:FF:FF:FF:FF:FF  93.83.73.70 who has 93.83.73.69? [empty, nothing]
Just for fun I disabled ARP on ether 1, and than sniffer starts showing DNS requests and pings from the computer, and src-address is 77.67.57.215 as expected because of src-nat. But they go nowhere.
/tool sniffer)?/ping 93.83.73.69 src-address=93.83.73.70directly from the router, that's the first thing you need to get working. If there's no response, check whether the ISP gateway's MAC address at least shows up under/ip arp print, and/or ping it witharp-ping=yes. (It's possible that the ISP's gateway will ignore regular ICMP ping, but it really cannot ignore ARP – it's always going to respond to ARP queries from the hAP.)/tool sniffer quick interface=ether1. Check if you can/ping 8.8.8.8 src-address=192.168.88.1from another terminal, this should get SNATed and the sniffer should show the packets as being sent from 77.67.57.215. 3) Try to ping 8.8.4.4 from another computer, this should also show up on the sniffer in the same way.
