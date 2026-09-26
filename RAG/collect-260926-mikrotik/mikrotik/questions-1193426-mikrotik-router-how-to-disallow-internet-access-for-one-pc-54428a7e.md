---
id: collect-260926-mikrotik/mikrotik/questions-1193426-mikrotik-router-how-to-disallow-internet-access-for-one-pc-54428a7e
title: "questions-1193426-mikrotik-router-how-to-disallow-internet-access-for-one-pc-54428a7e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1193426-mikrotik-router-how-to-disallow-internet-access-for-one-pc-54428a7e.md
source_anchor: ""
source_lines: [1, 9]
sha256: c27810cff3ecc98ca8c4aa0b07aeee8f492eb72b8dd8f5e5f086fabf6d10991d
---

# questions-1193426-mikrotik-router-how-to-disallow-internet-access-for-one-pc-54428a7e

This can be translated almost directly to firewall rules:
/ip firewall filter {
- allow from PC-2 to LAN: add chain=forward src-address=<PC2_IP> dst-address=<LAN_SUBNET> action=accept
- deny from PC-2 to everywhere else:  add chain=forward src-address=<PC2_IP> action=reject
Which can also be combined:
}
Here <LAN_SUBNET> should be the prefix you want to allow, e.g. 192.168.88.0/24 for the IPv4 rule, or 2001:db8:abcd:0::/64 for IPv6.
The rule checking goes from top to bottom until first match, so make sure the rule goes after "allow established" but before any "allow  everything" rules you might have.
Note: Within the same subnet, access will always be allowed, as communications only go through the built-in switch and don't reach the OS. (Although RouterOS allows overriding that if necessary – under /interface ethernet switch rule, you can find an option to redirect packets from PC-2 to the OS as well. However, it's generally best to assume that intra-subnet traffic is unfiltered.)
