---
id: collect-261001-general-networking/general-networking/manual-how-tos-fwcategory-html-9a4a7124
title: "manual-how-tos-fwcategory-html-9a4a7124"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-fwcategory-html-9a4a7124.md
source_anchor: ""
source_lines: [1, 16]
sha256: bc9b8b188c905aad6e3a2289f10543ea392e751a73915a0fa7c7b4fa2923b61f
---

# manual-how-tos-fwcategory-html-9a4a7124

Organize PF Rules by Category
OPNsense firewall rules can be organized per category. These categories can be freely chosen or selected.
Note
This feature was added in version 16.1.1. Always keep your system up to date.
Adding a category to a rule
To add a category to a rule, open or create a new rule and scroll to Category. Then just add you category, if this is the first rule with a category no selection options will be visible.
Firewall Rules Filter by category
Only when there are rules with a defined category, the Filter by category becomes visible at the bottom of the table.
If you click it is will look like this:
If you have a large number of categories, then just start typing and in search box to make a quick selection.
And after selection
Now when selecting our test category it will look like this:
That is all there is to it to organize your rules without messing anything up.
Multi Select
In a later release of OPNsense 16.1 multi selection has been added. This features makes it possible to select rules from more than one category.
Example:
