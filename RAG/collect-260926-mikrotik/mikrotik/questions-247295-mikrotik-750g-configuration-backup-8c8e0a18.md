---
id: collect-260926-mikrotik/mikrotik/questions-247295-mikrotik-750g-configuration-backup-8c8e0a18
title: "questions-247295-mikrotik-750g-configuration-backup-8c8e0a18"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-247295-mikrotik-750g-configuration-backup-8c8e0a18.md
source_anchor: ""
source_lines: [1, 9]
sha256: 2fa2174bcd0e8ada6bf467bd1f36084d4ce5edf079604143f5f33eff9144969c
---

# questions-247295-mikrotik-750g-configuration-backup-8c8e0a18

Use script to transfer configuration.
/export file=config.rsc
Transfer script to PC.
Fix script by removing mac-address options from interfaces at begining of script.
Transfer to other Mikrotik
And apply script.
/system reset-configuration no-defaults=yes run-after-reset=config.rsc
PS. For successful applying configuration new router must have enabled all packages which configuration exported.
Export file can be edited to cover some special cases (I know, user-manager configuration don't cleared by /system reset-configuration)
