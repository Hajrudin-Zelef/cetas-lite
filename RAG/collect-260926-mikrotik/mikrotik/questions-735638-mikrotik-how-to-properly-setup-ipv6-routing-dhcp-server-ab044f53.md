---
id: collect-260926-mikrotik/mikrotik/questions-735638-mikrotik-how-to-properly-setup-ipv6-routing-dhcp-server-ab044f53
title: "questions-735638-mikrotik-how-to-properly-setup-ipv6-routing-dhcp-server-ab044f53"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-735638-mikrotik-how-to-properly-setup-ipv6-routing-dhcp-server-ab044f53.md
source_anchor: ""
source_lines: [1, 26]
sha256: 813c8ecb0207269c291fe8b44199f39f9a9874af65fc9f916a25943036153a4f
---

# questions-735638-mikrotik-how-to-properly-setup-ipv6-routing-dhcp-server-ab044f53

I have the following configuration in my network:
ISP Router <- eth1 -> Mikrotik Router <- eth2-5 -> Internal network
My ISP router receives a /64 IPV6 block that is dynamic, changing whenever I reboot it.
The issue is that the devices in my internal network aren't receiving an IPV6 address, probably because the DHCP server in Mikrotik isn't configured correctly.
I already configured the DHCP client in Mikrotik ( IPV6 -> DHCP client ) and I can see the /64 block that was assigned by my ISP.
But I can't configure the DHCP Server. I tried to assign the pool in "IPV6 -> DHCP Server -> Add new" but my devices still doesn't receive any IPV6 addresses.
Any ideas about what is wrong? I'm using version 6.33.
Note: If i connect a device directly to my ISP router, IPV6 works correctly.
--- Edit ---
After Michael comment I was able to make some progress. I disabled the DHCP Server in Mikrotik and added a new address in "IPV6 -> Addresses" with advertise flag and using the pool that is created by the DHCP client.
All devices in my internal network are receiving IPV6 addresses and can ping each other. Unfortunately I can't ping external hosts.
In Mikrotik I can only ping external hosts only if I check the option "Request address" in /ipv6 dhcp-client. Otherwise ping doesn't work in Mikrotik.
Example: trying to ping ipv6.google.com.
/ping 2800:3f0:4001:801::200e
  SEQ HOST                                     SIZE TTL TIME  STATUS
    0 2800:3f0:4001:801::200e                                 timeout
    1 2800:3f0:4001:801::200e                                 timeout
    2 2800:3f0:4001:801::200e                                 timeout
    sent=3 received=0 packet-loss=100%
Here is how my route looks like:
/ipv6 route> print
Flags: X - disabled, A - active, D - dynamic, C - connect, S - static, r - rip, o - ospf, b - bgp, U - unreachable
 #      DST-ADDRESS              GATEWAY                  DISTANCE
 0 ADS  ::/0                     fe80::e297:96ff:fe6a:...        1
 1 ADC  2001:1284:****:90dd::/64 ether2-master-local             0
 2  DSU 2001:1284:****:90dd::/64                                 1
