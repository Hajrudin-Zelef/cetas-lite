---
id: collect-260926-mikrotik/mikrotik/questions-1388195-mikrotik-how-to-open-a-port-fd9b3947
title: "dec/26/2018 15:12:03 by RouterOS 6.43.4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1388195-mikrotik-how-to-open-a-port-fd9b3947.md
source_anchor: ""
source_lines: [1, 22]
sha256: 9d395f2e21045872b38a449dc11b1bac43083702b8458b2578ba839cd8ca053a
---

# dec/26/2018 15:12:03 by RouterOS 6.43.4

I have recently bought a Mikrotik hAP ac^2 and I am trying to figure out how to connect to my NUC QBittorrent from "outside". This is all new to me...
My ISP has bridged its router and I connected my new Mikrotik to it and configured some basic things like WiFi.
I found my public DNS name: XXXXXXXXcd.sn.mynetname.net.
These are my settings:
/ip firewall export
# dec/26/2018 15:12:03 by RouterOS 6.43.4
# software id = EBLA-R903
#
# model = RBD52G-5HacD2HnD
# serial number = XXXXXXXXXX
/ip firewall filter
add action=accept chain=input protocol=icmp
add action=accept chain=input connection-state=established,related
add action=accept chain=forward connection-state=established,related
add action=drop chain=input in-interface=ether1
add action=drop chain=forward connection-state=invalid
add action=drop chain=forward connection-nat-state=!dstnat connection-state=new in-interface=ether1
/ip firewall nat
add action=masquerade chain=srcnat out-interface=ether1
add action=dst-nat chain=dstnat comment="QBitTorrent Web UI" dst-port=8100 in-interface=ether1 protocol=tcp to-addresses=192.168.88.239 to-ports=8100
My NUC is 192.168.88.239. When I try to connect to XXXXXXXXXcd.sn.mynetname.net:8100 my connection is Unable to connect. I figured out I should be using ether1 because I connected my bridged router to it and confirmed it by unplugging and checking the state in Interfaces tab of my Mikrotik...
What am I doing wrong?
