---
id: collect-260926-mikrotik/mikrotik/site-to-site-wireguard-firewall-rules
title: "site-to-site-wireguard-firewall-rules"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/site-to-site-wireguard-firewall-rules.md
source_anchor: ""
source_lines: [1, 45]
sha256: e3a41f8f2701a2059209a4fdfafa192947e0d89e12e6972c3e6bdf32403c9c79
---

# site-to-site-wireguard-firewall-rules

I’ve been struggling with this all day, I feel like I’m missing something fundamental.

I have two identical RB5009 set up across the world from each other, each running 7.1.1. I have had wireguard running on each of them successfully for several months now, but just using mobile phones, laptops (192.168.100.0/24). I followed this guide to set up site to site, and finding that firewall rules don’t seem to affect the traffic between sites.

I can ping between subnets from ‘Office 1’ to ‘Office 2’, configured as shown:

My firewall config is below (most is defconf).  I expect Wireguard traffic to come in through rule 2, then go through the firewall rules again as local traffic. I expect that if the traffic is from my mobile Wireguard clients (each a /32 address in the 192.168.100.0/24 subnet), it is accepted and that if it’s from my other site (Office 2), then it will be dropped. However, all traffic gets through with this config.

I wonder if, because it’s two peers, that the traffic I’m expecting to be filtered in the firewall is considered an established connection due to the Office 1 to Office 2 tunnel? But it shouldn’t be, ‘Office 2’ traffic should be routed through its own Wireguard interface and not back through the already established connection from ‘Office 1’, right?

Firewall (Office 1)

```
/ip firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=accept chain=input comment="wireguard server" dst-port=13231 in-interface-list=WAN protocol=udp
add action=accept chain=forward comment="wireguard mobile subnet to LAN" dst-address=172.16.0.0/16 src-address=192.168.100.0/24
add action=accept chain=input comment="defconf: accept ICMP" in-interface-list=!WAN protocol=icmp
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept to local loopback (for CAPsMAN)" dst-address=127.0.0.1
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN log=yes
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related hw-offload=yes
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf: drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none out-interface-list=WAN
```

IP addresses (Office 1)

```
/ip address
add address=172.16.0.1/16 comment=defconf interface=bridge network=172.16.0.0
add address=192.168.100.1/24 interface=wireguard1 network=192.168.100.0
add address=10.255.255.1/30 comment="wireguard site-to-site bridge" interface=wireguard1 network=10.255.255.0
```

Routes (Office 1)

```
/ip route
add disabled=no distance=1 dst-address=192.168.0.0/24 gateway=10.255.255.2 pref-src=0.0.0.0 routing-table=main scope=30 \
    suppress-hw-offload=no target-scope=10
```
