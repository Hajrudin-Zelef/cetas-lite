---
id: collect-261001-meraki/meraki/r-meraki-comments-crvims-core-switch-arp-issue-7c2cc0ab
title: "r-meraki-comments-crvims-core-switch-arp-issue-7c2cc0ab"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-crvims-core-switch-arp-issue-7c2cc0ab.md
source_anchor: ""
source_lines: [1, 11]
sha256: b404aaa0a4f9bb85b852ed3774f8bc6f5aae255a6d2dd6b4e209e344558696fc
---

# r-meraki-comments-crvims-core-switch-arp-issue-7c2cc0ab

We have a pair of MS425-32 that handles Layer 3 routing and holds all the VLAN interfaces.

Seems that FW 11.22 has an issue where the ARP table doesn't sync between the two switches, and it causes some very hard to troubleshoot issues.

I started to suspect a FW bug when a device would "wake up" and become pingable if I ping it from a device on the same VLAN.

I would have devices fall offline or maintain a certain packet loss. I guess it depends on if that packet passes through the switch with the ARP entry or not.

Support says we have to wait for a software fix. I got so tired of the packet loss I just powered one of the core switches down until they fix it.

Anyone else experiencing this?
