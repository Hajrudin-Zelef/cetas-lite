---
id: collect-260926-mikrotik/mikrotik/questions-1735854-mikrotik-firewall-rule-to-drop-everything-11011080
title: "questions-1735854-mikrotik-firewall-rule-to-drop-everything-11011080"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-1735854-mikrotik-firewall-rule-to-drop-everything-11011080.md
source_anchor: ""
source_lines: [1, 4]
sha256: 3e922672941b18d36197c25d23c8de1f07f63b807c6bc442c8e0245d585a8dbc
---

# questions-1735854-mikrotik-firewall-rule-to-drop-everything-11011080

I would like to make rules that will allow traffic only from certain mac addresses, and drop everything else. Do you have any ideas? I have tried /ip firewall filter add action=drop chain=forward But it did not work.
1 Answer 1
Seems you're looking at the wrong tool. The IP firewall does (by default) not do anything based on MAC addresses - or any filtering at all if source and target are in the same subnet, as routing does not apply in this case. If you want to set up such rules, you need to use the Bridge Filter feature of RouterOS. Be aware that Layer2/Bridge Filters require quite a bit of processing power.
/ip firewall filter print where chain=forward?Flags: X - disabled, I - invalid, D - dynamic 0 D ;;; special dummy rule to show fasttrack counters chain=forward action=passthrough 1 ;;; TEst pc chain=forward action=accept src-mac-address=A4:4C:C8:6A:F2:F5 log=no log-prefix="" 2 ;;; Drop all Traffic chain=forward action=drop log=no log-prefix="". The second rule (1) is my test pc for mac filtering.
