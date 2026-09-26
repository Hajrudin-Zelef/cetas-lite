---
id: collect-260926-mikrotik/mikrotik/questions-181703-how-do-i-set-up-routeros-to-use-web-proxy-on-another-machine-8a4b82f0
title: "questions-181703-how-do-i-set-up-routeros-to-use-web-proxy-on-another-machine-8a4b82f0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-181703-how-do-i-set-up-routeros-to-use-web-proxy-on-another-machine-8a4b82f0.md
source_anchor: ""
source_lines: [1, 13]
sha256: 562096061f6045833f410e41c25164f7f0bf49d972327a3deeabc37c547f3ce3
---

# questions-181703-how-do-i-set-up-routeros-to-use-web-proxy-on-another-machine-8a4b82f0

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
I'd like to run a proxy on another machine, so that I can take advantage of more sophisticated filtering rules available in Squid or the like. However, if I use NAT to redirect traffic to another machine running Squid it won't work, since the HTTP request will need to be rewritten in order to be a proxy HTTP request; just redirecting the traffic gives bad request errors from Squid.
No need of setting proxy in RouterOS. You can route all outgoing HTTP traffic to the server directly thru NAT:
ip firewall nat add in-interface=eth1 src-address=!<IP of Squid machine> dst-port=80 protocol=tcp action=dst-nat to-addresses=<IP of Squid machine> to-ports=8080
The last parameter "src-address=!..." is needed in case which squid machine communicates thru same interface as the other machines. Otherwise it would go like this:
Computer sends HTTP request
RouterOS destinates this packet to squid
Squid sends HTTP request to webserver
RouterOS destinates squid request again to squid -> loop
/ip proxy
set parent-proxy=<IP of Squid machine> parent-proxy-port=3128
/ip firewall nat
chain=dstnat src-address=!<IP of Squid machine> protocol=tcp dst-port=80 src-address-list=<IP of Local machine> action=redirect to-ports=8080
