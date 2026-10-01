---
id: collect-261001-fortinet/fortinet/dharnn-fortigate-ssl-vpn-segmentation-lab-4e5ce6a3-2
title: "dharnn-fortigate-ssl-vpn-segmentation-lab-4e5ce6a3"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/dharnn-fortigate-ssl-vpn-segmentation-lab-4e5ce6a3.md
source_anchor: ""
source_lines: [91, 95]
sha256: 15b0079cc000ef18b2a6ebd92ded5c8a8123352e4af506346485d9668436acb2
---

# dharnn-fortigate-ssl-vpn-segmentation-lab-4e5ce6a3

- Compare what works against what doesn't. Web Mode succeeding on the identical IP/port was the single most useful clue — it eliminated an entire category of possible causes (network, cert, NAT, basic auth) before any debug command was even run.
- MySQL error codes matter. 1044 vs1045 is the difference between a privilege problem and a password problem — worth knowing before reaching formysql_secure_installation again.
- "Latest" isn't always compatible. Pulling the newest release of any software against an older underlying stack (PHP, in this case) is a common, avoidable source of otherwise-confusing failures.
- Routing and tunnel interfaces are independent of tunnel health. A healthy IKE/IPsec SA doesn't guarantee traffic is actually being routed into the tunnel — that's a separate, static-route-dependent decision the firewall makes per-packet.
Built in EVE-NG as a self-directed lab exercise. Screenshots in /images.
