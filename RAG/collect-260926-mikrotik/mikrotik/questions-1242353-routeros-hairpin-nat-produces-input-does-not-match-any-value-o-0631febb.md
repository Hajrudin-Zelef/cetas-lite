---
id: collect-260926-mikrotik/mikrotik/questions-1242353-routeros-hairpin-nat-produces-input-does-not-match-any-value-o-0631febb
title: "questions-1242353-routeros-hairpin-nat-produces-input-does-not-match-any-value-o-0631febb"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-1242353-routeros-hairpin-nat-produces-input-does-not-match-any-value-o-0631febb.md
source_anchor: ""
source_lines: [1, 11]
sha256: d99d3b96d21592ba35168f50ef6916b298193f24118b500beb1e31070dc9042b
---

# questions-1242353-routeros-hairpin-nat-produces-input-does-not-match-any-value-o-0631febb

I have RouterOS v6.39.2 (stable)
- external IP: 134.249.116.246
- router IP: 192.168.88.1
- internal server IP: 192.168.88.245
I setup a regular port forwarding via
  ip firewall nat add chain=dstnat dst-address=134.249.116.246 protocol=tcp dst-port=8080 action=dst-nat to-addresses=192.168.88.245 to-ports=8080 
Attempt to setup a hairpin NAT
 ip firewall nat add chain=srcnat src-address=192.168.88.0/24 dst-address=192.168.88.245 protocol=tcp dst-port=8080 out-interface=LAN action=masquerade
failed due to:
input does not match any value of interface
I can't understand what is wrong
