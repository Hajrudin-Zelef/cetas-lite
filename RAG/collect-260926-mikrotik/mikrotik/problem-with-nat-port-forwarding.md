---
id: collect-260926-mikrotik/mikrotik/problem-with-nat-port-forwarding
title: "problem-with-nat-port-forwarding"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/problem-with-nat-port-forwarding.md
source_anchor: ""
source_lines: [1, 27]
sha256: bf05412a273411cf6bdb41f362e604056a33a248b1ec5d60cf265062dd2e676e
---

# problem-with-nat-port-forwarding

its worked thank you alot i went with these config======= thank you again

6. NAT AND IP FIREWALL RULES —> normally consists of three rules.

For basic destination nat which includes basic port forwarding, if coming from other devices, one needs ONE firewall rule in the forward firewall filter chain. The rule (example below) basically allows port forwarding to a server on the LAN, that has a corresponding DST NAT rule in the config. Then you need DST-NAT rules fore every port forwarding, each with all the details. A source nat rule is a more general rule so that any traffic initiated by the LAN is natted by the Routers WANIP on the way out. Thus source-nat is important only if local users are directed to the server (not through LANIP) but through the WANIP.

a. Filter Rule: add chain=forward action=accept connection-nat-state=dstnat

{ any traffic with destination ports identified in DST NAT rules will be allowed through firewall }

b. Source Nat Rule:

case1: Dynamic WANIP add chain=srcnat action=masquerade out-interface-list=WAN { will also work for fixed/static WANIPs but not as technically correct }

case2: Fixed/Static WANIP add chain=srcnat action=src-nat out-interface=ether1 to-addresses=WANIP(static) { where out interface must be the active interface, pppoe1-out, vlan etc. }

c. Destination Nat Rule:

case1: Dynamic WANIP add chain=dstnat action=dst-nat dst-port=xxxx protocol=yyy in-interface-list=WAN 

to-addresses=IPof Server { to ports not required if same as dst-ports }

case2: Fixed Static WANIP add chain=dstnat action=dst-nat dst-address=WANIP(static) dst-port=xxxx 

protocol=yyy to-addresses=IPofServer to-ports=zzzz { in this case users come in on port xxxx but the port gets translated to zzzz before hitting the server }

c. Destination Port Ranges: Caution, ensure port ranges do NOT overlap with expected incoming VPN listenting ports as DST nat takes precedence over input chain rules.
