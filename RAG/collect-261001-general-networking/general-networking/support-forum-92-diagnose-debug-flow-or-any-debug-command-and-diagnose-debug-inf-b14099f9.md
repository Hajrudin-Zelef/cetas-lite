---
id: collect-261001-general-networking/general-networking/support-forum-92-diagnose-debug-flow-or-any-debug-command-and-diagnose-debug-inf-b14099f9
title: "support-forum-92-diagnose-debug-flow-or-any-debug-command-and-diagnose-debug-inf-b14099f9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/support-forum-92-diagnose-debug-flow-or-any-debug-command-and-diagnose-debug-inf-b14099f9.md
source_anchor: ""
source_lines: [1, 5]
sha256: ef8c2fea1ca4173a167b9810225b425bb2ae8e0c7098af5b102fc9fa05b9a879
---

# support-forum-92-diagnose-debug-flow-or-any-debug-command-and-diagnose-debug-inf-b14099f9

When I enable the various debugs as shown and I run diagnose debug info command I am expecting to see all currently enabled debugs in the location shown but I do not. Is this how it should be?
Where or how can I obtain feedback to confirm what debugs are turned on at any given point?
As you're showing, you filter set is filtering only protocol 6 (TCP) in. The protocol filter takes only one. So your last filter 6/TCP is there. IPsec never uses TCP. Just clear the filter with "diag debug flow filter clear" then specify only address to filter. If this is a spoke only with one IPsec, you don't have to specify even the address. I don't see any point specifying protocol to just debug IKE.
Toshi
Did this topic help you find an answer to your question?
