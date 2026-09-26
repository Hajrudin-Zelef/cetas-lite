---
id: collect-260926-mikrotik/mikrotik/questions-269769-manually-set-ipv6-neighbors-mac-address-in-mikrotiks-routeros-9542c757
title: "questions-269769-manually-set-ipv6-neighbors-mac-address-in-mikrotiks-routeros-9542c757"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/questions-269769-manually-set-ipv6-neighbors-mac-address-in-mikrotiks-routeros-9542c757.md
source_anchor: ""
source_lines: [1, 4]
sha256: a93b5d139af66614746f3336887fc6ed93734a5dd0df1d36bff59a8755b05509
---

# questions-269769-manually-set-ipv6-neighbors-mac-address-in-mikrotiks-routeros-9542c757

I have an interesting problem. An ISP of ours started providing native IPv6 on their network. They provided us with a /56 prefix and a /126 linking segment (::1 being their endpoint and ::2 being ours). However, the ISP's DSLAM does not support Neighbor Discovery, so they asked us to provide our router's MAC address to manually enter in their neighbor table and provided us with theirs.
We are using Mikrotik's RouterOS 5.2 on our router and we have not been able to figure out how this can be done. For IPv4, it is simple, we can create a mapping under /ip arp, but for IPv6 the neighbor list (/ip neighbor) seems to be read-only. 
We tried to work this around, by setting a fe80:: link-local address with the ISP's router's MAC address encoded in the EUI-64, but when we sent packets to that destination, RouterOS tried to resolve the address using ND nevertheless.
I am asking if there is a way to either create a static mapping IPv6 address -> MAC address or to create a route directly to a MAC address and the interface, so we can stop using tunnels for our IPv6 needs.
