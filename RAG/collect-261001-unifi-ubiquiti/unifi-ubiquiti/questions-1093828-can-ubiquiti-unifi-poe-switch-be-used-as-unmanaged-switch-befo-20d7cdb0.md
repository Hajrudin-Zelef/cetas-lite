---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1093828-can-ubiquiti-unifi-poe-switch-be-used-as-unmanaged-switch-befo-20d7cdb0
title: "questions-1093828-can-ubiquiti-unifi-poe-switch-be-used-as-unmanaged-switch-befo-20d7cdb0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1093828-can-ubiquiti-unifi-poe-switch-be-used-as-unmanaged-switch-befo-20d7cdb0.md
source_anchor: ""
source_lines: [1, 12]
sha256: 714c6f3b05f0174d587caf9979d53624123e242c951734a69a32ffe27c273859
---

# questions-1093828-can-ubiquiti-unifi-poe-switch-be-used-as-unmanaged-switch-befo-20d7cdb0

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I know some managed switches can just be used as an unmanaged switch out of the box, but I have a site with a UniFi US-16-150W PoE switch which nobody knows anything about how it is/was setup or passwords/controller information.
My question is: Does anyone know if I factory reset the switch can it just be used as an unmanaged PoE switch without adopting and without using the UniFi controller software etc?
I looked through the manual but it is not clear if the switch will operate or not out of the box (before adopting). And I would like to know the answer before I do a factory reset!
I'll answer my own question as I was forced to do a reset of this UniFi.
Yes, it does appear that a UniFi switch (at least this one which is actually a US-8-150W not as I originally posted) can be used as an unmanaged switch out of the box. Both data and PoE are working.
I reset it, it's showing a steady white LED which according to the manual: "Factory defaults, waiting for integration".
