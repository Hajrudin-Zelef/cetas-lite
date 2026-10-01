---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-39401-ubiquiti-edgerouter-pro-erpro-8-as-bgp-1-gbps-endpoint-c9f8222e
title: "questions-39401-ubiquiti-edgerouter-pro-erpro-8-as-bgp-1-gbps-endpoint-c9f8222e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "license", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-39401-ubiquiti-edgerouter-pro-erpro-8-as-bgp-1-gbps-endpoint-c9f8222e.md
source_anchor: ""
source_lines: [1, 4]
sha256: e3c0bfeb036c9f3b81b2b6272536bb2f210d16cb8dff722b7200be607245529f
---

# questions-39401-ubiquiti-edgerouter-pro-erpro-8-as-bgp-1-gbps-endpoint-c9f8222e

My company has had dual 100 Mbps Internet connections running BGP on two Cisco 2921 routers for several years. It works as expected... when there are issues on one line, BGP cuts over properly. If we want to do maintenance on one, we manually flip BGP over without missing a beat, and all is well. We are only getting default routes from our providers, so memory is not a concern.
We are about to upgrade one of the connections to 1 Gbps, but the 2921 routers cannot pass traffic at that speed. I am considering getting a Cisco 3560CX switch with an IP Services license which would work fine. However, I'm very curious to try a Ubiquiti EdgeRouter Pro (ERPro‑8) instead. I'm just so shocked as to how cheap it is, and it makes me wonder if something is wrong. It would be a straight-up replacement for one of the 2921s, so I wouldn't be doing any firewall, NAT, or any other fancy functionality.
Can you think of any reason why this sub-$400 router would not work as well for basic BGP routing as a Cisco 3560-CX IP Services (about $2,500) or a Cisco 4431 with the high-speed license that would cost about $10,000? Is it unreliable? Does it only last for a year before exploding? How can this be?!
Thanks.
