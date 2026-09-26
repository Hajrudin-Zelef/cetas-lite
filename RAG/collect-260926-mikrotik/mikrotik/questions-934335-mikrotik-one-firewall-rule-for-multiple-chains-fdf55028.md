---
id: collect-260926-mikrotik/mikrotik/questions-934335-mikrotik-one-firewall-rule-for-multiple-chains-fdf55028
title: "Mikrotik: One firewall rule for multiple chains?"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-934335-mikrotik-one-firewall-rule-for-multiple-chains-fdf55028.md
source_anchor: ""
source_lines: [1, 13]
sha256: 332a5b3bf1805a06d5a6e9c8ada2f7966f5dd54efdfaba1a0ea85dad5257c473
---

# Mikrotik: One firewall rule for multiple chains?

*Source : https://superuser.com/questions/934335/mikrotik-one-firewall-rule-for-multiple-chains | Site : superuser.com | Score : 1*

I created multiple firewall rules for input chain as described in an article and now it seems that I need to create the same set of rules for forward chain. Could I make a rule for multiple chains?

---

## Reponse — score 1

You can put the 'common' rules of both input and forward chains into a new chain and then from the input and forward chain you only need to jump rules to the newly created chain. ie: 1 rule per chain.

This way you maintain a single set of rules for both input and forward chains.
