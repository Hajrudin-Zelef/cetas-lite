---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-889889-ubiquity-unifi-wont-accept-login-after-migration-5443e6f7
title: "questions-889889-ubiquity-unifi-wont-accept-login-after-migration-5443e6f7"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-889889-ubiquity-unifi-wont-accept-login-after-migration-5443e6f7.md
source_anchor: ""
source_lines: [1, 12]
sha256: 17ef1f9a1074c852191740ffe3e07ca20a4faa6dca019fef7132196b26397c2b
---

# questions-889889-ubiquity-unifi-wont-accept-login-after-migration-5443e6f7

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
4
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've migrated a existing Unifi Controller to another server. When the configuration Wizard appeard, I choose "Restore from Backup" after the Import the login screen appears, but the Login isn't working.
I tried these versions
5.5.24 -> 5.5.24 and
5.5.24 -> 5.6.26
The issue was I made the wrong backup. I made the Backup at Settings-> Site -> Export Site which seems to be insufficient. I read that on a Migration how-to.
Settings -> Maintenance -> Backup is the right backup to take.
