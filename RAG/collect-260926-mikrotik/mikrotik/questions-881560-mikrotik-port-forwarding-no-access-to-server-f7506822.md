---
id: collect-260926-mikrotik/mikrotik/questions-881560-mikrotik-port-forwarding-no-access-to-server-f7506822
title: "questions-881560-mikrotik-port-forwarding-no-access-to-server-f7506822"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-881560-mikrotik-port-forwarding-no-access-to-server-f7506822.md
source_anchor: ""
source_lines: [1, 15]
sha256: 60f7f8940fd1c579a97353f1fce2a6005bd1744c39ff01e462c81fb87a593ba1
---

# questions-881560-mikrotik-port-forwarding-no-access-to-server-f7506822

I need your help concerning port forwarding with Mikrotik's RouterOS!
I have executed this command :
/ip firewall nat add action=dst-nat chain=dstnat  comment="Owncloud in NAT" dst-port=8079 in-interface=ether1 protocol=tcp to-addresses=172.16.0.50 to-ports=8079
My server is plugged in my ether3
It opened the port in question, but the request is not going to the server so when I try to open the web page hosted on that port I receive a 404 error.
If you are missing information, don't be shy to ask me to add more details! I am not used to debug Mikrotik product, since this is my first one!
** UPDATE **
I am not able to access my server from the WAN interface, I am able to access my server inside my network.
** SOLUTION **
I solved my problem, I have noticed that if I use the LTE on my phone to access my server, it works. So I added a static DNS to access my server inside my network
I did :
ip dns> set servers=.8.8.8.8,8.8.4.4 \
\... allow-remote-requests=yes
to set google dns servers and then I added my static dns :
ip dns static> add name=www.exemple.com address=172.16.0.50
