---
id: collect-261001-meraki/meraki/r-meraki-comments-8445ji-meraki-stack-core-switching-mac-table-arp-issue-6b939cb9
title: "r-meraki-comments-8445ji-meraki-stack-core-switching-mac-table-arp-issue-6b939cb9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-8445ji-meraki-stack-core-switching-mac-table-arp-issue-6b939cb9.md
source_anchor: ""
source_lines: [1, 3]
sha256: eec8fa5b136f92acedfe5ac86e404101c55afd489b1995abba79f663cb0c2219
---

# r-meraki-comments-8445ji-meraki-stack-core-switching-mac-table-arp-issue-6b939cb9

I have had numerous instances where meraki core switches when booted at the same time are showing odd issues. Most common issue is that some internal devices can communicate outbound, some can't. Changing either IP or mac address of endpoint resolves the problem. Rebooting core switches - one at the time - resolves the problem. We observe this issue only when meraki core stack is used for inter-VLAN routing.

Issue is very annoying because any power outage leads into very bizarre networking issues for which only complete solution is to reboot core switches, one at the time. Anyone else observed this issue?
