---
id: collect-260926-mikrotik/mikrotik/questions-999610-how-to-forward-gre-protocol-47-on-mikrotik-router-3658e0ba
title: "questions-999610-how-to-forward-gre-protocol-47-on-mikrotik-router-3658e0ba"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-999610-how-to-forward-gre-protocol-47-on-mikrotik-router-3658e0ba.md
source_anchor: ""
source_lines: [1, 32]
sha256: dbcddb67f4aa481598e87f7c657ac9ea48e74109d82b700f03671f6e57cd5464
---

# questions-999610-how-to-forward-gre-protocol-47-on-mikrotik-router-3658e0ba

GRE protocol on its own is not enough.
Depending on which type of VPN service you are using you'll have to port forward some other TCP or UDP ports.
For PPTP for example you need to forward port 1723 TCP.
For L2TP you need to forward port 1701 UDP.
For OpenVPN you need to forward port 1194 UDP+TCP (OpenVPN does not use GRE).
Here is a port forwarding example to use: http://wiki.mikrotik.com/wiki/Manual:Initial_Configuration#Port_forwarding
  Port forwarding
  
  To make services on local servers/hosts available to general public it
  is possible to forward ports from outside to inside your NATed
  network, that is done from /ip firewall nat menu. For example, to make
  possible for remote helpdesk to connect to your desktop and guide you,
  make your local file cache available for you when not at location etc.
  Static configuration
  
  A lot of users prefer to configure these rules statically, to have
  more control over what service is reachable from outside and what is
  not. This also has to be used when service you are using does not
  support dynamic configuration.
  
  Following rule will forward all connections to port 22 on the router
  external ip address to port 86 on your local host with set IP address:
  
  if you require other services to be accessible you can change protocol
  as required, but usually services are running TCP and dst-port. If
  change of port is not required, eg. remote service is 22 and local is
  also 22, then to-ports can be left unset.
  
  Comparable command line command:
 /ip firewall nat add chain=dstnat dst-address=172.16.88.67
 protocol=tcp dst-port=22 \  action=dst-nat to-address=192.168.88.22
 to-ports=86
