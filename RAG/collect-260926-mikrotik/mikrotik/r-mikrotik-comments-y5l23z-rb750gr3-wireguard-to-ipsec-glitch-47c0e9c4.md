---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-y5l23z-rb750gr3-wireguard-to-ipsec-glitch-47c0e9c4
title: "r-mikrotik-comments-y5l23z-rb750gr3-wireguard-to-ipsec-glitch-47c0e9c4"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-y5l23z-rb750gr3-wireguard-to-ipsec-glitch-47c0e9c4.md
source_anchor: ""
source_lines: [1, 13]
sha256: e374592a956175498d2701dd0abe27f1bfa0ef1b95e428cd9dd6102d1dca7a36
---

# r-mikrotik-comments-y5l23z-rb750gr3-wireguard-to-ipsec-glitch-47c0e9c4

RB750Gr3 WireGuard to IPsec Glitch 
        
    Has anyone used WireGuard and IPsec together successfully on a RG750Gr3 (hEX)?
I have a strange issue with my RB750Gr3. It has an IPsec site-to-site tunnel to another site (running RouterOS in a VM for now), which works fine in both directions, but there also are mobile clients connected to the RB750Gr3 via WireGuard. The problem is, whenever a packet is forwarded from the WireGuard interface to the second site via IPsec, there is a 90%+ chance of it being corrupted in some way (I get a very high "State Protocol Error" count). The other direction, meaning from IPsec to WireGuard, works fine.
After a lot of testing I found out, that this only happens when hardware acceleration on the RB750Gr3 is used. When using for example DES with any hashing function or AES with SHA-512, everything works flawlessly, but using for example AES with SHA-256, which can be hardware accelerated, the issue comes back immediately.
Section des commentaires
So their support responded very quickly and it turned out I should have read the latests patchnotes...
7.6beta10: "ipsec - fixed packet processing by hardware encryption engine on MMIPS devices"
So I upgraded to 7.6rc3 and surprise - the packet loss is gone :D
Smells like an MTU issue but like others have said, raise a ticket. Get a supout from each side of the tunnel.
Commentaire supprimé par le membre
No, a both sites there is a LAN subnet, and at the site where the RB750Gr3 is located the LAN devices are connected to it through its hardware offloaded bridge. Pinging from either LAN to the other LAN seems fine (packet loss below 1% when pinging at an interval of 50ms, which is to be expected I think).
Commentaire supprimé par le membre
