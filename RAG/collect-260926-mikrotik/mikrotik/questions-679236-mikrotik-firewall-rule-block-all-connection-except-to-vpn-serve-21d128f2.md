---
id: collect-260926-mikrotik/mikrotik/questions-679236-mikrotik-firewall-rule-block-all-connection-except-to-vpn-serve-21d128f2
title: "questions-679236-mikrotik-firewall-rule-block-all-connection-except-to-vpn-serve-21d128f2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-679236-mikrotik-firewall-rule-block-all-connection-except-to-vpn-serve-21d128f2.md
source_anchor: ""
source_lines: [1, 19]
sha256: 56cb168897778736772154b88988bdb20aa757f5a63555d94c94a3613a13c419
---

# questions-679236-mikrotik-firewall-rule-block-all-connection-except-to-vpn-serve-21d128f2

PPTP uses
- TCP port 1723 
- GRE (protocol ID 47) for tunneling
Accept PPTP in Mikrotik:
/ip firewall filter add chain=input action=accept protocol=tcp dst-port=1723
/ip firewall filter add chain=input action=accept protocol=gre
L2TP/IPSec uses 
- TCP port 1701 
- UDP port 500 for Security Association (SA) - to negotiate security method (password, certificate, kerberos)
- AH (Protocol ID 50) - Authentication Header
- ESP (Protocol ID 51) - Encapsulated Secure Payload
Accept L2TP/IPSec in Mikrotik:
/ip firewall filter add chain=input action=accept protocol=tcp dst-port=1701
/ip firewall filter add chain=input action=accept protocol=udp dst-port=500
/ip firewall filter add chain=input action=accept protocol=ipsec-ah
/ip firewall filter add chain=input action=accept protocol=ipsec-esp
Block all other incoming connection (TCP)
/ip firewall filter add chain=input protocol=tcp action=reject reject-with=tcp-reset
You can use action=drop instead of reject, but according to Hannes Schmidt, NMAP can still see the port is open but dropped (filtered) by firewall
