---
id: collect-260926-mikrotik/mikrotik/working-qos-configuration-1
title: "working-qos-configuration"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/qos/working-qos-configuration.md
source_anchor: ""
source_lines: [1, 117]
sha256: b46149228379eb7ff7d8a596b54a446bb7472df3d81b6a47adfc4f374a7f4dd5
---

# working-qos-configuration

Hi everybody,

I’ve been struggling for some time to create a working QoS configuration, reading through countless manual pages, wikis, forum topics etc. I’ve noticed two things; in order to do something great with a Mikrotik router, you have to be a network guru. I thought I knew relatively much about networking when I got my router, but it seems I was wrong. Second, it seems like there’s a tendency among those who actually do know something and who actually have working configurations to be pretty secretive about them, which is understandable, but very frustrating for novices. That’s why I thought I’d post my QoS configuration here as well as explain what it does (or at least what I want it to do).

1. **My network layout**

WAN connection: 110/5 Mbit (up/down). The actual downlink speed is limited to 100 Mbit by the fact that my RouterBOARD only has 100 Mbit Ethernet ports.

My stuff is hooked up like this:

Cable modem ← [bridge] → Routerboard → Switch(es) → computers.

1. **What I want to accomplish**

I’m a heavy BitTorrent user, which means I’m pretty much maxing out my upstream bandwidth all the time. I also have my own FTP server which some of my friends use to download some stuff from me, a HTTP server from which stream music from my collection when I’m at other peoples places (I have my own home-made system for this) and finally I’m always connected to my computer at home using SSH.

What happens with no QoS at all is that I constantly have to limit the torrent clients upload speed everytime I come somewhere and want to stream something, or a friend needs to get something fast over FTP. It also means that unless I cap the upload speed in the torrent client I get incredible lag when using SSH to chat on IRC (for example). Of course web browsing is a PITA too when seeding at full speed…

1. **The solution**

The solution to all this is to mark traffic from the most critical to the non-critical. I got most of my rules from this guide. This is my current mangle config and my queue tree config. Just to clearify: “Public” is the interface connected to the modem and “Local” is the interface connected to the switch. The key seems to be to mark the ACK packets correctly (only marking the ones with a size of 0-80 bytes like most people suggest is not sufficient as the ACK flag can be set on bigger packets as well).

Mangle rules:

```
[xxx@xxx] /ip firewall mangle> print
Flags: X - disabled, I - invalid, D - dynamic 
 0   ;;; Link-critical traffic (ARP, DHCP)
     chain=postrouting action=mark-packet new-packet-mark=link_critical passthrough=no 
     protocol=udp out-interface=Public src-port=68 dst-port=67 
 1   ;;; Time-critical traffic (DNS, control packets)
     chain=postrouting action=mark-packet new-packet-mark=time_critical passthrough=no 
     protocol=udp out-interface=Public dst-port=53 
 2   chain=postrouting action=mark-packet new-packet-mark=time_critical passthrough=no 
     tcp-flags=fin,syn,rst protocol=tcp out-interface=Public 
 3   chain=postrouting action=mark-packet new-packet-mark=time_critical passthrough=no 
     tcp-flags=ack protocol=tcp out-interface=Public packet-size=40-89 
 4   chain=postrouting action=mark-packet new-packet-mark=time_critical passthrough=no 
     connection-state=new protocol=tcp out-interface=Public 
 5   ;;; Critical traffic
     chain=postrouting action=mark-packet new-packet-mark=critical passthrough=no 
     tcp-flags=ack protocol=tcp out-interface=Public packet-size=90-159 
 6   ;;; High-priority interactive traffic (SSH)
     chain=postrouting action=mark-packet new-packet-mark=high_pri_interactive 
     passthrough=no protocol=tcp out-interface=Public port=22,2200 
 7   chain=postrouting action=mark-packet new-packet-mark=high_pri_interactive 
     passthrough=no tcp-flags=ack protocol=tcp out-interface=Public packet-size=160-249 
 8   chain=postrouting action=mark-packet new-packet-mark=high_pri_interactive 
     passthrough=no protocol=tcp out-interface=Public port=8291 
 9   ;;; Low-priority interactive traffic (HTTP, HTTPS, DelugeWebUI)
     chain=postrouting action=mark-packet new-packet-mark=low_pri_interactive 
     passthrough=no protocol=tcp out-interface=Public port=80,443,8112 
10   chain=postrouting action=mark-packet new-packet-mark=low_pri_interactive 
     passthrough=no tcp-flags=ack protocol=tcp out-interface=Public packet-size=250-359 
11   ;;; High-priority non-interactive traffic (FTP)
     chain=postrouting action=mark-packet new-packet-mark=high_pri_non_interactive 
     passthrough=no protocol=tcp out-interface=Public connection-type=ftp 
12   chain=postrouting action=mark-packet new-packet-mark=high_pri_non_interactive 
     passthrough=no tcp-flags=ack protocol=tcp out-interface=Public packet-size=360-489 
13   ;;; Low-priority non-interactive traffic (POP, SMTP)
     chain=postrouting action=mark-packet new-packet-mark=low_pri_non_interactive 
     passthrough=no protocol=tcp out-interface=Public port=25,110 
14   chain=postrouting action=mark-packet new-packet-mark=low_pri_non_interactive 
     passthrough=no tcp-flags=ack protocol=tcp out-interface=Public packet-size=490-639 
15   ;;; Non-critical traffic (P2P)
     chain=postrouting action=mark-packet new-packet-mark=non_critical passthrough=no 
     tcp-flags=ack protocol=tcp out-interface=Public packet-size=640-809 
16   chain=postrouting action=mark-packet new-packet-mark=non_critical passthrough=no 
     protocol=tcp out-interface=Public
```

Queue tree:

```
[admin@Neggelandia] /queue tree> print
Flags: X - disabled, I - invalid 
 0   name="Link-critical" parent=Outgoing queue packet-mark=link_critical limit-at=0 
     queue=default priority=1 max-limit=4350000 burst-limit=0 burst-threshold=0 
     burst-time=0s 
 1   name="Time-critical" parent=Outgoing queue packet-mark=time_critical limit-at=0 
     queue=default priority=2 max-limit=4350000 burst-limit=0 burst-threshold=0 
     burst-time=0s 
 2   name="Critical" parent=Outgoing queue packet-mark=critical limit-at=0 queue=default 
     priority=3 max-limit=4350000 burst-limit=0 burst-threshold=0 burst-time=0s 
 3   name="High-pri interactive" parent=Outgoing queue packet-mark=high_pri_interactive 
     limit-at=0 queue=default priority=4 max-limit=4350000 burst-limit=0 
     burst-threshold=0 burst-time=0s 
 4   name="Low-pri interactive" parent=Outgoing queue packet-mark=low_pri_interactive 
     limit-at=0 queue=default priority=5 max-limit=4350000 burst-limit=0 
     burst-threshold=0 burst-time=0s 
 5   name="High-pri non-interactive" parent=Outgoing queue 
     packet-mark=high_pri_non_interactive limit-at=0 queue=default priority=6 
     max-limit=4350000 burst-limit=0 burst-threshold=0 burst-time=0s 
 6   name="Outgoing queue" parent=Public packet-mark="" limit-at=0 queue=default 
     priority=8 max-limit=4350000 burst-limit=0 burst-threshold=0 burst-time=0s 
 7   name="Low-pri non-interactive" parent=Outgoing queue 
     packet-mark=low_pri_non_interactive limit-at=0 queue=default priority=7 
     max-limit=4350000 burst-limit=0 burst-threshold=0 burst-time=0s 
 8   name="Non-critical" parent=Outgoing queue packet-mark=non_critical limit-at=0 
     queue=default priority=8 max-limit=4350000 burst-limit=0 burst-threshold=0 
     burst-time=0s
```

Queue types:

```
[xxx@xxx] /queue type> print
 0 name="default" kind=pfifo pfifo-limit=200 
 1 name="ethernet-default" kind=pfifo pfifo-limit=200
```

1. **Result**

With the current settings, if I’m seeding at full speed and somebody starts downloading something over FTP, the FTP transfer gets all upload bandwidth while the torrents almost stop. If I’m connected with SSH, there’s absolutely no lag when I’m writing in the console. When surfing the web, DNS lookups happen instantly and pages load just like if I wasn’t seeding at all.

Feel free to comment on my config if you have any suggestions that would make it better.

