---
id: collect-261001-cisco/cisco/t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by-b0cf47c4-3
title: "t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by--b0cf47c4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-cisco/t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by--b0cf47c4.md
source_anchor: ""
source_lines: [196, 222]
sha256: c63ec6970bc43a91602efc106c19518e48befdc114897ac65a87bea54d5c4721
---

# t5-network-security-asa-loses-connectivity-to-hsrp-gateway-every-20-mi-fixed-by--b0cf47c4

- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 06:17 AM
@Scott12 Did you note the mac address in arp entry in moment of the issue ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-18-2026 08:51 AM
At the moment we have observed the ARP behavior during the incident windows, however the entry appears to become unreachable from the ASA perspective rather than showing a clear MAC change or flapping condition.
After clearing the ARP table, the ASA relearns the gateway MAC correctly and connectivity is immediately restored, which suggests the issue is related to ARP resolution/refresh rather than a persistent incorrect MAC entry.
We are currently working on capturing the exact ARP/MAC state at the precise moment of impact and correlating it with the switching side (MAC address-table) to determine if there is any inconsistency or asymmetric forwarding condition in the path.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2026 01:55 AM
Hello, have you resolved this issue? I have a guess: did you set the ASA next-hop to the HSRP real address on the Nexus? Will the issue recur?
