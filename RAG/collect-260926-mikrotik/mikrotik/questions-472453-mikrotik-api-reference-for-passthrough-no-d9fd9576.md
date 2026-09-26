---
id: collect-260926-mikrotik/mikrotik/questions-472453-mikrotik-api-reference-for-passthrough-no-d9fd9576
title: "questions-472453-mikrotik-api-reference-for-passthrough-no-d9fd9576"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/scripts-scheduler/questions-472453-mikrotik-api-reference-for-passthrough-no-d9fd9576.md
source_anchor: ""
source_lines: [1, 6]
sha256: 9d49efd31300eb60a3c8cad36500912e4f17df3cdaba4cacbca0b7a7901351ba
---

# questions-472453-mikrotik-api-reference-for-passthrough-no-d9fd9576

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
But nowhere have I seen a complete description of what this attribute does! Can someone link to documentation or give a complete description of the behavior of this attribute?
(Please note that I am NOT referring to the similarly named action: action=passthrough.)
In my experience, when you use passthrough=yes, for example in a mangle rule, then the packet if it gets matched by that rule it's processed by the subsequent rules.
Where in case you use passthrough=no, if the packet gets matched by that rule, it will not get processed by the subsequent rules which can be useful to save some CPU as stated at the bottom of the manual page you posted.
There are certain types of configurations that passthrough=yes is needed, but in most cases I believe that it's safe to use passthrough=no unless you need otherwise.
