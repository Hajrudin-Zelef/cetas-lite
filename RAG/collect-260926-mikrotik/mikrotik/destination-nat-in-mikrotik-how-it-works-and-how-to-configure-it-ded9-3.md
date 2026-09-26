---
id: collect-260926-mikrotik/mikrotik/destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9-3
title: "destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9.md
source_anchor: ""
source_lines: [166, 170]
sha256: 7257fd87d9fd7597d1a97e0151f0fe0039d942bf058225f6c2a98707efa6a70e
---

# destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9

It enables access from outside your network—such as the internet—to servers or services inside your LAN by translating external requests to internal hosts and ports.

### How do you configure a simple DST-NAT rule?

In RouterOS, create a firewall NAT rule on the dstnat chain specifying the external destination IP/port and set action=dst-nat with to-addresses and optional to-ports for the internal service.
