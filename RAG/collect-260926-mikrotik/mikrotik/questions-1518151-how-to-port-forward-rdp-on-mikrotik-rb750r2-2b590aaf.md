---
id: collect-260926-mikrotik/mikrotik/questions-1518151-how-to-port-forward-rdp-on-mikrotik-rb750r2-2b590aaf
title: "questions-1518151-how-to-port-forward-rdp-on-mikrotik-rb750r2-2b590aaf"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1518151-how-to-port-forward-rdp-on-mikrotik-rb750r2-2b590aaf.md
source_anchor: ""
source_lines: [1, 21]
sha256: 9a83aae9457d65628b0de5b10b2e9e9c5720cfb332992004e3fc892177d1923f
---

# questions-1518151-how-to-port-forward-rdp-on-mikrotik-rb750r2-2b590aaf

Not a power user. I got a new Mikrotik router with the latest firmware. I want to use it on my home PC with RDP. On the previous cheap Tenda router, port forwarding worked fine because the settings were simple. I read a lot of manuals/guides about port forwarding in RouterOS and tried to configure this by myself, but nothing happened. I can’t connect from the outside. Can someone please write in detail how to configure this correctly?
2 Answers 2
Please read my advice below before implementing this..
/ip firewall nat
add action=dst-nat chain=dstnat comment="RDP" disabled=no dst-port=3389 protocol=tcp to-addresses=(PRIVATE IP of RDPhost) to-ports=3389
It has become standard security practice to not use port forwarding for RDP as you will most likely be hacked.
Please at the least read up here how to secure with a source address here, also if you read up and see people telling you to change to an non standard port to secure RDP this just adds another 1% security not much.
- 
        Thanks! this worked, im going to be accessing from 4g hotspot so i cant use a source address as its not a set ipjunior– junior2020-01-18 02:21:59 +00:00Commented Jan 18, 2020 at 2:21
- 
            
            
- 
        @junior if you youtube No-IP and mikrotik setup theres alot out thereSQLTemp– SQLTemp2020-01-18 02:25:37 +00:00Commented Jan 18, 2020 at 2:25
- 
        You can setup a VPN tunnel instead.eckes– eckes2020-01-19 13:41:59 +00:00Commented Jan 19, 2020 at 13:41
Port Forwarding is a form of NAT. NAT interfaces do not provide much, if any, protection. The Microtik Router you have supports IPsec VPN and you could set up an IPsec connection and use an IPsec Client Application to access your home machine. I do this myself with a Cisco Router.
- 
        do you have the source code for thisjunior– junior2020-01-18 02:21:04 +00:00Commented Jan 18, 2020 at 2:21
- 
        You can do all the above with commercial applications. It would be significant programming jobanon– anon2020-01-18 02:22:43 +00:00Commented Jan 18, 2020 at 2:22
