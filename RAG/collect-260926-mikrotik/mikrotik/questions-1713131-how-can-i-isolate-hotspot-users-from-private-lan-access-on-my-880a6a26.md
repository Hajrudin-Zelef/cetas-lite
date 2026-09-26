---
id: collect-260926-mikrotik/mikrotik/questions-1713131-how-can-i-isolate-hotspot-users-from-private-lan-access-on-my-880a6a26
title: "questions-1713131-how-can-i-isolate-hotspot-users-from-private-lan-access-on-my--880a6a26"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1713131-how-can-i-isolate-hotspot-users-from-private-lan-access-on-my--880a6a26.md
source_anchor: ""
source_lines: [1, 11]
sha256: bd95d0493ce0bcfe238188178cdc6fd230ff2e21cf3ddf23d8b81af8479705c6
---

# questions-1713131-how-can-i-isolate-hotspot-users-from-private-lan-access-on-my--880a6a26

I am new to mikrotik , I have installed routerboard :
modem IP : 192.168.1.1 ====> to mikrotik port 1
mikrotik IP 192.168.1.10 gateway :192.168.1.1
bridge ports 2,3,4,5, as OUT port
Hotspot created :ip 10.10.10.1 pool : 10.10.10.0/24
Acess point 2 ===> to mikrotik port 2
Acess point 3 ===> to mikrotik port 3
Acess point 4 ===> to mikrotik port 4
Acess point 5 ===> to mikrotik port 5
everything is working fine, now if hotspot user is connected to the hotspot, they are able to login to any devices with 192.168.1.0
How to isolate bridge ports (hotspot) from Private network (IN) with allowing hotspot users to be connected to the internet only not the internal network?
