---
id: collect-260926-mikrotik/mikrotik/questions-21138-mikrotik-api-no-such-command-when-trying-to-get-count-only-of-pp-903379ed
title: "questions-21138-mikrotik-api-no-such-command-when-trying-to-get-count-only-of-pp-903379ed"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/scripts-scheduler/questions-21138-mikrotik-api-no-such-command-when-trying-to-get-count-only-of-pp-903379ed.md
source_anchor: ""
source_lines: [1, 9]
sha256: 2607b52dabdd4f173581303d59f17156da221c24078ad1b21293c086a54a2c66
---

# questions-21138-mikrotik-api-no-such-command-when-trying-to-get-count-only-of-pp-903379ed

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I don't know if I'm posting in the right forum. But well correct me if I am wrong.
I am developing a service that counts the PPPoE users per MikroTik Device in the whole network every X minutes. I'm using the API build in Python3 (Although I think the problem is not language-dependent, but I tell you in order for you to know).
When I connect to a MKT, I can execute the command ppp active print count-only. In the API the language changes a bit. I can't execute that command. It tells me =message=no such command prefix. I've tried with /ppp/active/print/count-only but only works when I execute /ppp/active/print but that is not the result I want.
