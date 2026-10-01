---
id: collect-261001-fortinet/fortinet/questions-609987-fortigate-100d-802-3ad-bonding-link-aggregation-de28ea7b
title: "questions-609987-fortigate-100d-802-3ad-bonding-link-aggregation-de28ea7b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "ethernet"]
source: docs/RAG/collect-261001-fortinet/questions-609987-fortigate-100d-802-3ad-bonding-link-aggregation-de28ea7b.md
source_anchor: ""
source_lines: [1, 6]
sha256: d756f1688b0a82ee7d1c5efdeb6d42740c1220af44a5032b69d62bd029ae81c3
---

# questions-609987-fortigate-100d-802-3ad-bonding-link-aggregation-de28ea7b

I'm going to go with the assumption (rather than asking, like @Shane Madden did) and assume you don't have your own address space and are just using IP addresses assigned by the ISP on both WAN links.
802.3ad is a layer 2 link aggregation protocol. It won't help you at all in the scenario you're describing. 802.3ad would be useful if you had, say, multiple metro-Ethernet terminations from the same ISP, all using the same IP space.
The "load-balancing" functionality built-in to the Fortigate devices assigns TCP connections to a WAN link (through various means-- weights assigned to interfaces, by source address, or by utilization). You're not going to see more than the bandwidth of a single WAN connection utilized for a single TCP connection, but you will see TCP connections spread across both WAN connections. This isn't wholly ineffective if your main concern is utilizing the speed of both interfaces for users accessing web sites (or other Internet resources).
Here's FortiNet's documentation describing the feature in more detail.
If you had your own IP address space, and were peered with your ISPs using BGP, you would be able to send and receive traffic across both WAN connections. (Most small to medium-sized businesses don't have this option, though.)
If your concern is inbound redundancy (like an on-site hosted server being accessible via the same public IP address via both ISPs) then you are going to need to look to getting your own address space and peering with your ISPs. (Typically, you're talking about a whole different level of expense, too, because consumer-grade ISPs don't offer this type of functionality.)
