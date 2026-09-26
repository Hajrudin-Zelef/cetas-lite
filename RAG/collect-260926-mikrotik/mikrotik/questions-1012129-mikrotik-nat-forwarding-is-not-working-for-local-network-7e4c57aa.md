---
id: collect-260926-mikrotik/mikrotik/questions-1012129-mikrotik-nat-forwarding-is-not-working-for-local-network-7e4c57aa
title: "questions-1012129-mikrotik-nat-forwarding-is-not-working-for-local-network-7e4c57aa"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-1012129-mikrotik-nat-forwarding-is-not-working-for-local-network-7e4c57aa.md
source_anchor: ""
source_lines: [1, 8]
sha256: 1cd7431c78ede8cb6d06a49c6e1b3aa7fa655256273f4161ad992b1a0c8529e4
---

# questions-1012129-mikrotik-nat-forwarding-is-not-working-for-local-network-7e4c57aa

I try to access my external IP address from the local network, but instead of reaching my webserver behind NAT - the webfig page shows up.
NAT forwarding is working when accessing from the internet.
EDIT: This is my current nat config
[admin@MikroTik] > ip firewall nat print
Flags: X - disabled, I - invalid, D - dynamic
 0    ;;; defconf: masquerade
      chain=srcnat action=masquerade out-interface-list=WAN ipsec-policy=out,none
 1    chain=dstnat action=dst-nat to-addresses=<server's ip> to-ports=80 protocol=tcp in-interface=ether1 dst-port=80
