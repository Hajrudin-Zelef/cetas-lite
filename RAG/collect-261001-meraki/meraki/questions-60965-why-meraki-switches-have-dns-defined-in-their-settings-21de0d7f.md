---
id: collect-261001-meraki/meraki/questions-60965-why-meraki-switches-have-dns-defined-in-their-settings-21de0d7f
title: "questions-60965-why-meraki-switches-have-dns-defined-in-their-settings-21de0d7f"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-60965-why-meraki-switches-have-dns-defined-in-their-settings-21de0d7f.md
source_anchor: ""
source_lines: [1, 10]
sha256: b3aa9ba5f94c614c0d8f2093cff10bf0ab537abe53dc9a60cfc4604581befd09
---

# questions-60965-why-meraki-switches-have-dns-defined-in-their-settings-21de0d7f

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
What a Meraki switch would do with DNS?
switching has nothing to do with Domain name resolution..
the only thing I can come up with, is that their switches probably using some type of domain address instead of IP to connect to their clouds???!!
Merakis are managed via a centralized system hosted in the cloud. DNS is the way by which the Meraki endpoints (...whether switches, firewalls, AP's or cameras) connect back to that cloud system. It's also key to said endpoints being able to reach alternate management points should failures or excessive load occur in Meraki's infrastructure.
