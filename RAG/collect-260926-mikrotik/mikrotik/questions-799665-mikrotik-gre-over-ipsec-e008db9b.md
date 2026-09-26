---
id: collect-260926-mikrotik/mikrotik/questions-799665-mikrotik-gre-over-ipsec-e008db9b
title: "questions-799665-mikrotik-gre-over-ipsec-e008db9b"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-799665-mikrotik-gre-over-ipsec-e008db9b.md
source_anchor: ""
source_lines: [1, 26]
sha256: 0e7923a05a24ccf2ecae38f7b66d901f4c1346dc1eaad59fe9f39497075dc7ff
---

# questions-799665-mikrotik-gre-over-ipsec-e008db9b

I'm trying to establish a GRE over IPSec tunnel between two MikroTik devices. Everything seems to work yet when I sniff the WAN interface I can clearly see the GRE packets which theoretically I shouldn't be able to see.
I've spent a few days on this and I'm at a loss on whats missing.
1.1.1.1 is the datacenter WAN, while 2.2.2.2 is the home WAN.
Router 1:
/interface gre
add allow-fast-path=no !keepalive local-address=1.1.1.1 name=\
    gre-tunnel-home remote-address=2.2.2.2
/ip ipsec peer
add address=2.2.2.2/32 dh-group=modp8192 enc-algorithm=blowfish \
    hash-algorithm=sha512 lifetime=30m local-address=1.1.1.1 \
    nat-traversal=no proposal-check=strict secret=secretcode
/ip ipsec policy
add dst-address=2.2.2.2/32 proposal=proposal1 sa-dst-address=2.2.2.2 \
    sa-src-address=1.1.1.1 src-address=1.1.1.1/32 tunnel=yes
Router 2:
/interface gre
add allow-fast-path=no !keepalive local-address=2.2.2.2 name=\
    gre-tunnel-datacenter remote-address=1.1.1.1
/ip ipsec peer
add address=1.1.1.1/32 dh-group=modp8192 enc-algorithm=blowfish \
    hash-algorithm=sha512 lifetime=30m local-address=2.2.2.2 \
    nat-traversal=no proposal-check=strict secret=secretcode
/ip ipsec policy
add dst-address=1.1.1.1/32 proposal=proposal1 sa-dst-address=\
    1.1.1.1 sa-src-address=2.2.2.2 src-address=2.2.2.2/32 \
    tunnel=yes
