---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1825935-opnsense-wont-forward-ports-from-wan-to-lan-7beff283
title: "OPNsense won&#39;t forward ports from WAN to LAN"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1825935-opnsense-wont-forward-ports-from-wan-to-lan-7beff283.md
source_anchor: ""
source_lines: [1, 29]
sha256: 0bf25490580e4aee857fc377332f85104c1f2fc28d76335c2140baf03e6acb02
---

# OPNsense won&#39;t forward ports from WAN to LAN

*Score : 1 | Source : https://superuser.com/questions/1825935/opnsense-wont-forward-ports-from-wan-to-lan*

I cannot get my xmpp client's ports to be forwarded from the WAN side of my FW to the LAN side chat server.
My configuration is:
Aliases:
xmpp_port=5222
xmpp_server=chat 
Firewall-->Port Forward
Interface: WAN
TCP/IP Version: IPV4
Protocol: TCP
Source: any
Source port range: any 
Destination: xmpp_server
Destination port range: xmpp_port
Redirect target: xmpp_server
Redirect target port: xmpp_port
Pool Options: Default
Nat reflection: Enable
Filter rule association: Rule
==== The #opnsense irc channel recommended I try:
Source port range: any
Destination: This Firewall
That didn't fix the problem.
Source port range set to xmpp_port.
I have done much searching on the internet and my config looks OK, so my hunch is that I have done something elsewhere that has messed up forwarding even though I do not have a complicated set up. Everything else seems to be functioning well.
Output from ifconfig -a|grep inet on my OPNsense firewall listed my LAN side gateway address among other addresses, none of which were my chat server.
