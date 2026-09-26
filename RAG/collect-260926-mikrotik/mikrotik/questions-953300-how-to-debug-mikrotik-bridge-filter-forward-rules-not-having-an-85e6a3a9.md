---
id: collect-260926-mikrotik/mikrotik/questions-953300-how-to-debug-mikrotik-bridge-filter-forward-rules-not-having-an-85e6a3a9
title: "questions-953300-how-to-debug-mikrotik-bridge-filter-forward-rules-not-having-an-85e6a3a9"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-953300-how-to-debug-mikrotik-bridge-filter-forward-rules-not-having-an-85e6a3a9.md
source_anchor: ""
source_lines: [1, 6]
sha256: 8e97911b2054ff64636a328ac9021c568747bd2366af95aaa10d65c0b6f423f6
---

# questions-953300-how-to-debug-mikrotik-bridge-filter-forward-rules-not-having-an-85e6a3a9

As part of diagnosing a different problem were trying to add a bridge filter rule that will stop all traffic from forwarding between two interfaces on a bridge.
The router has two interfaces ether1 and ether2 on a bridge.
we then added a rule with this
/interface bridge filter
add action=drop chain=forward in-interface=ether1
i had expected this to stop all traffic that arrived on ether1 from being forwarded across the bridge and going out ether2. However traffic continues to flow and this rule has no effect.
